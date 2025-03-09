package fitDecl

import strconv "strconv"

type DateMode byte

const (
	DateMode_DayMonth DateMode = 0
	DateMode_MonthDay DateMode = 1
	DateMode_Invalid  DateMode = 255
)

var DateModeNames = map[DateMode]string{
	DateMode_DayMonth: "DayMonth",
	DateMode_MonthDay: "MonthDay",
}

func (k DateMode) String() string {
	if uint(k) < uint(len(DateModeNames)) {
		if v, ok := DateModeNames[k]; ok {
			return v
		}
	}
	return "DateMode(" + strconv.Itoa(int(k)) + ")"
}

var DateModeValues = map[string]DateMode{
	"DayMonth": DateMode_DayMonth,
	"MonthDay": DateMode_MonthDay,
}

func DateModeFromString(s string) DateMode {
	if v, ok := DateModeValues[s]; ok {
		return v
	}
	return DateMode(DateMode_Invalid)
}
