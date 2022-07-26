package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewInstantIndicatorsFrequency() (protocol.PDU, *Frequency) {
	target := &Frequency{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterInstantIndicators),
		byte(protocol.BWRIModeFrequency),
	}, target
}

type Frequency struct {
	Frequency float32 `json:"frequency"`
}

func (r *Frequency) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 3 {
		return errors.New("pdu length mismatch")
	}

	r.Frequency = float32(protocol.UnpackInteger(pdu[0:3])) / 100

	return nil
}
