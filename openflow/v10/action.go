package v10

import (
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
	"net"
)

// action output interface
type ActionOutput interface {
	openflow.Action
	Port() uint16
	SetPort(uint16) error
	MaxLen() uint16
	SetMaxLen(uint16)
}

type ActionSetVLANVID interface {
	openflow.Action
	VLANVID() uint16
	SetVLANVID(uint16)
}

//  Set Priority Code Point (PCP)
type ActionSetVLANPCP interface {
	openflow.Action
	VLANPCP() uint8
	SetVLANPCP(uint8)
}

type ActionStripVLAN interface {
	openflow.Action
}

type ActionSetDLSrc interface {
	openflow.Action
	DLSrc() net.HardwareAddr // source mac address
	SetDLSrc(net.HardwareAddr)
}

type ActionSetDLDst interface {
	openflow.Action
	DLDst() net.HardwareAddr // destination mac address
	SetDLDst(net.HardwareAddr)
}

type ActionSetNWSrc interface {
	openflow.Action
	NWSrc() net.IP // source ip address
	SetNWSrc(net.IP)
}

type ActionSetNWDst interface {
	openflow.Action
	NWDst() net.IP // destination ip address
	SetNWDst(net.IP)
}

type ActionSetNWTos interface {
	openflow.Action
	NWTos() uint8
	SetNWTos(uint8)
}

type ActionSetTPSrc interface {
	openflow.Action
	Port() uint16
	SetPort(uint16) error
}

type ActionSetTPDst interface {
	openflow.Action
	Port() uint16
	SetPort(uint16) error
}

type ActionEnqueue interface {
	openflow.Action
	Port() uint16
	SetPort(uint16) error
	QueueID() uint32
	SetQueueID(uint32)
}

type ActionVendor interface {
	openflow.Action
	Vendor() uint32
	SetVendor(uint32)
}

type actionHead struct {
	actionType uint16
	length     uint16 // the length of entire action block
	payload    []byte
}

func (a *actionHead) Type() uint16 {
	return a.actionType
}

func (a *actionHead) SetType(t uint16) {
	a.actionType = t
}

func (a *actionHead) Length() uint16 {
	return a.length
}

func (a *actionHead) Payload() []byte {
	return a.payload
}

func (a *actionHead) SetPayload(payload []byte) error {
	if payload == nil {
		return openflow.ErrInvalidDataLength
	}
	a.payload = payload
	// actionType + length = 4 bytes
	a.length = uint16(4 + len(payload))
	return nil
}

func (a *actionHead) MarshalBinary() ([]byte, error) {
	if a.length == uint16(0) {
		a.length = uint16(4)
	}
	// length doesn't match payload size
	if a.length != uint16(4+len(a.payload)) {
		return nil, openflow.ErrInvalidDataLength
	}
	v := make([]byte, int(a.length))
	binary.BigEndian.PutUint16(v[0:2], a.actionType)
	binary.BigEndian.PutUint16(v[2:4], a.length)
	if len(a.payload) > 0 {
		copy(v[4:], a.payload)
	}
	return v, nil
}

func (a *actionHead) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) < 4 {
		return openflow.ErrInvalidPacketLength
	}
	a.actionType = binary.BigEndian.Uint16(data[0:2])
	a.length = binary.BigEndian.Uint16(data[2:4])
	if int(a.length) != len(data) {
		return openflow.ErrInvalidDataLength
	}
	if int(a.length) > 4 {
		a.payload = data[4:int(a.length)]
	}
	return nil
}

func NewActionHead() openflow.Action {
	return &actionHead{}
}

// action output definitions
type actionOutput struct {
	actionHead
	port   uint16
	maxLen uint16
}

func (ao *actionOutput) Port() uint16 {
	return ao.port
}

func (ao *actionOutput) SetPort(port uint16) error {
	if port > 0xffef {
		return openflow.ErrInvalidValueProvided
	}
	ao.port = port
	return nil
}

