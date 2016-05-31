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
	ErrUnsupportedIPProtocol = errors.New("unsupported IP protocol")
	ErrUnsupportedEtherType  = errors.New("unsupported Ethernet type")
	ErrMissingIPProtocol     = errors.New("missing IP protocol")
	ErrMissingEtherType      = errors.New("missing Ethernet type")
	ErrUnsupportedMatchType  = errors.New("unsupported flow match type")
	ErrNoDataProvided        = errors.New("No data provided")
)