package read_parameter

import (
	"errors"
	"fmt"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewSerialNumberAndBuildDate() (protocol.PDU, *SerialNumberAndBuildDate) {
	target := &SerialNumberAndBuildDate{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterSerialNumberAndBuildDate),
	}, target
}

type SerialNumberAndBuildDate struct {
	SerialNumber string `json:"serial_number"`
	Date         string `json:"date"`
}

func (r *SerialNumberAndBuildDate) GetRequest() []byte {
	return []byte{
		byte(protocol.ParameterSerialNumberAndBuildDate),
	}
}

func (r *SerialNumberAndBuildDate) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 7 {
		return errors.New("pdu length mismatch")
	}

	r.SerialNumber = fmt.Sprintf("%02d%02d%02d%02d", pdu[0], pdu[1], pdu[2], pdu[3])
	r.Date = fmt.Sprintf("20%02d-%02d-%02d", pdu[6], pdu[5], pdu[4])

	return nil
}
