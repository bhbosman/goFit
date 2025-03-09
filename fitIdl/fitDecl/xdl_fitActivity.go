package fitDecl

import strconv "strconv"

type Activity byte

const (
	Activity_Manual         Activity = 0
	Activity_AutoMultiSport Activity = 1
	Activity_Invalid        Activity = 255
)

var ActivityNames = map[Activity]string{
	Activity_Manual:         "Manual",
	Activity_AutoMultiSport: "AutoMultiSport",
}

func (k Activity) String() string {
	if uint(k) < uint(len(ActivityNames)) {
		if v, ok := ActivityNames[k]; ok {
			return v
		}
	}
	return "Activity(" + strconv.Itoa(int(k)) + ")"
}

var ActivityValues = map[string]Activity{
	"Manual":         Activity_Manual,
	"AutoMultiSport": Activity_AutoMultiSport,
}

func ActivityFromString(s string) Activity {
	if v, ok := ActivityValues[s]; ok {
		return v
	}
	return Activity(Activity_Invalid)
}
