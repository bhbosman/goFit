package fitDecl

import strconv "strconv"

type HipRaiseExerciseName uint16

const (
	HipRaiseExerciseName_BarbellHipThrustOnFloor                         HipRaiseExerciseName = 0
	HipRaiseExerciseName_BarbellHipThrustWithBench                       HipRaiseExerciseName = 1
	HipRaiseExerciseName_BentKneeSwissBallReverseHipRaise                HipRaiseExerciseName = 2
	HipRaiseExerciseName_WeightedBentKneeSwissBallReverseHipRaise        HipRaiseExerciseName = 3
	HipRaiseExerciseName_BridgeWithLegExtension                          HipRaiseExerciseName = 4
	HipRaiseExerciseName_WeightedBridgeWithLegExtension                  HipRaiseExerciseName = 5
	HipRaiseExerciseName_ClamBridge                                      HipRaiseExerciseName = 6
	HipRaiseExerciseName_FrontKickTabletop                               HipRaiseExerciseName = 7
	HipRaiseExerciseName_WeightedFrontKickTabletop                       HipRaiseExerciseName = 8
	HipRaiseExerciseName_HipExtensionAndCross                            HipRaiseExerciseName = 9
	HipRaiseExerciseName_WeightedHipExtensionAndCross                    HipRaiseExerciseName = 10
	HipRaiseExerciseName_HipRaise                                        HipRaiseExerciseName = 11
	HipRaiseExerciseName_WeightedHipRaise                                HipRaiseExerciseName = 12
	HipRaiseExerciseName_HipRaiseWithFeetOnSwissBall                     HipRaiseExerciseName = 13
	HipRaiseExerciseName_WeightedHipRaiseWithFeetOnSwissBall             HipRaiseExerciseName = 14
	HipRaiseExerciseName_HipRaiseWithHeadOnBosuBall                      HipRaiseExerciseName = 15
	HipRaiseExerciseName_WeightedHipRaiseWithHeadOnBosuBall              HipRaiseExerciseName = 16
	HipRaiseExerciseName_HipRaiseWithHeadOnSwissBall                     HipRaiseExerciseName = 17
	HipRaiseExerciseName_WeightedHipRaiseWithHeadOnSwissBall             HipRaiseExerciseName = 18
	HipRaiseExerciseName_HipRaiseWithKneeSqueeze                         HipRaiseExerciseName = 19
	HipRaiseExerciseName_WeightedHipRaiseWithKneeSqueeze                 HipRaiseExerciseName = 20
	HipRaiseExerciseName_InclineRearLegExtension                         HipRaiseExerciseName = 21
	HipRaiseExerciseName_WeightedInclineRearLegExtension                 HipRaiseExerciseName = 22
	HipRaiseExerciseName_KettlebellSwing                                 HipRaiseExerciseName = 23
	HipRaiseExerciseName_MarchingHipRaise                                HipRaiseExerciseName = 24
	HipRaiseExerciseName_WeightedMarchingHipRaise                        HipRaiseExerciseName = 25
	HipRaiseExerciseName_MarchingHipRaiseWithFeetOnASwissBall            HipRaiseExerciseName = 26
	HipRaiseExerciseName_WeightedMarchingHipRaiseWithFeetOnASwissBall    HipRaiseExerciseName = 27
	HipRaiseExerciseName_ReverseHipRaise                                 HipRaiseExerciseName = 28
	HipRaiseExerciseName_WeightedReverseHipRaise                         HipRaiseExerciseName = 29
	HipRaiseExerciseName_SingleLegHipRaise                               HipRaiseExerciseName = 30
	HipRaiseExerciseName_WeightedSingleLegHipRaise                       HipRaiseExerciseName = 31
	HipRaiseExerciseName_SingleLegHipRaiseWithFootOnBench                HipRaiseExerciseName = 32
	HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnBench        HipRaiseExerciseName = 33
	HipRaiseExerciseName_SingleLegHipRaiseWithFootOnBosuBall             HipRaiseExerciseName = 34
	HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnBosuBall     HipRaiseExerciseName = 35
	HipRaiseExerciseName_SingleLegHipRaiseWithFootOnFoamRoller           HipRaiseExerciseName = 36
	HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnFoamRoller   HipRaiseExerciseName = 37
	HipRaiseExerciseName_SingleLegHipRaiseWithFootOnMedicineBall         HipRaiseExerciseName = 38
	HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnMedicineBall HipRaiseExerciseName = 39
	HipRaiseExerciseName_SingleLegHipRaiseWithHeadOnBosuBall             HipRaiseExerciseName = 40
	HipRaiseExerciseName_WeightedSingleLegHipRaiseWithHeadOnBosuBall     HipRaiseExerciseName = 41
	HipRaiseExerciseName_WeightedClamBridge                              HipRaiseExerciseName = 42
	HipRaiseExerciseName_SingleLegSwissBallHipRaiseAndLegCurl            HipRaiseExerciseName = 43
	HipRaiseExerciseName_Clams                                           HipRaiseExerciseName = 44
	HipRaiseExerciseName_LegCircles                                      HipRaiseExerciseName = 47
	HipRaiseExerciseName_LegLift                                         HipRaiseExerciseName = 48
	HipRaiseExerciseName_LegLiftInExternalRotation                       HipRaiseExerciseName = 49
	HipRaiseExerciseName_Invalid                                         HipRaiseExerciseName = 65535
)

