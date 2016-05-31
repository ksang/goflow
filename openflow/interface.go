package openflow

import (
	"encoding"
)

// Echo message interface
type EchoDecoder interface {
	MessageDecoder
	Data() []byte
	SetData(data []byte) error
	Error() error
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
