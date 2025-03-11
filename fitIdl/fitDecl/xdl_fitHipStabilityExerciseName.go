package fitDecl

import strconv "strconv"

type HipStabilityExerciseName uint16

const (
	HipStabilityExerciseName_BandSideLyingLegRaise             HipStabilityExerciseName = 0
	HipStabilityExerciseName_DeadBug                           HipStabilityExerciseName = 1
	HipStabilityExerciseName_WeightedDeadBug                   HipStabilityExerciseName = 2
	HipStabilityExerciseName_ExternalHipRaise                  HipStabilityExerciseName = 3
	HipStabilityExerciseName_WeightedExternalHipRaise          HipStabilityExerciseName = 4
	HipStabilityExerciseName_FireHydrantKicks                  HipStabilityExerciseName = 5
	HipStabilityExerciseName_WeightedFireHydrantKicks          HipStabilityExerciseName = 6
	HipStabilityExerciseName_HipCircles                        HipStabilityExerciseName = 7
	HipStabilityExerciseName_WeightedHipCircles                HipStabilityExerciseName = 8
	HipStabilityExerciseName_InnerThighLift                    HipStabilityExerciseName = 9
	HipStabilityExerciseName_WeightedInnerThighLift            HipStabilityExerciseName = 10
	HipStabilityExerciseName_LateralWalksWithBandAtAnkles      HipStabilityExerciseName = 11
	HipStabilityExerciseName_PretzelSideKick                   HipStabilityExerciseName = 12
	HipStabilityExerciseName_WeightedPretzelSideKick           HipStabilityExerciseName = 13
	HipStabilityExerciseName_ProneHipInternalRotation          HipStabilityExerciseName = 14
	HipStabilityExerciseName_WeightedProneHipInternalRotation  HipStabilityExerciseName = 15
	HipStabilityExerciseName_Quadruped                         HipStabilityExerciseName = 16
	HipStabilityExerciseName_QuadrupedHipExtension             HipStabilityExerciseName = 17
	HipStabilityExerciseName_WeightedQuadrupedHipExtension     HipStabilityExerciseName = 18
	HipStabilityExerciseName_QuadrupedWithLegLift              HipStabilityExerciseName = 19
	HipStabilityExerciseName_WeightedQuadrupedWithLegLift      HipStabilityExerciseName = 20
	HipStabilityExerciseName_SideLyingLegRaise                 HipStabilityExerciseName = 21
	HipStabilityExerciseName_WeightedSideLyingLegRaise         HipStabilityExerciseName = 22
	HipStabilityExerciseName_SlidingHipAdduction               HipStabilityExerciseName = 23
	HipStabilityExerciseName_WeightedSlidingHipAdduction       HipStabilityExerciseName = 24
	HipStabilityExerciseName_StandingAdduction                 HipStabilityExerciseName = 25
	HipStabilityExerciseName_WeightedStandingAdduction         HipStabilityExerciseName = 26
	HipStabilityExerciseName_StandingCableHipAbduction         HipStabilityExerciseName = 27
	HipStabilityExerciseName_StandingHipAbduction              HipStabilityExerciseName = 28
	HipStabilityExerciseName_WeightedStandingHipAbduction      HipStabilityExerciseName = 29
	HipStabilityExerciseName_StandingRearLegRaise              HipStabilityExerciseName = 30
	HipStabilityExerciseName_WeightedStandingRearLegRaise      HipStabilityExerciseName = 31
	HipStabilityExerciseName_SupineHipInternalRotation         HipStabilityExerciseName = 32
	HipStabilityExerciseName_WeightedSupineHipInternalRotation HipStabilityExerciseName = 33
	HipStabilityExerciseName_Invalid                           HipStabilityExerciseName = 65535
)

var HipStabilityExerciseNameNames = map[HipStabilityExerciseName]string{
	HipStabilityExerciseName_BandSideLyingLegRaise:             "BandSideLyingLegRaise",
	HipStabilityExerciseName_DeadBug:                           "DeadBug",
	HipStabilityExerciseName_WeightedDeadBug:                   "WeightedDeadBug",
	HipStabilityExerciseName_ExternalHipRaise:                  "ExternalHipRaise",
	HipStabilityExerciseName_WeightedExternalHipRaise:          "WeightedExternalHipRaise",
	HipStabilityExerciseName_FireHydrantKicks:                  "FireHydrantKicks",
	HipStabilityExerciseName_WeightedFireHydrantKicks:          "WeightedFireHydrantKicks",
	HipStabilityExerciseName_HipCircles:                        "HipCircles",
	HipStabilityExerciseName_WeightedHipCircles:                "WeightedHipCircles",
	HipStabilityExerciseName_InnerThighLift:                    "InnerThighLift",
	HipStabilityExerciseName_WeightedInnerThighLift:            "WeightedInnerThighLift",
	HipStabilityExerciseName_LateralWalksWithBandAtAnkles:      "LateralWalksWithBandAtAnkles",
	HipStabilityExerciseName_PretzelSideKick:                   "PretzelSideKick",
	HipStabilityExerciseName_WeightedPretzelSideKick:           "WeightedPretzelSideKick",
	HipStabilityExerciseName_ProneHipInternalRotation:          "ProneHipInternalRotation",
	HipStabilityExerciseName_WeightedProneHipInternalRotation:  "WeightedProneHipInternalRotation",
	HipStabilityExerciseName_Quadruped:                         "Quadruped",
	HipStabilityExerciseName_QuadrupedHipExtension:             "QuadrupedHipExtension",
	HipStabilityExerciseName_WeightedQuadrupedHipExtension:     "WeightedQuadrupedHipExtension",
	HipStabilityExerciseName_QuadrupedWithLegLift:              "QuadrupedWithLegLift",
	HipStabilityExerciseName_WeightedQuadrupedWithLegLift:      "WeightedQuadrupedWithLegLift",
	HipStabilityExerciseName_SideLyingLegRaise:                 "SideLyingLegRaise",
	HipStabilityExerciseName_WeightedSideLyingLegRaise:         "WeightedSideLyingLegRaise",
	HipStabilityExerciseName_SlidingHipAdduction:               "SlidingHipAdduction",
	HipStabilityExerciseName_WeightedSlidingHipAdduction:       "WeightedSlidingHipAdduction",
	HipStabilityExerciseName_StandingAdduction:                 "StandingAdduction",
	HipStabilityExerciseName_WeightedStandingAdduction:         "WeightedStandingAdduction",
	HipStabilityExerciseName_StandingCableHipAbduction:         "StandingCableHipAbduction",
	HipStabilityExerciseName_StandingHipAbduction:              "StandingHipAbduction",
	HipStabilityExerciseName_WeightedStandingHipAbduction:      "WeightedStandingHipAbduction",
	HipStabilityExerciseName_StandingRearLegRaise:              "StandingRearLegRaise",
	HipStabilityExerciseName_WeightedStandingRearLegRaise:      "WeightedStandingRearLegRaise",
	HipStabilityExerciseName_SupineHipInternalRotation:         "SupineHipInternalRotation",
	HipStabilityExerciseName_WeightedSupineHipInternalRotation: "WeightedSupineHipInternalRotation",
}

