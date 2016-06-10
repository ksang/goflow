package openflow

// Openflow version number
const (
	OF10_VERSION = 0x01
	OF11_VERSION = 0x02
	OF12_VERSION = 0x03
	OF13_VERSION = 0x04
	OF14_VERSION = 0x05
)

const (
	OF_HEADER_SIZE = 8
)

// Port id reserved values, for other normal ports should specify their own id value
type PortID uint16
const (
	Max 	PortID = 0xFF00
	inPort 	PortID = 0xFFF8 + iota
	Table
	Normal
	Flood
	All
	Controller
	Local
	// workaround complier complaining overflow
	None  	PortID = 0xFFFF
)

// Port reason codes
type PortReason uint8
const (
	PortAdded PortReason = iota
	PortDeleted
	PortModified
)

// Port state
type PortState uint32
const (
	STPListen PortState = 0x00
	LinkDown  PortState = 0x01 << iota
	STPLearn
	STPForward
	STPBlock
	STPMask
)

// Port config
type PortConfig uint32
const (
	PortDown PortConfig = 0x01 << iota
	NoSTP
	NoRecv
	NoRecvSTP
	NoFlood
	NoFwd
	NoPacketIn
)
// Port feature
type PortFeature uint32
const (
	HD_10MB PortFeature = 0x01 << iota
	FD_10MB
	HD_100MB
	FD_100MB
	HD_1GB
	FD_1GB
	FD_10GB
	Copper
	Fiber
	AutoNeg
	Pause
	PauseAsym
)

type FlowCommand uint16

const (
	Add FlowCommand = iota
	Modify
	ModifyStrict
	Delete
	DeleteStrict
)

type FlowFlag uint16
const (
	SendFlowRem FlowFlag = iota + 0x01
	CheckOverlap
	Emerg
)

// Features Capabilities
type FeatureCapability uint32
const (
	FLOW_STATS 	FeatureCapability = 0x01 << iota
	TABLE_STATS
	PORT_STATS
	STP
	RESERVED
	IP_REASM
	QUEUE_STATS
	ARP_MATCH_IP
)

// Features Actions
type FeatureAction uint32
const (
	OUTPUT 	FeatureAction = 0x1 << iota
	SET_VLAN_VID
	SET_VLAN_PCP
	STRIP_VLAN
	SET_DL_SRC
	SET_DL_DST
	SET_NW_SRC
	SET_NW_DST
	SET_NW_TOS
	SET_TP_SRC
	SET_TP_DST
	ENQUEUE
)

// StatRequest/Response type
type StatsType uint16
const (
	STATS_Description StatsType = iota
	STATS_Flow
	STATS_Aggregate
	STATS_Table
	STATS_Port
	STATS_Queue
	STATS_Vendor
)