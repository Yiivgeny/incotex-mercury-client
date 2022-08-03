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

	r.Sum = protocol.SignedFloatDecode(pdu[0:3]) / 1000
	r.A = protocol.SignedFloatDecode(pdu[3:6]) / 1000
	r.B = protocol.SignedFloatDecode(pdu[6:9]) / 1000
	r.C = protocol.SignedFloatDecode(pdu[9:12]) / 1000

	return nil
}