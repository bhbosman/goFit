package fitDecl

import strconv "strconv"

type BikeLightBeamAngleMode uint8

const (
	BikeLightBeamAngleMode_Manual  BikeLightBeamAngleMode = 0
	BikeLightBeamAngleMode_Auto    BikeLightBeamAngleMode = 1
	BikeLightBeamAngleMode_Invalid BikeLightBeamAngleMode = 255
)

var BikeLightBeamAngleModeNames = map[BikeLightBeamAngleMode]string{
	BikeLightBeamAngleMode_Manual: "Manual",
	BikeLightBeamAngleMode_Auto:   "Auto",
}

func (k BikeLightBeamAngleMode) String() string {
	if uint(k) < uint(len(BikeLightBeamAngleModeNames)) {
		if v, ok := BikeLightBeamAngleModeNames[k]; ok {
			return v
		}
	}
	return "BikeLightBeamAngleMode(" + strconv.Itoa(int(k)) + ")"
}

var BikeLightBeamAngleModeValues = map[string]BikeLightBeamAngleMode{
	"Manual": BikeLightBeamAngleMode_Manual,
	"Auto":   BikeLightBeamAngleMode_Auto,
}

func BikeLightBeamAngleModeFromString(s string) BikeLightBeamAngleMode {
	if v, ok := BikeLightBeamAngleModeValues[s]; ok {
		return v
	}
	return BikeLightBeamAngleMode(BikeLightBeamAngleMode_Invalid)
}
