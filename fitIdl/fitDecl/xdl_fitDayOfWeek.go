package fitDecl

import strconv "strconv"

type DayOfWeek byte

const (
	DayOfWeek_Sunday    DayOfWeek = 0
	DayOfWeek_Monday    DayOfWeek = 1
	DayOfWeek_Tuesday   DayOfWeek = 2
	DayOfWeek_Wednesday DayOfWeek = 3
	DayOfWeek_Thursday  DayOfWeek = 4
	DayOfWeek_Friday    DayOfWeek = 5
	DayOfWeek_Saturday  DayOfWeek = 6
	DayOfWeek_Invalid   DayOfWeek = 255
)

var DayOfWeekNames = map[DayOfWeek]string{
	DayOfWeek_Sunday:    "Sunday",
	DayOfWeek_Monday:    "Monday",
	DayOfWeek_Tuesday:   "Tuesday",
	DayOfWeek_Wednesday: "Wednesday",
	DayOfWeek_Thursday:  "Thursday",
	DayOfWeek_Friday:    "Friday",
	DayOfWeek_Saturday:  "Saturday",
}

func (k DayOfWeek) String() string {
	if uint(k) < uint(len(DayOfWeekNames)) {
		if v, ok := DayOfWeekNames[k]; ok {
			return v
		}
	}
	return "DayOfWeek(" + strconv.Itoa(int(k)) + ")"
}

var DayOfWeekValues = map[string]DayOfWeek{
	"Sunday":    DayOfWeek_Sunday,
	"Monday":    DayOfWeek_Monday,
	"Tuesday":   DayOfWeek_Tuesday,
	"Wednesday": DayOfWeek_Wednesday,
	"Thursday":  DayOfWeek_Thursday,
	"Friday":    DayOfWeek_Friday,
	"Saturday":  DayOfWeek_Saturday,
}

func DayOfWeekFromString(s string) DayOfWeek {
	if v, ok := DayOfWeekValues[s]; ok {
		return v
	}
	return DayOfWeek(DayOfWeek_Invalid)
}
