package fitDecl

import strconv "strconv"

type Event byte

const (
	Event_Timer                 Event = 0
	Event_Workout               Event = 3
	Event_WorkoutStep           Event = 4
	Event_PowerDown             Event = 5
	Event_PowerUp               Event = 6
	Event_OffCourse             Event = 7
	Event_Session               Event = 8
	Event_Lap                   Event = 9
	Event_CoursePoint           Event = 10
	Event_Battery               Event = 11
	Event_VirtualPartnerPace    Event = 12
	Event_HrHighAlert           Event = 13
	Event_HrLowAlert            Event = 14
	Event_SpeedHighAlert        Event = 15
	Event_SpeedLowAlert         Event = 16
	Event_CadHighAlert          Event = 17
	Event_CadLowAlert           Event = 18
	Event_PowerHighAlert        Event = 19
	Event_PowerLowAlert         Event = 20
	Event_RecoveryHr            Event = 21
	Event_BatteryLow            Event = 22
	Event_TimeDurationAlert     Event = 23
	Event_DistanceDurationAlert Event = 24
	Event_CalorieDurationAlert  Event = 25
	Event_Activity              Event = 26
	Event_FitnessEquipment      Event = 27
	Event_Length                Event = 28
	Event_UserMarker            Event = 32
	Event_SportPoint            Event = 33
	Event_Calibration           Event = 36
	Event_FrontGearChange       Event = 42
	Event_RearGearChange        Event = 43
	Event_RiderPositionChange   Event = 44
	Event_ElevHighAlert         Event = 45
	Event_ElevLowAlert          Event = 46
	Event_CommTimeout           Event = 47
	Event_AutoActivityDetect    Event = 54
	Event_DiveAlert             Event = 56
	Event_DiveGasSwitched       Event = 57
	Event_TankPressureReserve   Event = 71
	Event_TankPressureCritical  Event = 72
	Event_TankLost              Event = 73
	Event_RadarThreatAlert      Event = 75
	Event_TankBatteryLow        Event = 76
	Event_TankPodConnected      Event = 81
	Event_TankPodDisconnected   Event = 82
	Event_Invalid               Event = 255
)

var EventNames = map[Event]string{
	Event_Timer:                 "Timer",
	Event_Workout:               "Workout",
	Event_WorkoutStep:           "WorkoutStep",
	Event_PowerDown:             "PowerDown",
	Event_PowerUp:               "PowerUp",
	Event_OffCourse:             "OffCourse",
	Event_Session:               "Session",
	Event_Lap:                   "Lap",
	Event_CoursePoint:           "CoursePoint",
	Event_Battery:               "Battery",
	Event_VirtualPartnerPace:    "VirtualPartnerPace",
	Event_HrHighAlert:           "HrHighAlert",
	Event_HrLowAlert:            "HrLowAlert",
	Event_SpeedHighAlert:        "SpeedHighAlert",
	Event_SpeedLowAlert:         "SpeedLowAlert",
	Event_CadHighAlert:          "CadHighAlert",
	Event_CadLowAlert:           "CadLowAlert",
	Event_PowerHighAlert:        "PowerHighAlert",
	Event_PowerLowAlert:         "PowerLowAlert",
	Event_RecoveryHr:            "RecoveryHr",
	Event_BatteryLow:            "BatteryLow",
	Event_TimeDurationAlert:     "TimeDurationAlert",
	Event_DistanceDurationAlert: "DistanceDurationAlert",
	Event_CalorieDurationAlert:  "CalorieDurationAlert",
	Event_Activity:              "Activity",
	Event_FitnessEquipment:      "FitnessEquipment",
	Event_Length:                "Length",
	Event_UserMarker:            "UserMarker",
	Event_SportPoint:            "SportPoint",
	Event_Calibration:           "Calibration",
	Event_FrontGearChange:       "FrontGearChange",
	Event_RearGearChange:        "RearGearChange",
	Event_RiderPositionChange:   "RiderPositionChange",
	Event_ElevHighAlert:         "ElevHighAlert",
	Event_ElevLowAlert:          "ElevLowAlert",
	Event_CommTimeout:           "CommTimeout",
	Event_AutoActivityDetect:    "AutoActivityDetect",
	Event_DiveAlert:             "DiveAlert",
	Event_DiveGasSwitched:       "DiveGasSwitched",
	Event_TankPressureReserve:   "TankPressureReserve",
	Event_TankPressureCritical:  "TankPressureCritical",
	Event_TankLost:              "TankLost",
	Event_RadarThreatAlert:      "RadarThreatAlert",
	Event_TankBatteryLow:        "TankBatteryLow",
	Event_TankPodConnected:      "TankPodConnected",
	Event_TankPodDisconnected:   "TankPodDisconnected",
}

