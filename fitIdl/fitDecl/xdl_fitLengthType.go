package fitDecl

import strconv "strconv"

type LengthType byte

const (
	LengthType_Idle    LengthType = 0
	LengthType_Active  LengthType = 1
	LengthType_Invalid LengthType = 255
)

var LengthTypeNames = map[LengthType]string{
	LengthType_Idle:   "Idle",
	LengthType_Active: "Active",
}

func (k LengthType) String() string {
	if uint(k) < uint(len(LengthTypeNames)) {
		if v, ok := LengthTypeNames[k]; ok {
			return v
		}
	}
	return "LengthType(" + strconv.Itoa(int(k)) + ")"
}

var LengthTypeValues = map[string]LengthType{
	"Idle":   LengthType_Idle,
	"Active": LengthType_Active,
}

func LengthTypeFromString(s string) LengthType {
	if v, ok := LengthTypeValues[s]; ok {
		return v
	}
	return LengthType(LengthType_Invalid)
}
