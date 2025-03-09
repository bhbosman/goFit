package fitDecl

import strconv "strconv"

type LungeExerciseName uint16

const (
	LungeExerciseName_OverheadLunge                                 LungeExerciseName = 0
	LungeExerciseName_LungeMatrix                                   LungeExerciseName = 1
	LungeExerciseName_WeightedLungeMatrix                           LungeExerciseName = 2
	LungeExerciseName_AlternatingBarbellForwardLunge                LungeExerciseName = 3
	LungeExerciseName_AlternatingDumbbellLungeWithReach             LungeExerciseName = 4
	LungeExerciseName_BackFootElevatedDumbbellSplitSquat            LungeExerciseName = 5
	LungeExerciseName_BarbellBoxLunge                               LungeExerciseName = 6
	LungeExerciseName_BarbellBulgarianSplitSquat                    LungeExerciseName = 7
	LungeExerciseName_BarbellCrossoverLunge                         LungeExerciseName = 8
	LungeExerciseName_BarbellFrontSplitSquat                        LungeExerciseName = 9
	LungeExerciseName_BarbellLunge                                  LungeExerciseName = 10
	LungeExerciseName_BarbellReverseLunge                           LungeExerciseName = 11
	LungeExerciseName_BarbellSideLunge                              LungeExerciseName = 12
	LungeExerciseName_BarbellSplitSquat                             LungeExerciseName = 13
	LungeExerciseName_CoreControlRearLunge                          LungeExerciseName = 14
	LungeExerciseName_DiagonalLunge                                 LungeExerciseName = 15
	LungeExerciseName_DropLunge                                     LungeExerciseName = 16
	LungeExerciseName_DumbbellBoxLunge                              LungeExerciseName = 17
	LungeExerciseName_DumbbellBulgarianSplitSquat                   LungeExerciseName = 18
	LungeExerciseName_DumbbellCrossoverLunge                        LungeExerciseName = 19
	LungeExerciseName_DumbbellDiagonalLunge                         LungeExerciseName = 20
	LungeExerciseName_DumbbellLunge                                 LungeExerciseName = 21
	LungeExerciseName_DumbbellLungeAndRotation                      LungeExerciseName = 22
	LungeExerciseName_DumbbellOverheadBulgarianSplitSquat           LungeExerciseName = 23
	LungeExerciseName_DumbbellReverseLungeToHighKneeAndPress        LungeExerciseName = 24
	LungeExerciseName_DumbbellSideLunge                             LungeExerciseName = 25
	LungeExerciseName_ElevatedFrontFootBarbellSplitSquat            LungeExerciseName = 26
	LungeExerciseName_FrontFootElevatedDumbbellSplitSquat           LungeExerciseName = 27
	LungeExerciseName_GunslingerLunge                               LungeExerciseName = 28
	LungeExerciseName_LawnmowerLunge                                LungeExerciseName = 29
	LungeExerciseName_LowLungeWithIsometricAdduction                LungeExerciseName = 30
	LungeExerciseName_LowSideToSideLunge                            LungeExerciseName = 31
	LungeExerciseName_Lunge                                         LungeExerciseName = 32
	LungeExerciseName_WeightedLunge                                 LungeExerciseName = 33
	LungeExerciseName_LungeWithArmReach                             LungeExerciseName = 34
	LungeExerciseName_LungeWithDiagonalReach                        LungeExerciseName = 35
	LungeExerciseName_LungeWithSideBend                             LungeExerciseName = 36
	LungeExerciseName_OffsetDumbbellLunge                           LungeExerciseName = 37
	LungeExerciseName_OffsetDumbbellReverseLunge                    LungeExerciseName = 38
	LungeExerciseName_OverheadBulgarianSplitSquat                   LungeExerciseName = 39
	LungeExerciseName_OverheadDumbbellReverseLunge                  LungeExerciseName = 40
	LungeExerciseName_OverheadDumbbellSplitSquat                    LungeExerciseName = 41
	LungeExerciseName_OverheadLungeWithRotation                     LungeExerciseName = 42
	LungeExerciseName_ReverseBarbellBoxLunge                        LungeExerciseName = 43
	LungeExerciseName_ReverseBoxLunge                               LungeExerciseName = 44
	LungeExerciseName_ReverseDumbbellBoxLunge                       LungeExerciseName = 45
	LungeExerciseName_ReverseDumbbellCrossoverLunge                 LungeExerciseName = 46
	LungeExerciseName_ReverseDumbbellDiagonalLunge                  LungeExerciseName = 47
	LungeExerciseName_ReverseLungeWithReachBack                     LungeExerciseName = 48
	LungeExerciseName_WeightedReverseLungeWithReachBack             LungeExerciseName = 49
	LungeExerciseName_ReverseLungeWithTwistAndOverheadReach         LungeExerciseName = 50
	LungeExerciseName_WeightedReverseLungeWithTwistAndOverheadReach LungeExerciseName = 51
	LungeExerciseName_ReverseSlidingBoxLunge                        LungeExerciseName = 52
	LungeExerciseName_WeightedReverseSlidingBoxLunge                LungeExerciseName = 53
	LungeExerciseName_ReverseSlidingLunge                           LungeExerciseName = 54
	LungeExerciseName_WeightedReverseSlidingLunge                   LungeExerciseName = 55
	LungeExerciseName_RunnersLungeToBalance                         LungeExerciseName = 56
	LungeExerciseName_WeightedRunnersLungeToBalance                 LungeExerciseName = 57
	LungeExerciseName_ShiftingSideLunge                             LungeExerciseName = 58
	LungeExerciseName_SideAndCrossoverLunge                         LungeExerciseName = 59
	LungeExerciseName_WeightedSideAndCrossoverLunge                 LungeExerciseName = 60
	LungeExerciseName_SideLunge                                     LungeExerciseName = 61
	LungeExerciseName_WeightedSideLunge                             LungeExerciseName = 62
	LungeExerciseName_SideLungeAndPress                             LungeExerciseName = 63
	LungeExerciseName_SideLungeJumpOff                              LungeExerciseName = 64
	LungeExerciseName_SideLungeSweep                                LungeExerciseName = 65
	LungeExerciseName_WeightedSideLungeSweep                        LungeExerciseName = 66
	LungeExerciseName_SideLungeToCrossoverTap                       LungeExerciseName = 67
	LungeExerciseName_WeightedSideLungeToCrossoverTap               LungeExerciseName = 68
	LungeExerciseName_SideToSideLungeChops                          LungeExerciseName = 69
	LungeExerciseName_WeightedSideToSideLungeChops                  LungeExerciseName = 70
	LungeExerciseName_SiffJumpLunge                                 LungeExerciseName = 71
	LungeExerciseName_WeightedSiffJumpLunge                         LungeExerciseName = 72
	LungeExerciseName_SingleArmReverseLungeAndPress                 LungeExerciseName = 73
	LungeExerciseName_SlidingLateralLunge                           LungeExerciseName = 74
	LungeExerciseName_WeightedSlidingLateralLunge                   LungeExerciseName = 75
	LungeExerciseName_WalkingBarbellLunge                           LungeExerciseName = 76
	LungeExerciseName_WalkingDumbbellLunge                          LungeExerciseName = 77
	LungeExerciseName_WalkingLunge                                  LungeExerciseName = 78
	LungeExerciseName_WeightedWalkingLunge                          LungeExerciseName = 79
	LungeExerciseName_WideGripOverheadBarbellSplitSquat             LungeExerciseName = 80
	LungeExerciseName_Invalid                                       LungeExerciseName = 65535
)

