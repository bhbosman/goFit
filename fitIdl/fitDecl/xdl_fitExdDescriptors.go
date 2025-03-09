package fitDecl

import strconv "strconv"

type ExdDescriptors byte

const (
	ExdDescriptors_BikeLightBatteryStatus           ExdDescriptors = 0
	ExdDescriptors_BeamAngleStatus                  ExdDescriptors = 1
	ExdDescriptors_BateryLevel                      ExdDescriptors = 2
	ExdDescriptors_LightNetworkMode                 ExdDescriptors = 3
	ExdDescriptors_NumberLightsConnected            ExdDescriptors = 4
	ExdDescriptors_Cadence                          ExdDescriptors = 5
	ExdDescriptors_Distance                         ExdDescriptors = 6
	ExdDescriptors_EstimatedTimeOfArrival           ExdDescriptors = 7
	ExdDescriptors_Heading                          ExdDescriptors = 8
	ExdDescriptors_Time                             ExdDescriptors = 9
	ExdDescriptors_BatteryLevel                     ExdDescriptors = 10
	ExdDescriptors_TrainerResistance                ExdDescriptors = 11
	ExdDescriptors_TrainerTargetPower               ExdDescriptors = 12
	ExdDescriptors_TimeSeated                       ExdDescriptors = 13
	ExdDescriptors_TimeStanding                     ExdDescriptors = 14
	ExdDescriptors_Elevation                        ExdDescriptors = 15
	ExdDescriptors_Grade                            ExdDescriptors = 16
	ExdDescriptors_Ascent                           ExdDescriptors = 17
	ExdDescriptors_Descent                          ExdDescriptors = 18
	ExdDescriptors_VerticalSpeed                    ExdDescriptors = 19
	ExdDescriptors_Di2BatteryLevel                  ExdDescriptors = 20
	ExdDescriptors_FrontGear                        ExdDescriptors = 21
	ExdDescriptors_RearGear                         ExdDescriptors = 22
	ExdDescriptors_GearRatio                        ExdDescriptors = 23
	ExdDescriptors_HeartRate                        ExdDescriptors = 24
	ExdDescriptors_HeartRateZone                    ExdDescriptors = 25
	ExdDescriptors_TimeInHeartRateZone              ExdDescriptors = 26
	ExdDescriptors_HeartRateReserve                 ExdDescriptors = 27
	ExdDescriptors_Calories                         ExdDescriptors = 28
	ExdDescriptors_GpsAccuracy                      ExdDescriptors = 29
	ExdDescriptors_GpsSignalStrength                ExdDescriptors = 30
	ExdDescriptors_Temperature                      ExdDescriptors = 31
	ExdDescriptors_TimeOfDay                        ExdDescriptors = 32
	ExdDescriptors_Balance                          ExdDescriptors = 33
	ExdDescriptors_PedalSmoothness                  ExdDescriptors = 34
	ExdDescriptors_Power                            ExdDescriptors = 35
	ExdDescriptors_FunctionalThresholdPower         ExdDescriptors = 36
	ExdDescriptors_IntensityFactor                  ExdDescriptors = 37
	ExdDescriptors_Work                             ExdDescriptors = 38
	ExdDescriptors_PowerRatio                       ExdDescriptors = 39
	ExdDescriptors_NormalizedPower                  ExdDescriptors = 40
	ExdDescriptors_TrainingStressScore              ExdDescriptors = 41
	ExdDescriptors_TimeOnZone                       ExdDescriptors = 42
	ExdDescriptors_Speed                            ExdDescriptors = 43
	ExdDescriptors_Laps                             ExdDescriptors = 44
	ExdDescriptors_Reps                             ExdDescriptors = 45
	ExdDescriptors_WorkoutStep                      ExdDescriptors = 46
	ExdDescriptors_CourseDistance                   ExdDescriptors = 47
	ExdDescriptors_NavigationDistance               ExdDescriptors = 48
	ExdDescriptors_CourseEstimatedTimeOfArrival     ExdDescriptors = 49
	ExdDescriptors_NavigationEstimatedTimeOfArrival ExdDescriptors = 50
	ExdDescriptors_CourseTime                       ExdDescriptors = 51
	ExdDescriptors_NavigationTime                   ExdDescriptors = 52
	ExdDescriptors_CourseHeading                    ExdDescriptors = 53
	ExdDescriptors_NavigationHeading                ExdDescriptors = 54
	ExdDescriptors_PowerZone                        ExdDescriptors = 55
	ExdDescriptors_TorqueEffectiveness              ExdDescriptors = 56
	ExdDescriptors_TimerTime                        ExdDescriptors = 57
	ExdDescriptors_PowerWeightRatio                 ExdDescriptors = 58
	ExdDescriptors_LeftPlatformCenterOffset         ExdDescriptors = 59
	ExdDescriptors_RightPlatformCenterOffset        ExdDescriptors = 60
	ExdDescriptors_LeftPowerPhaseStartAngle         ExdDescriptors = 61
	ExdDescriptors_RightPowerPhaseStartAngle        ExdDescriptors = 62
	ExdDescriptors_LeftPowerPhaseFinishAngle        ExdDescriptors = 63
	ExdDescriptors_RightPowerPhaseFinishAngle       ExdDescriptors = 64
	ExdDescriptors_Gears                            ExdDescriptors = 65
	ExdDescriptors_Pace                             ExdDescriptors = 66
	ExdDescriptors_TrainingEffect                   ExdDescriptors = 67
	ExdDescriptors_VerticalOscillation              ExdDescriptors = 68
	ExdDescriptors_VerticalRatio                    ExdDescriptors = 69
	ExdDescriptors_GroundContactTime                ExdDescriptors = 70
	ExdDescriptors_LeftGroundContactTimeBalance     ExdDescriptors = 71
	ExdDescriptors_RightGroundContactTimeBalance    ExdDescriptors = 72
	ExdDescriptors_StrideLength                     ExdDescriptors = 73
	ExdDescriptors_RunningCadence                   ExdDescriptors = 74
	ExdDescriptors_PerformanceCondition             ExdDescriptors = 75
	ExdDescriptors_CourseType                       ExdDescriptors = 76
	ExdDescriptors_TimeInPowerZone                  ExdDescriptors = 77
	ExdDescriptors_NavigationTurn                   ExdDescriptors = 78
	ExdDescriptors_CourseLocation                   ExdDescriptors = 79
	ExdDescriptors_NavigationLocation               ExdDescriptors = 80
	ExdDescriptors_Compass                          ExdDescriptors = 81
	ExdDescriptors_GearCombo                        ExdDescriptors = 82
	ExdDescriptors_MuscleOxygen                     ExdDescriptors = 83
	ExdDescriptors_Icon                             ExdDescriptors = 84
	ExdDescriptors_CompassHeading                   ExdDescriptors = 85
	ExdDescriptors_GpsHeading                       ExdDescriptors = 86
	ExdDescriptors_GpsElevation                     ExdDescriptors = 87
	ExdDescriptors_AnaerobicTrainingEffect          ExdDescriptors = 88
	ExdDescriptors_Course                           ExdDescriptors = 89
	ExdDescriptors_OffCourse                        ExdDescriptors = 90
	ExdDescriptors_GlideRatio                       ExdDescriptors = 91
	ExdDescriptors_VerticalDistance                 ExdDescriptors = 92
	ExdDescriptors_Vmg                              ExdDescriptors = 93
	ExdDescriptors_AmbientPressure                  ExdDescriptors = 94
	ExdDescriptors_Pressure                         ExdDescriptors = 95
	ExdDescriptors_Vam                              ExdDescriptors = 96
	ExdDescriptors_Invalid                          ExdDescriptors = 255
)

