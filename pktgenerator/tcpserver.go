package pktgenerator

import (
	"encoding/hex"
	"log"
	"net"
)

type TcpServer struct {
	laddr    string
	stopping chan chan error
	conn     chan net.Conn
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1500)
	len, err := conn.Read(buf)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("Packet Received: \n", hex.Dump(buf[:len]), "\n")
}

func NewTcpServer(laddr string) *TcpServer {
	return &TcpServer{
		laddr:    laddr,
		stopping: make(chan chan error),
		conn:     make(chan net.Conn),
	}
}

func (t *TcpServer) Loop(ln net.Listener, ready chan bool) {
	var err error
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			t.conn <- conn
		}
	}()
	ready <- true
	for {
		select {
		case errc := <-t.stopping:
			errc <- err
			break
		case conn := <-t.conn:
			go handleConnection(conn)
		}
	}
}

func (t *TcpServer) Start() error {
	ready := make(chan bool)
	ln, err := net.Listen("tcp", t.laddr)
	if err != nil {
		log.Println(err)
		return err
	}
	go t.Loop(ln, ready)
	<-ready
	return nil
}

func (t *TcpServer) Stop() error {
	errc := make(chan error)
	t.stopping <- errc
	return <-errc
}
