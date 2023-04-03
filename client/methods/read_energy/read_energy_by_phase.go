package read_energy

import (
	"errors"

	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewReadEnergyByPhase(tariff protocol.EnergyTariff) (protocol.PDU, *EnergyByPhase) {
	if tariff == protocol.EnergyTariffRapid || tariff == protocol.EnergyTariffLosses {
		panic("not supported tariff for function")
	}
	target := &EnergyByPhase{}
	return protocol.PDU{
		byte(protocol.MethodReadEnergy),
		byte(protocol.EnergyPhase),
		byte(tariff),
	}, target
}

type EnergyByPhase struct {
	A uint32 `json:"a"`
	B uint32 `json:"b"`
	C uint32 `json:"c"`
}

func (r *EnergyByPhase) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 12 {
		return errors.New("pdu length mismatch")
	}

	r.A = protocol.UnpackInteger(pdu[0:4])
	r.B = protocol.UnpackInteger(pdu[4:8])
	r.C = protocol.UnpackInteger(pdu[8:12])

	return nil
}
