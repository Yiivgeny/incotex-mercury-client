package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewNetworkAddress() (protocol.PDU, *NetworkAddress) {
	target := &NetworkAddress{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterNetworkAddress),
	}, target
}

type NetworkAddress struct {
	Address protocol.Address `json:"address"`
}

func (r *NetworkAddress) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 2 {
		return errors.New("pdu length mismatch")
	}
	if pdu[0] != 0 {
		return errors.New("protocol mismatch")
	}

	r.Address = protocol.Address(pdu[1])

	return nil
}
