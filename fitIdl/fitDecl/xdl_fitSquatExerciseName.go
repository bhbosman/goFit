package fitDecl

import strconv "strconv"

type SquatExerciseName uint16

const (
	SquatExerciseName_LegPress                                        SquatExerciseName = 0
	SquatExerciseName_BackSquatWithBodyBar                            SquatExerciseName = 1
	SquatExerciseName_BackSquats                                      SquatExerciseName = 2
	SquatExerciseName_WeightedBackSquats                              SquatExerciseName = 3
	SquatExerciseName_BalancingSquat                                  SquatExerciseName = 4
	SquatExerciseName_WeightedBalancingSquat                          SquatExerciseName = 5
	SquatExerciseName_BarbellBackSquat                                SquatExerciseName = 6
	SquatExerciseName_BarbellBoxSquat                                 SquatExerciseName = 7
	SquatExerciseName_BarbellFrontSquat                               SquatExerciseName = 8
	SquatExerciseName_BarbellHackSquat                                SquatExerciseName = 9
	SquatExerciseName_BarbellHangSquatSnatch                          SquatExerciseName = 10
	SquatExerciseName_BarbellLateralStepUp                            SquatExerciseName = 11
	SquatExerciseName_BarbellQuarterSquat                             SquatExerciseName = 12
	SquatExerciseName_BarbellSiffSquat                                SquatExerciseName = 13
	SquatExerciseName_BarbellSquatSnatch                              SquatExerciseName = 14
	SquatExerciseName_BarbellSquatWithHeelsRaised                     SquatExerciseName = 15
	SquatExerciseName_BarbellStepover                                 SquatExerciseName = 16
	SquatExerciseName_BarbellStepUp                                   SquatExerciseName = 17
	SquatExerciseName_BenchSquatWithRotationalChop                    SquatExerciseName = 18
	SquatExerciseName_WeightedBenchSquatWithRotationalChop            SquatExerciseName = 19
	SquatExerciseName_BodyWeightWallSquat                             SquatExerciseName = 20
	SquatExerciseName_WeightedWallSquat                               SquatExerciseName = 21
	SquatExerciseName_BoxStepSquat                                    SquatExerciseName = 22
	SquatExerciseName_WeightedBoxStepSquat                            SquatExerciseName = 23
	SquatExerciseName_BracedSquat                                     SquatExerciseName = 24
	SquatExerciseName_CrossedArmBarbellFrontSquat                     SquatExerciseName = 25
	SquatExerciseName_CrossoverDumbbellStepUp                         SquatExerciseName = 26
	SquatExerciseName_DumbbellFrontSquat                              SquatExerciseName = 27
	SquatExerciseName_DumbbellSplitSquat                              SquatExerciseName = 28
	SquatExerciseName_DumbbellSquat                                   SquatExerciseName = 29
	SquatExerciseName_DumbbellSquatClean                              SquatExerciseName = 30
	SquatExerciseName_DumbbellStepover                                SquatExerciseName = 31
	SquatExerciseName_DumbbellStepUp                                  SquatExerciseName = 32
	SquatExerciseName_ElevatedSingleLegSquat                          SquatExerciseName = 33
	SquatExerciseName_WeightedElevatedSingleLegSquat                  SquatExerciseName = 34
	SquatExerciseName_FigureFourSquats                                SquatExerciseName = 35
	SquatExerciseName_WeightedFigureFourSquats                        SquatExerciseName = 36
	SquatExerciseName_GobletSquat                                     SquatExerciseName = 37
	SquatExerciseName_KettlebellSquat                                 SquatExerciseName = 38
	SquatExerciseName_KettlebellSwingOverhead                         SquatExerciseName = 39
	SquatExerciseName_KettlebellSwingWithFlipToSquat                  SquatExerciseName = 40
	SquatExerciseName_LateralDumbbellStepUp                           SquatExerciseName = 41
	SquatExerciseName_OneLeggedSquat                                  SquatExerciseName = 42
	SquatExerciseName_OverheadDumbbellSquat                           SquatExerciseName = 43
	SquatExerciseName_OverheadSquat                                   SquatExerciseName = 44
	SquatExerciseName_PartialSingleLegSquat                           SquatExerciseName = 45
	SquatExerciseName_WeightedPartialSingleLegSquat                   SquatExerciseName = 46
	SquatExerciseName_PistolSquat                                     SquatExerciseName = 47
	SquatExerciseName_WeightedPistolSquat                             SquatExerciseName = 48
	SquatExerciseName_PlieSlides                                      SquatExerciseName = 49
	SquatExerciseName_WeightedPlieSlides                              SquatExerciseName = 50
	SquatExerciseName_PlieSquat                                       SquatExerciseName = 51
	SquatExerciseName_WeightedPlieSquat                               SquatExerciseName = 52
	SquatExerciseName_PrisonerSquat                                   SquatExerciseName = 53
	SquatExerciseName_WeightedPrisonerSquat                           SquatExerciseName = 54
	SquatExerciseName_SingleLegBenchGetUp                             SquatExerciseName = 55
	SquatExerciseName_WeightedSingleLegBenchGetUp                     SquatExerciseName = 56
	SquatExerciseName_SingleLegBenchSquat                             SquatExerciseName = 57
	SquatExerciseName_WeightedSingleLegBenchSquat                     SquatExerciseName = 58
	SquatExerciseName_SingleLegSquatOnSwissBall                       SquatExerciseName = 59
	SquatExerciseName_WeightedSingleLegSquatOnSwissBall               SquatExerciseName = 60
	SquatExerciseName_Squat                                           SquatExerciseName = 61
	SquatExerciseName_WeightedSquat                                   SquatExerciseName = 62
	SquatExerciseName_SquatsWithBand                                  SquatExerciseName = 63
	SquatExerciseName_StaggeredSquat                                  SquatExerciseName = 64
	SquatExerciseName_WeightedStaggeredSquat                          SquatExerciseName = 65
	SquatExerciseName_StepUp                                          SquatExerciseName = 66
	SquatExerciseName_WeightedStepUp                                  SquatExerciseName = 67
	SquatExerciseName_SuitcaseSquats                                  SquatExerciseName = 68
	SquatExerciseName_SumoSquat                                       SquatExerciseName = 69
	SquatExerciseName_SumoSquatSlideIn                                SquatExerciseName = 70
	SquatExerciseName_WeightedSumoSquatSlideIn                        SquatExerciseName = 71
	SquatExerciseName_SumoSquatToHighPull                             SquatExerciseName = 72
	SquatExerciseName_SumoSquatToStand                                SquatExerciseName = 73
	SquatExerciseName_WeightedSumoSquatToStand                        SquatExerciseName = 74
	SquatExerciseName_SumoSquatWithRotation                           SquatExerciseName = 75
	SquatExerciseName_WeightedSumoSquatWithRotation                   SquatExerciseName = 76
	SquatExerciseName_SwissBallBodyWeightWallSquat                    SquatExerciseName = 77
	SquatExerciseName_WeightedSwissBallWallSquat                      SquatExerciseName = 78
	SquatExerciseName_Thrusters                                       SquatExerciseName = 79
	SquatExerciseName_UnevenSquat                                     SquatExerciseName = 80
	SquatExerciseName_WeightedUnevenSquat                             SquatExerciseName = 81
	SquatExerciseName_WaistSlimmingSquat                              SquatExerciseName = 82
	SquatExerciseName_WallBall                                        SquatExerciseName = 83
	SquatExerciseName_WideStanceBarbellSquat                          SquatExerciseName = 84
	SquatExerciseName_WideStanceGobletSquat                           SquatExerciseName = 85
	SquatExerciseName_ZercherSquat                                    SquatExerciseName = 86
	SquatExerciseName_SquatAndSideKick                                SquatExerciseName = 88
	SquatExerciseName_SquatJumpsInNOut                                SquatExerciseName = 89
	SquatExerciseName_PilatesPlieSquatsParallelTurnedOutFlatAndHeels  SquatExerciseName = 90
	SquatExerciseName_ReleveStraightLegAndKneeBentWithOneLegVariation SquatExerciseName = 91
	SquatExerciseName_Invalid                                         SquatExerciseName = 65535
)

