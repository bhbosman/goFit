package fitDecl

import strconv "strconv"

type PowerPhaseType byte

const (
	PowerPhaseType_PowerPhaseStartAngle PowerPhaseType = 0
	PowerPhaseType_PowerPhaseEndAngle   PowerPhaseType = 1
	PowerPhaseType_PowerPhaseArcLength  PowerPhaseType = 2
	PowerPhaseType_PowerPhaseCenter     PowerPhaseType = 3
	PowerPhaseType_Invalid              PowerPhaseType = 255
)

var PowerPhaseTypeNames = map[PowerPhaseType]string{
	PowerPhaseType_PowerPhaseStartAngle: "PowerPhaseStartAngle",
	PowerPhaseType_PowerPhaseEndAngle:   "PowerPhaseEndAngle",
	PowerPhaseType_PowerPhaseArcLength:  "PowerPhaseArcLength",
	PowerPhaseType_PowerPhaseCenter:     "PowerPhaseCenter",
}

func (k PowerPhaseType) String() string {
	if uint(k) < uint(len(PowerPhaseTypeNames)) {
		if v, ok := PowerPhaseTypeNames[k]; ok {
			return v
		}
	}
	return "PowerPhaseType(" + strconv.Itoa(int(k)) + ")"
}

var PowerPhaseTypeValues = map[string]PowerPhaseType{
	"PowerPhaseStartAngle": PowerPhaseType_PowerPhaseStartAngle,
	"PowerPhaseEndAngle":   PowerPhaseType_PowerPhaseEndAngle,
	"PowerPhaseArcLength":  PowerPhaseType_PowerPhaseArcLength,
	"PowerPhaseCenter":     PowerPhaseType_PowerPhaseCenter,
}

func PowerPhaseTypeFromString(s string) PowerPhaseType {
	if v, ok := PowerPhaseTypeValues[s]; ok {
		return v
	}
	return PowerPhaseType(PowerPhaseType_Invalid)
}
