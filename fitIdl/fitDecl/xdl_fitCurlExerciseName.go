package fitDecl

import strconv "strconv"

type CurlExerciseName uint16

const (
	CurlExerciseName_AlternatingDumbbellBicepsCurl             CurlExerciseName = 0
	CurlExerciseName_AlternatingDumbbellBicepsCurlOnSwissBall  CurlExerciseName = 1
	CurlExerciseName_AlternatingInclineDumbbellBicepsCurl      CurlExerciseName = 2
	CurlExerciseName_BarbellBicepsCurl                         CurlExerciseName = 3
	CurlExerciseName_BarbellReverseWristCurl                   CurlExerciseName = 4
	CurlExerciseName_BarbellWristCurl                          CurlExerciseName = 5
	CurlExerciseName_BehindTheBackBarbellReverseWristCurl      CurlExerciseName = 6
	CurlExerciseName_BehindTheBackOneArmCableCurl              CurlExerciseName = 7
	CurlExerciseName_CableBicepsCurl                           CurlExerciseName = 8
	CurlExerciseName_CableHammerCurl                           CurlExerciseName = 9
	CurlExerciseName_CheatingBarbellBicepsCurl                 CurlExerciseName = 10
	CurlExerciseName_CloseGripEzBarBicepsCurl                  CurlExerciseName = 11
	CurlExerciseName_CrossBodyDumbbellHammerCurl               CurlExerciseName = 12
	CurlExerciseName_DeadHangBicepsCurl                        CurlExerciseName = 13
	CurlExerciseName_DeclineHammerCurl                         CurlExerciseName = 14
	CurlExerciseName_DumbbellBicepsCurlWithStaticHold          CurlExerciseName = 15
	CurlExerciseName_DumbbellHammerCurl                        CurlExerciseName = 16
	CurlExerciseName_DumbbellReverseWristCurl                  CurlExerciseName = 17
	CurlExerciseName_DumbbellWristCurl                         CurlExerciseName = 18
	CurlExerciseName_EzBarPreacherCurl                         CurlExerciseName = 19
	CurlExerciseName_ForwardBendBicepsCurl                     CurlExerciseName = 20
	CurlExerciseName_HammerCurlToPress                         CurlExerciseName = 21
	CurlExerciseName_InclineDumbbellBicepsCurl                 CurlExerciseName = 22
	CurlExerciseName_InclineOffsetThumbDumbbellCurl            CurlExerciseName = 23
	CurlExerciseName_KettlebellBicepsCurl                      CurlExerciseName = 24
	CurlExerciseName_LyingConcentrationCableCurl               CurlExerciseName = 25
	CurlExerciseName_OneArmPreacherCurl                        CurlExerciseName = 26
	CurlExerciseName_PlatePinchCurl                            CurlExerciseName = 27
	CurlExerciseName_PreacherCurlWithCable                     CurlExerciseName = 28
	CurlExerciseName_ReverseEzBarCurl                          CurlExerciseName = 29
	CurlExerciseName_ReverseGripWristCurl                      CurlExerciseName = 30
	CurlExerciseName_ReverseGripBarbellBicepsCurl              CurlExerciseName = 31
	CurlExerciseName_SeatedAlternatingDumbbellBicepsCurl       CurlExerciseName = 32
	CurlExerciseName_SeatedDumbbellBicepsCurl                  CurlExerciseName = 33
	CurlExerciseName_SeatedReverseDumbbellCurl                 CurlExerciseName = 34
	CurlExerciseName_SplitStanceOffsetPinkyDumbbellCurl        CurlExerciseName = 35
	CurlExerciseName_StandingAlternatingDumbbellCurls          CurlExerciseName = 36
	CurlExerciseName_StandingDumbbellBicepsCurl                CurlExerciseName = 37
	CurlExerciseName_StandingEzBarBicepsCurl                   CurlExerciseName = 38
	CurlExerciseName_StaticCurl                                CurlExerciseName = 39
	CurlExerciseName_SwissBallDumbbellOverheadTricepsExtension CurlExerciseName = 40
	CurlExerciseName_SwissBallEzBarPreacherCurl                CurlExerciseName = 41
	CurlExerciseName_TwistingStandingDumbbellBicepsCurl        CurlExerciseName = 42
	CurlExerciseName_WideGripEzBarBicepsCurl                   CurlExerciseName = 43
	CurlExerciseName_Invalid                                   CurlExerciseName = 65535
)

