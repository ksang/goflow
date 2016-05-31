package openflow

import (
	"encoding/binary"
)

type MessageDecoder interface {
	HeaderDecoder
	Payload() []byte
	SetPayload(payload []byte)
}

type Message struct {
	header  Header
	payload []byte
}

// Create a message with no payload.
func NewMessage(version uint8, msgType uint8, xid uint32) Message {
	return Message{
		header: Header{
			version: version,
			msgType: msgType,
			length:  OF_HEADER_SIZE,
			xid:     xid,
		},
	}
}

// Member functions to implement HeaderDecoder interface
func (m *Message) Version() uint8 {
	return m.header.version
}

func (m *Message) MsgType() uint8 {
	return m.header.msgType
}

func (m *Message) Length() uint16 {
	return m.header.length
}

func (m *Message) TransactionID() uint32 {
	return m.header.xid
}

func (m *Message) SetTransactionID(xid uint32) {
	m.header.xid = xid
}

// Member functions to implement MessageDecoder interface
func (m *Message) Payload() []byte {
	return m.payload
}

func (m *Message) SetPayload(payload []byte) {
	if payload == nil {
		return
	}
	m.payload = payload
	m.header.length = uint16(OF_HEADER_SIZE + len(payload))
}

// convert Message to binary data
func (m *Message) MarshalBinary() ([]byte, error) {
	var length uint16 = OF_HEADER_SIZE
	if m.payload != nil {
		length += uint16(len(m.payload))
	}

	v := make([]byte, length)
	v[0] = m.header.version
	v[1] = m.header.msgType
	binary.BigEndian.PutUint16(v[2:4], length)
	binary.BigEndian.PutUint32(v[4:8], m.header.xid)
	if length > OF_HEADER_SIZE {
		copy(v[8:], m.payload)
	}

	return v, nil
}

// convert binary data to message
func (m *Message) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) < OF_HEADER_SIZE {
		return ErrInvalidPacketLength
	}

	m.header.version = data[0]
	m.header.msgType = data[1]

	m.header.length = binary.BigEndian.Uint16(data[2:4])
	if m.header.length < OF_HEADER_SIZE || len(data) < int(m.header.length) {
		return ErrInvalidPacketLength
	}
	m.header.xid = binary.BigEndian.Uint32(data[4:8])
	m.payload = data[8:m.header.length]

	return nil
}
