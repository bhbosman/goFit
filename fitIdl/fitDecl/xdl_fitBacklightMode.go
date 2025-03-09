package fitDecl

import strconv "strconv"

type BacklightMode byte

const (
	BacklightMode_Off                                 BacklightMode = 0
	BacklightMode_Manual                              BacklightMode = 1
	BacklightMode_KeyAndMessages                      BacklightMode = 2
	BacklightMode_AutoBrightness                      BacklightMode = 3
	BacklightMode_SmartNotifications                  BacklightMode = 4
	BacklightMode_KeyAndMessagesNight                 BacklightMode = 5
	BacklightMode_KeyAndMessagesAndSmartNotifications BacklightMode = 6
	BacklightMode_Invalid                             BacklightMode = 255
)

var BacklightModeNames = map[BacklightMode]string{
	BacklightMode_Off:                                 "Off",
	BacklightMode_Manual:                              "Manual",
	BacklightMode_KeyAndMessages:                      "KeyAndMessages",
	BacklightMode_AutoBrightness:                      "AutoBrightness",
	BacklightMode_SmartNotifications:                  "SmartNotifications",
	BacklightMode_KeyAndMessagesNight:                 "KeyAndMessagesNight",
	BacklightMode_KeyAndMessagesAndSmartNotifications: "KeyAndMessagesAndSmartNotifications",
}

func (k BacklightMode) String() string {
	if uint(k) < uint(len(BacklightModeNames)) {
		if v, ok := BacklightModeNames[k]; ok {
			return v
		}
	}
	return "BacklightMode(" + strconv.Itoa(int(k)) + ")"
}

var BacklightModeValues = map[string]BacklightMode{
	"Off":                                 BacklightMode_Off,
	"Manual":                              BacklightMode_Manual,
	"KeyAndMessages":                      BacklightMode_KeyAndMessages,
	"AutoBrightness":                      BacklightMode_AutoBrightness,
	"SmartNotifications":                  BacklightMode_SmartNotifications,
	"KeyAndMessagesNight":                 BacklightMode_KeyAndMessagesNight,
	"KeyAndMessagesAndSmartNotifications": BacklightMode_KeyAndMessagesAndSmartNotifications,
}

func BacklightModeFromString(s string) BacklightMode {
	if v, ok := BacklightModeValues[s]; ok {
		return v
	}
	return BacklightMode(BacklightMode_Invalid)
}
