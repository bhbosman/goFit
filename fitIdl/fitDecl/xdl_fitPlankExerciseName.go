package fitDecl

import strconv "strconv"

type PlankExerciseName uint16

const (
	PlankExerciseName__45DegreePlank                                   PlankExerciseName = 0
	PlankExerciseName_Weighted45DegreePlank                            PlankExerciseName = 1
	PlankExerciseName__90DegreeStaticHold                              PlankExerciseName = 2
	PlankExerciseName_Weighted90DegreeStaticHold                       PlankExerciseName = 3
	PlankExerciseName_BearCrawl                                        PlankExerciseName = 4
	PlankExerciseName_WeightedBearCrawl                                PlankExerciseName = 5
	PlankExerciseName_CrossBodyMountainClimber                         PlankExerciseName = 6
	PlankExerciseName_WeightedCrossBodyMountainClimber                 PlankExerciseName = 7
	PlankExerciseName_ElbowPlankPikeJacks                              PlankExerciseName = 8
	PlankExerciseName_WeightedElbowPlankPikeJacks                      PlankExerciseName = 9
	PlankExerciseName_ElevatedFeetPlank                                PlankExerciseName = 10
	PlankExerciseName_WeightedElevatedFeetPlank                        PlankExerciseName = 11
	PlankExerciseName_ElevatorAbs                                      PlankExerciseName = 12
	PlankExerciseName_WeightedElevatorAbs                              PlankExerciseName = 13
	PlankExerciseName_ExtendedPlank                                    PlankExerciseName = 14
	PlankExerciseName_WeightedExtendedPlank                            PlankExerciseName = 15
	PlankExerciseName_FullPlankPasseTwist                              PlankExerciseName = 16
	PlankExerciseName_WeightedFullPlankPasseTwist                      PlankExerciseName = 17
	PlankExerciseName_InchingElbowPlank                                PlankExerciseName = 18
	PlankExerciseName_WeightedInchingElbowPlank                        PlankExerciseName = 19
	PlankExerciseName_InchwormToSidePlank                              PlankExerciseName = 20
	PlankExerciseName_WeightedInchwormToSidePlank                      PlankExerciseName = 21
	PlankExerciseName_KneelingPlank                                    PlankExerciseName = 22
	PlankExerciseName_WeightedKneelingPlank                            PlankExerciseName = 23
	PlankExerciseName_KneelingSidePlankWithLegLift                     PlankExerciseName = 24
	PlankExerciseName_WeightedKneelingSidePlankWithLegLift             PlankExerciseName = 25
	PlankExerciseName_LateralRoll                                      PlankExerciseName = 26
	PlankExerciseName_WeightedLateralRoll                              PlankExerciseName = 27
	PlankExerciseName_LyingReversePlank                                PlankExerciseName = 28
	PlankExerciseName_WeightedLyingReversePlank                        PlankExerciseName = 29
	PlankExerciseName_MedicineBallMountainClimber                      PlankExerciseName = 30
	PlankExerciseName_WeightedMedicineBallMountainClimber              PlankExerciseName = 31
	PlankExerciseName_ModifiedMountainClimberAndExtension              PlankExerciseName = 32
	PlankExerciseName_WeightedModifiedMountainClimberAndExtension      PlankExerciseName = 33
	PlankExerciseName_MountainClimber                                  PlankExerciseName = 34
	PlankExerciseName_WeightedMountainClimber                          PlankExerciseName = 35
	PlankExerciseName_MountainClimberOnSlidingDiscs                    PlankExerciseName = 36
	PlankExerciseName_WeightedMountainClimberOnSlidingDiscs            PlankExerciseName = 37
	PlankExerciseName_MountainClimberWithFeetOnBosuBall                PlankExerciseName = 38
	PlankExerciseName_WeightedMountainClimberWithFeetOnBosuBall        PlankExerciseName = 39
	PlankExerciseName_MountainClimberWithHandsOnBench                  PlankExerciseName = 40
	PlankExerciseName_MountainClimberWithHandsOnSwissBall              PlankExerciseName = 41
	PlankExerciseName_WeightedMountainClimberWithHandsOnSwissBall      PlankExerciseName = 42
	PlankExerciseName_Plank                                            PlankExerciseName = 43
	PlankExerciseName_PlankJacksWithFeetOnSlidingDiscs                 PlankExerciseName = 44
	PlankExerciseName_WeightedPlankJacksWithFeetOnSlidingDiscs         PlankExerciseName = 45
	PlankExerciseName_PlankKneeTwist                                   PlankExerciseName = 46
	PlankExerciseName_WeightedPlankKneeTwist                           PlankExerciseName = 47
	PlankExerciseName_PlankPikeJumps                                   PlankExerciseName = 48
	PlankExerciseName_WeightedPlankPikeJumps                           PlankExerciseName = 49
	PlankExerciseName_PlankPikes                                       PlankExerciseName = 50
	PlankExerciseName_WeightedPlankPikes                               PlankExerciseName = 51
	PlankExerciseName_PlankToStandUp                                   PlankExerciseName = 52
	PlankExerciseName_WeightedPlankToStandUp                           PlankExerciseName = 53
	PlankExerciseName_PlankWithArmRaise                                PlankExerciseName = 54
	PlankExerciseName_WeightedPlankWithArmRaise                        PlankExerciseName = 55
	PlankExerciseName_PlankWithKneeToElbow                             PlankExerciseName = 56
	PlankExerciseName_WeightedPlankWithKneeToElbow                     PlankExerciseName = 57
	PlankExerciseName_PlankWithObliqueCrunch                           PlankExerciseName = 58
	PlankExerciseName_WeightedPlankWithObliqueCrunch                   PlankExerciseName = 59
	PlankExerciseName_PlyometricSidePlank                              PlankExerciseName = 60
	PlankExerciseName_WeightedPlyometricSidePlank                      PlankExerciseName = 61
	PlankExerciseName_RollingSidePlank                                 PlankExerciseName = 62
	PlankExerciseName_WeightedRollingSidePlank                         PlankExerciseName = 63
	PlankExerciseName_SideKickPlank                                    PlankExerciseName = 64
	PlankExerciseName_WeightedSideKickPlank                            PlankExerciseName = 65
	PlankExerciseName_SidePlank                                        PlankExerciseName = 66
	PlankExerciseName_WeightedSidePlank                                PlankExerciseName = 67
	PlankExerciseName_SidePlankAndRow                                  PlankExerciseName = 68
	PlankExerciseName_WeightedSidePlankAndRow                          PlankExerciseName = 69
	PlankExerciseName_SidePlankLift                                    PlankExerciseName = 70
	PlankExerciseName_WeightedSidePlankLift                            PlankExerciseName = 71
	PlankExerciseName_SidePlankWithElbowOnBosuBall                     PlankExerciseName = 72
	PlankExerciseName_WeightedSidePlankWithElbowOnBosuBall             PlankExerciseName = 73
	PlankExerciseName_SidePlankWithFeetOnBench                         PlankExerciseName = 74
	PlankExerciseName_WeightedSidePlankWithFeetOnBench                 PlankExerciseName = 75
	PlankExerciseName_SidePlankWithKneeCircle                          PlankExerciseName = 76
	PlankExerciseName_WeightedSidePlankWithKneeCircle                  PlankExerciseName = 77
	PlankExerciseName_SidePlankWithKneeTuck                            PlankExerciseName = 78
	PlankExerciseName_WeightedSidePlankWithKneeTuck                    PlankExerciseName = 79
	PlankExerciseName_SidePlankWithLegLift                             PlankExerciseName = 80
	PlankExerciseName_WeightedSidePlankWithLegLift                     PlankExerciseName = 81
	PlankExerciseName_SidePlankWithReachUnder                          PlankExerciseName = 82
	PlankExerciseName_WeightedSidePlankWithReachUnder                  PlankExerciseName = 83
	PlankExerciseName_SingleLegElevatedFeetPlank                       PlankExerciseName = 84
	PlankExerciseName_WeightedSingleLegElevatedFeetPlank               PlankExerciseName = 85
	PlankExerciseName_SingleLegFlexAndExtend                           PlankExerciseName = 86
	PlankExerciseName_WeightedSingleLegFlexAndExtend                   PlankExerciseName = 87
	PlankExerciseName_SingleLegSidePlank                               PlankExerciseName = 88
	PlankExerciseName_WeightedSingleLegSidePlank                       PlankExerciseName = 89
	PlankExerciseName_SpidermanPlank                                   PlankExerciseName = 90
	PlankExerciseName_WeightedSpidermanPlank                           PlankExerciseName = 91
	PlankExerciseName_StraightArmPlank                                 PlankExerciseName = 92
	PlankExerciseName_WeightedStraightArmPlank                         PlankExerciseName = 93
	PlankExerciseName_StraightArmPlankWithShoulderTouch                PlankExerciseName = 94
	PlankExerciseName_WeightedStraightArmPlankWithShoulderTouch        PlankExerciseName = 95
	PlankExerciseName_SwissBallPlank                                   PlankExerciseName = 96
	PlankExerciseName_WeightedSwissBallPlank                           PlankExerciseName = 97
	PlankExerciseName_SwissBallPlankLegLift                            PlankExerciseName = 98
	PlankExerciseName_WeightedSwissBallPlankLegLift                    PlankExerciseName = 99
	PlankExerciseName_SwissBallPlankLegLiftAndHold                     PlankExerciseName = 100
	PlankExerciseName_SwissBallPlankWithFeetOnBench                    PlankExerciseName = 101
	PlankExerciseName_WeightedSwissBallPlankWithFeetOnBench            PlankExerciseName = 102
	PlankExerciseName_SwissBallProneJackknife                          PlankExerciseName = 103
	PlankExerciseName_WeightedSwissBallProneJackknife                  PlankExerciseName = 104
	PlankExerciseName_SwissBallSidePlank                               PlankExerciseName = 105
	PlankExerciseName_WeightedSwissBallSidePlank                       PlankExerciseName = 106
	PlankExerciseName_ThreeWayPlank                                    PlankExerciseName = 107
	PlankExerciseName_WeightedThreeWayPlank                            PlankExerciseName = 108
	PlankExerciseName_TowelPlankAndKneeIn                              PlankExerciseName = 109
	PlankExerciseName_WeightedTowelPlankAndKneeIn                      PlankExerciseName = 110
	PlankExerciseName_TStabilization                                   PlankExerciseName = 111
	PlankExerciseName_WeightedTStabilization                           PlankExerciseName = 112
	PlankExerciseName_TurkishGetUpToSidePlank                          PlankExerciseName = 113
	PlankExerciseName_WeightedTurkishGetUpToSidePlank                  PlankExerciseName = 114
	PlankExerciseName_TwoPointPlank                                    PlankExerciseName = 115
	PlankExerciseName_WeightedTwoPointPlank                            PlankExerciseName = 116
	PlankExerciseName_WeightedPlank                                    PlankExerciseName = 117
	PlankExerciseName_WideStancePlankWithDiagonalArmLift               PlankExerciseName = 118
	PlankExerciseName_WeightedWideStancePlankWithDiagonalArmLift       PlankExerciseName = 119
	PlankExerciseName_WideStancePlankWithDiagonalLegLift               PlankExerciseName = 120
	PlankExerciseName_WeightedWideStancePlankWithDiagonalLegLift       PlankExerciseName = 121
	PlankExerciseName_WideStancePlankWithLegLift                       PlankExerciseName = 122
	PlankExerciseName_WeightedWideStancePlankWithLegLift               PlankExerciseName = 123
	PlankExerciseName_WideStancePlankWithOppositeArmAndLegLift         PlankExerciseName = 124
	PlankExerciseName_WeightedMountainClimberWithHandsOnBench          PlankExerciseName = 125
	PlankExerciseName_WeightedSwissBallPlankLegLiftAndHold             PlankExerciseName = 126
	PlankExerciseName_WeightedWideStancePlankWithOppositeArmAndLegLift PlankExerciseName = 127
	PlankExerciseName_PlankWithFeetOnSwissBall                         PlankExerciseName = 128
	PlankExerciseName_SidePlankToPlankWithReachUnder                   PlankExerciseName = 129
	PlankExerciseName_BridgeWithGluteLowerLift                         PlankExerciseName = 130
	PlankExerciseName_BridgeOneLegBridge                               PlankExerciseName = 131
	PlankExerciseName_PlankWithArmVariations                           PlankExerciseName = 132
	PlankExerciseName_PlankWithLegLift                                 PlankExerciseName = 133
	PlankExerciseName_ReversePlankWithLegPull                          PlankExerciseName = 134
	PlankExerciseName_Invalid                                          PlankExerciseName = 65535
)

