package fitDecl

import strconv "strconv"

type HyperextensionExerciseName uint16

const (
	HyperextensionExerciseName_BackExtensionWithOppositeArmAndLegReach         HyperextensionExerciseName = 0
	HyperextensionExerciseName_WeightedBackExtensionWithOppositeArmAndLegReach HyperextensionExerciseName = 1
	HyperextensionExerciseName_BaseRotations                                   HyperextensionExerciseName = 2
	HyperextensionExerciseName_WeightedBaseRotations                           HyperextensionExerciseName = 3
	HyperextensionExerciseName_BentKneeReverseHyperextension                   HyperextensionExerciseName = 4
	HyperextensionExerciseName_WeightedBentKneeReverseHyperextension           HyperextensionExerciseName = 5
	HyperextensionExerciseName_HollowHoldAndRoll                               HyperextensionExerciseName = 6
	HyperextensionExerciseName_WeightedHollowHoldAndRoll                       HyperextensionExerciseName = 7
	HyperextensionExerciseName_Kicks                                           HyperextensionExerciseName = 8
	HyperextensionExerciseName_WeightedKicks                                   HyperextensionExerciseName = 9
	HyperextensionExerciseName_KneeRaises                                      HyperextensionExerciseName = 10
	HyperextensionExerciseName_WeightedKneeRaises                              HyperextensionExerciseName = 11
	HyperextensionExerciseName_KneelingSuperman                                HyperextensionExerciseName = 12
	HyperextensionExerciseName_WeightedKneelingSuperman                        HyperextensionExerciseName = 13
	HyperextensionExerciseName_LatPullDownWithRow                              HyperextensionExerciseName = 14
	HyperextensionExerciseName_MedicineBallDeadliftToReach                     HyperextensionExerciseName = 15
	HyperextensionExerciseName_OneArmOneLegRow                                 HyperextensionExerciseName = 16
	HyperextensionExerciseName_OneArmRowWithBand                               HyperextensionExerciseName = 17
	HyperextensionExerciseName_OverheadLungeWithMedicineBall                   HyperextensionExerciseName = 18
	HyperextensionExerciseName_PlankKneeTucks                                  HyperextensionExerciseName = 19
	HyperextensionExerciseName_WeightedPlankKneeTucks                          HyperextensionExerciseName = 20
	HyperextensionExerciseName_SideStep                                        HyperextensionExerciseName = 21
	HyperextensionExerciseName_WeightedSideStep                                HyperextensionExerciseName = 22
	HyperextensionExerciseName_SingleLegBackExtension                          HyperextensionExerciseName = 23
	HyperextensionExerciseName_WeightedSingleLegBackExtension                  HyperextensionExerciseName = 24
	HyperextensionExerciseName_SpineExtension                                  HyperextensionExerciseName = 25
	HyperextensionExerciseName_WeightedSpineExtension                          HyperextensionExerciseName = 26
	HyperextensionExerciseName_StaticBackExtension                             HyperextensionExerciseName = 27
	HyperextensionExerciseName_WeightedStaticBackExtension                     HyperextensionExerciseName = 28
	HyperextensionExerciseName_SupermanFromFloor                               HyperextensionExerciseName = 29
	HyperextensionExerciseName_WeightedSupermanFromFloor                       HyperextensionExerciseName = 30
	HyperextensionExerciseName_SwissBallBackExtension                          HyperextensionExerciseName = 31
	HyperextensionExerciseName_WeightedSwissBallBackExtension                  HyperextensionExerciseName = 32
	HyperextensionExerciseName_SwissBallHyperextension                         HyperextensionExerciseName = 33
	HyperextensionExerciseName_WeightedSwissBallHyperextension                 HyperextensionExerciseName = 34
	HyperextensionExerciseName_SwissBallOppositeArmAndLegLift                  HyperextensionExerciseName = 35
	HyperextensionExerciseName_WeightedSwissBallOppositeArmAndLegLift          HyperextensionExerciseName = 36
	HyperextensionExerciseName_SupermanOnSwissBall                             HyperextensionExerciseName = 37
	HyperextensionExerciseName_Cobra                                           HyperextensionExerciseName = 38
	HyperextensionExerciseName_Invalid                                         HyperextensionExerciseName = 65535
)

