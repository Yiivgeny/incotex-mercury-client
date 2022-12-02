package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewInstantIndicatorsTemperature() (protocol.PDU, *Temperature) {
	target := &Temperature{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterInstantIndicators),
		byte(protocol.BWRIModeTemperature),
	}, target
}

type Temperature struct {
	Temperature int16 `json:"temperature"`
}

func (r *Temperature) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 2 {
		return errors.New("pdu length mismatch")
	}

	r.Temperature = protocol.UnpackNormalShort(pdu[0:2])

	return nil
}
