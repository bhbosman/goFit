package fitDecl

import strconv "strconv"

type SitUpExerciseName uint16

const (
	SitUpExerciseName_AlternatingSitUp                    SitUpExerciseName = 0
	SitUpExerciseName_WeightedAlternatingSitUp            SitUpExerciseName = 1
	SitUpExerciseName_BentKneeVUp                         SitUpExerciseName = 2
	SitUpExerciseName_WeightedBentKneeVUp                 SitUpExerciseName = 3
	SitUpExerciseName_ButterflySitUp                      SitUpExerciseName = 4
	SitUpExerciseName_WeightedButterflySitup              SitUpExerciseName = 5
	SitUpExerciseName_CrossPunchRollUp                    SitUpExerciseName = 6
	SitUpExerciseName_WeightedCrossPunchRollUp            SitUpExerciseName = 7
	SitUpExerciseName_CrossedArmsSitUp                    SitUpExerciseName = 8
	SitUpExerciseName_WeightedCrossedArmsSitUp            SitUpExerciseName = 9
	SitUpExerciseName_GetUpSitUp                          SitUpExerciseName = 10
	SitUpExerciseName_WeightedGetUpSitUp                  SitUpExerciseName = 11
	SitUpExerciseName_HoveringSitUp                       SitUpExerciseName = 12
	SitUpExerciseName_WeightedHoveringSitUp               SitUpExerciseName = 13
	SitUpExerciseName_KettlebellSitUp                     SitUpExerciseName = 14
	SitUpExerciseName_MedicineBallAlternatingVUp          SitUpExerciseName = 15
	SitUpExerciseName_MedicineBallSitUp                   SitUpExerciseName = 16
	SitUpExerciseName_MedicineBallVUp                     SitUpExerciseName = 17
	SitUpExerciseName_ModifiedSitUp                       SitUpExerciseName = 18
	SitUpExerciseName_NegativeSitUp                       SitUpExerciseName = 19
	SitUpExerciseName_OneArmFullSitUp                     SitUpExerciseName = 20
	SitUpExerciseName_RecliningCircle                     SitUpExerciseName = 21
	SitUpExerciseName_WeightedRecliningCircle             SitUpExerciseName = 22
	SitUpExerciseName_ReverseCurlUp                       SitUpExerciseName = 23
	SitUpExerciseName_WeightedReverseCurlUp               SitUpExerciseName = 24
	SitUpExerciseName_SingleLegSwissBallJackknife         SitUpExerciseName = 25
	SitUpExerciseName_WeightedSingleLegSwissBallJackknife SitUpExerciseName = 26
	SitUpExerciseName_TheTeaser                           SitUpExerciseName = 27
	SitUpExerciseName_TheTeaserWeighted                   SitUpExerciseName = 28
	SitUpExerciseName_ThreePartRollDown                   SitUpExerciseName = 29
	SitUpExerciseName_WeightedThreePartRollDown           SitUpExerciseName = 30
	SitUpExerciseName_VUp                                 SitUpExerciseName = 31
	SitUpExerciseName_WeightedVUp                         SitUpExerciseName = 32
	SitUpExerciseName_WeightedRussianTwistOnSwissBall     SitUpExerciseName = 33
	SitUpExerciseName_WeightedSitUp                       SitUpExerciseName = 34
	SitUpExerciseName_XAbs                                SitUpExerciseName = 35
	SitUpExerciseName_WeightedXAbs                        SitUpExerciseName = 36
	SitUpExerciseName_SitUp                               SitUpExerciseName = 37
	SitUpExerciseName_Invalid                             SitUpExerciseName = 65535
)

var SitUpExerciseNameNames = map[SitUpExerciseName]string{
	SitUpExerciseName_AlternatingSitUp:                    "AlternatingSitUp",
	SitUpExerciseName_WeightedAlternatingSitUp:            "WeightedAlternatingSitUp",
	SitUpExerciseName_BentKneeVUp:                         "BentKneeVUp",
	SitUpExerciseName_WeightedBentKneeVUp:                 "WeightedBentKneeVUp",
	SitUpExerciseName_ButterflySitUp:                      "ButterflySitUp",
	SitUpExerciseName_WeightedButterflySitup:              "WeightedButterflySitup",
	SitUpExerciseName_CrossPunchRollUp:                    "CrossPunchRollUp",
	SitUpExerciseName_WeightedCrossPunchRollUp:            "WeightedCrossPunchRollUp",
	SitUpExerciseName_CrossedArmsSitUp:                    "CrossedArmsSitUp",
	SitUpExerciseName_WeightedCrossedArmsSitUp:            "WeightedCrossedArmsSitUp",
	SitUpExerciseName_GetUpSitUp:                          "GetUpSitUp",
	SitUpExerciseName_WeightedGetUpSitUp:                  "WeightedGetUpSitUp",
	SitUpExerciseName_HoveringSitUp:                       "HoveringSitUp",
	SitUpExerciseName_WeightedHoveringSitUp:               "WeightedHoveringSitUp",
	SitUpExerciseName_KettlebellSitUp:                     "KettlebellSitUp",
	SitUpExerciseName_MedicineBallAlternatingVUp:          "MedicineBallAlternatingVUp",
	SitUpExerciseName_MedicineBallSitUp:                   "MedicineBallSitUp",
	SitUpExerciseName_MedicineBallVUp:                     "MedicineBallVUp",
	SitUpExerciseName_ModifiedSitUp:                       "ModifiedSitUp",
	SitUpExerciseName_NegativeSitUp:                       "NegativeSitUp",
	SitUpExerciseName_OneArmFullSitUp:                     "OneArmFullSitUp",
	SitUpExerciseName_RecliningCircle:                     "RecliningCircle",
	SitUpExerciseName_WeightedRecliningCircle:             "WeightedRecliningCircle",
	SitUpExerciseName_ReverseCurlUp:                       "ReverseCurlUp",
	SitUpExerciseName_WeightedReverseCurlUp:               "WeightedReverseCurlUp",
	SitUpExerciseName_SingleLegSwissBallJackknife:         "SingleLegSwissBallJackknife",
	SitUpExerciseName_WeightedSingleLegSwissBallJackknife: "WeightedSingleLegSwissBallJackknife",
	SitUpExerciseName_TheTeaser:                           "TheTeaser",
	SitUpExerciseName_TheTeaserWeighted:                   "TheTeaserWeighted",
	SitUpExerciseName_ThreePartRollDown:                   "ThreePartRollDown",
	SitUpExerciseName_WeightedThreePartRollDown:           "WeightedThreePartRollDown",
	SitUpExerciseName_VUp:                                 "VUp",
	SitUpExerciseName_WeightedVUp:                         "WeightedVUp",
	SitUpExerciseName_WeightedRussianTwistOnSwissBall:     "WeightedRussianTwistOnSwissBall",
	SitUpExerciseName_WeightedSitUp:                       "WeightedSitUp",
	SitUpExerciseName_XAbs:                                "XAbs",
	SitUpExerciseName_WeightedXAbs:                        "WeightedXAbs",
	SitUpExerciseName_SitUp:                               "SitUp",
}

