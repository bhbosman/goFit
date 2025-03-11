package fitDecl

import strconv "strconv"

type TricepsExtensionExerciseName uint16

const (
	TricepsExtensionExerciseName_BenchDip                                     TricepsExtensionExerciseName = 0
	TricepsExtensionExerciseName_WeightedBenchDip                             TricepsExtensionExerciseName = 1
	TricepsExtensionExerciseName_BodyWeightDip                                TricepsExtensionExerciseName = 2
	TricepsExtensionExerciseName_CableKickback                                TricepsExtensionExerciseName = 3
	TricepsExtensionExerciseName_CableLyingTricepsExtension                   TricepsExtensionExerciseName = 4
	TricepsExtensionExerciseName_CableOverheadTricepsExtension                TricepsExtensionExerciseName = 5
	TricepsExtensionExerciseName_DumbbellKickback                             TricepsExtensionExerciseName = 6
	TricepsExtensionExerciseName_DumbbellLyingTricepsExtension                TricepsExtensionExerciseName = 7
	TricepsExtensionExerciseName_EzBarOverheadTricepsExtension                TricepsExtensionExerciseName = 8
	TricepsExtensionExerciseName_InclineDip                                   TricepsExtensionExerciseName = 9
	TricepsExtensionExerciseName_WeightedInclineDip                           TricepsExtensionExerciseName = 10
	TricepsExtensionExerciseName_InclineEzBarLyingTricepsExtension            TricepsExtensionExerciseName = 11
	TricepsExtensionExerciseName_LyingDumbbellPulloverToExtension             TricepsExtensionExerciseName = 12
	TricepsExtensionExerciseName_LyingEzBarTricepsExtension                   TricepsExtensionExerciseName = 13
	TricepsExtensionExerciseName_LyingTricepsExtensionToCloseGripBenchPress   TricepsExtensionExerciseName = 14
	TricepsExtensionExerciseName_OverheadDumbbellTricepsExtension             TricepsExtensionExerciseName = 15
	TricepsExtensionExerciseName_RecliningTricepsPress                        TricepsExtensionExerciseName = 16
	TricepsExtensionExerciseName_ReverseGripPressdown                         TricepsExtensionExerciseName = 17
	TricepsExtensionExerciseName_ReverseGripTricepsPressdown                  TricepsExtensionExerciseName = 18
	TricepsExtensionExerciseName_RopePressdown                                TricepsExtensionExerciseName = 19
	TricepsExtensionExerciseName_SeatedBarbellOverheadTricepsExtension        TricepsExtensionExerciseName = 20
	TricepsExtensionExerciseName_SeatedDumbbellOverheadTricepsExtension       TricepsExtensionExerciseName = 21
	TricepsExtensionExerciseName_SeatedEzBarOverheadTricepsExtension          TricepsExtensionExerciseName = 22
	TricepsExtensionExerciseName_SeatedSingleArmOverheadDumbbellExtension     TricepsExtensionExerciseName = 23
	TricepsExtensionExerciseName_SingleArmDumbbellOverheadTricepsExtension    TricepsExtensionExerciseName = 24
	TricepsExtensionExerciseName_SingleDumbbellSeatedOverheadTricepsExtension TricepsExtensionExerciseName = 25
	TricepsExtensionExerciseName_SingleLegBenchDipAndKick                     TricepsExtensionExerciseName = 26
	TricepsExtensionExerciseName_WeightedSingleLegBenchDipAndKick             TricepsExtensionExerciseName = 27
	TricepsExtensionExerciseName_SingleLegDip                                 TricepsExtensionExerciseName = 28
	TricepsExtensionExerciseName_WeightedSingleLegDip                         TricepsExtensionExerciseName = 29
	TricepsExtensionExerciseName_StaticLyingTricepsExtension                  TricepsExtensionExerciseName = 30
	TricepsExtensionExerciseName_SuspendedDip                                 TricepsExtensionExerciseName = 31
	TricepsExtensionExerciseName_WeightedSuspendedDip                         TricepsExtensionExerciseName = 32
	TricepsExtensionExerciseName_SwissBallDumbbellLyingTricepsExtension       TricepsExtensionExerciseName = 33
	TricepsExtensionExerciseName_SwissBallEzBarLyingTricepsExtension          TricepsExtensionExerciseName = 34
	TricepsExtensionExerciseName_SwissBallEzBarOverheadTricepsExtension       TricepsExtensionExerciseName = 35
	TricepsExtensionExerciseName_TabletopDip                                  TricepsExtensionExerciseName = 36
	TricepsExtensionExerciseName_WeightedTabletopDip                          TricepsExtensionExerciseName = 37
	TricepsExtensionExerciseName_TricepsExtensionOnFloor                      TricepsExtensionExerciseName = 38
	TricepsExtensionExerciseName_TricepsPressdown                             TricepsExtensionExerciseName = 39
	TricepsExtensionExerciseName_WeightedDip                                  TricepsExtensionExerciseName = 40
	TricepsExtensionExerciseName_Invalid                                      TricepsExtensionExerciseName = 65535
)

