package fitDecl

import strconv "strconv"

type ExdDataUnits byte

const (
	ExdDataUnits_NoUnits                        ExdDataUnits = 0
	ExdDataUnits_Laps                           ExdDataUnits = 1
	ExdDataUnits_MilesPerHour                   ExdDataUnits = 2
	ExdDataUnits_KilometersPerHour              ExdDataUnits = 3
	ExdDataUnits_FeetPerHour                    ExdDataUnits = 4
	ExdDataUnits_MetersPerHour                  ExdDataUnits = 5
	ExdDataUnits_DegreesCelsius                 ExdDataUnits = 6
	ExdDataUnits_DegreesFarenheit               ExdDataUnits = 7
	ExdDataUnits_Zone                           ExdDataUnits = 8
	ExdDataUnits_Gear                           ExdDataUnits = 9
	ExdDataUnits_Rpm                            ExdDataUnits = 10
	ExdDataUnits_Bpm                            ExdDataUnits = 11
	ExdDataUnits_Degrees                        ExdDataUnits = 12
	ExdDataUnits_Millimeters                    ExdDataUnits = 13
	ExdDataUnits_Meters                         ExdDataUnits = 14
	ExdDataUnits_Kilometers                     ExdDataUnits = 15
	ExdDataUnits_Feet                           ExdDataUnits = 16
	ExdDataUnits_Yards                          ExdDataUnits = 17
	ExdDataUnits_Kilofeet                       ExdDataUnits = 18
	ExdDataUnits_Miles                          ExdDataUnits = 19
	ExdDataUnits_Time                           ExdDataUnits = 20
	ExdDataUnits_EnumTurnType                   ExdDataUnits = 21
	ExdDataUnits_Percent                        ExdDataUnits = 22
	ExdDataUnits_Watts                          ExdDataUnits = 23
	ExdDataUnits_WattsPerKilogram               ExdDataUnits = 24
	ExdDataUnits_EnumBatteryStatus              ExdDataUnits = 25
	ExdDataUnits_EnumBikeLightBeamAngleMode     ExdDataUnits = 26
	ExdDataUnits_EnumBikeLightBatteryStatus     ExdDataUnits = 27
	ExdDataUnits_EnumBikeLightNetworkConfigType ExdDataUnits = 28
	ExdDataUnits_Lights                         ExdDataUnits = 29
	ExdDataUnits_Seconds                        ExdDataUnits = 30
	ExdDataUnits_Minutes                        ExdDataUnits = 31
	ExdDataUnits_Hours                          ExdDataUnits = 32
	ExdDataUnits_Calories                       ExdDataUnits = 33
	ExdDataUnits_Kilojoules                     ExdDataUnits = 34
	ExdDataUnits_Milliseconds                   ExdDataUnits = 35
	ExdDataUnits_SecondPerMile                  ExdDataUnits = 36
	ExdDataUnits_SecondPerKilometer             ExdDataUnits = 37
	ExdDataUnits_Centimeter                     ExdDataUnits = 38
	ExdDataUnits_EnumCoursePoint                ExdDataUnits = 39
	ExdDataUnits_Bradians                       ExdDataUnits = 40
	ExdDataUnits_EnumSport                      ExdDataUnits = 41
	ExdDataUnits_InchesHg                       ExdDataUnits = 42
	ExdDataUnits_MmHg                           ExdDataUnits = 43
	ExdDataUnits_Mbars                          ExdDataUnits = 44
	ExdDataUnits_HectoPascals                   ExdDataUnits = 45
	ExdDataUnits_FeetPerMin                     ExdDataUnits = 46
	ExdDataUnits_MetersPerMin                   ExdDataUnits = 47
	ExdDataUnits_MetersPerSec                   ExdDataUnits = 48
	ExdDataUnits_EightCardinal                  ExdDataUnits = 49
	ExdDataUnits_Invalid                        ExdDataUnits = 255
)

var ExdDataUnitsNames = map[ExdDataUnits]string{
	ExdDataUnits_NoUnits:                        "NoUnits",
	ExdDataUnits_Laps:                           "Laps",
	ExdDataUnits_MilesPerHour:                   "MilesPerHour",
	ExdDataUnits_KilometersPerHour:              "KilometersPerHour",
	ExdDataUnits_FeetPerHour:                    "FeetPerHour",
	ExdDataUnits_MetersPerHour:                  "MetersPerHour",
	ExdDataUnits_DegreesCelsius:                 "DegreesCelsius",
	ExdDataUnits_DegreesFarenheit:               "DegreesFarenheit",
	ExdDataUnits_Zone:                           "Zone",
	ExdDataUnits_Gear:                           "Gear",
	ExdDataUnits_Rpm:                            "Rpm",
	ExdDataUnits_Bpm:                            "Bpm",
	ExdDataUnits_Degrees:                        "Degrees",
	ExdDataUnits_Millimeters:                    "Millimeters",
	ExdDataUnits_Meters:                         "Meters",
	ExdDataUnits_Kilometers:                     "Kilometers",
	ExdDataUnits_Feet:                           "Feet",
	ExdDataUnits_Yards:                          "Yards",
	ExdDataUnits_Kilofeet:                       "Kilofeet",
	ExdDataUnits_Miles:                          "Miles",
	ExdDataUnits_Time:                           "Time",
	ExdDataUnits_EnumTurnType:                   "EnumTurnType",
	ExdDataUnits_Percent:                        "Percent",
	ExdDataUnits_Watts:                          "Watts",
	ExdDataUnits_WattsPerKilogram:               "WattsPerKilogram",
	ExdDataUnits_EnumBatteryStatus:              "EnumBatteryStatus",
	ExdDataUnits_EnumBikeLightBeamAngleMode:     "EnumBikeLightBeamAngleMode",
	ExdDataUnits_EnumBikeLightBatteryStatus:     "EnumBikeLightBatteryStatus",
	ExdDataUnits_EnumBikeLightNetworkConfigType: "EnumBikeLightNetworkConfigType",
	ExdDataUnits_Lights:                         "Lights",
	ExdDataUnits_Seconds:                        "Seconds",
	ExdDataUnits_Minutes:                        "Minutes",
	ExdDataUnits_Hours:                          "Hours",
	ExdDataUnits_Calories:                       "Calories",
	ExdDataUnits_Kilojoules:                     "Kilojoules",
	ExdDataUnits_Milliseconds:                   "Milliseconds",
	ExdDataUnits_SecondPerMile:                  "SecondPerMile",
	ExdDataUnits_SecondPerKilometer:             "SecondPerKilometer",
	ExdDataUnits_Centimeter:                     "Centimeter",
	ExdDataUnits_EnumCoursePoint:                "EnumCoursePoint",
	ExdDataUnits_Bradians:                       "Bradians",
	ExdDataUnits_EnumSport:                      "EnumSport",
	ExdDataUnits_InchesHg:                       "InchesHg",
	ExdDataUnits_MmHg:                           "MmHg",
	ExdDataUnits_Mbars:                          "Mbars",
	ExdDataUnits_HectoPascals:                   "HectoPascals",
	ExdDataUnits_FeetPerMin:                     "FeetPerMin",
	ExdDataUnits_MetersPerMin:                   "MetersPerMin",
	ExdDataUnits_MetersPerSec:                   "MetersPerSec",
	ExdDataUnits_EightCardinal:                  "EightCardinal",
}

