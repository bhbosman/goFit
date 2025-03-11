package fitDecl

import strconv "strconv"

type AutolapTrigger byte

const (
	AutolapTrigger_Time             AutolapTrigger = 0
	AutolapTrigger_Distance         AutolapTrigger = 1
	AutolapTrigger_PositionStart    AutolapTrigger = 2
	AutolapTrigger_PositionLap      AutolapTrigger = 3
	AutolapTrigger_PositionWaypoint AutolapTrigger = 4
	AutolapTrigger_PositionMarked   AutolapTrigger = 5
	AutolapTrigger_Off              AutolapTrigger = 6
	AutolapTrigger_AutoSelect       AutolapTrigger = 13
	AutolapTrigger_Invalid          AutolapTrigger = 255
)

var AutolapTriggerNames = map[AutolapTrigger]string{
	AutolapTrigger_Time:             "Time",
	AutolapTrigger_Distance:         "Distance",
	AutolapTrigger_PositionStart:    "PositionStart",
	AutolapTrigger_PositionLap:      "PositionLap",
	AutolapTrigger_PositionWaypoint: "PositionWaypoint",
	AutolapTrigger_PositionMarked:   "PositionMarked",
	AutolapTrigger_Off:              "Off",
	AutolapTrigger_AutoSelect:       "AutoSelect",
}

func (k AutolapTrigger) String() string {
	if uint(k) < uint(len(AutolapTriggerNames)) {
		if v, ok := AutolapTriggerNames[k]; ok {
			return v
		}
	}
	return "AutolapTrigger(" + strconv.Itoa(int(k)) + ")"
}

var AutolapTriggerValues = map[string]AutolapTrigger{
	"Time":             AutolapTrigger_Time,
	"Distance":         AutolapTrigger_Distance,
	"PositionStart":    AutolapTrigger_PositionStart,
	"PositionLap":      AutolapTrigger_PositionLap,
	"PositionWaypoint": AutolapTrigger_PositionWaypoint,
	"PositionMarked":   AutolapTrigger_PositionMarked,
	"Off":              AutolapTrigger_Off,
	"AutoSelect":       AutolapTrigger_AutoSelect,
}

func AutolapTriggerFromString(s string) AutolapTrigger {
	if v, ok := AutolapTriggerValues[s]; ok {
		return v
	}
	return AutolapTrigger(AutolapTrigger_Invalid)
}
