package fitDecl

import strconv "strconv"

type CarryExerciseName uint16

const (
	CarryExerciseName_BarHolds          CarryExerciseName = 0
	CarryExerciseName_FarmersWalk       CarryExerciseName = 1
	CarryExerciseName_FarmersWalkOnToes CarryExerciseName = 2
	CarryExerciseName_HexDumbbellHold   CarryExerciseName = 3
	CarryExerciseName_OverheadCarry     CarryExerciseName = 4
	CarryExerciseName_Invalid           CarryExerciseName = 65535
)

var CarryExerciseNameNames = map[CarryExerciseName]string{
	CarryExerciseName_BarHolds:          "BarHolds",
	CarryExerciseName_FarmersWalk:       "FarmersWalk",
	CarryExerciseName_FarmersWalkOnToes: "FarmersWalkOnToes",
	CarryExerciseName_HexDumbbellHold:   "HexDumbbellHold",
	CarryExerciseName_OverheadCarry:     "OverheadCarry",
}

func (k CarryExerciseName) String() string {
	if uint(k) < uint(len(CarryExerciseNameNames)) {
		if v, ok := CarryExerciseNameNames[k]; ok {
			return v
		}
	}
	return "CarryExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var CarryExerciseNameValues = map[string]CarryExerciseName{
	"BarHolds":          CarryExerciseName_BarHolds,
	"FarmersWalk":       CarryExerciseName_FarmersWalk,
	"FarmersWalkOnToes": CarryExerciseName_FarmersWalkOnToes,
	"HexDumbbellHold":   CarryExerciseName_HexDumbbellHold,
	"OverheadCarry":     CarryExerciseName_OverheadCarry,
}

func CarryExerciseNameFromString(s string) CarryExerciseName {
	if v, ok := CarryExerciseNameValues[s]; ok {
		return v
	}
	return CarryExerciseName(CarryExerciseName_Invalid)
}
