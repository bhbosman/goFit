package fitDecl

import strconv "strconv"

type DiveBacklightMode byte

const (
	DiveBacklightMode_AtDepth  DiveBacklightMode = 0
	DiveBacklightMode_AlwaysOn DiveBacklightMode = 1
	DiveBacklightMode_Invalid  DiveBacklightMode = 255
)

var DiveBacklightModeNames = map[DiveBacklightMode]string{
	DiveBacklightMode_AtDepth:  "AtDepth",
	DiveBacklightMode_AlwaysOn: "AlwaysOn",
}

func (k DiveBacklightMode) String() string {
	if uint(k) < uint(len(DiveBacklightModeNames)) {
		if v, ok := DiveBacklightModeNames[k]; ok {
			return v
		}
	}
	return "DiveBacklightMode(" + strconv.Itoa(int(k)) + ")"
}

var DiveBacklightModeValues = map[string]DiveBacklightMode{
	"AtDepth":  DiveBacklightMode_AtDepth,
	"AlwaysOn": DiveBacklightMode_AlwaysOn,
}

func DiveBacklightModeFromString(s string) DiveBacklightMode {
	if v, ok := DiveBacklightModeValues[s]; ok {
		return v
	}
	return DiveBacklightMode(DiveBacklightMode_Invalid)
}
