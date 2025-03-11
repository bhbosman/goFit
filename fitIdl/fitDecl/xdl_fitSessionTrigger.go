package fitDecl

import strconv "strconv"

type SessionTrigger byte

const (
	SessionTrigger_ActivityEnd      SessionTrigger = 0
	SessionTrigger_Manual           SessionTrigger = 1
	SessionTrigger_AutoMultiSport   SessionTrigger = 2
	SessionTrigger_FitnessEquipment SessionTrigger = 3
	SessionTrigger_Invalid          SessionTrigger = 255
)

var SessionTriggerNames = map[SessionTrigger]string{
	SessionTrigger_ActivityEnd:      "ActivityEnd",
	SessionTrigger_Manual:           "Manual",
	SessionTrigger_AutoMultiSport:   "AutoMultiSport",
	SessionTrigger_FitnessEquipment: "FitnessEquipment",
}

func (k SessionTrigger) String() string {
	if uint(k) < uint(len(SessionTriggerNames)) {
		if v, ok := SessionTriggerNames[k]; ok {
			return v
		}
	}
	return "SessionTrigger(" + strconv.Itoa(int(k)) + ")"
}

var SessionTriggerValues = map[string]SessionTrigger{
	"ActivityEnd":      SessionTrigger_ActivityEnd,
	"Manual":           SessionTrigger_Manual,
	"AutoMultiSport":   SessionTrigger_AutoMultiSport,
	"FitnessEquipment": SessionTrigger_FitnessEquipment,
}

func SessionTriggerFromString(s string) SessionTrigger {
	if v, ok := SessionTriggerValues[s]; ok {
		return v
	}
	return SessionTrigger(SessionTrigger_Invalid)
}