var CurlExerciseNameNames = map[CurlExerciseName]string{
	CurlExerciseName_AlternatingDumbbellBicepsCurl:             "AlternatingDumbbellBicepsCurl",
	CurlExerciseName_AlternatingDumbbellBicepsCurlOnSwissBall:  "AlternatingDumbbellBicepsCurlOnSwissBall",
	CurlExerciseName_AlternatingInclineDumbbellBicepsCurl:      "AlternatingInclineDumbbellBicepsCurl",
	CurlExerciseName_BarbellBicepsCurl:                         "BarbellBicepsCurl",
	CurlExerciseName_BarbellReverseWristCurl:                   "BarbellReverseWristCurl",
	CurlExerciseName_BarbellWristCurl:                          "BarbellWristCurl",
	CurlExerciseName_BehindTheBackBarbellReverseWristCurl:      "BehindTheBackBarbellReverseWristCurl",
	CurlExerciseName_BehindTheBackOneArmCableCurl:              "BehindTheBackOneArmCableCurl",
	CurlExerciseName_CableBicepsCurl:                           "CableBicepsCurl",
	CurlExerciseName_CableHammerCurl:                           "CableHammerCurl",
	CurlExerciseName_CheatingBarbellBicepsCurl:                 "CheatingBarbellBicepsCurl",
	CurlExerciseName_CloseGripEzBarBicepsCurl:                  "CloseGripEzBarBicepsCurl",
	CurlExerciseName_CrossBodyDumbbellHammerCurl:               "CrossBodyDumbbellHammerCurl",
	CurlExerciseName_DeadHangBicepsCurl:                        "DeadHangBicepsCurl",
	CurlExerciseName_DeclineHammerCurl:                         "DeclineHammerCurl",
	CurlExerciseName_DumbbellBicepsCurlWithStaticHold:          "DumbbellBicepsCurlWithStaticHold",
	CurlExerciseName_DumbbellHammerCurl:                        "DumbbellHammerCurl",
	CurlExerciseName_DumbbellReverseWristCurl:                  "DumbbellReverseWristCurl",
	CurlExerciseName_DumbbellWristCurl:                         "DumbbellWristCurl",
	CurlExerciseName_EzBarPreacherCurl:                         "EzBarPreacherCurl",
	CurlExerciseName_ForwardBendBicepsCurl:                     "ForwardBendBicepsCurl",
	CurlExerciseName_HammerCurlToPress:                         "HammerCurlToPress",
	CurlExerciseName_InclineDumbbellBicepsCurl:                 "InclineDumbbellBicepsCurl",
	CurlExerciseName_InclineOffsetThumbDumbbellCurl:            "InclineOffsetThumbDumbbellCurl",
	CurlExerciseName_KettlebellBicepsCurl:                      "KettlebellBicepsCurl",
	CurlExerciseName_LyingConcentrationCableCurl:               "LyingConcentrationCableCurl",
	CurlExerciseName_OneArmPreacherCurl:                        "OneArmPreacherCurl",
	CurlExerciseName_PlatePinchCurl:                            "PlatePinchCurl",
	CurlExerciseName_PreacherCurlWithCable:                     "PreacherCurlWithCable",
	CurlExerciseName_ReverseEzBarCurl:                          "ReverseEzBarCurl",
	CurlExerciseName_ReverseGripWristCurl:                      "ReverseGripWristCurl",
	CurlExerciseName_ReverseGripBarbellBicepsCurl:              "ReverseGripBarbellBicepsCurl",
	CurlExerciseName_SeatedAlternatingDumbbellBicepsCurl:       "SeatedAlternatingDumbbellBicepsCurl",
	CurlExerciseName_SeatedDumbbellBicepsCurl:                  "SeatedDumbbellBicepsCurl",
	CurlExerciseName_SeatedReverseDumbbellCurl:                 "SeatedReverseDumbbellCurl",
	CurlExerciseName_SplitStanceOffsetPinkyDumbbellCurl:        "SplitStanceOffsetPinkyDumbbellCurl",
	CurlExerciseName_StandingAlternatingDumbbellCurls:          "StandingAlternatingDumbbellCurls",
	CurlExerciseName_StandingDumbbellBicepsCurl:                "StandingDumbbellBicepsCurl",
	CurlExerciseName_StandingEzBarBicepsCurl:                   "StandingEzBarBicepsCurl",
	CurlExerciseName_StaticCurl:                                "StaticCurl",
	CurlExerciseName_SwissBallDumbbellOverheadTricepsExtension: "SwissBallDumbbellOverheadTricepsExtension",
	CurlExerciseName_SwissBallEzBarPreacherCurl:                "SwissBallEzBarPreacherCurl",
	CurlExerciseName_TwistingStandingDumbbellBicepsCurl:        "TwistingStandingDumbbellBicepsCurl",
	CurlExerciseName_WideGripEzBarBicepsCurl:                   "WideGripEzBarBicepsCurl",
}

