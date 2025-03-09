package fitDecl

import strconv "strconv"

type WaterType byte

const (
	WaterType_Fresh   WaterType = 0
	WaterType_Salt    WaterType = 1
	WaterType_En13319 WaterType = 2
	WaterType_Custom  WaterType = 3
	WaterType_Invalid WaterType = 255
)

var WaterTypeNames = map[WaterType]string{
	WaterType_Fresh:   "Fresh",
	WaterType_Salt:    "Salt",
	WaterType_En13319: "En13319",
	WaterType_Custom:  "Custom",
}

func (k WaterType) String() string {
	if uint(k) < uint(len(WaterTypeNames)) {
		if v, ok := WaterTypeNames[k]; ok {
			return v
		}
	}
	return "WaterType(" + strconv.Itoa(int(k)) + ")"
}

var WaterTypeValues = map[string]WaterType{
	"Fresh":   WaterType_Fresh,
	"Salt":    WaterType_Salt,
	"En13319": WaterType_En13319,
	"Custom":  WaterType_Custom,
}

func WaterTypeFromString(s string) WaterType {
	if v, ok := WaterTypeValues[s]; ok {
		return v
	}
	return WaterType(WaterType_Invalid)
}
