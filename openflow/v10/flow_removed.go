package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)

type flowRemoved struct {
	openflow.Message
	match           openflow.Match
	cookie          uint64
	priority        uint16
	reason          uint8
	durationSec     uint32
	durationNanoSec uint32
	idleTimeout     uint16
	packetCount     uint64
	byteCount       uint64
}

func (f *flowRemoved) Match() openflow.Match {
	return f.match
}
func (f *flowRemoved) SetMatch(m openflow.Match) {
	f.match = m
}

func (f *flowRemoved) Cookie() uint64 {
	return f.cookie
}

func (f *flowRemoved) SetCookie(cookie uint64) {
	f.cookie = cookie
}

func (f *flowRemoved) Priority() uint16 {
	return f.priority
}

func (f *flowRemoved) SetPriority(p uint16) {
	f.priority = p
}

func (f *flowRemoved) Reason() uint8 {
	return f.reason
}

func (f *flowRemoved) SetReason(r uint8) error {
	if r > 0x3 {
		return openflow.ErrInvalidValueProvided
	}
	f.reason = r
	return nil
}

func (f *flowRemoved) DurationSec() uint32 {
	return f.durationSec
}

func (f *flowRemoved) SetDurationSec(d uint32) {
	f.durationSec = d
}

func (f *flowRemoved) DurationNanoSec() uint32 {
	return f.durationNanoSec
}

func (f *flowRemoved) SetDurationNanoSec(d uint32) {
	f.durationNanoSec = d
}

func (f *flowRemoved) IdleTimeout() uint16 {
	return f.idleTimeout
}

func (f *flowRemoved) SetIdleTimeout(it uint16) {
	f.idleTimeout = it
}

func (f *flowRemoved) PacketCount() uint64 {
	return f.packetCount
}

func (f *flowRemoved) SetPacketCount(pc uint64) {
	f.packetCount = pc
}

func (f *flowRemoved) ByteCount() uint64 {
	return f.byteCount
}

func (f *flowRemoved) SetByteCount(bc uint64) {
	f.byteCount = bc
}

func (f *flowRemoved) MarshalBinary() ([]byte, error) {
	v := make([]byte, 80)
	m, err := f.match.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(v[0:40], m)
	binary.BigEndian.PutUint64(v[40:48], f.cookie)
	binary.BigEndian.PutUint16(v[48:50], f.priority)
	v[50] = f.reason
	//v[51] is pad
	binary.BigEndian.PutUint32(v[52:56], f.durationSec)
	binary.BigEndian.PutUint32(v[56:60], f.durationNanoSec)
	binary.BigEndian.PutUint16(v[60:62], f.idleTimeout)
	//v[62:64] is pad
	binary.BigEndian.PutUint64(v[64:72], f.packetCount)
	binary.BigEndian.PutUint64(v[72:80], f.byteCount)
	f.SetPayload(v)
	return f.Message.MarshalBinary()
}

func (f *flowRemoved) UnmarshalBinary(data []byte) error {
	if err := f.Message.UnmarshalBinary(data); err != nil {
		return err
	}

	payload := f.Payload()
	if payload == nil || len(payload) != 80 {
		return openflow.ErrInvalidPacketLength
	}
	f.match = NewMatch()
	if err := f.match.UnmarshalBinary(payload[0:40]); err != nil {
		return err
	}
	f.cookie = binary.BigEndian.Uint64(payload[40:48])
	f.priority = binary.BigEndian.Uint16(payload[48:50])
	f.reason = payload[50]
	// payload[51] is padding
	f.durationSec = binary.BigEndian.Uint32(payload[52:56])
	f.durationNanoSec = binary.BigEndian.Uint32(payload[56:60])
	f.idleTimeout = binary.BigEndian.Uint16(payload[60:62])
	// payload[62:64] is padding
	f.packetCount = binary.BigEndian.Uint64(payload[64:72])
	f.byteCount = binary.BigEndian.Uint64(payload[72:80])

	return nil
}

func NewFlowRemoved(xid uint32) openflow.FlowRemoved {
	return &flowRemoved{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_FLOW_REMOVED, xid),
		match : NewMatch(),
	}
}
