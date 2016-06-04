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

func NewEchoRequestPkt(dst string) (OpenFlowPkt, error) {
	echo := v10.NewEchoRequest(uint32(23333))
	echo.SetData([]byte(time.Now().Format(time.UnixDate)))
	data, err := echo.MarshalBinary()
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
	return OpenFlowPkt{
		TCPPacket: TCPPacket{
			dst:     dst,
			payload: data,
		},
	}, nil
}

func NewPacketOutPkt(dst string) (OpenFlowPkt, error) {
	packetOut := v10.NewPacketOut(uint32(23335))
	packetOut.SetBufferID(uint32(23333))
	packetOut.SetInPort(uint16(321))
	actOut := v10.NewActionOutput()
	actOut.SetPort(uint16(555))
	actOut.SetMaxLen(uint16(65535))

	actSetDL := v10.NewActionSetDLSrc()
	actSetDL.SetDLSrc([]byte{0xa8,0x66,0x7f,0x33,0x44,0x55})

	actSetNW := v10.NewActionSetNWSrc()
	actSetNW.SetNWSrc([]byte{0x1, 0x2, 0x3, 0x4})
	//packetOut.AddAction(actOut)
	packetOut.AddAction(actSetNW)
	//packetOut.SetData([]byte(time.Now().Format(time.UnixDate)))
	data, err := packetOut.MarshalBinary()
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