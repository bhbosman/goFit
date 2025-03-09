package fitDecl

import strconv "strconv"

type SplitType byte

const (
	SplitType_AscentSplit      SplitType = 1
	SplitType_DescentSplit     SplitType = 2
	SplitType_IntervalActive   SplitType = 3
	SplitType_IntervalRest     SplitType = 4
	SplitType_IntervalWarmup   SplitType = 5
	SplitType_IntervalCooldown SplitType = 6
	SplitType_IntervalRecovery SplitType = 7
	SplitType_IntervalOther    SplitType = 8
	SplitType_ClimbActive      SplitType = 9
	SplitType_ClimbRest        SplitType = 10
	SplitType_SurfActive       SplitType = 11
	SplitType_RunActive        SplitType = 12
	SplitType_RunRest          SplitType = 13
	SplitType_WorkoutRound     SplitType = 14
	SplitType_RwdRun           SplitType = 17
	SplitType_RwdWalk          SplitType = 18
	SplitType_WindsurfActive   SplitType = 21
	SplitType_RwdStand         SplitType = 22
	SplitType_Transition       SplitType = 23
	SplitType_SkiLiftSplit     SplitType = 28
	SplitType_SkiRunSplit      SplitType = 29
	SplitType_Invalid          SplitType = 255
)

var SplitTypeNames = map[SplitType]string{
	SplitType_AscentSplit:      "AscentSplit",
	SplitType_DescentSplit:     "DescentSplit",
	SplitType_IntervalActive:   "IntervalActive",
	SplitType_IntervalRest:     "IntervalRest",
	SplitType_IntervalWarmup:   "IntervalWarmup",
	SplitType_IntervalCooldown: "IntervalCooldown",
	SplitType_IntervalRecovery: "IntervalRecovery",
	SplitType_IntervalOther:    "IntervalOther",
	SplitType_ClimbActive:      "ClimbActive",
	SplitType_ClimbRest:        "ClimbRest",
	SplitType_SurfActive:       "SurfActive",
	SplitType_RunActive:        "RunActive",
	SplitType_RunRest:          "RunRest",
	SplitType_WorkoutRound:     "WorkoutRound",
	SplitType_RwdRun:           "RwdRun",
	SplitType_RwdWalk:          "RwdWalk",
	SplitType_WindsurfActive:   "WindsurfActive",
	SplitType_RwdStand:         "RwdStand",
	SplitType_Transition:       "Transition",
	SplitType_SkiLiftSplit:     "SkiLiftSplit",
	SplitType_SkiRunSplit:      "SkiRunSplit",
}

func (k SplitType) String() string {
	if uint(k) < uint(len(SplitTypeNames)) {
		if v, ok := SplitTypeNames[k]; ok {
			return v
		}
	}
	return "SplitType(" + strconv.Itoa(int(k)) + ")"
}

var SplitTypeValues = map[string]SplitType{
	"AscentSplit":      SplitType_AscentSplit,
	"DescentSplit":     SplitType_DescentSplit,
	"IntervalActive":   SplitType_IntervalActive,
	"IntervalRest":     SplitType_IntervalRest,
	"IntervalWarmup":   SplitType_IntervalWarmup,
	"IntervalCooldown": SplitType_IntervalCooldown,
	"IntervalRecovery": SplitType_IntervalRecovery,
	"IntervalOther":    SplitType_IntervalOther,
	"ClimbActive":      SplitType_ClimbActive,
	"ClimbRest":        SplitType_ClimbRest,
	"SurfActive":       SplitType_SurfActive,
	"RunActive":        SplitType_RunActive,
	"RunRest":          SplitType_RunRest,
	"WorkoutRound":     SplitType_WorkoutRound,
	"RwdRun":           SplitType_RwdRun,
	"RwdWalk":          SplitType_RwdWalk,
	"WindsurfActive":   SplitType_WindsurfActive,
	"RwdStand":         SplitType_RwdStand,
	"Transition":       SplitType_Transition,
	"SkiLiftSplit":     SplitType_SkiLiftSplit,
	"SkiRunSplit":      SplitType_SkiRunSplit,
}

func SplitTypeFromString(s string) SplitType {
	if v, ok := SplitTypeValues[s]; ok {
		return v
	}
	return SplitType(SplitType_Invalid)
}
