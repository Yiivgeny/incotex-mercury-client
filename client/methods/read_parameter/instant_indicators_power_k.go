package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewInstantIndicatorsPowerK() (protocol.PDU, *PowerK) {
	target := &PowerK{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterInstantIndicators),
		byte(protocol.BWRIModeKPower) | byte(protocol.BWRIPhaseA),
	}, target
}

type PowerK struct {
	A float32 `json:"a"`
	B float32 `json:"b"`
	C float32 `json:"c"`
}

func (r *PowerK) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 9 {
		return errors.New("pdu length mismatch")
	}

	r.A = protocol.SignedFloatDecode(pdu[0:3]) / 1000
	r.B = protocol.SignedFloatDecode(pdu[3:6]) / 1000
	r.C = protocol.SignedFloatDecode(pdu[6:9]) / 1000

	return nil
}
