package fitDecl

import strconv "strconv"

type ShrugExerciseName uint16

const (
	ShrugExerciseName_BarbellJumpShrug               ShrugExerciseName = 0
	ShrugExerciseName_BarbellShrug                   ShrugExerciseName = 1
	ShrugExerciseName_BarbellUprightRow              ShrugExerciseName = 2
	ShrugExerciseName_BehindTheBackSmithMachineShrug ShrugExerciseName = 3
	ShrugExerciseName_DumbbellJumpShrug              ShrugExerciseName = 4
	ShrugExerciseName_DumbbellShrug                  ShrugExerciseName = 5
	ShrugExerciseName_DumbbellUprightRow             ShrugExerciseName = 6
	ShrugExerciseName_InclineDumbbellShrug           ShrugExerciseName = 7
	ShrugExerciseName_OverheadBarbellShrug           ShrugExerciseName = 8
	ShrugExerciseName_OverheadDumbbellShrug          ShrugExerciseName = 9
	ShrugExerciseName_ScaptionAndShrug               ShrugExerciseName = 10
	ShrugExerciseName_ScapularRetraction             ShrugExerciseName = 11
	ShrugExerciseName_SerratusChairShrug             ShrugExerciseName = 12
	ShrugExerciseName_WeightedSerratusChairShrug     ShrugExerciseName = 13
	ShrugExerciseName_SerratusShrug                  ShrugExerciseName = 14
	ShrugExerciseName_WeightedSerratusShrug          ShrugExerciseName = 15
	ShrugExerciseName_WideGripJumpShrug              ShrugExerciseName = 16
	ShrugExerciseName_Invalid                        ShrugExerciseName = 65535
)

var ShrugExerciseNameNames = map[ShrugExerciseName]string{
	ShrugExerciseName_BarbellJumpShrug:               "BarbellJumpShrug",
	ShrugExerciseName_BarbellShrug:                   "BarbellShrug",
	ShrugExerciseName_BarbellUprightRow:              "BarbellUprightRow",
	ShrugExerciseName_BehindTheBackSmithMachineShrug: "BehindTheBackSmithMachineShrug",
	ShrugExerciseName_DumbbellJumpShrug:              "DumbbellJumpShrug",
	ShrugExerciseName_DumbbellShrug:                  "DumbbellShrug",
	ShrugExerciseName_DumbbellUprightRow:             "DumbbellUprightRow",
	ShrugExerciseName_InclineDumbbellShrug:           "InclineDumbbellShrug",
	ShrugExerciseName_OverheadBarbellShrug:           "OverheadBarbellShrug",
	ShrugExerciseName_OverheadDumbbellShrug:          "OverheadDumbbellShrug",
	ShrugExerciseName_ScaptionAndShrug:               "ScaptionAndShrug",
	ShrugExerciseName_ScapularRetraction:             "ScapularRetraction",
	ShrugExerciseName_SerratusChairShrug:             "SerratusChairShrug",
	ShrugExerciseName_WeightedSerratusChairShrug:     "WeightedSerratusChairShrug",
	ShrugExerciseName_SerratusShrug:                  "SerratusShrug",
	ShrugExerciseName_WeightedSerratusShrug:          "WeightedSerratusShrug",
	ShrugExerciseName_WideGripJumpShrug:              "WideGripJumpShrug",
}

func (k ShrugExerciseName) String() string {
	if uint(k) < uint(len(ShrugExerciseNameNames)) {
		if v, ok := ShrugExerciseNameNames[k]; ok {
			return v
		}
	}
	return "ShrugExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var ShrugExerciseNameValues = map[string]ShrugExerciseName{
	"BarbellJumpShrug":               ShrugExerciseName_BarbellJumpShrug,
	"BarbellShrug":                   ShrugExerciseName_BarbellShrug,
	"BarbellUprightRow":              ShrugExerciseName_BarbellUprightRow,
	"BehindTheBackSmithMachineShrug": ShrugExerciseName_BehindTheBackSmithMachineShrug,
	"DumbbellJumpShrug":              ShrugExerciseName_DumbbellJumpShrug,
	"DumbbellShrug":                  ShrugExerciseName_DumbbellShrug,
	"DumbbellUprightRow":             ShrugExerciseName_DumbbellUprightRow,
	"InclineDumbbellShrug":           ShrugExerciseName_InclineDumbbellShrug,
	"OverheadBarbellShrug":           ShrugExerciseName_OverheadBarbellShrug,
	"OverheadDumbbellShrug":          ShrugExerciseName_OverheadDumbbellShrug,
	"ScaptionAndShrug":               ShrugExerciseName_ScaptionAndShrug,
	"ScapularRetraction":             ShrugExerciseName_ScapularRetraction,
	"SerratusChairShrug":             ShrugExerciseName_SerratusChairShrug,
	"WeightedSerratusChairShrug":     ShrugExerciseName_WeightedSerratusChairShrug,
	"SerratusShrug":                  ShrugExerciseName_SerratusShrug,
	"WeightedSerratusShrug":          ShrugExerciseName_WeightedSerratusShrug,
	"WideGripJumpShrug":              ShrugExerciseName_WideGripJumpShrug,
}

func ShrugExerciseNameFromString(s string) ShrugExerciseName {
	if v, ok := ShrugExerciseNameValues[s]; ok {
		return v
	}
	return ShrugExerciseName(ShrugExerciseName_Invalid)
}
