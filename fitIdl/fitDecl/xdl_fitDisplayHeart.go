package fitDecl

import strconv "strconv"

type DisplayHeart byte

const (
	DisplayHeart_Bpm     DisplayHeart = 0
	DisplayHeart_Max     DisplayHeart = 1
	DisplayHeart_Reserve DisplayHeart = 2
	DisplayHeart_Invalid DisplayHeart = 255
)

var DisplayHeartNames = map[DisplayHeart]string{
	DisplayHeart_Bpm:     "Bpm",
	DisplayHeart_Max:     "Max",
	DisplayHeart_Reserve: "Reserve",
}

func (k DisplayHeart) String() string {
	if uint(k) < uint(len(DisplayHeartNames)) {
		if v, ok := DisplayHeartNames[k]; ok {
			return v
		}
	}
	return "DisplayHeart(" + strconv.Itoa(int(k)) + ")"
}

var DisplayHeartValues = map[string]DisplayHeart{
	"Bpm":     DisplayHeart_Bpm,
	"Max":     DisplayHeart_Max,
	"Reserve": DisplayHeart_Reserve,
}

func DisplayHeartFromString(s string) DisplayHeart {
	if v, ok := DisplayHeartValues[s]; ok {
		return v
	}
	return DisplayHeart(DisplayHeart_Invalid)
}
