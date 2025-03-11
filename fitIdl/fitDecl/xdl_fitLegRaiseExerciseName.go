package fitDecl

import strconv "strconv"

type LegRaiseExerciseName uint16

const (
	LegRaiseExerciseName_HangingKneeRaise                   LegRaiseExerciseName = 0
	LegRaiseExerciseName_HangingLegRaise                    LegRaiseExerciseName = 1
	LegRaiseExerciseName_WeightedHangingLegRaise            LegRaiseExerciseName = 2
	LegRaiseExerciseName_HangingSingleLegRaise              LegRaiseExerciseName = 3
	LegRaiseExerciseName_WeightedHangingSingleLegRaise      LegRaiseExerciseName = 4
	LegRaiseExerciseName_KettlebellLegRaises                LegRaiseExerciseName = 5
	LegRaiseExerciseName_LegLoweringDrill                   LegRaiseExerciseName = 6
	LegRaiseExerciseName_WeightedLegLoweringDrill           LegRaiseExerciseName = 7
	LegRaiseExerciseName_LyingStraightLegRaise              LegRaiseExerciseName = 8
	LegRaiseExerciseName_WeightedLyingStraightLegRaise      LegRaiseExerciseName = 9
	LegRaiseExerciseName_MedicineBallLegDrops               LegRaiseExerciseName = 10
	LegRaiseExerciseName_QuadrupedLegRaise                  LegRaiseExerciseName = 11
	LegRaiseExerciseName_WeightedQuadrupedLegRaise          LegRaiseExerciseName = 12
	LegRaiseExerciseName_ReverseLegRaise                    LegRaiseExerciseName = 13
	LegRaiseExerciseName_WeightedReverseLegRaise            LegRaiseExerciseName = 14
	LegRaiseExerciseName_ReverseLegRaiseOnSwissBall         LegRaiseExerciseName = 15
	LegRaiseExerciseName_WeightedReverseLegRaiseOnSwissBall LegRaiseExerciseName = 16
	LegRaiseExerciseName_SingleLegLoweringDrill             LegRaiseExerciseName = 17
	LegRaiseExerciseName_WeightedSingleLegLoweringDrill     LegRaiseExerciseName = 18
	LegRaiseExerciseName_WeightedHangingKneeRaise           LegRaiseExerciseName = 19
	LegRaiseExerciseName_LateralStepover                    LegRaiseExerciseName = 20
	LegRaiseExerciseName_WeightedLateralStepover            LegRaiseExerciseName = 21
	LegRaiseExerciseName_Invalid                            LegRaiseExerciseName = 65535
)

var LegRaiseExerciseNameNames = map[LegRaiseExerciseName]string{
	LegRaiseExerciseName_HangingKneeRaise:                   "HangingKneeRaise",
	LegRaiseExerciseName_HangingLegRaise:                    "HangingLegRaise",
	LegRaiseExerciseName_WeightedHangingLegRaise:            "WeightedHangingLegRaise",
	LegRaiseExerciseName_HangingSingleLegRaise:              "HangingSingleLegRaise",
	LegRaiseExerciseName_WeightedHangingSingleLegRaise:      "WeightedHangingSingleLegRaise",
	LegRaiseExerciseName_KettlebellLegRaises:                "KettlebellLegRaises",
	LegRaiseExerciseName_LegLoweringDrill:                   "LegLoweringDrill",
	LegRaiseExerciseName_WeightedLegLoweringDrill:           "WeightedLegLoweringDrill",
	LegRaiseExerciseName_LyingStraightLegRaise:              "LyingStraightLegRaise",
	LegRaiseExerciseName_WeightedLyingStraightLegRaise:      "WeightedLyingStraightLegRaise",
	LegRaiseExerciseName_MedicineBallLegDrops:               "MedicineBallLegDrops",
	LegRaiseExerciseName_QuadrupedLegRaise:                  "QuadrupedLegRaise",
	LegRaiseExerciseName_WeightedQuadrupedLegRaise:          "WeightedQuadrupedLegRaise",
	LegRaiseExerciseName_ReverseLegRaise:                    "ReverseLegRaise",
	LegRaiseExerciseName_WeightedReverseLegRaise:            "WeightedReverseLegRaise",
	LegRaiseExerciseName_ReverseLegRaiseOnSwissBall:         "ReverseLegRaiseOnSwissBall",
	LegRaiseExerciseName_WeightedReverseLegRaiseOnSwissBall: "WeightedReverseLegRaiseOnSwissBall",
	LegRaiseExerciseName_SingleLegLoweringDrill:             "SingleLegLoweringDrill",
	LegRaiseExerciseName_WeightedSingleLegLoweringDrill:     "WeightedSingleLegLoweringDrill",
	LegRaiseExerciseName_WeightedHangingKneeRaise:           "WeightedHangingKneeRaise",
	LegRaiseExerciseName_LateralStepover:                    "LateralStepover",
	LegRaiseExerciseName_WeightedLateralStepover:            "WeightedLateralStepover",
}

