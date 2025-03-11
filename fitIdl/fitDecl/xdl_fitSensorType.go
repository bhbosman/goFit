package fitDecl

import strconv "strconv"

type SensorType byte

const (
	SensorType_Accelerometer SensorType = 0
	SensorType_Gyroscope     SensorType = 1
	SensorType_Compass       SensorType = 2
	SensorType_Barometer     SensorType = 3
	SensorType_Invalid       SensorType = 255
)

var SensorTypeNames = map[SensorType]string{
	SensorType_Accelerometer: "Accelerometer",
	SensorType_Gyroscope:     "Gyroscope",
	SensorType_Compass:       "Compass",
	SensorType_Barometer:     "Barometer",
}

func (k SensorType) String() string {
	if uint(k) < uint(len(SensorTypeNames)) {
		if v, ok := SensorTypeNames[k]; ok {
			return v
		}
	}
	return "SensorType(" + strconv.Itoa(int(k)) + ")"
}

var SensorTypeValues = map[string]SensorType{
	"Accelerometer": SensorType_Accelerometer,
	"Gyroscope":     SensorType_Gyroscope,
	"Compass":       SensorType_Compass,
	"Barometer":     SensorType_Barometer,
}

func SensorTypeFromString(s string) SensorType {
	if v, ok := SensorTypeValues[s]; ok {
		return v
	}
	return SensorType(SensorType_Invalid)
}
