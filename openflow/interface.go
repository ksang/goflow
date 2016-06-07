package openflow

import (
	"net"
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

type Port interface {
	PortID() uint16
	SetPortID(uint16)
	HWAddr() net.HardwareAddr
	SetHWAddr(net.HardwareAddr) error
	Name() string
	SetName(string) error
	Config() uint32
	SetConfig(uint32)
	State() uint32
	SetState(uint32)
	Curr() uint32
	SetCurr(uint32)
	Advertised() uint32
	SetAdvertised(uint32)
	Supported() uint32
	SetSupported(uint32)
	Peer() uint32
	SetPeer(uint32)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
// Hello message interface
// Share the same interface and struct as echo
type Hello interface {
	Echo
}

// Echo message interface (request/reply)
type Echo interface {
	MessageDecoder
	Data() []byte
	SetData(data []byte) error
	Error() error
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// Error message interface
type Error interface {
	MessageDecoder
	Type() uint16
	SetType(uint16) error
	Code() uint16
	SetCode(uint16) error
	Data() []byte
	SetData(data []byte) error
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type Vendor interface {
	MessageDecoder
	VendorID() uint32
	SetVendorID(uint32)
	Data() []byte
	SetData(data []byte) error
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type FeatureRequest interface {
	MessageDecoder
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type FeatureReply interface {
	MessageDecoder
	DPID() uint64
	SetDPID(uint64)
	NumBuffers() uint32
	SetNumBuffers(uint32)
	NumTables() uint8
	SetNumTables(uint8)
	Capabilities() uint32
	SetCapabilities(uint32)
	Actions() uint32
	SetActions(uint32)
	Ports() []Port
	AddPort(Port)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type Config interface {
	Flags() uint16
	SetFlags(uint16)
	MissSendLength() uint16
	SetMissSendLength(uint16)
}

type SetConfig interface {
	MessageDecoder
	Config
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type GetConfigRequest interface {
	MessageDecoder
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type GetConfigReply interface {
	MessageDecoder
	Config
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
