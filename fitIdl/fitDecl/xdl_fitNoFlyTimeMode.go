package fitDecl

import strconv "strconv"

type NoFlyTimeMode byte

const (
	NoFlyTimeMode_Standard    NoFlyTimeMode = 0
	NoFlyTimeMode_Flat24Hours NoFlyTimeMode = 1
	NoFlyTimeMode_Invalid     NoFlyTimeMode = 255
)

var NoFlyTimeModeNames = map[NoFlyTimeMode]string{
	NoFlyTimeMode_Standard:    "Standard",
	NoFlyTimeMode_Flat24Hours: "Flat24Hours",
}

func (k NoFlyTimeMode) String() string {
	if uint(k) < uint(len(NoFlyTimeModeNames)) {
		if v, ok := NoFlyTimeModeNames[k]; ok {
			return v
		}
	}
	return "NoFlyTimeMode(" + strconv.Itoa(int(k)) + ")"
}

var NoFlyTimeModeValues = map[string]NoFlyTimeMode{
	"Standard":    NoFlyTimeMode_Standard,
	"Flat24Hours": NoFlyTimeMode_Flat24Hours,
}

func NoFlyTimeModeFromString(s string) NoFlyTimeMode {
	if v, ok := NoFlyTimeModeValues[s]; ok {
		return v
	}
	return NoFlyTimeMode(NoFlyTimeMode_Invalid)
}
