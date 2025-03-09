package fitDecl

import strconv "strconv"

type AttitudeStage byte

const (
	AttitudeStage_Failed   AttitudeStage = 0
	AttitudeStage_Aligning AttitudeStage = 1
	AttitudeStage_Degraded AttitudeStage = 2
	AttitudeStage_Valid    AttitudeStage = 3
	AttitudeStage_Invalid  AttitudeStage = 255
)

var AttitudeStageNames = map[AttitudeStage]string{
	AttitudeStage_Failed:   "Failed",
	AttitudeStage_Aligning: "Aligning",
	AttitudeStage_Degraded: "Degraded",
	AttitudeStage_Valid:    "Valid",
}

func (k AttitudeStage) String() string {
	if uint(k) < uint(len(AttitudeStageNames)) {
		if v, ok := AttitudeStageNames[k]; ok {
			return v
		}
	}
	return "AttitudeStage(" + strconv.Itoa(int(k)) + ")"
}

var AttitudeStageValues = map[string]AttitudeStage{
	"Failed":   AttitudeStage_Failed,
	"Aligning": AttitudeStage_Aligning,
	"Degraded": AttitudeStage_Degraded,
	"Valid":    AttitudeStage_Valid,
}

func AttitudeStageFromString(s string) AttitudeStage {
	if v, ok := AttitudeStageValues[s]; ok {
		return v
	}
	return AttitudeStage(AttitudeStage_Invalid)
}
