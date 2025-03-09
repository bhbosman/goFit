package fitDecl

import strconv "strconv"

type LocalDateTime uint32

const (
	LocalDateTime_Min     LocalDateTime = 268435456
	LocalDateTime_Invalid LocalDateTime = 4294967295
)

var LocalDateTimeNames = map[LocalDateTime]string{
	LocalDateTime_Min: "Min",
}

func (k LocalDateTime) String() string {
	if uint(k) < uint(len(LocalDateTimeNames)) {
		if v, ok := LocalDateTimeNames[k]; ok {
			return v
		}
	}
	return "LocalDateTime(" + strconv.Itoa(int(k)) + ")"
}

var LocalDateTimeValues = map[string]LocalDateTime{
	"Min": LocalDateTime_Min,
}

func LocalDateTimeFromString(s string) LocalDateTime {
	if v, ok := LocalDateTimeValues[s]; ok {
		return v
	}
	return LocalDateTime(LocalDateTime_Invalid)
}
