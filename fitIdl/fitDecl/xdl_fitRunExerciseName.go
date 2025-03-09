package fitDecl

import strconv "strconv"

type RunExerciseName uint16

const (
	RunExerciseName_Run     RunExerciseName = 0
	RunExerciseName_Walk    RunExerciseName = 1
	RunExerciseName_Jog     RunExerciseName = 2
	RunExerciseName_Sprint  RunExerciseName = 3
	RunExerciseName_Invalid RunExerciseName = 65535
)

var RunExerciseNameNames = map[RunExerciseName]string{
	RunExerciseName_Run:    "Run",
	RunExerciseName_Walk:   "Walk",
	RunExerciseName_Jog:    "Jog",
	RunExerciseName_Sprint: "Sprint",
}

func (k RunExerciseName) String() string {
	if uint(k) < uint(len(RunExerciseNameNames)) {
		if v, ok := RunExerciseNameNames[k]; ok {
			return v
		}
	}
	return "RunExerciseName(" + strconv.Itoa(int(k)) + ")"
}

var RunExerciseNameValues = map[string]RunExerciseName{
	"Run":    RunExerciseName_Run,
	"Walk":   RunExerciseName_Walk,
	"Jog":    RunExerciseName_Jog,
	"Sprint": RunExerciseName_Sprint,
}

func RunExerciseNameFromString(s string) RunExerciseName {
	if v, ok := RunExerciseNameValues[s]; ok {
		return v
	}
	return RunExerciseName(RunExerciseName_Invalid)
}
