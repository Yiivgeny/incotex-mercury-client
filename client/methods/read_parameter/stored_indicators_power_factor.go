package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewStoredIndicatorsPowerFactor() (protocol.PDU, *PowerFactorWithSum) {
	target := &PowerFactorWithSum{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterStoredIndicators),
		byte(protocol.BWRIModePowerFactor) | byte(protocol.BWRIPhaseA),
	}, target
}

type PowerFactorWithSum struct {
	A   float32 `json:"a"`
	B   float32 `json:"b"`
	C   float32 `json:"c"`
	Sum float32 `json:"sum"`
}

func (r *PowerFactorWithSum) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 12 {
		return errors.New("pdu length mismatch")
	}

	r.Sum = float32(protocol.UnpackSignedPower(pdu[0:3], 0, 0)) / 1000
	r.A = float32(protocol.UnpackSignedPower(pdu[3:6], 0, 0)) / 1000
	r.B = float32(protocol.UnpackSignedPower(pdu[6:9], 0, 0)) / 1000
	r.C = float32(protocol.UnpackSignedPower(pdu[9:12], 0, 0)) / 1000

	return nil
}
