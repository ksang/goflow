package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)

type flowMod struct {
	openflow.Message
	match       openflow.Match
	cookie      uint64
	command     openflow.FlowCommand
	idleTimeout uint16
	hardTimeout uint16
	priority    uint16
	bufferID    uint32
	outPort     uint16
	flags       openflow.FlowFlag
	action      openflow.Action
}

func (f *flowMod) Match() openflow.Match {
	return f.match
}

func (f *flowMod) SetMatch(m openflow.Match) {
	f.match = m
}

func (f *flowMod) Cookie() uint64 {
	return f.cookie
}

func (f *flowMod) SetCookie(cookie uint64) error {
	if cookie == 0xFFFFFFFFFFFFFFFF {
		return openflow.ErrInvalidValueProvided
	}
	f.cookie = cookie
	return nil
}

func (f *flowMod) Command() openflow.FlowCommand {
	return f.command
}

func (f *flowMod) SetCommand(fc openflow.FlowCommand) {
	f.command = fc
}

func (f *flowMod) IdleTimeout() uint16 {
	return f.idleTimeout
}

func (f *flowMod) SetIdleTimeout(it uint16) {
	f.idleTimeout = it
}

func (f *flowMod) HardTimeout() uint16 {
	return f.hardTimeout
}

func (f *flowMod) SetHardTimeout(ht uint16) {
	f.hardTimeout = ht
}

func (f *flowMod) Priority() uint16 {
	return f.priority
}

func (f *flowMod) SetPriority(pri uint16) {
	f.priority = pri
}

func (f *flowMod) BufferID() uint32 {
	return f.bufferID
}

func (f *flowMod) SetBufferID(bid uint32) {
	f.bufferID = bid
}

func (f *flowMod) OutPort() uint16 {
	return f.outPort
}

func (f *flowMod) SetOutPort(op uint16) {
	f.outPort = op
}

func (f *flowMod) Flags() openflow.FlowFlag {
	return f.flags
}

func (f *flowMod) SetFlags(ff openflow.FlowFlag) {
	f.flags = ff
}

func (f *flowMod) Action() openflow.Action {
	return f.action
}

func (f *flowMod) SetAction(a openflow.Action) {
	f.action = a
}

func (f *flowMod) MarshalBinary() ([]byte, error) {
	m, err := f.match.MarshalBinary()
	a, err := f.action.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v := make([]byte, 64+int(f.action.Length()))
	copy(v[0:40], m)
	binary.BigEndian.PutUint64(v[40:48], f.cookie)
	binary.BigEndian.PutUint16(v[48:50], uint16(f.command))
	binary.BigEndian.PutUint16(v[50:52], f.idleTimeout)
	binary.BigEndian.PutUint16(v[52:54], f.hardTimeout)
	binary.BigEndian.PutUint16(v[54:56], f.priority)
	binary.BigEndian.PutUint32(v[56:60], f.bufferID)
	binary.BigEndian.PutUint16(v[60:62], f.outPort)
	binary.BigEndian.PutUint16(v[62:64], uint16(f.flags))
	copy(v[64:], a)
	f.SetPayload(v)
	return f.Message.MarshalBinary()
}

func (f *flowMod) UnmarshalBinary(data []byte) error {
	if err := f.Message.UnmarshalBinary(data); err != nil {
		return err
	}

	payload := f.Payload()
	if payload == nil || len(payload) < 68 {
		return openflow.ErrInvalidPacketLength
	}
	f.match = NewMatch()
	f.action = NewActionHeader()
	if err := f.match.UnmarshalBinary(payload[0:40]); err != nil {
		return err
	}
	f.cookie = binary.BigEndian.Uint64(payload[40:48])
	f.command = openflow.FlowCommand(binary.BigEndian.Uint16(payload[48:50]))
	f.idleTimeout = binary.BigEndian.Uint16(payload[50:52])
	f.idleTimeout = binary.BigEndian.Uint16(payload[52:54])
	f.priority = binary.BigEndian.Uint16(payload[54:56])
	f.bufferID = binary.BigEndian.Uint32(payload[56:60])
	f.outPort = binary.BigEndian.Uint16(payload[60:62])
	f.flags = openflow.FlowFlag(binary.BigEndian.Uint16(payload[62:64]))
	if err := f.action.UnmarshalBinary(payload[64:]); err != nil {
		return err
	}
	return nil
}

func NewFlowMod(xid uint32) openflow.FlowMod {
	return &flowMod{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_FLOW_MOD, xid),
		match  : NewMatch(),
	}
}