var LungeExerciseNameNames = map[LungeExerciseName]string{
	LungeExerciseName_OverheadLunge:                                 "OverheadLunge",
	LungeExerciseName_LungeMatrix:                                   "LungeMatrix",
	LungeExerciseName_WeightedLungeMatrix:                           "WeightedLungeMatrix",
	LungeExerciseName_AlternatingBarbellForwardLunge:                "AlternatingBarbellForwardLunge",
	LungeExerciseName_AlternatingDumbbellLungeWithReach:             "AlternatingDumbbellLungeWithReach",
	LungeExerciseName_BackFootElevatedDumbbellSplitSquat:            "BackFootElevatedDumbbellSplitSquat",
	LungeExerciseName_BarbellBoxLunge:                               "BarbellBoxLunge",
	LungeExerciseName_BarbellBulgarianSplitSquat:                    "BarbellBulgarianSplitSquat",
	LungeExerciseName_BarbellCrossoverLunge:                         "BarbellCrossoverLunge",
	LungeExerciseName_BarbellFrontSplitSquat:                        "BarbellFrontSplitSquat",
	LungeExerciseName_BarbellLunge:                                  "BarbellLunge",
	LungeExerciseName_BarbellReverseLunge:                           "BarbellReverseLunge",
	LungeExerciseName_BarbellSideLunge:                              "BarbellSideLunge",
	LungeExerciseName_BarbellSplitSquat:                             "BarbellSplitSquat",
	LungeExerciseName_CoreControlRearLunge:                          "CoreControlRearLunge",
	LungeExerciseName_DiagonalLunge:                                 "DiagonalLunge",
	LungeExerciseName_DropLunge:                                     "DropLunge",
	LungeExerciseName_DumbbellBoxLunge:                              "DumbbellBoxLunge",
	LungeExerciseName_DumbbellBulgarianSplitSquat:                   "DumbbellBulgarianSplitSquat",
	LungeExerciseName_DumbbellCrossoverLunge:                        "DumbbellCrossoverLunge",
	LungeExerciseName_DumbbellDiagonalLunge:                         "DumbbellDiagonalLunge",
	LungeExerciseName_DumbbellLunge:                                 "DumbbellLunge",
	LungeExerciseName_DumbbellLungeAndRotation:                      "DumbbellLungeAndRotation",
	LungeExerciseName_DumbbellOverheadBulgarianSplitSquat:           "DumbbellOverheadBulgarianSplitSquat",
	LungeExerciseName_DumbbellReverseLungeToHighKneeAndPress:        "DumbbellReverseLungeToHighKneeAndPress",
	LungeExerciseName_DumbbellSideLunge:                             "DumbbellSideLunge",
	LungeExerciseName_ElevatedFrontFootBarbellSplitSquat:            "ElevatedFrontFootBarbellSplitSquat",
	LungeExerciseName_FrontFootElevatedDumbbellSplitSquat:           "FrontFootElevatedDumbbellSplitSquat",
	LungeExerciseName_GunslingerLunge:                               "GunslingerLunge",
	LungeExerciseName_LawnmowerLunge:                                "LawnmowerLunge",
	LungeExerciseName_LowLungeWithIsometricAdduction:                "LowLungeWithIsometricAdduction",
	LungeExerciseName_LowSideToSideLunge:                            "LowSideToSideLunge",
	LungeExerciseName_Lunge:                                         "Lunge",
	LungeExerciseName_WeightedLunge:                                 "WeightedLunge",
	LungeExerciseName_LungeWithArmReach:                             "LungeWithArmReach",
	LungeExerciseName_LungeWithDiagonalReach:                        "LungeWithDiagonalReach",
	LungeExerciseName_LungeWithSideBend:                             "LungeWithSideBend",
	LungeExerciseName_OffsetDumbbellLunge:                           "OffsetDumbbellLunge",
	LungeExerciseName_OffsetDumbbellReverseLunge:                    "OffsetDumbbellReverseLunge",
	LungeExerciseName_OverheadBulgarianSplitSquat:                   "OverheadBulgarianSplitSquat",
	LungeExerciseName_OverheadDumbbellReverseLunge:                  "OverheadDumbbellReverseLunge",
	LungeExerciseName_OverheadDumbbellSplitSquat:                    "OverheadDumbbellSplitSquat",
	LungeExerciseName_OverheadLungeWithRotation:                     "OverheadLungeWithRotation",
	LungeExerciseName_ReverseBarbellBoxLunge:                        "ReverseBarbellBoxLunge",
	LungeExerciseName_ReverseBoxLunge:                               "ReverseBoxLunge",
	LungeExerciseName_ReverseDumbbellBoxLunge:                       "ReverseDumbbellBoxLunge",
	LungeExerciseName_ReverseDumbbellCrossoverLunge:                 "ReverseDumbbellCrossoverLunge",
	LungeExerciseName_ReverseDumbbellDiagonalLunge:                  "ReverseDumbbellDiagonalLunge",
	LungeExerciseName_ReverseLungeWithReachBack:                     "ReverseLungeWithReachBack",
	LungeExerciseName_WeightedReverseLungeWithReachBack:             "WeightedReverseLungeWithReachBack",
	LungeExerciseName_ReverseLungeWithTwistAndOverheadReach:         "ReverseLungeWithTwistAndOverheadReach",
	LungeExerciseName_WeightedReverseLungeWithTwistAndOverheadReach: "WeightedReverseLungeWithTwistAndOverheadReach",
	LungeExerciseName_ReverseSlidingBoxLunge:                        "ReverseSlidingBoxLunge",
	LungeExerciseName_WeightedReverseSlidingBoxLunge:                "WeightedReverseSlidingBoxLunge",
	LungeExerciseName_ReverseSlidingLunge:                           "ReverseSlidingLunge",
	LungeExerciseName_WeightedReverseSlidingLunge:                   "WeightedReverseSlidingLunge",
	LungeExerciseName_RunnersLungeToBalance:                         "RunnersLungeToBalance",
	LungeExerciseName_WeightedRunnersLungeToBalance:                 "WeightedRunnersLungeToBalance",
	LungeExerciseName_ShiftingSideLunge:                             "ShiftingSideLunge",
	LungeExerciseName_SideAndCrossoverLunge:                         "SideAndCrossoverLunge",
	LungeExerciseName_WeightedSideAndCrossoverLunge:                 "WeightedSideAndCrossoverLunge",
	LungeExerciseName_SideLunge:                                     "SideLunge",
	LungeExerciseName_WeightedSideLunge:                             "WeightedSideLunge",
	LungeExerciseName_SideLungeAndPress:                             "SideLungeAndPress",
	LungeExerciseName_SideLungeJumpOff:                              "SideLungeJumpOff",
	LungeExerciseName_SideLungeSweep:                                "SideLungeSweep",
	LungeExerciseName_WeightedSideLungeSweep:                        "WeightedSideLungeSweep",
	LungeExerciseName_SideLungeToCrossoverTap:                       "SideLungeToCrossoverTap",
	LungeExerciseName_WeightedSideLungeToCrossoverTap:               "WeightedSideLungeToCrossoverTap",
	LungeExerciseName_SideToSideLungeChops:                          "SideToSideLungeChops",
	LungeExerciseName_WeightedSideToSideLungeChops:                  "WeightedSideToSideLungeChops",
	LungeExerciseName_SiffJumpLunge:                                 "SiffJumpLunge",
	LungeExerciseName_WeightedSiffJumpLunge:                         "WeightedSiffJumpLunge",
	LungeExerciseName_SingleArmReverseLungeAndPress:                 "SingleArmReverseLungeAndPress",
	LungeExerciseName_SlidingLateralLunge:                           "SlidingLateralLunge",
	LungeExerciseName_WeightedSlidingLateralLunge:                   "WeightedSlidingLateralLunge",
	LungeExerciseName_WalkingBarbellLunge:                           "WalkingBarbellLunge",
	LungeExerciseName_WalkingDumbbellLunge:                          "WalkingDumbbellLunge",
	LungeExerciseName_WalkingLunge:                                  "WalkingLunge",
	LungeExerciseName_WeightedWalkingLunge:                          "WeightedWalkingLunge",
	LungeExerciseName_WideGripOverheadBarbellSplitSquat:             "WideGripOverheadBarbellSplitSquat",
}

