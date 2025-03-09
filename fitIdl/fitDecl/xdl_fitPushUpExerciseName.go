package fitDecl

import strconv "strconv"

type PushUpExerciseName uint16

const (
	PushUpExerciseName_ChestPressWithBand                         PushUpExerciseName = 0
	PushUpExerciseName_AlternatingStaggeredPushUp                 PushUpExerciseName = 1
	PushUpExerciseName_WeightedAlternatingStaggeredPushUp         PushUpExerciseName = 2
	PushUpExerciseName_AlternatingHandsMedicineBallPushUp         PushUpExerciseName = 3
	PushUpExerciseName_WeightedAlternatingHandsMedicineBallPushUp PushUpExerciseName = 4
	PushUpExerciseName_BosuBallPushUp                             PushUpExerciseName = 5
	PushUpExerciseName_WeightedBosuBallPushUp                     PushUpExerciseName = 6
	PushUpExerciseName_ClappingPushUp                             PushUpExerciseName = 7
	PushUpExerciseName_WeightedClappingPushUp                     PushUpExerciseName = 8
	PushUpExerciseName_CloseGripMedicineBallPushUp                PushUpExerciseName = 9
	PushUpExerciseName_WeightedCloseGripMedicineBallPushUp        PushUpExerciseName = 10
	PushUpExerciseName_CloseHandsPushUp                           PushUpExerciseName = 11
	PushUpExerciseName_WeightedCloseHandsPushUp                   PushUpExerciseName = 12
	PushUpExerciseName_DeclinePushUp                              PushUpExerciseName = 13
	PushUpExerciseName_WeightedDeclinePushUp                      PushUpExerciseName = 14
	PushUpExerciseName_DiamondPushUp                              PushUpExerciseName = 15
	PushUpExerciseName_WeightedDiamondPushUp                      PushUpExerciseName = 16
	PushUpExerciseName_ExplosiveCrossoverPushUp                   PushUpExerciseName = 17
	PushUpExerciseName_WeightedExplosiveCrossoverPushUp           PushUpExerciseName = 18
	PushUpExerciseName_ExplosivePushUp                            PushUpExerciseName = 19
	PushUpExerciseName_WeightedExplosivePushUp                    PushUpExerciseName = 20
	PushUpExerciseName_FeetElevatedSideToSidePushUp               PushUpExerciseName = 21
	PushUpExerciseName_WeightedFeetElevatedSideToSidePushUp       PushUpExerciseName = 22
	PushUpExerciseName_HandReleasePushUp                          PushUpExerciseName = 23
	PushUpExerciseName_WeightedHandReleasePushUp                  PushUpExerciseName = 24
	PushUpExerciseName_HandstandPushUp                            PushUpExerciseName = 25
	PushUpExerciseName_WeightedHandstandPushUp                    PushUpExerciseName = 26
	PushUpExerciseName_InclinePushUp                              PushUpExerciseName = 27
	PushUpExerciseName_WeightedInclinePushUp                      PushUpExerciseName = 28
	PushUpExerciseName_IsometricExplosivePushUp                   PushUpExerciseName = 29
	PushUpExerciseName_WeightedIsometricExplosivePushUp           PushUpExerciseName = 30
	PushUpExerciseName_JudoPushUp                                 PushUpExerciseName = 31
	PushUpExerciseName_WeightedJudoPushUp                         PushUpExerciseName = 32
	PushUpExerciseName_KneelingPushUp                             PushUpExerciseName = 33
	PushUpExerciseName_WeightedKneelingPushUp                     PushUpExerciseName = 34
	PushUpExerciseName_MedicineBallChestPass                      PushUpExerciseName = 35
	PushUpExerciseName_MedicineBallPushUp                         PushUpExerciseName = 36
	PushUpExerciseName_WeightedMedicineBallPushUp                 PushUpExerciseName = 37
	PushUpExerciseName_OneArmPushUp                               PushUpExerciseName = 38
	PushUpExerciseName_WeightedOneArmPushUp                       PushUpExerciseName = 39
	PushUpExerciseName_WeightedPushUp                             PushUpExerciseName = 40
	PushUpExerciseName_PushUpAndRow                               PushUpExerciseName = 41
	PushUpExerciseName_WeightedPushUpAndRow                       PushUpExerciseName = 42
	PushUpExerciseName_PushUpPlus                                 PushUpExerciseName = 43
	PushUpExerciseName_WeightedPushUpPlus                         PushUpExerciseName = 44
	PushUpExerciseName_PushUpWithFeetOnSwissBall                  PushUpExerciseName = 45
	PushUpExerciseName_WeightedPushUpWithFeetOnSwissBall          PushUpExerciseName = 46
	PushUpExerciseName_PushUpWithOneHandOnMedicineBall            PushUpExerciseName = 47
	PushUpExerciseName_WeightedPushUpWithOneHandOnMedicineBall    PushUpExerciseName = 48
	PushUpExerciseName_ShoulderPushUp                             PushUpExerciseName = 49
	PushUpExerciseName_WeightedShoulderPushUp                     PushUpExerciseName = 50
	PushUpExerciseName_SingleArmMedicineBallPushUp                PushUpExerciseName = 51
	PushUpExerciseName_WeightedSingleArmMedicineBallPushUp        PushUpExerciseName = 52
	PushUpExerciseName_SpidermanPushUp                            PushUpExerciseName = 53
	PushUpExerciseName_WeightedSpidermanPushUp                    PushUpExerciseName = 54
	PushUpExerciseName_StackedFeetPushUp                          PushUpExerciseName = 55
	PushUpExerciseName_WeightedStackedFeetPushUp                  PushUpExerciseName = 56
	PushUpExerciseName_StaggeredHandsPushUp                       PushUpExerciseName = 57
	PushUpExerciseName_WeightedStaggeredHandsPushUp               PushUpExerciseName = 58
	PushUpExerciseName_SuspendedPushUp                            PushUpExerciseName = 59
	PushUpExerciseName_WeightedSuspendedPushUp                    PushUpExerciseName = 60
	PushUpExerciseName_SwissBallPushUp                            PushUpExerciseName = 61
	PushUpExerciseName_WeightedSwissBallPushUp                    PushUpExerciseName = 62
	PushUpExerciseName_SwissBallPushUpPlus                        PushUpExerciseName = 63
	PushUpExerciseName_WeightedSwissBallPushUpPlus                PushUpExerciseName = 64
	PushUpExerciseName_TPushUp                                    PushUpExerciseName = 65
	PushUpExerciseName_WeightedTPushUp                            PushUpExerciseName = 66
	PushUpExerciseName_TripleStopPushUp                           PushUpExerciseName = 67
	PushUpExerciseName_WeightedTripleStopPushUp                   PushUpExerciseName = 68
	PushUpExerciseName_WideHandsPushUp                            PushUpExerciseName = 69
	PushUpExerciseName_WeightedWideHandsPushUp                    PushUpExerciseName = 70
	PushUpExerciseName_ParalletteHandstandPushUp                  PushUpExerciseName = 71
	PushUpExerciseName_WeightedParalletteHandstandPushUp          PushUpExerciseName = 72
	PushUpExerciseName_RingHandstandPushUp                        PushUpExerciseName = 73
	PushUpExerciseName_WeightedRingHandstandPushUp                PushUpExerciseName = 74
	PushUpExerciseName_RingPushUp                                 PushUpExerciseName = 75
	PushUpExerciseName_WeightedRingPushUp                         PushUpExerciseName = 76
	PushUpExerciseName_PushUp                                     PushUpExerciseName = 77
	PushUpExerciseName_PilatesPushup                              PushUpExerciseName = 78
	PushUpExerciseName_Invalid                                    PushUpExerciseName = 65535
)

