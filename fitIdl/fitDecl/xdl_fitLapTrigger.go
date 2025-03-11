package fitDecl

import strconv "strconv"

type LapTrigger byte

const (
	LapTrigger_Manual           LapTrigger = 0
	LapTrigger_Time             LapTrigger = 1
	LapTrigger_Distance         LapTrigger = 2
	LapTrigger_PositionStart    LapTrigger = 3
	LapTrigger_PositionLap      LapTrigger = 4
	LapTrigger_PositionWaypoint LapTrigger = 5
	LapTrigger_PositionMarked   LapTrigger = 6
	LapTrigger_SessionEnd       LapTrigger = 7
	LapTrigger_FitnessEquipment LapTrigger = 8
	LapTrigger_Invalid          LapTrigger = 255
)

var LapTriggerNames = map[LapTrigger]string{
	LapTrigger_Manual:           "Manual",
	LapTrigger_Time:             "Time",
	LapTrigger_Distance:         "Distance",
	LapTrigger_PositionStart:    "PositionStart",
	LapTrigger_PositionLap:      "PositionLap",
	LapTrigger_PositionWaypoint: "PositionWaypoint",
	LapTrigger_PositionMarked:   "PositionMarked",
	LapTrigger_SessionEnd:       "SessionEnd",
	LapTrigger_FitnessEquipment: "FitnessEquipment",
}

func (k LapTrigger) String() string {
	if uint(k) < uint(len(LapTriggerNames)) {
		if v, ok := LapTriggerNames[k]; ok {
			return v
		}
	}
	return "LapTrigger(" + strconv.Itoa(int(k)) + ")"
}

var LapTriggerValues = map[string]LapTrigger{
	"Manual":           LapTrigger_Manual,
	"Time":             LapTrigger_Time,
	"Distance":         LapTrigger_Distance,
	"PositionStart":    LapTrigger_PositionStart,
	"PositionLap":      LapTrigger_PositionLap,
	"PositionWaypoint": LapTrigger_PositionWaypoint,
	"PositionMarked":   LapTrigger_PositionMarked,
	"SessionEnd":       LapTrigger_SessionEnd,
	"FitnessEquipment": LapTrigger_FitnessEquipment,
}

func LapTriggerFromString(s string) LapTrigger {
	if v, ok := LapTriggerValues[s]; ok {
		return v
	}
	return LapTrigger(LapTrigger_Invalid)
}
