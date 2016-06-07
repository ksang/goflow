package v10

import(
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)

type config struct {
	flags          uint16
	missSendLength uint16
}

func (c *config) Flags() uint16 {
	return c.flags
}

func (c *config) SetFlags(flags uint16) {
	c.flags = flags
}

func (c *config) MissSendLength() uint16 {
	return c.missSendLength
}

func (c *config) SetMissSendLength(length uint16) {
	c.missSendLength = length
}

type setConfig struct {
	openflow.Message
	config
}

func (s *setConfig) MarshalBinary() ([]byte, error) {
	v := make([]byte, 4)
	binary.BigEndian.PutUint16(v[0:2], s.flags)
	binary.BigEndian.PutUint16(v[2:4], s.missSendLength)
	s.SetPayload(v)

	return s.Message.MarshalBinary()
}

func (s *setConfig) UnmarshalBinary(data []byte) error {
	if err := s.Message.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := s.Payload()
	if payload == nil || len(payload) != 4 {
		return openflow.ErrInvalidPacketLength
	}
	s.flags = binary.BigEndian.Uint16(payload[0:2])
	s.missSendLength = binary.BigEndian.Uint16(payload[2:4])
	return nil	
}

func NewSetConfig(xid uint32) openflow.SetConfig {
	return &setConfig{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_SET_CONFIG, xid),
		config: config{
			flags:          OFPC_FRAG_NORMAL,
			missSendLength: 0xFFFF,
		},
	}
}

type getConfigRequest struct {
	openflow.Message
}

func NewGetConfigRequest(xid uint32) openflow.GetConfigRequest {
	return &getConfigRequest{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_GET_CONFIG_REQUEST, xid),
	}
}

type getConfigReply struct {
	setConfig
}

func NewGetConfigReply(xid uint32) openflow.GetConfigReply {
	return &setConfig{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_GET_CONFIG_REPLY, xid),
		config: config{},
	}
}


