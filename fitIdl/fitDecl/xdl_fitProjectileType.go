package fitDecl

import strconv "strconv"

type ProjectileType byte

const (
	ProjectileType_Arrow           ProjectileType = 0
	ProjectileType_RifleCartridge  ProjectileType = 1
	ProjectileType_PistolCartridge ProjectileType = 2
	ProjectileType_Shotshell       ProjectileType = 3
	ProjectileType_AirRiflePellet  ProjectileType = 4
	ProjectileType_Other           ProjectileType = 5
	ProjectileType_Invalid         ProjectileType = 255
)

var ProjectileTypeNames = map[ProjectileType]string{
	ProjectileType_Arrow:           "Arrow",
	ProjectileType_RifleCartridge:  "RifleCartridge",
	ProjectileType_PistolCartridge: "PistolCartridge",
	ProjectileType_Shotshell:       "Shotshell",
	ProjectileType_AirRiflePellet:  "AirRiflePellet",
	ProjectileType_Other:           "Other",
}

func (k ProjectileType) String() string {
	if uint(k) < uint(len(ProjectileTypeNames)) {
		if v, ok := ProjectileTypeNames[k]; ok {
			return v
		}
	}
	return "ProjectileType(" + strconv.Itoa(int(k)) + ")"
}

var ProjectileTypeValues = map[string]ProjectileType{
	"Arrow":           ProjectileType_Arrow,
	"RifleCartridge":  ProjectileType_RifleCartridge,
	"PistolCartridge": ProjectileType_PistolCartridge,
	"Shotshell":       ProjectileType_Shotshell,
	"AirRiflePellet":  ProjectileType_AirRiflePellet,
	"Other":           ProjectileType_Other,
}

func ProjectileTypeFromString(s string) ProjectileType {
	if v, ok := ProjectileTypeValues[s]; ok {
		return v
	}
	return ProjectileType(ProjectileType_Invalid)
}