var HipRaiseExerciseNameNames = map[HipRaiseExerciseName]string{
	HipRaiseExerciseName_BarbellHipThrustOnFloor:                         "BarbellHipThrustOnFloor",
	HipRaiseExerciseName_BarbellHipThrustWithBench:                       "BarbellHipThrustWithBench",
	HipRaiseExerciseName_BentKneeSwissBallReverseHipRaise:                "BentKneeSwissBallReverseHipRaise",
	HipRaiseExerciseName_WeightedBentKneeSwissBallReverseHipRaise:        "WeightedBentKneeSwissBallReverseHipRaise",
	HipRaiseExerciseName_BridgeWithLegExtension:                          "BridgeWithLegExtension",
	HipRaiseExerciseName_WeightedBridgeWithLegExtension:                  "WeightedBridgeWithLegExtension",
	HipRaiseExerciseName_ClamBridge:                                      "ClamBridge",
	HipRaiseExerciseName_FrontKickTabletop:                               "FrontKickTabletop",
	HipRaiseExerciseName_WeightedFrontKickTabletop:                       "WeightedFrontKickTabletop",
	HipRaiseExerciseName_HipExtensionAndCross:                            "HipExtensionAndCross",
	HipRaiseExerciseName_WeightedHipExtensionAndCross:                    "WeightedHipExtensionAndCross",
	HipRaiseExerciseName_HipRaise:                                        "HipRaise",
	HipRaiseExerciseName_WeightedHipRaise:                                "WeightedHipRaise",
	HipRaiseExerciseName_HipRaiseWithFeetOnSwissBall:                     "HipRaiseWithFeetOnSwissBall",
	HipRaiseExerciseName_WeightedHipRaiseWithFeetOnSwissBall:             "WeightedHipRaiseWithFeetOnSwissBall",
	HipRaiseExerciseName_HipRaiseWithHeadOnBosuBall:                      "HipRaiseWithHeadOnBosuBall",
	HipRaiseExerciseName_WeightedHipRaiseWithHeadOnBosuBall:              "WeightedHipRaiseWithHeadOnBosuBall",
	HipRaiseExerciseName_HipRaiseWithHeadOnSwissBall:                     "HipRaiseWithHeadOnSwissBall",
	HipRaiseExerciseName_WeightedHipRaiseWithHeadOnSwissBall:             "WeightedHipRaiseWithHeadOnSwissBall",
	HipRaiseExerciseName_HipRaiseWithKneeSqueeze:                         "HipRaiseWithKneeSqueeze",
	HipRaiseExerciseName_WeightedHipRaiseWithKneeSqueeze:                 "WeightedHipRaiseWithKneeSqueeze",
	HipRaiseExerciseName_InclineRearLegExtension:                         "InclineRearLegExtension",
	HipRaiseExerciseName_WeightedInclineRearLegExtension:                 "WeightedInclineRearLegExtension",
	HipRaiseExerciseName_KettlebellSwing:                                 "KettlebellSwing",
	HipRaiseExerciseName_MarchingHipRaise:                                "MarchingHipRaise",
	HipRaiseExerciseName_WeightedMarchingHipRaise:                        "WeightedMarchingHipRaise",
	HipRaiseExerciseName_MarchingHipRaiseWithFeetOnASwissBall:            "MarchingHipRaiseWithFeetOnASwissBall",
	HipRaiseExerciseName_WeightedMarchingHipRaiseWithFeetOnASwissBall:    "WeightedMarchingHipRaiseWithFeetOnASwissBall",
	HipRaiseExerciseName_ReverseHipRaise:                                 "ReverseHipRaise",
	HipRaiseExerciseName_WeightedReverseHipRaise:                         "WeightedReverseHipRaise",
	HipRaiseExerciseName_SingleLegHipRaise:                               "SingleLegHipRaise",
	HipRaiseExerciseName_WeightedSingleLegHipRaise:                       "WeightedSingleLegHipRaise",
	HipRaiseExerciseName_SingleLegHipRaiseWithFootOnBench:                "SingleLegHipRaiseWithFootOnBench",
	HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnBench:        "WeightedSingleLegHipRaiseWithFootOnBench",
	HipRaiseExerciseName_SingleLegHipRaiseWithFootOnBosuBall:             "SingleLegHipRaiseWithFootOnBosuBall",
	HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnBosuBall:     "WeightedSingleLegHipRaiseWithFootOnBosuBall",
	HipRaiseExerciseName_SingleLegHipRaiseWithFootOnFoamRoller:           "SingleLegHipRaiseWithFootOnFoamRoller",
	HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnFoamRoller:   "WeightedSingleLegHipRaiseWithFootOnFoamRoller",
	HipRaiseExerciseName_SingleLegHipRaiseWithFootOnMedicineBall:         "SingleLegHipRaiseWithFootOnMedicineBall",
	HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnMedicineBall: "WeightedSingleLegHipRaiseWithFootOnMedicineBall",
	HipRaiseExerciseName_SingleLegHipRaiseWithHeadOnBosuBall:             "SingleLegHipRaiseWithHeadOnBosuBall",
	HipRaiseExerciseName_WeightedSingleLegHipRaiseWithHeadOnBosuBall:     "WeightedSingleLegHipRaiseWithHeadOnBosuBall",
	HipRaiseExerciseName_WeightedClamBridge:                              "WeightedClamBridge",
	HipRaiseExerciseName_SingleLegSwissBallHipRaiseAndLegCurl:            "SingleLegSwissBallHipRaiseAndLegCurl",
	HipRaiseExerciseName_Clams:                                           "Clams",
	HipRaiseExerciseName_LegCircles:                                      "LegCircles",
	HipRaiseExerciseName_LegLift:                                         "LegLift",
	HipRaiseExerciseName_LegLiftInExternalRotation:                       "LegLiftInExternalRotation",
}

