package fitDecl

import strconv "strconv"

type AntplusDeviceType uint8

const (
	AntplusDeviceType_Antfs                   AntplusDeviceType = 1
	AntplusDeviceType_BikePower               AntplusDeviceType = 11
	AntplusDeviceType_EnvironmentSensorLegacy AntplusDeviceType = 12
	AntplusDeviceType_MultiSportSpeedDistance AntplusDeviceType = 15
	AntplusDeviceType_Control                 AntplusDeviceType = 16
	AntplusDeviceType_FitnessEquipment        AntplusDeviceType = 17
	AntplusDeviceType_BloodPressure           AntplusDeviceType = 18
	AntplusDeviceType_GeocacheNode            AntplusDeviceType = 19
	AntplusDeviceType_LightElectricVehicle    AntplusDeviceType = 20
	AntplusDeviceType_EnvSensor               AntplusDeviceType = 25
	AntplusDeviceType_Racquet                 AntplusDeviceType = 26
	AntplusDeviceType_ControlHub              AntplusDeviceType = 27
	AntplusDeviceType_MuscleOxygen            AntplusDeviceType = 31
	AntplusDeviceType_Shifting                AntplusDeviceType = 34
	AntplusDeviceType_BikeLightMain           AntplusDeviceType = 35
	AntplusDeviceType_BikeLightShared         AntplusDeviceType = 36
	AntplusDeviceType_Exd                     AntplusDeviceType = 38
	AntplusDeviceType_BikeRadar               AntplusDeviceType = 40
	AntplusDeviceType_BikeAero                AntplusDeviceType = 46
	AntplusDeviceType_WeightScale             AntplusDeviceType = 119
	AntplusDeviceType_HeartRate               AntplusDeviceType = 120
	AntplusDeviceType_BikeSpeedCadence        AntplusDeviceType = 121
	AntplusDeviceType_BikeCadence             AntplusDeviceType = 122
	AntplusDeviceType_BikeSpeed               AntplusDeviceType = 123
	AntplusDeviceType_StrideSpeedDistance     AntplusDeviceType = 124
	AntplusDeviceType_Invalid                 AntplusDeviceType = 255
)

var AntplusDeviceTypeNames = map[AntplusDeviceType]string{
	AntplusDeviceType_Antfs:                   "Antfs",
	AntplusDeviceType_BikePower:               "BikePower",
	AntplusDeviceType_EnvironmentSensorLegacy: "EnvironmentSensorLegacy",
	AntplusDeviceType_MultiSportSpeedDistance: "MultiSportSpeedDistance",
	AntplusDeviceType_Control:                 "Control",
	AntplusDeviceType_FitnessEquipment:        "FitnessEquipment",
	AntplusDeviceType_BloodPressure:           "BloodPressure",
	AntplusDeviceType_GeocacheNode:            "GeocacheNode",
	AntplusDeviceType_LightElectricVehicle:    "LightElectricVehicle",
	AntplusDeviceType_EnvSensor:               "EnvSensor",
	AntplusDeviceType_Racquet:                 "Racquet",
	AntplusDeviceType_ControlHub:              "ControlHub",
	AntplusDeviceType_MuscleOxygen:            "MuscleOxygen",
	AntplusDeviceType_Shifting:                "Shifting",
	AntplusDeviceType_BikeLightMain:           "BikeLightMain",
	AntplusDeviceType_BikeLightShared:         "BikeLightShared",
	AntplusDeviceType_Exd:                     "Exd",
	AntplusDeviceType_BikeRadar:               "BikeRadar",
	AntplusDeviceType_BikeAero:                "BikeAero",
	AntplusDeviceType_WeightScale:             "WeightScale",
	AntplusDeviceType_HeartRate:               "HeartRate",
	AntplusDeviceType_BikeSpeedCadence:        "BikeSpeedCadence",
	AntplusDeviceType_BikeCadence:             "BikeCadence",
	AntplusDeviceType_BikeSpeed:               "BikeSpeed",
	AntplusDeviceType_StrideSpeedDistance:     "StrideSpeedDistance",
}

func (k AntplusDeviceType) String() string {
	if uint(k) < uint(len(AntplusDeviceTypeNames)) {
		if v, ok := AntplusDeviceTypeNames[k]; ok {
			return v
		}
	}
	return "AntplusDeviceType(" + strconv.Itoa(int(k)) + ")"
}

var AntplusDeviceTypeValues = map[string]AntplusDeviceType{
	"Antfs":                   AntplusDeviceType_Antfs,
	"BikePower":               AntplusDeviceType_BikePower,
	"EnvironmentSensorLegacy": AntplusDeviceType_EnvironmentSensorLegacy,
	"MultiSportSpeedDistance": AntplusDeviceType_MultiSportSpeedDistance,
	"Control":                 AntplusDeviceType_Control,
	"FitnessEquipment":        AntplusDeviceType_FitnessEquipment,
	"BloodPressure":           AntplusDeviceType_BloodPressure,
	"GeocacheNode":            AntplusDeviceType_GeocacheNode,
	"LightElectricVehicle":    AntplusDeviceType_LightElectricVehicle,
	"EnvSensor":               AntplusDeviceType_EnvSensor,
	"Racquet":                 AntplusDeviceType_Racquet,
	"ControlHub":              AntplusDeviceType_ControlHub,
	"MuscleOxygen":            AntplusDeviceType_MuscleOxygen,
	"Shifting":                AntplusDeviceType_Shifting,
	"BikeLightMain":           AntplusDeviceType_BikeLightMain,
	"BikeLightShared":         AntplusDeviceType_BikeLightShared,
	"Exd":                     AntplusDeviceType_Exd,
	"BikeRadar":               AntplusDeviceType_BikeRadar,
	"BikeAero":                AntplusDeviceType_BikeAero,
	"WeightScale":             AntplusDeviceType_WeightScale,
	"HeartRate":               AntplusDeviceType_HeartRate,
	"BikeSpeedCadence":        AntplusDeviceType_BikeSpeedCadence,
	"BikeCadence":             AntplusDeviceType_BikeCadence,
	"BikeSpeed":               AntplusDeviceType_BikeSpeed,
	"StrideSpeedDistance":     AntplusDeviceType_StrideSpeedDistance,
}

func AntplusDeviceTypeFromString(s string) AntplusDeviceType {
	if v, ok := AntplusDeviceTypeValues[s]; ok {
		return v
	}
	return AntplusDeviceType(AntplusDeviceType_Invalid)
}
