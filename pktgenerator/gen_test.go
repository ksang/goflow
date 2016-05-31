package pktgenerator

import (
	"fmt"
	"testing"
)

func TestOpenFlow(t *testing.T) {
	server := NewTcpServer(":6633")
	err := server.Start()
	if err != nil {
		fmt.Print(err)
		return
	}
	ofpkt, err := NewEchoRequestPkt("127.0.0.1:6633")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Sending OF packets.")
	err = SendMany(10, &ofpkt)
	if err != nil {
		fmt.Print(err)
	}
	server.Stop()

}
