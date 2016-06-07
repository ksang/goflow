package pktgenerator

import (
	"fmt"
	"time"
	"testing"
)

func TestOpenFlow(t *testing.T) {
	server := NewTcpServer(":6633")
	if err := server.Start(); err != nil {
		fmt.Print(err)
		return
	}
	ofpkt, err := NewGetConfigReply("127.0.0.1:6633")
	if err != nil {
		fmt.Print(err)
		return 
	}
	fmt.Println("Sending OF packets.")
	if err = SendMany(1, &ofpkt); err != nil {
		fmt.Print(err)
	}
	time.Sleep(time.Second)
	server.Stop()

}
