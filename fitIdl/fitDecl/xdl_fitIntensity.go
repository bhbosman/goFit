package fitDecl

import strconv "strconv"

type Intensity byte

const (
	Intensity_Active   Intensity = 0
	Intensity_Rest     Intensity = 1
	Intensity_Warmup   Intensity = 2
	Intensity_Cooldown Intensity = 3
	Intensity_Recovery Intensity = 4
	Intensity_Interval Intensity = 5
	Intensity_Other    Intensity = 6
	Intensity_Invalid  Intensity = 255
)

var IntensityNames = map[Intensity]string{
	Intensity_Active:   "Active",
	Intensity_Rest:     "Rest",
	Intensity_Warmup:   "Warmup",
	Intensity_Cooldown: "Cooldown",
	Intensity_Recovery: "Recovery",
	Intensity_Interval: "Interval",
	Intensity_Other:    "Other",
}

func (k Intensity) String() string {
	if uint(k) < uint(len(IntensityNames)) {
		if v, ok := IntensityNames[k]; ok {
			return v
		}
	}
	return "Intensity(" + strconv.Itoa(int(k)) + ")"
}

var IntensityValues = map[string]Intensity{
	"Active":   Intensity_Active,
	"Rest":     Intensity_Rest,
	"Warmup":   Intensity_Warmup,
	"Cooldown": Intensity_Cooldown,
	"Recovery": Intensity_Recovery,
	"Interval": Intensity_Interval,
	"Other":    Intensity_Other,
}

func IntensityFromString(s string) Intensity {
	if v, ok := IntensityValues[s]; ok {
		return v
	}
	return Intensity(Intensity_Invalid)
}
