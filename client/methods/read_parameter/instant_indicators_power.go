package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewInstantIndicatorsPower(power protocol.BWRIPower) (protocol.PDU, *PowerInstant) {
	active := 0
	reactive := 0

	if power == protocol.BWRIPowerP {
		active = 1
	}
	if power == protocol.BWRIPowerQ {
		reactive = 1
	}

	target := &PowerInstant{
		decoder: func(pdu []byte) float32 {
			return float32(protocol.UnpackSignedPower(pdu, active, reactive)) / 100
		},
	}
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

	decoder func([]byte) float32
}

func (r *PowerInstant) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 12 {
		return errors.New("pdu length mismatch")
	}

	r.Sum = r.decoder(pdu[0:3])
	r.A = r.decoder(pdu[3:6])
	r.B = r.decoder(pdu[6:9])
	r.C = r.decoder(pdu[9:12])

	return nil
}
