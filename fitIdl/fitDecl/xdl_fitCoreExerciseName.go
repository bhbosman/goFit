package fitDecl

import strconv "strconv"

type CoreExerciseName uint16

const (
	CoreExerciseName_AbsJabs                          CoreExerciseName = 0
	CoreExerciseName_WeightedAbsJabs                  CoreExerciseName = 1
	CoreExerciseName_AlternatingPlateReach            CoreExerciseName = 2
	CoreExerciseName_BarbellRollout                   CoreExerciseName = 3
	CoreExerciseName_WeightedBarbellRollout           CoreExerciseName = 4
	CoreExerciseName_BodyBarObliqueTwist              CoreExerciseName = 5
	CoreExerciseName_CableCorePress                   CoreExerciseName = 6
	CoreExerciseName_CableSideBend                    CoreExerciseName = 7
	CoreExerciseName_SideBend                         CoreExerciseName = 8
	CoreExerciseName_WeightedSideBend                 CoreExerciseName = 9
	CoreExerciseName_CrescentCircle                   CoreExerciseName = 10
	CoreExerciseName_WeightedCrescentCircle           CoreExerciseName = 11
	CoreExerciseName_CyclingRussianTwist              CoreExerciseName = 12
	CoreExerciseName_WeightedCyclingRussianTwist      CoreExerciseName = 13
	CoreExerciseName_ElevatedFeetRussianTwist         CoreExerciseName = 14
	CoreExerciseName_WeightedElevatedFeetRussianTwist CoreExerciseName = 15
	CoreExerciseName_HalfTurkishGetUp                 CoreExerciseName = 16
	CoreExerciseName_KettlebellWindmill               CoreExerciseName = 17
	CoreExerciseName_KneelingAbWheel                  CoreExerciseName = 18
	CoreExerciseName_WeightedKneelingAbWheel          CoreExerciseName = 19
	CoreExerciseName_ModifiedFrontLever               CoreExerciseName = 20
	CoreExerciseName_OpenKneeTucks                    CoreExerciseName = 21
	CoreExerciseName_WeightedOpenKneeTucks            CoreExerciseName = 22
	CoreExerciseName_SideAbsLegLift                   CoreExerciseName = 23
	CoreExerciseName_WeightedSideAbsLegLift           CoreExerciseName = 24
	CoreExerciseName_SwissBallJackknife               CoreExerciseName = 25
	CoreExerciseName_WeightedSwissBallJackknife       CoreExerciseName = 26
	CoreExerciseName_SwissBallPike                    CoreExerciseName = 27
	CoreExerciseName_WeightedSwissBallPike            CoreExerciseName = 28
	CoreExerciseName_SwissBallRollout                 CoreExerciseName = 29
	CoreExerciseName_WeightedSwissBallRollout         CoreExerciseName = 30
	CoreExerciseName_TriangleHipPress                 CoreExerciseName = 31
	CoreExerciseName_WeightedTriangleHipPress         CoreExerciseName = 32
	CoreExerciseName_TrxSuspendedJackknife            CoreExerciseName = 33
	CoreExerciseName_WeightedTrxSuspendedJackknife    CoreExerciseName = 34
	CoreExerciseName_UBoat                            CoreExerciseName = 35
	CoreExerciseName_WeightedUBoat                    CoreExerciseName = 36
	CoreExerciseName_WindmillSwitches                 CoreExerciseName = 37
	CoreExerciseName_WeightedWindmillSwitches         CoreExerciseName = 38
	CoreExerciseName_AlternatingSlideOut              CoreExerciseName = 39
	CoreExerciseName_WeightedAlternatingSlideOut      CoreExerciseName = 40
	CoreExerciseName_GhdBackExtensions                CoreExerciseName = 41
	CoreExerciseName_WeightedGhdBackExtensions        CoreExerciseName = 42
	CoreExerciseName_OverheadWalk                     CoreExerciseName = 43
	CoreExerciseName_Inchworm                         CoreExerciseName = 44
	CoreExerciseName_WeightedModifiedFrontLever       CoreExerciseName = 45
	CoreExerciseName_RussianTwist                     CoreExerciseName = 46
	CoreExerciseName_ArmAndLegExtensionOnKnees        CoreExerciseName = 48
	CoreExerciseName_Bicycle                          CoreExerciseName = 49
	CoreExerciseName_BicepCurlWithLegExtension        CoreExerciseName = 50
	CoreExerciseName_CatCow                           CoreExerciseName = 51
	CoreExerciseName_Corkscrew                        CoreExerciseName = 52
	CoreExerciseName_CrissCross                       CoreExerciseName = 53
	CoreExerciseName_DoubleLegStretch                 CoreExerciseName = 55
	CoreExerciseName_KneeFolds                        CoreExerciseName = 56
	CoreExerciseName_LowerLift                        CoreExerciseName = 57
	CoreExerciseName_NeckPull                         CoreExerciseName = 58
	CoreExerciseName_PelvicClocks                     CoreExerciseName = 59
	CoreExerciseName_RollOver                         CoreExerciseName = 60
	CoreExerciseName_RollUp                           CoreExerciseName = 61
	CoreExerciseName_Rolling                          CoreExerciseName = 62
	CoreExerciseName_Rowing1                          CoreExerciseName = 63
	CoreExerciseName_Rowing2                          CoreExerciseName = 64
	CoreExerciseName_Scissors                         CoreExerciseName = 65
	CoreExerciseName_SingleLegCircles                 CoreExerciseName = 66
	CoreExerciseName_SingleLegStretch                 CoreExerciseName = 67
	CoreExerciseName_Swan                             CoreExerciseName = 69
	CoreExerciseName_Swimming                         CoreExerciseName = 70
	CoreExerciseName_Teaser                           CoreExerciseName = 71
	CoreExerciseName_TheHundred                       CoreExerciseName = 72
	CoreExerciseName_Invalid                          CoreExerciseName = 65535
)

