package read_parameter

import (
	"errors"
	"fmt"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewIndividualOptions() (protocol.PDU, *IndividualOptions) {
	target := &IndividualOptions{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterIndividualOptions),
	}, target
}

type IndividualOptions struct {
	SerialNumber string `json:"serial_number"`
	Date         string `json:"date"`
	Firmware     string `json:"firmware"`
	Modification string `json:"modification"`
}

func (r *IndividualOptions) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 16 {
		return errors.New("pdu length mismatch")
	}

	r.SerialNumber = fmt.Sprintf("%02d%02d%02d%02d", pdu[0], pdu[1], pdu[2], pdu[3])
	r.Date = fmt.Sprintf("20%02d-%02d-%02d", pdu[6], pdu[5], pdu[4])
	r.Firmware = fmt.Sprintf("%d.%d.%d", pdu[7], pdu[8], pdu[9])
	r.Modification = fmt.Sprintf("%X", pdu[10:])

	return nil
}
