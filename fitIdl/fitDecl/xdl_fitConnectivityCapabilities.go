package fitDecl

import strconv "strconv"

type ConnectivityCapabilities uint32

const (
	ConnectivityCapabilities_Bluetooth                       ConnectivityCapabilities = 1
	ConnectivityCapabilities_BluetoothLe                     ConnectivityCapabilities = 2
	ConnectivityCapabilities_Ant                             ConnectivityCapabilities = 4
	ConnectivityCapabilities_ActivityUpload                  ConnectivityCapabilities = 8
	ConnectivityCapabilities_CourseDownload                  ConnectivityCapabilities = 16
	ConnectivityCapabilities_WorkoutDownload                 ConnectivityCapabilities = 32
	ConnectivityCapabilities_LiveTrack                       ConnectivityCapabilities = 64
	ConnectivityCapabilities_WeatherConditions               ConnectivityCapabilities = 128
	ConnectivityCapabilities_WeatherAlerts                   ConnectivityCapabilities = 256
	ConnectivityCapabilities_GpsEphemerisDownload            ConnectivityCapabilities = 512
	ConnectivityCapabilities_ExplicitArchive                 ConnectivityCapabilities = 1024
	ConnectivityCapabilities_SetupIncomplete                 ConnectivityCapabilities = 2048
	ConnectivityCapabilities_ContinueSyncAfterSoftwareUpdate ConnectivityCapabilities = 4096
	ConnectivityCapabilities_ConnectIqAppDownload            ConnectivityCapabilities = 8192
	ConnectivityCapabilities_GolfCourseDownload              ConnectivityCapabilities = 16384
	ConnectivityCapabilities_DeviceInitiatesSync             ConnectivityCapabilities = 32768
	ConnectivityCapabilities_ConnectIqWatchAppDownload       ConnectivityCapabilities = 65536
	ConnectivityCapabilities_ConnectIqWidgetDownload         ConnectivityCapabilities = 131072
	ConnectivityCapabilities_ConnectIqWatchFaceDownload      ConnectivityCapabilities = 262144
	ConnectivityCapabilities_ConnectIqDataFieldDownload      ConnectivityCapabilities = 524288
	ConnectivityCapabilities_ConnectIqAppManagment           ConnectivityCapabilities = 1048576
	ConnectivityCapabilities_SwingSensor                     ConnectivityCapabilities = 2097152
	ConnectivityCapabilities_SwingSensorRemote               ConnectivityCapabilities = 4194304
	ConnectivityCapabilities_IncidentDetection               ConnectivityCapabilities = 8388608
	ConnectivityCapabilities_AudioPrompts                    ConnectivityCapabilities = 16777216
	ConnectivityCapabilities_WifiVerification                ConnectivityCapabilities = 33554432
	ConnectivityCapabilities_TrueUp                          ConnectivityCapabilities = 67108864
	ConnectivityCapabilities_FindMyWatch                     ConnectivityCapabilities = 134217728
	ConnectivityCapabilities_RemoteManualSync                ConnectivityCapabilities = 268435456
	ConnectivityCapabilities_LiveTrackAutoStart              ConnectivityCapabilities = 536870912
	ConnectivityCapabilities_LiveTrackMessaging              ConnectivityCapabilities = 1073741824
	ConnectivityCapabilities_InstantInput                    ConnectivityCapabilities = 2147483648
	ConnectivityCapabilities_Invalid                         ConnectivityCapabilities = 0
)

var ConnectivityCapabilitiesNames = map[ConnectivityCapabilities]string{
	ConnectivityCapabilities_Bluetooth:                       "Bluetooth",
	ConnectivityCapabilities_BluetoothLe:                     "BluetoothLe",
	ConnectivityCapabilities_Ant:                             "Ant",
	ConnectivityCapabilities_ActivityUpload:                  "ActivityUpload",
	ConnectivityCapabilities_CourseDownload:                  "CourseDownload",
	ConnectivityCapabilities_WorkoutDownload:                 "WorkoutDownload",
	ConnectivityCapabilities_LiveTrack:                       "LiveTrack",
	ConnectivityCapabilities_WeatherConditions:               "WeatherConditions",
	ConnectivityCapabilities_WeatherAlerts:                   "WeatherAlerts",
	ConnectivityCapabilities_GpsEphemerisDownload:            "GpsEphemerisDownload",
	ConnectivityCapabilities_ExplicitArchive:                 "ExplicitArchive",
	ConnectivityCapabilities_SetupIncomplete:                 "SetupIncomplete",
	ConnectivityCapabilities_ContinueSyncAfterSoftwareUpdate: "ContinueSyncAfterSoftwareUpdate",
	ConnectivityCapabilities_ConnectIqAppDownload:            "ConnectIqAppDownload",
	ConnectivityCapabilities_GolfCourseDownload:              "GolfCourseDownload",
	ConnectivityCapabilities_DeviceInitiatesSync:             "DeviceInitiatesSync",
	ConnectivityCapabilities_ConnectIqWatchAppDownload:       "ConnectIqWatchAppDownload",
	ConnectivityCapabilities_ConnectIqWidgetDownload:         "ConnectIqWidgetDownload",
	ConnectivityCapabilities_ConnectIqWatchFaceDownload:      "ConnectIqWatchFaceDownload",
	ConnectivityCapabilities_ConnectIqDataFieldDownload:      "ConnectIqDataFieldDownload",
	ConnectivityCapabilities_ConnectIqAppManagment:           "ConnectIqAppManagment",
	ConnectivityCapabilities_SwingSensor:                     "SwingSensor",
	ConnectivityCapabilities_SwingSensorRemote:               "SwingSensorRemote",
	ConnectivityCapabilities_IncidentDetection:               "IncidentDetection",
	ConnectivityCapabilities_AudioPrompts:                    "AudioPrompts",
	ConnectivityCapabilities_WifiVerification:                "WifiVerification",
	ConnectivityCapabilities_TrueUp:                          "TrueUp",
	ConnectivityCapabilities_FindMyWatch:                     "FindMyWatch",
	ConnectivityCapabilities_RemoteManualSync:                "RemoteManualSync",
	ConnectivityCapabilities_LiveTrackAutoStart:              "LiveTrackAutoStart",
	ConnectivityCapabilities_LiveTrackMessaging:              "LiveTrackMessaging",
	ConnectivityCapabilities_InstantInput:                    "InstantInput",
}

