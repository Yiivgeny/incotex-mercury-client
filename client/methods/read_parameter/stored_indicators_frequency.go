package read_parameter

import (
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewStoredIndicatorsFrequency() (protocol.PDU, *Frequency) {
	target := &Frequency{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterStoredIndicators),
		byte(protocol.BWRIModeFrequency),
	}, target
}
