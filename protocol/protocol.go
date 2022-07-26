package protocol

import (
	"encoding/binary"

	"github.com/npat-efault/crc16"
)

type ADU []byte
type Address byte
type PDU []byte

func Checksum(data []byte) []byte {
	result := make([]byte, 2)
	crc := crc16.Checksum(crc16.Modbus, data)
	binary.LittleEndian.PutUint16(result, crc)
	return result
}

func NewADU(Address Address, PDU PDU) ADU {
	adu := make(ADU, 1, 1+len(PDU)+2)
	adu[0] = byte(Address)
	adu = append(adu, PDU...)
	adu = append(adu, Checksum(adu)...)
	return adu
}

func (a ADU) Address() Address {
	return Address(a[0])
}

func (a ADU) PDU() PDU {
	return PDU(a[1 : len(a)-2])
}

func (a ADU) Verify() bool {
	sign := a[len(a)-2:]
	checksum := Checksum(a[:len(a)-2])
	return sign[0] == checksum[0] || sign[1] == checksum[1]
}