func (k ExdDataUnits) String() string {
	if uint(k) < uint(len(ExdDataUnitsNames)) {
		if v, ok := ExdDataUnitsNames[k]; ok {
			return v
		}
	}
	return "ExdDataUnits(" + strconv.Itoa(int(k)) + ")"
}

var ExdDataUnitsValues = map[string]ExdDataUnits{
	"NoUnits":                        ExdDataUnits_NoUnits,
	"Laps":                           ExdDataUnits_Laps,
	"MilesPerHour":                   ExdDataUnits_MilesPerHour,
	"KilometersPerHour":              ExdDataUnits_KilometersPerHour,
	"FeetPerHour":                    ExdDataUnits_FeetPerHour,
	"MetersPerHour":                  ExdDataUnits_MetersPerHour,
	"DegreesCelsius":                 ExdDataUnits_DegreesCelsius,
	"DegreesFarenheit":               ExdDataUnits_DegreesFarenheit,
	"Zone":                           ExdDataUnits_Zone,
	"Gear":                           ExdDataUnits_Gear,
	"Rpm":                            ExdDataUnits_Rpm,
	"Bpm":                            ExdDataUnits_Bpm,
	"Degrees":                        ExdDataUnits_Degrees,
	"Millimeters":                    ExdDataUnits_Millimeters,
	"Meters":                         ExdDataUnits_Meters,
	"Kilometers":                     ExdDataUnits_Kilometers,
	"Feet":                           ExdDataUnits_Feet,
	"Yards":                          ExdDataUnits_Yards,
	"Kilofeet":                       ExdDataUnits_Kilofeet,
	"Miles":                          ExdDataUnits_Miles,
	"Time":                           ExdDataUnits_Time,
	"EnumTurnType":                   ExdDataUnits_EnumTurnType,
	"Percent":                        ExdDataUnits_Percent,
	"Watts":                          ExdDataUnits_Watts,
	"WattsPerKilogram":               ExdDataUnits_WattsPerKilogram,
	"EnumBatteryStatus":              ExdDataUnits_EnumBatteryStatus,
	"EnumBikeLightBeamAngleMode":     ExdDataUnits_EnumBikeLightBeamAngleMode,
	"EnumBikeLightBatteryStatus":     ExdDataUnits_EnumBikeLightBatteryStatus,
	"EnumBikeLightNetworkConfigType": ExdDataUnits_EnumBikeLightNetworkConfigType,
	"Lights":                         ExdDataUnits_Lights,
	"Seconds":                        ExdDataUnits_Seconds,
	"Minutes":                        ExdDataUnits_Minutes,
	"Hours":                          ExdDataUnits_Hours,
	"Calories":                       ExdDataUnits_Calories,
	"Kilojoules":                     ExdDataUnits_Kilojoules,
	"Milliseconds":                   ExdDataUnits_Milliseconds,
	"SecondPerMile":                  ExdDataUnits_SecondPerMile,
	"SecondPerKilometer":             ExdDataUnits_SecondPerKilometer,
	"Centimeter":                     ExdDataUnits_Centimeter,
	"EnumCoursePoint":                ExdDataUnits_EnumCoursePoint,
	"Bradians":                       ExdDataUnits_Bradians,
	"EnumSport":                      ExdDataUnits_EnumSport,
	"InchesHg":                       ExdDataUnits_InchesHg,
	"MmHg":                           ExdDataUnits_MmHg,
	"Mbars":                          ExdDataUnits_Mbars,
	"HectoPascals":                   ExdDataUnits_HectoPascals,
	"FeetPerMin":                     ExdDataUnits_FeetPerMin,
	"MetersPerMin":                   ExdDataUnits_MetersPerMin,
	"MetersPerSec":                   ExdDataUnits_MetersPerSec,
	"EightCardinal":                  ExdDataUnits_EightCardinal,
}

func ExdDataUnitsFromString(s string) ExdDataUnits {
	if v, ok := ExdDataUnitsValues[s]; ok {
		return v
	}
	return ExdDataUnits(ExdDataUnits_Invalid)
}
