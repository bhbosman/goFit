package fitDecl

import strconv "strconv"

type ExdQualifiers byte

const (
	ExdQualifiers_NoQualifier              ExdQualifiers = 0
	ExdQualifiers_Instantaneous            ExdQualifiers = 1
	ExdQualifiers_Average                  ExdQualifiers = 2
	ExdQualifiers_Lap                      ExdQualifiers = 3
	ExdQualifiers_Maximum                  ExdQualifiers = 4
	ExdQualifiers_MaximumAverage           ExdQualifiers = 5
	ExdQualifiers_MaximumLap               ExdQualifiers = 6
	ExdQualifiers_LastLap                  ExdQualifiers = 7
	ExdQualifiers_AverageLap               ExdQualifiers = 8
	ExdQualifiers_ToDestination            ExdQualifiers = 9
	ExdQualifiers_ToGo                     ExdQualifiers = 10
	ExdQualifiers_ToNext                   ExdQualifiers = 11
	ExdQualifiers_NextCoursePoint          ExdQualifiers = 12
	ExdQualifiers_Total                    ExdQualifiers = 13
	ExdQualifiers_ThreeSecondAverage       ExdQualifiers = 14
	ExdQualifiers_TenSecondAverage         ExdQualifiers = 15
	ExdQualifiers_ThirtySecondAverage      ExdQualifiers = 16
	ExdQualifiers_PercentMaximum           ExdQualifiers = 17
	ExdQualifiers_PercentMaximumAverage    ExdQualifiers = 18
	ExdQualifiers_LapPercentMaximum        ExdQualifiers = 19
	ExdQualifiers_Elapsed                  ExdQualifiers = 20
	ExdQualifiers_Sunrise                  ExdQualifiers = 21
	ExdQualifiers_Sunset                   ExdQualifiers = 22
	ExdQualifiers_ComparedToVirtualPartner ExdQualifiers = 23
	ExdQualifiers_Maximum24h               ExdQualifiers = 24
	ExdQualifiers_Minimum24h               ExdQualifiers = 25
	ExdQualifiers_Minimum                  ExdQualifiers = 26
	ExdQualifiers_First                    ExdQualifiers = 27
	ExdQualifiers_Second                   ExdQualifiers = 28
	ExdQualifiers_Third                    ExdQualifiers = 29
	ExdQualifiers_Shifter                  ExdQualifiers = 30
	ExdQualifiers_LastSport                ExdQualifiers = 31
	ExdQualifiers_Moving                   ExdQualifiers = 32
	ExdQualifiers_Stopped                  ExdQualifiers = 33
	ExdQualifiers_EstimatedTotal           ExdQualifiers = 34
	ExdQualifiers_Zone9                    ExdQualifiers = 242
	ExdQualifiers_Zone8                    ExdQualifiers = 243
	ExdQualifiers_Zone7                    ExdQualifiers = 244
	ExdQualifiers_Zone6                    ExdQualifiers = 245
	ExdQualifiers_Zone5                    ExdQualifiers = 246
	ExdQualifiers_Zone4                    ExdQualifiers = 247
	ExdQualifiers_Zone3                    ExdQualifiers = 248
	ExdQualifiers_Zone2                    ExdQualifiers = 249
	ExdQualifiers_Zone1                    ExdQualifiers = 250
	ExdQualifiers_Invalid                  ExdQualifiers = 255
)

var ExdQualifiersNames = map[ExdQualifiers]string{
	ExdQualifiers_NoQualifier:              "NoQualifier",
	ExdQualifiers_Instantaneous:            "Instantaneous",
	ExdQualifiers_Average:                  "Average",
	ExdQualifiers_Lap:                      "Lap",
	ExdQualifiers_Maximum:                  "Maximum",
	ExdQualifiers_MaximumAverage:           "MaximumAverage",
	ExdQualifiers_MaximumLap:               "MaximumLap",
	ExdQualifiers_LastLap:                  "LastLap",
	ExdQualifiers_AverageLap:               "AverageLap",
	ExdQualifiers_ToDestination:            "ToDestination",
	ExdQualifiers_ToGo:                     "ToGo",
	ExdQualifiers_ToNext:                   "ToNext",
	ExdQualifiers_NextCoursePoint:          "NextCoursePoint",
	ExdQualifiers_Total:                    "Total",
	ExdQualifiers_ThreeSecondAverage:       "ThreeSecondAverage",
	ExdQualifiers_TenSecondAverage:         "TenSecondAverage",
	ExdQualifiers_ThirtySecondAverage:      "ThirtySecondAverage",
	ExdQualifiers_PercentMaximum:           "PercentMaximum",
	ExdQualifiers_PercentMaximumAverage:    "PercentMaximumAverage",
	ExdQualifiers_LapPercentMaximum:        "LapPercentMaximum",
	ExdQualifiers_Elapsed:                  "Elapsed",
	ExdQualifiers_Sunrise:                  "Sunrise",
	ExdQualifiers_Sunset:                   "Sunset",
	ExdQualifiers_ComparedToVirtualPartner: "ComparedToVirtualPartner",
	ExdQualifiers_Maximum24h:               "Maximum24h",
	ExdQualifiers_Minimum24h:               "Minimum24h",
	ExdQualifiers_Minimum:                  "Minimum",
	ExdQualifiers_First:                    "First",
	ExdQualifiers_Second:                   "Second",
	ExdQualifiers_Third:                    "Third",
	ExdQualifiers_Shifter:                  "Shifter",
	ExdQualifiers_LastSport:                "LastSport",
	ExdQualifiers_Moving:                   "Moving",
	ExdQualifiers_Stopped:                  "Stopped",
	ExdQualifiers_EstimatedTotal:           "EstimatedTotal",
	ExdQualifiers_Zone9:                    "Zone9",
	ExdQualifiers_Zone8:                    "Zone8",
	ExdQualifiers_Zone7:                    "Zone7",
	ExdQualifiers_Zone6:                    "Zone6",
	ExdQualifiers_Zone5:                    "Zone5",
	ExdQualifiers_Zone4:                    "Zone4",
	ExdQualifiers_Zone3:                    "Zone3",
	ExdQualifiers_Zone2:                    "Zone2",
	ExdQualifiers_Zone1:                    "Zone1",
}

