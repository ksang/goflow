package pktgenerator

import (
	"github.com/ksang/goflow/openflow"
	"github.com/ksang/goflow/openflow/v10"
	"time"
)

type OpenFlowPkt struct {
	TCPPacket
}

func NewOpenFlowPkt(dst string,
	version uint8,
	msgType uint8,
	xid uint32,
) (OpenFlowPkt, error) {
	msg := openflow.NewMessage(version, msgType, xid)
	data, err := msg.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return OpenFlowPkt{
		TCPPacket: TCPPacket{
			dst:     dst,
			payload: data,
		},
	}, nil
}

func CreateOFPkt(dst string, data []byte) OpenFlowPkt {
	return OpenFlowPkt{
		TCPPacket: TCPPacket{
			dst:     dst,
			payload: data,
		},
	}
}

func NewHelloPkt(dst string) (OpenFlowPkt, error) {
	hello := v10.NewHello(uint32(23333))
	hello.SetData([]byte(time.Now().Format(time.UnixDate)))
	data, err := hello.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewErrorPkt(dst string) (OpenFlowPkt, error) {
	error := v10.NewError(uint32(23333))
	error.SetType(uint16(1))
	error.SetCode(uint16(2))
	error.SetData([]byte(time.Now().Format(time.UnixDate)))
	data, err := error.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewEchoRequestPkt(dst string) (OpenFlowPkt, error) {
	echo := v10.NewEchoRequest(uint32(23333))
	echo.SetData([]byte(time.Now().Format(time.UnixDate)))
	data, err := echo.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewPacketInPkt(dst string) (OpenFlowPkt, error) {
	packetIn := v10.NewPacketIn(uint32(23334))
	packetIn.SetBufferID(uint32(23333))
	packetIn.SetInPort(uint16(123))
	packetIn.SetReason(uint8(2))
	packetIn.SetData([]byte(time.Now().Format(time.UnixDate)))
	data, err := packetIn.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewPacketOutPkt(dst string) (OpenFlowPkt, error) {
	packetOut := v10.NewPacketOut(uint32(23335))
	packetOut.SetBufferID(uint32(23333))
	packetOut.SetInPort(uint16(321))
	actOut := v10.NewActionOutput()
	actOut.SetPort(uint16(555))
	actOut.SetMaxLen(uint16(65535))

	actSetDL := v10.NewActionSetDLSrc()
	actSetDL.SetDLSrc([]byte{0xa8, 0x66, 0x7f, 0x33, 0x44, 0x55})

	actSetNW := v10.NewActionSetNWSrc()
	actSetNW.SetNWSrc([]byte{0x1, 0x2, 0x3, 0x4})

	actSetTP := v10.NewActionSetTPSrc()
	actSetTP.SetPort(uint16(123))
	//packetOut.AddAction(actOut)
	packetOut.AddAction(actSetTP)
	//packetOut.SetData([]byte(time.Now().Format(time.UnixDate)))
	data, err := packetOut.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewFeatureRequestPkt(dst string) (OpenFlowPkt, error) {
	feature := v10.NewFeatureRequest(uint32(23334))
	data, err := feature.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewFeatureReplyPkt(dst string) (OpenFlowPkt, error) {
	feature := v10.NewFeatureReply(uint32(23334))
	feature.SetDPID(uint64(111111))
	feature.SetNumBuffers(uint32(8))
	feature.SetNumTables(uint8(6))
	feature.SetCapabilities(openflow.FLOW_STATS | openflow.TABLE_STATS)
	feature.SetActions(openflow.OUTPUT | openflow.ENQUEUE)
	p1, err := v10.NewPort(openflow.PortID(1), []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}, "port no.1")
	p2, err := v10.NewPort(openflow.PortID(2), []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66}, "port no.2")
	if err != nil {
		return OpenFlowPkt{}, err
	}
	feature.AddPort(p1)
	feature.AddPort(p2)
	data, err := feature.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewGetConfigReplyPkt(dst string) (OpenFlowPkt, error) {
	config := v10.NewGetConfigReply(uint32(23334))
	config.SetFlags(v10.OFPC_FRAG_NORMAL)
	config.SetMissSendLength(0x10)
	data, err := config.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewFlowRemovedPkt(dst string) (OpenFlowPkt, error) {
	fr := v10.NewFlowRemoved(uint32(23334))
	m := v10.NewMatch()
	m.SetInPort(uint16(1))
	m.SetDLSrc([]byte{0x1, 0x1, 0x1, 0x1, 0x1, 0x1})
	m.SetDLDst([]byte{0x2, 0x2, 0x2, 0x2, 0x2, 0x2})
	m.SetNWSrc([]byte{0x1, 0x1, 0x1, 0x1})
	m.SetNWDst([]byte{0x2, 0x2, 0x2, 0x2})
	m.SetDLVlan(uint16(4096))
	m.SetWildcardNWSrc(24)
	fr.SetMatch(m)
	fr.SetCookie(uint64(111))
	fr.SetPacketCount(uint64(65535))
	data, err := fr.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewPortStatusPkt(dst string) (OpenFlowPkt, error) {
	ps := v10.NewPortStatus(uint32(23334))
	p1, err := v10.NewPort(openflow.PortID(1), []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}, "port no.1")
	ps.SetPort(p1)
	ps.SetReason(openflow.PortModified)
	data, err := ps.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewFlowModPkt(dst string) (OpenFlowPkt, error) {
	fm := v10.NewFlowMod(uint32(23334))
	m := v10.NewMatch()
	m.SetInPort(uint16(1))
	m.SetDLSrc([]byte{0x1, 0x1, 0x1, 0x1, 0x1, 0x1})
	m.SetDLDst([]byte{0x2, 0x2, 0x2, 0x2, 0x2, 0x2})
	m.SetNWSrc([]byte{0x1, 0x1, 0x1, 0x1})
	m.SetNWDst([]byte{0x2, 0x2, 0x2, 0x2})
	m.SetDLVlan(uint16(4096))
	m.SetWildcardNWSrc(24)
	fm.SetMatch(m)
	fm.SetCookie(uint64(111))
	fm.SetIdleTimeout(uint16(65535))
	actOut := v10.NewActionOutput()
	actOut.SetPort(uint16(555))
	actOut.SetMaxLen(uint16(65535))
	fm.SetAction(actOut)
	fm.SetFlags(openflow.CheckOverlap)
	data, err := fm.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewPortModPkt(dst string) (OpenFlowPkt, error) {
	pm := v10.NewPortMod(uint32(23334))
	pm.SetPort(openflow.Flood)
	pm.SetHWAddr([]byte{0x1, 0x1, 0x1, 0x1, 0x1, 0x1})
	pm.SetConfig(openflow.NoRecvSTP)
	pm.SetMask(0x7f)
	pm.SetAdvertise(openflow.FD_10GB|openflow.Fiber)
	data, err := pm.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewBarrierRequestPkt(dst string) (OpenFlowPkt, error) {
	bq := v10.NewBarrierRequest(uint32(1234))
	data, _ := bq.MarshalBinary()
	return CreateOFPkt(dst, data), nil
}

func NewStatsRequestFlowPkt(dst string) (OpenFlowPkt, error) {
	srf := v10.NewStatsReuqestFlow(uint32(23334))
	data, err := srf.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewQueueGetConfigRequestPkt(dst string) (OpenFlowPkt, error) {
	qgr := v10.NewQueueGetConfigRequest(uint32(23334))
	qgr.SetPort(uint16(233))
	data, err := qgr.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}

func NewQueueGetConfigReplyPkt(dst string) (OpenFlowPkt, error) {
	qgr := v10.NewQueueGetConfigReply(uint32(23334))
	qgr.SetPort(uint16(233))
	p1 := v10.NewQueue()
	p1.SetQueueID(uint32(111))
	p1.SetRate(uint16(111))
	p2 := v10.NewQueue()
	p2.SetQueueID(uint32(222))
	p2.SetRate(uint16(222))
	qgr.AddQueue(p1)
	qgr.AddQueue(p2)
	data, err := qgr.MarshalBinary()
	if err != nil {
		return OpenFlowPkt{}, err
	}
	return CreateOFPkt(dst, data), nil
}