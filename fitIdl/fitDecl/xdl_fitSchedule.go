package fitDecl

import strconv "strconv"

type Schedule byte

const (
	Schedule_Workout Schedule = 0
	Schedule_Course  Schedule = 1
	Schedule_Invalid Schedule = 255
)

var ScheduleNames = map[Schedule]string{
	Schedule_Workout: "Workout",
	Schedule_Course:  "Course",
}

func (k Schedule) String() string {
	if uint(k) < uint(len(ScheduleNames)) {
		if v, ok := ScheduleNames[k]; ok {
			return v
		}
	}
	return "Schedule(" + strconv.Itoa(int(k)) + ")"
}

var ScheduleValues = map[string]Schedule{
	"Workout": Schedule_Workout,
	"Course":  Schedule_Course,
}

func ScheduleFromString(s string) Schedule {
	if v, ok := ScheduleValues[s]; ok {
		return v
	}
	return Schedule(Schedule_Invalid)
}
