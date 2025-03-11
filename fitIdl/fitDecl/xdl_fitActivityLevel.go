package fitDecl

import strconv "strconv"

type ActivityLevel byte

const (
	ActivityLevel_Low     ActivityLevel = 0
	ActivityLevel_Medium  ActivityLevel = 1
	ActivityLevel_High    ActivityLevel = 2
	ActivityLevel_Invalid ActivityLevel = 255
)

var ActivityLevelNames = map[ActivityLevel]string{
	ActivityLevel_Low:    "Low",
	ActivityLevel_Medium: "Medium",
	ActivityLevel_High:   "High",
}

func (k ActivityLevel) String() string {
	if uint(k) < uint(len(ActivityLevelNames)) {
		if v, ok := ActivityLevelNames[k]; ok {
			return v
		}
	}
	return "ActivityLevel(" + strconv.Itoa(int(k)) + ")"
}

var ActivityLevelValues = map[string]ActivityLevel{
	"Low":    ActivityLevel_Low,
	"Medium": ActivityLevel_Medium,
	"High":   ActivityLevel_High,
}

func ActivityLevelFromString(s string) ActivityLevel {
	if v, ok := ActivityLevelValues[s]; ok {
		return v
	}
	return ActivityLevel(ActivityLevel_Invalid)
}
