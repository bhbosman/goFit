package fitDecl

import strconv "strconv"

type TimeZone byte

const (
	TimeZone_Almaty                   TimeZone = 0
	TimeZone_Bangkok                  TimeZone = 1
	TimeZone_Bombay                   TimeZone = 2
	TimeZone_Brasilia                 TimeZone = 3
	TimeZone_Cairo                    TimeZone = 4
	TimeZone_CapeVerdeIs              TimeZone = 5
	TimeZone_Darwin                   TimeZone = 6
	TimeZone_Eniwetok                 TimeZone = 7
	TimeZone_Fiji                     TimeZone = 8
	TimeZone_HongKong                 TimeZone = 9
	TimeZone_Islamabad                TimeZone = 10
	TimeZone_Kabul                    TimeZone = 11
	TimeZone_Magadan                  TimeZone = 12
	TimeZone_MidAtlantic              TimeZone = 13
	TimeZone_Moscow                   TimeZone = 14
	TimeZone_Muscat                   TimeZone = 15
	TimeZone_Newfoundland             TimeZone = 16
	TimeZone_Samoa                    TimeZone = 17
	TimeZone_Sydney                   TimeZone = 18
	TimeZone_Tehran                   TimeZone = 19
	TimeZone_Tokyo                    TimeZone = 20
	TimeZone_UsAlaska                 TimeZone = 21
	TimeZone_UsAtlantic               TimeZone = 22
	TimeZone_UsCentral                TimeZone = 23
	TimeZone_UsEastern                TimeZone = 24
	TimeZone_UsHawaii                 TimeZone = 25
	TimeZone_UsMountain               TimeZone = 26
	TimeZone_UsPacific                TimeZone = 27
	TimeZone_Other                    TimeZone = 28
	TimeZone_Auckland                 TimeZone = 29
	TimeZone_Kathmandu                TimeZone = 30
	TimeZone_EuropeWesternWet         TimeZone = 31
	TimeZone_EuropeCentralCet         TimeZone = 32
	TimeZone_EuropeEasternEet         TimeZone = 33
	TimeZone_Jakarta                  TimeZone = 34
	TimeZone_Perth                    TimeZone = 35
	TimeZone_Adelaide                 TimeZone = 36
	TimeZone_Brisbane                 TimeZone = 37
	TimeZone_Tasmania                 TimeZone = 38
	TimeZone_Iceland                  TimeZone = 39
	TimeZone_Amsterdam                TimeZone = 40
	TimeZone_Athens                   TimeZone = 41
	TimeZone_Barcelona                TimeZone = 42
	TimeZone_Berlin                   TimeZone = 43
	TimeZone_Brussels                 TimeZone = 44
	TimeZone_Budapest                 TimeZone = 45
	TimeZone_Copenhagen               TimeZone = 46
	TimeZone_Dublin                   TimeZone = 47
	TimeZone_Helsinki                 TimeZone = 48
	TimeZone_Lisbon                   TimeZone = 49
	TimeZone_London                   TimeZone = 50
	TimeZone_Madrid                   TimeZone = 51
	TimeZone_Munich                   TimeZone = 52
	TimeZone_Oslo                     TimeZone = 53
	TimeZone_Paris                    TimeZone = 54
	TimeZone_Prague                   TimeZone = 55
	TimeZone_Reykjavik                TimeZone = 56
	TimeZone_Rome                     TimeZone = 57
	TimeZone_Stockholm                TimeZone = 58
	TimeZone_Vienna                   TimeZone = 59
	TimeZone_Warsaw                   TimeZone = 60
	TimeZone_Zurich                   TimeZone = 61
	TimeZone_Quebec                   TimeZone = 62
	TimeZone_Ontario                  TimeZone = 63
	TimeZone_Manitoba                 TimeZone = 64
	TimeZone_Saskatchewan             TimeZone = 65
	TimeZone_Alberta                  TimeZone = 66
	TimeZone_BritishColumbia          TimeZone = 67
	TimeZone_Boise                    TimeZone = 68
	TimeZone_Boston                   TimeZone = 69
	TimeZone_Chicago                  TimeZone = 70
	TimeZone_Dallas                   TimeZone = 71
	TimeZone_Denver                   TimeZone = 72
	TimeZone_KansasCity               TimeZone = 73
	TimeZone_LasVegas                 TimeZone = 74
	TimeZone_LosAngeles               TimeZone = 75
	TimeZone_Miami                    TimeZone = 76
	TimeZone_Minneapolis              TimeZone = 77
	TimeZone_NewYork                  TimeZone = 78
	TimeZone_NewOrleans               TimeZone = 79
	TimeZone_Phoenix                  TimeZone = 80
	TimeZone_SantaFe                  TimeZone = 81
	TimeZone_Seattle                  TimeZone = 82
	TimeZone_WashingtonDc             TimeZone = 83
	TimeZone_UsArizona                TimeZone = 84
	TimeZone_Chita                    TimeZone = 85
	TimeZone_Ekaterinburg             TimeZone = 86
	TimeZone_Irkutsk                  TimeZone = 87
	TimeZone_Kaliningrad              TimeZone = 88
	TimeZone_Krasnoyarsk              TimeZone = 89
	TimeZone_Novosibirsk              TimeZone = 90
	TimeZone_PetropavlovskKamchatskiy TimeZone = 91
	TimeZone_Samara                   TimeZone = 92
	TimeZone_Vladivostok              TimeZone = 93
	TimeZone_MexicoCentral            TimeZone = 94
	TimeZone_MexicoMountain           TimeZone = 95
	TimeZone_MexicoPacific            TimeZone = 96
	TimeZone_CapeTown                 TimeZone = 97
	TimeZone_Winkhoek                 TimeZone = 98
	TimeZone_Lagos                    TimeZone = 99
	TimeZone_Riyahd                   TimeZone = 100
	TimeZone_Venezuela                TimeZone = 101
	TimeZone_AustraliaLh              TimeZone = 102
	TimeZone_Santiago                 TimeZone = 103
	TimeZone_Manual                   TimeZone = 253
	TimeZone_Automatic                TimeZone = 254
	TimeZone_Invalid                  TimeZone = 255
)