var HyperextensionExerciseNameNames = map[HyperextensionExerciseName]string{
	HyperextensionExerciseName_BackExtensionWithOppositeArmAndLegReach:         "BackExtensionWithOppositeArmAndLegReach",
	HyperextensionExerciseName_WeightedBackExtensionWithOppositeArmAndLegReach: "WeightedBackExtensionWithOppositeArmAndLegReach",
	HyperextensionExerciseName_BaseRotations:                                   "BaseRotations",
	HyperextensionExerciseName_WeightedBaseRotations:                           "WeightedBaseRotations",
	HyperextensionExerciseName_BentKneeReverseHyperextension:                   "BentKneeReverseHyperextension",
	HyperextensionExerciseName_WeightedBentKneeReverseHyperextension:           "WeightedBentKneeReverseHyperextension",
	HyperextensionExerciseName_HollowHoldAndRoll:                               "HollowHoldAndRoll",
	HyperextensionExerciseName_WeightedHollowHoldAndRoll:                       "WeightedHollowHoldAndRoll",
	HyperextensionExerciseName_Kicks:                                           "Kicks",
	HyperextensionExerciseName_WeightedKicks:                                   "WeightedKicks",
	HyperextensionExerciseName_KneeRaises:                                      "KneeRaises",
	HyperextensionExerciseName_WeightedKneeRaises:                              "WeightedKneeRaises",
	HyperextensionExerciseName_KneelingSuperman:                                "KneelingSuperman",
	HyperextensionExerciseName_WeightedKneelingSuperman:                        "WeightedKneelingSuperman",
	HyperextensionExerciseName_LatPullDownWithRow:                              "LatPullDownWithRow",
	HyperextensionExerciseName_MedicineBallDeadliftToReach:                     "MedicineBallDeadliftToReach",
	HyperextensionExerciseName_OneArmOneLegRow:                                 "OneArmOneLegRow",
	HyperextensionExerciseName_OneArmRowWithBand:                               "OneArmRowWithBand",
	HyperextensionExerciseName_OverheadLungeWithMedicineBall:                   "OverheadLungeWithMedicineBall",
	HyperextensionExerciseName_PlankKneeTucks:                                  "PlankKneeTucks",
	HyperextensionExerciseName_WeightedPlankKneeTucks:                          "WeightedPlankKneeTucks",
	HyperextensionExerciseName_SideStep:                                        "SideStep",
	HyperextensionExerciseName_WeightedSideStep:                                "WeightedSideStep",
	HyperextensionExerciseName_SingleLegBackExtension:                          "SingleLegBackExtension",
	HyperextensionExerciseName_WeightedSingleLegBackExtension:                  "WeightedSingleLegBackExtension",
	HyperextensionExerciseName_SpineExtension:                                  "SpineExtension",
	HyperextensionExerciseName_WeightedSpineExtension:                          "WeightedSpineExtension",
	HyperextensionExerciseName_StaticBackExtension:                             "StaticBackExtension",
	HyperextensionExerciseName_WeightedStaticBackExtension:                     "WeightedStaticBackExtension",
	HyperextensionExerciseName_SupermanFromFloor:                               "SupermanFromFloor",
	HyperextensionExerciseName_WeightedSupermanFromFloor:                       "WeightedSupermanFromFloor",
	HyperextensionExerciseName_SwissBallBackExtension:                          "SwissBallBackExtension",
	HyperextensionExerciseName_WeightedSwissBallBackExtension:                  "WeightedSwissBallBackExtension",
	HyperextensionExerciseName_SwissBallHyperextension:                         "SwissBallHyperextension",
	HyperextensionExerciseName_WeightedSwissBallHyperextension:                 "WeightedSwissBallHyperextension",
	HyperextensionExerciseName_SwissBallOppositeArmAndLegLift:                  "SwissBallOppositeArmAndLegLift",
	HyperextensionExerciseName_WeightedSwissBallOppositeArmAndLegLift:          "WeightedSwissBallOppositeArmAndLegLift",
	HyperextensionExerciseName_SupermanOnSwissBall:                             "SupermanOnSwissBall",
	HyperextensionExerciseName_Cobra:                                           "Cobra",
}

