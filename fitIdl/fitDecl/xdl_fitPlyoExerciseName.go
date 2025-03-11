package fitDecl

import strconv "strconv"

type PlyoExerciseName uint16

const (
	PlyoExerciseName_AlternatingJumpLunge                  PlyoExerciseName = 0
	PlyoExerciseName_WeightedAlternatingJumpLunge          PlyoExerciseName = 1
	PlyoExerciseName_BarbellJumpSquat                      PlyoExerciseName = 2
	PlyoExerciseName_BodyWeightJumpSquat                   PlyoExerciseName = 3
	PlyoExerciseName_WeightedJumpSquat                     PlyoExerciseName = 4
	PlyoExerciseName_CrossKneeStrike                       PlyoExerciseName = 5
	PlyoExerciseName_WeightedCrossKneeStrike               PlyoExerciseName = 6
	PlyoExerciseName_DepthJump                             PlyoExerciseName = 7
	PlyoExerciseName_WeightedDepthJump                     PlyoExerciseName = 8
	PlyoExerciseName_DumbbellJumpSquat                     PlyoExerciseName = 9
	PlyoExerciseName_DumbbellSplitJump                     PlyoExerciseName = 10
	PlyoExerciseName_FrontKneeStrike                       PlyoExerciseName = 11
	PlyoExerciseName_WeightedFrontKneeStrike               PlyoExerciseName = 12
	PlyoExerciseName_HighBoxJump                           PlyoExerciseName = 13
	PlyoExerciseName_WeightedHighBoxJump                   PlyoExerciseName = 14
	PlyoExerciseName_IsometricExplosiveBodyWeightJumpSquat PlyoExerciseName = 15
	PlyoExerciseName_WeightedIsometricExplosiveJumpSquat   PlyoExerciseName = 16
	PlyoExerciseName_LateralLeapAndHop                     PlyoExerciseName = 17
	PlyoExerciseName_WeightedLateralLeapAndHop             PlyoExerciseName = 18
	PlyoExerciseName_LateralPlyoSquats                     PlyoExerciseName = 19
	PlyoExerciseName_WeightedLateralPlyoSquats             PlyoExerciseName = 20
	PlyoExerciseName_LateralSlide                          PlyoExerciseName = 21
	PlyoExerciseName_WeightedLateralSlide                  PlyoExerciseName = 22
	PlyoExerciseName_MedicineBallOverheadThrows            PlyoExerciseName = 23
	PlyoExerciseName_MedicineBallSideThrow                 PlyoExerciseName = 24
	PlyoExerciseName_MedicineBallSlam                      PlyoExerciseName = 25
	PlyoExerciseName_SideToSideMedicineBallThrows          PlyoExerciseName = 26
	PlyoExerciseName_SideToSideShuffleJump                 PlyoExerciseName = 27
	PlyoExerciseName_WeightedSideToSideShuffleJump         PlyoExerciseName = 28
	PlyoExerciseName_SquatJumpOntoBox                      PlyoExerciseName = 29
	PlyoExerciseName_WeightedSquatJumpOntoBox              PlyoExerciseName = 30
	PlyoExerciseName_SquatJumpsInAndOut                    PlyoExerciseName = 31
	PlyoExerciseName_WeightedSquatJumpsInAndOut            PlyoExerciseName = 32
	PlyoExerciseName_Invalid                               PlyoExerciseName = 65535
)

var PlyoExerciseNameNames = map[PlyoExerciseName]string{
	PlyoExerciseName_AlternatingJumpLunge:                  "AlternatingJumpLunge",
	PlyoExerciseName_WeightedAlternatingJumpLunge:          "WeightedAlternatingJumpLunge",
	PlyoExerciseName_BarbellJumpSquat:                      "BarbellJumpSquat",
	PlyoExerciseName_BodyWeightJumpSquat:                   "BodyWeightJumpSquat",
	PlyoExerciseName_WeightedJumpSquat:                     "WeightedJumpSquat",
	PlyoExerciseName_CrossKneeStrike:                       "CrossKneeStrike",
	PlyoExerciseName_WeightedCrossKneeStrike:               "WeightedCrossKneeStrike",
	PlyoExerciseName_DepthJump:                             "DepthJump",
	PlyoExerciseName_WeightedDepthJump:                     "WeightedDepthJump",
	PlyoExerciseName_DumbbellJumpSquat:                     "DumbbellJumpSquat",
	PlyoExerciseName_DumbbellSplitJump:                     "DumbbellSplitJump",
	PlyoExerciseName_FrontKneeStrike:                       "FrontKneeStrike",
	PlyoExerciseName_WeightedFrontKneeStrike:               "WeightedFrontKneeStrike",
	PlyoExerciseName_HighBoxJump:                           "HighBoxJump",
	PlyoExerciseName_WeightedHighBoxJump:                   "WeightedHighBoxJump",
	PlyoExerciseName_IsometricExplosiveBodyWeightJumpSquat: "IsometricExplosiveBodyWeightJumpSquat",
	PlyoExerciseName_WeightedIsometricExplosiveJumpSquat:   "WeightedIsometricExplosiveJumpSquat",
	PlyoExerciseName_LateralLeapAndHop:                     "LateralLeapAndHop",
	PlyoExerciseName_WeightedLateralLeapAndHop:             "WeightedLateralLeapAndHop",
	PlyoExerciseName_LateralPlyoSquats:                     "LateralPlyoSquats",
	PlyoExerciseName_WeightedLateralPlyoSquats:             "WeightedLateralPlyoSquats",
	PlyoExerciseName_LateralSlide:                          "LateralSlide",
	PlyoExerciseName_WeightedLateralSlide:                  "WeightedLateralSlide",
	PlyoExerciseName_MedicineBallOverheadThrows:            "MedicineBallOverheadThrows",
	PlyoExerciseName_MedicineBallSideThrow:                 "MedicineBallSideThrow",
	PlyoExerciseName_MedicineBallSlam:                      "MedicineBallSlam",
	PlyoExerciseName_SideToSideMedicineBallThrows:          "SideToSideMedicineBallThrows",
	PlyoExerciseName_SideToSideShuffleJump:                 "SideToSideShuffleJump",
	PlyoExerciseName_WeightedSideToSideShuffleJump:         "WeightedSideToSideShuffleJump",
	PlyoExerciseName_SquatJumpOntoBox:                      "SquatJumpOntoBox",
	PlyoExerciseName_WeightedSquatJumpOntoBox:              "WeightedSquatJumpOntoBox",
	PlyoExerciseName_SquatJumpsInAndOut:                    "SquatJumpsInAndOut",
	PlyoExerciseName_WeightedSquatJumpsInAndOut:            "WeightedSquatJumpsInAndOut",
}

