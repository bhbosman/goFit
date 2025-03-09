package fitDecl

import strconv "strconv"

type WktStepDuration byte

const (
	WktStepDuration_Time                               WktStepDuration = 0
	WktStepDuration_Distance                           WktStepDuration = 1
	WktStepDuration_HrLessThan                         WktStepDuration = 2
	WktStepDuration_HrGreaterThan                      WktStepDuration = 3
	WktStepDuration_Calories                           WktStepDuration = 4
	WktStepDuration_Open                               WktStepDuration = 5
	WktStepDuration_RepeatUntilStepsCmplt              WktStepDuration = 6
	WktStepDuration_RepeatUntilTime                    WktStepDuration = 7
	WktStepDuration_RepeatUntilDistance                WktStepDuration = 8
	WktStepDuration_RepeatUntilCalories                WktStepDuration = 9
	WktStepDuration_RepeatUntilHrLessThan              WktStepDuration = 10
	WktStepDuration_RepeatUntilHrGreaterThan           WktStepDuration = 11
	WktStepDuration_RepeatUntilPowerLessThan           WktStepDuration = 12
	WktStepDuration_RepeatUntilPowerGreaterThan        WktStepDuration = 13
	WktStepDuration_PowerLessThan                      WktStepDuration = 14
	WktStepDuration_PowerGreaterThan                   WktStepDuration = 15
	WktStepDuration_TrainingPeaksTss                   WktStepDuration = 16
	WktStepDuration_RepeatUntilPowerLastLapLessThan    WktStepDuration = 17
	WktStepDuration_RepeatUntilMaxPowerLastLapLessThan WktStepDuration = 18
	WktStepDuration_Power3sLessThan                    WktStepDuration = 19
	WktStepDuration_Power10sLessThan                   WktStepDuration = 20
	WktStepDuration_Power30sLessThan                   WktStepDuration = 21
	WktStepDuration_Power3sGreaterThan                 WktStepDuration = 22
	WktStepDuration_Power10sGreaterThan                WktStepDuration = 23
	WktStepDuration_Power30sGreaterThan                WktStepDuration = 24
	WktStepDuration_PowerLapLessThan                   WktStepDuration = 25
	WktStepDuration_PowerLapGreaterThan                WktStepDuration = 26
	WktStepDuration_RepeatUntilTrainingPeaksTss        WktStepDuration = 27
	WktStepDuration_RepetitionTime                     WktStepDuration = 28
	WktStepDuration_Reps                               WktStepDuration = 29
	WktStepDuration_TimeOnly                           WktStepDuration = 31
	WktStepDuration_Invalid                            WktStepDuration = 255
)

var WktStepDurationNames = map[WktStepDuration]string{
	WktStepDuration_Time:                               "Time",
	WktStepDuration_Distance:                           "Distance",
	WktStepDuration_HrLessThan:                         "HrLessThan",
	WktStepDuration_HrGreaterThan:                      "HrGreaterThan",
	WktStepDuration_Calories:                           "Calories",
	WktStepDuration_Open:                               "Open",
	WktStepDuration_RepeatUntilStepsCmplt:              "RepeatUntilStepsCmplt",
	WktStepDuration_RepeatUntilTime:                    "RepeatUntilTime",
	WktStepDuration_RepeatUntilDistance:                "RepeatUntilDistance",
	WktStepDuration_RepeatUntilCalories:                "RepeatUntilCalories",
	WktStepDuration_RepeatUntilHrLessThan:              "RepeatUntilHrLessThan",
	WktStepDuration_RepeatUntilHrGreaterThan:           "RepeatUntilHrGreaterThan",
	WktStepDuration_RepeatUntilPowerLessThan:           "RepeatUntilPowerLessThan",
	WktStepDuration_RepeatUntilPowerGreaterThan:        "RepeatUntilPowerGreaterThan",
	WktStepDuration_PowerLessThan:                      "PowerLessThan",
	WktStepDuration_PowerGreaterThan:                   "PowerGreaterThan",
	WktStepDuration_TrainingPeaksTss:                   "TrainingPeaksTss",
	WktStepDuration_RepeatUntilPowerLastLapLessThan:    "RepeatUntilPowerLastLapLessThan",
	WktStepDuration_RepeatUntilMaxPowerLastLapLessThan: "RepeatUntilMaxPowerLastLapLessThan",
	WktStepDuration_Power3sLessThan:                    "Power3sLessThan",
	WktStepDuration_Power10sLessThan:                   "Power10sLessThan",
	WktStepDuration_Power30sLessThan:                   "Power30sLessThan",
	WktStepDuration_Power3sGreaterThan:                 "Power3sGreaterThan",
	WktStepDuration_Power10sGreaterThan:                "Power10sGreaterThan",
	WktStepDuration_Power30sGreaterThan:                "Power30sGreaterThan",
	WktStepDuration_PowerLapLessThan:                   "PowerLapLessThan",
	WktStepDuration_PowerLapGreaterThan:                "PowerLapGreaterThan",
	WktStepDuration_RepeatUntilTrainingPeaksTss:        "RepeatUntilTrainingPeaksTss",
	WktStepDuration_RepetitionTime:                     "RepetitionTime",
	WktStepDuration_Reps:                               "Reps",
	WktStepDuration_TimeOnly:                           "TimeOnly",
}