var PushUpExerciseNameNames = map[PushUpExerciseName]string{
	PushUpExerciseName_ChestPressWithBand:                         "ChestPressWithBand",
	PushUpExerciseName_AlternatingStaggeredPushUp:                 "AlternatingStaggeredPushUp",
	PushUpExerciseName_WeightedAlternatingStaggeredPushUp:         "WeightedAlternatingStaggeredPushUp",
	PushUpExerciseName_AlternatingHandsMedicineBallPushUp:         "AlternatingHandsMedicineBallPushUp",
	PushUpExerciseName_WeightedAlternatingHandsMedicineBallPushUp: "WeightedAlternatingHandsMedicineBallPushUp",
	PushUpExerciseName_BosuBallPushUp:                             "BosuBallPushUp",
	PushUpExerciseName_WeightedBosuBallPushUp:                     "WeightedBosuBallPushUp",
	PushUpExerciseName_ClappingPushUp:                             "ClappingPushUp",
	PushUpExerciseName_WeightedClappingPushUp:                     "WeightedClappingPushUp",
	PushUpExerciseName_CloseGripMedicineBallPushUp:                "CloseGripMedicineBallPushUp",
	PushUpExerciseName_WeightedCloseGripMedicineBallPushUp:        "WeightedCloseGripMedicineBallPushUp",
	PushUpExerciseName_CloseHandsPushUp:                           "CloseHandsPushUp",
	PushUpExerciseName_WeightedCloseHandsPushUp:                   "WeightedCloseHandsPushUp",
	PushUpExerciseName_DeclinePushUp:                              "DeclinePushUp",
	PushUpExerciseName_WeightedDeclinePushUp:                      "WeightedDeclinePushUp",
	PushUpExerciseName_DiamondPushUp:                              "DiamondPushUp",
	PushUpExerciseName_WeightedDiamondPushUp:                      "WeightedDiamondPushUp",
	PushUpExerciseName_ExplosiveCrossoverPushUp:                   "ExplosiveCrossoverPushUp",
	PushUpExerciseName_WeightedExplosiveCrossoverPushUp:           "WeightedExplosiveCrossoverPushUp",
	PushUpExerciseName_ExplosivePushUp:                            "ExplosivePushUp",
	PushUpExerciseName_WeightedExplosivePushUp:                    "WeightedExplosivePushUp",
	PushUpExerciseName_FeetElevatedSideToSidePushUp:               "FeetElevatedSideToSidePushUp",
	PushUpExerciseName_WeightedFeetElevatedSideToSidePushUp:       "WeightedFeetElevatedSideToSidePushUp",
	PushUpExerciseName_HandReleasePushUp:                          "HandReleasePushUp",
	PushUpExerciseName_WeightedHandReleasePushUp:                  "WeightedHandReleasePushUp",
	PushUpExerciseName_HandstandPushUp:                            "HandstandPushUp",
	PushUpExerciseName_WeightedHandstandPushUp:                    "WeightedHandstandPushUp",
	PushUpExerciseName_InclinePushUp:                              "InclinePushUp",
	PushUpExerciseName_WeightedInclinePushUp:                      "WeightedInclinePushUp",
	PushUpExerciseName_IsometricExplosivePushUp:                   "IsometricExplosivePushUp",
	PushUpExerciseName_WeightedIsometricExplosivePushUp:           "WeightedIsometricExplosivePushUp",
	PushUpExerciseName_JudoPushUp:                                 "JudoPushUp",
	PushUpExerciseName_WeightedJudoPushUp:                         "WeightedJudoPushUp",
	PushUpExerciseName_KneelingPushUp:                             "KneelingPushUp",
	PushUpExerciseName_WeightedKneelingPushUp:                     "WeightedKneelingPushUp",
	PushUpExerciseName_MedicineBallChestPass:                      "MedicineBallChestPass",
	PushUpExerciseName_MedicineBallPushUp:                         "MedicineBallPushUp",
	PushUpExerciseName_WeightedMedicineBallPushUp:                 "WeightedMedicineBallPushUp",
	PushUpExerciseName_OneArmPushUp:                               "OneArmPushUp",
	PushUpExerciseName_WeightedOneArmPushUp:                       "WeightedOneArmPushUp",
	PushUpExerciseName_WeightedPushUp:                             "WeightedPushUp",
	PushUpExerciseName_PushUpAndRow:                               "PushUpAndRow",
	PushUpExerciseName_WeightedPushUpAndRow:                       "WeightedPushUpAndRow",
	PushUpExerciseName_PushUpPlus:                                 "PushUpPlus",
	PushUpExerciseName_WeightedPushUpPlus:                         "WeightedPushUpPlus",
	PushUpExerciseName_PushUpWithFeetOnSwissBall:                  "PushUpWithFeetOnSwissBall",
	PushUpExerciseName_WeightedPushUpWithFeetOnSwissBall:          "WeightedPushUpWithFeetOnSwissBall",
	PushUpExerciseName_PushUpWithOneHandOnMedicineBall:            "PushUpWithOneHandOnMedicineBall",
	PushUpExerciseName_WeightedPushUpWithOneHandOnMedicineBall:    "WeightedPushUpWithOneHandOnMedicineBall",
	PushUpExerciseName_ShoulderPushUp:                             "ShoulderPushUp",
	PushUpExerciseName_WeightedShoulderPushUp:                     "WeightedShoulderPushUp",
	PushUpExerciseName_SingleArmMedicineBallPushUp:                "SingleArmMedicineBallPushUp",
	PushUpExerciseName_WeightedSingleArmMedicineBallPushUp:        "WeightedSingleArmMedicineBallPushUp",
	PushUpExerciseName_SpidermanPushUp:                            "SpidermanPushUp",
	PushUpExerciseName_WeightedSpidermanPushUp:                    "WeightedSpidermanPushUp",
	PushUpExerciseName_StackedFeetPushUp:                          "StackedFeetPushUp",
	PushUpExerciseName_WeightedStackedFeetPushUp:                  "WeightedStackedFeetPushUp",
	PushUpExerciseName_StaggeredHandsPushUp:                       "StaggeredHandsPushUp",
	PushUpExerciseName_WeightedStaggeredHandsPushUp:               "WeightedStaggeredHandsPushUp",
	PushUpExerciseName_SuspendedPushUp:                            "SuspendedPushUp",
	PushUpExerciseName_WeightedSuspendedPushUp:                    "WeightedSuspendedPushUp",
	PushUpExerciseName_SwissBallPushUp:                            "SwissBallPushUp",
	PushUpExerciseName_WeightedSwissBallPushUp:                    "WeightedSwissBallPushUp",
	PushUpExerciseName_SwissBallPushUpPlus:                        "SwissBallPushUpPlus",
	PushUpExerciseName_WeightedSwissBallPushUpPlus:                "WeightedSwissBallPushUpPlus",
	PushUpExerciseName_TPushUp:                                    "TPushUp",
	PushUpExerciseName_WeightedTPushUp:                            "WeightedTPushUp",
	PushUpExerciseName_TripleStopPushUp:                           "TripleStopPushUp",
	PushUpExerciseName_WeightedTripleStopPushUp:                   "WeightedTripleStopPushUp",
	PushUpExerciseName_WideHandsPushUp:                            "WideHandsPushUp",
	PushUpExerciseName_WeightedWideHandsPushUp:                    "WeightedWideHandsPushUp",
	PushUpExerciseName_ParalletteHandstandPushUp:                  "ParalletteHandstandPushUp",
	PushUpExerciseName_WeightedParalletteHandstandPushUp:          "WeightedParalletteHandstandPushUp",
	PushUpExerciseName_RingHandstandPushUp:                        "RingHandstandPushUp",
	PushUpExerciseName_WeightedRingHandstandPushUp:                "WeightedRingHandstandPushUp",
	PushUpExerciseName_RingPushUp:                                 "RingPushUp",
	PushUpExerciseName_WeightedRingPushUp:                         "WeightedRingPushUp",
	PushUpExerciseName_PushUp:                                     "PushUp",
	PushUpExerciseName_PilatesPushup:                              "PilatesPushup",
}