func (k SitUpExerciseName) String() string {
	if uint(k) < uint(len(SitUpExerciseNameNames)) {
		if v, ok := SitUpExerciseNameNames[k]; ok {
			return v
		}
	}
	return "SitUpExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var SitUpExerciseNameValues = map[string]SitUpExerciseName{
	"AlternatingSitUp":                    SitUpExerciseName_AlternatingSitUp,
	"WeightedAlternatingSitUp":            SitUpExerciseName_WeightedAlternatingSitUp,
	"BentKneeVUp":                         SitUpExerciseName_BentKneeVUp,
	"WeightedBentKneeVUp":                 SitUpExerciseName_WeightedBentKneeVUp,
	"ButterflySitUp":                      SitUpExerciseName_ButterflySitUp,
	"WeightedButterflySitup":              SitUpExerciseName_WeightedButterflySitup,
	"CrossPunchRollUp":                    SitUpExerciseName_CrossPunchRollUp,
	"WeightedCrossPunchRollUp":            SitUpExerciseName_WeightedCrossPunchRollUp,
	"CrossedArmsSitUp":                    SitUpExerciseName_CrossedArmsSitUp,
	"WeightedCrossedArmsSitUp":            SitUpExerciseName_WeightedCrossedArmsSitUp,
	"GetUpSitUp":                          SitUpExerciseName_GetUpSitUp,
	"WeightedGetUpSitUp":                  SitUpExerciseName_WeightedGetUpSitUp,
	"HoveringSitUp":                       SitUpExerciseName_HoveringSitUp,
	"WeightedHoveringSitUp":               SitUpExerciseName_WeightedHoveringSitUp,
	"KettlebellSitUp":                     SitUpExerciseName_KettlebellSitUp,
	"MedicineBallAlternatingVUp":          SitUpExerciseName_MedicineBallAlternatingVUp,
	"MedicineBallSitUp":                   SitUpExerciseName_MedicineBallSitUp,
	"MedicineBallVUp":                     SitUpExerciseName_MedicineBallVUp,
	"ModifiedSitUp":                       SitUpExerciseName_ModifiedSitUp,
	"NegativeSitUp":                       SitUpExerciseName_NegativeSitUp,
	"OneArmFullSitUp":                     SitUpExerciseName_OneArmFullSitUp,
	"RecliningCircle":                     SitUpExerciseName_RecliningCircle,
	"WeightedRecliningCircle":             SitUpExerciseName_WeightedRecliningCircle,
	"ReverseCurlUp":                       SitUpExerciseName_ReverseCurlUp,
	"WeightedReverseCurlUp":               SitUpExerciseName_WeightedReverseCurlUp,
	"SingleLegSwissBallJackknife":         SitUpExerciseName_SingleLegSwissBallJackknife,
	"WeightedSingleLegSwissBallJackknife": SitUpExerciseName_WeightedSingleLegSwissBallJackknife,
	"TheTeaser":                           SitUpExerciseName_TheTeaser,
	"TheTeaserWeighted":                   SitUpExerciseName_TheTeaserWeighted,
	"ThreePartRollDown":                   SitUpExerciseName_ThreePartRollDown,
	"WeightedThreePartRollDown":           SitUpExerciseName_WeightedThreePartRollDown,
	"VUp":                                 SitUpExerciseName_VUp,
	"WeightedVUp":                         SitUpExerciseName_WeightedVUp,
	"WeightedRussianTwistOnSwissBall":     SitUpExerciseName_WeightedRussianTwistOnSwissBall,
	"WeightedSitUp":                       SitUpExerciseName_WeightedSitUp,
	"XAbs":                                SitUpExerciseName_XAbs,
	"WeightedXAbs":                        SitUpExerciseName_WeightedXAbs,
	"SitUp":                               SitUpExerciseName_SitUp,
}

func SitUpExerciseNameFromString(s string) SitUpExerciseName {
	if v, ok := SitUpExerciseNameValues[s]; ok {
		return v
	}
	return SitUpExerciseName(SitUpExerciseName_Invalid)
}
