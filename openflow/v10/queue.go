package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)


type queue struct {
	queueID uint32
	length uint16
	rate uint16
}

func (q *queue) QueueID() uint32 {
	return q.queueID
}

func (q *queue) SetQueueID (qid uint32) {
	q.queueID = qid
}

func (q *queue) Length() uint16 {
	return q.length
}

func (q *queue) Rate() uint16 {
	return q.rate
}
	
func (q *queue) SetRate(r uint16) {
	q.rate = r
}

func (q *queue) MarshalBinary() ([]byte, error) {
	v := make([]byte, 24)
	binary.BigEndian.PutUint32(v[0:4], q.queueID)
	binary.BigEndian.PutUint16(v[4:6], q.length)
	// v[6:8] is pad
	binary.BigEndian.PutUint16(v[8:10], uint16(0x0001))
	binary.BigEndian.PutUint16(v[10:12], uint16(16))
	binary.BigEndian.PutUint16(v[16:18], q.rate)
	// v[18:24] is pad
	return v, nil
}

func (q *queue) UnmarshalBinary(data []byte) error {
	if len(data) != 24 {
		return openflow.ErrInvalidPacketLength
	}
	q.queueID = binary.BigEndian.Uint32(data[0:4])
	q.length = binary.BigEndian.Uint16(data[4:6])
	// data[6:8] is pad
	property := binary.BigEndian.Uint16(data[8:10])
	if property != 0x01 {
		// Unknown property
		return nil
	}
	q.rate = binary.BigEndian.Uint16(data[16:18])
	return nil
}

func NewQueue() openflow.Queue {
	return &queue{
			length : 0x18,
	}
}

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

type queueGetConfigReply struct {
	openflow.Message
	port uint16
	queue []openflow.Queue
}

func (q *queueGetConfigReply) Port() uint16 {
	return q.port
}

func (q *queueGetConfigReply) SetPort(p uint16) {
	q.port = p
}

func (q *queueGetConfigReply) Queue() []openflow.Queue {
	return q.queue
}

func (q *queueGetConfigReply) AddQueue(nq openflow.Queue)  {
	q.queue = append(q.queue, nq)
}

func (q *queueGetConfigReply) MarshalBinary() ([]byte, error) {
	v := make([]byte, 8 + 24*len(q.queue))
	binary.BigEndian.PutUint16(v[0:2], q.port)
	// v[2:8] is pad
	for i, nq := range q.queue {
		data, err := nq.MarshalBinary()
		if err != nil {
			return nil, err
		}
		copy(v[8+i*24:8+(i+1)*24], data)
	}
	q.SetPayload(v)
	return q.Message.MarshalBinary()
}

func (q *queueGetConfigReply) UnmarshalBinary(data []byte) error {
	if err := q.Message.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := q.Payload()
	if payload == nil || len(payload) < 8 {
		return openflow.ErrInvalidPacketLength
	}
	q.port = binary.BigEndian.Uint16(payload[0:2])
	if (len(payload) - 8) % 24 != 0 {
		return openflow.ErrInvalidPacketLength
	}
	// Unmarshal Queues
	for i := 8; i < len(payload); i += 24 {
		nq := NewQueue()
		if err := nq.UnmarshalBinary(payload[i:i+24]); err != nil {
			return err
		}
		q.queue = append(q.queue, nq)
	}

	return nil
}

func NewQueueGetConfigReply(xid uint32)	openflow.QueueGetConfigReply {
	return &queueGetConfigReply{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_QUEUE_GET_CONFIG_REPLY, xid),
	}
}