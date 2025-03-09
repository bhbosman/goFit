package fitDecl

import strconv "strconv"

type FitnessEquipmentState byte

const (
	FitnessEquipmentState_Ready   FitnessEquipmentState = 0
	FitnessEquipmentState_InUse   FitnessEquipmentState = 1
	FitnessEquipmentState_Paused  FitnessEquipmentState = 2
	FitnessEquipmentState_Unknown FitnessEquipmentState = 3
	FitnessEquipmentState_Invalid FitnessEquipmentState = 255
)

var FitnessEquipmentStateNames = map[FitnessEquipmentState]string{
	FitnessEquipmentState_Ready:   "Ready",
	FitnessEquipmentState_InUse:   "InUse",
	FitnessEquipmentState_Paused:  "Paused",
	FitnessEquipmentState_Unknown: "Unknown",
}

func (k FitnessEquipmentState) String() string {
	if uint(k) < uint(len(FitnessEquipmentStateNames)) {
		if v, ok := FitnessEquipmentStateNames[k]; ok {
			return v
		}
	}
	return "FitnessEquipmentState(" + strconv.Itoa(int(k)) + ")"
}

var FitnessEquipmentStateValues = map[string]FitnessEquipmentState{
	"Ready":   FitnessEquipmentState_Ready,
	"InUse":   FitnessEquipmentState_InUse,
	"Paused":  FitnessEquipmentState_Paused,
	"Unknown": FitnessEquipmentState_Unknown,
}

func FitnessEquipmentStateFromString(s string) FitnessEquipmentState {
	if v, ok := FitnessEquipmentStateValues[s]; ok {
		return v
	}
	return FitnessEquipmentState(FitnessEquipmentState_Invalid)
}
