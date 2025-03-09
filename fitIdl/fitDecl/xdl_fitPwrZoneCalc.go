package fitDecl

import strconv "strconv"

type PwrZoneCalc byte

const (
	PwrZoneCalc_Custom     PwrZoneCalc = 0
	PwrZoneCalc_PercentFtp PwrZoneCalc = 1
	PwrZoneCalc_Invalid    PwrZoneCalc = 255
)

var PwrZoneCalcNames = map[PwrZoneCalc]string{
	PwrZoneCalc_Custom:     "Custom",
	PwrZoneCalc_PercentFtp: "PercentFtp",
}

func (k PwrZoneCalc) String() string {
	if uint(k) < uint(len(PwrZoneCalcNames)) {
		if v, ok := PwrZoneCalcNames[k]; ok {
			return v
		}
	}
	return "PwrZoneCalc(" + strconv.Itoa(int(k)) + ")"
}

var PwrZoneCalcValues = map[string]PwrZoneCalc{
	"Custom":     PwrZoneCalc_Custom,
	"PercentFtp": PwrZoneCalc_PercentFtp,
}

func PwrZoneCalcFromString(s string) PwrZoneCalc {
	if v, ok := PwrZoneCalcValues[s]; ok {
		return v
	}
	return PwrZoneCalc(PwrZoneCalc_Invalid)
}
