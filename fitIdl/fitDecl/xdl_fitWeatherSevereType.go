package fitDecl

import strconv "strconv"

type WeatherSevereType byte

const (
	WeatherSevereType_Unspecified             WeatherSevereType = 0
	WeatherSevereType_Tornado                 WeatherSevereType = 1
	WeatherSevereType_Tsunami                 WeatherSevereType = 2
	WeatherSevereType_Hurricane               WeatherSevereType = 3
	WeatherSevereType_ExtremeWind             WeatherSevereType = 4
	WeatherSevereType_Typhoon                 WeatherSevereType = 5
	WeatherSevereType_InlandHurricane         WeatherSevereType = 6
	WeatherSevereType_HurricaneForceWind      WeatherSevereType = 7
	WeatherSevereType_Waterspout              WeatherSevereType = 8
	WeatherSevereType_SevereThunderstorm      WeatherSevereType = 9
	WeatherSevereType_WreckhouseWinds         WeatherSevereType = 10
	WeatherSevereType_LesSuetesWind           WeatherSevereType = 11
	WeatherSevereType_Avalanche               WeatherSevereType = 12
	WeatherSevereType_FlashFlood              WeatherSevereType = 13
	WeatherSevereType_TropicalStorm           WeatherSevereType = 14
	WeatherSevereType_InlandTropicalStorm     WeatherSevereType = 15
	WeatherSevereType_Blizzard                WeatherSevereType = 16
	WeatherSevereType_IceStorm                WeatherSevereType = 17
	WeatherSevereType_FreezingRain            WeatherSevereType = 18
	WeatherSevereType_DebrisFlow              WeatherSevereType = 19
	WeatherSevereType_FlashFreeze             WeatherSevereType = 20
	WeatherSevereType_DustStorm               WeatherSevereType = 21
	WeatherSevereType_HighWind                WeatherSevereType = 22
	WeatherSevereType_WinterStorm             WeatherSevereType = 23
	WeatherSevereType_HeavyFreezingSpray      WeatherSevereType = 24
	WeatherSevereType_ExtremeCold             WeatherSevereType = 25
	WeatherSevereType_WindChill               WeatherSevereType = 26
	WeatherSevereType_ColdWave                WeatherSevereType = 27
	WeatherSevereType_HeavySnowAlert          WeatherSevereType = 28
	WeatherSevereType_LakeEffectBlowingSnow   WeatherSevereType = 29
	WeatherSevereType_SnowSquall              WeatherSevereType = 30
	WeatherSevereType_LakeEffectSnow          WeatherSevereType = 31
	WeatherSevereType_WinterWeather           WeatherSevereType = 32
	WeatherSevereType_Sleet                   WeatherSevereType = 33
	WeatherSevereType_Snowfall                WeatherSevereType = 34
	WeatherSevereType_SnowAndBlowingSnow      WeatherSevereType = 35
	WeatherSevereType_BlowingSnow             WeatherSevereType = 36
	WeatherSevereType_SnowAlert               WeatherSevereType = 37
	WeatherSevereType_ArcticOutflow           WeatherSevereType = 38
	WeatherSevereType_FreezingDrizzle         WeatherSevereType = 39
	WeatherSevereType_Storm                   WeatherSevereType = 40
	WeatherSevereType_StormSurge              WeatherSevereType = 41
	WeatherSevereType_Rainfall                WeatherSevereType = 42
	WeatherSevereType_ArealFlood              WeatherSevereType = 43
	WeatherSevereType_CoastalFlood            WeatherSevereType = 44
	WeatherSevereType_LakeshoreFlood          WeatherSevereType = 45
	WeatherSevereType_ExcessiveHeat           WeatherSevereType = 46
	WeatherSevereType_Heat                    WeatherSevereType = 47
	WeatherSevereType_Weather                 WeatherSevereType = 48
	WeatherSevereType_HighHeatAndHumidity     WeatherSevereType = 49
	WeatherSevereType_HumidexAndHealth        WeatherSevereType = 50
	WeatherSevereType_Humidex                 WeatherSevereType = 51
	WeatherSevereType_Gale                    WeatherSevereType = 52
	WeatherSevereType_FreezingSpray           WeatherSevereType = 53
	WeatherSevereType_SpecialMarine           WeatherSevereType = 54
	WeatherSevereType_Squall                  WeatherSevereType = 55
	WeatherSevereType_StrongWind              WeatherSevereType = 56
	WeatherSevereType_LakeWind                WeatherSevereType = 57
	WeatherSevereType_MarineWeather           WeatherSevereType = 58
	WeatherSevereType_Wind                    WeatherSevereType = 59
	WeatherSevereType_SmallCraftHazardousSeas WeatherSevereType = 60
	WeatherSevereType_HazardousSeas           WeatherSevereType = 61
	WeatherSevereType_SmallCraft              WeatherSevereType = 62
	WeatherSevereType_SmallCraftWinds         WeatherSevereType = 63
	WeatherSevereType_SmallCraftRoughBar      WeatherSevereType = 64
	WeatherSevereType_HighWaterLevel          WeatherSevereType = 65
	WeatherSevereType_Ashfall                 WeatherSevereType = 66
	WeatherSevereType_FreezingFog             WeatherSevereType = 67
	WeatherSevereType_DenseFog                WeatherSevereType = 68
	WeatherSevereType_DenseSmoke              WeatherSevereType = 69
	WeatherSevereType_BlowingDust             WeatherSevereType = 70
	WeatherSevereType_HardFreeze              WeatherSevereType = 71
	WeatherSevereType_Freeze                  WeatherSevereType = 72
	WeatherSevereType_Frost                   WeatherSevereType = 73
	WeatherSevereType_FireWeather             WeatherSevereType = 74
	WeatherSevereType_Flood                   WeatherSevereType = 75
	WeatherSevereType_RipTide                 WeatherSevereType = 76
	WeatherSevereType_HighSurf                WeatherSevereType = 77
	WeatherSevereType_Smog                    WeatherSevereType = 78
	WeatherSevereType_AirQuality              WeatherSevereType = 79
	WeatherSevereType_BriskWind               WeatherSevereType = 80
	WeatherSevereType_AirStagnation           WeatherSevereType = 81
	WeatherSevereType_LowWater                WeatherSevereType = 82
	WeatherSevereType_Hydrological            WeatherSevereType = 83
	WeatherSevereType_SpecialWeather          WeatherSevereType = 84
	WeatherSevereType_Invalid                 WeatherSevereType = 255
)

