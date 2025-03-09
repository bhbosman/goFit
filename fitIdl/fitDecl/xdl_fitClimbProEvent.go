package fitDecl

import strconv "strconv"

type ClimbProEvent byte

const (
	ClimbProEvent_Approach ClimbProEvent = 0
	ClimbProEvent_Start    ClimbProEvent = 1
	ClimbProEvent_Complete ClimbProEvent = 2
	ClimbProEvent_Invalid  ClimbProEvent = 255
)

var ClimbProEventNames = map[ClimbProEvent]string{
	ClimbProEvent_Approach: "Approach",
	ClimbProEvent_Start:    "Start",
	ClimbProEvent_Complete: "Complete",
}

func (k ClimbProEvent) String() string {
	if uint(k) < uint(len(ClimbProEventNames)) {
		if v, ok := ClimbProEventNames[k]; ok {
			return v
		}
	}
	return "ClimbProEvent(" + strconv.Itoa(int(k)) + ")"
}

var ClimbProEventValues = map[string]ClimbProEvent{
	"Approach": ClimbProEvent_Approach,
	"Start":    ClimbProEvent_Start,
	"Complete": ClimbProEvent_Complete,
}

func ClimbProEventFromString(s string) ClimbProEvent {
	if v, ok := ClimbProEventValues[s]; ok {
		return v
	}
	return ClimbProEvent(ClimbProEvent_Invalid)
}
