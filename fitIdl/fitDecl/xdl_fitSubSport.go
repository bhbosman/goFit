package fitDecl

import strconv "strconv"

type SubSport byte

const (
	SubSport_Generic              SubSport = 0
	SubSport_Treadmill            SubSport = 1
	SubSport_Street               SubSport = 2
	SubSport_Trail                SubSport = 3
	SubSport_Track                SubSport = 4
	SubSport_Spin                 SubSport = 5
	SubSport_IndoorCycling        SubSport = 6
	SubSport_Road                 SubSport = 7
	SubSport_Mountain             SubSport = 8
	SubSport_Downhill             SubSport = 9
	SubSport_Recumbent            SubSport = 10
	SubSport_Cyclocross           SubSport = 11
	SubSport_HandCycling          SubSport = 12
	SubSport_TrackCycling         SubSport = 13
	SubSport_IndoorRowing         SubSport = 14
	SubSport_Elliptical           SubSport = 15
	SubSport_StairClimbing        SubSport = 16
	SubSport_LapSwimming          SubSport = 17
	SubSport_OpenWater            SubSport = 18
	SubSport_FlexibilityTraining  SubSport = 19
	SubSport_StrengthTraining     SubSport = 20
	SubSport_WarmUp               SubSport = 21
	SubSport_Match                SubSport = 22
	SubSport_Exercise             SubSport = 23
	SubSport_Challenge            SubSport = 24
	SubSport_IndoorSkiing         SubSport = 25
	SubSport_CardioTraining       SubSport = 26
	SubSport_IndoorWalking        SubSport = 27
	SubSport_EBikeFitness         SubSport = 28
	SubSport_Bmx                  SubSport = 29
	SubSport_CasualWalking        SubSport = 30
	SubSport_SpeedWalking         SubSport = 31
	SubSport_BikeToRunTransition  SubSport = 32
	SubSport_RunToBikeTransition  SubSport = 33
	SubSport_SwimToBikeTransition SubSport = 34
	SubSport_Atv                  SubSport = 35
	SubSport_Motocross            SubSport = 36
	SubSport_Backcountry          SubSport = 37
	SubSport_Resort               SubSport = 38
	SubSport_RcDrone              SubSport = 39
	SubSport_Wingsuit             SubSport = 40
	SubSport_Whitewater           SubSport = 41
	SubSport_SkateSkiing          SubSport = 42
	SubSport_Yoga                 SubSport = 43
	SubSport_Pilates              SubSport = 44
	SubSport_IndoorRunning        SubSport = 45
	SubSport_GravelCycling        SubSport = 46
	SubSport_EBikeMountain        SubSport = 47
	SubSport_Commuting            SubSport = 48
	SubSport_MixedSurface         SubSport = 49
	SubSport_Navigate             SubSport = 50
	SubSport_TrackMe              SubSport = 51
	SubSport_Map                  SubSport = 52
	SubSport_SingleGasDiving      SubSport = 53
	SubSport_MultiGasDiving       SubSport = 54
	SubSport_GaugeDiving          SubSport = 55
	SubSport_ApneaDiving          SubSport = 56
	SubSport_ApneaHunting         SubSport = 57
	SubSport_VirtualActivity      SubSport = 58
	SubSport_Obstacle             SubSport = 59
	SubSport_Breathing            SubSport = 62
	SubSport_SailRace             SubSport = 65
	SubSport_Ultra                SubSport = 67
	SubSport_IndoorClimbing       SubSport = 68
	SubSport_Bouldering           SubSport = 69
	SubSport_Hiit                 SubSport = 70
	SubSport_Amrap                SubSport = 73
	SubSport_Emom                 SubSport = 74
	SubSport_Tabata               SubSport = 75
	SubSport_Pickleball           SubSport = 84
	SubSport_Padel                SubSport = 85
	SubSport_IndoorWheelchairWalk SubSport = 86
	SubSport_IndoorWheelchairRun  SubSport = 87
	SubSport_IndoorHandCycling    SubSport = 88
	SubSport_Squash               SubSport = 94
	SubSport_Badminton            SubSport = 95
	SubSport_Racquetball          SubSport = 96
	SubSport_TableTennis          SubSport = 97
	SubSport_FlyCanopy            SubSport = 110
	SubSport_FlyParaglide         SubSport = 111
	SubSport_FlyParamotor         SubSport = 112
	SubSport_FlyPressurized       SubSport = 113
	SubSport_FlyNavigate          SubSport = 114
	SubSport_FlyTimer             SubSport = 115
	SubSport_FlyAltimeter         SubSport = 116
	SubSport_FlyWx                SubSport = 117
	SubSport_FlyVfr               SubSport = 118
	SubSport_FlyIfr               SubSport = 119
	SubSport_All                  SubSport = 254
	SubSport_Invalid              SubSport = 255
)

