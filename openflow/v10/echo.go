package v10

import(
	"github.com/ksang/goflow/openflow"
)

type Echo struct {
	openflow.Message
	err error
	data []byte
}

// Implement EchoDecoder interface
func (e *Echo) Data() []byte {
	return e.data
}

func (e *Echo) SetData(data []byte) error {
	if data == nil {
		return openflow.ErrNoDataProvided
	}
	e.data = data
	return nil
}

func (e *Echo) Error() error {
	return e.err
}

func (e *Echo) MarshalBinary() ([]byte, error) {
	if e.err != nil {
		return nil, e.err
	}
	e.SetPayload(e.data)
	return e.Message.MarshalBinary()
}

func (e *Echo) UnmarshalBinary(data []byte) error {
	if err := e.Message.UnmarshalBinary(data); err != nil {
		return err
	}
	e.data = e.Payload()
	return nil
}

func NewEchoRequest(xid uint32) openflow.EchoDecoder {
	return &Echo{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_ECHO_REQUEST, xid),
	}
}

func NewEchoReply(xid uint32) openflow.EchoDecoder {
	return &Echo{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_ECHO_REPLY, xid),
	}
}