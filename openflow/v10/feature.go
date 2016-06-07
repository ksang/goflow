package v10

import(
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)

type featureRequest struct {
	openflow.Message
}

func (f *featureRequest) MarshalBinary() ([]byte, error) {
	return f.Message.MarshalBinary()
}

func (f *featureRequest) UnmarshalBinary(data []byte) error {
	return f.Message.UnmarshalBinary(data)
}

func NewFeatureRequest(xid uint32) openflow.FeatureRequest {
	return &featureRequest{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_FEATURES_REQUEST, xid),
	}
}

type featureReply struct {
	openflow.Message
	dpid uint64
	numBuffers uint32
	numTables uint8
	capabilities uint32
	actions uint32
	ports []openflow.Port
}

func (f *featureReply) DPID() uint64 {
	return f.dpid
}

func (f *featureReply) SetDPID(dpid uint64) {
	f.dpid = dpid
} 

func (f *featureReply) NumBuffers() uint32 {
	return f.numBuffers
}

func (f *featureReply) SetNumBuffers(nb uint32) {
	f.numBuffers = nb
}

func (f *featureReply) NumTables() uint8 {
	return f.numTables
}

func (f *featureReply) SetNumTables(nt uint8) {
	f.numTables = nt
}

func (f *featureReply) Capabilities() uint32 {
	return f.capabilities
}

func (f *featureReply) SetCapabilities(c uint32) {
	f.capabilities = c
}

func (f *featureReply) Actions() uint32 {
	return f.actions
}

func (f *featureReply) SetActions(a uint32) {
	f.actions = a
}

func (f *featureReply) Ports() []openflow.Port {
	return f.ports
}

func (f *featureReply) AddPort(p openflow.Port) {
	f.ports = append(f.ports, p)
}

func (f *featureReply) MarshalBinary() ([]byte, error) {
	// lengh = feature reply fields + len(port)*portSize
	packetOutLen := 24 + len(f.ports)*48
	v := make([]byte, packetOutLen)
	binary.BigEndian.PutUint64(v[0:8], f.dpid)
	binary.BigEndian.PutUint32(v[8:12], f.numBuffers)
	v[12] = f.numTables
	// v[13:16] is padding
	binary.BigEndian.PutUint32(v[16:20], f.capabilities)
	binary.BigEndian.PutUint32(v[20:24], f.actions)
	// Marshal ports
	pos := 24
	for _, p := range f.ports {
		portPayload, err := p.MarshalBinary()
		if err != nil {
			return nil, err
		}
		copy(v[pos:pos+48], portPayload)
		pos += 48
	}
	// Call Message's SetPayload method
	f.SetPayload(v)
	return f.Message.MarshalBinary()
}

func (f *featureReply) UnmarshalBinary(data []byte) error {
	if err := f.Message.UnmarshalBinary(data); err != nil {
		return err
	}
	payload := f.Payload()
	if payload == nil || len(payload) < 24 || (len(payload)-24) % 48 != 0{
		return openflow.ErrInvalidPacketLength
	}
	f.dpid = binary.BigEndian.Uint64(payload[0:8])
	f.numBuffers = binary.BigEndian.Uint32(payload[8:12])
	f.numTables = payload[12]
	// payload[13:16] is padding
	f.capabilities = binary.BigEndian.Uint32(payload[16:20])
	f.actions = binary.BigEndian.Uint32(payload[20:24])
	for pos := 24; pos < len(payload); pos += 48{
		port := NewEmptyPort()
		if err := port.UnmarshalBinary(payload[pos:pos+48]); err != nil {
			return err
		}
	} 
	return nil
}

func NewFeatureReply(xid uint32) openflow.FeatureReply {
	return &featureReply{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_FEATURES_REPLY, xid),
	}
}
