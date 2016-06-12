package pktgenerator

import (
	"log"
	"testing"
	"time"
)

func TestOpenFlow(t *testing.T) {
	server := NewTcpServer(":6633")
	if err := server.Start(); err != nil {
		t.Error(err)
		return
	}
	ofpkt, err := NewQueueGetConfigReplyPkt("127.0.0.1:6633")
	if err != nil {
		t.Error(err)
		return
	}
	log.Println("Sending OF packets.")
	if err = SendMany(1, &ofpkt); err != nil {
		t.Error(err)
	}
	time.Sleep(time.Second)
	server.Stop()

}