func (k PushUpExerciseName) String() string {
	if uint(k) < uint(len(PushUpExerciseNameNames)) {
		if v, ok := PushUpExerciseNameNames[k]; ok {
			return v
		}
	}
	return "PushUpExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var PushUpExerciseNameValues = map[string]PushUpExerciseName{
	"ChestPressWithBand":                         PushUpExerciseName_ChestPressWithBand,
	"AlternatingStaggeredPushUp":                 PushUpExerciseName_AlternatingStaggeredPushUp,
	"WeightedAlternatingStaggeredPushUp":         PushUpExerciseName_WeightedAlternatingStaggeredPushUp,
	"AlternatingHandsMedicineBallPushUp":         PushUpExerciseName_AlternatingHandsMedicineBallPushUp,
	"WeightedAlternatingHandsMedicineBallPushUp": PushUpExerciseName_WeightedAlternatingHandsMedicineBallPushUp,
	"BosuBallPushUp":                             PushUpExerciseName_BosuBallPushUp,
	"WeightedBosuBallPushUp":                     PushUpExerciseName_WeightedBosuBallPushUp,
	"ClappingPushUp":                             PushUpExerciseName_ClappingPushUp,
	"WeightedClappingPushUp":                     PushUpExerciseName_WeightedClappingPushUp,
	"CloseGripMedicineBallPushUp":                PushUpExerciseName_CloseGripMedicineBallPushUp,
	"WeightedCloseGripMedicineBallPushUp":        PushUpExerciseName_WeightedCloseGripMedicineBallPushUp,
	"CloseHandsPushUp":                           PushUpExerciseName_CloseHandsPushUp,
	"WeightedCloseHandsPushUp":                   PushUpExerciseName_WeightedCloseHandsPushUp,
	"DeclinePushUp":                              PushUpExerciseName_DeclinePushUp,
	"WeightedDeclinePushUp":                      PushUpExerciseName_WeightedDeclinePushUp,
	"DiamondPushUp":                              PushUpExerciseName_DiamondPushUp,
	"WeightedDiamondPushUp":                      PushUpExerciseName_WeightedDiamondPushUp,
	"ExplosiveCrossoverPushUp":                   PushUpExerciseName_ExplosiveCrossoverPushUp,
	"WeightedExplosiveCrossoverPushUp":           PushUpExerciseName_WeightedExplosiveCrossoverPushUp,
	"ExplosivePushUp":                            PushUpExerciseName_ExplosivePushUp,
	"WeightedExplosivePushUp":                    PushUpExerciseName_WeightedExplosivePushUp,
	"FeetElevatedSideToSidePushUp":               PushUpExerciseName_FeetElevatedSideToSidePushUp,
	"WeightedFeetElevatedSideToSidePushUp":       PushUpExerciseName_WeightedFeetElevatedSideToSidePushUp,
	"HandReleasePushUp":                          PushUpExerciseName_HandReleasePushUp,
	"WeightedHandReleasePushUp":                  PushUpExerciseName_WeightedHandReleasePushUp,
	"HandstandPushUp":                            PushUpExerciseName_HandstandPushUp,
	"WeightedHandstandPushUp":                    PushUpExerciseName_WeightedHandstandPushUp,
	"InclinePushUp":                              PushUpExerciseName_InclinePushUp,
	"WeightedInclinePushUp":                      PushUpExerciseName_WeightedInclinePushUp,
	"IsometricExplosivePushUp":                   PushUpExerciseName_IsometricExplosivePushUp,
	"WeightedIsometricExplosivePushUp":           PushUpExerciseName_WeightedIsometricExplosivePushUp,
	"JudoPushUp":                                 PushUpExerciseName_JudoPushUp,
	"WeightedJudoPushUp":                         PushUpExerciseName_WeightedJudoPushUp,
	"KneelingPushUp":                             PushUpExerciseName_KneelingPushUp,
	"WeightedKneelingPushUp":                     PushUpExerciseName_WeightedKneelingPushUp,
	"MedicineBallChestPass":                      PushUpExerciseName_MedicineBallChestPass,
	"MedicineBallPushUp":                         PushUpExerciseName_MedicineBallPushUp,
	"WeightedMedicineBallPushUp":                 PushUpExerciseName_WeightedMedicineBallPushUp,
	"OneArmPushUp":                               PushUpExerciseName_OneArmPushUp,
	"WeightedOneArmPushUp":                       PushUpExerciseName_WeightedOneArmPushUp,
	"WeightedPushUp":                             PushUpExerciseName_WeightedPushUp,
	"PushUpAndRow":                               PushUpExerciseName_PushUpAndRow,
	"WeightedPushUpAndRow":                       PushUpExerciseName_WeightedPushUpAndRow,
	"PushUpPlus":                                 PushUpExerciseName_PushUpPlus,
	"WeightedPushUpPlus":                         PushUpExerciseName_WeightedPushUpPlus,
	"PushUpWithFeetOnSwissBall":                  PushUpExerciseName_PushUpWithFeetOnSwissBall,
	"WeightedPushUpWithFeetOnSwissBall":          PushUpExerciseName_WeightedPushUpWithFeetOnSwissBall,
	"PushUpWithOneHandOnMedicineBall":            PushUpExerciseName_PushUpWithOneHandOnMedicineBall,
	"WeightedPushUpWithOneHandOnMedicineBall":    PushUpExerciseName_WeightedPushUpWithOneHandOnMedicineBall,
	"ShoulderPushUp":                             PushUpExerciseName_ShoulderPushUp,
	"WeightedShoulderPushUp":                     PushUpExerciseName_WeightedShoulderPushUp,
	"SingleArmMedicineBallPushUp":                PushUpExerciseName_SingleArmMedicineBallPushUp,
	"WeightedSingleArmMedicineBallPushUp":        PushUpExerciseName_WeightedSingleArmMedicineBallPushUp,
	"SpidermanPushUp":                            PushUpExerciseName_SpidermanPushUp,
	"WeightedSpidermanPushUp":                    PushUpExerciseName_WeightedSpidermanPushUp,
	"StackedFeetPushUp":                          PushUpExerciseName_StackedFeetPushUp,
	"WeightedStackedFeetPushUp":                  PushUpExerciseName_WeightedStackedFeetPushUp,
	"StaggeredHandsPushUp":                       PushUpExerciseName_StaggeredHandsPushUp,
	"WeightedStaggeredHandsPushUp":               PushUpExerciseName_WeightedStaggeredHandsPushUp,
	"SuspendedPushUp":                            PushUpExerciseName_SuspendedPushUp,
	"WeightedSuspendedPushUp":                    PushUpExerciseName_WeightedSuspendedPushUp,
	"SwissBallPushUp":                            PushUpExerciseName_SwissBallPushUp,
	"WeightedSwissBallPushUp":                    PushUpExerciseName_WeightedSwissBallPushUp,
	"SwissBallPushUpPlus":                        PushUpExerciseName_SwissBallPushUpPlus,
	"WeightedSwissBallPushUpPlus":                PushUpExerciseName_WeightedSwissBallPushUpPlus,
	"TPushUp":                                    PushUpExerciseName_TPushUp,
	"WeightedTPushUp":                            PushUpExerciseName_WeightedTPushUp,
	"TripleStopPushUp":                           PushUpExerciseName_TripleStopPushUp,
	"WeightedTripleStopPushUp":                   PushUpExerciseName_WeightedTripleStopPushUp,
	"WideHandsPushUp":                            PushUpExerciseName_WideHandsPushUp,
	"WeightedWideHandsPushUp":                    PushUpExerciseName_WeightedWideHandsPushUp,
	"ParalletteHandstandPushUp":                  PushUpExerciseName_ParalletteHandstandPushUp,
	"WeightedParalletteHandstandPushUp":          PushUpExerciseName_WeightedParalletteHandstandPushUp,
	"RingHandstandPushUp":                        PushUpExerciseName_RingHandstandPushUp,
	"WeightedRingHandstandPushUp":                PushUpExerciseName_WeightedRingHandstandPushUp,
	"RingPushUp":                                 PushUpExerciseName_RingPushUp,
	"WeightedRingPushUp":                         PushUpExerciseName_WeightedRingPushUp,
	"PushUp":                                     PushUpExerciseName_PushUp,
	"PilatesPushup":                              PushUpExerciseName_PilatesPushup,
}

func PushUpExerciseNameFromString(s string) PushUpExerciseName {
	if v, ok := PushUpExerciseNameValues[s]; ok {
		return v
	}
	return PushUpExerciseName(PushUpExerciseName_Invalid)
}