func (k LegRaiseExerciseName) String() string {
	if uint(k) < uint(len(LegRaiseExerciseNameNames)) {
		if v, ok := LegRaiseExerciseNameNames[k]; ok {
			return v
		}
	}
	return "LegRaiseExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var LegRaiseExerciseNameValues = map[string]LegRaiseExerciseName{
	"HangingKneeRaise":                   LegRaiseExerciseName_HangingKneeRaise,
	"HangingLegRaise":                    LegRaiseExerciseName_HangingLegRaise,
	"WeightedHangingLegRaise":            LegRaiseExerciseName_WeightedHangingLegRaise,
	"HangingSingleLegRaise":              LegRaiseExerciseName_HangingSingleLegRaise,
	"WeightedHangingSingleLegRaise":      LegRaiseExerciseName_WeightedHangingSingleLegRaise,
	"KettlebellLegRaises":                LegRaiseExerciseName_KettlebellLegRaises,
	"LegLoweringDrill":                   LegRaiseExerciseName_LegLoweringDrill,
	"WeightedLegLoweringDrill":           LegRaiseExerciseName_WeightedLegLoweringDrill,
	"LyingStraightLegRaise":              LegRaiseExerciseName_LyingStraightLegRaise,
	"WeightedLyingStraightLegRaise":      LegRaiseExerciseName_WeightedLyingStraightLegRaise,
	"MedicineBallLegDrops":               LegRaiseExerciseName_MedicineBallLegDrops,
	"QuadrupedLegRaise":                  LegRaiseExerciseName_QuadrupedLegRaise,
	"WeightedQuadrupedLegRaise":          LegRaiseExerciseName_WeightedQuadrupedLegRaise,
	"ReverseLegRaise":                    LegRaiseExerciseName_ReverseLegRaise,
	"WeightedReverseLegRaise":            LegRaiseExerciseName_WeightedReverseLegRaise,
	"ReverseLegRaiseOnSwissBall":         LegRaiseExerciseName_ReverseLegRaiseOnSwissBall,
	"WeightedReverseLegRaiseOnSwissBall": LegRaiseExerciseName_WeightedReverseLegRaiseOnSwissBall,
	"SingleLegLoweringDrill":             LegRaiseExerciseName_SingleLegLoweringDrill,
	"WeightedSingleLegLoweringDrill":     LegRaiseExerciseName_WeightedSingleLegLoweringDrill,
	"WeightedHangingKneeRaise":           LegRaiseExerciseName_WeightedHangingKneeRaise,
	"LateralStepover":                    LegRaiseExerciseName_LateralStepover,
	"WeightedLateralStepover":            LegRaiseExerciseName_WeightedLateralStepover,
}

func LegRaiseExerciseNameFromString(s string) LegRaiseExerciseName {
	if v, ok := LegRaiseExerciseNameValues[s]; ok {
		return v
	}
	return LegRaiseExerciseName(LegRaiseExerciseName_Invalid)
}
