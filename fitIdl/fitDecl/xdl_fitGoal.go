package fitDecl

import strconv "strconv"

type Goal byte

const (
	Goal_Time          Goal = 0
	Goal_Distance      Goal = 1
	Goal_Calories      Goal = 2
	Goal_Frequency     Goal = 3
	Goal_Steps         Goal = 4
	Goal_Ascent        Goal = 5
	Goal_ActiveMinutes Goal = 6
	Goal_Invalid       Goal = 255
)

var GoalNames = map[Goal]string{
	Goal_Time:          "Time",
	Goal_Distance:      "Distance",
	Goal_Calories:      "Calories",
	Goal_Frequency:     "Frequency",
	Goal_Steps:         "Steps",
	Goal_Ascent:        "Ascent",
	Goal_ActiveMinutes: "ActiveMinutes",
}

func (k Goal) String() string {
	if uint(k) < uint(len(GoalNames)) {
		if v, ok := GoalNames[k]; ok {
			return v
		}
	}
	return "Goal(" + strconv.Itoa(int(k)) + ")"
}

var GoalValues = map[string]Goal{
	"Time":          Goal_Time,
	"Distance":      Goal_Distance,
	"Calories":      Goal_Calories,
	"Frequency":     Goal_Frequency,
	"Steps":         Goal_Steps,
	"Ascent":        Goal_Ascent,
	"ActiveMinutes": Goal_ActiveMinutes,
}

func GoalFromString(s string) Goal {
	if v, ok := GoalValues[s]; ok {
		return v
	}
	return Goal(Goal_Invalid)
}
