package fitDecl

import strconv "strconv"

type WeatherStatus byte

const (
	WeatherStatus_Clear                  WeatherStatus = 0
	WeatherStatus_PartlyCloudy           WeatherStatus = 1
	WeatherStatus_MostlyCloudy           WeatherStatus = 2
	WeatherStatus_Rain                   WeatherStatus = 3
	WeatherStatus_Snow                   WeatherStatus = 4
	WeatherStatus_Windy                  WeatherStatus = 5
	WeatherStatus_Thunderstorms          WeatherStatus = 6
	WeatherStatus_WintryMix              WeatherStatus = 7
	WeatherStatus_Fog                    WeatherStatus = 8
	WeatherStatus_Hazy                   WeatherStatus = 11
	WeatherStatus_Hail                   WeatherStatus = 12
	WeatherStatus_ScatteredShowers       WeatherStatus = 13
	WeatherStatus_ScatteredThunderstorms WeatherStatus = 14
	WeatherStatus_UnknownPrecipitation   WeatherStatus = 15
	WeatherStatus_LightRain              WeatherStatus = 16
	WeatherStatus_HeavyRain              WeatherStatus = 17
	WeatherStatus_LightSnow              WeatherStatus = 18
	WeatherStatus_HeavySnow              WeatherStatus = 19
	WeatherStatus_LightRainSnow          WeatherStatus = 20
	WeatherStatus_HeavyRainSnow          WeatherStatus = 21
	WeatherStatus_Cloudy                 WeatherStatus = 22
	WeatherStatus_Invalid                WeatherStatus = 255
)

var WeatherStatusNames = map[WeatherStatus]string{
	WeatherStatus_Clear:                  "Clear",
	WeatherStatus_PartlyCloudy:           "PartlyCloudy",
	WeatherStatus_MostlyCloudy:           "MostlyCloudy",
	WeatherStatus_Rain:                   "Rain",
	WeatherStatus_Snow:                   "Snow",
	WeatherStatus_Windy:                  "Windy",
	WeatherStatus_Thunderstorms:          "Thunderstorms",
	WeatherStatus_WintryMix:              "WintryMix",
	WeatherStatus_Fog:                    "Fog",
	WeatherStatus_Hazy:                   "Hazy",
	WeatherStatus_Hail:                   "Hail",
	WeatherStatus_ScatteredShowers:       "ScatteredShowers",
	WeatherStatus_ScatteredThunderstorms: "ScatteredThunderstorms",
	WeatherStatus_UnknownPrecipitation:   "UnknownPrecipitation",
	WeatherStatus_LightRain:              "LightRain",
	WeatherStatus_HeavyRain:              "HeavyRain",
	WeatherStatus_LightSnow:              "LightSnow",
	WeatherStatus_HeavySnow:              "HeavySnow",
	WeatherStatus_LightRainSnow:          "LightRainSnow",
	WeatherStatus_HeavyRainSnow:          "HeavyRainSnow",
	WeatherStatus_Cloudy:                 "Cloudy",
}

func (k WeatherStatus) String() string {
	if uint(k) < uint(len(WeatherStatusNames)) {
		if v, ok := WeatherStatusNames[k]; ok {
			return v
		}
	}
	return "WeatherStatus(" + strconv.Itoa(int(k)) + ")"
}

var WeatherStatusValues = map[string]WeatherStatus{
	"Clear":                  WeatherStatus_Clear,
	"PartlyCloudy":           WeatherStatus_PartlyCloudy,
	"MostlyCloudy":           WeatherStatus_MostlyCloudy,
	"Rain":                   WeatherStatus_Rain,
	"Snow":                   WeatherStatus_Snow,
	"Windy":                  WeatherStatus_Windy,
	"Thunderstorms":          WeatherStatus_Thunderstorms,
	"WintryMix":              WeatherStatus_WintryMix,
	"Fog":                    WeatherStatus_Fog,
	"Hazy":                   WeatherStatus_Hazy,
	"Hail":                   WeatherStatus_Hail,
	"ScatteredShowers":       WeatherStatus_ScatteredShowers,
	"ScatteredThunderstorms": WeatherStatus_ScatteredThunderstorms,
	"UnknownPrecipitation":   WeatherStatus_UnknownPrecipitation,
	"LightRain":              WeatherStatus_LightRain,
	"HeavyRain":              WeatherStatus_HeavyRain,
	"LightSnow":              WeatherStatus_LightSnow,
	"HeavySnow":              WeatherStatus_HeavySnow,
	"LightRainSnow":          WeatherStatus_LightRainSnow,
	"HeavyRainSnow":          WeatherStatus_HeavyRainSnow,
	"Cloudy":                 WeatherStatus_Cloudy,
}

func WeatherStatusFromString(s string) WeatherStatus {
	if v, ok := WeatherStatusValues[s]; ok {
		return v
	}
	return WeatherStatus(WeatherStatus_Invalid)
}
