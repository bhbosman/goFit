package fitDecl

import strconv "strconv"

type ShoulderPressExerciseName uint16

const (
	ShoulderPressExerciseName_AlternatingDumbbellShoulderPress         ShoulderPressExerciseName = 0
	ShoulderPressExerciseName_ArnoldPress                              ShoulderPressExerciseName = 1
	ShoulderPressExerciseName_BarbellFrontSquatToPushPress             ShoulderPressExerciseName = 2
	ShoulderPressExerciseName_BarbellPushPress                         ShoulderPressExerciseName = 3
	ShoulderPressExerciseName_BarbellShoulderPress                     ShoulderPressExerciseName = 4
	ShoulderPressExerciseName_DeadCurlPress                            ShoulderPressExerciseName = 5
	ShoulderPressExerciseName_DumbbellAlternatingShoulderPressAndTwist ShoulderPressExerciseName = 6
	ShoulderPressExerciseName_DumbbellHammerCurlToLungeToPress         ShoulderPressExerciseName = 7
	ShoulderPressExerciseName_DumbbellPushPress                        ShoulderPressExerciseName = 8
	ShoulderPressExerciseName_FloorInvertedShoulderPress               ShoulderPressExerciseName = 9
	ShoulderPressExerciseName_WeightedFloorInvertedShoulderPress       ShoulderPressExerciseName = 10
	ShoulderPressExerciseName_InvertedShoulderPress                    ShoulderPressExerciseName = 11
	ShoulderPressExerciseName_WeightedInvertedShoulderPress            ShoulderPressExerciseName = 12
	ShoulderPressExerciseName_OneArmPushPress                          ShoulderPressExerciseName = 13
	ShoulderPressExerciseName_OverheadBarbellPress                     ShoulderPressExerciseName = 14
	ShoulderPressExerciseName_OverheadDumbbellPress                    ShoulderPressExerciseName = 15
	ShoulderPressExerciseName_SeatedBarbellShoulderPress               ShoulderPressExerciseName = 16
	ShoulderPressExerciseName_SeatedDumbbellShoulderPress              ShoulderPressExerciseName = 17
	ShoulderPressExerciseName_SingleArmDumbbellShoulderPress           ShoulderPressExerciseName = 18
	ShoulderPressExerciseName_SingleArmStepUpAndPress                  ShoulderPressExerciseName = 19
	ShoulderPressExerciseName_SmithMachineOverheadPress                ShoulderPressExerciseName = 20
	ShoulderPressExerciseName_SplitStanceHammerCurlToPress             ShoulderPressExerciseName = 21
	ShoulderPressExerciseName_SwissBallDumbbellShoulderPress           ShoulderPressExerciseName = 22
	ShoulderPressExerciseName_WeightPlateFrontRaise                    ShoulderPressExerciseName = 23
	ShoulderPressExerciseName_Invalid                                  ShoulderPressExerciseName = 65535
)

var ShoulderPressExerciseNameNames = map[ShoulderPressExerciseName]string{
	ShoulderPressExerciseName_AlternatingDumbbellShoulderPress:         "AlternatingDumbbellShoulderPress",
	ShoulderPressExerciseName_ArnoldPress:                              "ArnoldPress",
	ShoulderPressExerciseName_BarbellFrontSquatToPushPress:             "BarbellFrontSquatToPushPress",
	ShoulderPressExerciseName_BarbellPushPress:                         "BarbellPushPress",
	ShoulderPressExerciseName_BarbellShoulderPress:                     "BarbellShoulderPress",
	ShoulderPressExerciseName_DeadCurlPress:                            "DeadCurlPress",
	ShoulderPressExerciseName_DumbbellAlternatingShoulderPressAndTwist: "DumbbellAlternatingShoulderPressAndTwist",
	ShoulderPressExerciseName_DumbbellHammerCurlToLungeToPress:         "DumbbellHammerCurlToLungeToPress",
	ShoulderPressExerciseName_DumbbellPushPress:                        "DumbbellPushPress",
	ShoulderPressExerciseName_FloorInvertedShoulderPress:               "FloorInvertedShoulderPress",
	ShoulderPressExerciseName_WeightedFloorInvertedShoulderPress:       "WeightedFloorInvertedShoulderPress",
	ShoulderPressExerciseName_InvertedShoulderPress:                    "InvertedShoulderPress",
	ShoulderPressExerciseName_WeightedInvertedShoulderPress:            "WeightedInvertedShoulderPress",
	ShoulderPressExerciseName_OneArmPushPress:                          "OneArmPushPress",
	ShoulderPressExerciseName_OverheadBarbellPress:                     "OverheadBarbellPress",
	ShoulderPressExerciseName_OverheadDumbbellPress:                    "OverheadDumbbellPress",
	ShoulderPressExerciseName_SeatedBarbellShoulderPress:               "SeatedBarbellShoulderPress",
	ShoulderPressExerciseName_SeatedDumbbellShoulderPress:              "SeatedDumbbellShoulderPress",
	ShoulderPressExerciseName_SingleArmDumbbellShoulderPress:           "SingleArmDumbbellShoulderPress",
	ShoulderPressExerciseName_SingleArmStepUpAndPress:                  "SingleArmStepUpAndPress",
	ShoulderPressExerciseName_SmithMachineOverheadPress:                "SmithMachineOverheadPress",
	ShoulderPressExerciseName_SplitStanceHammerCurlToPress:             "SplitStanceHammerCurlToPress",
	ShoulderPressExerciseName_SwissBallDumbbellShoulderPress:           "SwissBallDumbbellShoulderPress",
	ShoulderPressExerciseName_WeightPlateFrontRaise:                    "WeightPlateFrontRaise",
}

