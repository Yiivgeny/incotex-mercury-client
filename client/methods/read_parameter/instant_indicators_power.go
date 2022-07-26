package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewInstantIndicatorsPower(power protocol.BWRIPower) (protocol.PDU, *PowerInstant) {
	target := &PowerInstant{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterInstantIndicators),
		byte(protocol.BWRIModePower) | byte(power) | byte(protocol.BWRIPhaseAll),
	}, target
}

type PowerInstant struct {
	A   float32 `json:"a"`
	B   float32 `json:"b"`
	C   float32 `json:"c"`
	Sum float32 `json:"sum"`
}

func (r *PowerInstant) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 12 {
		return errors.New("pdu length mismatch")
	}

	r.Sum = protocol.SignedFloatDecode(pdu[0:3])
	r.A = protocol.SignedFloatDecode(pdu[3:6])
	r.B = protocol.SignedFloatDecode(pdu[6:9])
	r.C = protocol.SignedFloatDecode(pdu[9:12])

	return nil
}
