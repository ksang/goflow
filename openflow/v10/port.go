package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
	"net"
)

type port struct {
	portID     uint16
	hwAddr     net.HardwareAddr
	name       string
	config     uint32
	state      uint32
	curr       uint32
	advertised uint32
	supported  uint32
	peer       uint32
}

func (p *port) PortID() uint16 {
	return p.portID
}

func (p *port) SetPortID(pid uint16) {
	p.portID = pid
}

func (p *port) HWAddr() net.HardwareAddr {
	return p.hwAddr
}

func (p *port) SetHWAddr(mac net.HardwareAddr) error {
	if len(mac) != 6 {
		return openflow.ErrInvalidValueProvided
	}
	p.hwAddr = mac
	return nil
}

func (p *port) Name() string {
	return p.name
}

func (p *port) SetName(name string) error {
	if len(name) > 16 {
		return openflow.ErrInvalidValueProvided
	}
	p.name = name
	return nil
}

func (p *port) Config() uint32 {
	return p.config
}

func (p *port) SetConfig(config uint32) {
	p.config = config
}

func (p *port) State() uint32 {
	return p.state
}

func (p *port) SetState(state uint32) {
	p.state = state
}

func (p *port) Curr() uint32 {
	return p.curr
}

func (p *port) SetCurr(curr uint32) {
	p.curr = curr
}

func (p *port) Advertised() uint32 {
	return p.advertised
}

func (p *port) SetAdvertised(value uint32) {
	p.advertised = value
}

func (p *port) Supported() uint32 {
	return p.supported
}

func (p *port) SetSupported(value uint32) {
	p.supported = value
}

func (p *port) Peer() uint32 {
	return p.supported
}

func (p *port) SetPeer(value uint32) {
	p.peer = value
}

func (p *port) MarshalBinary() ([]byte, error) {
	v := make([]byte, 48)
	binary.BigEndian.PutUint16(v[0:2], p.portID)
	copy(v[2:8], p.hwAddr)
	copy(v[8:24], []byte(p.name))
	binary.BigEndian.PutUint32(v[24:28], p.config)
	binary.BigEndian.PutUint32(v[28:32], p.state)
	binary.BigEndian.PutUint32(v[32:36], p.curr)
	binary.BigEndian.PutUint32(v[36:40], p.advertised)
	binary.BigEndian.PutUint32(v[40:44], p.supported)
	binary.BigEndian.PutUint32(v[44:48], p.peer)
	return v, nil
}

func (p *port) UnmarshalBinary(data []byte) error {
	if len(data) != 48 {
		return openflow.ErrInvalidDataLength
	}
	p.portID = binary.BigEndian.Uint16(data[0:2])
	p.hwAddr = data[2:8]
	p.name = string(data[8:24])
	p.config = binary.BigEndian.Uint32(data[24:28])
	p.state = binary.BigEndian.Uint32(data[28:32])
	p.curr = binary.BigEndian.Uint32(data[32:36])
	p.advertised = binary.BigEndian.Uint32(data[36:40])
	p.supported = binary.BigEndian.Uint32(data[40:44])
	p.peer = binary.BigEndian.Uint32(data[44:48])
	return nil
}

func NewPort(pid uint16, hwAddr net.HardwareAddr, name string) (openflow.Port, error) {
	p := &port{}
	p.SetPortID(pid)
	if err := p.SetHWAddr(hwAddr); err != nil {
		return nil, err
	}
	if err := p.SetName(name); err != nil {
		return nil, err
	}
	return p, nil
}

func NewEmptyPort() openflow.Port {
	return &port{}
}
