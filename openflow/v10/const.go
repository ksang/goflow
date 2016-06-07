/*
Package openflow/v10 impelements openflow version 1.0.0
*/
package v10

// Openflow message types
const (
	/* Immutable messages. */
	OFPT_HELLO = iota		/* Symmetric message */
	OFPT_ERROR				/* Symmetric message */
	OFPT_ECHO_REQUEST		/* Symmetric message */
	OFPT_ECHO_REPLY			/* Symmetric message */
	OFPT_VENDOR 			/* Symmetric message */

	/* Switch configuration messages. */
	OFPT_FEATURES_REQUEST	/* Controller/switch message */
	OFPT_FEATURES_REPLY		/* Controller/switch message */
	OFPT_GET_CONFIG_REQUEST /* Controller/switch message */
	OFPT_GET_CONFIG_REPLY 	/* Controller/switch message */
	OFPT_SET_CONFIG 		/* Controller/switch message */

	/* Asynchronous messages. */
	OFPT_PACKET_IN 			/* Async message */
	OFPT_FLOW_REMOVED       /* Async message */
	OFPT_PORT_STATUS 		/* Async message */

	/* Controller command messages. */
	OFPT_PACKET_OUT  		/* Controller/switch message */
	OFPT_FLOW_MOD 			/* Controller/switch message */
	OFPT_PORT_MOD  			/* Controller/switch message */

	/* Statistics messages. */
	OFPT_STATS_REQUEST 		/* Controller/switch message */
	OFPT_STATS_REPLY        /* Controller/switch message */

	/* Barrier messages. */
	OFPT_BARRIER_REQUEST 	/* Controller/switch message */
	OFPT_BARRIER_REPLY 		/* Controller/switch message */

	/* Queue Configuration messages. */
	OFPT_QUEUE_GET_CONFIG_REQUEST 	/* Controller/switch message */
	OFPT_QUEUE_GET_CONFIG_REPLY   	/* Controller/switch message */
)

// Openflow action types
const (
	OFPAT_OUTPUT       = iota /* Output to switch port. */
	OFPAT_SET_VLAN_VID        /* Set the 802.1q VLAN id. */
	OFPAT_SET_VLAN_PCP        /* Set the 802.1q priority. */
	OFPAT_STRIP_VLAN          /* Strip the 802.1q header. */
	OFPAT_SET_DL_SRC          /* Ethernet source address. */
	OFPAT_SET_DL_DST          /* Ethernet destination address. */
	OFPAT_SET_NW_SRC          /* IP source address. */
	OFPAT_SET_NW_DST          /* IP destination address. */
	OFPAT_SET_NW_TOS          /* IP ToS (DSCP field, 6 bits). */
	OFPAT_SET_TP_SRC          /* TCP/UDP source port. */
	OFPAT_SET_TP_DST          /* TCP/UDP destination port. */
	OFPAT_ENQUEUE             /* Output to queue. */
	OFPAT_VENDOR       = 0xffff
)

// Openflow port settings
const (
	/* Maximum number of physical switch ports. */
	OFPP_MAX        = 0xff00	
	
    /* Fake output "ports". */
	OFPP_IN_PORT    = 0xfff8	/* Send the packet out the input port.  This
                                   virtual port must be explicitly used
                                   in order to send back out of the input
                                   port. */
	OFPP_TABLE      = 0xfff9	/* Perform actions in flow table.
								   NB: This can only be the destination
								   port for packet-out messages. */
	OFPP_FLOOD      = 0xfffb	/* Process with normal L2/L3 switching. */
	OFPP_ALL        = 0xfffc	/* All physical ports except input port and
              					   those disabled by STP. */
	OFPP_CONTROLLER = 0xfffd    /* Send to controller. */
	OFPP_LOCAL		= 0xfffe	/* Local openflow "port". */
	OFPP_NONE       = 0xffff 	/* Not associated with a physical port. */
)

const (
	OFPFW_IN_PORT     = 1 << 0  /* Switch input port. */
	OFPFW_DL_VLAN     = 1 << 1  /* VLAN id. */
	OFPFW_DL_SRC      = 1 << 2  /* Ethernet source address. */
	OFPFW_DL_DST      = 1 << 3  /* Ethernet destination address. */
	OFPFW_DL_TYPE     = 1 << 4  /* Ethernet frame type. */
	OFPFW_NW_PROTO    = 1 << 5  /* IP protocol. */
	OFPFW_TP_SRC      = 1 << 6  /* TCP/UDP source port. */
	OFPFW_TP_DST      = 1 << 7  /* TCP/UDP destination port. */
	OFPFW_DL_VLAN_PCP = 1 << 20 /* VLAN priority. */
	OFPFW_NW_TOS      = 1 << 21 /* IP ToS (DSCP field, 6 bits). */
)

const (
	OFPPF_10MB_HD    = 1 << 0  /* 10 Mb half-duplex rate support. */
	OFPPF_10MB_FD    = 1 << 1  /* 10 Mb full-duplex rate support. */
	OFPPF_100MB_HD   = 1 << 2  /* 100 Mb half-duplex rate support. */
	OFPPF_100MB_FD   = 1 << 3  /* 100 Mb full-duplex rate support. */
	OFPPF_1GB_HD     = 1 << 4  /* 1 Gb half-duplex rate support. */
	OFPPF_1GB_FD     = 1 << 5  /* 1 Gb full-duplex rate support. */
	OFPPF_10GB_FD    = 1 << 6  /* 10 Gb full-duplex rate support. */
	OFPPF_COPPER     = 1 << 7  /* Copper medium. */
	OFPPF_FIBER      = 1 << 8  /* Fiber medium. */
	OFPPF_AUTONEG    = 1 << 9  /* Auto-negotiation. */
	OFPPF_PAUSE      = 1 << 10 /* Pause. */
	OFPPF_PAUSE_ASYM = 1 << 11 /* Asymmetric pause. */
)

