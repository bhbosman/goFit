package fitDecl

import strconv "strconv"

type WarmUpExerciseName uint16

const (
	WarmUpExerciseName_QuadrupedRocking            WarmUpExerciseName = 0
	WarmUpExerciseName_NeckTilts                   WarmUpExerciseName = 1
	WarmUpExerciseName_AnkleCircles                WarmUpExerciseName = 2
	WarmUpExerciseName_AnkleDorsiflexionWithBand   WarmUpExerciseName = 3
	WarmUpExerciseName_AnkleInternalRotation       WarmUpExerciseName = 4
	WarmUpExerciseName_ArmCircles                  WarmUpExerciseName = 5
	WarmUpExerciseName_BentOverReachToSky          WarmUpExerciseName = 6
	WarmUpExerciseName_CatCamel                    WarmUpExerciseName = 7
	WarmUpExerciseName_ElbowToFootLunge            WarmUpExerciseName = 8
	WarmUpExerciseName_ForwardAndBackwardLegSwings WarmUpExerciseName = 9
	WarmUpExerciseName_Groiners                    WarmUpExerciseName = 10
	WarmUpExerciseName_InvertedHamstringStretch    WarmUpExerciseName = 11
	WarmUpExerciseName_LateralDuckUnder            WarmUpExerciseName = 12
	WarmUpExerciseName_NeckRotations               WarmUpExerciseName = 13
	WarmUpExerciseName_OppositeArmAndLegBalance    WarmUpExerciseName = 14
	WarmUpExerciseName_ReachRollAndLift            WarmUpExerciseName = 15
	WarmUpExerciseName_ShoulderCircles             WarmUpExerciseName = 17
	WarmUpExerciseName_SideToSideLegSwings         WarmUpExerciseName = 18
	WarmUpExerciseName_SleeperStretch              WarmUpExerciseName = 19
	WarmUpExerciseName_SlideOut                    WarmUpExerciseName = 20
	WarmUpExerciseName_SwissBallHipCrossover       WarmUpExerciseName = 21
	WarmUpExerciseName_SwissBallReachRollAndLift   WarmUpExerciseName = 22
	WarmUpExerciseName_SwissBallWindshieldWipers   WarmUpExerciseName = 23
	WarmUpExerciseName_ThoracicRotation            WarmUpExerciseName = 24
	WarmUpExerciseName_WalkingHighKicks            WarmUpExerciseName = 25
	WarmUpExerciseName_WalkingHighKnees            WarmUpExerciseName = 26
	WarmUpExerciseName_WalkingKneeHugs             WarmUpExerciseName = 27
	WarmUpExerciseName_WalkingLegCradles           WarmUpExerciseName = 28
	WarmUpExerciseName_Walkout                     WarmUpExerciseName = 29
	WarmUpExerciseName_WalkoutFromPushUpPosition   WarmUpExerciseName = 30
	WarmUpExerciseName_Invalid                     WarmUpExerciseName = 65535
)

var WarmUpExerciseNameNames = map[WarmUpExerciseName]string{
	WarmUpExerciseName_QuadrupedRocking:            "QuadrupedRocking",
	WarmUpExerciseName_NeckTilts:                   "NeckTilts",
	WarmUpExerciseName_AnkleCircles:                "AnkleCircles",
	WarmUpExerciseName_AnkleDorsiflexionWithBand:   "AnkleDorsiflexionWithBand",
	WarmUpExerciseName_AnkleInternalRotation:       "AnkleInternalRotation",
	WarmUpExerciseName_ArmCircles:                  "ArmCircles",
	WarmUpExerciseName_BentOverReachToSky:          "BentOverReachToSky",
	WarmUpExerciseName_CatCamel:                    "CatCamel",
	WarmUpExerciseName_ElbowToFootLunge:            "ElbowToFootLunge",
	WarmUpExerciseName_ForwardAndBackwardLegSwings: "ForwardAndBackwardLegSwings",
	WarmUpExerciseName_Groiners:                    "Groiners",
	WarmUpExerciseName_InvertedHamstringStretch:    "InvertedHamstringStretch",
	WarmUpExerciseName_LateralDuckUnder:            "LateralDuckUnder",
	WarmUpExerciseName_NeckRotations:               "NeckRotations",
	WarmUpExerciseName_OppositeArmAndLegBalance:    "OppositeArmAndLegBalance",
	WarmUpExerciseName_ReachRollAndLift:            "ReachRollAndLift",
	WarmUpExerciseName_ShoulderCircles:             "ShoulderCircles",
	WarmUpExerciseName_SideToSideLegSwings:         "SideToSideLegSwings",
	WarmUpExerciseName_SleeperStretch:              "SleeperStretch",
	WarmUpExerciseName_SlideOut:                    "SlideOut",
	WarmUpExerciseName_SwissBallHipCrossover:       "SwissBallHipCrossover",
	WarmUpExerciseName_SwissBallReachRollAndLift:   "SwissBallReachRollAndLift",
	WarmUpExerciseName_SwissBallWindshieldWipers:   "SwissBallWindshieldWipers",
	WarmUpExerciseName_ThoracicRotation:            "ThoracicRotation",
	WarmUpExerciseName_WalkingHighKicks:            "WalkingHighKicks",
	WarmUpExerciseName_WalkingHighKnees:            "WalkingHighKnees",
	WarmUpExerciseName_WalkingKneeHugs:             "WalkingKneeHugs",
	WarmUpExerciseName_WalkingLegCradles:           "WalkingLegCradles",
	WarmUpExerciseName_Walkout:                     "Walkout",
	WarmUpExerciseName_WalkoutFromPushUpPosition:   "WalkoutFromPushUpPosition",
}

