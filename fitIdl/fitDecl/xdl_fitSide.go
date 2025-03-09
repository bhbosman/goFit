package fitDecl

import strconv "strconv"

type Side byte

const (
	Side_Right   Side = 0
	Side_Left    Side = 1
	Side_Invalid Side = 255
)

var SideNames = map[Side]string{
	Side_Right: "Right",
	Side_Left:  "Left",
}

func (k Side) String() string {
	if uint(k) < uint(len(SideNames)) {
		if v, ok := SideNames[k]; ok {
			return v
		}
	}
	return "Side(" + strconv.Itoa(int(k)) + ")"
}

var SideValues = map[string]Side{
	"Right": Side_Right,
	"Left":  Side_Left,
}

func SideFromString(s string) Side {
	if v, ok := SideValues[s]; ok {
		return v
	}
	return Side(Side_Invalid)
}