var SquatExerciseNameNames = map[SquatExerciseName]string{
	SquatExerciseName_LegPress:                                        "LegPress",
	SquatExerciseName_BackSquatWithBodyBar:                            "BackSquatWithBodyBar",
	SquatExerciseName_BackSquats:                                      "BackSquats",
	SquatExerciseName_WeightedBackSquats:                              "WeightedBackSquats",
	SquatExerciseName_BalancingSquat:                                  "BalancingSquat",
	SquatExerciseName_WeightedBalancingSquat:                          "WeightedBalancingSquat",
	SquatExerciseName_BarbellBackSquat:                                "BarbellBackSquat",
	SquatExerciseName_BarbellBoxSquat:                                 "BarbellBoxSquat",
	SquatExerciseName_BarbellFrontSquat:                               "BarbellFrontSquat",
	SquatExerciseName_BarbellHackSquat:                                "BarbellHackSquat",
	SquatExerciseName_BarbellHangSquatSnatch:                          "BarbellHangSquatSnatch",
	SquatExerciseName_BarbellLateralStepUp:                            "BarbellLateralStepUp",
	SquatExerciseName_BarbellQuarterSquat:                             "BarbellQuarterSquat",
	SquatExerciseName_BarbellSiffSquat:                                "BarbellSiffSquat",
	SquatExerciseName_BarbellSquatSnatch:                              "BarbellSquatSnatch",
	SquatExerciseName_BarbellSquatWithHeelsRaised:                     "BarbellSquatWithHeelsRaised",
	SquatExerciseName_BarbellStepover:                                 "BarbellStepover",
	SquatExerciseName_BarbellStepUp:                                   "BarbellStepUp",
	SquatExerciseName_BenchSquatWithRotationalChop:                    "BenchSquatWithRotationalChop",
	SquatExerciseName_WeightedBenchSquatWithRotationalChop:            "WeightedBenchSquatWithRotationalChop",
	SquatExerciseName_BodyWeightWallSquat:                             "BodyWeightWallSquat",
	SquatExerciseName_WeightedWallSquat:                               "WeightedWallSquat",
	SquatExerciseName_BoxStepSquat:                                    "BoxStepSquat",
	SquatExerciseName_WeightedBoxStepSquat:                            "WeightedBoxStepSquat",
	SquatExerciseName_BracedSquat:                                     "BracedSquat",
	SquatExerciseName_CrossedArmBarbellFrontSquat:                     "CrossedArmBarbellFrontSquat",
	SquatExerciseName_CrossoverDumbbellStepUp:                         "CrossoverDumbbellStepUp",
	SquatExerciseName_DumbbellFrontSquat:                              "DumbbellFrontSquat",
	SquatExerciseName_DumbbellSplitSquat:                              "DumbbellSplitSquat",
	SquatExerciseName_DumbbellSquat:                                   "DumbbellSquat",
	SquatExerciseName_DumbbellSquatClean:                              "DumbbellSquatClean",
	SquatExerciseName_DumbbellStepover:                                "DumbbellStepover",
	SquatExerciseName_DumbbellStepUp:                                  "DumbbellStepUp",
	SquatExerciseName_ElevatedSingleLegSquat:                          "ElevatedSingleLegSquat",
	SquatExerciseName_WeightedElevatedSingleLegSquat:                  "WeightedElevatedSingleLegSquat",
	SquatExerciseName_FigureFourSquats:                                "FigureFourSquats",
	SquatExerciseName_WeightedFigureFourSquats:                        "WeightedFigureFourSquats",
	SquatExerciseName_GobletSquat:                                     "GobletSquat",
	SquatExerciseName_KettlebellSquat:                                 "KettlebellSquat",
	SquatExerciseName_KettlebellSwingOverhead:                         "KettlebellSwingOverhead",
	SquatExerciseName_KettlebellSwingWithFlipToSquat:                  "KettlebellSwingWithFlipToSquat",
	SquatExerciseName_LateralDumbbellStepUp:                           "LateralDumbbellStepUp",
	SquatExerciseName_OneLeggedSquat:                                  "OneLeggedSquat",
	SquatExerciseName_OverheadDumbbellSquat:                           "OverheadDumbbellSquat",
	SquatExerciseName_OverheadSquat:                                   "OverheadSquat",
	SquatExerciseName_PartialSingleLegSquat:                           "PartialSingleLegSquat",
	SquatExerciseName_WeightedPartialSingleLegSquat:                   "WeightedPartialSingleLegSquat",
	SquatExerciseName_PistolSquat:                                     "PistolSquat",
	SquatExerciseName_WeightedPistolSquat:                             "WeightedPistolSquat",
	SquatExerciseName_PlieSlides:                                      "PlieSlides",
	SquatExerciseName_WeightedPlieSlides:                              "WeightedPlieSlides",
	SquatExerciseName_PlieSquat:                                       "PlieSquat",
	SquatExerciseName_WeightedPlieSquat:                               "WeightedPlieSquat",
	SquatExerciseName_PrisonerSquat:                                   "PrisonerSquat",
	SquatExerciseName_WeightedPrisonerSquat:                           "WeightedPrisonerSquat",
	SquatExerciseName_SingleLegBenchGetUp:                             "SingleLegBenchGetUp",
	SquatExerciseName_WeightedSingleLegBenchGetUp:                     "WeightedSingleLegBenchGetUp",
	SquatExerciseName_SingleLegBenchSquat:                             "SingleLegBenchSquat",
	SquatExerciseName_WeightedSingleLegBenchSquat:                     "WeightedSingleLegBenchSquat",
	SquatExerciseName_SingleLegSquatOnSwissBall:                       "SingleLegSquatOnSwissBall",
	SquatExerciseName_WeightedSingleLegSquatOnSwissBall:               "WeightedSingleLegSquatOnSwissBall",
	SquatExerciseName_Squat:                                           "Squat",
	SquatExerciseName_WeightedSquat:                                   "WeightedSquat",
	SquatExerciseName_SquatsWithBand:                                  "SquatsWithBand",
	SquatExerciseName_StaggeredSquat:                                  "StaggeredSquat",
	SquatExerciseName_WeightedStaggeredSquat:                          "WeightedStaggeredSquat",
	SquatExerciseName_StepUp:                                          "StepUp",
	SquatExerciseName_WeightedStepUp:                                  "WeightedStepUp",
	SquatExerciseName_SuitcaseSquats:                                  "SuitcaseSquats",
	SquatExerciseName_SumoSquat:                                       "SumoSquat",
	SquatExerciseName_SumoSquatSlideIn:                                "SumoSquatSlideIn",
	SquatExerciseName_WeightedSumoSquatSlideIn:                        "WeightedSumoSquatSlideIn",
	SquatExerciseName_SumoSquatToHighPull:                             "SumoSquatToHighPull",
	SquatExerciseName_SumoSquatToStand:                                "SumoSquatToStand",
	SquatExerciseName_WeightedSumoSquatToStand:                        "WeightedSumoSquatToStand",
	SquatExerciseName_SumoSquatWithRotation:                           "SumoSquatWithRotation",
	SquatExerciseName_WeightedSumoSquatWithRotation:                   "WeightedSumoSquatWithRotation",
	SquatExerciseName_SwissBallBodyWeightWallSquat:                    "SwissBallBodyWeightWallSquat",
	SquatExerciseName_WeightedSwissBallWallSquat:                      "WeightedSwissBallWallSquat",
	SquatExerciseName_Thrusters:                                       "Thrusters",
	SquatExerciseName_UnevenSquat:                                     "UnevenSquat",
	SquatExerciseName_WeightedUnevenSquat:                             "WeightedUnevenSquat",
	SquatExerciseName_WaistSlimmingSquat:                              "WaistSlimmingSquat",
	SquatExerciseName_WallBall:                                        "WallBall",
	SquatExerciseName_WideStanceBarbellSquat:                          "WideStanceBarbellSquat",
	SquatExerciseName_WideStanceGobletSquat:                           "WideStanceGobletSquat",
	SquatExerciseName_ZercherSquat:                                    "ZercherSquat",
	SquatExerciseName_SquatAndSideKick:                                "SquatAndSideKick",
	SquatExerciseName_SquatJumpsInNOut:                                "SquatJumpsInNOut",
	SquatExerciseName_PilatesPlieSquatsParallelTurnedOutFlatAndHeels:  "PilatesPlieSquatsParallelTurnedOutFlatAndHeels",
	SquatExerciseName_ReleveStraightLegAndKneeBentWithOneLegVariation: "ReleveStraightLegAndKneeBentWithOneLegVariation",
}

