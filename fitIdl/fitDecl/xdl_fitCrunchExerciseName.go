package fitDecl

import strconv "strconv"

type CrunchExerciseName uint16

const (
	CrunchExerciseName_BicycleCrunch                           CrunchExerciseName = 0
	CrunchExerciseName_CableCrunch                             CrunchExerciseName = 1
	CrunchExerciseName_CircularArmCrunch                       CrunchExerciseName = 2
	CrunchExerciseName_CrossedArmsCrunch                       CrunchExerciseName = 3
	CrunchExerciseName_WeightedCrossedArmsCrunch               CrunchExerciseName = 4
	CrunchExerciseName_CrossLegReverseCrunch                   CrunchExerciseName = 5
	CrunchExerciseName_WeightedCrossLegReverseCrunch           CrunchExerciseName = 6
	CrunchExerciseName_CrunchChop                              CrunchExerciseName = 7
	CrunchExerciseName_WeightedCrunchChop                      CrunchExerciseName = 8
	CrunchExerciseName_DoubleCrunch                            CrunchExerciseName = 9
	CrunchExerciseName_WeightedDoubleCrunch                    CrunchExerciseName = 10
	CrunchExerciseName_ElbowToKneeCrunch                       CrunchExerciseName = 11
	CrunchExerciseName_WeightedElbowToKneeCrunch               CrunchExerciseName = 12
	CrunchExerciseName_FlutterKicks                            CrunchExerciseName = 13
	CrunchExerciseName_WeightedFlutterKicks                    CrunchExerciseName = 14
	CrunchExerciseName_FoamRollerReverseCrunchOnBench          CrunchExerciseName = 15
	CrunchExerciseName_WeightedFoamRollerReverseCrunchOnBench  CrunchExerciseName = 16
	CrunchExerciseName_FoamRollerReverseCrunchWithDumbbell     CrunchExerciseName = 17
	CrunchExerciseName_FoamRollerReverseCrunchWithMedicineBall CrunchExerciseName = 18
	CrunchExerciseName_FrogPress                               CrunchExerciseName = 19
	CrunchExerciseName_HangingKneeRaiseObliqueCrunch           CrunchExerciseName = 20
	CrunchExerciseName_WeightedHangingKneeRaiseObliqueCrunch   CrunchExerciseName = 21
	CrunchExerciseName_HipCrossover                            CrunchExerciseName = 22
	CrunchExerciseName_WeightedHipCrossover                    CrunchExerciseName = 23
	CrunchExerciseName_HollowRock                              CrunchExerciseName = 24
	CrunchExerciseName_WeightedHollowRock                      CrunchExerciseName = 25
	CrunchExerciseName_InclineReverseCrunch                    CrunchExerciseName = 26
	CrunchExerciseName_WeightedInclineReverseCrunch            CrunchExerciseName = 27
	CrunchExerciseName_KneelingCableCrunch                     CrunchExerciseName = 28
	CrunchExerciseName_KneelingCrossCrunch                     CrunchExerciseName = 29
	CrunchExerciseName_WeightedKneelingCrossCrunch             CrunchExerciseName = 30
	CrunchExerciseName_KneelingObliqueCableCrunch              CrunchExerciseName = 31
	CrunchExerciseName_KneesToElbow                            CrunchExerciseName = 32
	CrunchExerciseName_LegExtensions                           CrunchExerciseName = 33
	CrunchExerciseName_WeightedLegExtensions                   CrunchExerciseName = 34
	CrunchExerciseName_LegLevers                               CrunchExerciseName = 35
	CrunchExerciseName_McgillCurlUp                            CrunchExerciseName = 36
	CrunchExerciseName_WeightedMcgillCurlUp                    CrunchExerciseName = 37
	CrunchExerciseName_ModifiedPilatesRollUpWithBall           CrunchExerciseName = 38
	CrunchExerciseName_WeightedModifiedPilatesRollUpWithBall   CrunchExerciseName = 39
	CrunchExerciseName_PilatesCrunch                           CrunchExerciseName = 40
	CrunchExerciseName_WeightedPilatesCrunch                   CrunchExerciseName = 41
	CrunchExerciseName_PilatesRollUpWithBall                   CrunchExerciseName = 42
	CrunchExerciseName_WeightedPilatesRollUpWithBall           CrunchExerciseName = 43
	CrunchExerciseName_RaisedLegsCrunch                        CrunchExerciseName = 44
	CrunchExerciseName_WeightedRaisedLegsCrunch                CrunchExerciseName = 45
	CrunchExerciseName_ReverseCrunch                           CrunchExerciseName = 46
	CrunchExerciseName_WeightedReverseCrunch                   CrunchExerciseName = 47
	CrunchExerciseName_ReverseCrunchOnABench                   CrunchExerciseName = 48
	CrunchExerciseName_WeightedReverseCrunchOnABench           CrunchExerciseName = 49
	CrunchExerciseName_ReverseCurlAndLift                      CrunchExerciseName = 50
	CrunchExerciseName_WeightedReverseCurlAndLift              CrunchExerciseName = 51
	CrunchExerciseName_RotationalLift                          CrunchExerciseName = 52
	CrunchExerciseName_WeightedRotationalLift                  CrunchExerciseName = 53
	CrunchExerciseName_SeatedAlternatingReverseCrunch          CrunchExerciseName = 54
	CrunchExerciseName_WeightedSeatedAlternatingReverseCrunch  CrunchExerciseName = 55
	CrunchExerciseName_SeatedLegU                              CrunchExerciseName = 56
	CrunchExerciseName_WeightedSeatedLegU                      CrunchExerciseName = 57
	CrunchExerciseName_SideToSideCrunchAndWeave                CrunchExerciseName = 58
	CrunchExerciseName_WeightedSideToSideCrunchAndWeave        CrunchExerciseName = 59
	CrunchExerciseName_SingleLegReverseCrunch                  CrunchExerciseName = 60
	CrunchExerciseName_WeightedSingleLegReverseCrunch          CrunchExerciseName = 61
	CrunchExerciseName_SkaterCrunchCross                       CrunchExerciseName = 62
	CrunchExerciseName_WeightedSkaterCrunchCross               CrunchExerciseName = 63
	CrunchExerciseName_StandingCableCrunch                     CrunchExerciseName = 64
	CrunchExerciseName_StandingSideCrunch                      CrunchExerciseName = 65
	CrunchExerciseName_StepClimb                               CrunchExerciseName = 66
	CrunchExerciseName_WeightedStepClimb                       CrunchExerciseName = 67
	CrunchExerciseName_SwissBallCrunch                         CrunchExerciseName = 68
	CrunchExerciseName_SwissBallReverseCrunch                  CrunchExerciseName = 69
	CrunchExerciseName_WeightedSwissBallReverseCrunch          CrunchExerciseName = 70
	CrunchExerciseName_SwissBallRussianTwist                   CrunchExerciseName = 71
	CrunchExerciseName_WeightedSwissBallRussianTwist           CrunchExerciseName = 72
	CrunchExerciseName_SwissBallSideCrunch                     CrunchExerciseName = 73
	CrunchExerciseName_WeightedSwissBallSideCrunch             CrunchExerciseName = 74
	CrunchExerciseName_ThoracicCrunchesOnFoamRoller            CrunchExerciseName = 75
	CrunchExerciseName_WeightedThoracicCrunchesOnFoamRoller    CrunchExerciseName = 76
	CrunchExerciseName_TricepsCrunch                           CrunchExerciseName = 77
	CrunchExerciseName_WeightedBicycleCrunch                   CrunchExerciseName = 78
	CrunchExerciseName_WeightedCrunch                          CrunchExerciseName = 79
	CrunchExerciseName_WeightedSwissBallCrunch                 CrunchExerciseName = 80
	CrunchExerciseName_ToesToBar                               CrunchExerciseName = 81
	CrunchExerciseName_WeightedToesToBar                       CrunchExerciseName = 82
	CrunchExerciseName_Crunch                                  CrunchExerciseName = 83
	CrunchExerciseName_StraightLegCrunchWithBall               CrunchExerciseName = 84
	CrunchExerciseName_Invalid                                 CrunchExerciseName = 65535
)

