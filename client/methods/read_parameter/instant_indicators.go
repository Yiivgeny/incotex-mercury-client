package read_parameter

import (
	"errors"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func NewInstantIndicators() (protocol.PDU, *InstantIndicators) {
	target := &InstantIndicators{}
	return protocol.PDU{
		byte(protocol.MethodReadParameter),
		byte(protocol.ParameterInstantIndicators),
		byte(protocol.BWRIModeAccelerated),
	}, target
}

type InstantIndicators struct {
	PowerP        *PowerInstant       `json:"power_p,omitempty"`
	PowerQ        *PowerInstant       `json:"power_q,omitempty"`
	PowerS        *PowerInstant       `json:"power_s,omitempty"`
	Voltage       *Voltage            `json:"voltage,omitempty"`
	PhaseShift    *PhaseShift         `json:"phase_shift,omitempty"`
	Current       *Current            `json:"current,omitempty"`
	PowerFactor   *PowerFactorWithSum `json:"power_factor,omitempty"`
	Frequency     *Frequency          `json:"frequency,omitempty"`
	Distortion    *Distortion         `json:"distortion,omitempty"`
	Temperature   *Temperature        `json:"temperature,omitempty"`
	LinearVoltage *LinearVoltage      `json:"linear_voltage,omitempty"`
}

func (r *InstantIndicators) Unmarshall(pdu protocol.PDU) error {
	if len(pdu) != 86 && len(pdu) != 95 {
		return errors.New("pdu length mismatch")
	}

	r.PowerP = &PowerInstant{}
	if err := r.PowerP.Unmarshall(pdu[0:12]); err != nil {
		return err
	}
	r.PowerQ = &PowerInstant{}
	if err := r.PowerQ.Unmarshall(pdu[12:24]); err != nil {
		return err
	}
	r.PowerS = &PowerInstant{}
	if err := r.PowerS.Unmarshall(pdu[24:36]); err != nil {
		return err
	}
	r.Voltage = &Voltage{}
	if err := r.Voltage.Unmarshall(pdu[36:45]); err != nil {
		return err
	}
	r.PhaseShift = &PhaseShift{}
	if err := r.PhaseShift.Unmarshall(pdu[45:54]); err != nil {
		return err
	}
	r.Current = &Current{}
	if err := r.Current.Unmarshall(pdu[54:63]); err != nil {
		return err
	}
	r.PowerFactor = &PowerFactorWithSum{}
	if err := r.PowerFactor.Unmarshall(pdu[63:75]); err != nil {
		return err
	}
	r.Frequency = &Frequency{}
	if err := r.Frequency.Unmarshall(pdu[75:78]); err != nil {
		return err
	}
	r.Distortion = &Distortion{}
	if err := r.Distortion.Unmarshall(pdu[78:84]); err != nil {
		return err
	}
	r.Temperature = &Temperature{}
	if err := r.Temperature.Unmarshall(pdu[84:86]); err != nil {
		return err
	}
	if len(pdu) >= 95 {
		r.LinearVoltage = &LinearVoltage{}
		if err := r.LinearVoltage.Unmarshall(pdu[86:95]); err != nil {
			return err
		}
	}

	return nil
}