func (k SquatExerciseName) String() string {
	if uint(k) < uint(len(SquatExerciseNameNames)) {
		if v, ok := SquatExerciseNameNames[k]; ok {
			return v
		}
	}
	return "SquatExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var SquatExerciseNameValues = map[string]SquatExerciseName{
	"LegPress":                                       SquatExerciseName_LegPress,
	"BackSquatWithBodyBar":                           SquatExerciseName_BackSquatWithBodyBar,
	"BackSquats":                                     SquatExerciseName_BackSquats,
	"WeightedBackSquats":                             SquatExerciseName_WeightedBackSquats,
	"BalancingSquat":                                 SquatExerciseName_BalancingSquat,
	"WeightedBalancingSquat":                         SquatExerciseName_WeightedBalancingSquat,
	"BarbellBackSquat":                               SquatExerciseName_BarbellBackSquat,
	"BarbellBoxSquat":                                SquatExerciseName_BarbellBoxSquat,
	"BarbellFrontSquat":                              SquatExerciseName_BarbellFrontSquat,
	"BarbellHackSquat":                               SquatExerciseName_BarbellHackSquat,
	"BarbellHangSquatSnatch":                         SquatExerciseName_BarbellHangSquatSnatch,
	"BarbellLateralStepUp":                           SquatExerciseName_BarbellLateralStepUp,
	"BarbellQuarterSquat":                            SquatExerciseName_BarbellQuarterSquat,
	"BarbellSiffSquat":                               SquatExerciseName_BarbellSiffSquat,
	"BarbellSquatSnatch":                             SquatExerciseName_BarbellSquatSnatch,
	"BarbellSquatWithHeelsRaised":                    SquatExerciseName_BarbellSquatWithHeelsRaised,
	"BarbellStepover":                                SquatExerciseName_BarbellStepover,
	"BarbellStepUp":                                  SquatExerciseName_BarbellStepUp,
	"BenchSquatWithRotationalChop":                   SquatExerciseName_BenchSquatWithRotationalChop,
	"WeightedBenchSquatWithRotationalChop":           SquatExerciseName_WeightedBenchSquatWithRotationalChop,
	"BodyWeightWallSquat":                            SquatExerciseName_BodyWeightWallSquat,
	"WeightedWallSquat":                              SquatExerciseName_WeightedWallSquat,
	"BoxStepSquat":                                   SquatExerciseName_BoxStepSquat,
	"WeightedBoxStepSquat":                           SquatExerciseName_WeightedBoxStepSquat,
	"BracedSquat":                                    SquatExerciseName_BracedSquat,
	"CrossedArmBarbellFrontSquat":                    SquatExerciseName_CrossedArmBarbellFrontSquat,
	"CrossoverDumbbellStepUp":                        SquatExerciseName_CrossoverDumbbellStepUp,
	"DumbbellFrontSquat":                             SquatExerciseName_DumbbellFrontSquat,
	"DumbbellSplitSquat":                             SquatExerciseName_DumbbellSplitSquat,
	"DumbbellSquat":                                  SquatExerciseName_DumbbellSquat,
	"DumbbellSquatClean":                             SquatExerciseName_DumbbellSquatClean,
	"DumbbellStepover":                               SquatExerciseName_DumbbellStepover,
	"DumbbellStepUp":                                 SquatExerciseName_DumbbellStepUp,
	"ElevatedSingleLegSquat":                         SquatExerciseName_ElevatedSingleLegSquat,
	"WeightedElevatedSingleLegSquat":                 SquatExerciseName_WeightedElevatedSingleLegSquat,
	"FigureFourSquats":                               SquatExerciseName_FigureFourSquats,
	"WeightedFigureFourSquats":                       SquatExerciseName_WeightedFigureFourSquats,
	"GobletSquat":                                    SquatExerciseName_GobletSquat,
	"KettlebellSquat":                                SquatExerciseName_KettlebellSquat,
	"KettlebellSwingOverhead":                        SquatExerciseName_KettlebellSwingOverhead,
	"KettlebellSwingWithFlipToSquat":                 SquatExerciseName_KettlebellSwingWithFlipToSquat,
	"LateralDumbbellStepUp":                          SquatExerciseName_LateralDumbbellStepUp,
	"OneLeggedSquat":                                 SquatExerciseName_OneLeggedSquat,
	"OverheadDumbbellSquat":                          SquatExerciseName_OverheadDumbbellSquat,
	"OverheadSquat":                                  SquatExerciseName_OverheadSquat,
	"PartialSingleLegSquat":                          SquatExerciseName_PartialSingleLegSquat,
	"WeightedPartialSingleLegSquat":                  SquatExerciseName_WeightedPartialSingleLegSquat,
	"PistolSquat":                                    SquatExerciseName_PistolSquat,
	"WeightedPistolSquat":                            SquatExerciseName_WeightedPistolSquat,
	"PlieSlides":                                     SquatExerciseName_PlieSlides,
	"WeightedPlieSlides":                             SquatExerciseName_WeightedPlieSlides,
	"PlieSquat":                                      SquatExerciseName_PlieSquat,
	"WeightedPlieSquat":                              SquatExerciseName_WeightedPlieSquat,
	"PrisonerSquat":                                  SquatExerciseName_PrisonerSquat,
	"WeightedPrisonerSquat":                          SquatExerciseName_WeightedPrisonerSquat,
	"SingleLegBenchGetUp":                            SquatExerciseName_SingleLegBenchGetUp,
	"WeightedSingleLegBenchGetUp":                    SquatExerciseName_WeightedSingleLegBenchGetUp,
	"SingleLegBenchSquat":                            SquatExerciseName_SingleLegBenchSquat,
	"WeightedSingleLegBenchSquat":                    SquatExerciseName_WeightedSingleLegBenchSquat,
	"SingleLegSquatOnSwissBall":                      SquatExerciseName_SingleLegSquatOnSwissBall,
	"WeightedSingleLegSquatOnSwissBall":              SquatExerciseName_WeightedSingleLegSquatOnSwissBall,
	"Squat":                                          SquatExerciseName_Squat,
	"WeightedSquat":                                  SquatExerciseName_WeightedSquat,
	"SquatsWithBand":                                 SquatExerciseName_SquatsWithBand,
	"StaggeredSquat":                                 SquatExerciseName_StaggeredSquat,
	"WeightedStaggeredSquat":                         SquatExerciseName_WeightedStaggeredSquat,
	"StepUp":                                         SquatExerciseName_StepUp,
	"WeightedStepUp":                                 SquatExerciseName_WeightedStepUp,
	"SuitcaseSquats":                                 SquatExerciseName_SuitcaseSquats,
	"SumoSquat":                                      SquatExerciseName_SumoSquat,
	"SumoSquatSlideIn":                               SquatExerciseName_SumoSquatSlideIn,
	"WeightedSumoSquatSlideIn":                       SquatExerciseName_WeightedSumoSquatSlideIn,
	"SumoSquatToHighPull":                            SquatExerciseName_SumoSquatToHighPull,
	"SumoSquatToStand":                               SquatExerciseName_SumoSquatToStand,
	"WeightedSumoSquatToStand":                       SquatExerciseName_WeightedSumoSquatToStand,
	"SumoSquatWithRotation":                          SquatExerciseName_SumoSquatWithRotation,
	"WeightedSumoSquatWithRotation":                  SquatExerciseName_WeightedSumoSquatWithRotation,
	"SwissBallBodyWeightWallSquat":                   SquatExerciseName_SwissBallBodyWeightWallSquat,
	"WeightedSwissBallWallSquat":                     SquatExerciseName_WeightedSwissBallWallSquat,
	"Thrusters":                                      SquatExerciseName_Thrusters,
	"UnevenSquat":                                    SquatExerciseName_UnevenSquat,
	"WeightedUnevenSquat":                            SquatExerciseName_WeightedUnevenSquat,
	"WaistSlimmingSquat":                             SquatExerciseName_WaistSlimmingSquat,
	"WallBall":                                       SquatExerciseName_WallBall,
	"WideStanceBarbellSquat":                         SquatExerciseName_WideStanceBarbellSquat,
	"WideStanceGobletSquat":                          SquatExerciseName_WideStanceGobletSquat,
	"ZercherSquat":                                   SquatExerciseName_ZercherSquat,
	"SquatAndSideKick":                               SquatExerciseName_SquatAndSideKick,
	"SquatJumpsInNOut":                               SquatExerciseName_SquatJumpsInNOut,
	"PilatesPlieSquatsParallelTurnedOutFlatAndHeels": SquatExerciseName_PilatesPlieSquatsParallelTurnedOutFlatAndHeels,
	"ReleveStraightLegAndKneeBentWithOneLegVariation": SquatExerciseName_ReleveStraightLegAndKneeBentWithOneLegVariation,
}

func SquatExerciseNameFromString(s string) SquatExerciseName {
	if v, ok := SquatExerciseNameValues[s]; ok {
		return v
	}
	return SquatExerciseName(SquatExerciseName_Invalid)
}
