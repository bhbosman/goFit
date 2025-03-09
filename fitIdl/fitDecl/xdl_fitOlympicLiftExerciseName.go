package fitDecl

import strconv "strconv"

type OlympicLiftExerciseName uint16

const (
	OlympicLiftExerciseName_BarbellHangPowerClean      OlympicLiftExerciseName = 0
	OlympicLiftExerciseName_BarbellHangSquatClean      OlympicLiftExerciseName = 1
	OlympicLiftExerciseName_BarbellPowerClean          OlympicLiftExerciseName = 2
	OlympicLiftExerciseName_BarbellPowerSnatch         OlympicLiftExerciseName = 3
	OlympicLiftExerciseName_BarbellSquatClean          OlympicLiftExerciseName = 4
	OlympicLiftExerciseName_CleanAndJerk               OlympicLiftExerciseName = 5
	OlympicLiftExerciseName_BarbellHangPowerSnatch     OlympicLiftExerciseName = 6
	OlympicLiftExerciseName_BarbellHangPull            OlympicLiftExerciseName = 7
	OlympicLiftExerciseName_BarbellHighPull            OlympicLiftExerciseName = 8
	OlympicLiftExerciseName_BarbellSnatch              OlympicLiftExerciseName = 9
	OlympicLiftExerciseName_BarbellSplitJerk           OlympicLiftExerciseName = 10
	OlympicLiftExerciseName_Clean                      OlympicLiftExerciseName = 11
	OlympicLiftExerciseName_DumbbellClean              OlympicLiftExerciseName = 12
	OlympicLiftExerciseName_DumbbellHangPull           OlympicLiftExerciseName = 13
	OlympicLiftExerciseName_OneHandDumbbellSplitSnatch OlympicLiftExerciseName = 14
	OlympicLiftExerciseName_PushJerk                   OlympicLiftExerciseName = 15
	OlympicLiftExerciseName_SingleArmDumbbellSnatch    OlympicLiftExerciseName = 16
	OlympicLiftExerciseName_SingleArmHangSnatch        OlympicLiftExerciseName = 17
	OlympicLiftExerciseName_SingleArmKettlebellSnatch  OlympicLiftExerciseName = 18
	OlympicLiftExerciseName_SplitJerk                  OlympicLiftExerciseName = 19
	OlympicLiftExerciseName_SquatCleanAndJerk          OlympicLiftExerciseName = 20
	OlympicLiftExerciseName_Invalid                    OlympicLiftExerciseName = 65535
)

var OlympicLiftExerciseNameNames = map[OlympicLiftExerciseName]string{
	OlympicLiftExerciseName_BarbellHangPowerClean:      "BarbellHangPowerClean",
	OlympicLiftExerciseName_BarbellHangSquatClean:      "BarbellHangSquatClean",
	OlympicLiftExerciseName_BarbellPowerClean:          "BarbellPowerClean",
	OlympicLiftExerciseName_BarbellPowerSnatch:         "BarbellPowerSnatch",
	OlympicLiftExerciseName_BarbellSquatClean:          "BarbellSquatClean",
	OlympicLiftExerciseName_CleanAndJerk:               "CleanAndJerk",
	OlympicLiftExerciseName_BarbellHangPowerSnatch:     "BarbellHangPowerSnatch",
	OlympicLiftExerciseName_BarbellHangPull:            "BarbellHangPull",
	OlympicLiftExerciseName_BarbellHighPull:            "BarbellHighPull",
	OlympicLiftExerciseName_BarbellSnatch:              "BarbellSnatch",
	OlympicLiftExerciseName_BarbellSplitJerk:           "BarbellSplitJerk",
	OlympicLiftExerciseName_Clean:                      "Clean",
	OlympicLiftExerciseName_DumbbellClean:              "DumbbellClean",
	OlympicLiftExerciseName_DumbbellHangPull:           "DumbbellHangPull",
	OlympicLiftExerciseName_OneHandDumbbellSplitSnatch: "OneHandDumbbellSplitSnatch",
	OlympicLiftExerciseName_PushJerk:                   "PushJerk",
	OlympicLiftExerciseName_SingleArmDumbbellSnatch:    "SingleArmDumbbellSnatch",
	OlympicLiftExerciseName_SingleArmHangSnatch:        "SingleArmHangSnatch",
	OlympicLiftExerciseName_SingleArmKettlebellSnatch:  "SingleArmKettlebellSnatch",
	OlympicLiftExerciseName_SplitJerk:                  "SplitJerk",
	OlympicLiftExerciseName_SquatCleanAndJerk:          "SquatCleanAndJerk",
}

func (k OlympicLiftExerciseName) String() string {
	if uint(k) < uint(len(OlympicLiftExerciseNameNames)) {
		if v, ok := OlympicLiftExerciseNameNames[k]; ok {
			return v
		}
	}
	return "OlympicLiftExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var OlympicLiftExerciseNameValues = map[string]OlympicLiftExerciseName{
	"BarbellHangPowerClean":      OlympicLiftExerciseName_BarbellHangPowerClean,
	"BarbellHangSquatClean":      OlympicLiftExerciseName_BarbellHangSquatClean,
	"BarbellPowerClean":          OlympicLiftExerciseName_BarbellPowerClean,
	"BarbellPowerSnatch":         OlympicLiftExerciseName_BarbellPowerSnatch,
	"BarbellSquatClean":          OlympicLiftExerciseName_BarbellSquatClean,
	"CleanAndJerk":               OlympicLiftExerciseName_CleanAndJerk,
	"BarbellHangPowerSnatch":     OlympicLiftExerciseName_BarbellHangPowerSnatch,
	"BarbellHangPull":            OlympicLiftExerciseName_BarbellHangPull,
	"BarbellHighPull":            OlympicLiftExerciseName_BarbellHighPull,
	"BarbellSnatch":              OlympicLiftExerciseName_BarbellSnatch,
	"BarbellSplitJerk":           OlympicLiftExerciseName_BarbellSplitJerk,
	"Clean":                      OlympicLiftExerciseName_Clean,
	"DumbbellClean":              OlympicLiftExerciseName_DumbbellClean,
	"DumbbellHangPull":           OlympicLiftExerciseName_DumbbellHangPull,
	"OneHandDumbbellSplitSnatch": OlympicLiftExerciseName_OneHandDumbbellSplitSnatch,
	"PushJerk":                   OlympicLiftExerciseName_PushJerk,
	"SingleArmDumbbellSnatch":    OlympicLiftExerciseName_SingleArmDumbbellSnatch,
	"SingleArmHangSnatch":        OlympicLiftExerciseName_SingleArmHangSnatch,
	"SingleArmKettlebellSnatch":  OlympicLiftExerciseName_SingleArmKettlebellSnatch,
	"SplitJerk":                  OlympicLiftExerciseName_SplitJerk,
	"SquatCleanAndJerk":          OlympicLiftExerciseName_SquatCleanAndJerk,
}

func OlympicLiftExerciseNameFromString(s string) OlympicLiftExerciseName {
	if v, ok := OlympicLiftExerciseNameValues[s]; ok {
		return v
	}
	return OlympicLiftExerciseName(OlympicLiftExerciseName_Invalid)
}
