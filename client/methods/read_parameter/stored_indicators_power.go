package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewStoredIndicatorsPower(power protocol.BWRIPower) (protocol.PDU, *Power) {
	target := &Power{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterStoredIndicators),
		byte(protocol.BWRIModePower) | byte(power) | byte(protocol.BWRIPhaseAll),
	}, target
}

type Power struct {
	A float32 `json:"a"`
	B float32 `json:"b"`
	C float32 `json:"c"`
	S float32 `json:"sum"`
}

func (r *Power) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 16 {
		return errors.New("pdu length mismatch")
	}

	r.S = protocol.SignedFloatDecode(pdu[0:4])
	r.A = protocol.SignedFloatDecode(pdu[4:8])
	r.B = protocol.SignedFloatDecode(pdu[8:12])
	r.C = protocol.SignedFloatDecode(pdu[12:16])

	return nil
}
