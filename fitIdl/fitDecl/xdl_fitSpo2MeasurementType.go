package fitDecl

import strconv "strconv"

type Spo2MeasurementType byte

const (
	Spo2MeasurementType_OffWrist        Spo2MeasurementType = 0
	Spo2MeasurementType_SpotCheck       Spo2MeasurementType = 1
	Spo2MeasurementType_ContinuousCheck Spo2MeasurementType = 2
	Spo2MeasurementType_Periodic        Spo2MeasurementType = 3
	Spo2MeasurementType_Invalid         Spo2MeasurementType = 255
)

var Spo2MeasurementTypeNames = map[Spo2MeasurementType]string{
	Spo2MeasurementType_OffWrist:        "OffWrist",
	Spo2MeasurementType_SpotCheck:       "SpotCheck",
	Spo2MeasurementType_ContinuousCheck: "ContinuousCheck",
	Spo2MeasurementType_Periodic:        "Periodic",
}

func (k Spo2MeasurementType) String() string {
	if uint(k) < uint(len(Spo2MeasurementTypeNames)) {
		if v, ok := Spo2MeasurementTypeNames[k]; ok {
			return v
		}
	}
	return "Spo2MeasurementType(" + strconv.Itoa(int(k)) + ")"
}

var Spo2MeasurementTypeValues = map[string]Spo2MeasurementType{
	"OffWrist":        Spo2MeasurementType_OffWrist,
	"SpotCheck":       Spo2MeasurementType_SpotCheck,
	"ContinuousCheck": Spo2MeasurementType_ContinuousCheck,
	"Periodic":        Spo2MeasurementType_Periodic,
}

func Spo2MeasurementTypeFromString(s string) Spo2MeasurementType {
	if v, ok := Spo2MeasurementTypeValues[s]; ok {
		return v
	}
	return Spo2MeasurementType(Spo2MeasurementType_Invalid)
}
