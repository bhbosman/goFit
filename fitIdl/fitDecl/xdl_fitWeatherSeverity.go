package fitDecl

import strconv "strconv"

type WeatherSeverity byte

const (
	WeatherSeverity_Unknown   WeatherSeverity = 0
	WeatherSeverity_Warning   WeatherSeverity = 1
	WeatherSeverity_Watch     WeatherSeverity = 2
	WeatherSeverity_Advisory  WeatherSeverity = 3
	WeatherSeverity_Statement WeatherSeverity = 4
	WeatherSeverity_Invalid   WeatherSeverity = 255
)

var WeatherSeverityNames = map[WeatherSeverity]string{
	WeatherSeverity_Unknown:   "Unknown",
	WeatherSeverity_Warning:   "Warning",
	WeatherSeverity_Watch:     "Watch",
	WeatherSeverity_Advisory:  "Advisory",
	WeatherSeverity_Statement: "Statement",
}

func (k WeatherSeverity) String() string {
	if uint(k) < uint(len(WeatherSeverityNames)) {
		if v, ok := WeatherSeverityNames[k]; ok {
			return v
		}
	}
	return "WeatherSeverity(" + strconv.Itoa(int(k)) + ")"
}

var WeatherSeverityValues = map[string]WeatherSeverity{
	"Unknown":   WeatherSeverity_Unknown,
	"Warning":   WeatherSeverity_Warning,
	"Watch":     WeatherSeverity_Watch,
	"Advisory":  WeatherSeverity_Advisory,
	"Statement": WeatherSeverity_Statement,
}

func WeatherSeverityFromString(s string) WeatherSeverity {
	if v, ok := WeatherSeverityValues[s]; ok {
		return v
	}
	return WeatherSeverity(WeatherSeverity_Invalid)
}
