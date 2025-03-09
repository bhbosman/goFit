package fitDecl

import strconv "strconv"

type DisplayPower byte

const (
	DisplayPower_Watts      DisplayPower = 0
	DisplayPower_PercentFtp DisplayPower = 1
	DisplayPower_Invalid    DisplayPower = 255
)

var DisplayPowerNames = map[DisplayPower]string{
	DisplayPower_Watts:      "Watts",
	DisplayPower_PercentFtp: "PercentFtp",
}

func (k DisplayPower) String() string {
	if uint(k) < uint(len(DisplayPowerNames)) {
		if v, ok := DisplayPowerNames[k]; ok {
			return v
		}
	}
	return "DisplayPower(" + strconv.Itoa(int(k)) + ")"
}

var DisplayPowerValues = map[string]DisplayPower{
	"Watts":      DisplayPower_Watts,
	"PercentFtp": DisplayPower_PercentFtp,
}

func DisplayPowerFromString(s string) DisplayPower {
	if v, ok := DisplayPowerValues[s]; ok {
		return v
	}
	return DisplayPower(DisplayPower_Invalid)
}
