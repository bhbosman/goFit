package fitDecl

import strconv "strconv"

type DiveAlert byte

const (
	DiveAlert_NdlReached                DiveAlert = 0
	DiveAlert_GasSwitchPrompted         DiveAlert = 1
	DiveAlert_NearSurface               DiveAlert = 2
	DiveAlert_ApproachingNdl            DiveAlert = 3
	DiveAlert_Po2Warn                   DiveAlert = 4
	DiveAlert_Po2CritHigh               DiveAlert = 5
	DiveAlert_Po2CritLow                DiveAlert = 6
	DiveAlert_TimeAlert                 DiveAlert = 7
	DiveAlert_DepthAlert                DiveAlert = 8
	DiveAlert_DecoCeilingBroken         DiveAlert = 9
	DiveAlert_DecoComplete              DiveAlert = 10
	DiveAlert_SafetyStopBroken          DiveAlert = 11
	DiveAlert_SafetyStopComplete        DiveAlert = 12
	DiveAlert_CnsWarning                DiveAlert = 13
	DiveAlert_CnsCritical               DiveAlert = 14
	DiveAlert_OtuWarning                DiveAlert = 15
	DiveAlert_OtuCritical               DiveAlert = 16
	DiveAlert_AscentCritical            DiveAlert = 17
	DiveAlert_AlertDismissedByKey       DiveAlert = 18
	DiveAlert_AlertDismissedByTimeout   DiveAlert = 19
	DiveAlert_BatteryLow                DiveAlert = 20
	DiveAlert_BatteryCritical           DiveAlert = 21
	DiveAlert_SafetyStopStarted         DiveAlert = 22
	DiveAlert_ApproachingFirstDecoStop  DiveAlert = 23
	DiveAlert_SetpointSwitchAutoLow     DiveAlert = 24
	DiveAlert_SetpointSwitchAutoHigh    DiveAlert = 25
	DiveAlert_SetpointSwitchManualLow   DiveAlert = 26
	DiveAlert_SetpointSwitchManualHigh  DiveAlert = 27
	DiveAlert_AutoSetpointSwitchIgnored DiveAlert = 28
	DiveAlert_SwitchedToOpenCircuit     DiveAlert = 29
	DiveAlert_SwitchedToClosedCircuit   DiveAlert = 30
	DiveAlert_TankBatteryLow            DiveAlert = 32
	DiveAlert_Po2CcrDilLow              DiveAlert = 33
	DiveAlert_DecoStopCleared           DiveAlert = 34
	DiveAlert_ApneaNeutralBuoyancy      DiveAlert = 35
	DiveAlert_ApneaTargetDepth          DiveAlert = 36
	DiveAlert_ApneaSurface              DiveAlert = 37
	DiveAlert_ApneaHighSpeed            DiveAlert = 38
	DiveAlert_ApneaLowSpeed             DiveAlert = 39
	DiveAlert_Invalid                   DiveAlert = 255
)

var DiveAlertNames = map[DiveAlert]string{
	DiveAlert_NdlReached:                "NdlReached",
	DiveAlert_GasSwitchPrompted:         "GasSwitchPrompted",
	DiveAlert_NearSurface:               "NearSurface",
	DiveAlert_ApproachingNdl:            "ApproachingNdl",
	DiveAlert_Po2Warn:                   "Po2Warn",
	DiveAlert_Po2CritHigh:               "Po2CritHigh",
	DiveAlert_Po2CritLow:                "Po2CritLow",
	DiveAlert_TimeAlert:                 "TimeAlert",
	DiveAlert_DepthAlert:                "DepthAlert",
	DiveAlert_DecoCeilingBroken:         "DecoCeilingBroken",
	DiveAlert_DecoComplete:              "DecoComplete",
	DiveAlert_SafetyStopBroken:          "SafetyStopBroken",
	DiveAlert_SafetyStopComplete:        "SafetyStopComplete",
	DiveAlert_CnsWarning:                "CnsWarning",
	DiveAlert_CnsCritical:               "CnsCritical",
	DiveAlert_OtuWarning:                "OtuWarning",
	DiveAlert_OtuCritical:               "OtuCritical",
	DiveAlert_AscentCritical:            "AscentCritical",
	DiveAlert_AlertDismissedByKey:       "AlertDismissedByKey",
	DiveAlert_AlertDismissedByTimeout:   "AlertDismissedByTimeout",
	DiveAlert_BatteryLow:                "BatteryLow",
	DiveAlert_BatteryCritical:           "BatteryCritical",
	DiveAlert_SafetyStopStarted:         "SafetyStopStarted",
	DiveAlert_ApproachingFirstDecoStop:  "ApproachingFirstDecoStop",
	DiveAlert_SetpointSwitchAutoLow:     "SetpointSwitchAutoLow",
	DiveAlert_SetpointSwitchAutoHigh:    "SetpointSwitchAutoHigh",
	DiveAlert_SetpointSwitchManualLow:   "SetpointSwitchManualLow",
	DiveAlert_SetpointSwitchManualHigh:  "SetpointSwitchManualHigh",
	DiveAlert_AutoSetpointSwitchIgnored: "AutoSetpointSwitchIgnored",
	DiveAlert_SwitchedToOpenCircuit:     "SwitchedToOpenCircuit",
	DiveAlert_SwitchedToClosedCircuit:   "SwitchedToClosedCircuit",
	DiveAlert_TankBatteryLow:            "TankBatteryLow",
	DiveAlert_Po2CcrDilLow:              "Po2CcrDilLow",
	DiveAlert_DecoStopCleared:           "DecoStopCleared",
	DiveAlert_ApneaNeutralBuoyancy:      "ApneaNeutralBuoyancy",
	DiveAlert_ApneaTargetDepth:          "ApneaTargetDepth",
	DiveAlert_ApneaSurface:              "ApneaSurface",
	DiveAlert_ApneaHighSpeed:            "ApneaHighSpeed",
	DiveAlert_ApneaLowSpeed:             "ApneaLowSpeed",
}

