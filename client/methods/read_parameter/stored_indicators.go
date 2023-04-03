package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewStoredIndicators() (protocol.PDU, *StoredIndicators) {
	target := &StoredIndicators{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterStoredIndicators),
		byte(protocol.BWRIModeAccelerated),
	}, target
}

type StoredIndicators struct {
	DateTime    *StoreDateTime      `json:"date_time,omitempty"`
	Energy      *EnergyByTariff     `json:"energy,omitempty"`
	EnergyTotal *Energy             `json:"energy_total,omitempty"`
	PowerP      *Power              `json:"power_p,omitempty"`
	PowerQ      *Power              `json:"power_q,omitempty"`
	PowerS      *Power              `json:"power_s,omitempty"`
	Voltage     *Voltage            `json:"voltage,omitempty"`
	PhaseShift  *PhaseShift         `json:"phase_shift,omitempty"`
	Current     *Current            `json:"current,omitempty"`
	PowerK      *PowerFactorWithSum `json:"power_factor,omitempty"`
	Frequency   *Frequency          `json:"frequency,omitempty"`
}

func (r *StoredIndicators) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 178 {
		return errors.New("pdu length mismatch")
	}

	r.DateTime = &StoreDateTime{}
	if err := r.DateTime.Unmarshall(pdu[0:8]); err != nil {
		return err
	}
	r.Energy = &EnergyByTariff{}
	if err := r.Energy.Unmarshall(pdu[8:72]); err != nil {
		return err
	}
	r.EnergyTotal = &Energy{}
	if err := r.EnergyTotal.Unmarshall(pdu[72:88]); err != nil {
		return err
	}
	_, r.PowerP = NewStoredIndicatorsPower(protocol.BWRIPowerP)
	if err := r.PowerP.Unmarshall(pdu[88:104]); err != nil {
		return err
	}
	_, r.PowerQ = NewStoredIndicatorsPower(protocol.BWRIPowerQ)
	if err := r.PowerQ.Unmarshall(pdu[104:120]); err != nil {
		return err
	}
	_, r.PowerS = NewStoredIndicatorsPower(protocol.BWRIPowerS)
	if err := r.PowerS.Unmarshall(pdu[120:136]); err != nil {
		return err
	}
	r.Voltage = &Voltage{}
	if err := r.Voltage.Unmarshall(pdu[136:145]); err != nil {
		return err
	}
	r.PhaseShift = &PhaseShift{}
	if err := r.PhaseShift.Unmarshall(pdu[145:154]); err != nil {
		return err
	}
	r.Current = &Current{}
	if err := r.Current.Unmarshall(pdu[154:163]); err != nil {
		return err
	}
	r.PowerK = &PowerFactorWithSum{}
	if err := r.PowerK.Unmarshall(pdu[163:175]); err != nil {
		return err
	}
	r.Frequency = &Frequency{}
	if err := r.Frequency.Unmarshall(pdu[175:178]); err != nil {
		return err
	}

	return nil
}

type EnergyByTariff struct {
	Tariff1 *Energy `json:"tariff_1,omitempty"`
	Tariff2 *Energy `json:"tariff_2,omitempty"`
	Tariff3 *Energy `json:"tariff_3,omitempty"`
	Tariff4 *Energy `json:"tariff_4,omitempty"`
}

func (r *EnergyByTariff) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 64 {
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

	return nil
}
