package fitDecl

import strconv "strconv"

type CommTimeoutType uint16

const (
	CommTimeoutType_WildcardPairingTimeout CommTimeoutType = 0
	CommTimeoutType_PairingTimeout         CommTimeoutType = 1
	CommTimeoutType_ConnectionLost         CommTimeoutType = 2
	CommTimeoutType_ConnectionTimeout      CommTimeoutType = 3
	CommTimeoutType_Invalid                CommTimeoutType = 65535
)

var CommTimeoutTypeNames = map[CommTimeoutType]string{
	CommTimeoutType_WildcardPairingTimeout: "WildcardPairingTimeout",
	CommTimeoutType_PairingTimeout:         "PairingTimeout",
	CommTimeoutType_ConnectionLost:         "ConnectionLost",
	CommTimeoutType_ConnectionTimeout:      "ConnectionTimeout",
}

func (k CommTimeoutType) String() string {
	if uint(k) < uint(len(CommTimeoutTypeNames)) {
		if v, ok := CommTimeoutTypeNames[k]; ok {
			return v
		}
	}
	return "CommTimeoutType(" + strconv.Itoa(int(k)) + ")"
}

var CommTimeoutTypeValues = map[string]CommTimeoutType{
	"WildcardPairingTimeout": CommTimeoutType_WildcardPairingTimeout,
	"PairingTimeout":         CommTimeoutType_PairingTimeout,
	"ConnectionLost":         CommTimeoutType_ConnectionLost,
	"ConnectionTimeout":      CommTimeoutType_ConnectionTimeout,
}

func CommTimeoutTypeFromString(s string) CommTimeoutType {
	if v, ok := CommTimeoutTypeValues[s]; ok {
		return v
	}
	return CommTimeoutType(CommTimeoutType_Invalid)
}