var TimeZoneNames = map[TimeZone]string{
	TimeZone_Almaty:                   "Almaty",
	TimeZone_Bangkok:                  "Bangkok",
	TimeZone_Bombay:                   "Bombay",
	TimeZone_Brasilia:                 "Brasilia",
	TimeZone_Cairo:                    "Cairo",
	TimeZone_CapeVerdeIs:              "CapeVerdeIs",
	TimeZone_Darwin:                   "Darwin",
	TimeZone_Eniwetok:                 "Eniwetok",
	TimeZone_Fiji:                     "Fiji",
	TimeZone_HongKong:                 "HongKong",
	TimeZone_Islamabad:                "Islamabad",
	TimeZone_Kabul:                    "Kabul",
	TimeZone_Magadan:                  "Magadan",
	TimeZone_MidAtlantic:              "MidAtlantic",
	TimeZone_Moscow:                   "Moscow",
	TimeZone_Muscat:                   "Muscat",
	TimeZone_Newfoundland:             "Newfoundland",
	TimeZone_Samoa:                    "Samoa",
	TimeZone_Sydney:                   "Sydney",
	TimeZone_Tehran:                   "Tehran",
	TimeZone_Tokyo:                    "Tokyo",
	TimeZone_UsAlaska:                 "UsAlaska",
	TimeZone_UsAtlantic:               "UsAtlantic",
	TimeZone_UsCentral:                "UsCentral",
	TimeZone_UsEastern:                "UsEastern",
	TimeZone_UsHawaii:                 "UsHawaii",
	TimeZone_UsMountain:               "UsMountain",
	TimeZone_UsPacific:                "UsPacific",
	TimeZone_Other:                    "Other",
	TimeZone_Auckland:                 "Auckland",
	TimeZone_Kathmandu:                "Kathmandu",
	TimeZone_EuropeWesternWet:         "EuropeWesternWet",
	TimeZone_EuropeCentralCet:         "EuropeCentralCet",
	TimeZone_EuropeEasternEet:         "EuropeEasternEet",
	TimeZone_Jakarta:                  "Jakarta",
	TimeZone_Perth:                    "Perth",
	TimeZone_Adelaide:                 "Adelaide",
	TimeZone_Brisbane:                 "Brisbane",
	TimeZone_Tasmania:                 "Tasmania",
	TimeZone_Iceland:                  "Iceland",
	TimeZone_Amsterdam:                "Amsterdam",
	TimeZone_Athens:                   "Athens",
	TimeZone_Barcelona:                "Barcelona",
	TimeZone_Berlin:                   "Berlin",
	TimeZone_Brussels:                 "Brussels",
	TimeZone_Budapest:                 "Budapest",
	TimeZone_Copenhagen:               "Copenhagen",
	TimeZone_Dublin:                   "Dublin",
	TimeZone_Helsinki:                 "Helsinki",
	TimeZone_Lisbon:                   "Lisbon",
	TimeZone_London:                   "London",
	TimeZone_Madrid:                   "Madrid",
	TimeZone_Munich:                   "Munich",
	TimeZone_Oslo:                     "Oslo",
	TimeZone_Paris:                    "Paris",
	TimeZone_Prague:                   "Prague",
	TimeZone_Reykjavik:                "Reykjavik",
	TimeZone_Rome:                     "Rome",
	TimeZone_Stockholm:                "Stockholm",
	TimeZone_Vienna:                   "Vienna",
	TimeZone_Warsaw:                   "Warsaw",
	TimeZone_Zurich:                   "Zurich",
	TimeZone_Quebec:                   "Quebec",
	TimeZone_Ontario:                  "Ontario",
	TimeZone_Manitoba:                 "Manitoba",
	TimeZone_Saskatchewan:             "Saskatchewan",
	TimeZone_Alberta:                  "Alberta",
	TimeZone_BritishColumbia:          "BritishColumbia",
	TimeZone_Boise:                    "Boise",
	TimeZone_Boston:                   "Boston",
	TimeZone_Chicago:                  "Chicago",
	TimeZone_Dallas:                   "Dallas",
	TimeZone_Denver:                   "Denver",
	TimeZone_KansasCity:               "KansasCity",
	TimeZone_LasVegas:                 "LasVegas",
	TimeZone_LosAngeles:               "LosAngeles",
	TimeZone_Miami:                    "Miami",
	TimeZone_Minneapolis:              "Minneapolis",
	TimeZone_NewYork:                  "NewYork",
	TimeZone_NewOrleans:               "NewOrleans",
	TimeZone_Phoenix:                  "Phoenix",
	TimeZone_SantaFe:                  "SantaFe",
	TimeZone_Seattle:                  "Seattle",
	TimeZone_WashingtonDc:             "WashingtonDc",
	TimeZone_UsArizona:                "UsArizona",
	TimeZone_Chita:                    "Chita",
	TimeZone_Ekaterinburg:             "Ekaterinburg",
	TimeZone_Irkutsk:                  "Irkutsk",
	TimeZone_Kaliningrad:              "Kaliningrad",
	TimeZone_Krasnoyarsk:              "Krasnoyarsk",
	TimeZone_Novosibirsk:              "Novosibirsk",
	TimeZone_PetropavlovskKamchatskiy: "PetropavlovskKamchatskiy",
	TimeZone_Samara:                   "Samara",
	TimeZone_Vladivostok:              "Vladivostok",
	TimeZone_MexicoCentral:            "MexicoCentral",
	TimeZone_MexicoMountain:           "MexicoMountain",
	TimeZone_MexicoPacific:            "MexicoPacific",
	TimeZone_CapeTown:                 "CapeTown",
	TimeZone_Winkhoek:                 "Winkhoek",
	TimeZone_Lagos:                    "Lagos",
	TimeZone_Riyahd:                   "Riyahd",
	TimeZone_Venezuela:                "Venezuela",
	TimeZone_AustraliaLh:              "AustraliaLh",
	TimeZone_Santiago:                 "Santiago",
	TimeZone_Manual:                   "Manual",
	TimeZone_Automatic:                "Automatic",
}

