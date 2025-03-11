package fitDecl

import strconv "strconv"

type DisplayPosition byte

const (
	DisplayPosition_Degree               DisplayPosition = 0
	DisplayPosition_DegreeMinute         DisplayPosition = 1
	DisplayPosition_DegreeMinuteSecond   DisplayPosition = 2
	DisplayPosition_AustrianGrid         DisplayPosition = 3
	DisplayPosition_BritishGrid          DisplayPosition = 4
	DisplayPosition_DutchGrid            DisplayPosition = 5
	DisplayPosition_HungarianGrid        DisplayPosition = 6
	DisplayPosition_FinnishGrid          DisplayPosition = 7
	DisplayPosition_GermanGrid           DisplayPosition = 8
	DisplayPosition_IcelandicGrid        DisplayPosition = 9
	DisplayPosition_IndonesianEquatorial DisplayPosition = 10
	DisplayPosition_IndonesianIrian      DisplayPosition = 11
	DisplayPosition_IndonesianSouthern   DisplayPosition = 12
	DisplayPosition_IndiaZone0           DisplayPosition = 13
	DisplayPosition_IndiaZoneIA          DisplayPosition = 14
	DisplayPosition_IndiaZoneIB          DisplayPosition = 15
	DisplayPosition_IndiaZoneIIA         DisplayPosition = 16
	DisplayPosition_IndiaZoneIIB         DisplayPosition = 17
	DisplayPosition_IndiaZoneIIIA        DisplayPosition = 18
	DisplayPosition_IndiaZoneIIIB        DisplayPosition = 19
	DisplayPosition_IndiaZoneIVA         DisplayPosition = 20
	DisplayPosition_IndiaZoneIVB         DisplayPosition = 21
	DisplayPosition_IrishTransverse      DisplayPosition = 22
	DisplayPosition_IrishGrid            DisplayPosition = 23
	DisplayPosition_Loran                DisplayPosition = 24
	DisplayPosition_MaidenheadGrid       DisplayPosition = 25
	DisplayPosition_MgrsGrid             DisplayPosition = 26
	DisplayPosition_NewZealandGrid       DisplayPosition = 27
	DisplayPosition_NewZealandTransverse DisplayPosition = 28
	DisplayPosition_QatarGrid            DisplayPosition = 29
	DisplayPosition_ModifiedSwedishGrid  DisplayPosition = 30
	DisplayPosition_SwedishGrid          DisplayPosition = 31
	DisplayPosition_SouthAfricanGrid     DisplayPosition = 32
	DisplayPosition_SwissGrid            DisplayPosition = 33
	DisplayPosition_TaiwanGrid           DisplayPosition = 34
	DisplayPosition_UnitedStatesGrid     DisplayPosition = 35
	DisplayPosition_UtmUpsGrid           DisplayPosition = 36
	DisplayPosition_WestMalayan          DisplayPosition = 37
	DisplayPosition_BorneoRso            DisplayPosition = 38
	DisplayPosition_EstonianGrid         DisplayPosition = 39
	DisplayPosition_LatvianGrid          DisplayPosition = 40
	DisplayPosition_SwedishRef99Grid     DisplayPosition = 41
	DisplayPosition_Invalid              DisplayPosition = 255
)

var DisplayPositionNames = map[DisplayPosition]string{
	DisplayPosition_Degree:               "Degree",
	DisplayPosition_DegreeMinute:         "DegreeMinute",
	DisplayPosition_DegreeMinuteSecond:   "DegreeMinuteSecond",
	DisplayPosition_AustrianGrid:         "AustrianGrid",
	DisplayPosition_BritishGrid:          "BritishGrid",
	DisplayPosition_DutchGrid:            "DutchGrid",
	DisplayPosition_HungarianGrid:        "HungarianGrid",
	DisplayPosition_FinnishGrid:          "FinnishGrid",
	DisplayPosition_GermanGrid:           "GermanGrid",
	DisplayPosition_IcelandicGrid:        "IcelandicGrid",
	DisplayPosition_IndonesianEquatorial: "IndonesianEquatorial",
	DisplayPosition_IndonesianIrian:      "IndonesianIrian",
	DisplayPosition_IndonesianSouthern:   "IndonesianSouthern",
	DisplayPosition_IndiaZone0:           "IndiaZone0",
	DisplayPosition_IndiaZoneIA:          "IndiaZoneIA",
	DisplayPosition_IndiaZoneIB:          "IndiaZoneIB",
	DisplayPosition_IndiaZoneIIA:         "IndiaZoneIIA",
	DisplayPosition_IndiaZoneIIB:         "IndiaZoneIIB",
	DisplayPosition_IndiaZoneIIIA:        "IndiaZoneIIIA",
	DisplayPosition_IndiaZoneIIIB:        "IndiaZoneIIIB",
	DisplayPosition_IndiaZoneIVA:         "IndiaZoneIVA",
	DisplayPosition_IndiaZoneIVB:         "IndiaZoneIVB",
	DisplayPosition_IrishTransverse:      "IrishTransverse",
	DisplayPosition_IrishGrid:            "IrishGrid",
	DisplayPosition_Loran:                "Loran",
	DisplayPosition_MaidenheadGrid:       "MaidenheadGrid",
	DisplayPosition_MgrsGrid:             "MgrsGrid",
	DisplayPosition_NewZealandGrid:       "NewZealandGrid",
	DisplayPosition_NewZealandTransverse: "NewZealandTransverse",
	DisplayPosition_QatarGrid:            "QatarGrid",
	DisplayPosition_ModifiedSwedishGrid:  "ModifiedSwedishGrid",
	DisplayPosition_SwedishGrid:          "SwedishGrid",
	DisplayPosition_SouthAfricanGrid:     "SouthAfricanGrid",
	DisplayPosition_SwissGrid:            "SwissGrid",
	DisplayPosition_TaiwanGrid:           "TaiwanGrid",
	DisplayPosition_UnitedStatesGrid:     "UnitedStatesGrid",
	DisplayPosition_UtmUpsGrid:           "UtmUpsGrid",
	DisplayPosition_WestMalayan:          "WestMalayan",
	DisplayPosition_BorneoRso:            "BorneoRso",
	DisplayPosition_EstonianGrid:         "EstonianGrid",
	DisplayPosition_LatvianGrid:          "LatvianGrid",
	DisplayPosition_SwedishRef99Grid:     "SwedishRef99Grid",
}

