package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewInstantIndicatorsPowerFactor() (protocol.PDU, *PowerFactor) {
	target := &PowerFactor{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterInstantIndicators),
		byte(protocol.BWRIModePowerFactor) | byte(protocol.BWRIPhaseA),
	}, target
}

type PowerFactor struct {
	A float32 `json:"a"`
	B float32 `json:"b"`
	C float32 `json:"c"`
}

func (r *PowerFactor) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 9 {
		return errors.New("pdu length mismatch")
	}

	r.A = float32(protocol.UnpackSignedPower(pdu[0:3], 0, 0)) / 1000
	r.B = float32(protocol.UnpackSignedPower(pdu[3:6], 0, 0)) / 1000
	r.C = float32(protocol.UnpackSignedPower(pdu[6:9], 0, 0))/ 1000

	return nil
}