func (k ExdQualifiers) String() string {
	if uint(k) < uint(len(ExdQualifiersNames)) {
		if v, ok := ExdQualifiersNames[k]; ok {
			return v
		}
	}
	return "ExdQualifiers(" + strconv.Itoa(int(k)) + ")"
}

var ExdQualifiersValues = map[string]ExdQualifiers{
	"NoQualifier":              ExdQualifiers_NoQualifier,
	"Instantaneous":            ExdQualifiers_Instantaneous,
	"Average":                  ExdQualifiers_Average,
	"Lap":                      ExdQualifiers_Lap,
	"Maximum":                  ExdQualifiers_Maximum,
	"MaximumAverage":           ExdQualifiers_MaximumAverage,
	"MaximumLap":               ExdQualifiers_MaximumLap,
	"LastLap":                  ExdQualifiers_LastLap,
	"AverageLap":               ExdQualifiers_AverageLap,
	"ToDestination":            ExdQualifiers_ToDestination,
	"ToGo":                     ExdQualifiers_ToGo,
	"ToNext":                   ExdQualifiers_ToNext,
	"NextCoursePoint":          ExdQualifiers_NextCoursePoint,
	"Total":                    ExdQualifiers_Total,
	"ThreeSecondAverage":       ExdQualifiers_ThreeSecondAverage,
	"TenSecondAverage":         ExdQualifiers_TenSecondAverage,
	"ThirtySecondAverage":      ExdQualifiers_ThirtySecondAverage,
	"PercentMaximum":           ExdQualifiers_PercentMaximum,
	"PercentMaximumAverage":    ExdQualifiers_PercentMaximumAverage,
	"LapPercentMaximum":        ExdQualifiers_LapPercentMaximum,
	"Elapsed":                  ExdQualifiers_Elapsed,
	"Sunrise":                  ExdQualifiers_Sunrise,
	"Sunset":                   ExdQualifiers_Sunset,
	"ComparedToVirtualPartner": ExdQualifiers_ComparedToVirtualPartner,
	"Maximum24h":               ExdQualifiers_Maximum24h,
	"Minimum24h":               ExdQualifiers_Minimum24h,
	"Minimum":                  ExdQualifiers_Minimum,
	"First":                    ExdQualifiers_First,
	"Second":                   ExdQualifiers_Second,
	"Third":                    ExdQualifiers_Third,
	"Shifter":                  ExdQualifiers_Shifter,
	"LastSport":                ExdQualifiers_LastSport,
	"Moving":                   ExdQualifiers_Moving,
	"Stopped":                  ExdQualifiers_Stopped,
	"EstimatedTotal":           ExdQualifiers_EstimatedTotal,
	"Zone9":                    ExdQualifiers_Zone9,
	"Zone8":                    ExdQualifiers_Zone8,
	"Zone7":                    ExdQualifiers_Zone7,
	"Zone6":                    ExdQualifiers_Zone6,
	"Zone5":                    ExdQualifiers_Zone5,
	"Zone4":                    ExdQualifiers_Zone4,
	"Zone3":                    ExdQualifiers_Zone3,
	"Zone2":                    ExdQualifiers_Zone2,
	"Zone1":                    ExdQualifiers_Zone1,
}

func ExdQualifiersFromString(s string) ExdQualifiers {
	if v, ok := ExdQualifiersValues[s]; ok {
		return v
	}
	return ExdQualifiers(ExdQualifiers_Invalid)
}
