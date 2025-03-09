package fitDecl

import strconv "strconv"

type LocalDeviceType uint8

const (
	LocalDeviceType_Gps           LocalDeviceType = 0
	LocalDeviceType_Glonass       LocalDeviceType = 1
	LocalDeviceType_GpsGlonass    LocalDeviceType = 2
	LocalDeviceType_Accelerometer LocalDeviceType = 3
	LocalDeviceType_Barometer     LocalDeviceType = 4
	LocalDeviceType_Temperature   LocalDeviceType = 5
	LocalDeviceType_Whr           LocalDeviceType = 10
	LocalDeviceType_SensorHub     LocalDeviceType = 12
	LocalDeviceType_Invalid       LocalDeviceType = 255
)

var LocalDeviceTypeNames = map[LocalDeviceType]string{
	LocalDeviceType_Gps:           "Gps",
	LocalDeviceType_Glonass:       "Glonass",
	LocalDeviceType_GpsGlonass:    "GpsGlonass",
	LocalDeviceType_Accelerometer: "Accelerometer",
	LocalDeviceType_Barometer:     "Barometer",
	LocalDeviceType_Temperature:   "Temperature",
	LocalDeviceType_Whr:           "Whr",
	LocalDeviceType_SensorHub:     "SensorHub",
}

func (k LocalDeviceType) String() string {
	if uint(k) < uint(len(LocalDeviceTypeNames)) {
		if v, ok := LocalDeviceTypeNames[k]; ok {
			return v
		}
	}
	return "LocalDeviceType(" + strconv.Itoa(int(k)) + ")"
}

var LocalDeviceTypeValues = map[string]LocalDeviceType{
	"Gps":           LocalDeviceType_Gps,
	"Glonass":       LocalDeviceType_Glonass,
	"GpsGlonass":    LocalDeviceType_GpsGlonass,
	"Accelerometer": LocalDeviceType_Accelerometer,
	"Barometer":     LocalDeviceType_Barometer,
	"Temperature":   LocalDeviceType_Temperature,
	"Whr":           LocalDeviceType_Whr,
	"SensorHub":     LocalDeviceType_SensorHub,
}

func LocalDeviceTypeFromString(s string) LocalDeviceType {
	if v, ok := LocalDeviceTypeValues[s]; ok {
		return v
	}
	return LocalDeviceType(LocalDeviceType_Invalid)
}
