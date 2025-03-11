package fitDecl

import strconv "strconv"

type LateralRaiseExerciseName uint16

const (
	LateralRaiseExerciseName__45DegreeCableExternalRotation        LateralRaiseExerciseName = 0
	LateralRaiseExerciseName_AlternatingLateralRaiseWithStaticHold LateralRaiseExerciseName = 1
	LateralRaiseExerciseName_BarMuscleUp                           LateralRaiseExerciseName = 2
	LateralRaiseExerciseName_BentOverLateralRaise                  LateralRaiseExerciseName = 3
	LateralRaiseExerciseName_CableDiagonalRaise                    LateralRaiseExerciseName = 4
	LateralRaiseExerciseName_CableFrontRaise                       LateralRaiseExerciseName = 5
	LateralRaiseExerciseName_CalorieRow                            LateralRaiseExerciseName = 6
	LateralRaiseExerciseName_ComboShoulderRaise                    LateralRaiseExerciseName = 7
	LateralRaiseExerciseName_DumbbellDiagonalRaise                 LateralRaiseExerciseName = 8
	LateralRaiseExerciseName_DumbbellVRaise                        LateralRaiseExerciseName = 9
	LateralRaiseExerciseName_FrontRaise                            LateralRaiseExerciseName = 10
	LateralRaiseExerciseName_LeaningDumbbellLateralRaise           LateralRaiseExerciseName = 11
	LateralRaiseExerciseName_LyingDumbbellRaise                    LateralRaiseExerciseName = 12
	LateralRaiseExerciseName_MuscleUp                              LateralRaiseExerciseName = 13
	LateralRaiseExerciseName_OneArmCableLateralRaise               LateralRaiseExerciseName = 14
	LateralRaiseExerciseName_OverhandGripRearLateralRaise          LateralRaiseExerciseName = 15
	LateralRaiseExerciseName_PlateRaises                           LateralRaiseExerciseName = 16
	LateralRaiseExerciseName_RingDip                               LateralRaiseExerciseName = 17
	LateralRaiseExerciseName_WeightedRingDip                       LateralRaiseExerciseName = 18
	LateralRaiseExerciseName_RingMuscleUp                          LateralRaiseExerciseName = 19
	LateralRaiseExerciseName_WeightedRingMuscleUp                  LateralRaiseExerciseName = 20
	LateralRaiseExerciseName_RopeClimb                             LateralRaiseExerciseName = 21
	LateralRaiseExerciseName_WeightedRopeClimb                     LateralRaiseExerciseName = 22
	LateralRaiseExerciseName_Scaption                              LateralRaiseExerciseName = 23
	LateralRaiseExerciseName_SeatedLateralRaise                    LateralRaiseExerciseName = 24
	LateralRaiseExerciseName_SeatedRearLateralRaise                LateralRaiseExerciseName = 25
	LateralRaiseExerciseName_SideLyingLateralRaise                 LateralRaiseExerciseName = 26
	LateralRaiseExerciseName_StandingLift                          LateralRaiseExerciseName = 27
	LateralRaiseExerciseName_SuspendedRow                          LateralRaiseExerciseName = 28
	LateralRaiseExerciseName_UnderhandGripRearLateralRaise         LateralRaiseExerciseName = 29
	LateralRaiseExerciseName_WallSlide                             LateralRaiseExerciseName = 30
	LateralRaiseExerciseName_WeightedWallSlide                     LateralRaiseExerciseName = 31
	LateralRaiseExerciseName_ArmCircles                            LateralRaiseExerciseName = 32
	LateralRaiseExerciseName_ShavingTheHead                        LateralRaiseExerciseName = 33
	LateralRaiseExerciseName_Invalid                               LateralRaiseExerciseName = 65535
)

var LateralRaiseExerciseNameNames = map[LateralRaiseExerciseName]string{
	LateralRaiseExerciseName__45DegreeCableExternalRotation:        "_45DegreeCableExternalRotation",
	LateralRaiseExerciseName_AlternatingLateralRaiseWithStaticHold: "AlternatingLateralRaiseWithStaticHold",
	LateralRaiseExerciseName_BarMuscleUp:                           "BarMuscleUp",
	LateralRaiseExerciseName_BentOverLateralRaise:                  "BentOverLateralRaise",
	LateralRaiseExerciseName_CableDiagonalRaise:                    "CableDiagonalRaise",
	LateralRaiseExerciseName_CableFrontRaise:                       "CableFrontRaise",
	LateralRaiseExerciseName_CalorieRow:                            "CalorieRow",
	LateralRaiseExerciseName_ComboShoulderRaise:                    "ComboShoulderRaise",
	LateralRaiseExerciseName_DumbbellDiagonalRaise:                 "DumbbellDiagonalRaise",
	LateralRaiseExerciseName_DumbbellVRaise:                        "DumbbellVRaise",
	LateralRaiseExerciseName_FrontRaise:                            "FrontRaise",
	LateralRaiseExerciseName_LeaningDumbbellLateralRaise:           "LeaningDumbbellLateralRaise",
	LateralRaiseExerciseName_LyingDumbbellRaise:                    "LyingDumbbellRaise",
	LateralRaiseExerciseName_MuscleUp:                              "MuscleUp",
	LateralRaiseExerciseName_OneArmCableLateralRaise:               "OneArmCableLateralRaise",
	LateralRaiseExerciseName_OverhandGripRearLateralRaise:          "OverhandGripRearLateralRaise",
	LateralRaiseExerciseName_PlateRaises:                           "PlateRaises",
	LateralRaiseExerciseName_RingDip:                               "RingDip",
	LateralRaiseExerciseName_WeightedRingDip:                       "WeightedRingDip",
	LateralRaiseExerciseName_RingMuscleUp:                          "RingMuscleUp",
	LateralRaiseExerciseName_WeightedRingMuscleUp:                  "WeightedRingMuscleUp",
	LateralRaiseExerciseName_RopeClimb:                             "RopeClimb",
	LateralRaiseExerciseName_WeightedRopeClimb:                     "WeightedRopeClimb",
	LateralRaiseExerciseName_Scaption:                              "Scaption",
	LateralRaiseExerciseName_SeatedLateralRaise:                    "SeatedLateralRaise",
	LateralRaiseExerciseName_SeatedRearLateralRaise:                "SeatedRearLateralRaise",
	LateralRaiseExerciseName_SideLyingLateralRaise:                 "SideLyingLateralRaise",
	LateralRaiseExerciseName_StandingLift:                          "StandingLift",
	LateralRaiseExerciseName_SuspendedRow:                          "SuspendedRow",
	LateralRaiseExerciseName_UnderhandGripRearLateralRaise:         "UnderhandGripRearLateralRaise",
	LateralRaiseExerciseName_WallSlide:                             "WallSlide",
	LateralRaiseExerciseName_WeightedWallSlide:                     "WeightedWallSlide",
	LateralRaiseExerciseName_ArmCircles:                            "ArmCircles",
	LateralRaiseExerciseName_ShavingTheHead:                        "ShavingTheHead",
}

