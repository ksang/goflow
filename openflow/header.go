/*
Package openflow is an implementation of openflow protocol in go.
It supports multiple versions of openflow specifications.
For mode details about openflow, please visit Open Networking Foundation
https://opennetworking.org/sdn-resources/openflow
*/
package openflow

type HeaderDecoder interface {
	Version() uint8
	MsgType() uint8
	Length() uint16
	TransactionID() uint32
	SetTransactionID(xid uint32)
}

// Openflow header definition
type Header struct {
	version uint8
	msgType uint8
	length  uint16
	xid     uint32
}