var SubSportNames = map[SubSport]string{
	SubSport_Generic:              "Generic",
	SubSport_Treadmill:            "Treadmill",
	SubSport_Street:               "Street",
	SubSport_Trail:                "Trail",
	SubSport_Track:                "Track",
	SubSport_Spin:                 "Spin",
	SubSport_IndoorCycling:        "IndoorCycling",
	SubSport_Road:                 "Road",
	SubSport_Mountain:             "Mountain",
	SubSport_Downhill:             "Downhill",
	SubSport_Recumbent:            "Recumbent",
	SubSport_Cyclocross:           "Cyclocross",
	SubSport_HandCycling:          "HandCycling",
	SubSport_TrackCycling:         "TrackCycling",
	SubSport_IndoorRowing:         "IndoorRowing",
	SubSport_Elliptical:           "Elliptical",
	SubSport_StairClimbing:        "StairClimbing",
	SubSport_LapSwimming:          "LapSwimming",
	SubSport_OpenWater:            "OpenWater",
	SubSport_FlexibilityTraining:  "FlexibilityTraining",
	SubSport_StrengthTraining:     "StrengthTraining",
	SubSport_WarmUp:               "WarmUp",
	SubSport_Match:                "Match",
	SubSport_Exercise:             "Exercise",
	SubSport_Challenge:            "Challenge",
	SubSport_IndoorSkiing:         "IndoorSkiing",
	SubSport_CardioTraining:       "CardioTraining",
	SubSport_IndoorWalking:        "IndoorWalking",
	SubSport_EBikeFitness:         "EBikeFitness",
	SubSport_Bmx:                  "Bmx",
	SubSport_CasualWalking:        "CasualWalking",
	SubSport_SpeedWalking:         "SpeedWalking",
	SubSport_BikeToRunTransition:  "BikeToRunTransition",
	SubSport_RunToBikeTransition:  "RunToBikeTransition",
	SubSport_SwimToBikeTransition: "SwimToBikeTransition",
	SubSport_Atv:                  "Atv",
	SubSport_Motocross:            "Motocross",
	SubSport_Backcountry:          "Backcountry",
	SubSport_Resort:               "Resort",
	SubSport_RcDrone:              "RcDrone",
	SubSport_Wingsuit:             "Wingsuit",
	SubSport_Whitewater:           "Whitewater",
	SubSport_SkateSkiing:          "SkateSkiing",
	SubSport_Yoga:                 "Yoga",
	SubSport_Pilates:              "Pilates",
	SubSport_IndoorRunning:        "IndoorRunning",
	SubSport_GravelCycling:        "GravelCycling",
	SubSport_EBikeMountain:        "EBikeMountain",
	SubSport_Commuting:            "Commuting",
	SubSport_MixedSurface:         "MixedSurface",
	SubSport_Navigate:             "Navigate",
	SubSport_TrackMe:              "TrackMe",
	SubSport_Map:                  "Map",
	SubSport_SingleGasDiving:      "SingleGasDiving",
	SubSport_MultiGasDiving:       "MultiGasDiving",
	SubSport_GaugeDiving:          "GaugeDiving",
	SubSport_ApneaDiving:          "ApneaDiving",
	SubSport_ApneaHunting:         "ApneaHunting",
	SubSport_VirtualActivity:      "VirtualActivity",
	SubSport_Obstacle:             "Obstacle",
	SubSport_Breathing:            "Breathing",
	SubSport_SailRace:             "SailRace",
	SubSport_Ultra:                "Ultra",
	SubSport_IndoorClimbing:       "IndoorClimbing",
	SubSport_Bouldering:           "Bouldering",
	SubSport_Hiit:                 "Hiit",
	SubSport_Amrap:                "Amrap",
	SubSport_Emom:                 "Emom",
	SubSport_Tabata:               "Tabata",
	SubSport_Pickleball:           "Pickleball",
	SubSport_Padel:                "Padel",
	SubSport_IndoorWheelchairWalk: "IndoorWheelchairWalk",
	SubSport_IndoorWheelchairRun:  "IndoorWheelchairRun",
	SubSport_IndoorHandCycling:    "IndoorHandCycling",
	SubSport_Squash:               "Squash",
	SubSport_Badminton:            "Badminton",
	SubSport_Racquetball:          "Racquetball",
	SubSport_TableTennis:          "TableTennis",
	SubSport_FlyCanopy:            "FlyCanopy",
	SubSport_FlyParaglide:         "FlyParaglide",
	SubSport_FlyParamotor:         "FlyParamotor",
	SubSport_FlyPressurized:       "FlyPressurized",
	SubSport_FlyNavigate:          "FlyNavigate",
	SubSport_FlyTimer:             "FlyTimer",
	SubSport_FlyAltimeter:         "FlyAltimeter",
	SubSport_FlyWx:                "FlyWx",
	SubSport_FlyVfr:               "FlyVfr",
	SubSport_FlyIfr:               "FlyIfr",
	SubSport_All:                  "All",
}

func (k SubSport) String() string {
	if uint(k) < uint(len(SubSportNames)) {
		if v, ok := SubSportNames[k]; ok {
			return v
		}
	}
	return "SubSport(" + strconv.Itoa(int(k)) + ")"
}