const (
	OFPPC_PORT_DOWN    = 1 << 0
	OFPPC_NO_STP       = 1 << 1
	OFPPC_NO_RECV      = 1 << 2
	OFPPC_NO_RECV_STP  = 1 << 3
	OFPPC_NO_FLOOD     = 1 << 4
	OFPPC_NO_FWD       = 1 << 5
	OFPPC_NO_PACKET_IN = 1 << 6
)

const (
	OFPPS_LINK_DOWN   = 1 << 0
	OFPPS_STP_LISTEN  = 0 << 8 /* Not learning or relaying frames. */
	OFPPS_STP_LEARN   = 1 << 8 /* Learning but not relaying frames. */
	OFPPS_STP_FORWARD = 2 << 8 /* Learning and relaying frames. */
	OFPPS_STP_BLOCK   = 3 << 8 /* Not part of spanning tree. */
	OFPPS_STP_MASK    = 3 << 8 /* Bit mask for OFPPS_STP_* values. */
)

const (
	OFPFF_SEND_FLOW_REM = 1 << 0 /* Send flow removed message when flow expires or is deleted. */
	OFPFF_CHECK_OVERLAP = 1 << 1 /* Check for overlapping entries first. */
	OFPFF_EMERG         = 1 << 2 /* Remark this is for emergency. */
)

const (
	OFP_NO_BUFFER = 0xffffffff
)

const (
	OFPFC_ADD           = 0 /* New flow. */
	OFPFC_MODIFY        = 1 /* Modify all matching flows. */
	OFPFC_MODIFY_STRICT = 2 /* Modify entry strictly matching wildcards and priority. */
	OFPFC_DELETE        = 3 /* Delete all matching flows. */
	OFPFC_DELETE_STRICT = 4 /* Delete entry strictly matching wildcards and priority. */
)

const (
	/* Description of this OpenFlow switch.
	 * The request body is empty.
	 * The reply body is struct ofp_desc_stats. */
	OFPST_DESC = iota
	/* Individual flow statistics.
	 * The request body is struct ofp_flow_stats_request.
	 * The reply body is an array of struct ofp_flow_stats. */
	OFPST_FLOW
	/* Aggregate flow statistics.
	 * The request body is struct ofp_aggregate_stats_request.
	 * The reply body is struct ofp_aggregate_stats_reply. */
	OFPST_AGGREGATE
	/* Flow table statistics.
	 * The request body is empty.
	 * The reply body is an array of struct ofp_table_stats. */
	OFPST_TABLE
	/* Physical port statistics.
	 * The request body is struct ofp_port_stats_request.
	 * The reply body is an array of struct ofp_port_stats. */
	OFPST_PORT
	/* Queue statistics for a port
	 * The request body defines the port
	 * The reply body is an array of struct ofp_queue_stats */
	OFPST_QUEUE
	/* Vendor extension.
	 * The request and reply bodies begin with a 32-bit vendor ID, which takes
	 * the same form as in "struct ofp_vendor_header". The request and reply
	 * bodies are otherwise vendor-defined. */
	OFPST_VENDOR = 0xffff
)

// Config Flags
const (
	OFPC_FRAG_NORMAL = iota /* No special handling for fragments. */
	OFPC_FRAG_DROP          /* Drop fragments. */
	OFPC_FRAG_REASM         /* Reassemble (only if OFPC_IP_REASM set). */
	OFPC_FRAG_MASK
)

const (
	OFPPR_ADD    = 0
	OFPPR_DELETE = 1
	OFPPR_MODIFY = 2
)

// Features Capabilities
const (
	CAP_FLOW_STATS 		= 0x01 << iota
	CAP_TABLE_STATS
	CAP_PORT_STATS 
	CAP_STP 
	CAP_RESERVED
	CAP_IP_REASM
	CAP_QUEUE_STATS	
	CAP_ARP_MATCH_IP
)

// Features Actions
const (
	ACT_OUTPUT			= 0x1 << iota
	ACT_SET_VLAN_VID
	ACT_SET_VLAN_PCP
	ACT_STRIP_VLAN
	ACT_SET_DL_SRC
	ACT_SET_DL_DST	
	ACT_SET_NW_SRC
	ACT_SET_NW_DST
	ACT_SET_NW_TOS
	ACT_SET_TP_SRC
	ACT_SET_TP_DST
	ACT_ENQUEUE	
)

// Port Config
const (
	CONF_PortDown 		= 0x1 << iota
	CONF_NoSTP
	CONF_NoRecv
	CONF_NoRecvSTP
	CONF_NoFlood
	CONF_NoFwd
	CONF_NoPacketIn
)

// Port state
const (
	STATE_STPListen		= 0x0
	STATE_LinkDown		= 0x1 << iota
	STATE_STPLearn
	STATE_STPForward
	STATE_STPBlock
	STATE_STPMask
)

// Port feature
const (
	FEAT_10MB_HD		= 0x1 << iota
	FEAT_10MB_FD
	FEAT_100MB_HD
	FEAT_100MB_FD
	FEAT_1GB_HD	
	FEAT_1GB_FD
	FEAT_10GB_FD
	FEAT_Copper
	FEAT_Fiber
	FEAT_AutoNeg
	FEAT_Pause
	FEAT_PauseAsym
)