func (k PlyoExerciseName) String() string {
	if uint(k) < uint(len(PlyoExerciseNameNames)) {
		if v, ok := PlyoExerciseNameNames[k]; ok {
			return v
		}
	}
	return "PlyoExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var PlyoExerciseNameValues = map[string]PlyoExerciseName{
	"AlternatingJumpLunge":                  PlyoExerciseName_AlternatingJumpLunge,
	"WeightedAlternatingJumpLunge":          PlyoExerciseName_WeightedAlternatingJumpLunge,
	"BarbellJumpSquat":                      PlyoExerciseName_BarbellJumpSquat,
	"BodyWeightJumpSquat":                   PlyoExerciseName_BodyWeightJumpSquat,
	"WeightedJumpSquat":                     PlyoExerciseName_WeightedJumpSquat,
	"CrossKneeStrike":                       PlyoExerciseName_CrossKneeStrike,
	"WeightedCrossKneeStrike":               PlyoExerciseName_WeightedCrossKneeStrike,
	"DepthJump":                             PlyoExerciseName_DepthJump,
	"WeightedDepthJump":                     PlyoExerciseName_WeightedDepthJump,
	"DumbbellJumpSquat":                     PlyoExerciseName_DumbbellJumpSquat,
	"DumbbellSplitJump":                     PlyoExerciseName_DumbbellSplitJump,
	"FrontKneeStrike":                       PlyoExerciseName_FrontKneeStrike,
	"WeightedFrontKneeStrike":               PlyoExerciseName_WeightedFrontKneeStrike,
	"HighBoxJump":                           PlyoExerciseName_HighBoxJump,
	"WeightedHighBoxJump":                   PlyoExerciseName_WeightedHighBoxJump,
	"IsometricExplosiveBodyWeightJumpSquat": PlyoExerciseName_IsometricExplosiveBodyWeightJumpSquat,
	"WeightedIsometricExplosiveJumpSquat":   PlyoExerciseName_WeightedIsometricExplosiveJumpSquat,
	"LateralLeapAndHop":                     PlyoExerciseName_LateralLeapAndHop,
	"WeightedLateralLeapAndHop":             PlyoExerciseName_WeightedLateralLeapAndHop,
	"LateralPlyoSquats":                     PlyoExerciseName_LateralPlyoSquats,
	"WeightedLateralPlyoSquats":             PlyoExerciseName_WeightedLateralPlyoSquats,
	"LateralSlide":                          PlyoExerciseName_LateralSlide,
	"WeightedLateralSlide":                  PlyoExerciseName_WeightedLateralSlide,
	"MedicineBallOverheadThrows":            PlyoExerciseName_MedicineBallOverheadThrows,
	"MedicineBallSideThrow":                 PlyoExerciseName_MedicineBallSideThrow,
	"MedicineBallSlam":                      PlyoExerciseName_MedicineBallSlam,
	"SideToSideMedicineBallThrows":          PlyoExerciseName_SideToSideMedicineBallThrows,
	"SideToSideShuffleJump":                 PlyoExerciseName_SideToSideShuffleJump,
	"WeightedSideToSideShuffleJump":         PlyoExerciseName_WeightedSideToSideShuffleJump,
	"SquatJumpOntoBox":                      PlyoExerciseName_SquatJumpOntoBox,
	"WeightedSquatJumpOntoBox":              PlyoExerciseName_WeightedSquatJumpOntoBox,
	"SquatJumpsInAndOut":                    PlyoExerciseName_SquatJumpsInAndOut,
	"WeightedSquatJumpsInAndOut":            PlyoExerciseName_WeightedSquatJumpsInAndOut,
}

func PlyoExerciseNameFromString(s string) PlyoExerciseName {
	if v, ok := PlyoExerciseNameValues[s]; ok {
		return v
	}
	return PlyoExerciseName(PlyoExerciseName_Invalid)
}