var PlankExerciseNameNames = map[PlankExerciseName]string{
	PlankExerciseName__45DegreePlank:                                   "_45DegreePlank",
	PlankExerciseName_Weighted45DegreePlank:                            "Weighted45DegreePlank",
	PlankExerciseName__90DegreeStaticHold:                              "_90DegreeStaticHold",
	PlankExerciseName_Weighted90DegreeStaticHold:                       "Weighted90DegreeStaticHold",
	PlankExerciseName_BearCrawl:                                        "BearCrawl",
	PlankExerciseName_WeightedBearCrawl:                                "WeightedBearCrawl",
	PlankExerciseName_CrossBodyMountainClimber:                         "CrossBodyMountainClimber",
	PlankExerciseName_WeightedCrossBodyMountainClimber:                 "WeightedCrossBodyMountainClimber",
	PlankExerciseName_ElbowPlankPikeJacks:                              "ElbowPlankPikeJacks",
	PlankExerciseName_WeightedElbowPlankPikeJacks:                      "WeightedElbowPlankPikeJacks",
	PlankExerciseName_ElevatedFeetPlank:                                "ElevatedFeetPlank",
	PlankExerciseName_WeightedElevatedFeetPlank:                        "WeightedElevatedFeetPlank",
	PlankExerciseName_ElevatorAbs:                                      "ElevatorAbs",
	PlankExerciseName_WeightedElevatorAbs:                              "WeightedElevatorAbs",
	PlankExerciseName_ExtendedPlank:                                    "ExtendedPlank",
	PlankExerciseName_WeightedExtendedPlank:                            "WeightedExtendedPlank",
	PlankExerciseName_FullPlankPasseTwist:                              "FullPlankPasseTwist",
	PlankExerciseName_WeightedFullPlankPasseTwist:                      "WeightedFullPlankPasseTwist",
	PlankExerciseName_InchingElbowPlank:                                "InchingElbowPlank",
	PlankExerciseName_WeightedInchingElbowPlank:                        "WeightedInchingElbowPlank",
	PlankExerciseName_InchwormToSidePlank:                              "InchwormToSidePlank",
	PlankExerciseName_WeightedInchwormToSidePlank:                      "WeightedInchwormToSidePlank",
	PlankExerciseName_KneelingPlank:                                    "KneelingPlank",
	PlankExerciseName_WeightedKneelingPlank:                            "WeightedKneelingPlank",
	PlankExerciseName_KneelingSidePlankWithLegLift:                     "KneelingSidePlankWithLegLift",
	PlankExerciseName_WeightedKneelingSidePlankWithLegLift:             "WeightedKneelingSidePlankWithLegLift",
	PlankExerciseName_LateralRoll:                                      "LateralRoll",
	PlankExerciseName_WeightedLateralRoll:                              "WeightedLateralRoll",
	PlankExerciseName_LyingReversePlank:                                "LyingReversePlank",
	PlankExerciseName_WeightedLyingReversePlank:                        "WeightedLyingReversePlank",
	PlankExerciseName_MedicineBallMountainClimber:                      "MedicineBallMountainClimber",
	PlankExerciseName_WeightedMedicineBallMountainClimber:              "WeightedMedicineBallMountainClimber",
	PlankExerciseName_ModifiedMountainClimberAndExtension:              "ModifiedMountainClimberAndExtension",
	PlankExerciseName_WeightedModifiedMountainClimberAndExtension:      "WeightedModifiedMountainClimberAndExtension",
	PlankExerciseName_MountainClimber:                                  "MountainClimber",
	PlankExerciseName_WeightedMountainClimber:                          "WeightedMountainClimber",
	PlankExerciseName_MountainClimberOnSlidingDiscs:                    "MountainClimberOnSlidingDiscs",
	PlankExerciseName_WeightedMountainClimberOnSlidingDiscs:            "WeightedMountainClimberOnSlidingDiscs",
	PlankExerciseName_MountainClimberWithFeetOnBosuBall:                "MountainClimberWithFeetOnBosuBall",
	PlankExerciseName_WeightedMountainClimberWithFeetOnBosuBall:        "WeightedMountainClimberWithFeetOnBosuBall",
	PlankExerciseName_MountainClimberWithHandsOnBench:                  "MountainClimberWithHandsOnBench",
	PlankExerciseName_MountainClimberWithHandsOnSwissBall:              "MountainClimberWithHandsOnSwissBall",
	PlankExerciseName_WeightedMountainClimberWithHandsOnSwissBall:      "WeightedMountainClimberWithHandsOnSwissBall",
	PlankExerciseName_Plank:                                            "Plank",
	PlankExerciseName_PlankJacksWithFeetOnSlidingDiscs:                 "PlankJacksWithFeetOnSlidingDiscs",
	PlankExerciseName_WeightedPlankJacksWithFeetOnSlidingDiscs:         "WeightedPlankJacksWithFeetOnSlidingDiscs",
	PlankExerciseName_PlankKneeTwist:                                   "PlankKneeTwist",
	PlankExerciseName_WeightedPlankKneeTwist:                           "WeightedPlankKneeTwist",
	PlankExerciseName_PlankPikeJumps:                                   "PlankPikeJumps",
	PlankExerciseName_WeightedPlankPikeJumps:                           "WeightedPlankPikeJumps",
	PlankExerciseName_PlankPikes:                                       "PlankPikes",
	PlankExerciseName_WeightedPlankPikes:                               "WeightedPlankPikes",
	PlankExerciseName_PlankToStandUp:                                   "PlankToStandUp",
	PlankExerciseName_WeightedPlankToStandUp:                           "WeightedPlankToStandUp",
	PlankExerciseName_PlankWithArmRaise:                                "PlankWithArmRaise",
	PlankExerciseName_WeightedPlankWithArmRaise:                        "WeightedPlankWithArmRaise",
	PlankExerciseName_PlankWithKneeToElbow:                             "PlankWithKneeToElbow",
	PlankExerciseName_WeightedPlankWithKneeToElbow:                     "WeightedPlankWithKneeToElbow",
	PlankExerciseName_PlankWithObliqueCrunch:                           "PlankWithObliqueCrunch",
	PlankExerciseName_WeightedPlankWithObliqueCrunch:                   "WeightedPlankWithObliqueCrunch",
	PlankExerciseName_PlyometricSidePlank:                              "PlyometricSidePlank",
	PlankExerciseName_WeightedPlyometricSidePlank:                      "WeightedPlyometricSidePlank",
	PlankExerciseName_RollingSidePlank:                                 "RollingSidePlank",
	PlankExerciseName_WeightedRollingSidePlank:                         "WeightedRollingSidePlank",
	PlankExerciseName_SideKickPlank:                                    "SideKickPlank",
	PlankExerciseName_WeightedSideKickPlank:                            "WeightedSideKickPlank",
	PlankExerciseName_SidePlank:                                        "SidePlank",
	PlankExerciseName_WeightedSidePlank:                                "WeightedSidePlank",
	PlankExerciseName_SidePlankAndRow:                                  "SidePlankAndRow",
	PlankExerciseName_WeightedSidePlankAndRow:                          "WeightedSidePlankAndRow",
	PlankExerciseName_SidePlankLift:                                    "SidePlankLift",
	PlankExerciseName_WeightedSidePlankLift:                            "WeightedSidePlankLift",
	PlankExerciseName_SidePlankWithElbowOnBosuBall:                     "SidePlankWithElbowOnBosuBall",
	PlankExerciseName_WeightedSidePlankWithElbowOnBosuBall:             "WeightedSidePlankWithElbowOnBosuBall",
	PlankExerciseName_SidePlankWithFeetOnBench:                         "SidePlankWithFeetOnBench",
	PlankExerciseName_WeightedSidePlankWithFeetOnBench:                 "WeightedSidePlankWithFeetOnBench",
	PlankExerciseName_SidePlankWithKneeCircle:                          "SidePlankWithKneeCircle",
	PlankExerciseName_WeightedSidePlankWithKneeCircle:                  "WeightedSidePlankWithKneeCircle",
	PlankExerciseName_SidePlankWithKneeTuck:                            "SidePlankWithKneeTuck",
	PlankExerciseName_WeightedSidePlankWithKneeTuck:                    "WeightedSidePlankWithKneeTuck",
	PlankExerciseName_SidePlankWithLegLift:                             "SidePlankWithLegLift",
	PlankExerciseName_WeightedSidePlankWithLegLift:                     "WeightedSidePlankWithLegLift",
	PlankExerciseName_SidePlankWithReachUnder:                          "SidePlankWithReachUnder",
	PlankExerciseName_WeightedSidePlankWithReachUnder:                  "WeightedSidePlankWithReachUnder",
	PlankExerciseName_SingleLegElevatedFeetPlank:                       "SingleLegElevatedFeetPlank",
	PlankExerciseName_WeightedSingleLegElevatedFeetPlank:               "WeightedSingleLegElevatedFeetPlank",
	PlankExerciseName_SingleLegFlexAndExtend:                           "SingleLegFlexAndExtend",
	PlankExerciseName_WeightedSingleLegFlexAndExtend:                   "WeightedSingleLegFlexAndExtend",
	PlankExerciseName_SingleLegSidePlank:                               "SingleLegSidePlank",
	PlankExerciseName_WeightedSingleLegSidePlank:                       "WeightedSingleLegSidePlank",
	PlankExerciseName_SpidermanPlank:                                   "SpidermanPlank",
	PlankExerciseName_WeightedSpidermanPlank:                           "WeightedSpidermanPlank",
	PlankExerciseName_StraightArmPlank:                                 "StraightArmPlank",
	PlankExerciseName_WeightedStraightArmPlank:                         "WeightedStraightArmPlank",
	PlankExerciseName_StraightArmPlankWithShoulderTouch:                "StraightArmPlankWithShoulderTouch",
	PlankExerciseName_WeightedStraightArmPlankWithShoulderTouch:        "WeightedStraightArmPlankWithShoulderTouch",
	PlankExerciseName_SwissBallPlank:                                   "SwissBallPlank",
	PlankExerciseName_WeightedSwissBallPlank:                           "WeightedSwissBallPlank",
	PlankExerciseName_SwissBallPlankLegLift:                            "SwissBallPlankLegLift",
	PlankExerciseName_WeightedSwissBallPlankLegLift:                    "WeightedSwissBallPlankLegLift",
	PlankExerciseName_SwissBallPlankLegLiftAndHold:                     "SwissBallPlankLegLiftAndHold",
	PlankExerciseName_SwissBallPlankWithFeetOnBench:                    "SwissBallPlankWithFeetOnBench",
	PlankExerciseName_WeightedSwissBallPlankWithFeetOnBench:            "WeightedSwissBallPlankWithFeetOnBench",
	PlankExerciseName_SwissBallProneJackknife:                          "SwissBallProneJackknife",
	PlankExerciseName_WeightedSwissBallProneJackknife:                  "WeightedSwissBallProneJackknife",
	PlankExerciseName_SwissBallSidePlank:                               "SwissBallSidePlank",
	PlankExerciseName_WeightedSwissBallSidePlank:                       "WeightedSwissBallSidePlank",
	PlankExerciseName_ThreeWayPlank:                                    "ThreeWayPlank",
	PlankExerciseName_WeightedThreeWayPlank:                            "WeightedThreeWayPlank",
	PlankExerciseName_TowelPlankAndKneeIn:                              "TowelPlankAndKneeIn",
	PlankExerciseName_WeightedTowelPlankAndKneeIn:                      "WeightedTowelPlankAndKneeIn",
	PlankExerciseName_TStabilization:                                   "TStabilization",
	PlankExerciseName_WeightedTStabilization:                           "WeightedTStabilization",
	PlankExerciseName_TurkishGetUpToSidePlank:                          "TurkishGetUpToSidePlank",
	PlankExerciseName_WeightedTurkishGetUpToSidePlank:                  "WeightedTurkishGetUpToSidePlank",
	PlankExerciseName_TwoPointPlank:                                    "TwoPointPlank",
	PlankExerciseName_WeightedTwoPointPlank:                            "WeightedTwoPointPlank",
	PlankExerciseName_WeightedPlank:                                    "WeightedPlank",
	PlankExerciseName_WideStancePlankWithDiagonalArmLift:               "WideStancePlankWithDiagonalArmLift",
	PlankExerciseName_WeightedWideStancePlankWithDiagonalArmLift:       "WeightedWideStancePlankWithDiagonalArmLift",
	PlankExerciseName_WideStancePlankWithDiagonalLegLift:               "WideStancePlankWithDiagonalLegLift",
	PlankExerciseName_WeightedWideStancePlankWithDiagonalLegLift:       "WeightedWideStancePlankWithDiagonalLegLift",
	PlankExerciseName_WideStancePlankWithLegLift:                       "WideStancePlankWithLegLift",
	PlankExerciseName_WeightedWideStancePlankWithLegLift:               "WeightedWideStancePlankWithLegLift",
	PlankExerciseName_WideStancePlankWithOppositeArmAndLegLift:         "WideStancePlankWithOppositeArmAndLegLift",
	PlankExerciseName_WeightedMountainClimberWithHandsOnBench:          "WeightedMountainClimberWithHandsOnBench",
	PlankExerciseName_WeightedSwissBallPlankLegLiftAndHold:             "WeightedSwissBallPlankLegLiftAndHold",
	PlankExerciseName_WeightedWideStancePlankWithOppositeArmAndLegLift: "WeightedWideStancePlankWithOppositeArmAndLegLift",
	PlankExerciseName_PlankWithFeetOnSwissBall:                         "PlankWithFeetOnSwissBall",
	PlankExerciseName_SidePlankToPlankWithReachUnder:                   "SidePlankToPlankWithReachUnder",
	PlankExerciseName_BridgeWithGluteLowerLift:                         "BridgeWithGluteLowerLift",
	PlankExerciseName_BridgeOneLegBridge:                               "BridgeOneLegBridge",
	PlankExerciseName_PlankWithArmVariations:                           "PlankWithArmVariations",
	PlankExerciseName_PlankWithLegLift:                                 "PlankWithLegLift",
	PlankExerciseName_ReversePlankWithLegPull:                          "ReversePlankWithLegPull",
}

