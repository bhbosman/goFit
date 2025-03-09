package fitDecl

import strconv "strconv"

type BleDeviceType uint8

const (
	BleDeviceType_ConnectedGps     BleDeviceType = 0
	BleDeviceType_HeartRate        BleDeviceType = 1
	BleDeviceType_BikePower        BleDeviceType = 2
	BleDeviceType_BikeSpeedCadence BleDeviceType = 3
	BleDeviceType_BikeSpeed        BleDeviceType = 4
	BleDeviceType_BikeCadence      BleDeviceType = 5
	BleDeviceType_Footpod          BleDeviceType = 6
	BleDeviceType_BikeTrainer      BleDeviceType = 7
	BleDeviceType_Invalid          BleDeviceType = 255
)

var BleDeviceTypeNames = map[BleDeviceType]string{
	BleDeviceType_ConnectedGps:     "ConnectedGps",
	BleDeviceType_HeartRate:        "HeartRate",
	BleDeviceType_BikePower:        "BikePower",
	BleDeviceType_BikeSpeedCadence: "BikeSpeedCadence",
	BleDeviceType_BikeSpeed:        "BikeSpeed",
	BleDeviceType_BikeCadence:      "BikeCadence",
	BleDeviceType_Footpod:          "Footpod",
	BleDeviceType_BikeTrainer:      "BikeTrainer",
}

func (k BleDeviceType) String() string {
	if uint(k) < uint(len(BleDeviceTypeNames)) {
		if v, ok := BleDeviceTypeNames[k]; ok {
			return v
		}
	}
	return "BleDeviceType(" + strconv.Itoa(int(k)) + ")"
}

var BleDeviceTypeValues = map[string]BleDeviceType{
	"ConnectedGps":     BleDeviceType_ConnectedGps,
	"HeartRate":        BleDeviceType_HeartRate,
	"BikePower":        BleDeviceType_BikePower,
	"BikeSpeedCadence": BleDeviceType_BikeSpeedCadence,
	"BikeSpeed":        BleDeviceType_BikeSpeed,
	"BikeCadence":      BleDeviceType_BikeCadence,
	"Footpod":          BleDeviceType_Footpod,
	"BikeTrainer":      BleDeviceType_BikeTrainer,
}

func BleDeviceTypeFromString(s string) BleDeviceType {
	if v, ok := BleDeviceTypeValues[s]; ok {
		return v
	}
	return BleDeviceType(BleDeviceType_Invalid)
}
