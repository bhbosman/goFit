package fitDecl

import strconv "strconv"

type SetType uint8

const (
	SetType_Rest    SetType = 0
	SetType_Active  SetType = 1
	SetType_Invalid SetType = 255
)

var SetTypeNames = map[SetType]string{
	SetType_Rest:   "Rest",
	SetType_Active: "Active",
}

func (k SetType) String() string {
	if uint(k) < uint(len(SetTypeNames)) {
		if v, ok := SetTypeNames[k]; ok {
			return v
		}
	}
	return "SetType(" + strconv.Itoa(int(k)) + ")"
}

var SetTypeValues = map[string]SetType{
	"Rest":   SetType_Rest,
	"Active": SetType_Active,
}

func SetTypeFromString(s string) SetType {
	if v, ok := SetTypeValues[s]; ok {
		return v
	}
	return SetType(SetType_Invalid)
}
