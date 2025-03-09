package fitDecl

import strconv "strconv"

type AntNetwork byte

const (
	AntNetwork_Public  AntNetwork = 0
	AntNetwork_Antplus AntNetwork = 1
	AntNetwork_Antfs   AntNetwork = 2
	AntNetwork_Private AntNetwork = 3
	AntNetwork_Invalid AntNetwork = 255
)

var AntNetworkNames = map[AntNetwork]string{
	AntNetwork_Public:  "Public",
	AntNetwork_Antplus: "Antplus",
	AntNetwork_Antfs:   "Antfs",
	AntNetwork_Private: "Private",
}

func (k AntNetwork) String() string {
	if uint(k) < uint(len(AntNetworkNames)) {
		if v, ok := AntNetworkNames[k]; ok {
			return v
		}
	}
	return "AntNetwork(" + strconv.Itoa(int(k)) + ")"
}

var AntNetworkValues = map[string]AntNetwork{
	"Public":  AntNetwork_Public,
	"Antplus": AntNetwork_Antplus,
	"Antfs":   AntNetwork_Antfs,
	"Private": AntNetwork_Private,
}

func AntNetworkFromString(s string) AntNetwork {
	if v, ok := AntNetworkValues[s]; ok {
		return v
	}
	return AntNetwork(AntNetwork_Invalid)
}
