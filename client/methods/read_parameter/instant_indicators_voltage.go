package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewInstantIndicatorsVoltage() (protocol.PDU, *Voltage) {
	target := &Voltage{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterInstantIndicators),
		byte(protocol.BWRIModeVoltage) | byte(protocol.BWRIPhaseA),
	}, target
}

type Voltage struct {
	A float32 `json:"a"`
	B float32 `json:"b"`
	C float32 `json:"c"`
}

func (r *Voltage) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 9 {
		return errors.New("pdu length mismatch")
	}

	r.A = float32(protocol.UnpackInteger(pdu[0:3])) / 100
	r.B = float32(protocol.UnpackInteger(pdu[3:6])) / 100
	r.C = float32(protocol.UnpackInteger(pdu[6:9])) / 100

	return nil
}
