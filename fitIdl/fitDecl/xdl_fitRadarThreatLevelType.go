package fitDecl

import strconv "strconv"

type RadarThreatLevelType byte

const (
	RadarThreatLevelType_ThreatUnknown         RadarThreatLevelType = 0
	RadarThreatLevelType_ThreatNone            RadarThreatLevelType = 1
	RadarThreatLevelType_ThreatApproaching     RadarThreatLevelType = 2
	RadarThreatLevelType_ThreatApproachingFast RadarThreatLevelType = 3
	RadarThreatLevelType_Invalid               RadarThreatLevelType = 255
)

var RadarThreatLevelTypeNames = map[RadarThreatLevelType]string{
	RadarThreatLevelType_ThreatUnknown:         "ThreatUnknown",
	RadarThreatLevelType_ThreatNone:            "ThreatNone",
	RadarThreatLevelType_ThreatApproaching:     "ThreatApproaching",
	RadarThreatLevelType_ThreatApproachingFast: "ThreatApproachingFast",
}

func (k RadarThreatLevelType) String() string {
	if uint(k) < uint(len(RadarThreatLevelTypeNames)) {
		if v, ok := RadarThreatLevelTypeNames[k]; ok {
			return v
		}
	}
	return "RadarThreatLevelType(" + strconv.Itoa(int(k)) + ")"
}

var RadarThreatLevelTypeValues = map[string]RadarThreatLevelType{
	"ThreatUnknown":         RadarThreatLevelType_ThreatUnknown,
	"ThreatNone":            RadarThreatLevelType_ThreatNone,
	"ThreatApproaching":     RadarThreatLevelType_ThreatApproaching,
	"ThreatApproachingFast": RadarThreatLevelType_ThreatApproachingFast,
}

func RadarThreatLevelTypeFromString(s string) RadarThreatLevelType {
	if v, ok := RadarThreatLevelTypeValues[s]; ok {
		return v
	}
	return RadarThreatLevelType(RadarThreatLevelType_Invalid)
}
