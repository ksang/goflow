package openflow

import (
	"encoding"
	"net"
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

// Port is a structure describes a port
type Port interface {
	PortID() PortID
	SetPortID(PortID)
	HWAddr() net.HardwareAddr
	SetHWAddr(net.HardwareAddr) error
	Name() string
	SetName(string) error
	Config() PortConfig
	SetConfig(PortConfig)
	State() PortState
	SetState(PortState)
	Curr() PortFeature
	SetCurr(PortFeature)
	Advertised() PortFeature
	SetAdvertised(PortFeature)
	Supported() PortFeature
	SetSupported(PortFeature)
	Peer() PortFeature
	SetPeer(PortFeature)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// Match is a structure used in flow related messages
type Match interface {
	Wildcards() uint32
	// returns wildcard mask for Inport and inport number
	InPort() (bool, uint16)
	SetInPort(uint16)
	SetWildcardInPort()

	DLSrc() (bool, net.HardwareAddr)
	SetDLSrc(net.HardwareAddr)
	SetWildcardDLSrc()
	DLDst() (bool, net.HardwareAddr)
	SetDLDst(net.HardwareAddr)
	SetWildcardDLDst()

	DLVlan() (bool, uint16)
	SetDLVlan(uint16) error
	SetWildcardDLVlan()

	DLPCP() (bool, uint8)
	SetDLPCP(uint8)
	SetWildcardDLVlanPCP()

	DLType() (bool, uint16)
	SetDLType(uint16) error
	SetWildcardDLType()

	NWTos() (bool, uint8)
	SetNWTos(uint8)
	SetWildcardNWTos()

	NWProto() (bool, uint8)
	SetNWProto(uint8) error
	SetWildcardNWProto()
	// The wildcard mask here is to wildcard src/dst ip address
	// Set it to 32 is matching the full ip address
	// Set it to 0  will mask all the ip addr, returns 0.0.0.0
	NWSrc() net.IP
	SetNWSrc(net.IP)
	SetWildcardNWSrc(int)
	NWDst() net.IP
	SetNWDst(net.IP)
	SetWildcardNWDst(int)
	// Source and destination port
	TPSrc() (bool, uint16)
	SetTPSrc(uint16)
	SetWildcardTPSrc()
	TPDst() (bool, uint16)
	SetTPDst(uint16)
	SetWildcardTPDst()

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
	Capabilities() FeatureCapability
	SetCapabilities(FeatureCapability)
	Actions() FeatureAction
	SetActions(FeatureAction)
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

type FlowRemoved interface {
	MessageDecoder
	Match() Match
	SetMatch(Match)
	Cookie() uint64
	SetCookie(uint64)
	Priority() uint16
	SetPriority(uint16)
	Reason() uint8
	SetReason(uint8) error
	DurationSec() uint32
	SetDurationSec(uint32)
	DurationNanoSec() uint32
	SetDurationNanoSec(uint32)
	IdleTimeout() uint16
	SetIdleTimeout(uint16)
	PacketCount() uint64
	SetPacketCount(uint64)
	ByteCount() uint64
	SetByteCount(uint64)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type PortStatus interface {
	MessageDecoder
	Reason() PortReason
	SetReason(PortReason)
	Port() Port
	SetPort(Port)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type FlowMod interface {
	MessageDecoder
	Match() Match
	SetMatch(Match)
	Cookie() uint64
	SetCookie(uint64) error
	Command() FlowCommand
	SetCommand(FlowCommand)
	IdleTimeout() uint16
	SetIdleTimeout(uint16)
	HardTimeout() uint16
	SetHardTimeout(uint16)
	Priority() uint16
	SetPriority(uint16)
	BufferID() uint32
	SetBufferID(uint32)
	OutPort() uint16
	SetOutPort(uint16)
	Flags() FlowFlag
	SetFlags(FlowFlag)
	Action() Action
	SetAction(Action)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type PortMod interface {
	MessageDecoder
	Port() PortID
	SetPort(PortID)
	HWAddr() net.HardwareAddr
	SetHWAddr(net.HardwareAddr)
	Config() PortConfig
	SetConfig(PortConfig)
	Mask() uint32
	SetMask(uint32) error
	Advertise() PortFeature
	SetAdvertise(PortFeature)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}