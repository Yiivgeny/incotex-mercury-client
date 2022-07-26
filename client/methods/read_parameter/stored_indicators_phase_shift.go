package read_parameter

import (
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewStoredIndicatorsPhaseShift() (protocol.PDU, *PhaseShift) {
	target := &PhaseShift{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterStoredIndicators),
		byte(protocol.BWRIModePhaseShift) | byte(protocol.BWRIPhaseA),
	}, target
}