func (k Event) String() string {
	if uint(k) < uint(len(EventNames)) {
		if v, ok := EventNames[k]; ok {
			return v
		}
	}
	return "Event(" + strconv.Itoa(int(k)) + ")"
}

var EventValues = map[string]Event{
	"Timer":                 Event_Timer,
	"Workout":               Event_Workout,
	"WorkoutStep":           Event_WorkoutStep,
	"PowerDown":             Event_PowerDown,
	"PowerUp":               Event_PowerUp,
	"OffCourse":             Event_OffCourse,
	"Session":               Event_Session,
	"Lap":                   Event_Lap,
	"CoursePoint":           Event_CoursePoint,
	"Battery":               Event_Battery,
	"VirtualPartnerPace":    Event_VirtualPartnerPace,
	"HrHighAlert":           Event_HrHighAlert,
	"HrLowAlert":            Event_HrLowAlert,
	"SpeedHighAlert":        Event_SpeedHighAlert,
	"SpeedLowAlert":         Event_SpeedLowAlert,
	"CadHighAlert":          Event_CadHighAlert,
	"CadLowAlert":           Event_CadLowAlert,
	"PowerHighAlert":        Event_PowerHighAlert,
	"PowerLowAlert":         Event_PowerLowAlert,
	"RecoveryHr":            Event_RecoveryHr,
	"BatteryLow":            Event_BatteryLow,
	"TimeDurationAlert":     Event_TimeDurationAlert,
	"DistanceDurationAlert": Event_DistanceDurationAlert,
	"CalorieDurationAlert":  Event_CalorieDurationAlert,
	"Activity":              Event_Activity,
	"FitnessEquipment":      Event_FitnessEquipment,
	"Length":                Event_Length,
	"UserMarker":            Event_UserMarker,
	"SportPoint":            Event_SportPoint,
	"Calibration":           Event_Calibration,
	"FrontGearChange":       Event_FrontGearChange,
	"RearGearChange":        Event_RearGearChange,
	"RiderPositionChange":   Event_RiderPositionChange,
	"ElevHighAlert":         Event_ElevHighAlert,
	"ElevLowAlert":          Event_ElevLowAlert,
	"CommTimeout":           Event_CommTimeout,
	"AutoActivityDetect":    Event_AutoActivityDetect,
	"DiveAlert":             Event_DiveAlert,
	"DiveGasSwitched":       Event_DiveGasSwitched,
	"TankPressureReserve":   Event_TankPressureReserve,
	"TankPressureCritical":  Event_TankPressureCritical,
	"TankLost":              Event_TankLost,
	"RadarThreatAlert":      Event_RadarThreatAlert,
	"TankBatteryLow":        Event_TankBatteryLow,
	"TankPodConnected":      Event_TankPodConnected,
	"TankPodDisconnected":   Event_TankPodDisconnected,
}

func EventFromString(s string) Event {
	if v, ok := EventValues[s]; ok {
		return v
	}
	return Event(Event_Invalid)
}