func (k ConnectivityCapabilities) String() string {
	if uint(k) < uint(len(ConnectivityCapabilitiesNames)) {
		if v, ok := ConnectivityCapabilitiesNames[k]; ok {
			return v
		}
	}
	return "ConnectivityCapabilities(" + strconv.Itoa(int(k)) + ")"
}

var ConnectivityCapabilitiesValues = map[string]ConnectivityCapabilities{
	"Bluetooth":                       ConnectivityCapabilities_Bluetooth,
	"BluetoothLe":                     ConnectivityCapabilities_BluetoothLe,
	"Ant":                             ConnectivityCapabilities_Ant,
	"ActivityUpload":                  ConnectivityCapabilities_ActivityUpload,
	"CourseDownload":                  ConnectivityCapabilities_CourseDownload,
	"WorkoutDownload":                 ConnectivityCapabilities_WorkoutDownload,
	"LiveTrack":                       ConnectivityCapabilities_LiveTrack,
	"WeatherConditions":               ConnectivityCapabilities_WeatherConditions,
	"WeatherAlerts":                   ConnectivityCapabilities_WeatherAlerts,
	"GpsEphemerisDownload":            ConnectivityCapabilities_GpsEphemerisDownload,
	"ExplicitArchive":                 ConnectivityCapabilities_ExplicitArchive,
	"SetupIncomplete":                 ConnectivityCapabilities_SetupIncomplete,
	"ContinueSyncAfterSoftwareUpdate": ConnectivityCapabilities_ContinueSyncAfterSoftwareUpdate,
	"ConnectIqAppDownload":            ConnectivityCapabilities_ConnectIqAppDownload,
	"GolfCourseDownload":              ConnectivityCapabilities_GolfCourseDownload,
	"DeviceInitiatesSync":             ConnectivityCapabilities_DeviceInitiatesSync,
	"ConnectIqWatchAppDownload":       ConnectivityCapabilities_ConnectIqWatchAppDownload,
	"ConnectIqWidgetDownload":         ConnectivityCapabilities_ConnectIqWidgetDownload,
	"ConnectIqWatchFaceDownload":      ConnectivityCapabilities_ConnectIqWatchFaceDownload,
	"ConnectIqDataFieldDownload":      ConnectivityCapabilities_ConnectIqDataFieldDownload,
	"ConnectIqAppManagment":           ConnectivityCapabilities_ConnectIqAppManagment,
	"SwingSensor":                     ConnectivityCapabilities_SwingSensor,
	"SwingSensorRemote":               ConnectivityCapabilities_SwingSensorRemote,
	"IncidentDetection":               ConnectivityCapabilities_IncidentDetection,
	"AudioPrompts":                    ConnectivityCapabilities_AudioPrompts,
	"WifiVerification":                ConnectivityCapabilities_WifiVerification,
	"TrueUp":                          ConnectivityCapabilities_TrueUp,
	"FindMyWatch":                     ConnectivityCapabilities_FindMyWatch,
	"RemoteManualSync":                ConnectivityCapabilities_RemoteManualSync,
	"LiveTrackAutoStart":              ConnectivityCapabilities_LiveTrackAutoStart,
	"LiveTrackMessaging":              ConnectivityCapabilities_LiveTrackMessaging,
	"InstantInput":                    ConnectivityCapabilities_InstantInput,
}

func ConnectivityCapabilitiesFromString(s string) ConnectivityCapabilities {
	if v, ok := ConnectivityCapabilitiesValues[s]; ok {
		return v
	}
	return ConnectivityCapabilities(ConnectivityCapabilities_Invalid)
}
