package fitDecl

import strconv "strconv"

type RiderPositionType byte

const (
	RiderPositionType_Seated               RiderPositionType = 0
	RiderPositionType_Standing             RiderPositionType = 1
	RiderPositionType_TransitionToSeated   RiderPositionType = 2
	RiderPositionType_TransitionToStanding RiderPositionType = 3
	RiderPositionType_Invalid              RiderPositionType = 255
)

var RiderPositionTypeNames = map[RiderPositionType]string{
	RiderPositionType_Seated:               "Seated",
	RiderPositionType_Standing:             "Standing",
	RiderPositionType_TransitionToSeated:   "TransitionToSeated",
	RiderPositionType_TransitionToStanding: "TransitionToStanding",
}

func (k RiderPositionType) String() string {
	if uint(k) < uint(len(RiderPositionTypeNames)) {
		if v, ok := RiderPositionTypeNames[k]; ok {
			return v
		}
	}
	return "RiderPositionType(" + strconv.Itoa(int(k)) + ")"
}

var RiderPositionTypeValues = map[string]RiderPositionType{
	"Seated":               RiderPositionType_Seated,
	"Standing":             RiderPositionType_Standing,
	"TransitionToSeated":   RiderPositionType_TransitionToSeated,
	"TransitionToStanding": RiderPositionType_TransitionToStanding,
}

func RiderPositionTypeFromString(s string) RiderPositionType {
	if v, ok := RiderPositionTypeValues[s]; ok {
		return v
	}
	return RiderPositionType(RiderPositionType_Invalid)
}
