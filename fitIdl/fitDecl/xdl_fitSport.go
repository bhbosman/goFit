package fitDecl

import strconv "strconv"

type Sport byte

const (
	Sport_Generic               Sport = 0
	Sport_Running               Sport = 1
	Sport_Cycling               Sport = 2
	Sport_Transition            Sport = 3
	Sport_FitnessEquipment      Sport = 4
	Sport_Swimming              Sport = 5
	Sport_Basketball            Sport = 6
	Sport_Soccer                Sport = 7
	Sport_Tennis                Sport = 8
	Sport_AmericanFootball      Sport = 9
	Sport_Training              Sport = 10
	Sport_Walking               Sport = 11
	Sport_CrossCountrySkiing    Sport = 12
	Sport_AlpineSkiing          Sport = 13
	Sport_Snowboarding          Sport = 14
	Sport_Rowing                Sport = 15
	Sport_Mountaineering        Sport = 16
	Sport_Hiking                Sport = 17
	Sport_Multisport            Sport = 18
	Sport_Paddling              Sport = 19
	Sport_Flying                Sport = 20
	Sport_EBiking               Sport = 21
	Sport_Motorcycling          Sport = 22
	Sport_Boating               Sport = 23
	Sport_Driving               Sport = 24
	Sport_Golf                  Sport = 25
	Sport_HangGliding           Sport = 26
	Sport_HorsebackRiding       Sport = 27
	Sport_Hunting               Sport = 28
	Sport_Fishing               Sport = 29
	Sport_InlineSkating         Sport = 30
	Sport_RockClimbing          Sport = 31
	Sport_Sailing               Sport = 32
	Sport_IceSkating            Sport = 33
	Sport_SkyDiving             Sport = 34
	Sport_Snowshoeing           Sport = 35
	Sport_Snowmobiling          Sport = 36
	Sport_StandUpPaddleboarding Sport = 37
	Sport_Surfing               Sport = 38
	Sport_Wakeboarding          Sport = 39
	Sport_WaterSkiing           Sport = 40
	Sport_Kayaking              Sport = 41
	Sport_Rafting               Sport = 42
	Sport_Windsurfing           Sport = 43
	Sport_Kitesurfing           Sport = 44
	Sport_Tactical              Sport = 45
	Sport_Jumpmaster            Sport = 46
	Sport_Boxing                Sport = 47
	Sport_FloorClimbing         Sport = 48
	Sport_Baseball              Sport = 49
	Sport_Diving                Sport = 53
	Sport_Hiit                  Sport = 62
	Sport_Racket                Sport = 64
	Sport_WheelchairPushWalk    Sport = 65
	Sport_WheelchairPushRun     Sport = 66
	Sport_Meditation            Sport = 67
	Sport_DiscGolf              Sport = 69
	Sport_Cricket               Sport = 71
	Sport_Rugby                 Sport = 72
	Sport_Hockey                Sport = 73
	Sport_Lacrosse              Sport = 74
	Sport_Volleyball            Sport = 75
	Sport_WaterTubing           Sport = 76
	Sport_Wakesurfing           Sport = 77
	Sport_MixedMartialArts      Sport = 80
	Sport_Snorkeling            Sport = 82
	Sport_Dance                 Sport = 83
	Sport_JumpRope              Sport = 84
	Sport_All                   Sport = 254
	Sport_Invalid               Sport = 255
)

var SportNames = map[Sport]string{
	Sport_Generic:               "Generic",
	Sport_Running:               "Running",
	Sport_Cycling:               "Cycling",
	Sport_Transition:            "Transition",
	Sport_FitnessEquipment:      "FitnessEquipment",
	Sport_Swimming:              "Swimming",
	Sport_Basketball:            "Basketball",
	Sport_Soccer:                "Soccer",
	Sport_Tennis:                "Tennis",
	Sport_AmericanFootball:      "AmericanFootball",
	Sport_Training:              "Training",
	Sport_Walking:               "Walking",
	Sport_CrossCountrySkiing:    "CrossCountrySkiing",
	Sport_AlpineSkiing:          "AlpineSkiing",
	Sport_Snowboarding:          "Snowboarding",
	Sport_Rowing:                "Rowing",
	Sport_Mountaineering:        "Mountaineering",
	Sport_Hiking:                "Hiking",
	Sport_Multisport:            "Multisport",
	Sport_Paddling:              "Paddling",
	Sport_Flying:                "Flying",
	Sport_EBiking:               "EBiking",
	Sport_Motorcycling:          "Motorcycling",
	Sport_Boating:               "Boating",
	Sport_Driving:               "Driving",
	Sport_Golf:                  "Golf",
	Sport_HangGliding:           "HangGliding",
	Sport_HorsebackRiding:       "HorsebackRiding",
	Sport_Hunting:               "Hunting",
	Sport_Fishing:               "Fishing",
	Sport_InlineSkating:         "InlineSkating",
	Sport_RockClimbing:          "RockClimbing",
	Sport_Sailing:               "Sailing",
	Sport_IceSkating:            "IceSkating",
	Sport_SkyDiving:             "SkyDiving",
	Sport_Snowshoeing:           "Snowshoeing",
	Sport_Snowmobiling:          "Snowmobiling",
	Sport_StandUpPaddleboarding: "StandUpPaddleboarding",
	Sport_Surfing:               "Surfing",
	Sport_Wakeboarding:          "Wakeboarding",
	Sport_WaterSkiing:           "WaterSkiing",
	Sport_Kayaking:              "Kayaking",
	Sport_Rafting:               "Rafting",
	Sport_Windsurfing:           "Windsurfing",
	Sport_Kitesurfing:           "Kitesurfing",
	Sport_Tactical:              "Tactical",
	Sport_Jumpmaster:            "Jumpmaster",
	Sport_Boxing:                "Boxing",
	Sport_FloorClimbing:         "FloorClimbing",
	Sport_Baseball:              "Baseball",
	Sport_Diving:                "Diving",
	Sport_Hiit:                  "Hiit",
	Sport_Racket:                "Racket",
	Sport_WheelchairPushWalk:    "WheelchairPushWalk",
	Sport_WheelchairPushRun:     "WheelchairPushRun",
	Sport_Meditation:            "Meditation",
	Sport_DiscGolf:              "DiscGolf",
	Sport_Cricket:               "Cricket",
	Sport_Rugby:                 "Rugby",
	Sport_Hockey:                "Hockey",
	Sport_Lacrosse:              "Lacrosse",
	Sport_Volleyball:            "Volleyball",
	Sport_WaterTubing:           "WaterTubing",
	Sport_Wakesurfing:           "Wakesurfing",
	Sport_MixedMartialArts:      "MixedMartialArts",
	Sport_Snorkeling:            "Snorkeling",
	Sport_Dance:                 "Dance",
	Sport_JumpRope:              "JumpRope",
	Sport_All:                   "All",
}