func (ao *actionOutput) MaxLen() uint16 {
	return ao.maxLen
}

func (ao *actionOutput) SetMaxLen(maxLen uint16) {
	ao.maxLen = maxLen
}

func (ao *actionOutput) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4)
	binary.BigEndian.PutUint16(v[0:2], ao.port)
	binary.BigEndian.PutUint16(v[2:4], ao.maxLen)
	if err := ao.SetPayload(v); err != nil {
		return nil, err
	}
	return ao.actionHead.MarshalBinary()
}

func (ao *actionOutput) UnmarshalBinary(data []byte) error {
	if err := ao.actionHead.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := ao.Payload()
	if len(payload) != 4 {
		return openflow.ErrInvalidDataLength
	}
	ao.port = binary.BigEndian.Uint16(payload[0:2])
	ao.maxLen = binary.BigEndian.Uint16(payload[2:4])
	return nil
}

func NewActionOutput() ActionOutput {
	return &actionOutput{
		actionHead: actionHead{
			actionType: OFPAT_OUTPUT,
			length:     8,
		},
	}
}

// action SetVLANVID definitions
type actionSetVLANVID struct {
	actionHead
	vlanVID uint16
}

func (as *actionSetVLANVID) VLANVID() uint16 {
	return as.vlanVID
}

func (as *actionSetVLANVID) SetVLANVID(vid uint16) {
	as.vlanVID = vid
}

func (as *actionSetVLANVID) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4)
	binary.BigEndian.PutUint16(v[0:2], as.vlanVID)
	// v[2:4] is padding
	if err := as.SetPayload(v); err != nil {
		return nil, err
	}
	return as.actionHead.MarshalBinary()
}

func (as *actionSetVLANVID) UnmarshalBinary(data []byte) error {
	if err := as.actionHead.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := as.Payload()
	if len(payload) != 4 {
		return openflow.ErrInvalidDataLength
	}
	as.vlanVID = binary.BigEndian.Uint16(payload[0:2])
	// payload[2:4] is padding
	return nil
}

func NewActionSetVLANVID() ActionSetVLANVID {
	return &actionSetVLANVID{
		actionHead: actionHead{
			actionType: OFPAT_SET_VLAN_VID,
			length:     8,
		},
	}
}

// action SetVLANPCP definitions
type actionSetVLANPCP struct {
	actionHead
	vlanPCP uint8
}

func (as *actionSetVLANPCP) VLANPCP() uint8 {
	return as.vlanPCP
}

func (as *actionSetVLANPCP) SetVLANPCP(pcp uint8) {
	as.vlanPCP = pcp
}

func (as *actionSetVLANPCP) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4)
	v[0] = as.vlanPCP
	// v[1:4] is padding
	if err := as.SetPayload(v); err != nil {
		return nil, err
	}
	return as.actionHead.MarshalBinary()
}

func (as *actionSetVLANPCP) UnmarshalBinary(data []byte) error {
	if err := as.actionHead.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := as.Payload()
	if len(payload) != 4 {
		return openflow.ErrInvalidDataLength
	}
	as.vlanPCP = payload[0]
	// payload[1:4] is padding
	return nil
}

func NewActionSetVLANPCP() ActionSetVLANPCP {
	return &actionSetVLANPCP{
		actionHead: actionHead{
			actionType: OFPAT_SET_VLAN_PCP,
			length:     8,
		},
	}
}

// action StripVLAN definitions
type actionStripVLAN struct {
	actionHead
}

func (as *actionStripVLAN) MarshalBinary() ([]byte, error) {
	return as.actionHead.MarshalBinary()
}

func (as *actionStripVLAN) UnmarshalBinary(data []byte) error {
	return as.actionHead.UnmarshalBinary(data)
}

func NewActionStripVLAN() ActionStripVLAN {
	return &actionSetVLANVID{
		actionHead: actionHead{
			actionType: OFPAT_STRIP_VLAN,
			length:     4,
		},
	}
}

