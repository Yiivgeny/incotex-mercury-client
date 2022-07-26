package protocol

type Method byte

const (
	MethodTestCommunication  Method = 0x00
	MethodOpenChannel        Method = 0x01
	MethodCloseChannel       Method = 0x02
	MethodWriteParameter     Method = 0x03 // TODO
	MethodReadArrays         Method = 0x04 // TODO
	MethodReadEnergy         Method = 0x05 // TODO
	MethodReadRaw            Method = 0x06 // TODO
	MethodWriteRaw           Method = 0x07 // TODO
	MethodReadParameter      Method = 0x08
	MethodReadEnergyReactive Method = 0x15 // TODO
	MethodReadPowerPeaks     Method = 0x17 // TODO
	MethodReadEnergyExtended Method = 0x18 // TODO
)

type AccessLevel byte

const (
	AccessLevelUser  AccessLevel = 0x01
	AccessLevelAdmin AccessLevel = 0x02
)

type Parameter byte

const (
	ParameterSerialNumberAndBuildDate  Parameter = 0x00 // Deprecated
	ParameterIndividualOptions         Parameter = 0x01
	ParameterTransformationCoefficient Parameter = 0x02 // TODO
	ParameterSoftwareVersion           Parameter = 0x03 // TODO
	ParameterInterface2Multiplier      Parameter = 0x04 // TODO
	ParameterNetworkAddress            Parameter = 0x05
	ParameterIndicationModes           Parameter = 0x06 // TODO
	ParameterDST                       Parameter = 0x07 // TODO
	ParameterLimitControlTimeout       Parameter = 0x08 // TODO
	ParameterProgrammableFlags         Parameter = 0x09 // TODO
	ParameterState                     Parameter = 0x0A // TODO
	ParameterLocation                  Parameter = 0x0B // TODO
	ParameterPowerSchedules            Parameter = 0x0C // TODO
	ParameterPowerValues               Parameter = 0x0D // TODO
	ParameterElectricIndicators        Parameter = 0x11 // Deprecated
	ParameterHardwareModification      Parameter = 0x12 // TODO
	ParameterAvgPowersMain             Parameter = 0x13 // TODO
	ParameterStoredIndicators          Parameter = 0x14
	ParameterAvgPowersAdditional       Parameter = 0x15 // TODO
	ParameterInstantIndicators         Parameter = 0x16
	ParameterTarifficatorState         Parameter = 0x17 // TODO
	ParameterLoadSwitchState           Parameter = 0x18 // TODO
	ParameterPowerLimit                Parameter = 0x19 // TODO
	ParameterEnergyLimit               Parameter = 0x1A // TODO
	ParameterIndicationsTariff         Parameter = 0x1B // TODO
	ParameterIndicationsPeriod         Parameter = 0x1C // TODO
	ParameterInterfaceMultiplier       Parameter = 0x1D // TODO
	ParameterLossesOptions             Parameter = 0x1E // TODO
	ParameterLossesPower               Parameter = 0x1F // TODO
	ParameterAllowedValues             Parameter = 0x20 // TODO
	ParameterAveragingTimes            Parameter = 0x21 // TODO
	ParameterTariffSchedule            Parameter = 0x22 // TODO
	ParameterHolidaysSchedule          Parameter = 0x23 // TODO
	ParameterLongOperationOptions      Parameter = 0x24 // TODO
	ParameterCRC                       Parameter = 0x26 // TODO
	ParameterPLCOptions                Parameter = 0x27 // TODO
	ParameterRightChannelOptions       Parameter = 0x28 // TODO
	ParameterProgrammableOptions       Parameter = 0x2B // TODO
	ParameterSettlementDay             Parameter = 0x2D // TODO
	ParameterBreakerLimiterOptions     Parameter = 0x2E // TODO
	ParameterEvents                    Parameter = 0x2F // TODO
)

type BWRIMode byte

const (
	BWRIModePower         BWRIMode = 0x00 << 4
	BWRIModeVoltage       BWRIMode = 0x01 << 4
	BWRIModeCurrent       BWRIMode = 0x02 << 4
	BWRIModeKPower        BWRIMode = 0x03 << 4
	BWRIModeFrequency     BWRIMode = 0x04 << 4
	BWRIModePhaseShift    BWRIMode = 0x05 << 4
	BWRIModeDistortion    BWRIMode = 0x06 << 4
	BWRIModeTemperature   BWRIMode = 0x07 << 4
	BWRIModeLinearVoltage BWRIMode = 0x08 << 4 // not working
	BWRIModeAccelerated   BWRIMode = 0x0A << 4
	BWRIModeDateTime      BWRIMode = 0x0E << 4
	BWRIModeEnergy        BWRIMode = 0x0F << 4
)

type BWRIPower byte

const (
	BWRIPowerP BWRIPower = 0x00 << 2
	BWRIPowerQ BWRIPower = 0x01 << 2
	BWRIPowerS BWRIPower = 0x02 << 2
)

type BWRIPhase byte

const (
	BWRIPhaseAll BWRIPhase = 0x00
	BWRIPhaseA   BWRIPhase = 0x01
	BWRIPhaseB   BWRIPhase = 0x02
	BWRIPhaseC   BWRIPhase = 0x03
)

type BWRICurrentOptions byte

const (
	BWRICurrentOptionsDiff    BWRICurrentOptions = 0x00
	BWRICurrentOptionsPhase   BWRICurrentOptions = 0x01
	BWRICurrentOptionsNeutral BWRICurrentOptions = 0x01
)

type BWRITariff byte

const (
	BWRITariffTotal BWRITariff = 0x00
	BWRITariff1     BWRITariff = 0x01
	BWRITariff2     BWRITariff = 0x02
	BWRITariff3     BWRITariff = 0x03
	BWRITariff4     BWRITariff = 0x04
)
