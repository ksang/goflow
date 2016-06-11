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
	s.SetStatsPayload(v)
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
		match 			: NewMatch(),
	}
	srf.statsHeader.SetType(openflow.STATS_Flow)
	return srf
}

// Aggregate shares the same struct as Flow
type statsRequestAggregate struct {
	statsRequestFlow
}

type StatsRequestAggregate interface {
	StatsRequestFlow
}

func NewStatsReuqestAggregate(xid uint32) StatsRequestAggregate {
	sra := &statsRequestAggregate{
		statsRequestFlow{
			statsHeader 	: NewStatsRequestHeader(xid).(*statsHeader),
			match 			: NewMatch(),			
		},
	}
	sra.statsHeader.SetType(openflow.STATS_Aggregate)
	return sra
}

// StatsRequestTable shares same struct as description
type statsRequestTable struct {
	statsRequestDescription
}

type StatsRequestTable interface {
	StatsRequestDescription
}

func NewStatsReuqestTable(xid uint32) StatsRequestTable {
	srt := NewStatsRequestHeader(xid)
	srt.SetType(openflow.STATS_Table)
	return srt
}

type statsRequestPort struct {
	*statsHeader
	portNumber uint16
}

type StatsRequestPort interface {
	openflow.StatsRequest
	PortNumber() uint16
	SetPortNumber(uint16) error
}

func (s *statsRequestPort) PortNumber() uint16 {
	return s.portNumber
}

func (s *statsRequestPort) SetPortNumber(pn uint16) error {
	if pn > 0xffef {
		return openflow.ErrInvalidValueProvided
	}
	s.portNumber = pn
	return nil
}

func (s *statsRequestPort) MarshalBinary() ([]byte, error) {
	v := make([]byte, 8)
	binary.BigEndian.PutUint16(v[0:2], s.portNumber)
	// v[2:8] is pad
	s.SetStatsPayload(v)
	return s.statsHeader.MarshalBinary()
}

func (s *statsRequestPort) UnmarshalBinary(data []byte) error {
	if err := s.statsHeader.UnmarshalBinary(data); err != nil{
		return err
	}
	if len(s.statsPayload) != 8 {
		return openflow.ErrInvalidDataLength
	}
	payload := s.statsPayload
	s.portNumber = binary.BigEndian.Uint16(payload[0:2])
	// payload[2:] is pad
	return nil
}

func NewStatsReuqestPort(xid uint32) StatsRequestPort {
	srp := &statsRequestPort{
		statsHeader 	: NewStatsRequestHeader(xid).(*statsHeader),
	}
	srp.SetType(openflow.STATS_Port)
	return srp
}

type statsRequestQueue struct {
	*statsHeader
	portNumber uint16
	queueID uint32
}

type StatsRequestQueue interface {
	openflow.StatsRequest
	PortNumber() uint16
	SetPortNumber(uint16) error
	QueueID() uint32
	SetQueueID(uint32) 
}

func (s *statsRequestQueue) PortNumber() uint16 {
	return s.portNumber
}

func (s *statsRequestQueue) SetPortNumber(pn uint16) error {
	if pn > 0xffef {
		return openflow.ErrInvalidValueProvided
	}
	s.portNumber = pn
	return nil
}

func (s *statsRequestQueue) QueueID() uint32 {
	return s.queueID
}

func (s *statsRequestQueue) SetQueueID(qid uint32) {
	s.queueID = qid
}

func (s *statsRequestQueue) MarshalBinary() ([]byte, error) {
	v := make([]byte, 8)
	binary.BigEndian.PutUint16(v[0:2], s.portNumber)
	binary.BigEndian.PutUint32(v[4:8], s.queueID)
	// v[2:4] is pad
	s.SetStatsPayload(v)
	return s.statsHeader.MarshalBinary()
}

func (s *statsRequestQueue) UnmarshalBinary(data []byte) error {
	if err := s.statsHeader.UnmarshalBinary(data); err != nil{
		return err
	}
	if len(s.statsPayload) != 8 {
		return openflow.ErrInvalidDataLength
	}
	payload := s.statsPayload
	s.portNumber = binary.BigEndian.Uint16(payload[0:2])
	// payload[2:4] is pad
	s.queueID = binary.BigEndian.Uint32(payload[4:8])
	return nil
}

func NewStatsReuqestQueue(xid uint32) StatsRequestQueue {
	srq := &statsRequestQueue{
		statsHeader 	: NewStatsRequestHeader(xid).(*statsHeader),
	}
	srq.SetType(openflow.STATS_Queue)
	return srq
}

type statsRequestVendor struct {
	*statsHeader
	vendorID uint32
}

type StatsRequestVendor interface {
	openflow.StatsRequest
	VendorID() uint32
	SetVendorID(uint32)
}

