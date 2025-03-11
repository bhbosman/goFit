package fitDecl

import strconv "strconv"

type ActivityClass byte

const (
	ActivityClass_Level    ActivityClass = 127
	ActivityClass_LevelMax ActivityClass = 100
	ActivityClass_Athlete  ActivityClass = 128
	ActivityClass_Invalid  ActivityClass = 255
)

var ActivityClassNames = map[ActivityClass]string{
	ActivityClass_Level:    "Level",
	ActivityClass_LevelMax: "LevelMax",
	ActivityClass_Athlete:  "Athlete",
}

func (k ActivityClass) String() string {
	if uint(k) < uint(len(ActivityClassNames)) {
		if v, ok := ActivityClassNames[k]; ok {
			return v
		}
	}
	return "ActivityClass(" + strconv.Itoa(int(k)) + ")"
}

var ActivityClassValues = map[string]ActivityClass{
	"Level":    ActivityClass_Level,
	"LevelMax": ActivityClass_LevelMax,
	"Athlete":  ActivityClass_Athlete,
}

func ActivityClassFromString(s string) ActivityClass {
	if v, ok := ActivityClassValues[s]; ok {
		return v
	}
	return ActivityClass(ActivityClass_Invalid)
}
