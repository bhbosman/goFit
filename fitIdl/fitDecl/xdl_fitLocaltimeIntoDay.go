package fitDecl

import strconv "strconv"

type LocaltimeIntoDay uint32

const (
	LocaltimeIntoDay_Invalid LocaltimeIntoDay = 4294967295
)

var LocaltimeIntoDayNames = map[LocaltimeIntoDay]string{
}

func (k LocaltimeIntoDay) String() string {
	if uint(k) < uint(len(LocaltimeIntoDayNames)) {
		if v, ok := LocaltimeIntoDayNames[k]; ok {
			return v
		}
	}
	return "LocaltimeIntoDay(" + strconv.Itoa(int(k)) + ")"
}

var LocaltimeIntoDayValues = map[string]LocaltimeIntoDay{
}

func LocaltimeIntoDayFromString(s string) LocaltimeIntoDay {
		if v, ok := LocaltimeIntoDayValues[s]; ok {
			return v
		}
	return LocaltimeIntoDay(LocaltimeIntoDay_Invalid)
}