func (k HipRaiseExerciseName) String() string {
	if uint(k) < uint(len(HipRaiseExerciseNameNames)) {
		if v, ok := HipRaiseExerciseNameNames[k]; ok {
			return v
		}
	}
	return "HipRaiseExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var HipRaiseExerciseNameValues = map[string]HipRaiseExerciseName{
	"BarbellHipThrustOnFloor":                         HipRaiseExerciseName_BarbellHipThrustOnFloor,
	"BarbellHipThrustWithBench":                       HipRaiseExerciseName_BarbellHipThrustWithBench,
	"BentKneeSwissBallReverseHipRaise":                HipRaiseExerciseName_BentKneeSwissBallReverseHipRaise,
	"WeightedBentKneeSwissBallReverseHipRaise":        HipRaiseExerciseName_WeightedBentKneeSwissBallReverseHipRaise,
	"BridgeWithLegExtension":                          HipRaiseExerciseName_BridgeWithLegExtension,
	"WeightedBridgeWithLegExtension":                  HipRaiseExerciseName_WeightedBridgeWithLegExtension,
	"ClamBridge":                                      HipRaiseExerciseName_ClamBridge,
	"FrontKickTabletop":                               HipRaiseExerciseName_FrontKickTabletop,
	"WeightedFrontKickTabletop":                       HipRaiseExerciseName_WeightedFrontKickTabletop,
	"HipExtensionAndCross":                            HipRaiseExerciseName_HipExtensionAndCross,
	"WeightedHipExtensionAndCross":                    HipRaiseExerciseName_WeightedHipExtensionAndCross,
	"HipRaise":                                        HipRaiseExerciseName_HipRaise,
	"WeightedHipRaise":                                HipRaiseExerciseName_WeightedHipRaise,
	"HipRaiseWithFeetOnSwissBall":                     HipRaiseExerciseName_HipRaiseWithFeetOnSwissBall,
	"WeightedHipRaiseWithFeetOnSwissBall":             HipRaiseExerciseName_WeightedHipRaiseWithFeetOnSwissBall,
	"HipRaiseWithHeadOnBosuBall":                      HipRaiseExerciseName_HipRaiseWithHeadOnBosuBall,
	"WeightedHipRaiseWithHeadOnBosuBall":              HipRaiseExerciseName_WeightedHipRaiseWithHeadOnBosuBall,
	"HipRaiseWithHeadOnSwissBall":                     HipRaiseExerciseName_HipRaiseWithHeadOnSwissBall,
	"WeightedHipRaiseWithHeadOnSwissBall":             HipRaiseExerciseName_WeightedHipRaiseWithHeadOnSwissBall,
	"HipRaiseWithKneeSqueeze":                         HipRaiseExerciseName_HipRaiseWithKneeSqueeze,
	"WeightedHipRaiseWithKneeSqueeze":                 HipRaiseExerciseName_WeightedHipRaiseWithKneeSqueeze,
	"InclineRearLegExtension":                         HipRaiseExerciseName_InclineRearLegExtension,
	"WeightedInclineRearLegExtension":                 HipRaiseExerciseName_WeightedInclineRearLegExtension,
	"KettlebellSwing":                                 HipRaiseExerciseName_KettlebellSwing,
	"MarchingHipRaise":                                HipRaiseExerciseName_MarchingHipRaise,
	"WeightedMarchingHipRaise":                        HipRaiseExerciseName_WeightedMarchingHipRaise,
	"MarchingHipRaiseWithFeetOnASwissBall":            HipRaiseExerciseName_MarchingHipRaiseWithFeetOnASwissBall,
	"WeightedMarchingHipRaiseWithFeetOnASwissBall":    HipRaiseExerciseName_WeightedMarchingHipRaiseWithFeetOnASwissBall,
	"ReverseHipRaise":                                 HipRaiseExerciseName_ReverseHipRaise,
	"WeightedReverseHipRaise":                         HipRaiseExerciseName_WeightedReverseHipRaise,
	"SingleLegHipRaise":                               HipRaiseExerciseName_SingleLegHipRaise,
	"WeightedSingleLegHipRaise":                       HipRaiseExerciseName_WeightedSingleLegHipRaise,
	"SingleLegHipRaiseWithFootOnBench":                HipRaiseExerciseName_SingleLegHipRaiseWithFootOnBench,
	"WeightedSingleLegHipRaiseWithFootOnBench":        HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnBench,
	"SingleLegHipRaiseWithFootOnBosuBall":             HipRaiseExerciseName_SingleLegHipRaiseWithFootOnBosuBall,
	"WeightedSingleLegHipRaiseWithFootOnBosuBall":     HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnBosuBall,
	"SingleLegHipRaiseWithFootOnFoamRoller":           HipRaiseExerciseName_SingleLegHipRaiseWithFootOnFoamRoller,
	"WeightedSingleLegHipRaiseWithFootOnFoamRoller":   HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnFoamRoller,
	"SingleLegHipRaiseWithFootOnMedicineBall":         HipRaiseExerciseName_SingleLegHipRaiseWithFootOnMedicineBall,
	"WeightedSingleLegHipRaiseWithFootOnMedicineBall": HipRaiseExerciseName_WeightedSingleLegHipRaiseWithFootOnMedicineBall,
	"SingleLegHipRaiseWithHeadOnBosuBall":             HipRaiseExerciseName_SingleLegHipRaiseWithHeadOnBosuBall,
	"WeightedSingleLegHipRaiseWithHeadOnBosuBall":     HipRaiseExerciseName_WeightedSingleLegHipRaiseWithHeadOnBosuBall,
	"WeightedClamBridge":                              HipRaiseExerciseName_WeightedClamBridge,
	"SingleLegSwissBallHipRaiseAndLegCurl":            HipRaiseExerciseName_SingleLegSwissBallHipRaiseAndLegCurl,
	"Clams":                                           HipRaiseExerciseName_Clams,
	"LegCircles":                                      HipRaiseExerciseName_LegCircles,
	"LegLift":                                         HipRaiseExerciseName_LegLift,
	"LegLiftInExternalRotation":                       HipRaiseExerciseName_LegLiftInExternalRotation,
}

func HipRaiseExerciseNameFromString(s string) HipRaiseExerciseName {
	if v, ok := HipRaiseExerciseNameValues[s]; ok {
		return v
	}
	return HipRaiseExerciseName(HipRaiseExerciseName_Invalid)
}
