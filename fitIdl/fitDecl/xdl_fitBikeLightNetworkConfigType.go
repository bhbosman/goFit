package fitDecl

import strconv "strconv"

type BikeLightNetworkConfigType byte

const (
	BikeLightNetworkConfigType_Auto           BikeLightNetworkConfigType = 0
	BikeLightNetworkConfigType_Individual     BikeLightNetworkConfigType = 4
	BikeLightNetworkConfigType_HighVisibility BikeLightNetworkConfigType = 5
	BikeLightNetworkConfigType_Trail          BikeLightNetworkConfigType = 6
	BikeLightNetworkConfigType_Invalid        BikeLightNetworkConfigType = 255
)

var BikeLightNetworkConfigTypeNames = map[BikeLightNetworkConfigType]string{
	BikeLightNetworkConfigType_Auto:           "Auto",
	BikeLightNetworkConfigType_Individual:     "Individual",
	BikeLightNetworkConfigType_HighVisibility: "HighVisibility",
	BikeLightNetworkConfigType_Trail:          "Trail",
}

func (k BikeLightNetworkConfigType) String() string {
	if uint(k) < uint(len(BikeLightNetworkConfigTypeNames)) {
		if v, ok := BikeLightNetworkConfigTypeNames[k]; ok {
			return v
		}
	}
	return "BikeLightNetworkConfigType(" + strconv.Itoa(int(k)) + ")"
}

var BikeLightNetworkConfigTypeValues = map[string]BikeLightNetworkConfigType{
	"Auto":           BikeLightNetworkConfigType_Auto,
	"Individual":     BikeLightNetworkConfigType_Individual,
	"HighVisibility": BikeLightNetworkConfigType_HighVisibility,
	"Trail":          BikeLightNetworkConfigType_Trail,
}

func BikeLightNetworkConfigTypeFromString(s string) BikeLightNetworkConfigType {
	if v, ok := BikeLightNetworkConfigTypeValues[s]; ok {
		return v
	}
	return BikeLightNetworkConfigType(BikeLightNetworkConfigType_Invalid)
}