func (k TimeZone) String() string {
	if uint(k) < uint(len(TimeZoneNames)) {
		if v, ok := TimeZoneNames[k]; ok {
			return v
		}
	}
	return "TimeZone(" + strconv.Itoa(int(k)) + ")"
}

var TimeZoneValues = map[string]TimeZone{
	"Almaty":                   TimeZone_Almaty,
	"Bangkok":                  TimeZone_Bangkok,
	"Bombay":                   TimeZone_Bombay,
	"Brasilia":                 TimeZone_Brasilia,
	"Cairo":                    TimeZone_Cairo,
	"CapeVerdeIs":              TimeZone_CapeVerdeIs,
	"Darwin":                   TimeZone_Darwin,
	"Eniwetok":                 TimeZone_Eniwetok,
	"Fiji":                     TimeZone_Fiji,
	"HongKong":                 TimeZone_HongKong,
	"Islamabad":                TimeZone_Islamabad,
	"Kabul":                    TimeZone_Kabul,
	"Magadan":                  TimeZone_Magadan,
	"MidAtlantic":              TimeZone_MidAtlantic,
	"Moscow":                   TimeZone_Moscow,
	"Muscat":                   TimeZone_Muscat,
	"Newfoundland":             TimeZone_Newfoundland,
	"Samoa":                    TimeZone_Samoa,
	"Sydney":                   TimeZone_Sydney,
	"Tehran":                   TimeZone_Tehran,
	"Tokyo":                    TimeZone_Tokyo,
	"UsAlaska":                 TimeZone_UsAlaska,
	"UsAtlantic":               TimeZone_UsAtlantic,
	"UsCentral":                TimeZone_UsCentral,
	"UsEastern":                TimeZone_UsEastern,
	"UsHawaii":                 TimeZone_UsHawaii,
	"UsMountain":               TimeZone_UsMountain,
	"UsPacific":                TimeZone_UsPacific,
	"Other":                    TimeZone_Other,
	"Auckland":                 TimeZone_Auckland,
	"Kathmandu":                TimeZone_Kathmandu,
	"EuropeWesternWet":         TimeZone_EuropeWesternWet,
	"EuropeCentralCet":         TimeZone_EuropeCentralCet,
	"EuropeEasternEet":         TimeZone_EuropeEasternEet,
	"Jakarta":                  TimeZone_Jakarta,
	"Perth":                    TimeZone_Perth,
	"Adelaide":                 TimeZone_Adelaide,
	"Brisbane":                 TimeZone_Brisbane,
	"Tasmania":                 TimeZone_Tasmania,
	"Iceland":                  TimeZone_Iceland,
	"Amsterdam":                TimeZone_Amsterdam,
	"Athens":                   TimeZone_Athens,
	"Barcelona":                TimeZone_Barcelona,
	"Berlin":                   TimeZone_Berlin,
	"Brussels":                 TimeZone_Brussels,
	"Budapest":                 TimeZone_Budapest,
	"Copenhagen":               TimeZone_Copenhagen,
	"Dublin":                   TimeZone_Dublin,
	"Helsinki":                 TimeZone_Helsinki,
	"Lisbon":                   TimeZone_Lisbon,
	"London":                   TimeZone_London,
	"Madrid":                   TimeZone_Madrid,
	"Munich":                   TimeZone_Munich,
	"Oslo":                     TimeZone_Oslo,
	"Paris":                    TimeZone_Paris,
	"Prague":                   TimeZone_Prague,
	"Reykjavik":                TimeZone_Reykjavik,
	"Rome":                     TimeZone_Rome,
	"Stockholm":                TimeZone_Stockholm,
	"Vienna":                   TimeZone_Vienna,
	"Warsaw":                   TimeZone_Warsaw,
	"Zurich":                   TimeZone_Zurich,
	"Quebec":                   TimeZone_Quebec,
	"Ontario":                  TimeZone_Ontario,
	"Manitoba":                 TimeZone_Manitoba,
	"Saskatchewan":             TimeZone_Saskatchewan,
	"Alberta":                  TimeZone_Alberta,
	"BritishColumbia":          TimeZone_BritishColumbia,
	"Boise":                    TimeZone_Boise,
	"Boston":                   TimeZone_Boston,
	"Chicago":                  TimeZone_Chicago,
	"Dallas":                   TimeZone_Dallas,
	"Denver":                   TimeZone_Denver,
	"KansasCity":               TimeZone_KansasCity,
	"LasVegas":                 TimeZone_LasVegas,
	"LosAngeles":               TimeZone_LosAngeles,
	"Miami":                    TimeZone_Miami,
	"Minneapolis":              TimeZone_Minneapolis,
	"NewYork":                  TimeZone_NewYork,
	"NewOrleans":               TimeZone_NewOrleans,
	"Phoenix":                  TimeZone_Phoenix,
	"SantaFe":                  TimeZone_SantaFe,
	"Seattle":                  TimeZone_Seattle,
	"WashingtonDc":             TimeZone_WashingtonDc,
	"UsArizona":                TimeZone_UsArizona,
	"Chita":                    TimeZone_Chita,
	"Ekaterinburg":             TimeZone_Ekaterinburg,
	"Irkutsk":                  TimeZone_Irkutsk,
	"Kaliningrad":              TimeZone_Kaliningrad,
	"Krasnoyarsk":              TimeZone_Krasnoyarsk,
	"Novosibirsk":              TimeZone_Novosibirsk,
	"PetropavlovskKamchatskiy": TimeZone_PetropavlovskKamchatskiy,
	"Samara":                   TimeZone_Samara,
	"Vladivostok":              TimeZone_Vladivostok,
	"MexicoCentral":            TimeZone_MexicoCentral,
	"MexicoMountain":           TimeZone_MexicoMountain,
	"MexicoPacific":            TimeZone_MexicoPacific,
	"CapeTown":                 TimeZone_CapeTown,
	"Winkhoek":                 TimeZone_Winkhoek,
	"Lagos":                    TimeZone_Lagos,
	"Riyahd":                   TimeZone_Riyahd,
	"Venezuela":                TimeZone_Venezuela,
	"AustraliaLh":              TimeZone_AustraliaLh,
	"Santiago":                 TimeZone_Santiago,
	"Manual":                   TimeZone_Manual,
	"Automatic":                TimeZone_Automatic,
}

func TimeZoneFromString(s string) TimeZone {
	if v, ok := TimeZoneValues[s]; ok {
		return v
	}
	return TimeZone(TimeZone_Invalid)
}
