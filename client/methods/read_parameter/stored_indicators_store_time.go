package read_parameter

import (
	"errors"
	"fmt"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewStoredIndicatorsStoreTime() (protocol.PDU, *StoreDateTime) {
	target := &StoreDateTime{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterStoredIndicators),
		byte(protocol.BWRIModeDateTime),
	}, target
}

type StoreDateTime struct {
	DateTime string `json:"datetime"`
	Weekday  int    `json:"weekday"`
	IsWinter bool   `json:"is_winter"`
}

func (r *StoreDateTime) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 8 {
		return errors.New("pdu length mismatch")
	}

	r.DateTime = fmt.Sprintf(
		"20%02d-%02d-%02d %02d:%02d:%02d",
		protocol.Bdc2Int(pdu[6]),
		protocol.Bdc2Int(pdu[5]),
		protocol.Bdc2Int(pdu[4]),
		protocol.Bdc2Int(pdu[2]),
		protocol.Bdc2Int(pdu[1]),
		protocol.Bdc2Int(pdu[0]),
	)
	r.Weekday = int(pdu[3])
	r.IsWinter = pdu[7] == 0x01

	return nil
}
