package fitDecl

import strconv "strconv"

type ActivityType byte

const (
	ActivityType_Generic          ActivityType = 0
	ActivityType_Running          ActivityType = 1
	ActivityType_Cycling          ActivityType = 2
	ActivityType_Transition       ActivityType = 3
	ActivityType_FitnessEquipment ActivityType = 4
	ActivityType_Swimming         ActivityType = 5
	ActivityType_Walking          ActivityType = 6
	ActivityType_Sedentary        ActivityType = 8
	ActivityType_All              ActivityType = 254
	ActivityType_Invalid          ActivityType = 255
)

var ActivityTypeNames = map[ActivityType]string{
	ActivityType_Generic:          "Generic",
	ActivityType_Running:          "Running",
	ActivityType_Cycling:          "Cycling",
	ActivityType_Transition:       "Transition",
	ActivityType_FitnessEquipment: "FitnessEquipment",
	ActivityType_Swimming:         "Swimming",
	ActivityType_Walking:          "Walking",
	ActivityType_Sedentary:        "Sedentary",
	ActivityType_All:              "All",
}

func (k ActivityType) String() string {
	if uint(k) < uint(len(ActivityTypeNames)) {
		if v, ok := ActivityTypeNames[k]; ok {
			return v
		}
	}
	return "ActivityType(" + strconv.Itoa(int(k)) + ")"
}

var ActivityTypeValues = map[string]ActivityType{
	"Generic":          ActivityType_Generic,
	"Running":          ActivityType_Running,
	"Cycling":          ActivityType_Cycling,
	"Transition":       ActivityType_Transition,
	"FitnessEquipment": ActivityType_FitnessEquipment,
	"Swimming":         ActivityType_Swimming,
	"Walking":          ActivityType_Walking,
	"Sedentary":        ActivityType_Sedentary,
	"All":              ActivityType_All,
}

func ActivityTypeFromString(s string) ActivityType {
	if v, ok := ActivityTypeValues[s]; ok {
		return v
	}
	return ActivityType(ActivityType_Invalid)
}
