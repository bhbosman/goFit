package fitDecl

import strconv "strconv"

type HrZoneCalc byte

const (
	HrZoneCalc_Custom       HrZoneCalc = 0
	HrZoneCalc_PercentMaxHr HrZoneCalc = 1
	HrZoneCalc_PercentHrr   HrZoneCalc = 2
	HrZoneCalc_PercentLthr  HrZoneCalc = 3
	HrZoneCalc_Invalid      HrZoneCalc = 255
)

var HrZoneCalcNames = map[HrZoneCalc]string{
	HrZoneCalc_Custom:       "Custom",
	HrZoneCalc_PercentMaxHr: "PercentMaxHr",
	HrZoneCalc_PercentHrr:   "PercentHrr",
	HrZoneCalc_PercentLthr:  "PercentLthr",
}

func (k HrZoneCalc) String() string {
	if uint(k) < uint(len(HrZoneCalcNames)) {
		if v, ok := HrZoneCalcNames[k]; ok {
			return v
		}
	}
	return "HrZoneCalc(" + strconv.Itoa(int(k)) + ")"
}

var HrZoneCalcValues = map[string]HrZoneCalc{
	"Custom":       HrZoneCalc_Custom,
	"PercentMaxHr": HrZoneCalc_PercentMaxHr,
	"PercentHrr":   HrZoneCalc_PercentHrr,
	"PercentLthr":  HrZoneCalc_PercentLthr,
}

func HrZoneCalcFromString(s string) HrZoneCalc {
	if v, ok := HrZoneCalcValues[s]; ok {
		return v
	}
	return HrZoneCalc(HrZoneCalc_Invalid)
}