func (k PlankExerciseName) String() string {
	if uint(k) < uint(len(PlankExerciseNameNames)) {
		if v, ok := PlankExerciseNameNames[k]; ok {
			return v
		}
	}
	return "PlankExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var PlankExerciseNameValues = map[string]PlankExerciseName{
	"_45DegreePlank":                              PlankExerciseName__45DegreePlank,
	"Weighted45DegreePlank":                       PlankExerciseName_Weighted45DegreePlank,
	"_90DegreeStaticHold":                         PlankExerciseName__90DegreeStaticHold,
	"Weighted90DegreeStaticHold":                  PlankExerciseName_Weighted90DegreeStaticHold,
	"BearCrawl":                                   PlankExerciseName_BearCrawl,
	"WeightedBearCrawl":                           PlankExerciseName_WeightedBearCrawl,
	"CrossBodyMountainClimber":                    PlankExerciseName_CrossBodyMountainClimber,
	"WeightedCrossBodyMountainClimber":            PlankExerciseName_WeightedCrossBodyMountainClimber,
	"ElbowPlankPikeJacks":                         PlankExerciseName_ElbowPlankPikeJacks,
	"WeightedElbowPlankPikeJacks":                 PlankExerciseName_WeightedElbowPlankPikeJacks,
	"ElevatedFeetPlank":                           PlankExerciseName_ElevatedFeetPlank,
	"WeightedElevatedFeetPlank":                   PlankExerciseName_WeightedElevatedFeetPlank,
	"ElevatorAbs":                                 PlankExerciseName_ElevatorAbs,
	"WeightedElevatorAbs":                         PlankExerciseName_WeightedElevatorAbs,
	"ExtendedPlank":                               PlankExerciseName_ExtendedPlank,
	"WeightedExtendedPlank":                       PlankExerciseName_WeightedExtendedPlank,
	"FullPlankPasseTwist":                         PlankExerciseName_FullPlankPasseTwist,
	"WeightedFullPlankPasseTwist":                 PlankExerciseName_WeightedFullPlankPasseTwist,
	"InchingElbowPlank":                           PlankExerciseName_InchingElbowPlank,
	"WeightedInchingElbowPlank":                   PlankExerciseName_WeightedInchingElbowPlank,
	"InchwormToSidePlank":                         PlankExerciseName_InchwormToSidePlank,
	"WeightedInchwormToSidePlank":                 PlankExerciseName_WeightedInchwormToSidePlank,
	"KneelingPlank":                               PlankExerciseName_KneelingPlank,
	"WeightedKneelingPlank":                       PlankExerciseName_WeightedKneelingPlank,
	"KneelingSidePlankWithLegLift":                PlankExerciseName_KneelingSidePlankWithLegLift,
	"WeightedKneelingSidePlankWithLegLift":        PlankExerciseName_WeightedKneelingSidePlankWithLegLift,
	"LateralRoll":                                 PlankExerciseName_LateralRoll,
	"WeightedLateralRoll":                         PlankExerciseName_WeightedLateralRoll,
	"LyingReversePlank":                           PlankExerciseName_LyingReversePlank,
	"WeightedLyingReversePlank":                   PlankExerciseName_WeightedLyingReversePlank,
	"MedicineBallMountainClimber":                 PlankExerciseName_MedicineBallMountainClimber,
	"WeightedMedicineBallMountainClimber":         PlankExerciseName_WeightedMedicineBallMountainClimber,
	"ModifiedMountainClimberAndExtension":         PlankExerciseName_ModifiedMountainClimberAndExtension,
	"WeightedModifiedMountainClimberAndExtension": PlankExerciseName_WeightedModifiedMountainClimberAndExtension,
	"MountainClimber":                             PlankExerciseName_MountainClimber,
	"WeightedMountainClimber":                     PlankExerciseName_WeightedMountainClimber,
	"MountainClimberOnSlidingDiscs":               PlankExerciseName_MountainClimberOnSlidingDiscs,
	"WeightedMountainClimberOnSlidingDiscs":       PlankExerciseName_WeightedMountainClimberOnSlidingDiscs,
	"MountainClimberWithFeetOnBosuBall":           PlankExerciseName_MountainClimberWithFeetOnBosuBall,
	"WeightedMountainClimberWithFeetOnBosuBall":   PlankExerciseName_WeightedMountainClimberWithFeetOnBosuBall,
	"MountainClimberWithHandsOnBench":             PlankExerciseName_MountainClimberWithHandsOnBench,
	"MountainClimberWithHandsOnSwissBall":         PlankExerciseName_MountainClimberWithHandsOnSwissBall,
	"WeightedMountainClimberWithHandsOnSwissBall": PlankExerciseName_WeightedMountainClimberWithHandsOnSwissBall,
	"Plank":                            PlankExerciseName_Plank,
	"PlankJacksWithFeetOnSlidingDiscs": PlankExerciseName_PlankJacksWithFeetOnSlidingDiscs,
	"WeightedPlankJacksWithFeetOnSlidingDiscs": PlankExerciseName_WeightedPlankJacksWithFeetOnSlidingDiscs,
	"PlankKneeTwist":                                   PlankExerciseName_PlankKneeTwist,
	"WeightedPlankKneeTwist":                           PlankExerciseName_WeightedPlankKneeTwist,
	"PlankPikeJumps":                                   PlankExerciseName_PlankPikeJumps,
	"WeightedPlankPikeJumps":                           PlankExerciseName_WeightedPlankPikeJumps,
	"PlankPikes":                                       PlankExerciseName_PlankPikes,
	"WeightedPlankPikes":                               PlankExerciseName_WeightedPlankPikes,
	"PlankToStandUp":                                   PlankExerciseName_PlankToStandUp,
	"WeightedPlankToStandUp":                           PlankExerciseName_WeightedPlankToStandUp,
	"PlankWithArmRaise":                                PlankExerciseName_PlankWithArmRaise,
	"WeightedPlankWithArmRaise":                        PlankExerciseName_WeightedPlankWithArmRaise,
	"PlankWithKneeToElbow":                             PlankExerciseName_PlankWithKneeToElbow,
	"WeightedPlankWithKneeToElbow":                     PlankExerciseName_WeightedPlankWithKneeToElbow,
	"PlankWithObliqueCrunch":                           PlankExerciseName_PlankWithObliqueCrunch,
	"WeightedPlankWithObliqueCrunch":                   PlankExerciseName_WeightedPlankWithObliqueCrunch,
	"PlyometricSidePlank":                              PlankExerciseName_PlyometricSidePlank,
	"WeightedPlyometricSidePlank":                      PlankExerciseName_WeightedPlyometricSidePlank,
	"RollingSidePlank":                                 PlankExerciseName_RollingSidePlank,
	"WeightedRollingSidePlank":                         PlankExerciseName_WeightedRollingSidePlank,
	"SideKickPlank":                                    PlankExerciseName_SideKickPlank,
	"WeightedSideKickPlank":                            PlankExerciseName_WeightedSideKickPlank,
	"SidePlank":                                        PlankExerciseName_SidePlank,
	"WeightedSidePlank":                                PlankExerciseName_WeightedSidePlank,
	"SidePlankAndRow":                                  PlankExerciseName_SidePlankAndRow,
	"WeightedSidePlankAndRow":                          PlankExerciseName_WeightedSidePlankAndRow,
	"SidePlankLift":                                    PlankExerciseName_SidePlankLift,
	"WeightedSidePlankLift":                            PlankExerciseName_WeightedSidePlankLift,
	"SidePlankWithElbowOnBosuBall":                     PlankExerciseName_SidePlankWithElbowOnBosuBall,
	"WeightedSidePlankWithElbowOnBosuBall":             PlankExerciseName_WeightedSidePlankWithElbowOnBosuBall,
	"SidePlankWithFeetOnBench":                         PlankExerciseName_SidePlankWithFeetOnBench,
	"WeightedSidePlankWithFeetOnBench":                 PlankExerciseName_WeightedSidePlankWithFeetOnBench,
	"SidePlankWithKneeCircle":                          PlankExerciseName_SidePlankWithKneeCircle,
	"WeightedSidePlankWithKneeCircle":                  PlankExerciseName_WeightedSidePlankWithKneeCircle,
	"SidePlankWithKneeTuck":                            PlankExerciseName_SidePlankWithKneeTuck,
	"WeightedSidePlankWithKneeTuck":                    PlankExerciseName_WeightedSidePlankWithKneeTuck,
	"SidePlankWithLegLift":                             PlankExerciseName_SidePlankWithLegLift,
	"WeightedSidePlankWithLegLift":                     PlankExerciseName_WeightedSidePlankWithLegLift,
	"SidePlankWithReachUnder":                          PlankExerciseName_SidePlankWithReachUnder,
	"WeightedSidePlankWithReachUnder":                  PlankExerciseName_WeightedSidePlankWithReachUnder,
	"SingleLegElevatedFeetPlank":                       PlankExerciseName_SingleLegElevatedFeetPlank,
	"WeightedSingleLegElevatedFeetPlank":               PlankExerciseName_WeightedSingleLegElevatedFeetPlank,
	"SingleLegFlexAndExtend":                           PlankExerciseName_SingleLegFlexAndExtend,
	"WeightedSingleLegFlexAndExtend":                   PlankExerciseName_WeightedSingleLegFlexAndExtend,
	"SingleLegSidePlank":                               PlankExerciseName_SingleLegSidePlank,
	"WeightedSingleLegSidePlank":                       PlankExerciseName_WeightedSingleLegSidePlank,
	"SpidermanPlank":                                   PlankExerciseName_SpidermanPlank,
	"WeightedSpidermanPlank":                           PlankExerciseName_WeightedSpidermanPlank,
	"StraightArmPlank":                                 PlankExerciseName_StraightArmPlank,
	"WeightedStraightArmPlank":                         PlankExerciseName_WeightedStraightArmPlank,
	"StraightArmPlankWithShoulderTouch":                PlankExerciseName_StraightArmPlankWithShoulderTouch,
	"WeightedStraightArmPlankWithShoulderTouch":        PlankExerciseName_WeightedStraightArmPlankWithShoulderTouch,
	"SwissBallPlank":                                   PlankExerciseName_SwissBallPlank,
	"WeightedSwissBallPlank":                           PlankExerciseName_WeightedSwissBallPlank,
	"SwissBallPlankLegLift":                            PlankExerciseName_SwissBallPlankLegLift,
	"WeightedSwissBallPlankLegLift":                    PlankExerciseName_WeightedSwissBallPlankLegLift,
	"SwissBallPlankLegLiftAndHold":                     PlankExerciseName_SwissBallPlankLegLiftAndHold,
	"SwissBallPlankWithFeetOnBench":                    PlankExerciseName_SwissBallPlankWithFeetOnBench,
	"WeightedSwissBallPlankWithFeetOnBench":            PlankExerciseName_WeightedSwissBallPlankWithFeetOnBench,
	"SwissBallProneJackknife":                          PlankExerciseName_SwissBallProneJackknife,
	"WeightedSwissBallProneJackknife":                  PlankExerciseName_WeightedSwissBallProneJackknife,
	"SwissBallSidePlank":                               PlankExerciseName_SwissBallSidePlank,
	"WeightedSwissBallSidePlank":                       PlankExerciseName_WeightedSwissBallSidePlank,
	"ThreeWayPlank":                                    PlankExerciseName_ThreeWayPlank,
	"WeightedThreeWayPlank":                            PlankExerciseName_WeightedThreeWayPlank,
	"TowelPlankAndKneeIn":                              PlankExerciseName_TowelPlankAndKneeIn,
	"WeightedTowelPlankAndKneeIn":                      PlankExerciseName_WeightedTowelPlankAndKneeIn,
	"TStabilization":                                   PlankExerciseName_TStabilization,
	"WeightedTStabilization":                           PlankExerciseName_WeightedTStabilization,
	"TurkishGetUpToSidePlank":                          PlankExerciseName_TurkishGetUpToSidePlank,
	"WeightedTurkishGetUpToSidePlank":                  PlankExerciseName_WeightedTurkishGetUpToSidePlank,
	"TwoPointPlank":                                    PlankExerciseName_TwoPointPlank,
	"WeightedTwoPointPlank":                            PlankExerciseName_WeightedTwoPointPlank,
	"WeightedPlank":                                    PlankExerciseName_WeightedPlank,
	"WideStancePlankWithDiagonalArmLift":               PlankExerciseName_WideStancePlankWithDiagonalArmLift,
	"WeightedWideStancePlankWithDiagonalArmLift":       PlankExerciseName_WeightedWideStancePlankWithDiagonalArmLift,
	"WideStancePlankWithDiagonalLegLift":               PlankExerciseName_WideStancePlankWithDiagonalLegLift,
	"WeightedWideStancePlankWithDiagonalLegLift":       PlankExerciseName_WeightedWideStancePlankWithDiagonalLegLift,
	"WideStancePlankWithLegLift":                       PlankExerciseName_WideStancePlankWithLegLift,
	"WeightedWideStancePlankWithLegLift":               PlankExerciseName_WeightedWideStancePlankWithLegLift,
	"WideStancePlankWithOppositeArmAndLegLift":         PlankExerciseName_WideStancePlankWithOppositeArmAndLegLift,
	"WeightedMountainClimberWithHandsOnBench":          PlankExerciseName_WeightedMountainClimberWithHandsOnBench,
	"WeightedSwissBallPlankLegLiftAndHold":             PlankExerciseName_WeightedSwissBallPlankLegLiftAndHold,
	"WeightedWideStancePlankWithOppositeArmAndLegLift": PlankExerciseName_WeightedWideStancePlankWithOppositeArmAndLegLift,
	"PlankWithFeetOnSwissBall":                         PlankExerciseName_PlankWithFeetOnSwissBall,
	"SidePlankToPlankWithReachUnder":                   PlankExerciseName_SidePlankToPlankWithReachUnder,
	"BridgeWithGluteLowerLift":                         PlankExerciseName_BridgeWithGluteLowerLift,
	"BridgeOneLegBridge":                               PlankExerciseName_BridgeOneLegBridge,
	"PlankWithArmVariations":                           PlankExerciseName_PlankWithArmVariations,
	"PlankWithLegLift":                                 PlankExerciseName_PlankWithLegLift,
	"ReversePlankWithLegPull":                          PlankExerciseName_ReversePlankWithLegPull,
}

func PlankExerciseNameFromString(s string) PlankExerciseName {
	if v, ok := PlankExerciseNameValues[s]; ok {
		return v
	}
	return PlankExerciseName(PlankExerciseName_Invalid)
}