func (k WarmUpExerciseName) String() string {
	if uint(k) < uint(len(WarmUpExerciseNameNames)) {
		if v, ok := WarmUpExerciseNameNames[k]; ok {
			return v
		}
	}
	return "WarmUpExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var WarmUpExerciseNameValues = map[string]WarmUpExerciseName{
	"QuadrupedRocking":            WarmUpExerciseName_QuadrupedRocking,
	"NeckTilts":                   WarmUpExerciseName_NeckTilts,
	"AnkleCircles":                WarmUpExerciseName_AnkleCircles,
	"AnkleDorsiflexionWithBand":   WarmUpExerciseName_AnkleDorsiflexionWithBand,
	"AnkleInternalRotation":       WarmUpExerciseName_AnkleInternalRotation,
	"ArmCircles":                  WarmUpExerciseName_ArmCircles,
	"BentOverReachToSky":          WarmUpExerciseName_BentOverReachToSky,
	"CatCamel":                    WarmUpExerciseName_CatCamel,
	"ElbowToFootLunge":            WarmUpExerciseName_ElbowToFootLunge,
	"ForwardAndBackwardLegSwings": WarmUpExerciseName_ForwardAndBackwardLegSwings,
	"Groiners":                    WarmUpExerciseName_Groiners,
	"InvertedHamstringStretch":    WarmUpExerciseName_InvertedHamstringStretch,
	"LateralDuckUnder":            WarmUpExerciseName_LateralDuckUnder,
	"NeckRotations":               WarmUpExerciseName_NeckRotations,
	"OppositeArmAndLegBalance":    WarmUpExerciseName_OppositeArmAndLegBalance,
	"ReachRollAndLift":            WarmUpExerciseName_ReachRollAndLift,
	"ShoulderCircles":             WarmUpExerciseName_ShoulderCircles,
	"SideToSideLegSwings":         WarmUpExerciseName_SideToSideLegSwings,
	"SleeperStretch":              WarmUpExerciseName_SleeperStretch,
	"SlideOut":                    WarmUpExerciseName_SlideOut,
	"SwissBallHipCrossover":       WarmUpExerciseName_SwissBallHipCrossover,
	"SwissBallReachRollAndLift":   WarmUpExerciseName_SwissBallReachRollAndLift,
	"SwissBallWindshieldWipers":   WarmUpExerciseName_SwissBallWindshieldWipers,
	"ThoracicRotation":            WarmUpExerciseName_ThoracicRotation,
	"WalkingHighKicks":            WarmUpExerciseName_WalkingHighKicks,
	"WalkingHighKnees":            WarmUpExerciseName_WalkingHighKnees,
	"WalkingKneeHugs":             WarmUpExerciseName_WalkingKneeHugs,
	"WalkingLegCradles":           WarmUpExerciseName_WalkingLegCradles,
	"Walkout":                     WarmUpExerciseName_Walkout,
	"WalkoutFromPushUpPosition":   WarmUpExerciseName_WalkoutFromPushUpPosition,
}

func WarmUpExerciseNameFromString(s string) WarmUpExerciseName {
	if v, ok := WarmUpExerciseNameValues[s]; ok {
		return v
	}
	return WarmUpExerciseName(WarmUpExerciseName_Invalid)
}