var CoreExerciseNameNames = map[CoreExerciseName]string{
	CoreExerciseName_AbsJabs:                          "AbsJabs",
	CoreExerciseName_WeightedAbsJabs:                  "WeightedAbsJabs",
	CoreExerciseName_AlternatingPlateReach:            "AlternatingPlateReach",
	CoreExerciseName_BarbellRollout:                   "BarbellRollout",
	CoreExerciseName_WeightedBarbellRollout:           "WeightedBarbellRollout",
	CoreExerciseName_BodyBarObliqueTwist:              "BodyBarObliqueTwist",
	CoreExerciseName_CableCorePress:                   "CableCorePress",
	CoreExerciseName_CableSideBend:                    "CableSideBend",
	CoreExerciseName_SideBend:                         "SideBend",
	CoreExerciseName_WeightedSideBend:                 "WeightedSideBend",
	CoreExerciseName_CrescentCircle:                   "CrescentCircle",
	CoreExerciseName_WeightedCrescentCircle:           "WeightedCrescentCircle",
	CoreExerciseName_CyclingRussianTwist:              "CyclingRussianTwist",
	CoreExerciseName_WeightedCyclingRussianTwist:      "WeightedCyclingRussianTwist",
	CoreExerciseName_ElevatedFeetRussianTwist:         "ElevatedFeetRussianTwist",
	CoreExerciseName_WeightedElevatedFeetRussianTwist: "WeightedElevatedFeetRussianTwist",
	CoreExerciseName_HalfTurkishGetUp:                 "HalfTurkishGetUp",
	CoreExerciseName_KettlebellWindmill:               "KettlebellWindmill",
	CoreExerciseName_KneelingAbWheel:                  "KneelingAbWheel",
	CoreExerciseName_WeightedKneelingAbWheel:          "WeightedKneelingAbWheel",
	CoreExerciseName_ModifiedFrontLever:               "ModifiedFrontLever",
	CoreExerciseName_OpenKneeTucks:                    "OpenKneeTucks",
	CoreExerciseName_WeightedOpenKneeTucks:            "WeightedOpenKneeTucks",
	CoreExerciseName_SideAbsLegLift:                   "SideAbsLegLift",
	CoreExerciseName_WeightedSideAbsLegLift:           "WeightedSideAbsLegLift",
	CoreExerciseName_SwissBallJackknife:               "SwissBallJackknife",
	CoreExerciseName_WeightedSwissBallJackknife:       "WeightedSwissBallJackknife",
	CoreExerciseName_SwissBallPike:                    "SwissBallPike",
	CoreExerciseName_WeightedSwissBallPike:            "WeightedSwissBallPike",
	CoreExerciseName_SwissBallRollout:                 "SwissBallRollout",
	CoreExerciseName_WeightedSwissBallRollout:         "WeightedSwissBallRollout",
	CoreExerciseName_TriangleHipPress:                 "TriangleHipPress",
	CoreExerciseName_WeightedTriangleHipPress:         "WeightedTriangleHipPress",
	CoreExerciseName_TrxSuspendedJackknife:            "TrxSuspendedJackknife",
	CoreExerciseName_WeightedTrxSuspendedJackknife:    "WeightedTrxSuspendedJackknife",
	CoreExerciseName_UBoat:                            "UBoat",
	CoreExerciseName_WeightedUBoat:                    "WeightedUBoat",
	CoreExerciseName_WindmillSwitches:                 "WindmillSwitches",
	CoreExerciseName_WeightedWindmillSwitches:         "WeightedWindmillSwitches",
	CoreExerciseName_AlternatingSlideOut:              "AlternatingSlideOut",
	CoreExerciseName_WeightedAlternatingSlideOut:      "WeightedAlternatingSlideOut",
	CoreExerciseName_GhdBackExtensions:                "GhdBackExtensions",
	CoreExerciseName_WeightedGhdBackExtensions:        "WeightedGhdBackExtensions",
	CoreExerciseName_OverheadWalk:                     "OverheadWalk",
	CoreExerciseName_Inchworm:                         "Inchworm",
	CoreExerciseName_WeightedModifiedFrontLever:       "WeightedModifiedFrontLever",
	CoreExerciseName_RussianTwist:                     "RussianTwist",
	CoreExerciseName_ArmAndLegExtensionOnKnees:        "ArmAndLegExtensionOnKnees",
	CoreExerciseName_Bicycle:                          "Bicycle",
	CoreExerciseName_BicepCurlWithLegExtension:        "BicepCurlWithLegExtension",
	CoreExerciseName_CatCow:                           "CatCow",
	CoreExerciseName_Corkscrew:                        "Corkscrew",
	CoreExerciseName_CrissCross:                       "CrissCross",
	CoreExerciseName_DoubleLegStretch:                 "DoubleLegStretch",
	CoreExerciseName_KneeFolds:                        "KneeFolds",
	CoreExerciseName_LowerLift:                        "LowerLift",
	CoreExerciseName_NeckPull:                         "NeckPull",
	CoreExerciseName_PelvicClocks:                     "PelvicClocks",
	CoreExerciseName_RollOver:                         "RollOver",
	CoreExerciseName_RollUp:                           "RollUp",
	CoreExerciseName_Rolling:                          "Rolling",
	CoreExerciseName_Rowing1:                          "Rowing1",
	CoreExerciseName_Rowing2:                          "Rowing2",
	CoreExerciseName_Scissors:                         "Scissors",
	CoreExerciseName_SingleLegCircles:                 "SingleLegCircles",
	CoreExerciseName_SingleLegStretch:                 "SingleLegStretch",
	CoreExerciseName_Swan:                             "Swan",
	CoreExerciseName_Swimming:                         "Swimming",
	CoreExerciseName_Teaser:                           "Teaser",
	CoreExerciseName_TheHundred:                       "TheHundred",
}

