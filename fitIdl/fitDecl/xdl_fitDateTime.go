package fitDecl

import strconv "strconv"

type DateTime uint32

const (
	DateTime_Min     DateTime = 268435456
	DateTime_Invalid DateTime = 4294967295
)

var DateTimeNames = map[DateTime]string{
	DateTime_Min: "Min",
}

func (k DateTime) String() string {
	if uint(k) < uint(len(DateTimeNames)) {
		if v, ok := DateTimeNames[k]; ok {
			return v
		}
	}
	return "DateTime(" + strconv.Itoa(int(k)) + ")"
}

var DateTimeValues = map[string]DateTime{
	"Min": DateTime_Min,
}

func DateTimeFromString(s string) DateTime {
	if v, ok := DateTimeValues[s]; ok {
		return v
	}
	return DateTime(DateTime_Invalid)
}
