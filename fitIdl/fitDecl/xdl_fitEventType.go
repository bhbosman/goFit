package fitDecl

import strconv "strconv"

type EventType byte

const (
	EventType_Start                  EventType = 0
	EventType_Stop                   EventType = 1
	EventType_ConsecutiveDepreciated EventType = 2
	EventType_Marker                 EventType = 3
	EventType_StopAll                EventType = 4
	EventType_BeginDepreciated       EventType = 5
	EventType_EndDepreciated         EventType = 6
	EventType_EndAllDepreciated      EventType = 7
	EventType_StopDisable            EventType = 8
	EventType_StopDisableAll         EventType = 9
	EventType_Invalid                EventType = 255
)

var EventTypeNames = map[EventType]string{
	EventType_Start:                  "Start",
	EventType_Stop:                   "Stop",
	EventType_ConsecutiveDepreciated: "ConsecutiveDepreciated",
	EventType_Marker:                 "Marker",
	EventType_StopAll:                "StopAll",
	EventType_BeginDepreciated:       "BeginDepreciated",
	EventType_EndDepreciated:         "EndDepreciated",
	EventType_EndAllDepreciated:      "EndAllDepreciated",
	EventType_StopDisable:            "StopDisable",
	EventType_StopDisableAll:         "StopDisableAll",
}

func (k EventType) String() string {
	if uint(k) < uint(len(EventTypeNames)) {
		if v, ok := EventTypeNames[k]; ok {
			return v
		}
	}
	return "EventType(" + strconv.Itoa(int(k)) + ")"
}

var EventTypeValues = map[string]EventType{
	"Start":                  EventType_Start,
	"Stop":                   EventType_Stop,
	"ConsecutiveDepreciated": EventType_ConsecutiveDepreciated,
	"Marker":                 EventType_Marker,
	"StopAll":                EventType_StopAll,
	"BeginDepreciated":       EventType_BeginDepreciated,
	"EndDepreciated":         EventType_EndDepreciated,
	"EndAllDepreciated":      EventType_EndAllDepreciated,
	"StopDisable":            EventType_StopDisable,
	"StopDisableAll":         EventType_StopDisableAll,
}

func EventTypeFromString(s string) EventType {
	if v, ok := EventTypeValues[s]; ok {
		return v
	}
	return EventType(EventType_Invalid)
}
