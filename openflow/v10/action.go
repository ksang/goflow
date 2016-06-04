package v10

import(
	"net"
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
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

type actionHead struct {
	actionType uint16
	length uint16 		// the length of entire action block
	payload []byte
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

func (a *actionHead) MarshalActionHead() ([]byte, error) {
	if a.length == uint16(0) {
		a.length = uint16(4)
	}
	// length doesn't match payload size
	if a.length != uint16(4 + len(a.payload)) {
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

func (a *actionHead) UnmarshalActionHead(data []byte) error {
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

func NewActionHead() openflow.ActionHead {
	return &actionHead{}
}

// action output definitions
type actionOutput struct {
	actionHead
	port uint16
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
	return ao.MarshalActionHead()
}

func (ao *actionOutput) UnmarshalBinary(data []byte) error {
	if err := ao.UnmarshalActionHead(data); err != nil {
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
		actionHead 	: actionHead {
			actionType 	: OFPAT_OUTPUT,
			length 		: 8,
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
	return as.MarshalActionHead()
}

func (as *actionSetVLANVID) UnmarshalBinary(data []byte) error {
	if err := as.UnmarshalActionHead(data); err != nil {
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
		actionHead 	: actionHead {
			actionType 	: OFPAT_SET_VLAN_VID,
			length 		: 8,
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
	return as.MarshalActionHead()
}

func (as *actionSetVLANPCP) UnmarshalBinary(data []byte) error {
	if err := as.UnmarshalActionHead(data); err != nil {
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
		actionHead 	: actionHead {
			actionType 	: OFPAT_SET_VLAN_PCP,
			length 		: 8,
		},
	}
}

// action StripVLAN definitions
type actionStripVLAN struct {
	actionHead
}

func (as *actionStripVLAN) MarshalBinary() ([]byte, error) {
	return as.MarshalActionHead()
}

func (as *actionStripVLAN) UnmarshalBinary(data []byte) error {
	return as.UnmarshalActionHead(data)
}

func NewActionStripVLAN() ActionStripVLAN {
	return &actionSetVLANVID{
		actionHead 	: actionHead {
			actionType 	: OFPAT_STRIP_VLAN,
			length 		: 4,
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
	return as.MarshalActionHead()
}

func (as *actionSetDL) UnmarshalBinary(data []byte) error {
	if err := as.UnmarshalActionHead(data); err != nil {
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
		actionHead 	: actionHead {
			actionType 	: OFPAT_SET_DL_SRC,
			length 		: 16,
		},
	}
}

func NewActionSetDLDst() ActionSetDLDst {
	return &actionSetDL{
		actionHead 	: actionHead {
			actionType 	: OFPAT_SET_DL_DST,
			length 		: 16,
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
	return as.MarshalActionHead()
}

func (as *actionSetNW) UnmarshalBinary(data []byte) error {
	if err := as.UnmarshalActionHead(data); err != nil {
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
		actionHead 	: actionHead {
			actionType 	: OFPAT_SET_NW_SRC,
			length 		: 8,
		},
	}
}

func NewActionSetNWDst() ActionSetNWDst {
	return &actionSetNW{
		actionHead 	: actionHead {
			actionType 	: OFPAT_SET_NW_DST,
			length 		: 8,
		},
	}
}