func (k DiveAlert) String() string {
	if uint(k) < uint(len(DiveAlertNames)) {
		if v, ok := DiveAlertNames[k]; ok {
			return v
		}
	}
	return "DiveAlert(" + strconv.Itoa(int(k)) + ")"
}

var DiveAlertValues = map[string]DiveAlert{
	"NdlReached":                DiveAlert_NdlReached,
	"GasSwitchPrompted":         DiveAlert_GasSwitchPrompted,
	"NearSurface":               DiveAlert_NearSurface,
	"ApproachingNdl":            DiveAlert_ApproachingNdl,
	"Po2Warn":                   DiveAlert_Po2Warn,
	"Po2CritHigh":               DiveAlert_Po2CritHigh,
	"Po2CritLow":                DiveAlert_Po2CritLow,
	"TimeAlert":                 DiveAlert_TimeAlert,
	"DepthAlert":                DiveAlert_DepthAlert,
	"DecoCeilingBroken":         DiveAlert_DecoCeilingBroken,
	"DecoComplete":              DiveAlert_DecoComplete,
	"SafetyStopBroken":          DiveAlert_SafetyStopBroken,
	"SafetyStopComplete":        DiveAlert_SafetyStopComplete,
	"CnsWarning":                DiveAlert_CnsWarning,
	"CnsCritical":               DiveAlert_CnsCritical,
	"OtuWarning":                DiveAlert_OtuWarning,
	"OtuCritical":               DiveAlert_OtuCritical,
	"AscentCritical":            DiveAlert_AscentCritical,
	"AlertDismissedByKey":       DiveAlert_AlertDismissedByKey,
	"AlertDismissedByTimeout":   DiveAlert_AlertDismissedByTimeout,
	"BatteryLow":                DiveAlert_BatteryLow,
	"BatteryCritical":           DiveAlert_BatteryCritical,
	"SafetyStopStarted":         DiveAlert_SafetyStopStarted,
	"ApproachingFirstDecoStop":  DiveAlert_ApproachingFirstDecoStop,
	"SetpointSwitchAutoLow":     DiveAlert_SetpointSwitchAutoLow,
	"SetpointSwitchAutoHigh":    DiveAlert_SetpointSwitchAutoHigh,
	"SetpointSwitchManualLow":   DiveAlert_SetpointSwitchManualLow,
	"SetpointSwitchManualHigh":  DiveAlert_SetpointSwitchManualHigh,
	"AutoSetpointSwitchIgnored": DiveAlert_AutoSetpointSwitchIgnored,
	"SwitchedToOpenCircuit":     DiveAlert_SwitchedToOpenCircuit,
	"SwitchedToClosedCircuit":   DiveAlert_SwitchedToClosedCircuit,
	"TankBatteryLow":            DiveAlert_TankBatteryLow,
	"Po2CcrDilLow":              DiveAlert_Po2CcrDilLow,
	"DecoStopCleared":           DiveAlert_DecoStopCleared,
	"ApneaNeutralBuoyancy":      DiveAlert_ApneaNeutralBuoyancy,
	"ApneaTargetDepth":          DiveAlert_ApneaTargetDepth,
	"ApneaSurface":              DiveAlert_ApneaSurface,
	"ApneaHighSpeed":            DiveAlert_ApneaHighSpeed,
	"ApneaLowSpeed":             DiveAlert_ApneaLowSpeed,
}

func DiveAlertFromString(s string) DiveAlert {
	if v, ok := DiveAlertValues[s]; ok {
		return v
	}
	return DiveAlert(DiveAlert_Invalid)
}
