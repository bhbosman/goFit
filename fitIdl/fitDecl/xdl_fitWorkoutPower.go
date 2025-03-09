package fitDecl

import strconv "strconv"

type WorkoutPower uint32

const (
	WorkoutPower_WattsOffset WorkoutPower = 1000
	WorkoutPower_Invalid     WorkoutPower = 4294967295
)

var WorkoutPowerNames = map[WorkoutPower]string{
	WorkoutPower_WattsOffset: "WattsOffset",
}

func (k WorkoutPower) String() string {
	if uint(k) < uint(len(WorkoutPowerNames)) {
		if v, ok := WorkoutPowerNames[k]; ok {
			return v
		}
	}
	return "WorkoutPower(" + strconv.Itoa(int(k)) + ")"
}

var WorkoutPowerValues = map[string]WorkoutPower{
	"WattsOffset": WorkoutPower_WattsOffset,
}

func WorkoutPowerFromString(s string) WorkoutPower {
	if v, ok := WorkoutPowerValues[s]; ok {
		return v
	}
	return WorkoutPower(WorkoutPower_Invalid)
}
