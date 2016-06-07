package v10

import(
	"encoding/binary"
	"github.com/ksang/goflow/openflow"
)

type vendor struct {
	openflow.Message
	vendorID uint32
	data []byte
}

func (v *vendor) VendorID() uint32 {
	return v.vendorID
}

func (v *vendor) SetVendorID(vid uint32) {
	v.vendorID = vid
}

func (v *vendor) Data() []byte {
	return v.data
}

func (v *vendor) SetData(data []byte) error {
	if data == nil {
		return openflow.ErrNoDataProvided
	}
	v.data = data
	return nil
}

func (v *vendor) MarshalBinary() ([]byte, error) {
	tmp := make([]byte, 4 + len(v.data))
	binary.BigEndian.PutUint32(tmp[0:4], v.vendorID)
	if len(v.data) > 0{
		copy(tmp[4:], v.data)
	}
	v.SetPayload(tmp)
	return v.Message.MarshalBinary()
}

func (v *vendor) UnmarshalBinary(data []byte) error {
	if err := v.Message.UnmarshalBinary(data); err != nil {
		return err
	}

	payload := v.Payload()
	if payload == nil || len(payload) < 4 {
		return openflow.ErrInvalidPacketLength
	}
	v.vendorID = binary.BigEndian.Uint32(payload[0:4])
	if len(payload) > 4 {
		v.data = payload[4:]
	}
	return nil
}

func NewVendor(xid uint32) openflow.Vendor {
	return &vendor{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_VENDOR, xid),
	}
}