func (k LungeExerciseName) String() string {
	if uint(k) < uint(len(LungeExerciseNameNames)) {
		if v, ok := LungeExerciseNameNames[k]; ok {
			return v
		}
	}
	return "LungeExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var LungeExerciseNameValues = map[string]LungeExerciseName{
	"OverheadLunge":                                 LungeExerciseName_OverheadLunge,
	"LungeMatrix":                                   LungeExerciseName_LungeMatrix,
	"WeightedLungeMatrix":                           LungeExerciseName_WeightedLungeMatrix,
	"AlternatingBarbellForwardLunge":                LungeExerciseName_AlternatingBarbellForwardLunge,
	"AlternatingDumbbellLungeWithReach":             LungeExerciseName_AlternatingDumbbellLungeWithReach,
	"BackFootElevatedDumbbellSplitSquat":            LungeExerciseName_BackFootElevatedDumbbellSplitSquat,
	"BarbellBoxLunge":                               LungeExerciseName_BarbellBoxLunge,
	"BarbellBulgarianSplitSquat":                    LungeExerciseName_BarbellBulgarianSplitSquat,
	"BarbellCrossoverLunge":                         LungeExerciseName_BarbellCrossoverLunge,
	"BarbellFrontSplitSquat":                        LungeExerciseName_BarbellFrontSplitSquat,
	"BarbellLunge":                                  LungeExerciseName_BarbellLunge,
	"BarbellReverseLunge":                           LungeExerciseName_BarbellReverseLunge,
	"BarbellSideLunge":                              LungeExerciseName_BarbellSideLunge,
	"BarbellSplitSquat":                             LungeExerciseName_BarbellSplitSquat,
	"CoreControlRearLunge":                          LungeExerciseName_CoreControlRearLunge,
	"DiagonalLunge":                                 LungeExerciseName_DiagonalLunge,
	"DropLunge":                                     LungeExerciseName_DropLunge,
	"DumbbellBoxLunge":                              LungeExerciseName_DumbbellBoxLunge,
	"DumbbellBulgarianSplitSquat":                   LungeExerciseName_DumbbellBulgarianSplitSquat,
	"DumbbellCrossoverLunge":                        LungeExerciseName_DumbbellCrossoverLunge,
	"DumbbellDiagonalLunge":                         LungeExerciseName_DumbbellDiagonalLunge,
	"DumbbellLunge":                                 LungeExerciseName_DumbbellLunge,
	"DumbbellLungeAndRotation":                      LungeExerciseName_DumbbellLungeAndRotation,
	"DumbbellOverheadBulgarianSplitSquat":           LungeExerciseName_DumbbellOverheadBulgarianSplitSquat,
	"DumbbellReverseLungeToHighKneeAndPress":        LungeExerciseName_DumbbellReverseLungeToHighKneeAndPress,
	"DumbbellSideLunge":                             LungeExerciseName_DumbbellSideLunge,
	"ElevatedFrontFootBarbellSplitSquat":            LungeExerciseName_ElevatedFrontFootBarbellSplitSquat,
	"FrontFootElevatedDumbbellSplitSquat":           LungeExerciseName_FrontFootElevatedDumbbellSplitSquat,
	"GunslingerLunge":                               LungeExerciseName_GunslingerLunge,
	"LawnmowerLunge":                                LungeExerciseName_LawnmowerLunge,
	"LowLungeWithIsometricAdduction":                LungeExerciseName_LowLungeWithIsometricAdduction,
	"LowSideToSideLunge":                            LungeExerciseName_LowSideToSideLunge,
	"Lunge":                                         LungeExerciseName_Lunge,
	"WeightedLunge":                                 LungeExerciseName_WeightedLunge,
	"LungeWithArmReach":                             LungeExerciseName_LungeWithArmReach,
	"LungeWithDiagonalReach":                        LungeExerciseName_LungeWithDiagonalReach,
	"LungeWithSideBend":                             LungeExerciseName_LungeWithSideBend,
	"OffsetDumbbellLunge":                           LungeExerciseName_OffsetDumbbellLunge,
	"OffsetDumbbellReverseLunge":                    LungeExerciseName_OffsetDumbbellReverseLunge,
	"OverheadBulgarianSplitSquat":                   LungeExerciseName_OverheadBulgarianSplitSquat,
	"OverheadDumbbellReverseLunge":                  LungeExerciseName_OverheadDumbbellReverseLunge,
	"OverheadDumbbellSplitSquat":                    LungeExerciseName_OverheadDumbbellSplitSquat,
	"OverheadLungeWithRotation":                     LungeExerciseName_OverheadLungeWithRotation,
	"ReverseBarbellBoxLunge":                        LungeExerciseName_ReverseBarbellBoxLunge,
	"ReverseBoxLunge":                               LungeExerciseName_ReverseBoxLunge,
	"ReverseDumbbellBoxLunge":                       LungeExerciseName_ReverseDumbbellBoxLunge,
	"ReverseDumbbellCrossoverLunge":                 LungeExerciseName_ReverseDumbbellCrossoverLunge,
	"ReverseDumbbellDiagonalLunge":                  LungeExerciseName_ReverseDumbbellDiagonalLunge,
	"ReverseLungeWithReachBack":                     LungeExerciseName_ReverseLungeWithReachBack,
	"WeightedReverseLungeWithReachBack":             LungeExerciseName_WeightedReverseLungeWithReachBack,
	"ReverseLungeWithTwistAndOverheadReach":         LungeExerciseName_ReverseLungeWithTwistAndOverheadReach,
	"WeightedReverseLungeWithTwistAndOverheadReach": LungeExerciseName_WeightedReverseLungeWithTwistAndOverheadReach,
	"ReverseSlidingBoxLunge":                        LungeExerciseName_ReverseSlidingBoxLunge,
	"WeightedReverseSlidingBoxLunge":                LungeExerciseName_WeightedReverseSlidingBoxLunge,
	"ReverseSlidingLunge":                           LungeExerciseName_ReverseSlidingLunge,
	"WeightedReverseSlidingLunge":                   LungeExerciseName_WeightedReverseSlidingLunge,
	"RunnersLungeToBalance":                         LungeExerciseName_RunnersLungeToBalance,
	"WeightedRunnersLungeToBalance":                 LungeExerciseName_WeightedRunnersLungeToBalance,
	"ShiftingSideLunge":                             LungeExerciseName_ShiftingSideLunge,
	"SideAndCrossoverLunge":                         LungeExerciseName_SideAndCrossoverLunge,
	"WeightedSideAndCrossoverLunge":                 LungeExerciseName_WeightedSideAndCrossoverLunge,
	"SideLunge":                                     LungeExerciseName_SideLunge,
	"WeightedSideLunge":                             LungeExerciseName_WeightedSideLunge,
	"SideLungeAndPress":                             LungeExerciseName_SideLungeAndPress,
	"SideLungeJumpOff":                              LungeExerciseName_SideLungeJumpOff,
	"SideLungeSweep":                                LungeExerciseName_SideLungeSweep,
	"WeightedSideLungeSweep":                        LungeExerciseName_WeightedSideLungeSweep,
	"SideLungeToCrossoverTap":                       LungeExerciseName_SideLungeToCrossoverTap,
	"WeightedSideLungeToCrossoverTap":               LungeExerciseName_WeightedSideLungeToCrossoverTap,
	"SideToSideLungeChops":                          LungeExerciseName_SideToSideLungeChops,
	"WeightedSideToSideLungeChops":                  LungeExerciseName_WeightedSideToSideLungeChops,
	"SiffJumpLunge":                                 LungeExerciseName_SiffJumpLunge,
	"WeightedSiffJumpLunge":                         LungeExerciseName_WeightedSiffJumpLunge,
	"SingleArmReverseLungeAndPress":                 LungeExerciseName_SingleArmReverseLungeAndPress,
	"SlidingLateralLunge":                           LungeExerciseName_SlidingLateralLunge,
	"WeightedSlidingLateralLunge":                   LungeExerciseName_WeightedSlidingLateralLunge,
	"WalkingBarbellLunge":                           LungeExerciseName_WalkingBarbellLunge,
	"WalkingDumbbellLunge":                          LungeExerciseName_WalkingDumbbellLunge,
	"WalkingLunge":                                  LungeExerciseName_WalkingLunge,
	"WeightedWalkingLunge":                          LungeExerciseName_WeightedWalkingLunge,
	"WideGripOverheadBarbellSplitSquat":             LungeExerciseName_WideGripOverheadBarbellSplitSquat,
}

func LungeExerciseNameFromString(s string) LungeExerciseName {
	if v, ok := LungeExerciseNameValues[s]; ok {
		return v
	}
	return LungeExerciseName(LungeExerciseName_Invalid)
}
