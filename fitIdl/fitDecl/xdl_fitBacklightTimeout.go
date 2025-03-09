package fitDecl

import strconv "strconv"

type BacklightTimeout uint8

const (
	BacklightTimeout_Infinite BacklightTimeout = 0
	BacklightTimeout_Invalid  BacklightTimeout = 255
)

var BacklightTimeoutNames = map[BacklightTimeout]string{
	BacklightTimeout_Infinite: "Infinite",
}

func (k BacklightTimeout) String() string {
	if uint(k) < uint(len(BacklightTimeoutNames)) {
		if v, ok := BacklightTimeoutNames[k]; ok {
			return v
		}
	}
	return "BacklightTimeout(" + strconv.Itoa(int(k)) + ")"
}

var BacklightTimeoutValues = map[string]BacklightTimeout{
	"Infinite": BacklightTimeout_Infinite,
}

func BacklightTimeoutFromString(s string) BacklightTimeout {
	if v, ok := BacklightTimeoutValues[s]; ok {
		return v
	}
	return BacklightTimeout(BacklightTimeout_Invalid)
}
