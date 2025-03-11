package fitDecl

import strconv "strconv"

type WktStepTarget byte

const (
	WktStepTarget_Speed        WktStepTarget = 0
	WktStepTarget_HeartRate    WktStepTarget = 1
	WktStepTarget_Open         WktStepTarget = 2
	WktStepTarget_Cadence      WktStepTarget = 3
	WktStepTarget_Power        WktStepTarget = 4
	WktStepTarget_Grade        WktStepTarget = 5
	WktStepTarget_Resistance   WktStepTarget = 6
	WktStepTarget_Power3s      WktStepTarget = 7
	WktStepTarget_Power10s     WktStepTarget = 8
	WktStepTarget_Power30s     WktStepTarget = 9
	WktStepTarget_PowerLap     WktStepTarget = 10
	WktStepTarget_SwimStroke   WktStepTarget = 11
	WktStepTarget_SpeedLap     WktStepTarget = 12
	WktStepTarget_HeartRateLap WktStepTarget = 13
	WktStepTarget_Invalid      WktStepTarget = 255
)

var WktStepTargetNames = map[WktStepTarget]string{
	WktStepTarget_Speed:        "Speed",
	WktStepTarget_HeartRate:    "HeartRate",
	WktStepTarget_Open:         "Open",
	WktStepTarget_Cadence:      "Cadence",
	WktStepTarget_Power:        "Power",
	WktStepTarget_Grade:        "Grade",
	WktStepTarget_Resistance:   "Resistance",
	WktStepTarget_Power3s:      "Power3s",
	WktStepTarget_Power10s:     "Power10s",
	WktStepTarget_Power30s:     "Power30s",
	WktStepTarget_PowerLap:     "PowerLap",
	WktStepTarget_SwimStroke:   "SwimStroke",
	WktStepTarget_SpeedLap:     "SpeedLap",
	WktStepTarget_HeartRateLap: "HeartRateLap",
}

func (k WktStepTarget) String() string {
	if uint(k) < uint(len(WktStepTargetNames)) {
		if v, ok := WktStepTargetNames[k]; ok {
			return v
		}
	}
	return "WktStepTarget(" + strconv.Itoa(int(k)) + ")"
}

var WktStepTargetValues = map[string]WktStepTarget{
	"Speed":        WktStepTarget_Speed,
	"HeartRate":    WktStepTarget_HeartRate,
	"Open":         WktStepTarget_Open,
	"Cadence":      WktStepTarget_Cadence,
	"Power":        WktStepTarget_Power,
	"Grade":        WktStepTarget_Grade,
	"Resistance":   WktStepTarget_Resistance,
	"Power3s":      WktStepTarget_Power3s,
	"Power10s":     WktStepTarget_Power10s,
	"Power30s":     WktStepTarget_Power30s,
	"PowerLap":     WktStepTarget_PowerLap,
	"SwimStroke":   WktStepTarget_SwimStroke,
	"SpeedLap":     WktStepTarget_SpeedLap,
	"HeartRateLap": WktStepTarget_HeartRateLap,
}

func WktStepTargetFromString(s string) WktStepTarget {
	if v, ok := WktStepTargetValues[s]; ok {
		return v
	}
	return WktStepTarget(WktStepTarget_Invalid)
}