// action SetDLSrc/Dst definitions
type actionSetDL struct {
	actionHead
	mac net.HardwareAddr
}

func (as *actionSetDL) DLSrc() net.HardwareAddr {
	return as.mac
}

func (as *actionSetDL) DLDst() net.HardwareAddr {
	return as.mac
}

func (as *actionSetDL) SetDLSrc(mac net.HardwareAddr) {
	as.mac = mac
}

func (as *actionSetDL) SetDLDst(mac net.HardwareAddr) {
	as.mac = mac
}

func (as *actionSetDL) MarshalBinary() ([]byte, error) {
	v := make([]byte, 12)
	copy(v[0:6], as.mac)
	// v[6:12] is padding
	if err := as.SetPayload(v); err != nil {
		return nil, err
	}
	return as.actionHead.MarshalBinary()
}

func (as *actionSetDL) UnmarshalBinary(data []byte) error {
	if err := as.actionHead.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := as.Payload()
	if len(payload) != 12 {
		return openflow.ErrInvalidDataLength
	}
	as.mac = payload[0:6]
	// payload[6:12] is padding
	return nil
}

func NewActionSetDLSrc() ActionSetDLSrc {
	return &actionSetDL{
		actionHead: actionHead{
			actionType: OFPAT_SET_DL_SRC,
			length:     16,
		},
	}
}

func NewActionSetDLDst() ActionSetDLDst {
	return &actionSetDL{
		actionHead: actionHead{
			actionType: OFPAT_SET_DL_DST,
			length:     16,
		},
	}
}

// action SetNWSrc/Dst definitions
type actionSetNW struct {
	actionHead
	ip net.IP
}

func (as *actionSetNW) NWSrc() net.IP {
	return as.ip
}

func (as *actionSetNW) NWDst() net.IP {
	return as.ip
}

func (as *actionSetNW) SetNWSrc(ip net.IP) {
	as.ip = ip
}

func (as *actionSetNW) SetNWDst(ip net.IP) {
	as.ip = ip
}

func (as *actionSetNW) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4)
	copy(v[0:4], as.ip)
	if err := as.SetPayload(v); err != nil {
		return nil, err
	}
	return as.actionHead.MarshalBinary()
}

func (as *actionSetNW) UnmarshalBinary(data []byte) error {
	if err := as.actionHead.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := as.Payload()
	if len(payload) != 4 {
		return openflow.ErrInvalidDataLength
	}
	as.ip = payload
	return nil
}

func NewActionSetNWSrc() ActionSetNWSrc {
	return &actionSetNW{
		actionHead: actionHead{
			actionType: OFPAT_SET_NW_SRC,
			length:     8,
		},
	}
}

func NewActionSetNWDst() ActionSetNWDst {
	return &actionSetNW{
		actionHead: actionHead{
			actionType: OFPAT_SET_NW_DST,
			length:     8,
		},
	}
}

// action SetVLANPCP definitions
type actionSetNWTos struct {
	actionHead
	nwTos uint8
}

func (as *actionSetNWTos) NWTos() uint8 {
	return as.nwTos
}

func (as *actionSetNWTos) SetNWTos(tos uint8) {
	as.nwTos = tos
}

func (as *actionSetNWTos) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4)
	v[0] = as.nwTos
	// v[1:4] is padding
	if err := as.SetPayload(v); err != nil {
		return nil, err
	}
	return as.actionHead.MarshalBinary()
}

func (as *actionSetNWTos) UnmarshalBinary(data []byte) error {
	if err := as.actionHead.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := as.Payload()
	if len(payload) != 4 {
		return openflow.ErrInvalidDataLength
	}
	as.nwTos = payload[0]
	// payload[1:4] is padding
	return nil
}

func NewActionSetNWTos() ActionSetNWTos {
	return &actionSetNWTos{
		actionHead: actionHead{
			actionType: OFPAT_SET_NW_TOS,
			length:     8,
		},
	}
}

