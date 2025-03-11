package fitDecl

import strconv "strconv"

type TimeIntoDay uint32

const (
	TimeIntoDay_Invalid TimeIntoDay = 4294967295
)

var TimeIntoDayNames = map[TimeIntoDay]string{
}

func (k TimeIntoDay) String() string {
	if uint(k) < uint(len(TimeIntoDayNames)) {
		if v, ok := TimeIntoDayNames[k]; ok {
			return v
		}
	}
	return "TimeIntoDay(" + strconv.Itoa(int(k)) + ")"
}

var TimeIntoDayValues = map[string]TimeIntoDay{
}

func TimeIntoDayFromString(s string) TimeIntoDay {
		if v, ok := TimeIntoDayValues[s]; ok {
			return v
		}
	return TimeIntoDay(TimeIntoDay_Invalid)
}


