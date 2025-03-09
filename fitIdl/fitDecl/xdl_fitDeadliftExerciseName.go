package fitDecl

import strconv "strconv"

type DeadliftExerciseName uint16

const (
	DeadliftExerciseName_BarbellDeadlift                       DeadliftExerciseName = 0
	DeadliftExerciseName_BarbellStraightLegDeadlift            DeadliftExerciseName = 1
	DeadliftExerciseName_DumbbellDeadlift                      DeadliftExerciseName = 2
	DeadliftExerciseName_DumbbellSingleLegDeadliftToRow        DeadliftExerciseName = 3
	DeadliftExerciseName_DumbbellStraightLegDeadlift           DeadliftExerciseName = 4
	DeadliftExerciseName_KettlebellFloorToShelf                DeadliftExerciseName = 5
	DeadliftExerciseName_OneArmOneLegDeadlift                  DeadliftExerciseName = 6
	DeadliftExerciseName_RackPull                              DeadliftExerciseName = 7
	DeadliftExerciseName_RotationalDumbbellStraightLegDeadlift DeadliftExerciseName = 8
	DeadliftExerciseName_SingleArmDeadlift                     DeadliftExerciseName = 9
	DeadliftExerciseName_SingleLegBarbellDeadlift              DeadliftExerciseName = 10
	DeadliftExerciseName_SingleLegBarbellStraightLegDeadlift   DeadliftExerciseName = 11
	DeadliftExerciseName_SingleLegDeadliftWithBarbell          DeadliftExerciseName = 12
	DeadliftExerciseName_SingleLegRdlCircuit                   DeadliftExerciseName = 13
	DeadliftExerciseName_SingleLegRomanianDeadliftWithDumbbell DeadliftExerciseName = 14
	DeadliftExerciseName_SumoDeadlift                          DeadliftExerciseName = 15
	DeadliftExerciseName_SumoDeadliftHighPull                  DeadliftExerciseName = 16
	DeadliftExerciseName_TrapBarDeadlift                       DeadliftExerciseName = 17
	DeadliftExerciseName_WideGripBarbellDeadlift               DeadliftExerciseName = 18
	DeadliftExerciseName_Invalid                               DeadliftExerciseName = 65535
)

var DeadliftExerciseNameNames = map[DeadliftExerciseName]string{
	DeadliftExerciseName_BarbellDeadlift:                       "BarbellDeadlift",
	DeadliftExerciseName_BarbellStraightLegDeadlift:            "BarbellStraightLegDeadlift",
	DeadliftExerciseName_DumbbellDeadlift:                      "DumbbellDeadlift",
	DeadliftExerciseName_DumbbellSingleLegDeadliftToRow:        "DumbbellSingleLegDeadliftToRow",
	DeadliftExerciseName_DumbbellStraightLegDeadlift:           "DumbbellStraightLegDeadlift",
	DeadliftExerciseName_KettlebellFloorToShelf:                "KettlebellFloorToShelf",
	DeadliftExerciseName_OneArmOneLegDeadlift:                  "OneArmOneLegDeadlift",
	DeadliftExerciseName_RackPull:                              "RackPull",
	DeadliftExerciseName_RotationalDumbbellStraightLegDeadlift: "RotationalDumbbellStraightLegDeadlift",
	DeadliftExerciseName_SingleArmDeadlift:                     "SingleArmDeadlift",
	DeadliftExerciseName_SingleLegBarbellDeadlift:              "SingleLegBarbellDeadlift",
	DeadliftExerciseName_SingleLegBarbellStraightLegDeadlift:   "SingleLegBarbellStraightLegDeadlift",
	DeadliftExerciseName_SingleLegDeadliftWithBarbell:          "SingleLegDeadliftWithBarbell",
	DeadliftExerciseName_SingleLegRdlCircuit:                   "SingleLegRdlCircuit",
	DeadliftExerciseName_SingleLegRomanianDeadliftWithDumbbell: "SingleLegRomanianDeadliftWithDumbbell",
	DeadliftExerciseName_SumoDeadlift:                          "SumoDeadlift",
	DeadliftExerciseName_SumoDeadliftHighPull:                  "SumoDeadliftHighPull",
	DeadliftExerciseName_TrapBarDeadlift:                       "TrapBarDeadlift",
	DeadliftExerciseName_WideGripBarbellDeadlift:               "WideGripBarbellDeadlift",
}

func (k DeadliftExerciseName) String() string {
	if uint(k) < uint(len(DeadliftExerciseNameNames)) {
		if v, ok := DeadliftExerciseNameNames[k]; ok {
			return v
		}
	}
	return "DeadliftExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var DeadliftExerciseNameValues = map[string]DeadliftExerciseName{
	"BarbellDeadlift":                       DeadliftExerciseName_BarbellDeadlift,
	"BarbellStraightLegDeadlift":            DeadliftExerciseName_BarbellStraightLegDeadlift,
	"DumbbellDeadlift":                      DeadliftExerciseName_DumbbellDeadlift,
	"DumbbellSingleLegDeadliftToRow":        DeadliftExerciseName_DumbbellSingleLegDeadliftToRow,
	"DumbbellStraightLegDeadlift":           DeadliftExerciseName_DumbbellStraightLegDeadlift,
	"KettlebellFloorToShelf":                DeadliftExerciseName_KettlebellFloorToShelf,
	"OneArmOneLegDeadlift":                  DeadliftExerciseName_OneArmOneLegDeadlift,
	"RackPull":                              DeadliftExerciseName_RackPull,
	"RotationalDumbbellStraightLegDeadlift": DeadliftExerciseName_RotationalDumbbellStraightLegDeadlift,
	"SingleArmDeadlift":                     DeadliftExerciseName_SingleArmDeadlift,
	"SingleLegBarbellDeadlift":              DeadliftExerciseName_SingleLegBarbellDeadlift,
	"SingleLegBarbellStraightLegDeadlift":   DeadliftExerciseName_SingleLegBarbellStraightLegDeadlift,
	"SingleLegDeadliftWithBarbell":          DeadliftExerciseName_SingleLegDeadliftWithBarbell,
	"SingleLegRdlCircuit":                   DeadliftExerciseName_SingleLegRdlCircuit,
	"SingleLegRomanianDeadliftWithDumbbell": DeadliftExerciseName_SingleLegRomanianDeadliftWithDumbbell,
	"SumoDeadlift":                          DeadliftExerciseName_SumoDeadlift,
	"SumoDeadliftHighPull":                  DeadliftExerciseName_SumoDeadliftHighPull,
	"TrapBarDeadlift":                       DeadliftExerciseName_TrapBarDeadlift,
	"WideGripBarbellDeadlift":               DeadliftExerciseName_WideGripBarbellDeadlift,
}

func DeadliftExerciseNameFromString(s string) DeadliftExerciseName {
	if v, ok := DeadliftExerciseNameValues[s]; ok {
		return v
	}
	return DeadliftExerciseName(DeadliftExerciseName_Invalid)
}