var TricepsExtensionExerciseNameNames = map[TricepsExtensionExerciseName]string{
	TricepsExtensionExerciseName_BenchDip:                                     "BenchDip",
	TricepsExtensionExerciseName_WeightedBenchDip:                             "WeightedBenchDip",
	TricepsExtensionExerciseName_BodyWeightDip:                                "BodyWeightDip",
	TricepsExtensionExerciseName_CableKickback:                                "CableKickback",
	TricepsExtensionExerciseName_CableLyingTricepsExtension:                   "CableLyingTricepsExtension",
	TricepsExtensionExerciseName_CableOverheadTricepsExtension:                "CableOverheadTricepsExtension",
	TricepsExtensionExerciseName_DumbbellKickback:                             "DumbbellKickback",
	TricepsExtensionExerciseName_DumbbellLyingTricepsExtension:                "DumbbellLyingTricepsExtension",
	TricepsExtensionExerciseName_EzBarOverheadTricepsExtension:                "EzBarOverheadTricepsExtension",
	TricepsExtensionExerciseName_InclineDip:                                   "InclineDip",
	TricepsExtensionExerciseName_WeightedInclineDip:                           "WeightedInclineDip",
	TricepsExtensionExerciseName_InclineEzBarLyingTricepsExtension:            "InclineEzBarLyingTricepsExtension",
	TricepsExtensionExerciseName_LyingDumbbellPulloverToExtension:             "LyingDumbbellPulloverToExtension",
	TricepsExtensionExerciseName_LyingEzBarTricepsExtension:                   "LyingEzBarTricepsExtension",
	TricepsExtensionExerciseName_LyingTricepsExtensionToCloseGripBenchPress:   "LyingTricepsExtensionToCloseGripBenchPress",
	TricepsExtensionExerciseName_OverheadDumbbellTricepsExtension:             "OverheadDumbbellTricepsExtension",
	TricepsExtensionExerciseName_RecliningTricepsPress:                        "RecliningTricepsPress",
	TricepsExtensionExerciseName_ReverseGripPressdown:                         "ReverseGripPressdown",
	TricepsExtensionExerciseName_ReverseGripTricepsPressdown:                  "ReverseGripTricepsPressdown",
	TricepsExtensionExerciseName_RopePressdown:                                "RopePressdown",
	TricepsExtensionExerciseName_SeatedBarbellOverheadTricepsExtension:        "SeatedBarbellOverheadTricepsExtension",
	TricepsExtensionExerciseName_SeatedDumbbellOverheadTricepsExtension:       "SeatedDumbbellOverheadTricepsExtension",
	TricepsExtensionExerciseName_SeatedEzBarOverheadTricepsExtension:          "SeatedEzBarOverheadTricepsExtension",
	TricepsExtensionExerciseName_SeatedSingleArmOverheadDumbbellExtension:     "SeatedSingleArmOverheadDumbbellExtension",
	TricepsExtensionExerciseName_SingleArmDumbbellOverheadTricepsExtension:    "SingleArmDumbbellOverheadTricepsExtension",
	TricepsExtensionExerciseName_SingleDumbbellSeatedOverheadTricepsExtension: "SingleDumbbellSeatedOverheadTricepsExtension",
	TricepsExtensionExerciseName_SingleLegBenchDipAndKick:                     "SingleLegBenchDipAndKick",
	TricepsExtensionExerciseName_WeightedSingleLegBenchDipAndKick:             "WeightedSingleLegBenchDipAndKick",
	TricepsExtensionExerciseName_SingleLegDip:                                 "SingleLegDip",
	TricepsExtensionExerciseName_WeightedSingleLegDip:                         "WeightedSingleLegDip",
	TricepsExtensionExerciseName_StaticLyingTricepsExtension:                  "StaticLyingTricepsExtension",
	TricepsExtensionExerciseName_SuspendedDip:                                 "SuspendedDip",
	TricepsExtensionExerciseName_WeightedSuspendedDip:                         "WeightedSuspendedDip",
	TricepsExtensionExerciseName_SwissBallDumbbellLyingTricepsExtension:       "SwissBallDumbbellLyingTricepsExtension",
	TricepsExtensionExerciseName_SwissBallEzBarLyingTricepsExtension:          "SwissBallEzBarLyingTricepsExtension",
	TricepsExtensionExerciseName_SwissBallEzBarOverheadTricepsExtension:       "SwissBallEzBarOverheadTricepsExtension",
	TricepsExtensionExerciseName_TabletopDip:                                  "TabletopDip",
	TricepsExtensionExerciseName_WeightedTabletopDip:                          "WeightedTabletopDip",
	TricepsExtensionExerciseName_TricepsExtensionOnFloor:                      "TricepsExtensionOnFloor",
	TricepsExtensionExerciseName_TricepsPressdown:                             "TricepsPressdown",
	TricepsExtensionExerciseName_WeightedDip:                                  "WeightedDip",
}

