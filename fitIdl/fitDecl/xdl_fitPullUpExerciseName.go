package fitDecl

import strconv "strconv"

type PullUpExerciseName uint16

const (
	PullUpExerciseName_BandedPullUps                    PullUpExerciseName = 0
	PullUpExerciseName__30DegreeLatPulldown             PullUpExerciseName = 1
	PullUpExerciseName_BandAssistedChinUp               PullUpExerciseName = 2
	PullUpExerciseName_CloseGripChinUp                  PullUpExerciseName = 3
	PullUpExerciseName_WeightedCloseGripChinUp          PullUpExerciseName = 4
	PullUpExerciseName_CloseGripLatPulldown             PullUpExerciseName = 5
	PullUpExerciseName_CrossoverChinUp                  PullUpExerciseName = 6
	PullUpExerciseName_WeightedCrossoverChinUp          PullUpExerciseName = 7
	PullUpExerciseName_EzBarPullover                    PullUpExerciseName = 8
	PullUpExerciseName_HangingHurdle                    PullUpExerciseName = 9
	PullUpExerciseName_WeightedHangingHurdle            PullUpExerciseName = 10
	PullUpExerciseName_KneelingLatPulldown              PullUpExerciseName = 11
	PullUpExerciseName_KneelingUnderhandGripLatPulldown PullUpExerciseName = 12
	PullUpExerciseName_LatPulldown                      PullUpExerciseName = 13
	PullUpExerciseName_MixedGripChinUp                  PullUpExerciseName = 14
	PullUpExerciseName_WeightedMixedGripChinUp          PullUpExerciseName = 15
	PullUpExerciseName_MixedGripPullUp                  PullUpExerciseName = 16
	PullUpExerciseName_WeightedMixedGripPullUp          PullUpExerciseName = 17
	PullUpExerciseName_ReverseGripPulldown              PullUpExerciseName = 18
	PullUpExerciseName_StandingCablePullover            PullUpExerciseName = 19
	PullUpExerciseName_StraightArmPulldown              PullUpExerciseName = 20
	PullUpExerciseName_SwissBallEzBarPullover           PullUpExerciseName = 21
	PullUpExerciseName_TowelPullUp                      PullUpExerciseName = 22
	PullUpExerciseName_WeightedTowelPullUp              PullUpExerciseName = 23
	PullUpExerciseName_WeightedPullUp                   PullUpExerciseName = 24
	PullUpExerciseName_WideGripLatPulldown              PullUpExerciseName = 25
	PullUpExerciseName_WideGripPullUp                   PullUpExerciseName = 26
	PullUpExerciseName_WeightedWideGripPullUp           PullUpExerciseName = 27
	PullUpExerciseName_BurpeePullUp                     PullUpExerciseName = 28
	PullUpExerciseName_WeightedBurpeePullUp             PullUpExerciseName = 29
	PullUpExerciseName_JumpingPullUps                   PullUpExerciseName = 30
	PullUpExerciseName_WeightedJumpingPullUps           PullUpExerciseName = 31
	PullUpExerciseName_KippingPullUp                    PullUpExerciseName = 32
	PullUpExerciseName_WeightedKippingPullUp            PullUpExerciseName = 33
	PullUpExerciseName_LPullUp                          PullUpExerciseName = 34
	PullUpExerciseName_WeightedLPullUp                  PullUpExerciseName = 35
	PullUpExerciseName_SuspendedChinUp                  PullUpExerciseName = 36
	PullUpExerciseName_WeightedSuspendedChinUp          PullUpExerciseName = 37
	PullUpExerciseName_PullUp                           PullUpExerciseName = 38
	PullUpExerciseName_Invalid                          PullUpExerciseName = 65535
)

var PullUpExerciseNameNames = map[PullUpExerciseName]string{
	PullUpExerciseName_BandedPullUps:                    "BandedPullUps",
	PullUpExerciseName__30DegreeLatPulldown:             "_30DegreeLatPulldown",
	PullUpExerciseName_BandAssistedChinUp:               "BandAssistedChinUp",
	PullUpExerciseName_CloseGripChinUp:                  "CloseGripChinUp",
	PullUpExerciseName_WeightedCloseGripChinUp:          "WeightedCloseGripChinUp",
	PullUpExerciseName_CloseGripLatPulldown:             "CloseGripLatPulldown",
	PullUpExerciseName_CrossoverChinUp:                  "CrossoverChinUp",
	PullUpExerciseName_WeightedCrossoverChinUp:          "WeightedCrossoverChinUp",
	PullUpExerciseName_EzBarPullover:                    "EzBarPullover",
	PullUpExerciseName_HangingHurdle:                    "HangingHurdle",
	PullUpExerciseName_WeightedHangingHurdle:            "WeightedHangingHurdle",
	PullUpExerciseName_KneelingLatPulldown:              "KneelingLatPulldown",
	PullUpExerciseName_KneelingUnderhandGripLatPulldown: "KneelingUnderhandGripLatPulldown",
	PullUpExerciseName_LatPulldown:                      "LatPulldown",
	PullUpExerciseName_MixedGripChinUp:                  "MixedGripChinUp",
	PullUpExerciseName_WeightedMixedGripChinUp:          "WeightedMixedGripChinUp",
	PullUpExerciseName_MixedGripPullUp:                  "MixedGripPullUp",
	PullUpExerciseName_WeightedMixedGripPullUp:          "WeightedMixedGripPullUp",
	PullUpExerciseName_ReverseGripPulldown:              "ReverseGripPulldown",
	PullUpExerciseName_StandingCablePullover:            "StandingCablePullover",
	PullUpExerciseName_StraightArmPulldown:              "StraightArmPulldown",
	PullUpExerciseName_SwissBallEzBarPullover:           "SwissBallEzBarPullover",
	PullUpExerciseName_TowelPullUp:                      "TowelPullUp",
	PullUpExerciseName_WeightedTowelPullUp:              "WeightedTowelPullUp",
	PullUpExerciseName_WeightedPullUp:                   "WeightedPullUp",
	PullUpExerciseName_WideGripLatPulldown:              "WideGripLatPulldown",
	PullUpExerciseName_WideGripPullUp:                   "WideGripPullUp",
	PullUpExerciseName_WeightedWideGripPullUp:           "WeightedWideGripPullUp",
	PullUpExerciseName_BurpeePullUp:                     "BurpeePullUp",
	PullUpExerciseName_WeightedBurpeePullUp:             "WeightedBurpeePullUp",
	PullUpExerciseName_JumpingPullUps:                   "JumpingPullUps",
	PullUpExerciseName_WeightedJumpingPullUps:           "WeightedJumpingPullUps",
	PullUpExerciseName_KippingPullUp:                    "KippingPullUp",
	PullUpExerciseName_WeightedKippingPullUp:            "WeightedKippingPullUp",
	PullUpExerciseName_LPullUp:                          "LPullUp",
	PullUpExerciseName_WeightedLPullUp:                  "WeightedLPullUp",
	PullUpExerciseName_SuspendedChinUp:                  "SuspendedChinUp",
	PullUpExerciseName_WeightedSuspendedChinUp:          "WeightedSuspendedChinUp",
	PullUpExerciseName_PullUp:                           "PullUp",
}

