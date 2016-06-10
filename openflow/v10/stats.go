package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)

type statsHeader struct {
	openflow.Message
	typ openflow.StatsType
	flags uint16
	statsPayload []byte
}

func (s *statsHeader) Type() openflow.StatsType {
	return s.typ
}

func (s *statsHeader) SetType(t openflow.StatsType) {
	s.typ = t
}

func (s *statsHeader) Flags() uint16 {
	return s.flags
}

func (s *statsHeader) SetFlags(f uint16) {
	s.flags = f
}

func (s *statsHeader) StatsPayload() []byte {
	return s.statsPayload
}

func (s *statsHeader) SetStatsPayload(pl []byte) {
	s.statsPayload = pl
}

func (s *statsHeader) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4 + len(s.statsPayload))
	binary.BigEndian.PutUint16(v[0:2], uint16(s.typ))
	binary.BigEndian.PutUint16(v[2:4], uint16(s.flags))
	if len(s.statsPayload) > 0 {
		copy(v[4:], s.statsPayload)
	}
	s.SetPayload(v)
	return s.Message.MarshalBinary()
}

func (s *statsHeader) UnmarshalBinary(data []byte) error {
	if err := s.Message.UnmarshalBinary(data); err != nil {
		return err
	}

	payload := s.Payload()
	if payload == nil || len(payload) < 4 {
		return openflow.ErrInvalidPacketLength
	}
	s.typ = openflow.StatsType(binary.BigEndian.Uint16(payload[0:2]))
	s.flags = binary.BigEndian.Uint16(payload[2:4])
	if len(payload) > 4 {
		s.SetStatsPayload(payload[4:])
	}
	return nil
}

func NewStatsRequestHeader(xid uint32) openflow.StatsRequest {
	return &statsHeader{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_STATS_REQUEST, xid),
	}
}

func NewStatsReplyHeader(xid uint32) openflow.StatsRequest {
	return &statsHeader{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_STATS_REPLY, xid),
	}
}

type statsRequestDescription struct {
	*statsHeader
}

type StatsRequestDescription interface {
	openflow.StatsRequest
}

func (s *statsRequestDescription) MarshalBinary() ([]byte, error) {
	return s.statsHeader.MarshalBinary()
}

func (s *statsRequestDescription) UnmarshalBinary(data []byte) error {
	return s.statsHeader.UnmarshalBinary(data)
}

func NewStatsRequestDescription(xid uint32) StatsRequestDescription {
	srd := NewStatsRequestHeader(xid)
	srd.SetType(openflow.STATS_Description)
	return srd
}

type statsRequestFlow struct {
	*statsHeader
	match openflow.Match
	tableID uint8
	outPort uint16
}

type StatsRequestFlow interface {
	openflow.StatsRequest
	Match() openflow.Match
	SetMatch(openflow.Match)
	TableID() uint8
	SetTableID(uint8)
	OutPort() uint16
	SetOutPort(uint16)
}

func (s *statsRequestFlow) Match() openflow.Match {
	return s.match
}

func (s *statsRequestFlow) SetMatch(m openflow.Match) {
	s.match = m
}

func (s *statsRequestFlow) TableID() uint8 {
	return s.tableID
}

func (s *statsRequestFlow) SetTableID(tid uint8) {
	s.tableID = tid
}

func (s *statsRequestFlow) OutPort() uint16 {
	return s.outPort
}

func (s *statsRequestFlow) SetOutPort(op uint16) {
	s.outPort = op
}

func (s *statsRequestFlow) MarshalBinary() ([]byte, error) {
	v := make([]byte, 44)
	m ,err := s.match.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(v[0:40], m)
	v[40] = s.tableID
	//v[41] is pad
	binary.BigEndian.PutUint16(v[42:44], s.outPort)
	return s.statsHeader.MarshalBinary()
}

func (s *statsRequestFlow) UnmarshalBinary(data []byte) error {
	if err := s.statsHeader.UnmarshalBinary(data); err != nil{
		return err
	}
	if len(s.statsPayload) != 44 {
		return openflow.ErrInvalidDataLength
	}
	payload := s.statsPayload
	s.match = NewMatch()
	if err := s.match.UnmarshalBinary(payload[0:40]); err != nil {
		return err
	}
	s.tableID = payload[40]
	//payload[41] is pad
	s.outPort = binary.BigEndian.Uint16(payload[42:44])
	return nil
}

func NewStatsReuqestFlow(xid uint32) StatsRequestFlow {
	srf := &statsRequestFlow{
		statsHeader 	: NewStatsRequestHeader(xid).(*statsHeader),
		match 			: NewMatch()
	}
	srf.statsHeader.SetType(openflow.STATS_Flow)
	return srf
}