func (k Sport) String() string {
	if uint(k) < uint(len(SportNames)) {
		if v, ok := SportNames[k]; ok {
			return v
		}
	}
	return "Sport(" + strconv.Itoa(int(k)) + ")"
}

var SportValues = map[string]Sport{
	"Generic":               Sport_Generic,
	"Running":               Sport_Running,
	"Cycling":               Sport_Cycling,
	"Transition":            Sport_Transition,
	"FitnessEquipment":      Sport_FitnessEquipment,
	"Swimming":              Sport_Swimming,
	"Basketball":            Sport_Basketball,
	"Soccer":                Sport_Soccer,
	"Tennis":                Sport_Tennis,
	"AmericanFootball":      Sport_AmericanFootball,
	"Training":              Sport_Training,
	"Walking":               Sport_Walking,
	"CrossCountrySkiing":    Sport_CrossCountrySkiing,
	"AlpineSkiing":          Sport_AlpineSkiing,
	"Snowboarding":          Sport_Snowboarding,
	"Rowing":                Sport_Rowing,
	"Mountaineering":        Sport_Mountaineering,
	"Hiking":                Sport_Hiking,
	"Multisport":            Sport_Multisport,
	"Paddling":              Sport_Paddling,
	"Flying":                Sport_Flying,
	"EBiking":               Sport_EBiking,
	"Motorcycling":          Sport_Motorcycling,
	"Boating":               Sport_Boating,
	"Driving":               Sport_Driving,
	"Golf":                  Sport_Golf,
	"HangGliding":           Sport_HangGliding,
	"HorsebackRiding":       Sport_HorsebackRiding,
	"Hunting":               Sport_Hunting,
	"Fishing":               Sport_Fishing,
	"InlineSkating":         Sport_InlineSkating,
	"RockClimbing":          Sport_RockClimbing,
	"Sailing":               Sport_Sailing,
	"IceSkating":            Sport_IceSkating,
	"SkyDiving":             Sport_SkyDiving,
	"Snowshoeing":           Sport_Snowshoeing,
	"Snowmobiling":          Sport_Snowmobiling,
	"StandUpPaddleboarding": Sport_StandUpPaddleboarding,
	"Surfing":               Sport_Surfing,
	"Wakeboarding":          Sport_Wakeboarding,
	"WaterSkiing":           Sport_WaterSkiing,
	"Kayaking":              Sport_Kayaking,
	"Rafting":               Sport_Rafting,
	"Windsurfing":           Sport_Windsurfing,
	"Kitesurfing":           Sport_Kitesurfing,
	"Tactical":              Sport_Tactical,
	"Jumpmaster":            Sport_Jumpmaster,
	"Boxing":                Sport_Boxing,
	"FloorClimbing":         Sport_FloorClimbing,
	"Baseball":              Sport_Baseball,
	"Diving":                Sport_Diving,
	"Hiit":                  Sport_Hiit,
	"Racket":                Sport_Racket,
	"WheelchairPushWalk":    Sport_WheelchairPushWalk,
	"WheelchairPushRun":     Sport_WheelchairPushRun,
	"Meditation":            Sport_Meditation,
	"DiscGolf":              Sport_DiscGolf,
	"Cricket":               Sport_Cricket,
	"Rugby":                 Sport_Rugby,
	"Hockey":                Sport_Hockey,
	"Lacrosse":              Sport_Lacrosse,
	"Volleyball":            Sport_Volleyball,
	"WaterTubing":           Sport_WaterTubing,
	"Wakesurfing":           Sport_Wakesurfing,
	"MixedMartialArts":      Sport_MixedMartialArts,
	"Snorkeling":            Sport_Snorkeling,
	"Dance":                 Sport_Dance,
	"JumpRope":              Sport_JumpRope,
	"All":                   Sport_All,
}

func SportFromString(s string) Sport {
	if v, ok := SportValues[s]; ok {
		return v
	}
	return Sport(Sport_Invalid)
}
