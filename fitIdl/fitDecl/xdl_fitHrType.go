package fitDecl

import strconv "strconv"

type HrType byte

const (
	HrType_Normal    HrType = 0
	HrType_Irregular HrType = 1
	HrType_Invalid   HrType = 255
)

var HrTypeNames = map[HrType]string{
	HrType_Normal:    "Normal",
	HrType_Irregular: "Irregular",
}

func (k HrType) String() string {
	if uint(k) < uint(len(HrTypeNames)) {
		if v, ok := HrTypeNames[k]; ok {
			return v
		}
	}
	return "HrType(" + strconv.Itoa(int(k)) + ")"
}

var HrTypeValues = map[string]HrType{
	"Normal":    HrType_Normal,
	"Irregular": HrType_Irregular,
}

func HrTypeFromString(s string) HrType {
	if v, ok := HrTypeValues[s]; ok {
		return v
	}
	return HrType(HrType_Invalid)
}
