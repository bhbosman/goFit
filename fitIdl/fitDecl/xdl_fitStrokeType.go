package fitDecl

import strconv "strconv"

type StrokeType byte

const (
	StrokeType_NoEvent  StrokeType = 0
	StrokeType_Other    StrokeType = 1
	StrokeType_Serve    StrokeType = 2
	StrokeType_Forehand StrokeType = 3
	StrokeType_Backhand StrokeType = 4
	StrokeType_Smash    StrokeType = 5
	StrokeType_Invalid  StrokeType = 255
)

var StrokeTypeNames = map[StrokeType]string{
	StrokeType_NoEvent:  "NoEvent",
	StrokeType_Other:    "Other",
	StrokeType_Serve:    "Serve",
	StrokeType_Forehand: "Forehand",
	StrokeType_Backhand: "Backhand",
	StrokeType_Smash:    "Smash",
}

func (k StrokeType) String() string {
	if uint(k) < uint(len(StrokeTypeNames)) {
		if v, ok := StrokeTypeNames[k]; ok {
			return v
		}
	}
	return "StrokeType(" + strconv.Itoa(int(k)) + ")"
}

var StrokeTypeValues = map[string]StrokeType{
	"NoEvent":  StrokeType_NoEvent,
	"Other":    StrokeType_Other,
	"Serve":    StrokeType_Serve,
	"Forehand": StrokeType_Forehand,
	"Backhand": StrokeType_Backhand,
	"Smash":    StrokeType_Smash,
}

func StrokeTypeFromString(s string) StrokeType {
	if v, ok := StrokeTypeValues[s]; ok {
		return v
	}
	return StrokeType(StrokeType_Invalid)
}
