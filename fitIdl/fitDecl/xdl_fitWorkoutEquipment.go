package fitDecl

import strconv "strconv"

type WorkoutEquipment byte

const (
	WorkoutEquipment_None          WorkoutEquipment = 0
	WorkoutEquipment_SwimFins      WorkoutEquipment = 1
	WorkoutEquipment_SwimKickboard WorkoutEquipment = 2
	WorkoutEquipment_SwimPaddles   WorkoutEquipment = 3
	WorkoutEquipment_SwimPullBuoy  WorkoutEquipment = 4
	WorkoutEquipment_SwimSnorkel   WorkoutEquipment = 5
	WorkoutEquipment_Invalid       WorkoutEquipment = 255
)

var WorkoutEquipmentNames = map[WorkoutEquipment]string{
	WorkoutEquipment_None:          "None",
	WorkoutEquipment_SwimFins:      "SwimFins",
	WorkoutEquipment_SwimKickboard: "SwimKickboard",
	WorkoutEquipment_SwimPaddles:   "SwimPaddles",
	WorkoutEquipment_SwimPullBuoy:  "SwimPullBuoy",
	WorkoutEquipment_SwimSnorkel:   "SwimSnorkel",
}

func (k WorkoutEquipment) String() string {
	if uint(k) < uint(len(WorkoutEquipmentNames)) {
		if v, ok := WorkoutEquipmentNames[k]; ok {
			return v
		}
	}
	return "WorkoutEquipment(" + strconv.Itoa(int(k)) + ")"
}

var WorkoutEquipmentValues = map[string]WorkoutEquipment{
	"None":          WorkoutEquipment_None,
	"SwimFins":      WorkoutEquipment_SwimFins,
	"SwimKickboard": WorkoutEquipment_SwimKickboard,
	"SwimPaddles":   WorkoutEquipment_SwimPaddles,
	"SwimPullBuoy":  WorkoutEquipment_SwimPullBuoy,
	"SwimSnorkel":   WorkoutEquipment_SwimSnorkel,
}

func WorkoutEquipmentFromString(s string) WorkoutEquipment {
	if v, ok := WorkoutEquipmentValues[s]; ok {
		return v
	}
	return WorkoutEquipment(WorkoutEquipment_Invalid)
}