var CrunchExerciseNameNames = map[CrunchExerciseName]string{
	CrunchExerciseName_BicycleCrunch:                           "BicycleCrunch",
	CrunchExerciseName_CableCrunch:                             "CableCrunch",
	CrunchExerciseName_CircularArmCrunch:                       "CircularArmCrunch",
	CrunchExerciseName_CrossedArmsCrunch:                       "CrossedArmsCrunch",
	CrunchExerciseName_WeightedCrossedArmsCrunch:               "WeightedCrossedArmsCrunch",
	CrunchExerciseName_CrossLegReverseCrunch:                   "CrossLegReverseCrunch",
	CrunchExerciseName_WeightedCrossLegReverseCrunch:           "WeightedCrossLegReverseCrunch",
	CrunchExerciseName_CrunchChop:                              "CrunchChop",
	CrunchExerciseName_WeightedCrunchChop:                      "WeightedCrunchChop",
	CrunchExerciseName_DoubleCrunch:                            "DoubleCrunch",
	CrunchExerciseName_WeightedDoubleCrunch:                    "WeightedDoubleCrunch",
	CrunchExerciseName_ElbowToKneeCrunch:                       "ElbowToKneeCrunch",
	CrunchExerciseName_WeightedElbowToKneeCrunch:               "WeightedElbowToKneeCrunch",
	CrunchExerciseName_FlutterKicks:                            "FlutterKicks",
	CrunchExerciseName_WeightedFlutterKicks:                    "WeightedFlutterKicks",
	CrunchExerciseName_FoamRollerReverseCrunchOnBench:          "FoamRollerReverseCrunchOnBench",
	CrunchExerciseName_WeightedFoamRollerReverseCrunchOnBench:  "WeightedFoamRollerReverseCrunchOnBench",
	CrunchExerciseName_FoamRollerReverseCrunchWithDumbbell:     "FoamRollerReverseCrunchWithDumbbell",
	CrunchExerciseName_FoamRollerReverseCrunchWithMedicineBall: "FoamRollerReverseCrunchWithMedicineBall",
	CrunchExerciseName_FrogPress:                               "FrogPress",
	CrunchExerciseName_HangingKneeRaiseObliqueCrunch:           "HangingKneeRaiseObliqueCrunch",
	CrunchExerciseName_WeightedHangingKneeRaiseObliqueCrunch:   "WeightedHangingKneeRaiseObliqueCrunch",
	CrunchExerciseName_HipCrossover:                            "HipCrossover",
	CrunchExerciseName_WeightedHipCrossover:                    "WeightedHipCrossover",
	CrunchExerciseName_HollowRock:                              "HollowRock",
	CrunchExerciseName_WeightedHollowRock:                      "WeightedHollowRock",
	CrunchExerciseName_InclineReverseCrunch:                    "InclineReverseCrunch",
	CrunchExerciseName_WeightedInclineReverseCrunch:            "WeightedInclineReverseCrunch",
	CrunchExerciseName_KneelingCableCrunch:                     "KneelingCableCrunch",
	CrunchExerciseName_KneelingCrossCrunch:                     "KneelingCrossCrunch",
	CrunchExerciseName_WeightedKneelingCrossCrunch:             "WeightedKneelingCrossCrunch",
	CrunchExerciseName_KneelingObliqueCableCrunch:              "KneelingObliqueCableCrunch",
	CrunchExerciseName_KneesToElbow:                            "KneesToElbow",
	CrunchExerciseName_LegExtensions:                           "LegExtensions",
	CrunchExerciseName_WeightedLegExtensions:                   "WeightedLegExtensions",
	CrunchExerciseName_LegLevers:                               "LegLevers",
	CrunchExerciseName_McgillCurlUp:                            "McgillCurlUp",
	CrunchExerciseName_WeightedMcgillCurlUp:                    "WeightedMcgillCurlUp",
	CrunchExerciseName_ModifiedPilatesRollUpWithBall:           "ModifiedPilatesRollUpWithBall",
	CrunchExerciseName_WeightedModifiedPilatesRollUpWithBall:   "WeightedModifiedPilatesRollUpWithBall",
	CrunchExerciseName_PilatesCrunch:                           "PilatesCrunch",
	CrunchExerciseName_WeightedPilatesCrunch:                   "WeightedPilatesCrunch",
	CrunchExerciseName_PilatesRollUpWithBall:                   "PilatesRollUpWithBall",
	CrunchExerciseName_WeightedPilatesRollUpWithBall:           "WeightedPilatesRollUpWithBall",
	CrunchExerciseName_RaisedLegsCrunch:                        "RaisedLegsCrunch",
	CrunchExerciseName_WeightedRaisedLegsCrunch:                "WeightedRaisedLegsCrunch",
	CrunchExerciseName_ReverseCrunch:                           "ReverseCrunch",
	CrunchExerciseName_WeightedReverseCrunch:                   "WeightedReverseCrunch",
	CrunchExerciseName_ReverseCrunchOnABench:                   "ReverseCrunchOnABench",
	CrunchExerciseName_WeightedReverseCrunchOnABench:           "WeightedReverseCrunchOnABench",
	CrunchExerciseName_ReverseCurlAndLift:                      "ReverseCurlAndLift",
	CrunchExerciseName_WeightedReverseCurlAndLift:              "WeightedReverseCurlAndLift",
	CrunchExerciseName_RotationalLift:                          "RotationalLift",
	CrunchExerciseName_WeightedRotationalLift:                  "WeightedRotationalLift",
	CrunchExerciseName_SeatedAlternatingReverseCrunch:          "SeatedAlternatingReverseCrunch",
	CrunchExerciseName_WeightedSeatedAlternatingReverseCrunch:  "WeightedSeatedAlternatingReverseCrunch",
	CrunchExerciseName_SeatedLegU:                              "SeatedLegU",
	CrunchExerciseName_WeightedSeatedLegU:                      "WeightedSeatedLegU",
	CrunchExerciseName_SideToSideCrunchAndWeave:                "SideToSideCrunchAndWeave",
	CrunchExerciseName_WeightedSideToSideCrunchAndWeave:        "WeightedSideToSideCrunchAndWeave",
	CrunchExerciseName_SingleLegReverseCrunch:                  "SingleLegReverseCrunch",
	CrunchExerciseName_WeightedSingleLegReverseCrunch:          "WeightedSingleLegReverseCrunch",
	CrunchExerciseName_SkaterCrunchCross:                       "SkaterCrunchCross",
	CrunchExerciseName_WeightedSkaterCrunchCross:               "WeightedSkaterCrunchCross",
	CrunchExerciseName_StandingCableCrunch:                     "StandingCableCrunch",
	CrunchExerciseName_StandingSideCrunch:                      "StandingSideCrunch",
	CrunchExerciseName_StepClimb:                               "StepClimb",
	CrunchExerciseName_WeightedStepClimb:                       "WeightedStepClimb",
	CrunchExerciseName_SwissBallCrunch:                         "SwissBallCrunch",
	CrunchExerciseName_SwissBallReverseCrunch:                  "SwissBallReverseCrunch",
	CrunchExerciseName_WeightedSwissBallReverseCrunch:          "WeightedSwissBallReverseCrunch",
	CrunchExerciseName_SwissBallRussianTwist:                   "SwissBallRussianTwist",
	CrunchExerciseName_WeightedSwissBallRussianTwist:           "WeightedSwissBallRussianTwist",
	CrunchExerciseName_SwissBallSideCrunch:                     "SwissBallSideCrunch",
	CrunchExerciseName_WeightedSwissBallSideCrunch:             "WeightedSwissBallSideCrunch",
	CrunchExerciseName_ThoracicCrunchesOnFoamRoller:            "ThoracicCrunchesOnFoamRoller",
	CrunchExerciseName_WeightedThoracicCrunchesOnFoamRoller:    "WeightedThoracicCrunchesOnFoamRoller",
	CrunchExerciseName_TricepsCrunch:                           "TricepsCrunch",
	CrunchExerciseName_WeightedBicycleCrunch:                   "WeightedBicycleCrunch",
	CrunchExerciseName_WeightedCrunch:                          "WeightedCrunch",
	CrunchExerciseName_WeightedSwissBallCrunch:                 "WeightedSwissBallCrunch",
	CrunchExerciseName_ToesToBar:                               "ToesToBar",
	CrunchExerciseName_WeightedToesToBar:                       "WeightedToesToBar",
	CrunchExerciseName_Crunch:                                  "Crunch",
	CrunchExerciseName_StraightLegCrunchWithBall:               "StraightLegCrunchWithBall",
}

