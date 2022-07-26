package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewInstantIndicatorsCurrent() (protocol.PDU, *Current) {
	target := &Current{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterInstantIndicators),
		byte(protocol.BWRIModeCurrent) | byte(protocol.BWRIPhaseA),
	}, target
}

type Current struct {
	A float32 `json:"a"`
	B float32 `json:"b"`
	C float32 `json:"c"`
}

func (r *Current) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 9 {
		return errors.New("pdu length mismatch")
	}

	r.A = float32(protocol.UnpackInteger(pdu[0:3])) / 1000
	r.B = float32(protocol.UnpackInteger(pdu[3:6])) / 1000
	r.C = float32(protocol.UnpackInteger(pdu[6:9])) / 1000

	return nil
}
