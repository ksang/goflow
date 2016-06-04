package openflow

import (
	"encoding"
)

// general action interface
type Action interface {
	Type() uint16
	SetType(uint16)
	Length() uint16
	Payload() []byte
	SetPayload([]byte) error
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type ActionHead interface {
	Type() uint16
	SetType(uint16)
	Length() uint16
	Payload() []byte
	SetPayload([]byte) error
	MarshalActionHead() ([]byte, error)
	UnmarshalActionHead([]byte) error
}

// Echo message interface
type Echo interface {
	MessageDecoder
	Data() []byte
	SetData(data []byte) error
	Error() error
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// Packet in message interface
type PacketIn interface {
	MessageDecoder
	BufferID() uint32
	SetBufferID(uint32)
	TotalLength() uint16
	InPort() uint16
	SetInPort(uint16)
	TableID() uint8
	SetTableID(uint8)
	Reason() uint8
	SetReason(uint8)
	Cookie() uint64
	SetCookie(uint64)
	Data() []byte
	SetData([]byte)
	encoding.BinaryUnmarshaler
	encoding.BinaryMarshaler
}

// Packet in message interface
type PacketOut interface {
	MessageDecoder
	BufferID() uint32
	SetBufferID(uint32)
	ActionsLength() uint16
	InPort() uint16
	SetInPort(uint16) error
	Data() []byte
	SetData([]byte)
	Action() []Action
	AddAction(Action)
	encoding.BinaryUnmarshaler
	encoding.BinaryMarshaler
}