var ExdDescriptorsNames = map[ExdDescriptors]string{
	ExdDescriptors_BikeLightBatteryStatus:           "BikeLightBatteryStatus",
	ExdDescriptors_BeamAngleStatus:                  "BeamAngleStatus",
	ExdDescriptors_BateryLevel:                      "BateryLevel",
	ExdDescriptors_LightNetworkMode:                 "LightNetworkMode",
	ExdDescriptors_NumberLightsConnected:            "NumberLightsConnected",
	ExdDescriptors_Cadence:                          "Cadence",
	ExdDescriptors_Distance:                         "Distance",
	ExdDescriptors_EstimatedTimeOfArrival:           "EstimatedTimeOfArrival",
	ExdDescriptors_Heading:                          "Heading",
	ExdDescriptors_Time:                             "Time",
	ExdDescriptors_BatteryLevel:                     "BatteryLevel",
	ExdDescriptors_TrainerResistance:                "TrainerResistance",
	ExdDescriptors_TrainerTargetPower:               "TrainerTargetPower",
	ExdDescriptors_TimeSeated:                       "TimeSeated",
	ExdDescriptors_TimeStanding:                     "TimeStanding",
	ExdDescriptors_Elevation:                        "Elevation",
	ExdDescriptors_Grade:                            "Grade",
	ExdDescriptors_Ascent:                           "Ascent",
	ExdDescriptors_Descent:                          "Descent",
	ExdDescriptors_VerticalSpeed:                    "VerticalSpeed",
	ExdDescriptors_Di2BatteryLevel:                  "Di2BatteryLevel",
	ExdDescriptors_FrontGear:                        "FrontGear",
	ExdDescriptors_RearGear:                         "RearGear",
	ExdDescriptors_GearRatio:                        "GearRatio",
	ExdDescriptors_HeartRate:                        "HeartRate",
	ExdDescriptors_HeartRateZone:                    "HeartRateZone",
	ExdDescriptors_TimeInHeartRateZone:              "TimeInHeartRateZone",
	ExdDescriptors_HeartRateReserve:                 "HeartRateReserve",
	ExdDescriptors_Calories:                         "Calories",
	ExdDescriptors_GpsAccuracy:                      "GpsAccuracy",
	ExdDescriptors_GpsSignalStrength:                "GpsSignalStrength",
	ExdDescriptors_Temperature:                      "Temperature",
	ExdDescriptors_TimeOfDay:                        "TimeOfDay",
	ExdDescriptors_Balance:                          "Balance",
	ExdDescriptors_PedalSmoothness:                  "PedalSmoothness",
	ExdDescriptors_Power:                            "Power",
	ExdDescriptors_FunctionalThresholdPower:         "FunctionalThresholdPower",
	ExdDescriptors_IntensityFactor:                  "IntensityFactor",
	ExdDescriptors_Work:                             "Work",
	ExdDescriptors_PowerRatio:                       "PowerRatio",
	ExdDescriptors_NormalizedPower:                  "NormalizedPower",
	ExdDescriptors_TrainingStressScore:              "TrainingStressScore",
	ExdDescriptors_TimeOnZone:                       "TimeOnZone",
	ExdDescriptors_Speed:                            "Speed",
	ExdDescriptors_Laps:                             "Laps",
	ExdDescriptors_Reps:                             "Reps",
	ExdDescriptors_WorkoutStep:                      "WorkoutStep",
	ExdDescriptors_CourseDistance:                   "CourseDistance",
	ExdDescriptors_NavigationDistance:               "NavigationDistance",
	ExdDescriptors_CourseEstimatedTimeOfArrival:     "CourseEstimatedTimeOfArrival",
	ExdDescriptors_NavigationEstimatedTimeOfArrival: "NavigationEstimatedTimeOfArrival",
	ExdDescriptors_CourseTime:                       "CourseTime",
	ExdDescriptors_NavigationTime:                   "NavigationTime",
	ExdDescriptors_CourseHeading:                    "CourseHeading",
	ExdDescriptors_NavigationHeading:                "NavigationHeading",
	ExdDescriptors_PowerZone:                        "PowerZone",
	ExdDescriptors_TorqueEffectiveness:              "TorqueEffectiveness",
	ExdDescriptors_TimerTime:                        "TimerTime",
	ExdDescriptors_PowerWeightRatio:                 "PowerWeightRatio",
	ExdDescriptors_LeftPlatformCenterOffset:         "LeftPlatformCenterOffset",
	ExdDescriptors_RightPlatformCenterOffset:        "RightPlatformCenterOffset",
	ExdDescriptors_LeftPowerPhaseStartAngle:         "LeftPowerPhaseStartAngle",
	ExdDescriptors_RightPowerPhaseStartAngle:        "RightPowerPhaseStartAngle",
	ExdDescriptors_LeftPowerPhaseFinishAngle:        "LeftPowerPhaseFinishAngle",
	ExdDescriptors_RightPowerPhaseFinishAngle:       "RightPowerPhaseFinishAngle",
	ExdDescriptors_Gears:                            "Gears",
	ExdDescriptors_Pace:                             "Pace",
	ExdDescriptors_TrainingEffect:                   "TrainingEffect",
	ExdDescriptors_VerticalOscillation:              "VerticalOscillation",
	ExdDescriptors_VerticalRatio:                    "VerticalRatio",
	ExdDescriptors_GroundContactTime:                "GroundContactTime",
	ExdDescriptors_LeftGroundContactTimeBalance:     "LeftGroundContactTimeBalance",
	ExdDescriptors_RightGroundContactTimeBalance:    "RightGroundContactTimeBalance",
	ExdDescriptors_StrideLength:                     "StrideLength",
	ExdDescriptors_RunningCadence:                   "RunningCadence",
	ExdDescriptors_PerformanceCondition:             "PerformanceCondition",
	ExdDescriptors_CourseType:                       "CourseType",
	ExdDescriptors_TimeInPowerZone:                  "TimeInPowerZone",
	ExdDescriptors_NavigationTurn:                   "NavigationTurn",
	ExdDescriptors_CourseLocation:                   "CourseLocation",
	ExdDescriptors_NavigationLocation:               "NavigationLocation",
	ExdDescriptors_Compass:                          "Compass",
	ExdDescriptors_GearCombo:                        "GearCombo",
	ExdDescriptors_MuscleOxygen:                     "MuscleOxygen",
	ExdDescriptors_Icon:                             "Icon",
	ExdDescriptors_CompassHeading:                   "CompassHeading",
	ExdDescriptors_GpsHeading:                       "GpsHeading",
	ExdDescriptors_GpsElevation:                     "GpsElevation",
	ExdDescriptors_AnaerobicTrainingEffect:          "AnaerobicTrainingEffect",
	ExdDescriptors_Course:                           "Course",
	ExdDescriptors_OffCourse:                        "OffCourse",
	ExdDescriptors_GlideRatio:                       "GlideRatio",
	ExdDescriptors_VerticalDistance:                 "VerticalDistance",
	ExdDescriptors_Vmg:                              "Vmg",
	ExdDescriptors_AmbientPressure:                  "AmbientPressure",
	ExdDescriptors_Pressure:                         "Pressure",
	ExdDescriptors_Vam:                              "Vam",
}