func (k DisplayPosition) String() string {
	if uint(k) < uint(len(DisplayPositionNames)) {
		if v, ok := DisplayPositionNames[k]; ok {
			return v
		}
	}
	return "DisplayPosition(" + strconv.Itoa(int(k)) + ")"
}

var DisplayPositionValues = map[string]DisplayPosition{
	"Degree":               DisplayPosition_Degree,
	"DegreeMinute":         DisplayPosition_DegreeMinute,
	"DegreeMinuteSecond":   DisplayPosition_DegreeMinuteSecond,
	"AustrianGrid":         DisplayPosition_AustrianGrid,
	"BritishGrid":          DisplayPosition_BritishGrid,
	"DutchGrid":            DisplayPosition_DutchGrid,
	"HungarianGrid":        DisplayPosition_HungarianGrid,
	"FinnishGrid":          DisplayPosition_FinnishGrid,
	"GermanGrid":           DisplayPosition_GermanGrid,
	"IcelandicGrid":        DisplayPosition_IcelandicGrid,
	"IndonesianEquatorial": DisplayPosition_IndonesianEquatorial,
	"IndonesianIrian":      DisplayPosition_IndonesianIrian,
	"IndonesianSouthern":   DisplayPosition_IndonesianSouthern,
	"IndiaZone0":           DisplayPosition_IndiaZone0,
	"IndiaZoneIA":          DisplayPosition_IndiaZoneIA,
	"IndiaZoneIB":          DisplayPosition_IndiaZoneIB,
	"IndiaZoneIIA":         DisplayPosition_IndiaZoneIIA,
	"IndiaZoneIIB":         DisplayPosition_IndiaZoneIIB,
	"IndiaZoneIIIA":        DisplayPosition_IndiaZoneIIIA,
	"IndiaZoneIIIB":        DisplayPosition_IndiaZoneIIIB,
	"IndiaZoneIVA":         DisplayPosition_IndiaZoneIVA,
	"IndiaZoneIVB":         DisplayPosition_IndiaZoneIVB,
	"IrishTransverse":      DisplayPosition_IrishTransverse,
	"IrishGrid":            DisplayPosition_IrishGrid,
	"Loran":                DisplayPosition_Loran,
	"MaidenheadGrid":       DisplayPosition_MaidenheadGrid,
	"MgrsGrid":             DisplayPosition_MgrsGrid,
	"NewZealandGrid":       DisplayPosition_NewZealandGrid,
	"NewZealandTransverse": DisplayPosition_NewZealandTransverse,
	"QatarGrid":            DisplayPosition_QatarGrid,
	"ModifiedSwedishGrid":  DisplayPosition_ModifiedSwedishGrid,
	"SwedishGrid":          DisplayPosition_SwedishGrid,
	"SouthAfricanGrid":     DisplayPosition_SouthAfricanGrid,
	"SwissGrid":            DisplayPosition_SwissGrid,
	"TaiwanGrid":           DisplayPosition_TaiwanGrid,
	"UnitedStatesGrid":     DisplayPosition_UnitedStatesGrid,
	"UtmUpsGrid":           DisplayPosition_UtmUpsGrid,
	"WestMalayan":          DisplayPosition_WestMalayan,
	"BorneoRso":            DisplayPosition_BorneoRso,
	"EstonianGrid":         DisplayPosition_EstonianGrid,
	"LatvianGrid":          DisplayPosition_LatvianGrid,
	"SwedishRef99Grid":     DisplayPosition_SwedishRef99Grid,
}

func DisplayPositionFromString(s string) DisplayPosition {
	if v, ok := DisplayPositionValues[s]; ok {
		return v
	}
	return DisplayPosition(DisplayPosition_Invalid)
}
