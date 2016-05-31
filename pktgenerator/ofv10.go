package pktgenerator

import (
	"github.com/ksang/goflow/openflow"
	"github.com/ksang/goflow/openflow/v10"
	"time"
)

type OpenFlowPkt struct {
	TCPPacket
	message openflow.Message
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
		message: msg,
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
		message: echo.Message,
	}, nil
}
