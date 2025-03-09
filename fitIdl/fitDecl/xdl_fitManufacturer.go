package fitDecl

import strconv "strconv"

type Manufacturer uint16

const (
	Manufacturer_Garmin                 Manufacturer = 1
	Manufacturer_GarminFr405Antfs       Manufacturer = 2
	Manufacturer_Zephyr                 Manufacturer = 3
	Manufacturer_Dayton                 Manufacturer = 4
	Manufacturer_Idt                    Manufacturer = 5
	Manufacturer_Srm                    Manufacturer = 6
	Manufacturer_Quarq                  Manufacturer = 7
	Manufacturer_Ibike                  Manufacturer = 8
	Manufacturer_Saris                  Manufacturer = 9
	Manufacturer_SparkHk                Manufacturer = 10
	Manufacturer_Tanita                 Manufacturer = 11
	Manufacturer_Echowell               Manufacturer = 12
	Manufacturer_DynastreamOem          Manufacturer = 13
	Manufacturer_Nautilus               Manufacturer = 14
	Manufacturer_Dynastream             Manufacturer = 15
	Manufacturer_Timex                  Manufacturer = 16
	Manufacturer_Metrigear              Manufacturer = 17
	Manufacturer_Xelic                  Manufacturer = 18
	Manufacturer_Beurer                 Manufacturer = 19
	Manufacturer_Cardiosport            Manufacturer = 20
	Manufacturer_AAndD                  Manufacturer = 21
	Manufacturer_Hmm                    Manufacturer = 22
	Manufacturer_Suunto                 Manufacturer = 23
	Manufacturer_ThitaElektronik        Manufacturer = 24
	Manufacturer_Gpulse                 Manufacturer = 25
	Manufacturer_CleanMobile            Manufacturer = 26
	Manufacturer_PedalBrain             Manufacturer = 27
	Manufacturer_Peaksware              Manufacturer = 28
	Manufacturer_Saxonar                Manufacturer = 29
	Manufacturer_LemondFitness          Manufacturer = 30
	Manufacturer_Dexcom                 Manufacturer = 31
	Manufacturer_WahooFitness           Manufacturer = 32
	Manufacturer_OctaneFitness          Manufacturer = 33
	Manufacturer_Archinoetics           Manufacturer = 34
	Manufacturer_TheHurtBox             Manufacturer = 35
	Manufacturer_CitizenSystems         Manufacturer = 36
	Manufacturer_Magellan               Manufacturer = 37
	Manufacturer_Osynce                 Manufacturer = 38
	Manufacturer_Holux                  Manufacturer = 39
	Manufacturer_Concept2               Manufacturer = 40
	Manufacturer_Shimano                Manufacturer = 41
	Manufacturer_OneGiantLeap           Manufacturer = 42
	Manufacturer_AceSensor              Manufacturer = 43
	Manufacturer_BrimBrothers           Manufacturer = 44
	Manufacturer_Xplova                 Manufacturer = 45
	Manufacturer_PerceptionDigital      Manufacturer = 46
	Manufacturer_Bf1systems             Manufacturer = 47
	Manufacturer_Pioneer                Manufacturer = 48
	Manufacturer_Spantec                Manufacturer = 49
	Manufacturer_Metalogics             Manufacturer = 50
	Manufacturer__4iiiis                Manufacturer = 51
	Manufacturer_SeikoEpson             Manufacturer = 52
	Manufacturer_SeikoEpsonOem          Manufacturer = 53
	Manufacturer_IforPowell             Manufacturer = 54
	Manufacturer_MaxwellGuider          Manufacturer = 55
	Manufacturer_StarTrac               Manufacturer = 56
	Manufacturer_Breakaway              Manufacturer = 57
	Manufacturer_AlatechTechnologyLtd   Manufacturer = 58
	Manufacturer_MioTechnologyEurope    Manufacturer = 59
	Manufacturer_Rotor                  Manufacturer = 60
	Manufacturer_Geonaute               Manufacturer = 61
	Manufacturer_IdBike                 Manufacturer = 62
	Manufacturer_Specialized            Manufacturer = 63
	Manufacturer_Wtek                   Manufacturer = 64
	Manufacturer_PhysicalEnterprises    Manufacturer = 65
	Manufacturer_NorthPoleEngineering   Manufacturer = 66
	Manufacturer_Bkool                  Manufacturer = 67
	Manufacturer_Cateye                 Manufacturer = 68
	Manufacturer_StagesCycling          Manufacturer = 69
	Manufacturer_Sigmasport             Manufacturer = 70
	Manufacturer_Tomtom                 Manufacturer = 71
	Manufacturer_Peripedal              Manufacturer = 72
	Manufacturer_Wattbike               Manufacturer = 73
	Manufacturer_Moxy                   Manufacturer = 76
	Manufacturer_Ciclosport             Manufacturer = 77
	Manufacturer_Powerbahn              Manufacturer = 78
	Manufacturer_AcornProjectsAps       Manufacturer = 79
	Manufacturer_Lifebeam               Manufacturer = 80
	Manufacturer_Bontrager              Manufacturer = 81
	Manufacturer_Wellgo                 Manufacturer = 82
	Manufacturer_Scosche                Manufacturer = 83
	Manufacturer_Magura                 Manufacturer = 84
	Manufacturer_Woodway                Manufacturer = 85
	Manufacturer_Elite                  Manufacturer = 86
	Manufacturer_NielsenKellerman       Manufacturer = 87
	Manufacturer_DkCity                 Manufacturer = 88
	Manufacturer_Tacx                   Manufacturer = 89
	Manufacturer_DirectionTechnology    Manufacturer = 90
	Manufacturer_Magtonic               Manufacturer = 91
	Manufacturer__1partcarbon           Manufacturer = 92
	Manufacturer_InsideRideTechnologies Manufacturer = 93
	Manufacturer_SoundOfMotion          Manufacturer = 94
	Manufacturer_Stryd                  Manufacturer = 95
	Manufacturer_Icg                    Manufacturer = 96
	Manufacturer_MiPulse                Manufacturer = 97
	Manufacturer_BsxAthletics           Manufacturer = 98
	Manufacturer_Look                   Manufacturer = 99
	Manufacturer_CampagnoloSrl          Manufacturer = 100
	Manufacturer_BodyBikeSmart          Manufacturer = 101
	Manufacturer_Praxisworks            Manufacturer = 102
	Manufacturer_LimitsTechnology       Manufacturer = 103
	Manufacturer_TopactionTechnology    Manufacturer = 104
	Manufacturer_Cosinuss               Manufacturer = 105
	Manufacturer_Fitcare                Manufacturer = 106
	Manufacturer_Magene                 Manufacturer = 107
	Manufacturer_GiantManufacturingCo   Manufacturer = 108
	Manufacturer_Tigrasport             Manufacturer = 109
	Manufacturer_Salutron               Manufacturer = 110
	Manufacturer_Technogym              Manufacturer = 111
	Manufacturer_BrytonSensors          Manufacturer = 112
	Manufacturer_LatitudeLimited        Manufacturer = 113
	Manufacturer_SoaringTechnology      Manufacturer = 114
	Manufacturer_Igpsport               Manufacturer = 115
	Manufacturer_Thinkrider             Manufacturer = 116
	Manufacturer_GopherSport            Manufacturer = 117
	Manufacturer_Waterrower             Manufacturer = 118
	Manufacturer_Orangetheory           Manufacturer = 119
	Manufacturer_Inpeak                 Manufacturer = 120
	Manufacturer_Kinetic                Manufacturer = 121
	Manufacturer_JohnsonHealthTech      Manufacturer = 122
	Manufacturer_PolarElectro           Manufacturer = 123
	Manufacturer_Seesense               Manufacturer = 124
	Manufacturer_NciTechnology          Manufacturer = 125
	Manufacturer_Iqsquare               Manufacturer = 126
	Manufacturer_Leomo                  Manufacturer = 127
	Manufacturer_IfitCom                Manufacturer = 128
	Manufacturer_CorosByte              Manufacturer = 129
	Manufacturer_VersaDesign            Manufacturer = 130
	Manufacturer_Chileaf                Manufacturer = 131
	Manufacturer_Cycplus                Manufacturer = 132
	Manufacturer_GravaaByte             Manufacturer = 133
	Manufacturer_Sigeyi                 Manufacturer = 134
	Manufacturer_Coospo                 Manufacturer = 135
	Manufacturer_Geoid                  Manufacturer = 136
	Manufacturer_Bosch                  Manufacturer = 137
	Manufacturer_Kyto                   Manufacturer = 138
	Manufacturer_KineticSports          Manufacturer = 139
	Manufacturer_DecathlonByte          Manufacturer = 140
	Manufacturer_TqSystems              Manufacturer = 141
	Manufacturer_TagHeuer               Manufacturer = 142
	Manufacturer_KeiserFitness          Manufacturer = 143
	Manufacturer_ZwiftByte              Manufacturer = 144
	Manufacturer_PorscheEp              Manufacturer = 145
	Manufacturer_Blackbird              Manufacturer = 146
	Manufacturer_MeilanByte             Manufacturer = 147
	Manufacturer_Ezon                   Manufacturer = 148
	Manufacturer_Laisi                  Manufacturer = 149
	Manufacturer_Myzone                 Manufacturer = 150
	Manufacturer_Abawo                  Manufacturer = 151
	Manufacturer_Bafang                 Manufacturer = 152
	Manufacturer_Development            Manufacturer = 255
	Manufacturer_Healthandlife          Manufacturer = 257
	Manufacturer_Lezyne                 Manufacturer = 258
	Manufacturer_ScribeLabs             Manufacturer = 259
	Manufacturer_Zwift                  Manufacturer = 260
	Manufacturer_Watteam                Manufacturer = 261
	Manufacturer_Recon                  Manufacturer = 262
	Manufacturer_FaveroElectronics      Manufacturer = 263
	Manufacturer_Dynovelo               Manufacturer = 264
	Manufacturer_Strava                 Manufacturer = 265
	Manufacturer_Precor                 Manufacturer = 266
	Manufacturer_Bryton                 Manufacturer = 267
	Manufacturer_Sram                   Manufacturer = 268
	Manufacturer_Navman                 Manufacturer = 269
	Manufacturer_Cobi                   Manufacturer = 270
	Manufacturer_Spivi                  Manufacturer = 271
	Manufacturer_MioMagellan            Manufacturer = 272
	Manufacturer_Evesports              Manufacturer = 273
	Manufacturer_SensitivusGauge        Manufacturer = 274
	Manufacturer_Podoon                 Manufacturer = 275
	Manufacturer_LifeTimeFitness        Manufacturer = 276
	Manufacturer_FalcoEMotors           Manufacturer = 277
	Manufacturer_Minoura                Manufacturer = 278
	Manufacturer_Cycliq                 Manufacturer = 279
	Manufacturer_Luxottica              Manufacturer = 280
	Manufacturer_TrainerRoad            Manufacturer = 281
	Manufacturer_TheSufferfest          Manufacturer = 282
	Manufacturer_Fullspeedahead         Manufacturer = 283
	Manufacturer_Virtualtraining        Manufacturer = 284
	Manufacturer_Feedbacksports         Manufacturer = 285
	Manufacturer_Omata                  Manufacturer = 286
	Manufacturer_Vdo                    Manufacturer = 287
	Manufacturer_Magneticdays           Manufacturer = 288
	Manufacturer_Hammerhead             Manufacturer = 289
	Manufacturer_KineticByKurt          Manufacturer = 290
	Manufacturer_Shapelog               Manufacturer = 291
	Manufacturer_Dabuziduo              Manufacturer = 292
	Manufacturer_Jetblack               Manufacturer = 293
	Manufacturer_Coros                  Manufacturer = 294
	Manufacturer_Virtugo                Manufacturer = 295
	Manufacturer_Velosense              Manufacturer = 296
	Manufacturer_Cycligentinc           Manufacturer = 297
	Manufacturer_Trailforks             Manufacturer = 298
	Manufacturer_MahleEbikemotion       Manufacturer = 299
	Manufacturer_Nurvv                  Manufacturer = 300
	Manufacturer_Microprogram           Manufacturer = 301
	Manufacturer_Zone5cloud             Manufacturer = 302
	Manufacturer_Greenteg               Manufacturer = 303
	Manufacturer_YamahaMotors           Manufacturer = 304
	Manufacturer_Whoop                  Manufacturer = 305
	Manufacturer_Gravaa                 Manufacturer = 306
	Manufacturer_Onelap                 Manufacturer = 307
	Manufacturer_MonarkExercise         Manufacturer = 308
	Manufacturer_Form                   Manufacturer = 309
	Manufacturer_Decathlon              Manufacturer = 310
	Manufacturer_Syncros                Manufacturer = 311
	Manufacturer_Heatup                 Manufacturer = 312
	Manufacturer_Cannondale             Manufacturer = 313
	Manufacturer_TrueFitness            Manufacturer = 314
	Manufacturer_RGTCycling             Manufacturer = 315
	Manufacturer_Vasa                   Manufacturer = 316
	Manufacturer_RaceRepublic           Manufacturer = 317
	Manufacturer_Fazua                  Manufacturer = 318
	Manufacturer_OrekaTraining          Manufacturer = 319
	Manufacturer_Lsec                   Manufacturer = 320
	Manufacturer_LululemonStudio        Manufacturer = 321
	Manufacturer_Shanyue                Manufacturer = 322
	Manufacturer_SpinningMda            Manufacturer = 323
	Manufacturer_Hilldating             Manufacturer = 324
	Manufacturer_AeroSensor             Manufacturer = 325
	Manufacturer_Nike                   Manufacturer = 326
	Manufacturer_Magicshine             Manufacturer = 327
	Manufacturer_Ictrainer              Manufacturer = 328
	Manufacturer_AbsoluteCycling        Manufacturer = 329
	Manufacturer_EoSwimbetter           Manufacturer = 330
	Manufacturer_Mywhoosh               Manufacturer = 331
	Manufacturer_Ravemen                Manufacturer = 332
	Manufacturer_Actigraphcorp          Manufacturer = 5759
	Manufacturer_Invalid                Manufacturer = 65535
)

