package read_parameter

import (
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewStoredIndicatorsVoltage() (protocol.PDU, *Voltage) {
	target := &Voltage{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterStoredIndicators),
		byte(protocol.BWRIModeVoltage) | byte(protocol.BWRIPhaseA),
	}, target
}
