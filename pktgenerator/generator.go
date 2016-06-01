/*
Package pktgenerator implements a package generator to test message format.
Using this tool we can test openflow packets.
*/
package pktgenerator

import (
	"errors"
	"net"
)

// Base interface for sending packets
type Generator interface {
	Send() error
}

type UDPPacket struct {
	dst     string
	payload []byte
}

type TCPPacket struct {
	dst     string
	payload []byte
}

func (pkt *UDPPacket) Send() error {
	if pkt.dst == "" {
		return errors.New("No target specified.")
	}
	raddr, err := net.ResolveUDPAddr("udp", pkt.dst)
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	defer conn.Close()
	_, err = conn.Write(pkt.payload)
	return err
}

func (pkt *TCPPacket) Send() error {
	if pkt.dst == "" {
		return errors.New("No target specified.")
	}
	raddr, err := net.ResolveTCPAddr("tcp", pkt.dst)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	defer conn.Close()
	_, err = conn.Write(pkt.payload)
	return err
}

// Send multiple identical packets.
func SendMany(number int, pkt Generator) error {
	for i := 0; i < number; i++ {
		if err := pkt.Send(); err != nil {
			return err
		}
	}

	return nil
}