func (s *statsRequestVendor) VendorID() uint32 {
	return s.vendorID
}

func (s *statsRequestVendor) SetVendorID(vid uint32) {
	s.vendorID = vid
}

func (s *statsRequestVendor) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4)
	binary.BigEndian.PutUint32(v[0:4], s.vendorID)
	s.SetStatsPayload(v)
	return s.statsHeader.MarshalBinary()
}

func (s *statsRequestVendor) UnmarshalBinary(data []byte) error {
	if err := s.statsHeader.UnmarshalBinary(data); err != nil{
		return err
	}
	if len(s.statsPayload) != 4 {
		return openflow.ErrInvalidDataLength
	}
	payload := s.statsPayload
	s.vendorID = binary.BigEndian.Uint32(payload[0:4])
	return nil
}

func NewStatsReuqestVendor(xid uint32) StatsRequestVendor {
	srv := &statsRequestVendor{
		statsHeader 	: NewStatsRequestHeader(xid).(*statsHeader),
	}
	srv.statsHeader.SetType(openflow.STATS_Vendor)
	return srv
}

type statsReplyDescription struct {
	*statsHeader
	mfrDesc *[256]byte
	hwDesc *[256]byte
	swDesc *[256]byte
	serialNum *[32]byte
	dpDesc *[256]byte
}

type StatsReplyDescription interface {
	openflow.StatsReply
	MfrDesc() *[256]byte
	SetMfrDesc(*[256]byte)
	HwDesc() *[256]byte
	SetHwDesc(*[256]byte)
	SwDesc() *[256]byte
	SetSwDesc(*[256]byte)
	SerialNum() *[32]byte
	SetSerialNum(*[32]byte)
	DpDesc() *[256]byte
	SetDpDesc(*[256]byte)
}

func (s *statsReplyDescription) MfrDesc() *[256]byte {
	return s.mfrDesc
}

func (s *statsReplyDescription) SetMfrDesc(mfr *[256]byte) {
	s.mfrDesc = mfr
}

func (s *statsReplyDescription) HwDesc() *[256]byte {
	return s.hwDesc
}

func (s *statsReplyDescription) SetHwDesc(hd *[256]byte) {
	s.hwDesc = hd
}

func (s *statsReplyDescription) SwDesc() *[256]byte {
	return s.swDesc
}

func (s *statsReplyDescription) SetSwDesc(sd *[256]byte) {
	s.swDesc = sd
}

func (s *statsReplyDescription) SerialNum() *[32]byte {
	return s.serialNum
}

func (s *statsReplyDescription) SetSerialNum(sn *[32]byte) {
	s.serialNum = sn
}

func (s *statsReplyDescription) DpDesc() *[256]byte {
	return s.dpDesc
}

func (s *statsReplyDescription) SetDpDesc(dd *[256]byte) {
	s.dpDesc = dd
}

func (s *statsReplyDescription) MarshalBinary() ([]byte, error) {
	v := make([]byte, 1056)
	copy(v, s.mfrDesc[:])
	copy(v, s.hwDesc[:])
	copy(v, s.swDesc[:])
	copy(v, s.serialNum[:])
	copy(v, s.dpDesc[:])
	s.SetStatsPayload(v)
	return s.statsHeader.MarshalBinary()
}

func (s *statsReplyDescription) UnmarshalBinary(data []byte) error {
	if err := s.statsHeader.UnmarshalBinary(data); err != nil{
		return err
	}
	if len(s.statsPayload) != 1056 {
		return openflow.ErrInvalidDataLength
	}
	payload := s.statsPayload
	var (
		md [256]byte
	 	hd [256]byte
	 	sd [256]byte
 		sn [32]byte
 		dd [256]byte
 	)
	copy(md[:], payload[0:256])
	copy(hd[:], payload[256:512])
	copy(sd[:], payload[512:768])
	copy(sn[:], payload[768:800])
	copy(dd[:], payload[800:1056])
	s.mfrDesc = &md
	s.hwDesc = &hd
	s.swDesc = &sd
	s.serialNum = &sn
	s.dpDesc = &dd
	return nil
}

func NewStatsReplyDescription(xid uint32) StatsReplyDescription {
	srd := &statsReplyDescription{
		statsHeader 	: NewStatsRequestHeader(xid).(*statsHeader),
	}
	srd.statsHeader.SetType(openflow.STATS_Description)
	var (
		md [256]byte
	 	hd [256]byte
	 	sd [256]byte
 		sn [32]byte
 		dd [256]byte
 	)
 	srd.mfrDesc = &md
	srd.hwDesc = &hd
	srd.swDesc = &sd
	srd.serialNum = &sn
	srd.dpDesc = &dd
	return srd
}

//TODO other StatsReply types