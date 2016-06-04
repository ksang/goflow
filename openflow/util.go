package openflow

import (
	"net"
	"encoding/binary"
)

// Not used
func FlipMacEndian(mac net.HardwareAddr) (net.HardwareAddr, error) {
	if len(mac) != 6{
		return nil, ErrInvalidDataLength
	}
	// pad it to uint64
	mac = append(mac, 0x00, 0x00)
	macuint := binary.BigEndian.Uint64(mac)
	res := make([]byte, 8)
	binary.LittleEndian.PutUint64(res,macuint)
	return res[2:], nil
}

func FlipIpEndian(ip net.IP) (net.IP, error) {
	if len(ip) != 4{
		return nil, ErrInvalidDataLength
	}
	ipuint := binary.BigEndian.Uint32(ip)
	res := make([]byte, 4)
	binary.LittleEndian.PutUint32(res,ipuint)
	return res, nil
}