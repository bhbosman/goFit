package fitDecl

import strconv "strconv"

type TimerTrigger byte

const (
	TimerTrigger_Manual           TimerTrigger = 0
	TimerTrigger_Auto             TimerTrigger = 1
	TimerTrigger_FitnessEquipment TimerTrigger = 2
	TimerTrigger_Invalid          TimerTrigger = 255
)

var TimerTriggerNames = map[TimerTrigger]string{
	TimerTrigger_Manual:           "Manual",
	TimerTrigger_Auto:             "Auto",
	TimerTrigger_FitnessEquipment: "FitnessEquipment",
}

func (k TimerTrigger) String() string {
	if uint(k) < uint(len(TimerTriggerNames)) {
		if v, ok := TimerTriggerNames[k]; ok {
			return v
		}
	}
	return "TimerTrigger(" + strconv.Itoa(int(k)) + ")"
}

var TimerTriggerValues = map[string]TimerTrigger{
	"Manual":           TimerTrigger_Manual,
	"Auto":             TimerTrigger_Auto,
	"FitnessEquipment": TimerTrigger_FitnessEquipment,
}

func TimerTriggerFromString(s string) TimerTrigger {
	if v, ok := TimerTriggerValues[s]; ok {
		return v
	}
	return TimerTrigger(TimerTrigger_Invalid)
}