func (k TricepsExtensionExerciseName) String() string {
	if uint(k) < uint(len(TricepsExtensionExerciseNameNames)) {
		if v, ok := TricepsExtensionExerciseNameNames[k]; ok {
			return v
		}
	}
	return "TricepsExtensionExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var TricepsExtensionExerciseNameValues = map[string]TricepsExtensionExerciseName{
	"BenchDip":                                     TricepsExtensionExerciseName_BenchDip,
	"WeightedBenchDip":                             TricepsExtensionExerciseName_WeightedBenchDip,
	"BodyWeightDip":                                TricepsExtensionExerciseName_BodyWeightDip,
	"CableKickback":                                TricepsExtensionExerciseName_CableKickback,
	"CableLyingTricepsExtension":                   TricepsExtensionExerciseName_CableLyingTricepsExtension,
	"CableOverheadTricepsExtension":                TricepsExtensionExerciseName_CableOverheadTricepsExtension,
	"DumbbellKickback":                             TricepsExtensionExerciseName_DumbbellKickback,
	"DumbbellLyingTricepsExtension":                TricepsExtensionExerciseName_DumbbellLyingTricepsExtension,
	"EzBarOverheadTricepsExtension":                TricepsExtensionExerciseName_EzBarOverheadTricepsExtension,
	"InclineDip":                                   TricepsExtensionExerciseName_InclineDip,
	"WeightedInclineDip":                           TricepsExtensionExerciseName_WeightedInclineDip,
	"InclineEzBarLyingTricepsExtension":            TricepsExtensionExerciseName_InclineEzBarLyingTricepsExtension,
	"LyingDumbbellPulloverToExtension":             TricepsExtensionExerciseName_LyingDumbbellPulloverToExtension,
	"LyingEzBarTricepsExtension":                   TricepsExtensionExerciseName_LyingEzBarTricepsExtension,
	"LyingTricepsExtensionToCloseGripBenchPress":   TricepsExtensionExerciseName_LyingTricepsExtensionToCloseGripBenchPress,
	"OverheadDumbbellTricepsExtension":             TricepsExtensionExerciseName_OverheadDumbbellTricepsExtension,
	"RecliningTricepsPress":                        TricepsExtensionExerciseName_RecliningTricepsPress,
	"ReverseGripPressdown":                         TricepsExtensionExerciseName_ReverseGripPressdown,
	"ReverseGripTricepsPressdown":                  TricepsExtensionExerciseName_ReverseGripTricepsPressdown,
	"RopePressdown":                                TricepsExtensionExerciseName_RopePressdown,
	"SeatedBarbellOverheadTricepsExtension":        TricepsExtensionExerciseName_SeatedBarbellOverheadTricepsExtension,
	"SeatedDumbbellOverheadTricepsExtension":       TricepsExtensionExerciseName_SeatedDumbbellOverheadTricepsExtension,
	"SeatedEzBarOverheadTricepsExtension":          TricepsExtensionExerciseName_SeatedEzBarOverheadTricepsExtension,
	"SeatedSingleArmOverheadDumbbellExtension":     TricepsExtensionExerciseName_SeatedSingleArmOverheadDumbbellExtension,
	"SingleArmDumbbellOverheadTricepsExtension":    TricepsExtensionExerciseName_SingleArmDumbbellOverheadTricepsExtension,
	"SingleDumbbellSeatedOverheadTricepsExtension": TricepsExtensionExerciseName_SingleDumbbellSeatedOverheadTricepsExtension,
	"SingleLegBenchDipAndKick":                     TricepsExtensionExerciseName_SingleLegBenchDipAndKick,
	"WeightedSingleLegBenchDipAndKick":             TricepsExtensionExerciseName_WeightedSingleLegBenchDipAndKick,
	"SingleLegDip":                                 TricepsExtensionExerciseName_SingleLegDip,
	"WeightedSingleLegDip":                         TricepsExtensionExerciseName_WeightedSingleLegDip,
	"StaticLyingTricepsExtension":                  TricepsExtensionExerciseName_StaticLyingTricepsExtension,
	"SuspendedDip":                                 TricepsExtensionExerciseName_SuspendedDip,
	"WeightedSuspendedDip":                         TricepsExtensionExerciseName_WeightedSuspendedDip,
	"SwissBallDumbbellLyingTricepsExtension":       TricepsExtensionExerciseName_SwissBallDumbbellLyingTricepsExtension,
	"SwissBallEzBarLyingTricepsExtension":          TricepsExtensionExerciseName_SwissBallEzBarLyingTricepsExtension,
	"SwissBallEzBarOverheadTricepsExtension":       TricepsExtensionExerciseName_SwissBallEzBarOverheadTricepsExtension,
	"TabletopDip":                                  TricepsExtensionExerciseName_TabletopDip,
	"WeightedTabletopDip":                          TricepsExtensionExerciseName_WeightedTabletopDip,
	"TricepsExtensionOnFloor":                      TricepsExtensionExerciseName_TricepsExtensionOnFloor,
	"TricepsPressdown":                             TricepsExtensionExerciseName_TricepsPressdown,
	"WeightedDip":                                  TricepsExtensionExerciseName_WeightedDip,
}

func TricepsExtensionExerciseNameFromString(s string) TricepsExtensionExerciseName {
	if v, ok := TricepsExtensionExerciseNameValues[s]; ok {
		return v
	}
	return TricepsExtensionExerciseName(TricepsExtensionExerciseName_Invalid)
}
