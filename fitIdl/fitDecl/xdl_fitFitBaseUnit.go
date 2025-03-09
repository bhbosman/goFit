package fitDecl

import strconv "strconv"

type FitBaseUnit uint16

const (
	FitBaseUnit_Other    FitBaseUnit = 0
	FitBaseUnit_Kilogram FitBaseUnit = 1
	FitBaseUnit_Pound    FitBaseUnit = 2
	FitBaseUnit_Invalid  FitBaseUnit = 65535
)

var FitBaseUnitNames = map[FitBaseUnit]string{
	FitBaseUnit_Other:    "Other",
	FitBaseUnit_Kilogram: "Kilogram",
	FitBaseUnit_Pound:    "Pound",
}

func (k FitBaseUnit) String() string {
	if uint(k) < uint(len(FitBaseUnitNames)) {
		if v, ok := FitBaseUnitNames[k]; ok {
			return v
		}
	}
	return "FitBaseUnit(" + strconv.Itoa(int(k)) + ")"
}

var FitBaseUnitValues = map[string]FitBaseUnit{
	"Other":    FitBaseUnit_Other,
	"Kilogram": FitBaseUnit_Kilogram,
	"Pound":    FitBaseUnit_Pound,
}

func FitBaseUnitFromString(s string) FitBaseUnit {
	if v, ok := FitBaseUnitValues[s]; ok {
		return v
	}
	return FitBaseUnit(FitBaseUnit_Invalid)
}
