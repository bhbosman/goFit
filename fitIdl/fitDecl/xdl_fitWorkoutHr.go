package fitDecl

import strconv "strconv"

type WorkoutHr uint32

const (
	WorkoutHr_BpmOffset WorkoutHr = 100
	WorkoutHr_Invalid   WorkoutHr = 4294967295
)

var WorkoutHrNames = map[WorkoutHr]string{
	WorkoutHr_BpmOffset: "BpmOffset",
}

func (k WorkoutHr) String() string {
	if uint(k) < uint(len(WorkoutHrNames)) {
		if v, ok := WorkoutHrNames[k]; ok {
			return v
		}
	}
	return "WorkoutHr(" + strconv.Itoa(int(k)) + ")"
}

var WorkoutHrValues = map[string]WorkoutHr{
	"BpmOffset": WorkoutHr_BpmOffset,
}

func WorkoutHrFromString(s string) WorkoutHr {
	if v, ok := WorkoutHrValues[s]; ok {
		return v
	}
	return WorkoutHr(WorkoutHr_Invalid)
}
