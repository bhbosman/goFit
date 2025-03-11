package fitDecl

import strconv "strconv"

type SourceType byte

const (
	SourceType_Ant                SourceType = 0
	SourceType_Antplus            SourceType = 1
	SourceType_Bluetooth          SourceType = 2
	SourceType_BluetoothLowEnergy SourceType = 3
	SourceType_Wifi               SourceType = 4
	SourceType_Local              SourceType = 5
	SourceType_Invalid            SourceType = 255
)

var SourceTypeNames = map[SourceType]string{
	SourceType_Ant:                "Ant",
	SourceType_Antplus:            "Antplus",
	SourceType_Bluetooth:          "Bluetooth",
	SourceType_BluetoothLowEnergy: "BluetoothLowEnergy",
	SourceType_Wifi:               "Wifi",
	SourceType_Local:              "Local",
}

func (k SourceType) String() string {
	if uint(k) < uint(len(SourceTypeNames)) {
		if v, ok := SourceTypeNames[k]; ok {
			return v
		}
	}
	return "SourceType(" + strconv.Itoa(int(k)) + ")"
}

var SourceTypeValues = map[string]SourceType{
	"Ant":                SourceType_Ant,
	"Antplus":            SourceType_Antplus,
	"Bluetooth":          SourceType_Bluetooth,
	"BluetoothLowEnergy": SourceType_BluetoothLowEnergy,
	"Wifi":               SourceType_Wifi,
	"Local":              SourceType_Local,
}

func SourceTypeFromString(s string) SourceType {
	if v, ok := SourceTypeValues[s]; ok {
		return v
	}
	return SourceType(SourceType_Invalid)
}
