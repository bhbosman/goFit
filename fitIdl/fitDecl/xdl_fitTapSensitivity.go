package fitDecl

import strconv "strconv"

type TapSensitivity byte

const (
	TapSensitivity_High    TapSensitivity = 0
	TapSensitivity_Medium  TapSensitivity = 1
	TapSensitivity_Low     TapSensitivity = 2
	TapSensitivity_Invalid TapSensitivity = 255
)

var TapSensitivityNames = map[TapSensitivity]string{
	TapSensitivity_High:   "High",
	TapSensitivity_Medium: "Medium",
	TapSensitivity_Low:    "Low",
}

func (k TapSensitivity) String() string {
	if uint(k) < uint(len(TapSensitivityNames)) {
		if v, ok := TapSensitivityNames[k]; ok {
			return v
		}
	}
	return "TapSensitivity(" + strconv.Itoa(int(k)) + ")"
}

var TapSensitivityValues = map[string]TapSensitivity{
	"High":   TapSensitivity_High,
	"Medium": TapSensitivity_Medium,
	"Low":    TapSensitivity_Low,
}

func TapSensitivityFromString(s string) TapSensitivity {
	if v, ok := TapSensitivityValues[s]; ok {
		return v
	}
	return TapSensitivity(TapSensitivity_Invalid)
}
