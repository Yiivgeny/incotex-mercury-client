package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewStoredIndicatorsPower(power protocol.BWRIPower) (protocol.PDU, *Power) {
	active := 0
	reactive := 0

	if power == protocol.BWRIPowerP {
		active = 1
	}
	if power == protocol.BWRIPowerQ {
		reactive = 1
	}

	target := &Power{
		decoder: func(pdu []byte) float32 {
			return float32(protocol.UnpackSignedPower(pdu, active, reactive)) / 100
		},
	}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterStoredIndicators),
		byte(protocol.BWRIModePower) | byte(power) | byte(protocol.BWRIPhaseAll),
	}, target
}

type Power struct {
	A   float32 `json:"a"`
	B   float32 `json:"b"`
	C   float32 `json:"c"`
	Sum float32 `json:"sum"`

	decoder func([]byte) float32
}

func (r *Power) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 16 {
		return errors.New("pdu length mismatch")
	}

	r.Sum = r.decoder(pdu[0:4])
	r.A = r.decoder(pdu[4:8])
	r.B = r.decoder(pdu[8:12])
	r.C = r.decoder(pdu[12:16])

	return nil
}
