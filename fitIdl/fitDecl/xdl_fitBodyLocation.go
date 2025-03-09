package fitDecl

import strconv "strconv"

type BodyLocation byte

const (
	BodyLocation_LeftLeg               BodyLocation = 0
	BodyLocation_LeftCalf              BodyLocation = 1
	BodyLocation_LeftShin              BodyLocation = 2
	BodyLocation_LeftHamstring         BodyLocation = 3
	BodyLocation_LeftQuad              BodyLocation = 4
	BodyLocation_LeftGlute             BodyLocation = 5
	BodyLocation_RightLeg              BodyLocation = 6
	BodyLocation_RightCalf             BodyLocation = 7
	BodyLocation_RightShin             BodyLocation = 8
	BodyLocation_RightHamstring        BodyLocation = 9
	BodyLocation_RightQuad             BodyLocation = 10
	BodyLocation_RightGlute            BodyLocation = 11
	BodyLocation_TorsoBack             BodyLocation = 12
	BodyLocation_LeftLowerBack         BodyLocation = 13
	BodyLocation_LeftUpperBack         BodyLocation = 14
	BodyLocation_RightLowerBack        BodyLocation = 15
	BodyLocation_RightUpperBack        BodyLocation = 16
	BodyLocation_TorsoFront            BodyLocation = 17
	BodyLocation_LeftAbdomen           BodyLocation = 18
	BodyLocation_LeftChest             BodyLocation = 19
	BodyLocation_RightAbdomen          BodyLocation = 20
	BodyLocation_RightChest            BodyLocation = 21
	BodyLocation_LeftArm               BodyLocation = 22
	BodyLocation_LeftShoulder          BodyLocation = 23
	BodyLocation_LeftBicep             BodyLocation = 24
	BodyLocation_LeftTricep            BodyLocation = 25
	BodyLocation_LeftBrachioradialis   BodyLocation = 26
	BodyLocation_LeftForearmExtensors  BodyLocation = 27
	BodyLocation_RightArm              BodyLocation = 28
	BodyLocation_RightShoulder         BodyLocation = 29
	BodyLocation_RightBicep            BodyLocation = 30
	BodyLocation_RightTricep           BodyLocation = 31
	BodyLocation_RightBrachioradialis  BodyLocation = 32
	BodyLocation_RightForearmExtensors BodyLocation = 33
	BodyLocation_Neck                  BodyLocation = 34
	BodyLocation_Throat                BodyLocation = 35
	BodyLocation_WaistMidBack          BodyLocation = 36
	BodyLocation_WaistFront            BodyLocation = 37
	BodyLocation_WaistLeft             BodyLocation = 38
	BodyLocation_WaistRight            BodyLocation = 39
	BodyLocation_Invalid               BodyLocation = 255
)

var BodyLocationNames = map[BodyLocation]string{
	BodyLocation_LeftLeg:               "LeftLeg",
	BodyLocation_LeftCalf:              "LeftCalf",
	BodyLocation_LeftShin:              "LeftShin",
	BodyLocation_LeftHamstring:         "LeftHamstring",
	BodyLocation_LeftQuad:              "LeftQuad",
	BodyLocation_LeftGlute:             "LeftGlute",
	BodyLocation_RightLeg:              "RightLeg",
	BodyLocation_RightCalf:             "RightCalf",
	BodyLocation_RightShin:             "RightShin",
	BodyLocation_RightHamstring:        "RightHamstring",
	BodyLocation_RightQuad:             "RightQuad",
	BodyLocation_RightGlute:            "RightGlute",
	BodyLocation_TorsoBack:             "TorsoBack",
	BodyLocation_LeftLowerBack:         "LeftLowerBack",
	BodyLocation_LeftUpperBack:         "LeftUpperBack",
	BodyLocation_RightLowerBack:        "RightLowerBack",
	BodyLocation_RightUpperBack:        "RightUpperBack",
	BodyLocation_TorsoFront:            "TorsoFront",
	BodyLocation_LeftAbdomen:           "LeftAbdomen",
	BodyLocation_LeftChest:             "LeftChest",
	BodyLocation_RightAbdomen:          "RightAbdomen",
	BodyLocation_RightChest:            "RightChest",
	BodyLocation_LeftArm:               "LeftArm",
	BodyLocation_LeftShoulder:          "LeftShoulder",
	BodyLocation_LeftBicep:             "LeftBicep",
	BodyLocation_LeftTricep:            "LeftTricep",
	BodyLocation_LeftBrachioradialis:   "LeftBrachioradialis",
	BodyLocation_LeftForearmExtensors:  "LeftForearmExtensors",
	BodyLocation_RightArm:              "RightArm",
	BodyLocation_RightShoulder:         "RightShoulder",
	BodyLocation_RightBicep:            "RightBicep",
	BodyLocation_RightTricep:           "RightTricep",
	BodyLocation_RightBrachioradialis:  "RightBrachioradialis",
	BodyLocation_RightForearmExtensors: "RightForearmExtensors",
	BodyLocation_Neck:                  "Neck",
	BodyLocation_Throat:                "Throat",
	BodyLocation_WaistMidBack:          "WaistMidBack",
	BodyLocation_WaistFront:            "WaistFront",
	BodyLocation_WaistLeft:             "WaistLeft",
	BodyLocation_WaistRight:            "WaistRight",
}