func (k ExdDescriptors) String() string {
	if uint(k) < uint(len(ExdDescriptorsNames)) {
		if v, ok := ExdDescriptorsNames[k]; ok {
			return v
		}
	}
	return "ExdDescriptors(" + strconv.Itoa(int(k)) + ")"
}

var ExdDescriptorsValues = map[string]ExdDescriptors{
	"BikeLightBatteryStatus":           ExdDescriptors_BikeLightBatteryStatus,
	"BeamAngleStatus":                  ExdDescriptors_BeamAngleStatus,
	"BateryLevel":                      ExdDescriptors_BateryLevel,
	"LightNetworkMode":                 ExdDescriptors_LightNetworkMode,
	"NumberLightsConnected":            ExdDescriptors_NumberLightsConnected,
	"Cadence":                          ExdDescriptors_Cadence,
	"Distance":                         ExdDescriptors_Distance,
	"EstimatedTimeOfArrival":           ExdDescriptors_EstimatedTimeOfArrival,
	"Heading":                          ExdDescriptors_Heading,
	"Time":                             ExdDescriptors_Time,
	"BatteryLevel":                     ExdDescriptors_BatteryLevel,
	"TrainerResistance":                ExdDescriptors_TrainerResistance,
	"TrainerTargetPower":               ExdDescriptors_TrainerTargetPower,
	"TimeSeated":                       ExdDescriptors_TimeSeated,
	"TimeStanding":                     ExdDescriptors_TimeStanding,
	"Elevation":                        ExdDescriptors_Elevation,
	"Grade":                            ExdDescriptors_Grade,
	"Ascent":                           ExdDescriptors_Ascent,
	"Descent":                          ExdDescriptors_Descent,
	"VerticalSpeed":                    ExdDescriptors_VerticalSpeed,
	"Di2BatteryLevel":                  ExdDescriptors_Di2BatteryLevel,
	"FrontGear":                        ExdDescriptors_FrontGear,
	"RearGear":                         ExdDescriptors_RearGear,
	"GearRatio":                        ExdDescriptors_GearRatio,
	"HeartRate":                        ExdDescriptors_HeartRate,
	"HeartRateZone":                    ExdDescriptors_HeartRateZone,
	"TimeInHeartRateZone":              ExdDescriptors_TimeInHeartRateZone,
	"HeartRateReserve":                 ExdDescriptors_HeartRateReserve,
	"Calories":                         ExdDescriptors_Calories,
	"GpsAccuracy":                      ExdDescriptors_GpsAccuracy,
	"GpsSignalStrength":                ExdDescriptors_GpsSignalStrength,
	"Temperature":                      ExdDescriptors_Temperature,
	"TimeOfDay":                        ExdDescriptors_TimeOfDay,
	"Balance":                          ExdDescriptors_Balance,
	"PedalSmoothness":                  ExdDescriptors_PedalSmoothness,
	"Power":                            ExdDescriptors_Power,
	"FunctionalThresholdPower":         ExdDescriptors_FunctionalThresholdPower,
	"IntensityFactor":                  ExdDescriptors_IntensityFactor,
	"Work":                             ExdDescriptors_Work,
	"PowerRatio":                       ExdDescriptors_PowerRatio,
	"NormalizedPower":                  ExdDescriptors_NormalizedPower,
	"TrainingStressScore":              ExdDescriptors_TrainingStressScore,
	"TimeOnZone":                       ExdDescriptors_TimeOnZone,
	"Speed":                            ExdDescriptors_Speed,
	"Laps":                             ExdDescriptors_Laps,
	"Reps":                             ExdDescriptors_Reps,
	"WorkoutStep":                      ExdDescriptors_WorkoutStep,
	"CourseDistance":                   ExdDescriptors_CourseDistance,
	"NavigationDistance":               ExdDescriptors_NavigationDistance,
	"CourseEstimatedTimeOfArrival":     ExdDescriptors_CourseEstimatedTimeOfArrival,
	"NavigationEstimatedTimeOfArrival": ExdDescriptors_NavigationEstimatedTimeOfArrival,
	"CourseTime":                       ExdDescriptors_CourseTime,
	"NavigationTime":                   ExdDescriptors_NavigationTime,
	"CourseHeading":                    ExdDescriptors_CourseHeading,
	"NavigationHeading":                ExdDescriptors_NavigationHeading,
	"PowerZone":                        ExdDescriptors_PowerZone,
	"TorqueEffectiveness":              ExdDescriptors_TorqueEffectiveness,
	"TimerTime":                        ExdDescriptors_TimerTime,
	"PowerWeightRatio":                 ExdDescriptors_PowerWeightRatio,
	"LeftPlatformCenterOffset":         ExdDescriptors_LeftPlatformCenterOffset,
	"RightPlatformCenterOffset":        ExdDescriptors_RightPlatformCenterOffset,
	"LeftPowerPhaseStartAngle":         ExdDescriptors_LeftPowerPhaseStartAngle,
	"RightPowerPhaseStartAngle":        ExdDescriptors_RightPowerPhaseStartAngle,
	"LeftPowerPhaseFinishAngle":        ExdDescriptors_LeftPowerPhaseFinishAngle,
	"RightPowerPhaseFinishAngle":       ExdDescriptors_RightPowerPhaseFinishAngle,
	"Gears":                            ExdDescriptors_Gears,
	"Pace":                             ExdDescriptors_Pace,
	"TrainingEffect":                   ExdDescriptors_TrainingEffect,
	"VerticalOscillation":              ExdDescriptors_VerticalOscillation,
	"VerticalRatio":                    ExdDescriptors_VerticalRatio,
	"GroundContactTime":                ExdDescriptors_GroundContactTime,
	"LeftGroundContactTimeBalance":     ExdDescriptors_LeftGroundContactTimeBalance,
	"RightGroundContactTimeBalance":    ExdDescriptors_RightGroundContactTimeBalance,
	"StrideLength":                     ExdDescriptors_StrideLength,
	"RunningCadence":                   ExdDescriptors_RunningCadence,
	"PerformanceCondition":             ExdDescriptors_PerformanceCondition,
	"CourseType":                       ExdDescriptors_CourseType,
	"TimeInPowerZone":                  ExdDescriptors_TimeInPowerZone,
	"NavigationTurn":                   ExdDescriptors_NavigationTurn,
	"CourseLocation":                   ExdDescriptors_CourseLocation,
	"NavigationLocation":               ExdDescriptors_NavigationLocation,
	"Compass":                          ExdDescriptors_Compass,
	"GearCombo":                        ExdDescriptors_GearCombo,
	"MuscleOxygen":                     ExdDescriptors_MuscleOxygen,
	"Icon":                             ExdDescriptors_Icon,
	"CompassHeading":                   ExdDescriptors_CompassHeading,
	"GpsHeading":                       ExdDescriptors_GpsHeading,
	"GpsElevation":                     ExdDescriptors_GpsElevation,
	"AnaerobicTrainingEffect":          ExdDescriptors_AnaerobicTrainingEffect,
	"Course":                           ExdDescriptors_Course,
	"OffCourse":                        ExdDescriptors_OffCourse,
	"GlideRatio":                       ExdDescriptors_GlideRatio,
	"VerticalDistance":                 ExdDescriptors_VerticalDistance,
	"Vmg":                              ExdDescriptors_Vmg,
	"AmbientPressure":                  ExdDescriptors_AmbientPressure,
	"Pressure":                         ExdDescriptors_Pressure,
	"Vam":                              ExdDescriptors_Vam,
}

func ExdDescriptorsFromString(s string) ExdDescriptors {
	if v, ok := ExdDescriptorsValues[s]; ok {
		return v
	}
	return ExdDescriptors(ExdDescriptors_Invalid)
}
