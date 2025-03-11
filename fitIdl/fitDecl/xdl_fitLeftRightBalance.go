package fitDecl

import strconv "strconv"

type LeftRightBalance uint8

const (
	LeftRightBalance_Mask    LeftRightBalance = 127
	LeftRightBalance_Right   LeftRightBalance = 128
	LeftRightBalance_Invalid LeftRightBalance = 255
)

var LeftRightBalanceNames = map[LeftRightBalance]string{
	LeftRightBalance_Mask:  "Mask",
	LeftRightBalance_Right: "Right",
}

func (k LeftRightBalance) String() string {
	if uint(k) < uint(len(LeftRightBalanceNames)) {
		if v, ok := LeftRightBalanceNames[k]; ok {
			return v
		}
	}
	return "LeftRightBalance(" + strconv.Itoa(int(k)) + ")"
}

var LeftRightBalanceValues = map[string]LeftRightBalance{
	"Mask":  LeftRightBalance_Mask,
	"Right": LeftRightBalance_Right,
}

func LeftRightBalanceFromString(s string) LeftRightBalance {
	if v, ok := LeftRightBalanceValues[s]; ok {
		return v
	}
	return LeftRightBalance(LeftRightBalance_Invalid)
}
