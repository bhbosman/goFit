package fitDecl

import strconv "strconv"

type WorkoutCapabilities uint32

const (
	WorkoutCapabilities_Interval         WorkoutCapabilities = 1
	WorkoutCapabilities_Custom           WorkoutCapabilities = 2
	WorkoutCapabilities_FitnessEquipment WorkoutCapabilities = 4
	WorkoutCapabilities_Firstbeat        WorkoutCapabilities = 8
	WorkoutCapabilities_NewLeaf          WorkoutCapabilities = 16
	WorkoutCapabilities_Tcx              WorkoutCapabilities = 32
	WorkoutCapabilities_Speed            WorkoutCapabilities = 128
	WorkoutCapabilities_HeartRate        WorkoutCapabilities = 256
	WorkoutCapabilities_Distance         WorkoutCapabilities = 512
	WorkoutCapabilities_Cadence          WorkoutCapabilities = 1024
	WorkoutCapabilities_Power            WorkoutCapabilities = 2048
	WorkoutCapabilities_Grade            WorkoutCapabilities = 4096
	WorkoutCapabilities_Resistance       WorkoutCapabilities = 8192
	WorkoutCapabilities_Protected        WorkoutCapabilities = 16384
	WorkoutCapabilities_Invalid          WorkoutCapabilities = 0
)

var WorkoutCapabilitiesNames = map[WorkoutCapabilities]string{
	WorkoutCapabilities_Interval:         "Interval",
	WorkoutCapabilities_Custom:           "Custom",
	WorkoutCapabilities_FitnessEquipment: "FitnessEquipment",
	WorkoutCapabilities_Firstbeat:        "Firstbeat",
	WorkoutCapabilities_NewLeaf:          "NewLeaf",
	WorkoutCapabilities_Tcx:              "Tcx",
	WorkoutCapabilities_Speed:            "Speed",
	WorkoutCapabilities_HeartRate:        "HeartRate",
	WorkoutCapabilities_Distance:         "Distance",
	WorkoutCapabilities_Cadence:          "Cadence",
	WorkoutCapabilities_Power:            "Power",
	WorkoutCapabilities_Grade:            "Grade",
	WorkoutCapabilities_Resistance:       "Resistance",
	WorkoutCapabilities_Protected:        "Protected",
}

func (k WorkoutCapabilities) String() string {
	if uint(k) < uint(len(WorkoutCapabilitiesNames)) {
		if v, ok := WorkoutCapabilitiesNames[k]; ok {
			return v
		}
	}
	return "WorkoutCapabilities(" + strconv.Itoa(int(k)) + ")"
}

var WorkoutCapabilitiesValues = map[string]WorkoutCapabilities{
	"Interval":         WorkoutCapabilities_Interval,
	"Custom":           WorkoutCapabilities_Custom,
	"FitnessEquipment": WorkoutCapabilities_FitnessEquipment,
	"Firstbeat":        WorkoutCapabilities_Firstbeat,
	"NewLeaf":          WorkoutCapabilities_NewLeaf,
	"Tcx":              WorkoutCapabilities_Tcx,
	"Speed":            WorkoutCapabilities_Speed,
	"HeartRate":        WorkoutCapabilities_HeartRate,
	"Distance":         WorkoutCapabilities_Distance,
	"Cadence":          WorkoutCapabilities_Cadence,
	"Power":            WorkoutCapabilities_Power,
	"Grade":            WorkoutCapabilities_Grade,
	"Resistance":       WorkoutCapabilities_Resistance,
	"Protected":        WorkoutCapabilities_Protected,
}

func WorkoutCapabilitiesFromString(s string) WorkoutCapabilities {
	if v, ok := WorkoutCapabilitiesValues[s]; ok {
		return v
	}
	return WorkoutCapabilities(WorkoutCapabilities_Invalid)
}
