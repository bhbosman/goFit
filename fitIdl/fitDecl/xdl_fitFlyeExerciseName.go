package fitDecl

import strconv "strconv"

type FlyeExerciseName uint16

const (
	FlyeExerciseName_CableCrossover                    FlyeExerciseName = 0
	FlyeExerciseName_DeclineDumbbellFlye               FlyeExerciseName = 1
	FlyeExerciseName_DumbbellFlye                      FlyeExerciseName = 2
	FlyeExerciseName_InclineDumbbellFlye               FlyeExerciseName = 3
	FlyeExerciseName_KettlebellFlye                    FlyeExerciseName = 4
	FlyeExerciseName_KneelingRearFlye                  FlyeExerciseName = 5
	FlyeExerciseName_SingleArmStandingCableReverseFlye FlyeExerciseName = 6
	FlyeExerciseName_SwissBallDumbbellFlye             FlyeExerciseName = 7
	FlyeExerciseName_ArmRotations                      FlyeExerciseName = 8
	FlyeExerciseName_HugATree                          FlyeExerciseName = 9
	FlyeExerciseName_Invalid                           FlyeExerciseName = 65535
)

var FlyeExerciseNameNames = map[FlyeExerciseName]string{
	FlyeExerciseName_CableCrossover:                    "CableCrossover",
	FlyeExerciseName_DeclineDumbbellFlye:               "DeclineDumbbellFlye",
	FlyeExerciseName_DumbbellFlye:                      "DumbbellFlye",
	FlyeExerciseName_InclineDumbbellFlye:               "InclineDumbbellFlye",
	FlyeExerciseName_KettlebellFlye:                    "KettlebellFlye",
	FlyeExerciseName_KneelingRearFlye:                  "KneelingRearFlye",
	FlyeExerciseName_SingleArmStandingCableReverseFlye: "SingleArmStandingCableReverseFlye",
	FlyeExerciseName_SwissBallDumbbellFlye:             "SwissBallDumbbellFlye",
	FlyeExerciseName_ArmRotations:                      "ArmRotations",
	FlyeExerciseName_HugATree:                          "HugATree",
}

func (k FlyeExerciseName) String() string {
	if uint(k) < uint(len(FlyeExerciseNameNames)) {
		if v, ok := FlyeExerciseNameNames[k]; ok {
			return v
		}
	}
	return "FlyeExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var FlyeExerciseNameValues = map[string]FlyeExerciseName{
	"CableCrossover":                    FlyeExerciseName_CableCrossover,
	"DeclineDumbbellFlye":               FlyeExerciseName_DeclineDumbbellFlye,
	"DumbbellFlye":                      FlyeExerciseName_DumbbellFlye,
	"InclineDumbbellFlye":               FlyeExerciseName_InclineDumbbellFlye,
	"KettlebellFlye":                    FlyeExerciseName_KettlebellFlye,
	"KneelingRearFlye":                  FlyeExerciseName_KneelingRearFlye,
	"SingleArmStandingCableReverseFlye": FlyeExerciseName_SingleArmStandingCableReverseFlye,
	"SwissBallDumbbellFlye":             FlyeExerciseName_SwissBallDumbbellFlye,
	"ArmRotations":                      FlyeExerciseName_ArmRotations,
	"HugATree":                          FlyeExerciseName_HugATree,
}

func FlyeExerciseNameFromString(s string) FlyeExerciseName {
	if v, ok := FlyeExerciseNameValues[s]; ok {
		return v
	}
	return FlyeExerciseName(FlyeExerciseName_Invalid)
}