func (k HyperextensionExerciseName) String() string {
	if uint(k) < uint(len(HyperextensionExerciseNameNames)) {
		if v, ok := HyperextensionExerciseNameNames[k]; ok {
			return v
		}
	}
	return "HyperextensionExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var HyperextensionExerciseNameValues = map[string]HyperextensionExerciseName{
	"BackExtensionWithOppositeArmAndLegReach":         HyperextensionExerciseName_BackExtensionWithOppositeArmAndLegReach,
	"WeightedBackExtensionWithOppositeArmAndLegReach": HyperextensionExerciseName_WeightedBackExtensionWithOppositeArmAndLegReach,
	"BaseRotations":                          HyperextensionExerciseName_BaseRotations,
	"WeightedBaseRotations":                  HyperextensionExerciseName_WeightedBaseRotations,
	"BentKneeReverseHyperextension":          HyperextensionExerciseName_BentKneeReverseHyperextension,
	"WeightedBentKneeReverseHyperextension":  HyperextensionExerciseName_WeightedBentKneeReverseHyperextension,
	"HollowHoldAndRoll":                      HyperextensionExerciseName_HollowHoldAndRoll,
	"WeightedHollowHoldAndRoll":              HyperextensionExerciseName_WeightedHollowHoldAndRoll,
	"Kicks":                                  HyperextensionExerciseName_Kicks,
	"WeightedKicks":                          HyperextensionExerciseName_WeightedKicks,
	"KneeRaises":                             HyperextensionExerciseName_KneeRaises,
	"WeightedKneeRaises":                     HyperextensionExerciseName_WeightedKneeRaises,
	"KneelingSuperman":                       HyperextensionExerciseName_KneelingSuperman,
	"WeightedKneelingSuperman":               HyperextensionExerciseName_WeightedKneelingSuperman,
	"LatPullDownWithRow":                     HyperextensionExerciseName_LatPullDownWithRow,
	"MedicineBallDeadliftToReach":            HyperextensionExerciseName_MedicineBallDeadliftToReach,
	"OneArmOneLegRow":                        HyperextensionExerciseName_OneArmOneLegRow,
	"OneArmRowWithBand":                      HyperextensionExerciseName_OneArmRowWithBand,
	"OverheadLungeWithMedicineBall":          HyperextensionExerciseName_OverheadLungeWithMedicineBall,
	"PlankKneeTucks":                         HyperextensionExerciseName_PlankKneeTucks,
	"WeightedPlankKneeTucks":                 HyperextensionExerciseName_WeightedPlankKneeTucks,
	"SideStep":                               HyperextensionExerciseName_SideStep,
	"WeightedSideStep":                       HyperextensionExerciseName_WeightedSideStep,
	"SingleLegBackExtension":                 HyperextensionExerciseName_SingleLegBackExtension,
	"WeightedSingleLegBackExtension":         HyperextensionExerciseName_WeightedSingleLegBackExtension,
	"SpineExtension":                         HyperextensionExerciseName_SpineExtension,
	"WeightedSpineExtension":                 HyperextensionExerciseName_WeightedSpineExtension,
	"StaticBackExtension":                    HyperextensionExerciseName_StaticBackExtension,
	"WeightedStaticBackExtension":            HyperextensionExerciseName_WeightedStaticBackExtension,
	"SupermanFromFloor":                      HyperextensionExerciseName_SupermanFromFloor,
	"WeightedSupermanFromFloor":              HyperextensionExerciseName_WeightedSupermanFromFloor,
	"SwissBallBackExtension":                 HyperextensionExerciseName_SwissBallBackExtension,
	"WeightedSwissBallBackExtension":         HyperextensionExerciseName_WeightedSwissBallBackExtension,
	"SwissBallHyperextension":                HyperextensionExerciseName_SwissBallHyperextension,
	"WeightedSwissBallHyperextension":        HyperextensionExerciseName_WeightedSwissBallHyperextension,
	"SwissBallOppositeArmAndLegLift":         HyperextensionExerciseName_SwissBallOppositeArmAndLegLift,
	"WeightedSwissBallOppositeArmAndLegLift": HyperextensionExerciseName_WeightedSwissBallOppositeArmAndLegLift,
	"SupermanOnSwissBall":                    HyperextensionExerciseName_SupermanOnSwissBall,
	"Cobra":                                  HyperextensionExerciseName_Cobra,
}

func HyperextensionExerciseNameFromString(s string) HyperextensionExerciseName {
	if v, ok := HyperextensionExerciseNameValues[s]; ok {
		return v
	}
	return HyperextensionExerciseName(HyperextensionExerciseName_Invalid)
}
