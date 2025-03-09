package fitDecl

import strconv "strconv"

type ExerciseCategory uint16

const (
	ExerciseCategory_BenchPress        ExerciseCategory = 0
	ExerciseCategory_CalfRaise         ExerciseCategory = 1
	ExerciseCategory_Cardio            ExerciseCategory = 2
	ExerciseCategory_Carry             ExerciseCategory = 3
	ExerciseCategory_Chop              ExerciseCategory = 4
	ExerciseCategory_Core              ExerciseCategory = 5
	ExerciseCategory_Crunch            ExerciseCategory = 6
	ExerciseCategory_Curl              ExerciseCategory = 7
	ExerciseCategory_Deadlift          ExerciseCategory = 8
	ExerciseCategory_Flye              ExerciseCategory = 9
	ExerciseCategory_HipRaise          ExerciseCategory = 10
	ExerciseCategory_HipStability      ExerciseCategory = 11
	ExerciseCategory_HipSwing          ExerciseCategory = 12
	ExerciseCategory_Hyperextension    ExerciseCategory = 13
	ExerciseCategory_LateralRaise      ExerciseCategory = 14
	ExerciseCategory_LegCurl           ExerciseCategory = 15
	ExerciseCategory_LegRaise          ExerciseCategory = 16
	ExerciseCategory_Lunge             ExerciseCategory = 17
	ExerciseCategory_OlympicLift       ExerciseCategory = 18
	ExerciseCategory_Plank             ExerciseCategory = 19
	ExerciseCategory_Plyo              ExerciseCategory = 20
	ExerciseCategory_PullUp            ExerciseCategory = 21
	ExerciseCategory_PushUp            ExerciseCategory = 22
	ExerciseCategory_Row               ExerciseCategory = 23
	ExerciseCategory_ShoulderPress     ExerciseCategory = 24
	ExerciseCategory_ShoulderStability ExerciseCategory = 25
	ExerciseCategory_Shrug             ExerciseCategory = 26
	ExerciseCategory_SitUp             ExerciseCategory = 27
	ExerciseCategory_Squat             ExerciseCategory = 28
	ExerciseCategory_TotalBody         ExerciseCategory = 29
	ExerciseCategory_TricepsExtension  ExerciseCategory = 30
	ExerciseCategory_WarmUp            ExerciseCategory = 31
	ExerciseCategory_Run               ExerciseCategory = 32
	ExerciseCategory_Unknown           ExerciseCategory = 65534
	ExerciseCategory_Invalid           ExerciseCategory = 65535
)

var ExerciseCategoryNames = map[ExerciseCategory]string{
	ExerciseCategory_BenchPress:        "BenchPress",
	ExerciseCategory_CalfRaise:         "CalfRaise",
	ExerciseCategory_Cardio:            "Cardio",
	ExerciseCategory_Carry:             "Carry",
	ExerciseCategory_Chop:              "Chop",
	ExerciseCategory_Core:              "Core",
	ExerciseCategory_Crunch:            "Crunch",
	ExerciseCategory_Curl:              "Curl",
	ExerciseCategory_Deadlift:          "Deadlift",
	ExerciseCategory_Flye:              "Flye",
	ExerciseCategory_HipRaise:          "HipRaise",
	ExerciseCategory_HipStability:      "HipStability",
	ExerciseCategory_HipSwing:          "HipSwing",
	ExerciseCategory_Hyperextension:    "Hyperextension",
	ExerciseCategory_LateralRaise:      "LateralRaise",
	ExerciseCategory_LegCurl:           "LegCurl",
	ExerciseCategory_LegRaise:          "LegRaise",
	ExerciseCategory_Lunge:             "Lunge",
	ExerciseCategory_OlympicLift:       "OlympicLift",
	ExerciseCategory_Plank:             "Plank",
	ExerciseCategory_Plyo:              "Plyo",
	ExerciseCategory_PullUp:            "PullUp",
	ExerciseCategory_PushUp:            "PushUp",
	ExerciseCategory_Row:               "Row",
	ExerciseCategory_ShoulderPress:     "ShoulderPress",
	ExerciseCategory_ShoulderStability: "ShoulderStability",
	ExerciseCategory_Shrug:             "Shrug",
	ExerciseCategory_SitUp:             "SitUp",
	ExerciseCategory_Squat:             "Squat",
	ExerciseCategory_TotalBody:         "TotalBody",
	ExerciseCategory_TricepsExtension:  "TricepsExtension",
	ExerciseCategory_WarmUp:            "WarmUp",
	ExerciseCategory_Run:               "Run",
	ExerciseCategory_Unknown:           "Unknown",
}

func (k ExerciseCategory) String() string {
	if uint(k) < uint(len(ExerciseCategoryNames)) {
		if v, ok := ExerciseCategoryNames[k]; ok {
			return v
		}
	}
	return "ExerciseCategory(" + strconv.Itoa(int(k)) + ")"
}

var ExerciseCategoryValues = map[string]ExerciseCategory{
	"BenchPress":        ExerciseCategory_BenchPress,
	"CalfRaise":         ExerciseCategory_CalfRaise,
	"Cardio":            ExerciseCategory_Cardio,
	"Carry":             ExerciseCategory_Carry,
	"Chop":              ExerciseCategory_Chop,
	"Core":              ExerciseCategory_Core,
	"Crunch":            ExerciseCategory_Crunch,
	"Curl":              ExerciseCategory_Curl,
	"Deadlift":          ExerciseCategory_Deadlift,
	"Flye":              ExerciseCategory_Flye,
	"HipRaise":          ExerciseCategory_HipRaise,
	"HipStability":      ExerciseCategory_HipStability,
	"HipSwing":          ExerciseCategory_HipSwing,
	"Hyperextension":    ExerciseCategory_Hyperextension,
	"LateralRaise":      ExerciseCategory_LateralRaise,
	"LegCurl":           ExerciseCategory_LegCurl,
	"LegRaise":          ExerciseCategory_LegRaise,
	"Lunge":             ExerciseCategory_Lunge,
	"OlympicLift":       ExerciseCategory_OlympicLift,
	"Plank":             ExerciseCategory_Plank,
	"Plyo":              ExerciseCategory_Plyo,
	"PullUp":            ExerciseCategory_PullUp,
	"PushUp":            ExerciseCategory_PushUp,
	"Row":               ExerciseCategory_Row,
	"ShoulderPress":     ExerciseCategory_ShoulderPress,
	"ShoulderStability": ExerciseCategory_ShoulderStability,
	"Shrug":             ExerciseCategory_Shrug,
	"SitUp":             ExerciseCategory_SitUp,
	"Squat":             ExerciseCategory_Squat,
	"TotalBody":         ExerciseCategory_TotalBody,
	"TricepsExtension":  ExerciseCategory_TricepsExtension,
	"WarmUp":            ExerciseCategory_WarmUp,
	"Run":               ExerciseCategory_Run,
	"Unknown":           ExerciseCategory_Unknown,
}

func ExerciseCategoryFromString(s string) ExerciseCategory {
	if v, ok := ExerciseCategoryValues[s]; ok {
		return v
	}
	return ExerciseCategory(ExerciseCategory_Invalid)
}