var ManufacturerNames = map[Manufacturer]string{
	Manufacturer_Garmin:                 "Garmin",
	Manufacturer_GarminFr405Antfs:       "GarminFr405Antfs",
	Manufacturer_Zephyr:                 "Zephyr",
	Manufacturer_Dayton:                 "Dayton",
	Manufacturer_Idt:                    "Idt",
	Manufacturer_Srm:                    "Srm",
	Manufacturer_Quarq:                  "Quarq",
	Manufacturer_Ibike:                  "Ibike",
	Manufacturer_Saris:                  "Saris",
	Manufacturer_SparkHk:                "SparkHk",
	Manufacturer_Tanita:                 "Tanita",
	Manufacturer_Echowell:               "Echowell",
	Manufacturer_DynastreamOem:          "DynastreamOem",
	Manufacturer_Nautilus:               "Nautilus",
	Manufacturer_Dynastream:             "Dynastream",
	Manufacturer_Timex:                  "Timex",
	Manufacturer_Metrigear:              "Metrigear",
	Manufacturer_Xelic:                  "Xelic",
	Manufacturer_Beurer:                 "Beurer",
	Manufacturer_Cardiosport:            "Cardiosport",
	Manufacturer_AAndD:                  "AAndD",
	Manufacturer_Hmm:                    "Hmm",
	Manufacturer_Suunto:                 "Suunto",
	Manufacturer_ThitaElektronik:        "ThitaElektronik",
	Manufacturer_Gpulse:                 "Gpulse",
	Manufacturer_CleanMobile:            "CleanMobile",
	Manufacturer_PedalBrain:             "PedalBrain",
	Manufacturer_Peaksware:              "Peaksware",
	Manufacturer_Saxonar:                "Saxonar",
	Manufacturer_LemondFitness:          "LemondFitness",
	Manufacturer_Dexcom:                 "Dexcom",
	Manufacturer_WahooFitness:           "WahooFitness",
	Manufacturer_OctaneFitness:          "OctaneFitness",
	Manufacturer_Archinoetics:           "Archinoetics",
	Manufacturer_TheHurtBox:             "TheHurtBox",
	Manufacturer_CitizenSystems:         "CitizenSystems",
	Manufacturer_Magellan:               "Magellan",
	Manufacturer_Osynce:                 "Osynce",
	Manufacturer_Holux:                  "Holux",
	Manufacturer_Concept2:               "Concept2",
	Manufacturer_Shimano:                "Shimano",
	Manufacturer_OneGiantLeap:           "OneGiantLeap",
	Manufacturer_AceSensor:              "AceSensor",
	Manufacturer_BrimBrothers:           "BrimBrothers",
	Manufacturer_Xplova:                 "Xplova",
	Manufacturer_PerceptionDigital:      "PerceptionDigital",
	Manufacturer_Bf1systems:             "Bf1systems",
	Manufacturer_Pioneer:                "Pioneer",
	Manufacturer_Spantec:                "Spantec",
	Manufacturer_Metalogics:             "Metalogics",
	Manufacturer__4iiiis:                "_4iiiis",
	Manufacturer_SeikoEpson:             "SeikoEpson",
	Manufacturer_SeikoEpsonOem:          "SeikoEpsonOem",
	Manufacturer_IforPowell:             "IforPowell",
	Manufacturer_MaxwellGuider:          "MaxwellGuider",
	Manufacturer_StarTrac:               "StarTrac",
	Manufacturer_Breakaway:              "Breakaway",
	Manufacturer_AlatechTechnologyLtd:   "AlatechTechnologyLtd",
	Manufacturer_MioTechnologyEurope:    "MioTechnologyEurope",
	Manufacturer_Rotor:                  "Rotor",
	Manufacturer_Geonaute:               "Geonaute",
	Manufacturer_IdBike:                 "IdBike",
	Manufacturer_Specialized:            "Specialized",
	Manufacturer_Wtek:                   "Wtek",
	Manufacturer_PhysicalEnterprises:    "PhysicalEnterprises",
	Manufacturer_NorthPoleEngineering:   "NorthPoleEngineering",
	Manufacturer_Bkool:                  "Bkool",
	Manufacturer_Cateye:                 "Cateye",
	Manufacturer_StagesCycling:          "StagesCycling",
	Manufacturer_Sigmasport:             "Sigmasport",
	Manufacturer_Tomtom:                 "Tomtom",
	Manufacturer_Peripedal:              "Peripedal",
	Manufacturer_Wattbike:               "Wattbike",
	Manufacturer_Moxy:                   "Moxy",
	Manufacturer_Ciclosport:             "Ciclosport",
	Manufacturer_Powerbahn:              "Powerbahn",
	Manufacturer_AcornProjectsAps:       "AcornProjectsAps",
	Manufacturer_Lifebeam:               "Lifebeam",
	Manufacturer_Bontrager:              "Bontrager",
	Manufacturer_Wellgo:                 "Wellgo",
	Manufacturer_Scosche:                "Scosche",
	Manufacturer_Magura:                 "Magura",
	Manufacturer_Woodway:                "Woodway",
	Manufacturer_Elite:                  "Elite",
	Manufacturer_NielsenKellerman:       "NielsenKellerman",
	Manufacturer_DkCity:                 "DkCity",
	Manufacturer_Tacx:                   "Tacx",
	Manufacturer_DirectionTechnology:    "DirectionTechnology",
	Manufacturer_Magtonic:               "Magtonic",
	Manufacturer__1partcarbon:           "_1partcarbon",
	Manufacturer_InsideRideTechnologies: "InsideRideTechnologies",
	Manufacturer_SoundOfMotion:          "SoundOfMotion",
	Manufacturer_Stryd:                  "Stryd",
	Manufacturer_Icg:                    "Icg",
	Manufacturer_MiPulse:                "MiPulse",
	Manufacturer_BsxAthletics:           "BsxAthletics",
	Manufacturer_Look:                   "Look",
	Manufacturer_CampagnoloSrl:          "CampagnoloSrl",
	Manufacturer_BodyBikeSmart:          "BodyBikeSmart",
	Manufacturer_Praxisworks:            "Praxisworks",
	Manufacturer_LimitsTechnology:       "LimitsTechnology",
	Manufacturer_TopactionTechnology:    "TopactionTechnology",
	Manufacturer_Cosinuss:               "Cosinuss",
	Manufacturer_Fitcare:                "Fitcare",
	Manufacturer_Magene:                 "Magene",
	Manufacturer_GiantManufacturingCo:   "GiantManufacturingCo",
	Manufacturer_Tigrasport:             "Tigrasport",
	Manufacturer_Salutron:               "Salutron",
	Manufacturer_Technogym:              "Technogym",
	Manufacturer_BrytonSensors:          "BrytonSensors",
	Manufacturer_LatitudeLimited:        "LatitudeLimited",
	Manufacturer_SoaringTechnology:      "SoaringTechnology",
	Manufacturer_Igpsport:               "Igpsport",
	Manufacturer_Thinkrider:             "Thinkrider",
	Manufacturer_GopherSport:            "GopherSport",
	Manufacturer_Waterrower:             "Waterrower",
	Manufacturer_Orangetheory:           "Orangetheory",
	Manufacturer_Inpeak:                 "Inpeak",
	Manufacturer_Kinetic:                "Kinetic",
	Manufacturer_JohnsonHealthTech:      "JohnsonHealthTech",
	Manufacturer_PolarElectro:           "PolarElectro",
	Manufacturer_Seesense:               "Seesense",
	Manufacturer_NciTechnology:          "NciTechnology",
	Manufacturer_Iqsquare:               "Iqsquare",
	Manufacturer_Leomo:                  "Leomo",
	Manufacturer_IfitCom:                "IfitCom",
	Manufacturer_CorosByte:              "CorosByte",
	Manufacturer_VersaDesign:            "VersaDesign",
	Manufacturer_Chileaf:                "Chileaf",
	Manufacturer_Cycplus:                "Cycplus",
	Manufacturer_GravaaByte:             "GravaaByte",
	Manufacturer_Sigeyi:                 "Sigeyi",
	Manufacturer_Coospo:                 "Coospo",
	Manufacturer_Geoid:                  "Geoid",
	Manufacturer_Bosch:                  "Bosch",
	Manufacturer_Kyto:                   "Kyto",
	Manufacturer_KineticSports:          "KineticSports",
	Manufacturer_DecathlonByte:          "DecathlonByte",
	Manufacturer_TqSystems:              "TqSystems",
	Manufacturer_TagHeuer:               "TagHeuer",
	Manufacturer_KeiserFitness:          "KeiserFitness",
	Manufacturer_ZwiftByte:              "ZwiftByte",
	Manufacturer_PorscheEp:              "PorscheEp",
	Manufacturer_Blackbird:              "Blackbird",
	Manufacturer_MeilanByte:             "MeilanByte",
	Manufacturer_Ezon:                   "Ezon",
	Manufacturer_Laisi:                  "Laisi",
	Manufacturer_Myzone:                 "Myzone",
	Manufacturer_Abawo:                  "Abawo",
	Manufacturer_Bafang:                 "Bafang",
	Manufacturer_Development:            "Development",
	Manufacturer_Healthandlife:          "Healthandlife",
	Manufacturer_Lezyne:                 "Lezyne",
	Manufacturer_ScribeLabs:             "ScribeLabs",
	Manufacturer_Zwift:                  "Zwift",
	Manufacturer_Watteam:                "Watteam",
	Manufacturer_Recon:                  "Recon",
	Manufacturer_FaveroElectronics:      "FaveroElectronics",
	Manufacturer_Dynovelo:               "Dynovelo",
	Manufacturer_Strava:                 "Strava",
	Manufacturer_Precor:                 "Precor",
	Manufacturer_Bryton:                 "Bryton",
	Manufacturer_Sram:                   "Sram",
	Manufacturer_Navman:                 "Navman",
	Manufacturer_Cobi:                   "Cobi",
	Manufacturer_Spivi:                  "Spivi",
	Manufacturer_MioMagellan:            "MioMagellan",
	Manufacturer_Evesports:              "Evesports",
	Manufacturer_SensitivusGauge:        "SensitivusGauge",
	Manufacturer_Podoon:                 "Podoon",
	Manufacturer_LifeTimeFitness:        "LifeTimeFitness",
	Manufacturer_FalcoEMotors:           "FalcoEMotors",
	Manufacturer_Minoura:                "Minoura",
	Manufacturer_Cycliq:                 "Cycliq",
	Manufacturer_Luxottica:              "Luxottica",
	Manufacturer_TrainerRoad:            "TrainerRoad",
	Manufacturer_TheSufferfest:          "TheSufferfest",
	Manufacturer_Fullspeedahead:         "Fullspeedahead",
	Manufacturer_Virtualtraining:        "Virtualtraining",
	Manufacturer_Feedbacksports:         "Feedbacksports",
	Manufacturer_Omata:                  "Omata",
	Manufacturer_Vdo:                    "Vdo",
	Manufacturer_Magneticdays:           "Magneticdays",
	Manufacturer_Hammerhead:             "Hammerhead",
	Manufacturer_KineticByKurt:          "KineticByKurt",
	Manufacturer_Shapelog:               "Shapelog",
	Manufacturer_Dabuziduo:              "Dabuziduo",
	Manufacturer_Jetblack:               "Jetblack",
	Manufacturer_Coros:                  "Coros",
	Manufacturer_Virtugo:                "Virtugo",
	Manufacturer_Velosense:              "Velosense",
	Manufacturer_Cycligentinc:           "Cycligentinc",
	Manufacturer_Trailforks:             "Trailforks",
	Manufacturer_MahleEbikemotion:       "MahleEbikemotion",
	Manufacturer_Nurvv:                  "Nurvv",
	Manufacturer_Microprogram:           "Microprogram",
	Manufacturer_Zone5cloud:             "Zone5cloud",
	Manufacturer_Greenteg:               "Greenteg",
	Manufacturer_YamahaMotors:           "YamahaMotors",
	Manufacturer_Whoop:                  "Whoop",
	Manufacturer_Gravaa:                 "Gravaa",
	Manufacturer_Onelap:                 "Onelap",
	Manufacturer_MonarkExercise:         "MonarkExercise",
	Manufacturer_Form:                   "Form",
	Manufacturer_Decathlon:              "Decathlon",
	Manufacturer_Syncros:                "Syncros",
	Manufacturer_Heatup:                 "Heatup",
	Manufacturer_Cannondale:             "Cannondale",
	Manufacturer_TrueFitness:            "TrueFitness",
	Manufacturer_RGTCycling:             "RGTCycling",
	Manufacturer_Vasa:                   "Vasa",
	Manufacturer_RaceRepublic:           "RaceRepublic",
	Manufacturer_Fazua:                  "Fazua",
	Manufacturer_OrekaTraining:          "OrekaTraining",
	Manufacturer_Lsec:                   "Lsec",
	Manufacturer_LululemonStudio:        "LululemonStudio",
	Manufacturer_Shanyue:                "Shanyue",
	Manufacturer_SpinningMda:            "SpinningMda",
	Manufacturer_Hilldating:             "Hilldating",
	Manufacturer_AeroSensor:             "AeroSensor",
	Manufacturer_Nike:                   "Nike",
	Manufacturer_Magicshine:             "Magicshine",
	Manufacturer_Ictrainer:              "Ictrainer",
	Manufacturer_AbsoluteCycling:        "AbsoluteCycling",
	Manufacturer_EoSwimbetter:           "EoSwimbetter",
	Manufacturer_Mywhoosh:               "Mywhoosh",
	Manufacturer_Ravemen:                "Ravemen",
	Manufacturer_Actigraphcorp:          "Actigraphcorp",
}

