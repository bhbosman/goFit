package fitDecl

import strconv "strconv"

type LeftRightBalance100 uint16

const (
	LeftRightBalance100_Mask    LeftRightBalance100 = 16383
	LeftRightBalance100_Right   LeftRightBalance100 = 32768
	LeftRightBalance100_Invalid LeftRightBalance100 = 65535
)

var LeftRightBalance100Names = map[LeftRightBalance100]string{
	LeftRightBalance100_Mask:  "Mask",
	LeftRightBalance100_Right: "Right",
}

func (k LeftRightBalance100) String() string {
	if uint(k) < uint(len(LeftRightBalance100Names)) {
		if v, ok := LeftRightBalance100Names[k]; ok {
			return v
		}
	}
	return "LeftRightBalance100(" + strconv.Itoa(int(k)) + ")"
}

var LeftRightBalance100Values = map[string]LeftRightBalance100{
	"Mask":  LeftRightBalance100_Mask,
	"Right": LeftRightBalance100_Right,
}

func LeftRightBalance100FromString(s string) LeftRightBalance100 {
	if v, ok := LeftRightBalance100Values[s]; ok {
		return v
	}
	return LeftRightBalance100(LeftRightBalance100_Invalid)
}
