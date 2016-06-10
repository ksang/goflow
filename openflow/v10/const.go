/*
Package openflow/v10 implements openflow version 1.0.0
*/
package v10

// Openflow message types
const (
	/* Immutable messages. */
	OFPT_HELLO        = iota /* Symmetric message */
	OFPT_ERROR               /* Symmetric message */
	OFPT_ECHO_REQUEST        /* Symmetric message */
	OFPT_ECHO_REPLY          /* Symmetric message */
	OFPT_VENDOR              /* Symmetric message */

	/* Switch configuration messages. */
	OFPT_FEATURES_REQUEST   /* Controller/switch message */
	OFPT_FEATURES_REPLY     /* Controller/switch message */
	OFPT_GET_CONFIG_REQUEST /* Controller/switch message */
	OFPT_GET_CONFIG_REPLY   /* Controller/switch message */
	OFPT_SET_CONFIG         /* Controller/switch message */

	/* Asynchronous messages. */
	OFPT_PACKET_IN    /* Async message */
	OFPT_FLOW_REMOVED /* Async message */
	OFPT_PORT_STATUS  /* Async message */

	/* Controller command messages. */
	OFPT_PACKET_OUT /* Controller/switch message */
	OFPT_FLOW_MOD   /* Controller/switch message */
	OFPT_PORT_MOD   /* Controller/switch message */

	/* Statistics messages. */
	OFPT_STATS_REQUEST /* Controller/switch message */
	OFPT_STATS_REPLY   /* Controller/switch message */

	/* Barrier messages. */
	OFPT_BARRIER_REQUEST /* Controller/switch message */
	OFPT_BARRIER_REPLY   /* Controller/switch message */

	/* Queue Configuration messages. */
	OFPT_QUEUE_GET_CONFIG_REQUEST /* Controller/switch message */
	OFPT_QUEUE_GET_CONFIG_REPLY   /* Controller/switch message */
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

