package v10

import (
	"github.com/ksang/goflow/openflow"
)

type barrier struct {
	openflow.Message
}

func (b *barrier) MarshalBinary() ([]byte, error) {
	return b.Message.MarshalBinary()
}

func (b *barrier) UnmarshalBinary(data []byte) error {
	return b.Message.UnmarshalBinary(data)
}

func NewBarrierRequest(xid uint32) openflow.BarrierRequest {
	return &barrier{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_BARRIER_REQUEST, xid),
	}
}

func NewBarrierReply(xid uint32) openflow.BarrierReply {
	return &barrier{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_BARRIER_REPLY, xid),
	}
}