func (k BodyLocation) String() string {
	if uint(k) < uint(len(BodyLocationNames)) {
		if v, ok := BodyLocationNames[k]; ok {
			return v
		}
	}
	return "BodyLocation(" + strconv.Itoa(int(k)) + ")"
}

var BodyLocationValues = map[string]BodyLocation{
	"LeftLeg":               BodyLocation_LeftLeg,
	"LeftCalf":              BodyLocation_LeftCalf,
	"LeftShin":              BodyLocation_LeftShin,
	"LeftHamstring":         BodyLocation_LeftHamstring,
	"LeftQuad":              BodyLocation_LeftQuad,
	"LeftGlute":             BodyLocation_LeftGlute,
	"RightLeg":              BodyLocation_RightLeg,
	"RightCalf":             BodyLocation_RightCalf,
	"RightShin":             BodyLocation_RightShin,
	"RightHamstring":        BodyLocation_RightHamstring,
	"RightQuad":             BodyLocation_RightQuad,
	"RightGlute":            BodyLocation_RightGlute,
	"TorsoBack":             BodyLocation_TorsoBack,
	"LeftLowerBack":         BodyLocation_LeftLowerBack,
	"LeftUpperBack":         BodyLocation_LeftUpperBack,
	"RightLowerBack":        BodyLocation_RightLowerBack,
	"RightUpperBack":        BodyLocation_RightUpperBack,
	"TorsoFront":            BodyLocation_TorsoFront,
	"LeftAbdomen":           BodyLocation_LeftAbdomen,
	"LeftChest":             BodyLocation_LeftChest,
	"RightAbdomen":          BodyLocation_RightAbdomen,
	"RightChest":            BodyLocation_RightChest,
	"LeftArm":               BodyLocation_LeftArm,
	"LeftShoulder":          BodyLocation_LeftShoulder,
	"LeftBicep":             BodyLocation_LeftBicep,
	"LeftTricep":            BodyLocation_LeftTricep,
	"LeftBrachioradialis":   BodyLocation_LeftBrachioradialis,
	"LeftForearmExtensors":  BodyLocation_LeftForearmExtensors,
	"RightArm":              BodyLocation_RightArm,
	"RightShoulder":         BodyLocation_RightShoulder,
	"RightBicep":            BodyLocation_RightBicep,
	"RightTricep":           BodyLocation_RightTricep,
	"RightBrachioradialis":  BodyLocation_RightBrachioradialis,
	"RightForearmExtensors": BodyLocation_RightForearmExtensors,
	"Neck":                  BodyLocation_Neck,
	"Throat":                BodyLocation_Throat,
	"WaistMidBack":          BodyLocation_WaistMidBack,
	"WaistFront":            BodyLocation_WaistFront,
	"WaistLeft":             BodyLocation_WaistLeft,
	"WaistRight":            BodyLocation_WaistRight,
}

func BodyLocationFromString(s string) BodyLocation {
	if v, ok := BodyLocationValues[s]; ok {
		return v
	}
	return BodyLocation(BodyLocation_Invalid)
}