func (k CurlExerciseName) String() string {
	if uint(k) < uint(len(CurlExerciseNameNames)) {
		if v, ok := CurlExerciseNameNames[k]; ok {
			return v
		}
	}
	return "CurlExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var CurlExerciseNameValues = map[string]CurlExerciseName{
	"AlternatingDumbbellBicepsCurl":             CurlExerciseName_AlternatingDumbbellBicepsCurl,
	"AlternatingDumbbellBicepsCurlOnSwissBall":  CurlExerciseName_AlternatingDumbbellBicepsCurlOnSwissBall,
	"AlternatingInclineDumbbellBicepsCurl":      CurlExerciseName_AlternatingInclineDumbbellBicepsCurl,
	"BarbellBicepsCurl":                         CurlExerciseName_BarbellBicepsCurl,
	"BarbellReverseWristCurl":                   CurlExerciseName_BarbellReverseWristCurl,
	"BarbellWristCurl":                          CurlExerciseName_BarbellWristCurl,
	"BehindTheBackBarbellReverseWristCurl":      CurlExerciseName_BehindTheBackBarbellReverseWristCurl,
	"BehindTheBackOneArmCableCurl":              CurlExerciseName_BehindTheBackOneArmCableCurl,
	"CableBicepsCurl":                           CurlExerciseName_CableBicepsCurl,
	"CableHammerCurl":                           CurlExerciseName_CableHammerCurl,
	"CheatingBarbellBicepsCurl":                 CurlExerciseName_CheatingBarbellBicepsCurl,
	"CloseGripEzBarBicepsCurl":                  CurlExerciseName_CloseGripEzBarBicepsCurl,
	"CrossBodyDumbbellHammerCurl":               CurlExerciseName_CrossBodyDumbbellHammerCurl,
	"DeadHangBicepsCurl":                        CurlExerciseName_DeadHangBicepsCurl,
	"DeclineHammerCurl":                         CurlExerciseName_DeclineHammerCurl,
	"DumbbellBicepsCurlWithStaticHold":          CurlExerciseName_DumbbellBicepsCurlWithStaticHold,
	"DumbbellHammerCurl":                        CurlExerciseName_DumbbellHammerCurl,
	"DumbbellReverseWristCurl":                  CurlExerciseName_DumbbellReverseWristCurl,
	"DumbbellWristCurl":                         CurlExerciseName_DumbbellWristCurl,
	"EzBarPreacherCurl":                         CurlExerciseName_EzBarPreacherCurl,
	"ForwardBendBicepsCurl":                     CurlExerciseName_ForwardBendBicepsCurl,
	"HammerCurlToPress":                         CurlExerciseName_HammerCurlToPress,
	"InclineDumbbellBicepsCurl":                 CurlExerciseName_InclineDumbbellBicepsCurl,
	"InclineOffsetThumbDumbbellCurl":            CurlExerciseName_InclineOffsetThumbDumbbellCurl,
	"KettlebellBicepsCurl":                      CurlExerciseName_KettlebellBicepsCurl,
	"LyingConcentrationCableCurl":               CurlExerciseName_LyingConcentrationCableCurl,
	"OneArmPreacherCurl":                        CurlExerciseName_OneArmPreacherCurl,
	"PlatePinchCurl":                            CurlExerciseName_PlatePinchCurl,
	"PreacherCurlWithCable":                     CurlExerciseName_PreacherCurlWithCable,
	"ReverseEzBarCurl":                          CurlExerciseName_ReverseEzBarCurl,
	"ReverseGripWristCurl":                      CurlExerciseName_ReverseGripWristCurl,
	"ReverseGripBarbellBicepsCurl":              CurlExerciseName_ReverseGripBarbellBicepsCurl,
	"SeatedAlternatingDumbbellBicepsCurl":       CurlExerciseName_SeatedAlternatingDumbbellBicepsCurl,
	"SeatedDumbbellBicepsCurl":                  CurlExerciseName_SeatedDumbbellBicepsCurl,
	"SeatedReverseDumbbellCurl":                 CurlExerciseName_SeatedReverseDumbbellCurl,
	"SplitStanceOffsetPinkyDumbbellCurl":        CurlExerciseName_SplitStanceOffsetPinkyDumbbellCurl,
	"StandingAlternatingDumbbellCurls":          CurlExerciseName_StandingAlternatingDumbbellCurls,
	"StandingDumbbellBicepsCurl":                CurlExerciseName_StandingDumbbellBicepsCurl,
	"StandingEzBarBicepsCurl":                   CurlExerciseName_StandingEzBarBicepsCurl,
	"StaticCurl":                                CurlExerciseName_StaticCurl,
	"SwissBallDumbbellOverheadTricepsExtension": CurlExerciseName_SwissBallDumbbellOverheadTricepsExtension,
	"SwissBallEzBarPreacherCurl":                CurlExerciseName_SwissBallEzBarPreacherCurl,
	"TwistingStandingDumbbellBicepsCurl":        CurlExerciseName_TwistingStandingDumbbellBicepsCurl,
	"WideGripEzBarBicepsCurl":                   CurlExerciseName_WideGripEzBarBicepsCurl,
}

func CurlExerciseNameFromString(s string) CurlExerciseName {
	if v, ok := CurlExerciseNameValues[s]; ok {
		return v
	}
	return CurlExerciseName(CurlExerciseName_Invalid)
}