func (k Manufacturer) String() string {
	if uint(k) < uint(len(ManufacturerNames)) {
		if v, ok := ManufacturerNames[k]; ok {
			return v
		}
	}
	return "Manufacturer(" + strconv.Itoa(int(k)) + ")"
}

var ManufacturerValues = map[string]Manufacturer{
	"Garmin":                 Manufacturer_Garmin,
	"GarminFr405Antfs":       Manufacturer_GarminFr405Antfs,
	"Zephyr":                 Manufacturer_Zephyr,
	"Dayton":                 Manufacturer_Dayton,
	"Idt":                    Manufacturer_Idt,
	"Srm":                    Manufacturer_Srm,
	"Quarq":                  Manufacturer_Quarq,
	"Ibike":                  Manufacturer_Ibike,
	"Saris":                  Manufacturer_Saris,
	"SparkHk":                Manufacturer_SparkHk,
	"Tanita":                 Manufacturer_Tanita,
	"Echowell":               Manufacturer_Echowell,
	"DynastreamOem":          Manufacturer_DynastreamOem,
	"Nautilus":               Manufacturer_Nautilus,
	"Dynastream":             Manufacturer_Dynastream,
	"Timex":                  Manufacturer_Timex,
	"Metrigear":              Manufacturer_Metrigear,
	"Xelic":                  Manufacturer_Xelic,
	"Beurer":                 Manufacturer_Beurer,
	"Cardiosport":            Manufacturer_Cardiosport,
	"AAndD":                  Manufacturer_AAndD,
	"Hmm":                    Manufacturer_Hmm,
	"Suunto":                 Manufacturer_Suunto,
	"ThitaElektronik":        Manufacturer_ThitaElektronik,
	"Gpulse":                 Manufacturer_Gpulse,
	"CleanMobile":            Manufacturer_CleanMobile,
	"PedalBrain":             Manufacturer_PedalBrain,
	"Peaksware":              Manufacturer_Peaksware,
	"Saxonar":                Manufacturer_Saxonar,
	"LemondFitness":          Manufacturer_LemondFitness,
	"Dexcom":                 Manufacturer_Dexcom,
	"WahooFitness":           Manufacturer_WahooFitness,
	"OctaneFitness":          Manufacturer_OctaneFitness,
	"Archinoetics":           Manufacturer_Archinoetics,
	"TheHurtBox":             Manufacturer_TheHurtBox,
	"CitizenSystems":         Manufacturer_CitizenSystems,
	"Magellan":               Manufacturer_Magellan,
	"Osynce":                 Manufacturer_Osynce,
	"Holux":                  Manufacturer_Holux,
	"Concept2":               Manufacturer_Concept2,
	"Shimano":                Manufacturer_Shimano,
	"OneGiantLeap":           Manufacturer_OneGiantLeap,
	"AceSensor":              Manufacturer_AceSensor,
	"BrimBrothers":           Manufacturer_BrimBrothers,
	"Xplova":                 Manufacturer_Xplova,
	"PerceptionDigital":      Manufacturer_PerceptionDigital,
	"Bf1systems":             Manufacturer_Bf1systems,
	"Pioneer":                Manufacturer_Pioneer,
	"Spantec":                Manufacturer_Spantec,
	"Metalogics":             Manufacturer_Metalogics,
	"_4iiiis":                Manufacturer__4iiiis,
	"SeikoEpson":             Manufacturer_SeikoEpson,
	"SeikoEpsonOem":          Manufacturer_SeikoEpsonOem,
	"IforPowell":             Manufacturer_IforPowell,
	"MaxwellGuider":          Manufacturer_MaxwellGuider,
	"StarTrac":               Manufacturer_StarTrac,
	"Breakaway":              Manufacturer_Breakaway,
	"AlatechTechnologyLtd":   Manufacturer_AlatechTechnologyLtd,
	"MioTechnologyEurope":    Manufacturer_MioTechnologyEurope,
	"Rotor":                  Manufacturer_Rotor,
	"Geonaute":               Manufacturer_Geonaute,
	"IdBike":                 Manufacturer_IdBike,
	"Specialized":            Manufacturer_Specialized,
	"Wtek":                   Manufacturer_Wtek,
	"PhysicalEnterprises":    Manufacturer_PhysicalEnterprises,
	"NorthPoleEngineering":   Manufacturer_NorthPoleEngineering,
	"Bkool":                  Manufacturer_Bkool,
	"Cateye":                 Manufacturer_Cateye,
	"StagesCycling":          Manufacturer_StagesCycling,
	"Sigmasport":             Manufacturer_Sigmasport,
	"Tomtom":                 Manufacturer_Tomtom,
	"Peripedal":              Manufacturer_Peripedal,
	"Wattbike":               Manufacturer_Wattbike,
	"Moxy":                   Manufacturer_Moxy,
	"Ciclosport":             Manufacturer_Ciclosport,
	"Powerbahn":              Manufacturer_Powerbahn,
	"AcornProjectsAps":       Manufacturer_AcornProjectsAps,
	"Lifebeam":               Manufacturer_Lifebeam,
	"Bontrager":              Manufacturer_Bontrager,
	"Wellgo":                 Manufacturer_Wellgo,
	"Scosche":                Manufacturer_Scosche,
	"Magura":                 Manufacturer_Magura,
	"Woodway":                Manufacturer_Woodway,
	"Elite":                  Manufacturer_Elite,
	"NielsenKellerman":       Manufacturer_NielsenKellerman,
	"DkCity":                 Manufacturer_DkCity,
	"Tacx":                   Manufacturer_Tacx,
	"DirectionTechnology":    Manufacturer_DirectionTechnology,
	"Magtonic":               Manufacturer_Magtonic,
	"_1partcarbon":           Manufacturer__1partcarbon,
	"InsideRideTechnologies": Manufacturer_InsideRideTechnologies,
	"SoundOfMotion":          Manufacturer_SoundOfMotion,
	"Stryd":                  Manufacturer_Stryd,
	"Icg":                    Manufacturer_Icg,
	"MiPulse":                Manufacturer_MiPulse,
	"BsxAthletics":           Manufacturer_BsxAthletics,
	"Look":                   Manufacturer_Look,
	"CampagnoloSrl":          Manufacturer_CampagnoloSrl,
	"BodyBikeSmart":          Manufacturer_BodyBikeSmart,
	"Praxisworks":            Manufacturer_Praxisworks,
	"LimitsTechnology":       Manufacturer_LimitsTechnology,
	"TopactionTechnology":    Manufacturer_TopactionTechnology,
	"Cosinuss":               Manufacturer_Cosinuss,
	"Fitcare":                Manufacturer_Fitcare,
	"Magene":                 Manufacturer_Magene,
	"GiantManufacturingCo":   Manufacturer_GiantManufacturingCo,
	"Tigrasport":             Manufacturer_Tigrasport,
	"Salutron":               Manufacturer_Salutron,
	"Technogym":              Manufacturer_Technogym,
	"BrytonSensors":          Manufacturer_BrytonSensors,
	"LatitudeLimited":        Manufacturer_LatitudeLimited,
	"SoaringTechnology":      Manufacturer_SoaringTechnology,
	"Igpsport":               Manufacturer_Igpsport,
	"Thinkrider":             Manufacturer_Thinkrider,
	"GopherSport":            Manufacturer_GopherSport,
	"Waterrower":             Manufacturer_Waterrower,
	"Orangetheory":           Manufacturer_Orangetheory,
	"Inpeak":                 Manufacturer_Inpeak,
	"Kinetic":                Manufacturer_Kinetic,
	"JohnsonHealthTech":      Manufacturer_JohnsonHealthTech,
	"PolarElectro":           Manufacturer_PolarElectro,
	"Seesense":               Manufacturer_Seesense,
	"NciTechnology":          Manufacturer_NciTechnology,
	"Iqsquare":               Manufacturer_Iqsquare,
	"Leomo":                  Manufacturer_Leomo,
	"IfitCom":                Manufacturer_IfitCom,
	"CorosByte":              Manufacturer_CorosByte,
	"VersaDesign":            Manufacturer_VersaDesign,
	"Chileaf":                Manufacturer_Chileaf,
	"Cycplus":                Manufacturer_Cycplus,
	"GravaaByte":             Manufacturer_GravaaByte,
	"Sigeyi":                 Manufacturer_Sigeyi,
	"Coospo":                 Manufacturer_Coospo,
	"Geoid":                  Manufacturer_Geoid,
	"Bosch":                  Manufacturer_Bosch,
	"Kyto":                   Manufacturer_Kyto,
	"KineticSports":          Manufacturer_KineticSports,
	"DecathlonByte":          Manufacturer_DecathlonByte,
	"TqSystems":              Manufacturer_TqSystems,
	"TagHeuer":               Manufacturer_TagHeuer,
	"KeiserFitness":          Manufacturer_KeiserFitness,
	"ZwiftByte":              Manufacturer_ZwiftByte,
	"PorscheEp":              Manufacturer_PorscheEp,
	"Blackbird":              Manufacturer_Blackbird,
	"MeilanByte":             Manufacturer_MeilanByte,
	"Ezon":                   Manufacturer_Ezon,
	"Laisi":                  Manufacturer_Laisi,
	"Myzone":                 Manufacturer_Myzone,
	"Abawo":                  Manufacturer_Abawo,
	"Bafang":                 Manufacturer_Bafang,
	"Development":            Manufacturer_Development,
	"Healthandlife":          Manufacturer_Healthandlife,
	"Lezyne":                 Manufacturer_Lezyne,
	"ScribeLabs":             Manufacturer_ScribeLabs,
	"Zwift":                  Manufacturer_Zwift,
	"Watteam":                Manufacturer_Watteam,
	"Recon":                  Manufacturer_Recon,
	"FaveroElectronics":      Manufacturer_FaveroElectronics,
	"Dynovelo":               Manufacturer_Dynovelo,
	"Strava":                 Manufacturer_Strava,
	"Precor":                 Manufacturer_Precor,
	"Bryton":                 Manufacturer_Bryton,
	"Sram":                   Manufacturer_Sram,
	"Navman":                 Manufacturer_Navman,
	"Cobi":                   Manufacturer_Cobi,
	"Spivi":                  Manufacturer_Spivi,
	"MioMagellan":            Manufacturer_MioMagellan,
	"Evesports":              Manufacturer_Evesports,
	"SensitivusGauge":        Manufacturer_SensitivusGauge,
	"Podoon":                 Manufacturer_Podoon,
	"LifeTimeFitness":        Manufacturer_LifeTimeFitness,
	"FalcoEMotors":           Manufacturer_FalcoEMotors,
	"Minoura":                Manufacturer_Minoura,
	"Cycliq":                 Manufacturer_Cycliq,
	"Luxottica":              Manufacturer_Luxottica,
	"TrainerRoad":            Manufacturer_TrainerRoad,
	"TheSufferfest":          Manufacturer_TheSufferfest,
	"Fullspeedahead":         Manufacturer_Fullspeedahead,
	"Virtualtraining":        Manufacturer_Virtualtraining,
	"Feedbacksports":         Manufacturer_Feedbacksports,
	"Omata":                  Manufacturer_Omata,
	"Vdo":                    Manufacturer_Vdo,
	"Magneticdays":           Manufacturer_Magneticdays,
	"Hammerhead":             Manufacturer_Hammerhead,
	"KineticByKurt":          Manufacturer_KineticByKurt,
	"Shapelog":               Manufacturer_Shapelog,
	"Dabuziduo":              Manufacturer_Dabuziduo,
	"Jetblack":               Manufacturer_Jetblack,
	"Coros":                  Manufacturer_Coros,
	"Virtugo":                Manufacturer_Virtugo,
	"Velosense":              Manufacturer_Velosense,
	"Cycligentinc":           Manufacturer_Cycligentinc,
	"Trailforks":             Manufacturer_Trailforks,
	"MahleEbikemotion":       Manufacturer_MahleEbikemotion,
	"Nurvv":                  Manufacturer_Nurvv,
	"Microprogram":           Manufacturer_Microprogram,
	"Zone5cloud":             Manufacturer_Zone5cloud,
	"Greenteg":               Manufacturer_Greenteg,
	"YamahaMotors":           Manufacturer_YamahaMotors,
	"Whoop":                  Manufacturer_Whoop,
	"Gravaa":                 Manufacturer_Gravaa,
	"Onelap":                 Manufacturer_Onelap,
	"MonarkExercise":         Manufacturer_MonarkExercise,
	"Form":                   Manufacturer_Form,
	"Decathlon":              Manufacturer_Decathlon,
	"Syncros":                Manufacturer_Syncros,
	"Heatup":                 Manufacturer_Heatup,
	"Cannondale":             Manufacturer_Cannondale,
	"TrueFitness":            Manufacturer_TrueFitness,
	"RGTCycling":             Manufacturer_RGTCycling,
	"Vasa":                   Manufacturer_Vasa,
	"RaceRepublic":           Manufacturer_RaceRepublic,
	"Fazua":                  Manufacturer_Fazua,
	"OrekaTraining":          Manufacturer_OrekaTraining,
	"Lsec":                   Manufacturer_Lsec,
	"LululemonStudio":        Manufacturer_LululemonStudio,
	"Shanyue":                Manufacturer_Shanyue,
	"SpinningMda":            Manufacturer_SpinningMda,
	"Hilldating":             Manufacturer_Hilldating,
	"AeroSensor":             Manufacturer_AeroSensor,
	"Nike":                   Manufacturer_Nike,
	"Magicshine":             Manufacturer_Magicshine,
	"Ictrainer":              Manufacturer_Ictrainer,
	"AbsoluteCycling":        Manufacturer_AbsoluteCycling,
	"EoSwimbetter":           Manufacturer_EoSwimbetter,
	"Mywhoosh":               Manufacturer_Mywhoosh,
	"Ravemen":                Manufacturer_Ravemen,
	"Actigraphcorp":          Manufacturer_Actigraphcorp,
}

func ManufacturerFromString(s string) Manufacturer {
	if v, ok := ManufacturerValues[s]; ok {
		return v
	}
	return Manufacturer(Manufacturer_Invalid)
}
