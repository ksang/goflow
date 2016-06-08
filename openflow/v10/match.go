package v10

import (
	"encoding/binary"
	"net"
	"github.com/ksang/goflow/openflow"
)

type wildcard struct {
	inPort    bool /* Switch input port. */
	dlVlan    bool /* VLAN id. */
	dlSrc     bool /* Ethernet source address. */
	dlDst     bool /* Ethernet destination address. */
	dlType    bool /* Ethernet frame type. */
	nwProto   bool /* IP protocol. */
	tpSrc     bool /* TCP/UDP source port. */
	tpDst     bool /* TCP/UDP destination port. */
	// IP wildcards are inversed VLSM numbers, it needs to be converted
	// e.g 0  is 255.255.255.255
	//	   32 is 0.0.0.0
	nwSrc        uint8
	nwDst        uint8
	dlVlanPCP    bool /* VLAN priority. */
	nwTos        bool
	all 		 bool
}

func (w *wildcard) MarshalBinary() ([]byte, error) {
	data := make([]byte, 4)
	// if wildcard all is set, return 0x003fffff directly
	if w.all {
		binary.BigEndian.PutUint32(data, 0x003fffff)
		return data, nil
	}

	var v uint32 = 0

	if w.inPort {
		v = v | OFPFW_IN_PORT
	}
	if w.dlVlan {
		v = v | OFPFW_DL_VLAN
	}
	if w.dlSrc {
		v = v | OFPFW_DL_SRC
	}
	if w.dlDst {
		v = v | OFPFW_DL_DST
	}
	if w.dlType {
		v = v | OFPFW_DL_TYPE
	}
	if w.nwProto {
		v = v | OFPFW_NW_PROTO
	}
	if w.tpSrc {
		v = v | OFPFW_TP_SRC
	}
	if w.tpDst {
		v = v | OFPFW_TP_DST
	}
	if w.nwSrc > 0 {
		v = v | (uint32(w.nwSrc) << 8)
	}
	if w.nwDst > 0 {
		v = v | (uint32(w.nwDst) << 14)
	}
	if w.dlVlanPCP {
		v = v | OFPFW_DL_VLAN_PCP
	}
	if w.nwTos {
		v = v | OFPFW_NW_TOS
	}
	binary.BigEndian.PutUint32(data[0:4], v)
	return data, nil
}

func (w *wildcard) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) != 4 {
		return openflow.ErrInvalidPacketLength
	}

	v := binary.BigEndian.Uint32(data[0:4])
	if v&OFPFW_IN_PORT != 0 {
		w.inPort = true
	}
	if v&OFPFW_DL_VLAN != 0 {
		w.dlVlan = true
	}
	if v&OFPFW_DL_SRC != 0 {
		w.dlSrc = true
	}
	if v&OFPFW_DL_DST != 0 {
		w.dlDst = true
	}
	if v&OFPFW_DL_TYPE != 0 {
		w.dlType = true
	}
	if v&OFPFW_NW_PROTO != 0 {
		w.nwProto = true
	}
	if v&OFPFW_TP_SRC != 0 {
		w.tpSrc = true
	}
	if v&OFPFW_TP_DST != 0 {
		w.tpDst = true
	}
	w.nwSrc = uint8((v & (uint32(0x3F) << 8)) >> 8)
	w.nwDst = uint8((v & (uint32(0x3F) << 14)) >> 14)
	if v&OFPFW_DL_VLAN_PCP != 0 {
		w.dlVlanPCP = true
	}
	if v&OFPFW_NW_TOS != 0 {
		w.nwTos = true
	}
	return nil
}

type match struct {
	wildcards 	wildcard
	inPort 		uint16
	dlSrc     	net.HardwareAddr
	dlDst     	net.HardwareAddr
	dlVlan     	uint16
	dlPCP 		uint8
	dlType   	uint16
	nwTos 		uint8
	nwProto   	uint8
	nwSrc     	net.IP
	nwDst    	net.IP
	tpSrc    	uint16
	tpDst   	uint16
}

