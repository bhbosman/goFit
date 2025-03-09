package fitDecl

import strconv "strconv"

type LegCurlExerciseName uint16

const (
	LegCurlExerciseName_LegCurl                     LegCurlExerciseName = 0
	LegCurlExerciseName_WeightedLegCurl             LegCurlExerciseName = 1
	LegCurlExerciseName_GoodMorning                 LegCurlExerciseName = 2
	LegCurlExerciseName_SeatedBarbellGoodMorning    LegCurlExerciseName = 3
	LegCurlExerciseName_SingleLegBarbellGoodMorning LegCurlExerciseName = 4
	LegCurlExerciseName_SingleLegSlidingLegCurl     LegCurlExerciseName = 5
	LegCurlExerciseName_SlidingLegCurl              LegCurlExerciseName = 6
	LegCurlExerciseName_SplitBarbellGoodMorning     LegCurlExerciseName = 7
	LegCurlExerciseName_SplitStanceExtension        LegCurlExerciseName = 8
	LegCurlExerciseName_StaggeredStanceGoodMorning  LegCurlExerciseName = 9
	LegCurlExerciseName_SwissBallHipRaiseAndLegCurl LegCurlExerciseName = 10
	LegCurlExerciseName_ZercherGoodMorning          LegCurlExerciseName = 11
	LegCurlExerciseName_Invalid                     LegCurlExerciseName = 65535
)

var LegCurlExerciseNameNames = map[LegCurlExerciseName]string{
	LegCurlExerciseName_LegCurl:                     "LegCurl",
	LegCurlExerciseName_WeightedLegCurl:             "WeightedLegCurl",
	LegCurlExerciseName_GoodMorning:                 "GoodMorning",
	LegCurlExerciseName_SeatedBarbellGoodMorning:    "SeatedBarbellGoodMorning",
	LegCurlExerciseName_SingleLegBarbellGoodMorning: "SingleLegBarbellGoodMorning",
	LegCurlExerciseName_SingleLegSlidingLegCurl:     "SingleLegSlidingLegCurl",
	LegCurlExerciseName_SlidingLegCurl:              "SlidingLegCurl",
	LegCurlExerciseName_SplitBarbellGoodMorning:     "SplitBarbellGoodMorning",
	LegCurlExerciseName_SplitStanceExtension:        "SplitStanceExtension",
	LegCurlExerciseName_StaggeredStanceGoodMorning:  "StaggeredStanceGoodMorning",
	LegCurlExerciseName_SwissBallHipRaiseAndLegCurl: "SwissBallHipRaiseAndLegCurl",
	LegCurlExerciseName_ZercherGoodMorning:          "ZercherGoodMorning",
}

func (k LegCurlExerciseName) String() string {
	if uint(k) < uint(len(LegCurlExerciseNameNames)) {
		if v, ok := LegCurlExerciseNameNames[k]; ok {
			return v
		}
	}
	return "LegCurlExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var LegCurlExerciseNameValues = map[string]LegCurlExerciseName{
	"LegCurl":                     LegCurlExerciseName_LegCurl,
	"WeightedLegCurl":             LegCurlExerciseName_WeightedLegCurl,
	"GoodMorning":                 LegCurlExerciseName_GoodMorning,
	"SeatedBarbellGoodMorning":    LegCurlExerciseName_SeatedBarbellGoodMorning,
	"SingleLegBarbellGoodMorning": LegCurlExerciseName_SingleLegBarbellGoodMorning,
	"SingleLegSlidingLegCurl":     LegCurlExerciseName_SingleLegSlidingLegCurl,
	"SlidingLegCurl":              LegCurlExerciseName_SlidingLegCurl,
	"SplitBarbellGoodMorning":     LegCurlExerciseName_SplitBarbellGoodMorning,
	"SplitStanceExtension":        LegCurlExerciseName_SplitStanceExtension,
	"StaggeredStanceGoodMorning":  LegCurlExerciseName_StaggeredStanceGoodMorning,
	"SwissBallHipRaiseAndLegCurl": LegCurlExerciseName_SwissBallHipRaiseAndLegCurl,
	"ZercherGoodMorning":          LegCurlExerciseName_ZercherGoodMorning,
}

func LegCurlExerciseNameFromString(s string) LegCurlExerciseName {
	if v, ok := LegCurlExerciseNameValues[s]; ok {
		return v
	}
	return LegCurlExerciseName(LegCurlExerciseName_Invalid)
}