func (k CoreExerciseName) String() string {
	if uint(k) < uint(len(CoreExerciseNameNames)) {
		if v, ok := CoreExerciseNameNames[k]; ok {
			return v
		}
	}
	return "CoreExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var CoreExerciseNameValues = map[string]CoreExerciseName{
	"AbsJabs":                          CoreExerciseName_AbsJabs,
	"WeightedAbsJabs":                  CoreExerciseName_WeightedAbsJabs,
	"AlternatingPlateReach":            CoreExerciseName_AlternatingPlateReach,
	"BarbellRollout":                   CoreExerciseName_BarbellRollout,
	"WeightedBarbellRollout":           CoreExerciseName_WeightedBarbellRollout,
	"BodyBarObliqueTwist":              CoreExerciseName_BodyBarObliqueTwist,
	"CableCorePress":                   CoreExerciseName_CableCorePress,
	"CableSideBend":                    CoreExerciseName_CableSideBend,
	"SideBend":                         CoreExerciseName_SideBend,
	"WeightedSideBend":                 CoreExerciseName_WeightedSideBend,
	"CrescentCircle":                   CoreExerciseName_CrescentCircle,
	"WeightedCrescentCircle":           CoreExerciseName_WeightedCrescentCircle,
	"CyclingRussianTwist":              CoreExerciseName_CyclingRussianTwist,
	"WeightedCyclingRussianTwist":      CoreExerciseName_WeightedCyclingRussianTwist,
	"ElevatedFeetRussianTwist":         CoreExerciseName_ElevatedFeetRussianTwist,
	"WeightedElevatedFeetRussianTwist": CoreExerciseName_WeightedElevatedFeetRussianTwist,
	"HalfTurkishGetUp":                 CoreExerciseName_HalfTurkishGetUp,
	"KettlebellWindmill":               CoreExerciseName_KettlebellWindmill,
	"KneelingAbWheel":                  CoreExerciseName_KneelingAbWheel,
	"WeightedKneelingAbWheel":          CoreExerciseName_WeightedKneelingAbWheel,
	"ModifiedFrontLever":               CoreExerciseName_ModifiedFrontLever,
	"OpenKneeTucks":                    CoreExerciseName_OpenKneeTucks,
	"WeightedOpenKneeTucks":            CoreExerciseName_WeightedOpenKneeTucks,
	"SideAbsLegLift":                   CoreExerciseName_SideAbsLegLift,
	"WeightedSideAbsLegLift":           CoreExerciseName_WeightedSideAbsLegLift,
	"SwissBallJackknife":               CoreExerciseName_SwissBallJackknife,
	"WeightedSwissBallJackknife":       CoreExerciseName_WeightedSwissBallJackknife,
	"SwissBallPike":                    CoreExerciseName_SwissBallPike,
	"WeightedSwissBallPike":            CoreExerciseName_WeightedSwissBallPike,
	"SwissBallRollout":                 CoreExerciseName_SwissBallRollout,
	"WeightedSwissBallRollout":         CoreExerciseName_WeightedSwissBallRollout,
	"TriangleHipPress":                 CoreExerciseName_TriangleHipPress,
	"WeightedTriangleHipPress":         CoreExerciseName_WeightedTriangleHipPress,
	"TrxSuspendedJackknife":            CoreExerciseName_TrxSuspendedJackknife,
	"WeightedTrxSuspendedJackknife":    CoreExerciseName_WeightedTrxSuspendedJackknife,
	"UBoat":                            CoreExerciseName_UBoat,
	"WeightedUBoat":                    CoreExerciseName_WeightedUBoat,
	"WindmillSwitches":                 CoreExerciseName_WindmillSwitches,
	"WeightedWindmillSwitches":         CoreExerciseName_WeightedWindmillSwitches,
	"AlternatingSlideOut":              CoreExerciseName_AlternatingSlideOut,
	"WeightedAlternatingSlideOut":      CoreExerciseName_WeightedAlternatingSlideOut,
	"GhdBackExtensions":                CoreExerciseName_GhdBackExtensions,
	"WeightedGhdBackExtensions":        CoreExerciseName_WeightedGhdBackExtensions,
	"OverheadWalk":                     CoreExerciseName_OverheadWalk,
	"Inchworm":                         CoreExerciseName_Inchworm,
	"WeightedModifiedFrontLever":       CoreExerciseName_WeightedModifiedFrontLever,
	"RussianTwist":                     CoreExerciseName_RussianTwist,
	"ArmAndLegExtensionOnKnees":        CoreExerciseName_ArmAndLegExtensionOnKnees,
	"Bicycle":                          CoreExerciseName_Bicycle,
	"BicepCurlWithLegExtension":        CoreExerciseName_BicepCurlWithLegExtension,
	"CatCow":                           CoreExerciseName_CatCow,
	"Corkscrew":                        CoreExerciseName_Corkscrew,
	"CrissCross":                       CoreExerciseName_CrissCross,
	"DoubleLegStretch":                 CoreExerciseName_DoubleLegStretch,
	"KneeFolds":                        CoreExerciseName_KneeFolds,
	"LowerLift":                        CoreExerciseName_LowerLift,
	"NeckPull":                         CoreExerciseName_NeckPull,
	"PelvicClocks":                     CoreExerciseName_PelvicClocks,
	"RollOver":                         CoreExerciseName_RollOver,
	"RollUp":                           CoreExerciseName_RollUp,
	"Rolling":                          CoreExerciseName_Rolling,
	"Rowing1":                          CoreExerciseName_Rowing1,
	"Rowing2":                          CoreExerciseName_Rowing2,
	"Scissors":                         CoreExerciseName_Scissors,
	"SingleLegCircles":                 CoreExerciseName_SingleLegCircles,
	"SingleLegStretch":                 CoreExerciseName_SingleLegStretch,
	"Swan":                             CoreExerciseName_Swan,
	"Swimming":                         CoreExerciseName_Swimming,
	"Teaser":                           CoreExerciseName_Teaser,
	"TheHundred":                       CoreExerciseName_TheHundred,
}

func CoreExerciseNameFromString(s string) CoreExerciseName {
	if v, ok := CoreExerciseNameValues[s]; ok {
		return v
	}
	return CoreExerciseName(CoreExerciseName_Invalid)
}
