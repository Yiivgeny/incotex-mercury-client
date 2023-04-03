package read_energy

import (
	"errors"

	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewReadEnergy(registry protocol.EnergyRegistry, tariff protocol.EnergyTariff) (protocol.PDU, *Energy) {
	if registry == protocol.EnergyPhase {
		panic("not supported registry for function")
	}
	if tariff == protocol.EnergyTariffRapid || tariff == protocol.EnergyTariffLosses {
		panic("not supported tariff for function")
	}
	target := &Energy{}
	return protocol.PDU{
		byte(protocol.MethodReadEnergy),
		byte(registry),
		byte(tariff),
	}, target
}

type Energy struct {
	ActiveDirect    uint32 `json:"active_direct"`
	ReactiveDirect  uint32 `json:"reactive_direct"`
	ActiveReverse   uint32 `json:"active_reverse"`
	ReactiveReverse uint32 `json:"reactive_reverse"`
}

func (r *Energy) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 16 {
		return errors.New("pdu length mismatch")
	}

	r.ActiveDirect = protocol.UnpackInteger(pdu[0:4])
	r.ActiveReverse = protocol.UnpackInteger(pdu[4:8])
	r.ReactiveDirect = protocol.UnpackInteger(pdu[8:12])
	r.ReactiveReverse = protocol.UnpackInteger(pdu[12:16])

	return nil
}