var SubSportValues = map[string]SubSport{
	"Generic":              SubSport_Generic,
	"Treadmill":            SubSport_Treadmill,
	"Street":               SubSport_Street,
	"Trail":                SubSport_Trail,
	"Track":                SubSport_Track,
	"Spin":                 SubSport_Spin,
	"IndoorCycling":        SubSport_IndoorCycling,
	"Road":                 SubSport_Road,
	"Mountain":             SubSport_Mountain,
	"Downhill":             SubSport_Downhill,
	"Recumbent":            SubSport_Recumbent,
	"Cyclocross":           SubSport_Cyclocross,
	"HandCycling":          SubSport_HandCycling,
	"TrackCycling":         SubSport_TrackCycling,
	"IndoorRowing":         SubSport_IndoorRowing,
	"Elliptical":           SubSport_Elliptical,
	"StairClimbing":        SubSport_StairClimbing,
	"LapSwimming":          SubSport_LapSwimming,
	"OpenWater":            SubSport_OpenWater,
	"FlexibilityTraining":  SubSport_FlexibilityTraining,
	"StrengthTraining":     SubSport_StrengthTraining,
	"WarmUp":               SubSport_WarmUp,
	"Match":                SubSport_Match,
	"Exercise":             SubSport_Exercise,
	"Challenge":            SubSport_Challenge,
	"IndoorSkiing":         SubSport_IndoorSkiing,
	"CardioTraining":       SubSport_CardioTraining,
	"IndoorWalking":        SubSport_IndoorWalking,
	"EBikeFitness":         SubSport_EBikeFitness,
	"Bmx":                  SubSport_Bmx,
	"CasualWalking":        SubSport_CasualWalking,
	"SpeedWalking":         SubSport_SpeedWalking,
	"BikeToRunTransition":  SubSport_BikeToRunTransition,
	"RunToBikeTransition":  SubSport_RunToBikeTransition,
	"SwimToBikeTransition": SubSport_SwimToBikeTransition,
	"Atv":                  SubSport_Atv,
	"Motocross":            SubSport_Motocross,
	"Backcountry":          SubSport_Backcountry,
	"Resort":               SubSport_Resort,
	"RcDrone":              SubSport_RcDrone,
	"Wingsuit":             SubSport_Wingsuit,
	"Whitewater":           SubSport_Whitewater,
	"SkateSkiing":          SubSport_SkateSkiing,
	"Yoga":                 SubSport_Yoga,
	"Pilates":              SubSport_Pilates,
	"IndoorRunning":        SubSport_IndoorRunning,
	"GravelCycling":        SubSport_GravelCycling,
	"EBikeMountain":        SubSport_EBikeMountain,
	"Commuting":            SubSport_Commuting,
	"MixedSurface":         SubSport_MixedSurface,
	"Navigate":             SubSport_Navigate,
	"TrackMe":              SubSport_TrackMe,
	"Map":                  SubSport_Map,
	"SingleGasDiving":      SubSport_SingleGasDiving,
	"MultiGasDiving":       SubSport_MultiGasDiving,
	"GaugeDiving":          SubSport_GaugeDiving,
	"ApneaDiving":          SubSport_ApneaDiving,
	"ApneaHunting":         SubSport_ApneaHunting,
	"VirtualActivity":      SubSport_VirtualActivity,
	"Obstacle":             SubSport_Obstacle,
	"Breathing":            SubSport_Breathing,
	"SailRace":             SubSport_SailRace,
	"Ultra":                SubSport_Ultra,
	"IndoorClimbing":       SubSport_IndoorClimbing,
	"Bouldering":           SubSport_Bouldering,
	"Hiit":                 SubSport_Hiit,
	"Amrap":                SubSport_Amrap,
	"Emom":                 SubSport_Emom,
	"Tabata":               SubSport_Tabata,
	"Pickleball":           SubSport_Pickleball,
	"Padel":                SubSport_Padel,
	"IndoorWheelchairWalk": SubSport_IndoorWheelchairWalk,
	"IndoorWheelchairRun":  SubSport_IndoorWheelchairRun,
	"IndoorHandCycling":    SubSport_IndoorHandCycling,
	"Squash":               SubSport_Squash,
	"Badminton":            SubSport_Badminton,
	"Racquetball":          SubSport_Racquetball,
	"TableTennis":          SubSport_TableTennis,
	"FlyCanopy":            SubSport_FlyCanopy,
	"FlyParaglide":         SubSport_FlyParaglide,
	"FlyParamotor":         SubSport_FlyParamotor,
	"FlyPressurized":       SubSport_FlyPressurized,
	"FlyNavigate":          SubSport_FlyNavigate,
	"FlyTimer":             SubSport_FlyTimer,
	"FlyAltimeter":         SubSport_FlyAltimeter,
	"FlyWx":                SubSport_FlyWx,
	"FlyVfr":               SubSport_FlyVfr,
	"FlyIfr":               SubSport_FlyIfr,
	"All":                  SubSport_All,
}

func SubSportFromString(s string) SubSport {
	if v, ok := SubSportValues[s]; ok {
		return v
	}
	return SubSport(SubSport_Invalid)
}
