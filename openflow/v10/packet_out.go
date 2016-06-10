package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)

type packetOut struct {
	openflow.Message
	bufferID      uint32
	inPort        uint16
	actionsLength uint16
	action        []openflow.Action
	data          []byte
}

func (p *packetOut) BufferID() uint32 {
	return p.bufferID
}

func (p *packetOut) SetBufferID(bid uint32) {
	p.bufferID = bid
}

func (p *packetOut) InPort() uint16 {
	return p.inPort
}

func (p *packetOut) SetInPort(ip uint16) error {
	if ip > 0xff00 {
		return openflow.ErrInvalidValueProvided
	}
	p.inPort = ip
	return nil
}

func (p *packetOut) ActionsLength() uint16 {
	return p.actionsLength
}

func (p *packetOut) Data() []byte {
	return p.data
}

func (p *packetOut) SetData(data []byte) {
	p.data = data
}

func (p *packetOut) Action() []openflow.Action {
	return p.action
}

func (p *packetOut) AddAction(action openflow.Action) {
	p.action = append(p.action, action)
	p.actionsLength += action.Length()
}

func (p *packetOut) MarshalBinary() ([]byte, error) {
	actionsLen := 0
	for _, act := range p.action {
		actionsLen += int(act.Length())
	}
	// actionsLength not matching sum of action's length field
	if int(p.actionsLength) != actionsLen {
		return nil, openflow.ErrInvalidDataLength
	}
	// packetOutLen = bufferID(4 bytes) + inPort(2 bytes) +
	//			actionsLength(2 bytes) + actionsLen + dataLen
	packetOutLen := 8 + int(actionsLen) + len(p.data)
	v := make([]byte, packetOutLen)
	binary.BigEndian.PutUint32(v[0:4], p.bufferID)
	binary.BigEndian.PutUint16(v[4:6], p.inPort)
	binary.BigEndian.PutUint16(v[6:8], p.actionsLength)
	// Marshal actions
	lastPosition := 8
	for _, act := range p.action {
		actLen := int(act.Length())
		actPayload, err := act.MarshalBinary()
		if err != nil {
			return nil, err
		}
		copy(v[lastPosition:lastPosition+actLen], actPayload)
		lastPosition += actLen
	}
	if len(p.data) > 0 {
		copy(v[lastPosition:], p.data)
	}
	// Call Message's SetPayload method
	p.SetPayload(v)
	return p.Message.MarshalBinary()
}

func (p *packetOut) UnmarshalBinary(data []byte) error {
	if err := p.Message.UnmarshalBinary(data); err != nil {
		return err
	}

	payload := p.Payload()
	if payload == nil || len(payload) < 8 {
		return openflow.ErrInvalidPacketLength
	}
	p.bufferID = binary.BigEndian.Uint32(payload[0:4])
	if err := p.SetInPort(binary.BigEndian.Uint16(payload[4:6])); err != nil {
		return err
	}
	p.actionsLength = binary.BigEndian.Uint16(payload[6:8])
	actLen := int(p.actionsLength)
	// packet size smaller than actions length
	if actLen+8 > len(payload) {
		return openflow.ErrInvalidDataLength
	}
	for i := 8; i < actLen+8; {
		act := NewActionHeader()
		if err := act.UnmarshalBinary(payload[i:]); err != nil {
			return err
		}
		i += int(act.Length())
		p.action = append(p.action, act)
	}
	// has data
	if len(payload) > actLen+8 {
		p.data = payload[actLen+8:]
	}
	return nil
}

func NewPacketOut(xid uint32) openflow.PacketOut {
	return &packetOut{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_PACKET_OUT, xid),
	}
}
