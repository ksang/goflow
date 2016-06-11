package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)

type queueGetConfigRequest struct {
	openflow.Message
	port uint16
}

func (q *queueGetConfigRequest) Port() uint16 {
	return q.port
}

func (q *queueGetConfigRequest) SetPort(p uint16) error {
	if p > 0xff00 {
		return openflow.ErrInvalidValueProvided
	}
	q.port = p
	return nil
}

func (q *queueGetConfigRequest) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4)
	binary.BigEndian.PutUint16(v[0:2], q.port)
	// v[2:4] is pad
	q.SetPayload(v)
	return q.Message.MarshalBinary()
}

func (q *queueGetConfigRequest) UnmarshalBinary(data []byte) error {
	if err := q.Message.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := q.Payload()
	q.port = binary.BigEndian.Uint16(payload[0:2])
	// payload[2:4] is pad
	return nil
}

func NewQueueGetConfigRequest(xid uint32) openflow.QueueGetConfigRequest {
	return &queueGetConfigRequest{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_QUEUE_GET_CONFIG_REQUEST, xid),
	}
}