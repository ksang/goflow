package v10

import(
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)

type errorMessage struct {
	openflow.Message
	typ uint16
	code uint16
	data []byte
}

func (e *errorMessage) Type() uint16 {
	return e.typ
}

func (e *errorMessage) SetType(typ uint16) error {
	if typ > 0x5 {
		return openflow.ErrInvalidValueProvided
	}
	e.typ = typ
	return nil
}

func (e *errorMessage) Code() uint16 {
	return e.code
}

func (e *errorMessage) SetCode(code uint16) error {
	// TODO: not checking error code accurately
	if code > 0x8 {
		return openflow.ErrInvalidValueProvided
	}
	e.code = code
	return nil
}

// Implement EchoDecoder interface
func (e *errorMessage) Data() []byte {
	return e.data
}

func (e *errorMessage) SetData(data []byte) error {
	if data == nil {
		return openflow.ErrNoDataProvided
	}
	e.data = data
	return nil
}

func (e *errorMessage) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4 + len(e.data))
	binary.BigEndian.PutUint16(v[0:2], e.typ)
	binary.BigEndian.PutUint16(v[2:4], e.code)
	if len(e.data) > 0{
		copy(v[4:], e.data)
	}
	e.SetPayload(v)
	return e.Message.MarshalBinary()
}

func (e *errorMessage) UnmarshalBinary(data []byte) error {
	if err := e.Message.UnmarshalBinary(data); err != nil {
		return err
	}

	payload := e.Payload()
	if payload == nil || len(payload) < 4 {
		return openflow.ErrInvalidPacketLength
	}
	e.typ = binary.BigEndian.Uint16(payload[0:2])
	e.code = binary.BigEndian.Uint16(payload[2:4])
	if len(payload) > 4 {
		e.data = payload[4:]
	}
	return nil
}

func NewError(xid uint32) openflow.Error {
	return &errorMessage{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_ERROR, xid),
	}
}
