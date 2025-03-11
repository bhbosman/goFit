package fitDecl

import strconv "strconv"

type TotalBodyExerciseName uint16

const (
	TotalBodyExerciseName_Burpee                           TotalBodyExerciseName = 0
	TotalBodyExerciseName_WeightedBurpee                   TotalBodyExerciseName = 1
	TotalBodyExerciseName_BurpeeBoxJump                    TotalBodyExerciseName = 2
	TotalBodyExerciseName_WeightedBurpeeBoxJump            TotalBodyExerciseName = 3
	TotalBodyExerciseName_HighPullBurpee                   TotalBodyExerciseName = 4
	TotalBodyExerciseName_ManMakers                        TotalBodyExerciseName = 5
	TotalBodyExerciseName_OneArmBurpee                     TotalBodyExerciseName = 6
	TotalBodyExerciseName_SquatThrusts                     TotalBodyExerciseName = 7
	TotalBodyExerciseName_WeightedSquatThrusts             TotalBodyExerciseName = 8
	TotalBodyExerciseName_SquatPlankPushUp                 TotalBodyExerciseName = 9
	TotalBodyExerciseName_WeightedSquatPlankPushUp         TotalBodyExerciseName = 10
	TotalBodyExerciseName_StandingTRotationBalance         TotalBodyExerciseName = 11
	TotalBodyExerciseName_WeightedStandingTRotationBalance TotalBodyExerciseName = 12
	TotalBodyExerciseName_Invalid                          TotalBodyExerciseName = 65535
)

var TotalBodyExerciseNameNames = map[TotalBodyExerciseName]string{
	TotalBodyExerciseName_Burpee:                           "Burpee",
	TotalBodyExerciseName_WeightedBurpee:                   "WeightedBurpee",
	TotalBodyExerciseName_BurpeeBoxJump:                    "BurpeeBoxJump",
	TotalBodyExerciseName_WeightedBurpeeBoxJump:            "WeightedBurpeeBoxJump",
	TotalBodyExerciseName_HighPullBurpee:                   "HighPullBurpee",
	TotalBodyExerciseName_ManMakers:                        "ManMakers",
	TotalBodyExerciseName_OneArmBurpee:                     "OneArmBurpee",
	TotalBodyExerciseName_SquatThrusts:                     "SquatThrusts",
	TotalBodyExerciseName_WeightedSquatThrusts:             "WeightedSquatThrusts",
	TotalBodyExerciseName_SquatPlankPushUp:                 "SquatPlankPushUp",
	TotalBodyExerciseName_WeightedSquatPlankPushUp:         "WeightedSquatPlankPushUp",
	TotalBodyExerciseName_StandingTRotationBalance:         "StandingTRotationBalance",
	TotalBodyExerciseName_WeightedStandingTRotationBalance: "WeightedStandingTRotationBalance",
}

func (k TotalBodyExerciseName) String() string {
	if uint(k) < uint(len(TotalBodyExerciseNameNames)) {
		if v, ok := TotalBodyExerciseNameNames[k]; ok {
			return v
		}
	}
	return "TotalBodyExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var TotalBodyExerciseNameValues = map[string]TotalBodyExerciseName{
	"Burpee":                           TotalBodyExerciseName_Burpee,
	"WeightedBurpee":                   TotalBodyExerciseName_WeightedBurpee,
	"BurpeeBoxJump":                    TotalBodyExerciseName_BurpeeBoxJump,
	"WeightedBurpeeBoxJump":            TotalBodyExerciseName_WeightedBurpeeBoxJump,
	"HighPullBurpee":                   TotalBodyExerciseName_HighPullBurpee,
	"ManMakers":                        TotalBodyExerciseName_ManMakers,
	"OneArmBurpee":                     TotalBodyExerciseName_OneArmBurpee,
	"SquatThrusts":                     TotalBodyExerciseName_SquatThrusts,
	"WeightedSquatThrusts":             TotalBodyExerciseName_WeightedSquatThrusts,
	"SquatPlankPushUp":                 TotalBodyExerciseName_SquatPlankPushUp,
	"WeightedSquatPlankPushUp":         TotalBodyExerciseName_WeightedSquatPlankPushUp,
	"StandingTRotationBalance":         TotalBodyExerciseName_StandingTRotationBalance,
	"WeightedStandingTRotationBalance": TotalBodyExerciseName_WeightedStandingTRotationBalance,
}

func TotalBodyExerciseNameFromString(s string) TotalBodyExerciseName {
	if v, ok := TotalBodyExerciseNameValues[s]; ok {
		return v
	}
	return TotalBodyExerciseName(TotalBodyExerciseName_Invalid)
}