var WeatherSevereTypeNames = map[WeatherSevereType]string{
	WeatherSevereType_Unspecified:             "Unspecified",
	WeatherSevereType_Tornado:                 "Tornado",
	WeatherSevereType_Tsunami:                 "Tsunami",
	WeatherSevereType_Hurricane:               "Hurricane",
	WeatherSevereType_ExtremeWind:             "ExtremeWind",
	WeatherSevereType_Typhoon:                 "Typhoon",
	WeatherSevereType_InlandHurricane:         "InlandHurricane",
	WeatherSevereType_HurricaneForceWind:      "HurricaneForceWind",
	WeatherSevereType_Waterspout:              "Waterspout",
	WeatherSevereType_SevereThunderstorm:      "SevereThunderstorm",
	WeatherSevereType_WreckhouseWinds:         "WreckhouseWinds",
	WeatherSevereType_LesSuetesWind:           "LesSuetesWind",
	WeatherSevereType_Avalanche:               "Avalanche",
	WeatherSevereType_FlashFlood:              "FlashFlood",
	WeatherSevereType_TropicalStorm:           "TropicalStorm",
	WeatherSevereType_InlandTropicalStorm:     "InlandTropicalStorm",
	WeatherSevereType_Blizzard:                "Blizzard",
	WeatherSevereType_IceStorm:                "IceStorm",
	WeatherSevereType_FreezingRain:            "FreezingRain",
	WeatherSevereType_DebrisFlow:              "DebrisFlow",
	WeatherSevereType_FlashFreeze:             "FlashFreeze",
	WeatherSevereType_DustStorm:               "DustStorm",
	WeatherSevereType_HighWind:                "HighWind",
	WeatherSevereType_WinterStorm:             "WinterStorm",
	WeatherSevereType_HeavyFreezingSpray:      "HeavyFreezingSpray",
	WeatherSevereType_ExtremeCold:             "ExtremeCold",
	WeatherSevereType_WindChill:               "WindChill",
	WeatherSevereType_ColdWave:                "ColdWave",
	WeatherSevereType_HeavySnowAlert:          "HeavySnowAlert",
	WeatherSevereType_LakeEffectBlowingSnow:   "LakeEffectBlowingSnow",
	WeatherSevereType_SnowSquall:              "SnowSquall",
	WeatherSevereType_LakeEffectSnow:          "LakeEffectSnow",
	WeatherSevereType_WinterWeather:           "WinterWeather",
	WeatherSevereType_Sleet:                   "Sleet",
	WeatherSevereType_Snowfall:                "Snowfall",
	WeatherSevereType_SnowAndBlowingSnow:      "SnowAndBlowingSnow",
	WeatherSevereType_BlowingSnow:             "BlowingSnow",
	WeatherSevereType_SnowAlert:               "SnowAlert",
	WeatherSevereType_ArcticOutflow:           "ArcticOutflow",
	WeatherSevereType_FreezingDrizzle:         "FreezingDrizzle",
	WeatherSevereType_Storm:                   "Storm",
	WeatherSevereType_StormSurge:              "StormSurge",
	WeatherSevereType_Rainfall:                "Rainfall",
	WeatherSevereType_ArealFlood:              "ArealFlood",
	WeatherSevereType_CoastalFlood:            "CoastalFlood",
	WeatherSevereType_LakeshoreFlood:          "LakeshoreFlood",
	WeatherSevereType_ExcessiveHeat:           "ExcessiveHeat",
	WeatherSevereType_Heat:                    "Heat",
	WeatherSevereType_Weather:                 "Weather",
	WeatherSevereType_HighHeatAndHumidity:     "HighHeatAndHumidity",
	WeatherSevereType_HumidexAndHealth:        "HumidexAndHealth",
	WeatherSevereType_Humidex:                 "Humidex",
	WeatherSevereType_Gale:                    "Gale",
	WeatherSevereType_FreezingSpray:           "FreezingSpray",
	WeatherSevereType_SpecialMarine:           "SpecialMarine",
	WeatherSevereType_Squall:                  "Squall",
	WeatherSevereType_StrongWind:              "StrongWind",
	WeatherSevereType_LakeWind:                "LakeWind",
	WeatherSevereType_MarineWeather:           "MarineWeather",
	WeatherSevereType_Wind:                    "Wind",
	WeatherSevereType_SmallCraftHazardousSeas: "SmallCraftHazardousSeas",
	WeatherSevereType_HazardousSeas:           "HazardousSeas",
	WeatherSevereType_SmallCraft:              "SmallCraft",
	WeatherSevereType_SmallCraftWinds:         "SmallCraftWinds",
	WeatherSevereType_SmallCraftRoughBar:      "SmallCraftRoughBar",
	WeatherSevereType_HighWaterLevel:          "HighWaterLevel",
	WeatherSevereType_Ashfall:                 "Ashfall",
	WeatherSevereType_FreezingFog:             "FreezingFog",
	WeatherSevereType_DenseFog:                "DenseFog",
	WeatherSevereType_DenseSmoke:              "DenseSmoke",
	WeatherSevereType_BlowingDust:             "BlowingDust",
	WeatherSevereType_HardFreeze:              "HardFreeze",
	WeatherSevereType_Freeze:                  "Freeze",
	WeatherSevereType_Frost:                   "Frost",
	WeatherSevereType_FireWeather:             "FireWeather",
	WeatherSevereType_Flood:                   "Flood",
	WeatherSevereType_RipTide:                 "RipTide",
	WeatherSevereType_HighSurf:                "HighSurf",
	WeatherSevereType_Smog:                    "Smog",
	WeatherSevereType_AirQuality:              "AirQuality",
	WeatherSevereType_BriskWind:               "BriskWind",
	WeatherSevereType_AirStagnation:           "AirStagnation",
	WeatherSevereType_LowWater:                "LowWater",
	WeatherSevereType_Hydrological:            "Hydrological",
	WeatherSevereType_SpecialWeather:          "SpecialWeather",
}