func (k PullUpExerciseName) String() string {
	if uint(k) < uint(len(PullUpExerciseNameNames)) {
		if v, ok := PullUpExerciseNameNames[k]; ok {
			return v
		}
	}
	return "PullUpExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var PullUpExerciseNameValues = map[string]PullUpExerciseName{
	"BandedPullUps":                    PullUpExerciseName_BandedPullUps,
	"_30DegreeLatPulldown":             PullUpExerciseName__30DegreeLatPulldown,
	"BandAssistedChinUp":               PullUpExerciseName_BandAssistedChinUp,
	"CloseGripChinUp":                  PullUpExerciseName_CloseGripChinUp,
	"WeightedCloseGripChinUp":          PullUpExerciseName_WeightedCloseGripChinUp,
	"CloseGripLatPulldown":             PullUpExerciseName_CloseGripLatPulldown,
	"CrossoverChinUp":                  PullUpExerciseName_CrossoverChinUp,
	"WeightedCrossoverChinUp":          PullUpExerciseName_WeightedCrossoverChinUp,
	"EzBarPullover":                    PullUpExerciseName_EzBarPullover,
	"HangingHurdle":                    PullUpExerciseName_HangingHurdle,
	"WeightedHangingHurdle":            PullUpExerciseName_WeightedHangingHurdle,
	"KneelingLatPulldown":              PullUpExerciseName_KneelingLatPulldown,
	"KneelingUnderhandGripLatPulldown": PullUpExerciseName_KneelingUnderhandGripLatPulldown,
	"LatPulldown":                      PullUpExerciseName_LatPulldown,
	"MixedGripChinUp":                  PullUpExerciseName_MixedGripChinUp,
	"WeightedMixedGripChinUp":          PullUpExerciseName_WeightedMixedGripChinUp,
	"MixedGripPullUp":                  PullUpExerciseName_MixedGripPullUp,
	"WeightedMixedGripPullUp":          PullUpExerciseName_WeightedMixedGripPullUp,
	"ReverseGripPulldown":              PullUpExerciseName_ReverseGripPulldown,
	"StandingCablePullover":            PullUpExerciseName_StandingCablePullover,
	"StraightArmPulldown":              PullUpExerciseName_StraightArmPulldown,
	"SwissBallEzBarPullover":           PullUpExerciseName_SwissBallEzBarPullover,
	"TowelPullUp":                      PullUpExerciseName_TowelPullUp,
	"WeightedTowelPullUp":              PullUpExerciseName_WeightedTowelPullUp,
	"WeightedPullUp":                   PullUpExerciseName_WeightedPullUp,
	"WideGripLatPulldown":              PullUpExerciseName_WideGripLatPulldown,
	"WideGripPullUp":                   PullUpExerciseName_WideGripPullUp,
	"WeightedWideGripPullUp":           PullUpExerciseName_WeightedWideGripPullUp,
	"BurpeePullUp":                     PullUpExerciseName_BurpeePullUp,
	"WeightedBurpeePullUp":             PullUpExerciseName_WeightedBurpeePullUp,
	"JumpingPullUps":                   PullUpExerciseName_JumpingPullUps,
	"WeightedJumpingPullUps":           PullUpExerciseName_WeightedJumpingPullUps,
	"KippingPullUp":                    PullUpExerciseName_KippingPullUp,
	"WeightedKippingPullUp":            PullUpExerciseName_WeightedKippingPullUp,
	"LPullUp":                          PullUpExerciseName_LPullUp,
	"WeightedLPullUp":                  PullUpExerciseName_WeightedLPullUp,
	"SuspendedChinUp":                  PullUpExerciseName_SuspendedChinUp,
	"WeightedSuspendedChinUp":          PullUpExerciseName_WeightedSuspendedChinUp,
	"PullUp":                           PullUpExerciseName_PullUp,
}

func PullUpExerciseNameFromString(s string) PullUpExerciseName {
	if v, ok := PullUpExerciseNameValues[s]; ok {
		return v
	}
	return PullUpExerciseName(PullUpExerciseName_Invalid)
}
