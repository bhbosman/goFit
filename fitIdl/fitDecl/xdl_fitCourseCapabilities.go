package fitDecl

import strconv "strconv"

type CourseCapabilities uint32

const (
	CourseCapabilities_Processed  CourseCapabilities = 1
	CourseCapabilities_Valid      CourseCapabilities = 2
	CourseCapabilities_Time       CourseCapabilities = 4
	CourseCapabilities_Distance   CourseCapabilities = 8
	CourseCapabilities_Position   CourseCapabilities = 16
	CourseCapabilities_HeartRate  CourseCapabilities = 32
	CourseCapabilities_Power      CourseCapabilities = 64
	CourseCapabilities_Cadence    CourseCapabilities = 128
	CourseCapabilities_Training   CourseCapabilities = 256
	CourseCapabilities_Navigation CourseCapabilities = 512
	CourseCapabilities_Bikeway    CourseCapabilities = 1024
	CourseCapabilities_Aviation   CourseCapabilities = 4096
	CourseCapabilities_Invalid    CourseCapabilities = 0
)

var CourseCapabilitiesNames = map[CourseCapabilities]string{
	CourseCapabilities_Processed:  "Processed",
	CourseCapabilities_Valid:      "Valid",
	CourseCapabilities_Time:       "Time",
	CourseCapabilities_Distance:   "Distance",
	CourseCapabilities_Position:   "Position",
	CourseCapabilities_HeartRate:  "HeartRate",
	CourseCapabilities_Power:      "Power",
	CourseCapabilities_Cadence:    "Cadence",
	CourseCapabilities_Training:   "Training",
	CourseCapabilities_Navigation: "Navigation",
	CourseCapabilities_Bikeway:    "Bikeway",
	CourseCapabilities_Aviation:   "Aviation",
}

func (k CourseCapabilities) String() string {
	if uint(k) < uint(len(CourseCapabilitiesNames)) {
		if v, ok := CourseCapabilitiesNames[k]; ok {
			return v
		}
	}
	return "CourseCapabilities(" + strconv.Itoa(int(k)) + ")"
}

var CourseCapabilitiesValues = map[string]CourseCapabilities{
	"Processed":  CourseCapabilities_Processed,
	"Valid":      CourseCapabilities_Valid,
	"Time":       CourseCapabilities_Time,
	"Distance":   CourseCapabilities_Distance,
	"Position":   CourseCapabilities_Position,
	"HeartRate":  CourseCapabilities_HeartRate,
	"Power":      CourseCapabilities_Power,
	"Cadence":    CourseCapabilities_Cadence,
	"Training":   CourseCapabilities_Training,
	"Navigation": CourseCapabilities_Navigation,
	"Bikeway":    CourseCapabilities_Bikeway,
	"Aviation":   CourseCapabilities_Aviation,
}

func CourseCapabilitiesFromString(s string) CourseCapabilities {
	if v, ok := CourseCapabilitiesValues[s]; ok {
		return v
	}
	return CourseCapabilities(CourseCapabilities_Invalid)
}
