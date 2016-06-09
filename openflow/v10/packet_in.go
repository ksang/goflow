package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)

type packetIn struct {
	openflow.Message
	bufferID    uint32
	totalLength uint16
	inPort      uint16
	reason      uint8
	data        []byte
}

func (p *packetIn) BufferID() uint32 {
	return p.bufferID
}

func (p *packetIn) SetBufferID(bid uint32) {
	p.bufferID = bid
}

func (p *packetIn) InPort() uint16 {
	return p.inPort
}

func (p *packetIn) SetInPort(ip uint16) {
	p.inPort = ip
}

func (p *packetIn) Data() []byte {
	return p.data
}

func (p *packetIn) SetData(data []byte) {
	p.data = data
	p.totalLength = uint16(len(p.data))
}

func (p *packetIn) TotalLength() uint16 {
	return p.totalLength
}

func (p *packetIn) TableID() uint8 {
	// OpenFlow 1.0 does not have table ID
	return 0
}

func (p *packetIn) SetTableID(tid uint8) {
	// OpenFlow 1.0 does not have table ID
	return
}

func (p *packetIn) Reason() uint8 {
	return p.reason
}

func (p *packetIn) SetReason(r uint8) {
	p.reason = r
}

func (p *packetIn) Cookie() uint64 {
	// OpenFlow 1.0 does not have cookie
	return 0
}

func (p *packetIn) SetCookie(c uint64) {
	// OpenFlow 1.0 does not have cookie
	return
}

func (p *packetIn) MarshalBinary() ([]byte, error) {
	if int(p.totalLength) != len(p.data) {
		return nil, openflow.ErrInvalidDataLength
	}
	v := make([]byte, p.totalLength+10)
	binary.BigEndian.PutUint32(v[0:4], p.bufferID)
	binary.BigEndian.PutUint16(v[4:6], p.totalLength)
	binary.BigEndian.PutUint16(v[6:8], p.inPort)
	v[8] = p.reason
	//v[9] is padding
	if p.totalLength > 0 {
		copy(v[10:], p.data)
	}
	p.SetPayload(v)
	return p.Message.MarshalBinary()
}

func (p *packetIn) UnmarshalBinary(data []byte) error {
	if err := p.Message.UnmarshalBinary(data); err != nil {
		return err
	}

	payload := p.Payload()
	if payload == nil || len(payload) < 10 {
		return openflow.ErrInvalidPacketLength
	}
	p.bufferID = binary.BigEndian.Uint32(payload[0:4])
	p.totalLength = binary.BigEndian.Uint16(payload[4:6])
	p.inPort = binary.BigEndian.Uint16(payload[6:8])
	p.reason = payload[8]
	// payload[9] is padding
	if len(payload) >= 10 {
		p.data = payload[10:]
	}
	if int(p.totalLength) != len(p.data) {
		return openflow.ErrInvalidDataLength
	}

	return nil
}

func NewPacketIn(xid uint32) openflow.PacketIn {
	return &packetIn{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_PACKET_IN, xid),
	}
}
