package fitDecl

import strconv "strconv"

type MaxMetSpeedSource byte

const (
	MaxMetSpeedSource_OnboardGps   MaxMetSpeedSource = 0
	MaxMetSpeedSource_ConnectedGps MaxMetSpeedSource = 1
	MaxMetSpeedSource_Cadence      MaxMetSpeedSource = 2
	MaxMetSpeedSource_Invalid      MaxMetSpeedSource = 255
)

var MaxMetSpeedSourceNames = map[MaxMetSpeedSource]string{
	MaxMetSpeedSource_OnboardGps:   "OnboardGps",
	MaxMetSpeedSource_ConnectedGps: "ConnectedGps",
	MaxMetSpeedSource_Cadence:      "Cadence",
}

func (k MaxMetSpeedSource) String() string {
	if uint(k) < uint(len(MaxMetSpeedSourceNames)) {
		if v, ok := MaxMetSpeedSourceNames[k]; ok {
			return v
		}
	}
	return "MaxMetSpeedSource(" + strconv.Itoa(int(k)) + ")"
}

var MaxMetSpeedSourceValues = map[string]MaxMetSpeedSource{
	"OnboardGps":   MaxMetSpeedSource_OnboardGps,
	"ConnectedGps": MaxMetSpeedSource_ConnectedGps,
	"Cadence":      MaxMetSpeedSource_Cadence,
}

func MaxMetSpeedSourceFromString(s string) MaxMetSpeedSource {
	if v, ok := MaxMetSpeedSourceValues[s]; ok {
		return v
	}
	return MaxMetSpeedSource(MaxMetSpeedSource_Invalid)
}