func (k WeatherSevereType) String() string {
	if uint(k) < uint(len(WeatherSevereTypeNames)) {
		if v, ok := WeatherSevereTypeNames[k]; ok {
			return v
		}
	}
	return "WeatherSevereType(" + strconv.Itoa(int(k)) + ")"
}

var WeatherSevereTypeValues = map[string]WeatherSevereType{
	"Unspecified":             WeatherSevereType_Unspecified,
	"Tornado":                 WeatherSevereType_Tornado,
	"Tsunami":                 WeatherSevereType_Tsunami,
	"Hurricane":               WeatherSevereType_Hurricane,
	"ExtremeWind":             WeatherSevereType_ExtremeWind,
	"Typhoon":                 WeatherSevereType_Typhoon,
	"InlandHurricane":         WeatherSevereType_InlandHurricane,
	"HurricaneForceWind":      WeatherSevereType_HurricaneForceWind,
	"Waterspout":              WeatherSevereType_Waterspout,
	"SevereThunderstorm":      WeatherSevereType_SevereThunderstorm,
	"WreckhouseWinds":         WeatherSevereType_WreckhouseWinds,
	"LesSuetesWind":           WeatherSevereType_LesSuetesWind,
	"Avalanche":               WeatherSevereType_Avalanche,
	"FlashFlood":              WeatherSevereType_FlashFlood,
	"TropicalStorm":           WeatherSevereType_TropicalStorm,
	"InlandTropicalStorm":     WeatherSevereType_InlandTropicalStorm,
	"Blizzard":                WeatherSevereType_Blizzard,
	"IceStorm":                WeatherSevereType_IceStorm,
	"FreezingRain":            WeatherSevereType_FreezingRain,
	"DebrisFlow":              WeatherSevereType_DebrisFlow,
	"FlashFreeze":             WeatherSevereType_FlashFreeze,
	"DustStorm":               WeatherSevereType_DustStorm,
	"HighWind":                WeatherSevereType_HighWind,
	"WinterStorm":             WeatherSevereType_WinterStorm,
	"HeavyFreezingSpray":      WeatherSevereType_HeavyFreezingSpray,
	"ExtremeCold":             WeatherSevereType_ExtremeCold,
	"WindChill":               WeatherSevereType_WindChill,
	"ColdWave":                WeatherSevereType_ColdWave,
	"HeavySnowAlert":          WeatherSevereType_HeavySnowAlert,
	"LakeEffectBlowingSnow":   WeatherSevereType_LakeEffectBlowingSnow,
	"SnowSquall":              WeatherSevereType_SnowSquall,
	"LakeEffectSnow":          WeatherSevereType_LakeEffectSnow,
	"WinterWeather":           WeatherSevereType_WinterWeather,
	"Sleet":                   WeatherSevereType_Sleet,
	"Snowfall":                WeatherSevereType_Snowfall,
	"SnowAndBlowingSnow":      WeatherSevereType_SnowAndBlowingSnow,
	"BlowingSnow":             WeatherSevereType_BlowingSnow,
	"SnowAlert":               WeatherSevereType_SnowAlert,
	"ArcticOutflow":           WeatherSevereType_ArcticOutflow,
	"FreezingDrizzle":         WeatherSevereType_FreezingDrizzle,
	"Storm":                   WeatherSevereType_Storm,
	"StormSurge":              WeatherSevereType_StormSurge,
	"Rainfall":                WeatherSevereType_Rainfall,
	"ArealFlood":              WeatherSevereType_ArealFlood,
	"CoastalFlood":            WeatherSevereType_CoastalFlood,
	"LakeshoreFlood":          WeatherSevereType_LakeshoreFlood,
	"ExcessiveHeat":           WeatherSevereType_ExcessiveHeat,
	"Heat":                    WeatherSevereType_Heat,
	"Weather":                 WeatherSevereType_Weather,
	"HighHeatAndHumidity":     WeatherSevereType_HighHeatAndHumidity,
	"HumidexAndHealth":        WeatherSevereType_HumidexAndHealth,
	"Humidex":                 WeatherSevereType_Humidex,
	"Gale":                    WeatherSevereType_Gale,
	"FreezingSpray":           WeatherSevereType_FreezingSpray,
	"SpecialMarine":           WeatherSevereType_SpecialMarine,
	"Squall":                  WeatherSevereType_Squall,
	"StrongWind":              WeatherSevereType_StrongWind,
	"LakeWind":                WeatherSevereType_LakeWind,
	"MarineWeather":           WeatherSevereType_MarineWeather,
	"Wind":                    WeatherSevereType_Wind,
	"SmallCraftHazardousSeas": WeatherSevereType_SmallCraftHazardousSeas,
	"HazardousSeas":           WeatherSevereType_HazardousSeas,
	"SmallCraft":              WeatherSevereType_SmallCraft,
	"SmallCraftWinds":         WeatherSevereType_SmallCraftWinds,
	"SmallCraftRoughBar":      WeatherSevereType_SmallCraftRoughBar,
	"HighWaterLevel":          WeatherSevereType_HighWaterLevel,
	"Ashfall":                 WeatherSevereType_Ashfall,
	"FreezingFog":             WeatherSevereType_FreezingFog,
	"DenseFog":                WeatherSevereType_DenseFog,
	"DenseSmoke":              WeatherSevereType_DenseSmoke,
	"BlowingDust":             WeatherSevereType_BlowingDust,
	"HardFreeze":              WeatherSevereType_HardFreeze,
	"Freeze":                  WeatherSevereType_Freeze,
	"Frost":                   WeatherSevereType_Frost,
	"FireWeather":             WeatherSevereType_FireWeather,
	"Flood":                   WeatherSevereType_Flood,
	"RipTide":                 WeatherSevereType_RipTide,
	"HighSurf":                WeatherSevereType_HighSurf,
	"Smog":                    WeatherSevereType_Smog,
	"AirQuality":              WeatherSevereType_AirQuality,
	"BriskWind":               WeatherSevereType_BriskWind,
	"AirStagnation":           WeatherSevereType_AirStagnation,
	"LowWater":                WeatherSevereType_LowWater,
	"Hydrological":            WeatherSevereType_Hydrological,
	"SpecialWeather":          WeatherSevereType_SpecialWeather,
}

func WeatherSevereTypeFromString(s string) WeatherSevereType {
	if v, ok := WeatherSevereTypeValues[s]; ok {
		return v
	}
	return WeatherSevereType(WeatherSevereType_Invalid)
}
