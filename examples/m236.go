package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Yiivgeny/incotex-mercury-client/client"
	"github.com/Yiivgeny/incotex-mercury-client/client/methods/read_energy"
	"github.com/Yiivgeny/incotex-mercury-client/client/methods/read_parameter"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

func main() {
	address := protocol.Address(0)

	cfg := client.NewConfig(9600, 2)
	cfg.Host = "192.168.90.231:8899"
	cfg.ResponseTimeout += time.Second
	auth := client.Auth{
		AccessLevel: protocol.AccessLevelUser,
		Password:    []byte{1, 1, 1, 1, 1, 1},
	}

	transport, err := client.NewTransportTCP(&cfg)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = transport.Close()
	}()

	c := client.NewClient(&cfg, transport)
	if err := c.TestCommunication(address); err != nil {
		panic(err)
	}
	fmt.Printf("Test communication OK\n")

	if err := c.OpenChannel(address, auth); err != nil {
		panic(err)
	}
	fmt.Printf("Open channel OK\n")

	defer func() {
		_ = c.CloseChannel(address)
		fmt.Printf("Close channel OK\n")
	}()

	if err := methodReadParameter(c, address); err != nil {
		panic(err)
	}
}

func methodReadParameter(c client.Client, address protocol.Address) error {

	{
		request, result := read_parameter.NewSerialNumberAndBuildDate()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: serial number and build date %s\n", str)
	}

	{
		request, result := read_parameter.NewIndividualOptions()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: individual options %s\n", str)
	}

	{
		request, result := read_parameter.NewNetworkAddress()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: network address %s\n", str)
	}

	{
		request, result := read_parameter.NewInstantIndicatorsPower(protocol.BWRIPowerP)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators power P %s\n", str)
	}

	{
		request, result := read_parameter.NewInstantIndicatorsPower(protocol.BWRIPowerQ)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators power Q %s\n", str)
	}

	{
		request, result := read_parameter.NewInstantIndicatorsPower(protocol.BWRIPowerS)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators power Sum %s\n", str)
	}

	{
		request, result := read_parameter.NewInstantIndicatorsVoltage()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators voltage %s\n", str)
	}

	{
		request, result := read_parameter.NewInstantIndicatorsCurrent()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators current %s\n", str)
	}

	{
		request, result := read_parameter.NewInstantIndicatorsPowerFactor()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators power factor %s\n", str)
	}

	{
		request, result := read_parameter.NewInstantIndicatorsFrequency()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators frequency %s\n", str)
	}

	{
		request, result := read_parameter.NewInstantIndicatorsPhaseShift()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators phase shift %s\n", str)
	}

	{
		request, result := read_parameter.NewInstantIndicatorsDistortion()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators distortion %s\n", str)
	}

	{
		request, result := read_parameter.NewInstantIndicatorsTemperature()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators temparature %s\n", str)
	}

	/* Not supported in Mercury-238
	{
		request, result := read_parameter.NewInstantIndicatorsLinearVoltage()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators linear voltage %s\n", str)
	}
	*/

	{
		request, result := read_parameter.NewInstantIndicators()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: instant indicators %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsStoreTime()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators store time %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsPower(protocol.BWRIPowerP)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators power P %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsPower(protocol.BWRIPowerQ)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators power Q %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsPower(protocol.BWRIPowerS)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators power Sum %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsVoltage()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators voltage %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsCurrent()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators current %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsPowerFactor()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators power k %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsFrequency()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators frequency %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsPhaseShift()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators phase shift %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsEnergy(protocol.BWRITariffTotal)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators energy total %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsEnergy(protocol.BWRITariff1)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators energy tariff 1 %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsEnergy(protocol.BWRITariff2)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators energy tariff 2 %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsEnergy(protocol.BWRITariff3)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators energy tariff 3 %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicatorsEnergy(protocol.BWRITariff4)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators energy tariff 4 %s\n", str)
	}

	{
		request, result := read_parameter.NewStoredIndicators()
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read parameter: stored indicators %s\n", str)
	}

	{
		request, result := read_energy.NewReadEnergy(protocol.EnergyFromReset, protocol.EnergyTariffTotal)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read energy: from reset total %s\n", str)
	}

	{
		request, result := read_energy.NewReadEnergy(protocol.EnergyFromReset, protocol.EnergyTariff1)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read energy: from reset tariff 1 %s\n", str)
	}

	{
		request, result := read_energy.NewReadEnergy(protocol.EnergyFromReset, protocol.EnergyTariff2)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read energy: from reset tariff 2 %s\n", str)
	}

	{
		request, result := read_energy.NewReadEnergy(protocol.EnergyFromReset, protocol.EnergyTariff3)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read energy: from reset tariff 3 %s\n", str)
	}

	{
		request, result := read_energy.NewReadEnergy(protocol.EnergyFromReset, protocol.EnergyTariff4)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read energy: from reset tariff 4 %s\n", str)
	}

	{
		request, result := read_energy.NewReadEnergyByPhase(protocol.EnergyTariffTotal)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read energy phases: from reset total %s\n", str)
	}

	{
		request, result := read_energy.NewReadEnergyRapid(protocol.EnergyFromReset)
		if err := c.Request(address, request, result); err != nil {
			return err
		}
		str, _ := json.Marshal(result)
		fmt.Printf("Read energy rapid: from reset total %s\n", str)
	}

	return nil
}