// action SetTPSrc/Dst definitions
type actionSetTP struct {
	actionHead
	port uint16
}

func (as *actionSetTP) Port() uint16 {
	return as.port
}

func (as *actionSetTP) SetPort(port uint16) error {
	if port > 0xffef {
		return openflow.ErrInvalidValueProvided
	}
	as.port = port
	return nil
}

func (as *actionSetTP) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4)
	binary.BigEndian.PutUint16(v[0:2], as.port)
	// v[2:4] is pad
	if err := as.SetPayload(v); err != nil {
		return nil, err
	}
	return as.actionHead.MarshalBinary()
}

func (as *actionSetTP) UnmarshalBinary(data []byte) error {
	if err := as.actionHead.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := as.Payload()
	if len(payload) != 4 {
		return openflow.ErrInvalidDataLength
	}
	as.port = binary.BigEndian.Uint16(payload[0:2])
	//payload[2:4] is pad
	return nil
}

func NewActionSetTPSrc() ActionSetTPSrc {
	return &actionSetTP{
		actionHead: actionHead{
			actionType: OFPAT_SET_TP_SRC,
			length:     8,
		},
	}
}

func NewActionSetTPDst() ActionSetTPDst {
	return &actionSetTP{
		actionHead: actionHead{
			actionType: OFPAT_SET_TP_DST,
			length:     8,
		},
	}
}

// action Enqueue definitions
type actionEnqueue struct {
	actionHead
	port    uint16
	queueID uint32
}

func (ae *actionEnqueue) Port() uint16 {
	return ae.port
}

func (ae *actionEnqueue) SetPort(port uint16) error {
	if port > 0xffef {
		return openflow.ErrInvalidValueProvided
	}
	ae.port = port
	return nil
}

func (ae *actionEnqueue) QueueID() uint32 {
	return ae.queueID
}

func (ae *actionEnqueue) SetQueueID(qid uint32) {
	ae.queueID = qid
}

func (ae *actionEnqueue) MarshalBinary() ([]byte, error) {
	v := make([]byte, 12)
	binary.BigEndian.PutUint16(v[0:2], ae.port)
	// v[2:8] is pad
	binary.BigEndian.PutUint32(v[8:12], ae.queueID)
	if err := ae.SetPayload(v); err != nil {
		return nil, err
	}
	return ae.actionHead.MarshalBinary()
}

func (ae *actionEnqueue) UnmarshalBinary(data []byte) error {
	if err := ae.actionHead.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := ae.Payload()
	if len(payload) != 12 {
		return openflow.ErrInvalidDataLength
	}
	ae.port = binary.BigEndian.Uint16(payload[0:2])
	//payload[2:8] is pad
	ae.queueID = binary.BigEndian.Uint32(payload[8:12])
	return nil
}

func NewActionEnqueue() ActionEnqueue {
	return &actionEnqueue{
		actionHead: actionHead{
			actionType: OFPAT_ENQUEUE,
			length:     16,
		},
	}
}

// action Vendor definitions
type actionVendor struct {
	actionHead
	vendor uint32
}

func (av *actionVendor) Vendor() uint32 {
	return av.vendor
}

func (av *actionVendor) SetVendor(vendor uint32) {
	av.vendor = vendor
}

func (av *actionVendor) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4)
	binary.BigEndian.PutUint32(v[0:4], av.vendor)
	if err := av.SetPayload(v); err != nil {
		return nil, err
	}
	return av.actionHead.MarshalBinary()
}

func (av *actionVendor) UnmarshalBinary(data []byte) error {
	if err := av.actionHead.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := av.Payload()
	if len(payload) != 4 {
		return openflow.ErrInvalidDataLength
	}
	av.vendor = binary.BigEndian.Uint32(payload[0:4])
	return nil
}

func NewActionVendor() ActionVendor {
	return &actionVendor{
		actionHead: actionHead{
			actionType: OFPAT_VENDOR,
			length:     8,
		},
	}
}
