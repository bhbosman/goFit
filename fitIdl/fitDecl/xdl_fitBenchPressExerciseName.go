package fitDecl

import strconv "strconv"

type BenchPressExerciseName uint16

const (
	BenchPressExerciseName_AlternatingDumbbellChestPressOnSwissBall BenchPressExerciseName = 0
	BenchPressExerciseName_BarbellBenchPress                        BenchPressExerciseName = 1
	BenchPressExerciseName_BarbellBoardBenchPress                   BenchPressExerciseName = 2
	BenchPressExerciseName_BarbellFloorPress                        BenchPressExerciseName = 3
	BenchPressExerciseName_CloseGripBarbellBenchPress               BenchPressExerciseName = 4
	BenchPressExerciseName_DeclineDumbbellBenchPress                BenchPressExerciseName = 5
	BenchPressExerciseName_DumbbellBenchPress                       BenchPressExerciseName = 6
	BenchPressExerciseName_DumbbellFloorPress                       BenchPressExerciseName = 7
	BenchPressExerciseName_InclineBarbellBenchPress                 BenchPressExerciseName = 8
	BenchPressExerciseName_InclineDumbbellBenchPress                BenchPressExerciseName = 9
	BenchPressExerciseName_InclineSmithMachineBenchPress            BenchPressExerciseName = 10
	BenchPressExerciseName_IsometricBarbellBenchPress               BenchPressExerciseName = 11
	BenchPressExerciseName_KettlebellChestPress                     BenchPressExerciseName = 12
	BenchPressExerciseName_NeutralGripDumbbellBenchPress            BenchPressExerciseName = 13
	BenchPressExerciseName_NeutralGripDumbbellInclineBenchPress     BenchPressExerciseName = 14
	BenchPressExerciseName_OneArmFloorPress                         BenchPressExerciseName = 15
	BenchPressExerciseName_WeightedOneArmFloorPress                 BenchPressExerciseName = 16
	BenchPressExerciseName_PartialLockout                           BenchPressExerciseName = 17
	BenchPressExerciseName_ReverseGripBarbellBenchPress             BenchPressExerciseName = 18
	BenchPressExerciseName_ReverseGripInclineBenchPress             BenchPressExerciseName = 19
	BenchPressExerciseName_SingleArmCableChestPress                 BenchPressExerciseName = 20
	BenchPressExerciseName_SingleArmDumbbellBenchPress              BenchPressExerciseName = 21
	BenchPressExerciseName_SmithMachineBenchPress                   BenchPressExerciseName = 22
	BenchPressExerciseName_SwissBallDumbbellChestPress              BenchPressExerciseName = 23
	BenchPressExerciseName_TripleStopBarbellBenchPress              BenchPressExerciseName = 24
	BenchPressExerciseName_WideGripBarbellBenchPress                BenchPressExerciseName = 25
	BenchPressExerciseName_AlternatingDumbbellChestPress            BenchPressExerciseName = 26
	BenchPressExerciseName_Invalid                                  BenchPressExerciseName = 65535
)

var BenchPressExerciseNameNames = map[BenchPressExerciseName]string{
	BenchPressExerciseName_AlternatingDumbbellChestPressOnSwissBall: "AlternatingDumbbellChestPressOnSwissBall",
	BenchPressExerciseName_BarbellBenchPress:                        "BarbellBenchPress",
	BenchPressExerciseName_BarbellBoardBenchPress:                   "BarbellBoardBenchPress",
	BenchPressExerciseName_BarbellFloorPress:                        "BarbellFloorPress",
	BenchPressExerciseName_CloseGripBarbellBenchPress:               "CloseGripBarbellBenchPress",
	BenchPressExerciseName_DeclineDumbbellBenchPress:                "DeclineDumbbellBenchPress",
	BenchPressExerciseName_DumbbellBenchPress:                       "DumbbellBenchPress",
	BenchPressExerciseName_DumbbellFloorPress:                       "DumbbellFloorPress",
	BenchPressExerciseName_InclineBarbellBenchPress:                 "InclineBarbellBenchPress",
	BenchPressExerciseName_InclineDumbbellBenchPress:                "InclineDumbbellBenchPress",
	BenchPressExerciseName_InclineSmithMachineBenchPress:            "InclineSmithMachineBenchPress",
	BenchPressExerciseName_IsometricBarbellBenchPress:               "IsometricBarbellBenchPress",
	BenchPressExerciseName_KettlebellChestPress:                     "KettlebellChestPress",
	BenchPressExerciseName_NeutralGripDumbbellBenchPress:            "NeutralGripDumbbellBenchPress",
	BenchPressExerciseName_NeutralGripDumbbellInclineBenchPress:     "NeutralGripDumbbellInclineBenchPress",
	BenchPressExerciseName_OneArmFloorPress:                         "OneArmFloorPress",
	BenchPressExerciseName_WeightedOneArmFloorPress:                 "WeightedOneArmFloorPress",
	BenchPressExerciseName_PartialLockout:                           "PartialLockout",
	BenchPressExerciseName_ReverseGripBarbellBenchPress:             "ReverseGripBarbellBenchPress",
	BenchPressExerciseName_ReverseGripInclineBenchPress:             "ReverseGripInclineBenchPress",
	BenchPressExerciseName_SingleArmCableChestPress:                 "SingleArmCableChestPress",
	BenchPressExerciseName_SingleArmDumbbellBenchPress:              "SingleArmDumbbellBenchPress",
	BenchPressExerciseName_SmithMachineBenchPress:                   "SmithMachineBenchPress",
	BenchPressExerciseName_SwissBallDumbbellChestPress:              "SwissBallDumbbellChestPress",
	BenchPressExerciseName_TripleStopBarbellBenchPress:              "TripleStopBarbellBenchPress",
	BenchPressExerciseName_WideGripBarbellBenchPress:                "WideGripBarbellBenchPress",
	BenchPressExerciseName_AlternatingDumbbellChestPress:            "AlternatingDumbbellChestPress",
}

