package fitDecl

import strconv "strconv"

type TimeMode byte

const (
	TimeMode_Hour12            TimeMode = 0
	TimeMode_Hour24            TimeMode = 1
	TimeMode_Military          TimeMode = 2
	TimeMode_Hour12WithSeconds TimeMode = 3
	TimeMode_Hour24WithSeconds TimeMode = 4
	TimeMode_Utc               TimeMode = 5
	TimeMode_Invalid           TimeMode = 255
)

var TimeModeNames = map[TimeMode]string{
	TimeMode_Hour12:            "Hour12",
	TimeMode_Hour24:            "Hour24",
	TimeMode_Military:          "Military",
	TimeMode_Hour12WithSeconds: "Hour12WithSeconds",
	TimeMode_Hour24WithSeconds: "Hour24WithSeconds",
	TimeMode_Utc:               "Utc",
}

func (k TimeMode) String() string {
	if uint(k) < uint(len(TimeModeNames)) {
		if v, ok := TimeModeNames[k]; ok {
			return v
		}
	}
	return "TimeMode(" + strconv.Itoa(int(k)) + ")"
}

var TimeModeValues = map[string]TimeMode{
	"Hour12":            TimeMode_Hour12,
	"Hour24":            TimeMode_Hour24,
	"Military":          TimeMode_Military,
	"Hour12WithSeconds": TimeMode_Hour12WithSeconds,
	"Hour24WithSeconds": TimeMode_Hour24WithSeconds,
	"Utc":               TimeMode_Utc,
}

func TimeModeFromString(s string) TimeMode {
	if v, ok := TimeModeValues[s]; ok {
		return v
	}
	return TimeMode(TimeMode_Invalid)
}