func (k HipStabilityExerciseName) String() string {
	if uint(k) < uint(len(HipStabilityExerciseNameNames)) {
		if v, ok := HipStabilityExerciseNameNames[k]; ok {
			return v
		}
	}
	return "HipStabilityExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var HipStabilityExerciseNameValues = map[string]HipStabilityExerciseName{
	"BandSideLyingLegRaise":             HipStabilityExerciseName_BandSideLyingLegRaise,
	"DeadBug":                           HipStabilityExerciseName_DeadBug,
	"WeightedDeadBug":                   HipStabilityExerciseName_WeightedDeadBug,
	"ExternalHipRaise":                  HipStabilityExerciseName_ExternalHipRaise,
	"WeightedExternalHipRaise":          HipStabilityExerciseName_WeightedExternalHipRaise,
	"FireHydrantKicks":                  HipStabilityExerciseName_FireHydrantKicks,
	"WeightedFireHydrantKicks":          HipStabilityExerciseName_WeightedFireHydrantKicks,
	"HipCircles":                        HipStabilityExerciseName_HipCircles,
	"WeightedHipCircles":                HipStabilityExerciseName_WeightedHipCircles,
	"InnerThighLift":                    HipStabilityExerciseName_InnerThighLift,
	"WeightedInnerThighLift":            HipStabilityExerciseName_WeightedInnerThighLift,
	"LateralWalksWithBandAtAnkles":      HipStabilityExerciseName_LateralWalksWithBandAtAnkles,
	"PretzelSideKick":                   HipStabilityExerciseName_PretzelSideKick,
	"WeightedPretzelSideKick":           HipStabilityExerciseName_WeightedPretzelSideKick,
	"ProneHipInternalRotation":          HipStabilityExerciseName_ProneHipInternalRotation,
	"WeightedProneHipInternalRotation":  HipStabilityExerciseName_WeightedProneHipInternalRotation,
	"Quadruped":                         HipStabilityExerciseName_Quadruped,
	"QuadrupedHipExtension":             HipStabilityExerciseName_QuadrupedHipExtension,
	"WeightedQuadrupedHipExtension":     HipStabilityExerciseName_WeightedQuadrupedHipExtension,
	"QuadrupedWithLegLift":              HipStabilityExerciseName_QuadrupedWithLegLift,
	"WeightedQuadrupedWithLegLift":      HipStabilityExerciseName_WeightedQuadrupedWithLegLift,
	"SideLyingLegRaise":                 HipStabilityExerciseName_SideLyingLegRaise,
	"WeightedSideLyingLegRaise":         HipStabilityExerciseName_WeightedSideLyingLegRaise,
	"SlidingHipAdduction":               HipStabilityExerciseName_SlidingHipAdduction,
	"WeightedSlidingHipAdduction":       HipStabilityExerciseName_WeightedSlidingHipAdduction,
	"StandingAdduction":                 HipStabilityExerciseName_StandingAdduction,
	"WeightedStandingAdduction":         HipStabilityExerciseName_WeightedStandingAdduction,
	"StandingCableHipAbduction":         HipStabilityExerciseName_StandingCableHipAbduction,
	"StandingHipAbduction":              HipStabilityExerciseName_StandingHipAbduction,
	"WeightedStandingHipAbduction":      HipStabilityExerciseName_WeightedStandingHipAbduction,
	"StandingRearLegRaise":              HipStabilityExerciseName_StandingRearLegRaise,
	"WeightedStandingRearLegRaise":      HipStabilityExerciseName_WeightedStandingRearLegRaise,
	"SupineHipInternalRotation":         HipStabilityExerciseName_SupineHipInternalRotation,
	"WeightedSupineHipInternalRotation": HipStabilityExerciseName_WeightedSupineHipInternalRotation,
}

func HipStabilityExerciseNameFromString(s string) HipStabilityExerciseName {
	if v, ok := HipStabilityExerciseNameValues[s]; ok {
		return v
	}
	return HipStabilityExerciseName(HipStabilityExerciseName_Invalid)
}
