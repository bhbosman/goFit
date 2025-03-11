package fitDecl

import strconv "strconv"

type GasConsumptionRateType byte

const (
	GasConsumptionRateType_PressureSac GasConsumptionRateType = 0
	GasConsumptionRateType_VolumeSac   GasConsumptionRateType = 1
	GasConsumptionRateType_Rmv         GasConsumptionRateType = 2
	GasConsumptionRateType_Invalid     GasConsumptionRateType = 255
)

var GasConsumptionRateTypeNames = map[GasConsumptionRateType]string{
	GasConsumptionRateType_PressureSac: "PressureSac",
	GasConsumptionRateType_VolumeSac:   "VolumeSac",
	GasConsumptionRateType_Rmv:         "Rmv",
}

func (k GasConsumptionRateType) String() string {
	if uint(k) < uint(len(GasConsumptionRateTypeNames)) {
		if v, ok := GasConsumptionRateTypeNames[k]; ok {
			return v
		}
	}
	return "GasConsumptionRateType(" + strconv.Itoa(int(k)) + ")"
}

var GasConsumptionRateTypeValues = map[string]GasConsumptionRateType{
	"PressureSac": GasConsumptionRateType_PressureSac,
	"VolumeSac":   GasConsumptionRateType_VolumeSac,
	"Rmv":         GasConsumptionRateType_Rmv,
}

func GasConsumptionRateTypeFromString(s string) GasConsumptionRateType {
	if v, ok := GasConsumptionRateTypeValues[s]; ok {
		return v
	}
	return GasConsumptionRateType(GasConsumptionRateType_Invalid)
}
