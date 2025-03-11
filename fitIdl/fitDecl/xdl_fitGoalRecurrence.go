package fitDecl

import strconv "strconv"

type GoalRecurrence byte

const (
	GoalRecurrence_Off     GoalRecurrence = 0
	GoalRecurrence_Daily   GoalRecurrence = 1
	GoalRecurrence_Weekly  GoalRecurrence = 2
	GoalRecurrence_Monthly GoalRecurrence = 3
	GoalRecurrence_Yearly  GoalRecurrence = 4
	GoalRecurrence_Custom  GoalRecurrence = 5
	GoalRecurrence_Invalid GoalRecurrence = 255
)

var GoalRecurrenceNames = map[GoalRecurrence]string{
	GoalRecurrence_Off:     "Off",
	GoalRecurrence_Daily:   "Daily",
	GoalRecurrence_Weekly:  "Weekly",
	GoalRecurrence_Monthly: "Monthly",
	GoalRecurrence_Yearly:  "Yearly",
	GoalRecurrence_Custom:  "Custom",
}

func (k GoalRecurrence) String() string {
	if uint(k) < uint(len(GoalRecurrenceNames)) {
		if v, ok := GoalRecurrenceNames[k]; ok {
			return v
		}
	}
	return "GoalRecurrence(" + strconv.Itoa(int(k)) + ")"
}

var GoalRecurrenceValues = map[string]GoalRecurrence{
	"Off":     GoalRecurrence_Off,
	"Daily":   GoalRecurrence_Daily,
	"Weekly":  GoalRecurrence_Weekly,
	"Monthly": GoalRecurrence_Monthly,
	"Yearly":  GoalRecurrence_Yearly,
	"Custom":  GoalRecurrence_Custom,
}

func GoalRecurrenceFromString(s string) GoalRecurrence {
	if v, ok := GoalRecurrenceValues[s]; ok {
		return v
	}
	return GoalRecurrence(GoalRecurrence_Invalid)
}