func (k BenchPressExerciseName) String() string {
	if uint(k) < uint(len(BenchPressExerciseNameNames)) {
		if v, ok := BenchPressExerciseNameNames[k]; ok {
			return v
		}
	}
	return "BenchPressExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var BenchPressExerciseNameValues = map[string]BenchPressExerciseName{
	"AlternatingDumbbellChestPressOnSwissBall": BenchPressExerciseName_AlternatingDumbbellChestPressOnSwissBall,
	"BarbellBenchPress":                        BenchPressExerciseName_BarbellBenchPress,
	"BarbellBoardBenchPress":                   BenchPressExerciseName_BarbellBoardBenchPress,
	"BarbellFloorPress":                        BenchPressExerciseName_BarbellFloorPress,
	"CloseGripBarbellBenchPress":               BenchPressExerciseName_CloseGripBarbellBenchPress,
	"DeclineDumbbellBenchPress":                BenchPressExerciseName_DeclineDumbbellBenchPress,
	"DumbbellBenchPress":                       BenchPressExerciseName_DumbbellBenchPress,
	"DumbbellFloorPress":                       BenchPressExerciseName_DumbbellFloorPress,
	"InclineBarbellBenchPress":                 BenchPressExerciseName_InclineBarbellBenchPress,
	"InclineDumbbellBenchPress":                BenchPressExerciseName_InclineDumbbellBenchPress,
	"InclineSmithMachineBenchPress":            BenchPressExerciseName_InclineSmithMachineBenchPress,
	"IsometricBarbellBenchPress":               BenchPressExerciseName_IsometricBarbellBenchPress,
	"KettlebellChestPress":                     BenchPressExerciseName_KettlebellChestPress,
	"NeutralGripDumbbellBenchPress":            BenchPressExerciseName_NeutralGripDumbbellBenchPress,
	"NeutralGripDumbbellInclineBenchPress":     BenchPressExerciseName_NeutralGripDumbbellInclineBenchPress,
	"OneArmFloorPress":                         BenchPressExerciseName_OneArmFloorPress,
	"WeightedOneArmFloorPress":                 BenchPressExerciseName_WeightedOneArmFloorPress,
	"PartialLockout":                           BenchPressExerciseName_PartialLockout,
	"ReverseGripBarbellBenchPress":             BenchPressExerciseName_ReverseGripBarbellBenchPress,
	"ReverseGripInclineBenchPress":             BenchPressExerciseName_ReverseGripInclineBenchPress,
	"SingleArmCableChestPress":                 BenchPressExerciseName_SingleArmCableChestPress,
	"SingleArmDumbbellBenchPress":              BenchPressExerciseName_SingleArmDumbbellBenchPress,
	"SmithMachineBenchPress":                   BenchPressExerciseName_SmithMachineBenchPress,
	"SwissBallDumbbellChestPress":              BenchPressExerciseName_SwissBallDumbbellChestPress,
	"TripleStopBarbellBenchPress":              BenchPressExerciseName_TripleStopBarbellBenchPress,
	"WideGripBarbellBenchPress":                BenchPressExerciseName_WideGripBarbellBenchPress,
	"AlternatingDumbbellChestPress":            BenchPressExerciseName_AlternatingDumbbellChestPress,
}

func BenchPressExerciseNameFromString(s string) BenchPressExerciseName {
	if v, ok := BenchPressExerciseNameValues[s]; ok {
		return v
	}
	return BenchPressExerciseName(BenchPressExerciseName_Invalid)
}