func (k LateralRaiseExerciseName) String() string {
	if uint(k) < uint(len(LateralRaiseExerciseNameNames)) {
		if v, ok := LateralRaiseExerciseNameNames[k]; ok {
			return v
		}
	}
	return "LateralRaiseExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var LateralRaiseExerciseNameValues = map[string]LateralRaiseExerciseName{
	"_45DegreeCableExternalRotation":        LateralRaiseExerciseName__45DegreeCableExternalRotation,
	"AlternatingLateralRaiseWithStaticHold": LateralRaiseExerciseName_AlternatingLateralRaiseWithStaticHold,
	"BarMuscleUp":                           LateralRaiseExerciseName_BarMuscleUp,
	"BentOverLateralRaise":                  LateralRaiseExerciseName_BentOverLateralRaise,
	"CableDiagonalRaise":                    LateralRaiseExerciseName_CableDiagonalRaise,
	"CableFrontRaise":                       LateralRaiseExerciseName_CableFrontRaise,
	"CalorieRow":                            LateralRaiseExerciseName_CalorieRow,
	"ComboShoulderRaise":                    LateralRaiseExerciseName_ComboShoulderRaise,
	"DumbbellDiagonalRaise":                 LateralRaiseExerciseName_DumbbellDiagonalRaise,
	"DumbbellVRaise":                        LateralRaiseExerciseName_DumbbellVRaise,
	"FrontRaise":                            LateralRaiseExerciseName_FrontRaise,
	"LeaningDumbbellLateralRaise":           LateralRaiseExerciseName_LeaningDumbbellLateralRaise,
	"LyingDumbbellRaise":                    LateralRaiseExerciseName_LyingDumbbellRaise,
	"MuscleUp":                              LateralRaiseExerciseName_MuscleUp,
	"OneArmCableLateralRaise":               LateralRaiseExerciseName_OneArmCableLateralRaise,
	"OverhandGripRearLateralRaise":          LateralRaiseExerciseName_OverhandGripRearLateralRaise,
	"PlateRaises":                           LateralRaiseExerciseName_PlateRaises,
	"RingDip":                               LateralRaiseExerciseName_RingDip,
	"WeightedRingDip":                       LateralRaiseExerciseName_WeightedRingDip,
	"RingMuscleUp":                          LateralRaiseExerciseName_RingMuscleUp,
	"WeightedRingMuscleUp":                  LateralRaiseExerciseName_WeightedRingMuscleUp,
	"RopeClimb":                             LateralRaiseExerciseName_RopeClimb,
	"WeightedRopeClimb":                     LateralRaiseExerciseName_WeightedRopeClimb,
	"Scaption":                              LateralRaiseExerciseName_Scaption,
	"SeatedLateralRaise":                    LateralRaiseExerciseName_SeatedLateralRaise,
	"SeatedRearLateralRaise":                LateralRaiseExerciseName_SeatedRearLateralRaise,
	"SideLyingLateralRaise":                 LateralRaiseExerciseName_SideLyingLateralRaise,
	"StandingLift":                          LateralRaiseExerciseName_StandingLift,
	"SuspendedRow":                          LateralRaiseExerciseName_SuspendedRow,
	"UnderhandGripRearLateralRaise":         LateralRaiseExerciseName_UnderhandGripRearLateralRaise,
	"WallSlide":                             LateralRaiseExerciseName_WallSlide,
	"WeightedWallSlide":                     LateralRaiseExerciseName_WeightedWallSlide,
	"ArmCircles":                            LateralRaiseExerciseName_ArmCircles,
	"ShavingTheHead":                        LateralRaiseExerciseName_ShavingTheHead,
}

func LateralRaiseExerciseNameFromString(s string) LateralRaiseExerciseName {
	if v, ok := LateralRaiseExerciseNameValues[s]; ok {
		return v
	}
	return LateralRaiseExerciseName(LateralRaiseExerciseName_Invalid)
}
