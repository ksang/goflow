package openflow

import (
	"errors"
)

var (
	ErrInvalidPacketLength   = errors.New("invalid packet length")
	ErrUnsupportedVersion    = errors.New("unsupported protocol version")
	ErrUnsupportedMessage    = errors.New("unsupported message type")
	ErrInvalidMACAddress     = errors.New("invalid MAC address")
	ErrInvalidIPAddress      = errors.New("invalid IP address")
	ErrInvalidVlanID		 = errors.New("invalid vlan id")
	ErrUnsupportedIPProtocol = errors.New("unsupported IP protocol")
	ErrUnsupportedEtherType  = errors.New("unsupported Ethernet type")
	ErrMissingIPProtocol     = errors.New("missing IP protocol")
	ErrMissingEtherType      = errors.New("missing Ethernet type")
	ErrUnsupportedMatchType  = errors.New("unsupported flow match type")
	ErrNoDataProvided        = errors.New("no data provided")
	ErrInvalidDataLength 	 = errors.New("invalid data length")
	ErrInvalidValueProvided  = errors.New("invalid value provided")
)
