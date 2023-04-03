package read_energy

import (
	"errors"

	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewReadEnergyRapid(registry protocol.EnergyRegistry) (protocol.PDU, *EnergyRapid) {
	if registry == protocol.EnergyPhase {
		panic("not supported registry for function")
	}
	target := &EnergyRapid{}
	return protocol.PDU{
		byte(protocol.MethodReadEnergy),
		byte(registry),
		byte(protocol.EnergyTariffRapid),
	}, target
}

type EnergyRapid struct {
	Tariff1 *Energy `json:"tariff1,omitempty"`
	Tariff2 *Energy `json:"tariff2,omitempty"`
	Tariff3 *Energy `json:"tariff3,omitempty"`
	Tariff4 *Energy `json:"tariff4,omitempty"`
	Total   *Energy `json:"total,omitempty"`
	Losses  *Energy `json:"losses,omitempty"` // TODO: Работа со функцией учета потерь
}

func (r *EnergyRapid) Unmarshall(pdu protocol.PDU) error {
	// TODO: Работа с однофазными
	if len(pdu) != 80 {
		return errors.New("pdu length mismatch")
	}

	r.Tariff1 = &Energy{}
	if err := r.Tariff1.Unmarshall(pdu[0:16]); err != nil {
		return err
	}
	r.Tariff2 = &Energy{}
	if err := r.Tariff2.Unmarshall(pdu[16:32]); err != nil {
		return err
	}
	r.Tariff3 = &Energy{}
	if err := r.Tariff3.Unmarshall(pdu[32:48]); err != nil {
		return err
	}
	r.Tariff4 = &Energy{}
	if err := r.Tariff4.Unmarshall(pdu[48:64]); err != nil {
		return err
	}
	r.Total = &Energy{}
	if err := r.Total.Unmarshall(pdu[64:80]); err != nil {
		return err
	}

	return nil
}