func (k CrunchExerciseName) String() string {
	if uint(k) < uint(len(CrunchExerciseNameNames)) {
		if v, ok := CrunchExerciseNameNames[k]; ok {
			return v
		}
	}
	return "CrunchExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var CrunchExerciseNameValues = map[string]CrunchExerciseName{
	"BicycleCrunch":                           CrunchExerciseName_BicycleCrunch,
	"CableCrunch":                             CrunchExerciseName_CableCrunch,
	"CircularArmCrunch":                       CrunchExerciseName_CircularArmCrunch,
	"CrossedArmsCrunch":                       CrunchExerciseName_CrossedArmsCrunch,
	"WeightedCrossedArmsCrunch":               CrunchExerciseName_WeightedCrossedArmsCrunch,
	"CrossLegReverseCrunch":                   CrunchExerciseName_CrossLegReverseCrunch,
	"WeightedCrossLegReverseCrunch":           CrunchExerciseName_WeightedCrossLegReverseCrunch,
	"CrunchChop":                              CrunchExerciseName_CrunchChop,
	"WeightedCrunchChop":                      CrunchExerciseName_WeightedCrunchChop,
	"DoubleCrunch":                            CrunchExerciseName_DoubleCrunch,
	"WeightedDoubleCrunch":                    CrunchExerciseName_WeightedDoubleCrunch,
	"ElbowToKneeCrunch":                       CrunchExerciseName_ElbowToKneeCrunch,
	"WeightedElbowToKneeCrunch":               CrunchExerciseName_WeightedElbowToKneeCrunch,
	"FlutterKicks":                            CrunchExerciseName_FlutterKicks,
	"WeightedFlutterKicks":                    CrunchExerciseName_WeightedFlutterKicks,
	"FoamRollerReverseCrunchOnBench":          CrunchExerciseName_FoamRollerReverseCrunchOnBench,
	"WeightedFoamRollerReverseCrunchOnBench":  CrunchExerciseName_WeightedFoamRollerReverseCrunchOnBench,
	"FoamRollerReverseCrunchWithDumbbell":     CrunchExerciseName_FoamRollerReverseCrunchWithDumbbell,
	"FoamRollerReverseCrunchWithMedicineBall": CrunchExerciseName_FoamRollerReverseCrunchWithMedicineBall,
	"FrogPress":                               CrunchExerciseName_FrogPress,
	"HangingKneeRaiseObliqueCrunch":           CrunchExerciseName_HangingKneeRaiseObliqueCrunch,
	"WeightedHangingKneeRaiseObliqueCrunch":   CrunchExerciseName_WeightedHangingKneeRaiseObliqueCrunch,
	"HipCrossover":                            CrunchExerciseName_HipCrossover,
	"WeightedHipCrossover":                    CrunchExerciseName_WeightedHipCrossover,
	"HollowRock":                              CrunchExerciseName_HollowRock,
	"WeightedHollowRock":                      CrunchExerciseName_WeightedHollowRock,
	"InclineReverseCrunch":                    CrunchExerciseName_InclineReverseCrunch,
	"WeightedInclineReverseCrunch":            CrunchExerciseName_WeightedInclineReverseCrunch,
	"KneelingCableCrunch":                     CrunchExerciseName_KneelingCableCrunch,
	"KneelingCrossCrunch":                     CrunchExerciseName_KneelingCrossCrunch,
	"WeightedKneelingCrossCrunch":             CrunchExerciseName_WeightedKneelingCrossCrunch,
	"KneelingObliqueCableCrunch":              CrunchExerciseName_KneelingObliqueCableCrunch,
	"KneesToElbow":                            CrunchExerciseName_KneesToElbow,
	"LegExtensions":                           CrunchExerciseName_LegExtensions,
	"WeightedLegExtensions":                   CrunchExerciseName_WeightedLegExtensions,
	"LegLevers":                               CrunchExerciseName_LegLevers,
	"McgillCurlUp":                            CrunchExerciseName_McgillCurlUp,
	"WeightedMcgillCurlUp":                    CrunchExerciseName_WeightedMcgillCurlUp,
	"ModifiedPilatesRollUpWithBall":           CrunchExerciseName_ModifiedPilatesRollUpWithBall,
	"WeightedModifiedPilatesRollUpWithBall":   CrunchExerciseName_WeightedModifiedPilatesRollUpWithBall,
	"PilatesCrunch":                           CrunchExerciseName_PilatesCrunch,
	"WeightedPilatesCrunch":                   CrunchExerciseName_WeightedPilatesCrunch,
	"PilatesRollUpWithBall":                   CrunchExerciseName_PilatesRollUpWithBall,
	"WeightedPilatesRollUpWithBall":           CrunchExerciseName_WeightedPilatesRollUpWithBall,
	"RaisedLegsCrunch":                        CrunchExerciseName_RaisedLegsCrunch,
	"WeightedRaisedLegsCrunch":                CrunchExerciseName_WeightedRaisedLegsCrunch,
	"ReverseCrunch":                           CrunchExerciseName_ReverseCrunch,
	"WeightedReverseCrunch":                   CrunchExerciseName_WeightedReverseCrunch,
	"ReverseCrunchOnABench":                   CrunchExerciseName_ReverseCrunchOnABench,
	"WeightedReverseCrunchOnABench":           CrunchExerciseName_WeightedReverseCrunchOnABench,
	"ReverseCurlAndLift":                      CrunchExerciseName_ReverseCurlAndLift,
	"WeightedReverseCurlAndLift":              CrunchExerciseName_WeightedReverseCurlAndLift,
	"RotationalLift":                          CrunchExerciseName_RotationalLift,
	"WeightedRotationalLift":                  CrunchExerciseName_WeightedRotationalLift,
	"SeatedAlternatingReverseCrunch":          CrunchExerciseName_SeatedAlternatingReverseCrunch,
	"WeightedSeatedAlternatingReverseCrunch":  CrunchExerciseName_WeightedSeatedAlternatingReverseCrunch,
	"SeatedLegU":                              CrunchExerciseName_SeatedLegU,
	"WeightedSeatedLegU":                      CrunchExerciseName_WeightedSeatedLegU,
	"SideToSideCrunchAndWeave":                CrunchExerciseName_SideToSideCrunchAndWeave,
	"WeightedSideToSideCrunchAndWeave":        CrunchExerciseName_WeightedSideToSideCrunchAndWeave,
	"SingleLegReverseCrunch":                  CrunchExerciseName_SingleLegReverseCrunch,
	"WeightedSingleLegReverseCrunch":          CrunchExerciseName_WeightedSingleLegReverseCrunch,
	"SkaterCrunchCross":                       CrunchExerciseName_SkaterCrunchCross,
	"WeightedSkaterCrunchCross":               CrunchExerciseName_WeightedSkaterCrunchCross,
	"StandingCableCrunch":                     CrunchExerciseName_StandingCableCrunch,
	"StandingSideCrunch":                      CrunchExerciseName_StandingSideCrunch,
	"StepClimb":                               CrunchExerciseName_StepClimb,
	"WeightedStepClimb":                       CrunchExerciseName_WeightedStepClimb,
	"SwissBallCrunch":                         CrunchExerciseName_SwissBallCrunch,
	"SwissBallReverseCrunch":                  CrunchExerciseName_SwissBallReverseCrunch,
	"WeightedSwissBallReverseCrunch":          CrunchExerciseName_WeightedSwissBallReverseCrunch,
	"SwissBallRussianTwist":                   CrunchExerciseName_SwissBallRussianTwist,
	"WeightedSwissBallRussianTwist":           CrunchExerciseName_WeightedSwissBallRussianTwist,
	"SwissBallSideCrunch":                     CrunchExerciseName_SwissBallSideCrunch,
	"WeightedSwissBallSideCrunch":             CrunchExerciseName_WeightedSwissBallSideCrunch,
	"ThoracicCrunchesOnFoamRoller":            CrunchExerciseName_ThoracicCrunchesOnFoamRoller,
	"WeightedThoracicCrunchesOnFoamRoller":    CrunchExerciseName_WeightedThoracicCrunchesOnFoamRoller,
	"TricepsCrunch":                           CrunchExerciseName_TricepsCrunch,
	"WeightedBicycleCrunch":                   CrunchExerciseName_WeightedBicycleCrunch,
	"WeightedCrunch":                          CrunchExerciseName_WeightedCrunch,
	"WeightedSwissBallCrunch":                 CrunchExerciseName_WeightedSwissBallCrunch,
	"ToesToBar":                               CrunchExerciseName_ToesToBar,
	"WeightedToesToBar":                       CrunchExerciseName_WeightedToesToBar,
	"Crunch":                                  CrunchExerciseName_Crunch,
	"StraightLegCrunchWithBall":               CrunchExerciseName_StraightLegCrunchWithBall,
}

func CrunchExerciseNameFromString(s string) CrunchExerciseName {
	if v, ok := CrunchExerciseNameValues[s]; ok {
		return v
	}
	return CrunchExerciseName(CrunchExerciseName_Invalid)
}
