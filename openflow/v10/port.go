package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
	"net"
)

type port struct {
	portID     openflow.PortID
	hwAddr     net.HardwareAddr
	name       string
	config     openflow.PortConfig
	state      openflow.PortState
	curr       openflow.PortFeature
	advertised openflow.PortFeature
	supported  openflow.PortFeature
	peer       openflow.PortFeature
}

func (p *port) PortID() openflow.PortID {
	return p.portID
}

func (p *port) SetPortID(pid openflow.PortID) {
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

func (p *port) Config() openflow.PortConfig {
	return p.config
}

func (p *port) SetConfig(config openflow.PortConfig) {
	p.config = config
}

func (p *port) State() openflow.PortState {
	return p.state
}

func (p *port) SetState(state openflow.PortState) {
	p.state = state
}

func (p *port) Curr() openflow.PortFeature {
	return p.curr
}

func (p *port) SetCurr(curr openflow.PortFeature) {
	p.curr = curr
}

func (p *port) Advertised() openflow.PortFeature {
	return p.advertised
}

func (p *port) SetAdvertised(value openflow.PortFeature) {
	p.advertised = value
}

func (p *port) Supported() openflow.PortFeature {
	return p.supported
}

func (p *port) SetSupported(value openflow.PortFeature) {
	p.supported = value
}

func (p *port) Peer() openflow.PortFeature {
	return p.supported
}

func (p *port) SetPeer(value openflow.PortFeature) {
	p.peer = value
}

func (p *port) MarshalBinary() ([]byte, error) {
	v := make([]byte, 48)
	binary.BigEndian.PutUint16(v[0:2], uint16(p.portID))
	copy(v[2:8], p.hwAddr)
	copy(v[8:24], []byte(p.name))
	binary.BigEndian.PutUint32(v[24:28], uint32(p.config))
	binary.BigEndian.PutUint32(v[28:32], uint32(p.state))
	binary.BigEndian.PutUint32(v[32:36], uint32(p.curr))
	binary.BigEndian.PutUint32(v[36:40], uint32(p.advertised))
	binary.BigEndian.PutUint32(v[40:44], uint32(p.supported))
	binary.BigEndian.PutUint32(v[44:48], uint32(p.peer))
	return v, nil
}

func (p *port) UnmarshalBinary(data []byte) error {
	if len(data) != 48 {
		return openflow.ErrInvalidDataLength
	}
	p.portID = openflow.PortID(binary.BigEndian.Uint16(data[0:2]))
	p.hwAddr = data[2:8]
	p.name = string(data[8:24])
	p.config = openflow.PortConfig(binary.BigEndian.Uint32(data[24:28]))
	p.state = openflow.PortState(binary.BigEndian.Uint32(data[28:32]))
	p.curr = openflow.PortFeature(binary.BigEndian.Uint32(data[32:36]))
	p.advertised = openflow.PortFeature(binary.BigEndian.Uint32(data[36:40]))
	p.supported = openflow.PortFeature(binary.BigEndian.Uint32(data[40:44]))
	p.peer = openflow.PortFeature(binary.BigEndian.Uint32(data[44:48]))
	return nil
}

func NewPort(pid openflow.PortID, hwAddr net.HardwareAddr, name string) (openflow.Port, error) {
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
