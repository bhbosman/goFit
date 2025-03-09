package fitDecl

import strconv "strconv"

type WeatherReport byte

const (
	WeatherReport_Current        WeatherReport = 0
	WeatherReport_HourlyForecast WeatherReport = 1
	WeatherReport_DailyForecast  WeatherReport = 2
	WeatherReport_Invalid        WeatherReport = 255
)

var WeatherReportNames = map[WeatherReport]string{
	WeatherReport_Current:        "Current",
	WeatherReport_HourlyForecast: "HourlyForecast",
	WeatherReport_DailyForecast:  "DailyForecast",
}

func (k WeatherReport) String() string {
	if uint(k) < uint(len(WeatherReportNames)) {
		if v, ok := WeatherReportNames[k]; ok {
			return v
		}
	}
	return "WeatherReport(" + strconv.Itoa(int(k)) + ")"
}

var WeatherReportValues = map[string]WeatherReport{
	"Current":        WeatherReport_Current,
	"HourlyForecast": WeatherReport_HourlyForecast,
	"DailyForecast":  WeatherReport_DailyForecast,
}

func WeatherReportFromString(s string) WeatherReport {
	if v, ok := WeatherReportValues[s]; ok {
		return v
	}
	return WeatherReport(WeatherReport_Invalid)
}
