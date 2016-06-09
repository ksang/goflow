package v10

import (
	"github.com/ksang/goflow/openflow"
)

type portStatus struct {
	openflow.Message
	reason openflow.PortReason
	port   openflow.Port
}

func (p *portStatus) Reason() openflow.PortReason {
	return p.reason
}

func (p *portStatus) SetReason(r openflow.PortReason) {
	p.reason = r
}

func (p *portStatus) Port() openflow.Port {
	return p.port
}

func (p *portStatus) SetPort(port openflow.Port) {
	p.port = port
}

func (p *portStatus) MarshalBinary() ([]byte, error) {
	v := make([]byte, 56)
	v[0] = uint8(p.reason)
	//v[1:8] is pad
	port, err := p.port.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(v[8:], port)
	p.SetPayload(v)
	return p.Message.MarshalBinary()
}

func (p *portStatus) UnmarshalBinary(data []byte) error {
	if err := p.Message.UnmarshalBinary(data); err != nil {
		return err
	}

	payload := p.Payload()
	if payload == nil || len(payload) != 56 {
		return openflow.ErrInvalidPacketLength
	}
	p.reason = openflow.PortReason(payload[0])
	// payload[1:8] is padding
	p.port = new(port)
	if err := p.port.UnmarshalBinary(payload[8:]); err != nil {
		return err
	}

	return nil
}

func NewPortStatus(xid uint32) openflow.PortStatus {
	return &portStatus{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_PORT_STATUS, xid),
	}
}
