package fitDecl

import strconv "strconv"

type HipSwingExerciseName uint16

const (
	HipSwingExerciseName_SingleArmKettlebellSwing HipSwingExerciseName = 0
	HipSwingExerciseName_SingleArmDumbbellSwing   HipSwingExerciseName = 1
	HipSwingExerciseName_StepOutSwing             HipSwingExerciseName = 2
	HipSwingExerciseName_Invalid                  HipSwingExerciseName = 65535
)

var HipSwingExerciseNameNames = map[HipSwingExerciseName]string{
	HipSwingExerciseName_SingleArmKettlebellSwing: "SingleArmKettlebellSwing",
	HipSwingExerciseName_SingleArmDumbbellSwing:   "SingleArmDumbbellSwing",
	HipSwingExerciseName_StepOutSwing:             "StepOutSwing",
}

func (k HipSwingExerciseName) String() string {
	if uint(k) < uint(len(HipSwingExerciseNameNames)) {
		if v, ok := HipSwingExerciseNameNames[k]; ok {
			return v
		}
	}
	return "HipSwingExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var HipSwingExerciseNameValues = map[string]HipSwingExerciseName{
	"SingleArmKettlebellSwing": HipSwingExerciseName_SingleArmKettlebellSwing,
	"SingleArmDumbbellSwing":   HipSwingExerciseName_SingleArmDumbbellSwing,
	"StepOutSwing":             HipSwingExerciseName_StepOutSwing,
}

func HipSwingExerciseNameFromString(s string) HipSwingExerciseName {
	if v, ok := HipSwingExerciseNameValues[s]; ok {
		return v
	}
	return HipSwingExerciseName(HipSwingExerciseName_Invalid)
}