func (k WktStepDuration) String() string {
	if uint(k) < uint(len(WktStepDurationNames)) {
		if v, ok := WktStepDurationNames[k]; ok {
			return v
		}
	}
	return "WktStepDuration(" + strconv.Itoa(int(k)) + ")"
}

var WktStepDurationValues = map[string]WktStepDuration{
	"Time":                               WktStepDuration_Time,
	"Distance":                           WktStepDuration_Distance,
	"HrLessThan":                         WktStepDuration_HrLessThan,
	"HrGreaterThan":                      WktStepDuration_HrGreaterThan,
	"Calories":                           WktStepDuration_Calories,
	"Open":                               WktStepDuration_Open,
	"RepeatUntilStepsCmplt":              WktStepDuration_RepeatUntilStepsCmplt,
	"RepeatUntilTime":                    WktStepDuration_RepeatUntilTime,
	"RepeatUntilDistance":                WktStepDuration_RepeatUntilDistance,
	"RepeatUntilCalories":                WktStepDuration_RepeatUntilCalories,
	"RepeatUntilHrLessThan":              WktStepDuration_RepeatUntilHrLessThan,
	"RepeatUntilHrGreaterThan":           WktStepDuration_RepeatUntilHrGreaterThan,
	"RepeatUntilPowerLessThan":           WktStepDuration_RepeatUntilPowerLessThan,
	"RepeatUntilPowerGreaterThan":        WktStepDuration_RepeatUntilPowerGreaterThan,
	"PowerLessThan":                      WktStepDuration_PowerLessThan,
	"PowerGreaterThan":                   WktStepDuration_PowerGreaterThan,
	"TrainingPeaksTss":                   WktStepDuration_TrainingPeaksTss,
	"RepeatUntilPowerLastLapLessThan":    WktStepDuration_RepeatUntilPowerLastLapLessThan,
	"RepeatUntilMaxPowerLastLapLessThan": WktStepDuration_RepeatUntilMaxPowerLastLapLessThan,
	"Power3sLessThan":                    WktStepDuration_Power3sLessThan,
	"Power10sLessThan":                   WktStepDuration_Power10sLessThan,
	"Power30sLessThan":                   WktStepDuration_Power30sLessThan,
	"Power3sGreaterThan":                 WktStepDuration_Power3sGreaterThan,
	"Power10sGreaterThan":                WktStepDuration_Power10sGreaterThan,
	"Power30sGreaterThan":                WktStepDuration_Power30sGreaterThan,
	"PowerLapLessThan":                   WktStepDuration_PowerLapLessThan,
	"PowerLapGreaterThan":                WktStepDuration_PowerLapGreaterThan,
	"RepeatUntilTrainingPeaksTss":        WktStepDuration_RepeatUntilTrainingPeaksTss,
	"RepetitionTime":                     WktStepDuration_RepetitionTime,
	"Reps":                               WktStepDuration_Reps,
	"TimeOnly":                           WktStepDuration_TimeOnly,
}

func WktStepDurationFromString(s string) WktStepDuration {
	if v, ok := WktStepDurationValues[s]; ok {
		return v
	}
	return WktStepDuration(WktStepDuration_Invalid)
}