func (k ShoulderPressExerciseName) String() string {
	if uint(k) < uint(len(ShoulderPressExerciseNameNames)) {
		if v, ok := ShoulderPressExerciseNameNames[k]; ok {
			return v
		}
	}
	return "ShoulderPressExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var ShoulderPressExerciseNameValues = map[string]ShoulderPressExerciseName{
	"AlternatingDumbbellShoulderPress":         ShoulderPressExerciseName_AlternatingDumbbellShoulderPress,
	"ArnoldPress":                              ShoulderPressExerciseName_ArnoldPress,
	"BarbellFrontSquatToPushPress":             ShoulderPressExerciseName_BarbellFrontSquatToPushPress,
	"BarbellPushPress":                         ShoulderPressExerciseName_BarbellPushPress,
	"BarbellShoulderPress":                     ShoulderPressExerciseName_BarbellShoulderPress,
	"DeadCurlPress":                            ShoulderPressExerciseName_DeadCurlPress,
	"DumbbellAlternatingShoulderPressAndTwist": ShoulderPressExerciseName_DumbbellAlternatingShoulderPressAndTwist,
	"DumbbellHammerCurlToLungeToPress":         ShoulderPressExerciseName_DumbbellHammerCurlToLungeToPress,
	"DumbbellPushPress":                        ShoulderPressExerciseName_DumbbellPushPress,
	"FloorInvertedShoulderPress":               ShoulderPressExerciseName_FloorInvertedShoulderPress,
	"WeightedFloorInvertedShoulderPress":       ShoulderPressExerciseName_WeightedFloorInvertedShoulderPress,
	"InvertedShoulderPress":                    ShoulderPressExerciseName_InvertedShoulderPress,
	"WeightedInvertedShoulderPress":            ShoulderPressExerciseName_WeightedInvertedShoulderPress,
	"OneArmPushPress":                          ShoulderPressExerciseName_OneArmPushPress,
	"OverheadBarbellPress":                     ShoulderPressExerciseName_OverheadBarbellPress,
	"OverheadDumbbellPress":                    ShoulderPressExerciseName_OverheadDumbbellPress,
	"SeatedBarbellShoulderPress":               ShoulderPressExerciseName_SeatedBarbellShoulderPress,
	"SeatedDumbbellShoulderPress":              ShoulderPressExerciseName_SeatedDumbbellShoulderPress,
	"SingleArmDumbbellShoulderPress":           ShoulderPressExerciseName_SingleArmDumbbellShoulderPress,
	"SingleArmStepUpAndPress":                  ShoulderPressExerciseName_SingleArmStepUpAndPress,
	"SmithMachineOverheadPress":                ShoulderPressExerciseName_SmithMachineOverheadPress,
	"SplitStanceHammerCurlToPress":             ShoulderPressExerciseName_SplitStanceHammerCurlToPress,
	"SwissBallDumbbellShoulderPress":           ShoulderPressExerciseName_SwissBallDumbbellShoulderPress,
	"WeightPlateFrontRaise":                    ShoulderPressExerciseName_WeightPlateFrontRaise,
}

func ShoulderPressExerciseNameFromString(s string) ShoulderPressExerciseName {
	if v, ok := ShoulderPressExerciseNameValues[s]; ok {
		return v
	}
	return ShoulderPressExerciseName(ShoulderPressExerciseName_Invalid)
}
