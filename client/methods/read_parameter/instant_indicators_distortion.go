package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewInstantIndicatorsDistortion() (protocol.PDU, *Distortion) {
	target := &Distortion{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterInstantIndicators),
		byte(protocol.BWRIModeDistortion) | byte(protocol.BWRIPhaseA),
	}, target
}

type Distortion struct {
	A float32 `json:"a"`
	B float32 `json:"b"`
	C float32 `json:"c"`
}

func (r *Distortion) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 6 {
		return errors.New("pdu length mismatch")
	}

	r.A = float32(protocol.UnpackInteger(pdu[0:2])) / 100
	r.B = float32(protocol.UnpackInteger(pdu[2:4])) / 100
	r.C = float32(protocol.UnpackInteger(pdu[4:6])) / 100

	return nil
}
