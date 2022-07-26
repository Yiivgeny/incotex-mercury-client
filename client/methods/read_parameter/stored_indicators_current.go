package read_parameter

import (
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewStoredIndicatorsCurrent() (protocol.PDU, *Current) {
	target := &Current{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterStoredIndicators),
		byte(protocol.BWRIModeCurrent) | byte(protocol.BWRIPhaseA),
	}, target
}