// NewMatch returns a Match whose fields are all wildcarded
func NewMatch() openflow.Match {
	return &match{
		wildcards: wildcard{
					 all: true,
					 },
		dlSrc:		net.HardwareAddr([]byte{0, 0, 0, 0, 0, 0}),
		dlDst:		net.HardwareAddr([]byte{0, 0, 0, 0, 0, 0}),
		nwSrc:		net.IPv4zero,
		nwDst:		net.IPv4zero,
	}
}

func (m *match)	Wildcards() uint32 {
	v, err := m.wildcards.MarshalBinary()
	if err != nil {
		// Not possible to happen
		return 0x0
	}
	return binary.BigEndian.Uint32(v)
}

func (m *match) InPort() (bool, uint16) {
	return m.wildcards.inPort, m.inPort
}
	
func (m* match) SetInPort(ip uint16) {
	m.inPort = ip
	m.wildcards.inPort = false
}

func (m *match)	SetWildcardInPort() {
	m.wildcards.inPort = true
}

func (m *match)	DLSrc() (bool, net.HardwareAddr) {
	return m.wildcards.dlSrc, m.dlSrc
}

func (m *match)	SetDLSrc(mac net.HardwareAddr) {
	m.dlSrc = mac
	m.wildcards.dlSrc = false
}

func (m *match)	SetWildcardDLSrc() {
	m.wildcards.dlSrc = true
}

func (m *match)	DLDst() (bool, net.HardwareAddr) {
	return m.wildcards.dlDst, m.dlDst
}

func (m *match)	SetDLDst(mac net.HardwareAddr) {
	m.dlDst = mac
	m.wildcards.dlDst = false
}
func (m *match)	SetWildcardDLDst() {
	m.wildcards.dlDst = true
}

func (m *match)	DLVlan() (bool, uint16) {
	return m.wildcards.dlVlan, m.dlVlan
}

func (m *match)	SetDLVlan(vlan uint16) error {
	if int(vlan) > 4096 {
		return openflow.ErrInvalidVlanID
	}
	m.dlVlan = vlan
	m.wildcards.dlVlan = false
	return nil
}

func (m *match)	SetWildcardDLVlan() {
	m.wildcards.dlVlan = true
}

func (m *match)	DLPCP() (bool, uint8) {
	return m.wildcards.dlVlanPCP, m.dlPCP
}

func (m *match)	SetDLPCP(pcp uint8) {
	m.dlPCP = pcp
	m.wildcards.dlVlanPCP = false
}

func (m *match)	SetWildcardDLVlanPCP(){
	m.wildcards.dlVlanPCP = true
}

func (m *match)	DLType() (bool, uint16) {
	return m.wildcards.dlType, m.dlType
}

func (m *match)	SetDLType(dlt uint16) error {
	if dlt != 0x0800 {
		return openflow.ErrUnsupportedEtherType
	}
	m.dlType = dlt
	m.wildcards.dlType = false
	return nil
}
	
func (m *match)	SetWildcardDLType() {
	m.wildcards.dlType = true
}

func (m *match)	NWTos() (bool, uint8) {
	return m.wildcards.nwTos, m.nwTos
}

func (m *match)	SetNWTos(tos uint8) {
	m.nwTos = tos
	m.wildcards.nwTos = false
}

func (m *match)	SetWildcardNWTos() {
	m.wildcards.nwTos = true
}

func (m *match)	NWProto() (bool, uint8) {
	return m.wildcards.nwProto, m.nwProto
}

func (m *match)	SetNWProto(proto uint8) error {
	if proto != 0x06 && proto != 0x11 {
		return openflow.ErrUnsupportedIPProtocol
	}
	m.nwProto = proto
	m.wildcards.nwProto = false
	return nil
}
	
func (m *match)	SetWildcardNWProto() {
	m.wildcards.nwProto = true
}

