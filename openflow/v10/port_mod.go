package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
	"net"
)

type portMod struct {
	openflow.Message
	port openflow.PortID
	hwAddr net.HardwareAddr
	config openflow.PortConfig
	mask uint32
	advertise openflow.PortFeature
}

func (p *portMod) Port() openflow.PortID {
	return p.port
}

func (p *portMod) SetPort(pid openflow.PortID){
	p.port = pid
}

func (p *portMod) HWAddr() net.HardwareAddr {
	return p.hwAddr
}

func (p *portMod) SetHWAddr(hwa net.HardwareAddr) {
	p.hwAddr = hwa
}

func (p *portMod) Config() openflow.PortConfig {
	return p.config
}
	
func (p *portMod) SetConfig(pc openflow.PortConfig) {
	p.config = pc
}

func (p *portMod) Mask() uint32 {
	return p.mask
}

func (p *portMod) SetMask(m uint32) error {
	if m > 0x7f {
		return openflow.ErrInvalidValueProvided
	}
	p.mask = m
	return nil
}

func (p *portMod) Advertise() openflow.PortFeature {
	return p. advertise
}

func (p *portMod) SetAdvertise(pf openflow.PortFeature) {
	p.advertise = pf
}

func (p *portMod) MarshalBinary() ([]byte, error) {
	v := make([]byte, 24)
	binary.BigEndian.PutUint16(v[0:2], uint16(p.port))
	copy(v[2:8], p.hwAddr)
	binary.BigEndian.PutUint32(v[8:12], uint32(p.config))
	binary.BigEndian.PutUint32(v[12:16], p.mask)
	binary.BigEndian.PutUint32(v[16:20], uint32(p.advertise))
	//v[20:24] is pad
	p.SetPayload(v)
	return p.Message.MarshalBinary()
}

func (p *portMod) UnmarshalBinary(data []byte) error {
	if err := p.Message.UnmarshalBinary(data); err != nil {
		return err
	}

	payload := p.Payload()
	if payload == nil || len(payload) != 24 {
		return openflow.ErrInvalidPacketLength
	}
	p.port = openflow.PortID(binary.BigEndian.Uint16(payload[0:2]))
	p.hwAddr = payload[2:8]
	p.config = openflow.PortConfig(binary.BigEndian.Uint32(payload[8:12]))
	p.mask = binary.BigEndian.Uint32(payload[12:16])
	p.advertise = openflow.PortFeature(binary.BigEndian.Uint32(payload[16:20]))
	// payload[20:24] is padding
	return nil
}

func NewPortMod(xid uint32) openflow.PortMod {
	return &portMod{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_PORT_MOD, xid),
	}
}