func (m *match)	NWSrc() net.IP {
	mask := net.CIDRMask(32 - int(m.wildcards.nwSrc), 32)
	return m.nwSrc.Mask(mask)
}

func (m *match) SetNWSrc(ip net.IP) {
	m.nwSrc = ip
}

func (m *match) SetWildcardNWSrc(vlsm int) {
	if vlsm < 0{
		vlsm = 0
	}
	if vlsm > 32 {
		vlsm = 32
	}
	// convert it from vlsm to our mask format
	m.wildcards.nwSrc = 0x20 - uint8(vlsm)
}

func (m *match) NWDst() net.IP {
	mask := net.CIDRMask(32 - int(m.wildcards.nwDst), 32)
	return m.nwDst.Mask(mask)
}

func (m *match) SetNWDst(ip net.IP) {
	m.nwDst = ip
}
func (m *match) SetWildcardNWDst(vlsm int) {
	if vlsm < 0{
		vlsm = 0
	}
	if vlsm > 32 {
		vlsm = 32
	}
	// convert it from vlsm to our mask format
	m.wildcards.nwDst = 0x20 - uint8(vlsm)
}
	// Source and destination port
func (m *match) TPSrc() (bool, uint16) {
	return m.wildcards.tpSrc, m.tpSrc
}

func (m *match) SetTPSrc(port uint16) {
	m.tpSrc = port
	m.wildcards.tpSrc = false
}

func (m *match) SetWildcardTPSrc() {
	m.wildcards.tpSrc = true
}

func (m *match) TPDst() (bool, uint16) {
	return m.wildcards.tpDst, m.tpDst
}

func (m *match) SetTPDst(port uint16) {
	m.tpDst = port
	m.wildcards.tpDst = false
}
func (m *match) SetWildcardTPDst() {
	m.wildcards.tpDst = true
}

func (m *match) MarshalBinary() ([]byte, error) {
	wildcard, err := m.wildcards.MarshalBinary()
	if err != nil {
		return nil, err
	}

	data := make([]byte, 40)
	copy(data[0:4], wildcard)
	binary.BigEndian.PutUint16(data[4:6], m.inPort)
	copy(data[6:12], m.dlSrc)
	copy(data[12:18], m.dlDst)
	binary.BigEndian.PutUint16(data[18:20], m.dlVlan)
	data[20] = m.dlPCP
	// data[21] is padding
	binary.BigEndian.PutUint16(data[22:24], m.dlType)
	data[24] = m.nwTos
	data[25] = m.nwProto
	// data[26:28] is padding
	copy(data[28:32], m.NWSrc())
	copy(data[32:36], m.NWDst())
	binary.BigEndian.PutUint16(data[36:38], m.tpSrc)
	binary.BigEndian.PutUint16(data[38:40], m.tpDst)

	return data, nil
}

//TODO
func (m *match) UnmarshalBinary(data []byte) error {
	if len(data) != 40 {
		return openflow.ErrInvalidPacketLength
	}

	m.wildcards = wildcard{}
	if err := m.wildcards.UnmarshalBinary(data[0:4]); err != nil {
		return err
	}
	m.inPort = binary.BigEndian.Uint16(data[4:6])
	m.dlSrc = make(net.HardwareAddr, 6)
	copy(m.dlSrc, data[6:12])
	m.dlDst = make(net.HardwareAddr, 6)
	copy(m.dlDst, data[12:18])
	m.dlVlan = binary.BigEndian.Uint16(data[18:20])
	m.dlPCP = data[20]
	// data[21] = padding
	m.dlType = binary.BigEndian.Uint16(data[22:24])
	m.nwTos = data[24]
	m.nwProto = data[25]
	// data[26:28] = padding
	m.nwSrc = net.IPv4(data[28], data[29], data[30], data[31])
	m.nwDst = net.IPv4(data[32], data[33], data[34], data[35])
	m.tpSrc = binary.BigEndian.Uint16(data[36:38])
	m.tpDst = binary.BigEndian.Uint16(data[38:40])

	return nil
}
