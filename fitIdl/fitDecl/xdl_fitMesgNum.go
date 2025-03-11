package fitDecl

import strconv "strconv"

type MesgNum uint16

const (
	MesgNum_FileId                      MesgNum = 0
	MesgNum_Capabilities                MesgNum = 1
	MesgNum_DeviceSettings              MesgNum = 2
	MesgNum_UserProfile                 MesgNum = 3
	MesgNum_HrmProfile                  MesgNum = 4
	MesgNum_SdmProfile                  MesgNum = 5
	MesgNum_BikeProfile                 MesgNum = 6
	MesgNum_ZonesTarget                 MesgNum = 7
	MesgNum_HrZone                      MesgNum = 8
	MesgNum_PowerZone                   MesgNum = 9
	MesgNum_MetZone                     MesgNum = 10
	MesgNum_Sport                       MesgNum = 12
	MesgNum_Goal                        MesgNum = 15
	MesgNum_Session                     MesgNum = 18
	MesgNum_Lap                         MesgNum = 19
	MesgNum_Record                      MesgNum = 20
	MesgNum_Event                       MesgNum = 21
	MesgNum_DeviceInfo                  MesgNum = 23
	MesgNum_Workout                     MesgNum = 26
	MesgNum_WorkoutStep                 MesgNum = 27
	MesgNum_Schedule                    MesgNum = 28
	MesgNum_WeightScale                 MesgNum = 30
	MesgNum_Course                      MesgNum = 31
	MesgNum_CoursePoint                 MesgNum = 32
	MesgNum_Totals                      MesgNum = 33
	MesgNum_Activity                    MesgNum = 34
	MesgNum_Software                    MesgNum = 35
	MesgNum_FileCapabilities            MesgNum = 37
	MesgNum_MesgCapabilities            MesgNum = 38
	MesgNum_FieldCapabilities           MesgNum = 39
	MesgNum_FileCreator                 MesgNum = 49
	MesgNum_BloodPressure               MesgNum = 51
	MesgNum_SpeedZone                   MesgNum = 53
	MesgNum_Monitoring                  MesgNum = 55
	MesgNum_TrainingFile                MesgNum = 72
	MesgNum_Hrv                         MesgNum = 78
	MesgNum_AntRx                       MesgNum = 80
	MesgNum_AntTx                       MesgNum = 81
	MesgNum_AntChannelId                MesgNum = 82
	MesgNum_Length                      MesgNum = 101
	MesgNum_MonitoringInfo              MesgNum = 103
	MesgNum_Pad                         MesgNum = 105
	MesgNum_SlaveDevice                 MesgNum = 106
	MesgNum_Connectivity                MesgNum = 127
	MesgNum_WeatherConditions           MesgNum = 128
	MesgNum_WeatherAlert                MesgNum = 129
	MesgNum_CadenceZone                 MesgNum = 131
	MesgNum_Hr                          MesgNum = 132
	MesgNum_SegmentLap                  MesgNum = 142
	MesgNum_MemoGlob                    MesgNum = 145
	MesgNum_SegmentId                   MesgNum = 148
	MesgNum_SegmentLeaderboardEntry     MesgNum = 149
	MesgNum_SegmentPoint                MesgNum = 150
	MesgNum_SegmentFile                 MesgNum = 151
	MesgNum_WorkoutSession              MesgNum = 158
	MesgNum_WatchfaceSettings           MesgNum = 159
	MesgNum_GpsMetadata                 MesgNum = 160
	MesgNum_CameraEvent                 MesgNum = 161
	MesgNum_TimestampCorrelation        MesgNum = 162
	MesgNum_GyroscopeData               MesgNum = 164
	MesgNum_AccelerometerData           MesgNum = 165
	MesgNum_ThreeDSensorCalibration     MesgNum = 167
	MesgNum_VideoFrame                  MesgNum = 169
	MesgNum_ObdiiData                   MesgNum = 174
	MesgNum_NmeaSentence                MesgNum = 177
	MesgNum_AviationAttitude            MesgNum = 178
	MesgNum_Video                       MesgNum = 184
	MesgNum_VideoTitle                  MesgNum = 185
	MesgNum_VideoDescription            MesgNum = 186
	MesgNum_VideoClip                   MesgNum = 187
	MesgNum_OhrSettings                 MesgNum = 188
	MesgNum_ExdScreenConfiguration      MesgNum = 200
	MesgNum_ExdDataFieldConfiguration   MesgNum = 201
	MesgNum_ExdDataConceptConfiguration MesgNum = 202
	MesgNum_FieldDescription            MesgNum = 206
	MesgNum_DeveloperDataId             MesgNum = 207
	MesgNum_MagnetometerData            MesgNum = 208
	MesgNum_BarometerData               MesgNum = 209
	MesgNum_OneDSensorCalibration       MesgNum = 210
	MesgNum_MonitoringHrData            MesgNum = 211
	MesgNum_TimeInZone                  MesgNum = 216
	MesgNum_Set                         MesgNum = 225
	MesgNum_StressLevel                 MesgNum = 227
	MesgNum_MaxMetData                  MesgNum = 229
	MesgNum_DiveSettings                MesgNum = 258
	MesgNum_DiveGas                     MesgNum = 259
	MesgNum_DiveAlarm                   MesgNum = 262
	MesgNum_ExerciseTitle               MesgNum = 264
	MesgNum_DiveSummary                 MesgNum = 268
	MesgNum_Spo2Data                    MesgNum = 269
	MesgNum_SleepLevel                  MesgNum = 275
	MesgNum_Jump                        MesgNum = 285
	MesgNum_AadAccelFeatures            MesgNum = 289
	MesgNum_BeatIntervals               MesgNum = 290
	MesgNum_RespirationRate             MesgNum = 297
	MesgNum_HsaAccelerometerData        MesgNum = 302
	MesgNum_HsaStepData                 MesgNum = 304
	MesgNum_HsaSpo2Data                 MesgNum = 305
	MesgNum_HsaStressData               MesgNum = 306
	MesgNum_HsaRespirationData          MesgNum = 307
	MesgNum_HsaHeartRateData            MesgNum = 308
	MesgNum_Split                       MesgNum = 312
	MesgNum_SplitSummary                MesgNum = 313
	MesgNum_HsaBodyBatteryData          MesgNum = 314
	MesgNum_HsaEvent                    MesgNum = 315
	MesgNum_ClimbPro                    MesgNum = 317
	MesgNum_TankUpdate                  MesgNum = 319
	MesgNum_TankSummary                 MesgNum = 323
	MesgNum_SleepAssessment             MesgNum = 346
	MesgNum_HrvStatusSummary            MesgNum = 370
	MesgNum_HrvValue                    MesgNum = 371
	MesgNum_RawBbi                      MesgNum = 372
	MesgNum_DeviceAuxBatteryInfo        MesgNum = 375
	MesgNum_HsaGyroscopeData            MesgNum = 376
	MesgNum_ChronoShotSession           MesgNum = 387
	MesgNum_ChronoShotData              MesgNum = 388
	MesgNum_HsaConfigurationData        MesgNum = 389
	MesgNum_DiveApneaAlarm              MesgNum = 393
	MesgNum_SkinTempOvernight           MesgNum = 398
	MesgNum_HsaWristTemperatureData     MesgNum = 409
	MesgNum_MfgRangeMin                 MesgNum = 65280
	MesgNum_MfgRangeMax                 MesgNum = 65534
	MesgNum_Invalid                     MesgNum = 65535
)

var MesgNumNames = map[MesgNum]string{
	MesgNum_FileId:                      "FileId",
	MesgNum_Capabilities:                "Capabilities",
	MesgNum_DeviceSettings:              "DeviceSettings",
	MesgNum_UserProfile:                 "UserProfile",
	MesgNum_HrmProfile:                  "HrmProfile",
	MesgNum_SdmProfile:                  "SdmProfile",
	MesgNum_BikeProfile:                 "BikeProfile",
	MesgNum_ZonesTarget:                 "ZonesTarget",
	MesgNum_HrZone:                      "HrZone",
	MesgNum_PowerZone:                   "PowerZone",
	MesgNum_MetZone:                     "MetZone",
	MesgNum_Sport:                       "Sport",
	MesgNum_Goal:                        "Goal",
	MesgNum_Session:                     "Session",
	MesgNum_Lap:                         "Lap",
	MesgNum_Record:                      "Record",
	MesgNum_Event:                       "Event",
	MesgNum_DeviceInfo:                  "DeviceInfo",
	MesgNum_Workout:                     "Workout",
	MesgNum_WorkoutStep:                 "WorkoutStep",
	MesgNum_Schedule:                    "Schedule",
	MesgNum_WeightScale:                 "WeightScale",
	MesgNum_Course:                      "Course",
	MesgNum_CoursePoint:                 "CoursePoint",
	MesgNum_Totals:                      "Totals",
	MesgNum_Activity:                    "Activity",
	MesgNum_Software:                    "Software",
	MesgNum_FileCapabilities:            "FileCapabilities",
	MesgNum_MesgCapabilities:            "MesgCapabilities",
	MesgNum_FieldCapabilities:           "FieldCapabilities",
	MesgNum_FileCreator:                 "FileCreator",
	MesgNum_BloodPressure:               "BloodPressure",
	MesgNum_SpeedZone:                   "SpeedZone",
	MesgNum_Monitoring:                  "Monitoring",
	MesgNum_TrainingFile:                "TrainingFile",
	MesgNum_Hrv:                         "Hrv",
	MesgNum_AntRx:                       "AntRx",
	MesgNum_AntTx:                       "AntTx",
	MesgNum_AntChannelId:                "AntChannelId",
	MesgNum_Length:                      "Length",
	MesgNum_MonitoringInfo:              "MonitoringInfo",
	MesgNum_Pad:                         "Pad",
	MesgNum_SlaveDevice:                 "SlaveDevice",
	MesgNum_Connectivity:                "Connectivity",
	MesgNum_WeatherConditions:           "WeatherConditions",
	MesgNum_WeatherAlert:                "WeatherAlert",
	MesgNum_CadenceZone:                 "CadenceZone",
	MesgNum_Hr:                          "Hr",
	MesgNum_SegmentLap:                  "SegmentLap",
	MesgNum_MemoGlob:                    "MemoGlob",
	MesgNum_SegmentId:                   "SegmentId",
	MesgNum_SegmentLeaderboardEntry:     "SegmentLeaderboardEntry",
	MesgNum_SegmentPoint:                "SegmentPoint",
	MesgNum_SegmentFile:                 "SegmentFile",
	MesgNum_WorkoutSession:              "WorkoutSession",
	MesgNum_WatchfaceSettings:           "WatchfaceSettings",
	MesgNum_GpsMetadata:                 "GpsMetadata",
	MesgNum_CameraEvent:                 "CameraEvent",
	MesgNum_TimestampCorrelation:        "TimestampCorrelation",
	MesgNum_GyroscopeData:               "GyroscopeData",
	MesgNum_AccelerometerData:           "AccelerometerData",
	MesgNum_ThreeDSensorCalibration:     "ThreeDSensorCalibration",
	MesgNum_VideoFrame:                  "VideoFrame",
	MesgNum_ObdiiData:                   "ObdiiData",
	MesgNum_NmeaSentence:                "NmeaSentence",
	MesgNum_AviationAttitude:            "AviationAttitude",
	MesgNum_Video:                       "Video",
	MesgNum_VideoTitle:                  "VideoTitle",
	MesgNum_VideoDescription:            "VideoDescription",
	MesgNum_VideoClip:                   "VideoClip",
	MesgNum_OhrSettings:                 "OhrSettings",
	MesgNum_ExdScreenConfiguration:      "ExdScreenConfiguration",
	MesgNum_ExdDataFieldConfiguration:   "ExdDataFieldConfiguration",
	MesgNum_ExdDataConceptConfiguration: "ExdDataConceptConfiguration",
	MesgNum_FieldDescription:            "FieldDescription",
	MesgNum_DeveloperDataId:             "DeveloperDataId",
	MesgNum_MagnetometerData:            "MagnetometerData",
	MesgNum_BarometerData:               "BarometerData",
	MesgNum_OneDSensorCalibration:       "OneDSensorCalibration",
	MesgNum_MonitoringHrData:            "MonitoringHrData",
	MesgNum_TimeInZone:                  "TimeInZone",
	MesgNum_Set:                         "Set",
	MesgNum_StressLevel:                 "StressLevel",
	MesgNum_MaxMetData:                  "MaxMetData",
	MesgNum_DiveSettings:                "DiveSettings",
	MesgNum_DiveGas:                     "DiveGas",
	MesgNum_DiveAlarm:                   "DiveAlarm",
	MesgNum_ExerciseTitle:               "ExerciseTitle",
	MesgNum_DiveSummary:                 "DiveSummary",
	MesgNum_Spo2Data:                    "Spo2Data",
	MesgNum_SleepLevel:                  "SleepLevel",
	MesgNum_Jump:                        "Jump",
	MesgNum_AadAccelFeatures:            "AadAccelFeatures",
	MesgNum_BeatIntervals:               "BeatIntervals",
	MesgNum_RespirationRate:             "RespirationRate",
	MesgNum_HsaAccelerometerData:        "HsaAccelerometerData",
	MesgNum_HsaStepData:                 "HsaStepData",
	MesgNum_HsaSpo2Data:                 "HsaSpo2Data",
	MesgNum_HsaStressData:               "HsaStressData",
	MesgNum_HsaRespirationData:          "HsaRespirationData",
	MesgNum_HsaHeartRateData:            "HsaHeartRateData",
	MesgNum_Split:                       "Split",
	MesgNum_SplitSummary:                "SplitSummary",
	MesgNum_HsaBodyBatteryData:          "HsaBodyBatteryData",
	MesgNum_HsaEvent:                    "HsaEvent",
	MesgNum_ClimbPro:                    "ClimbPro",
	MesgNum_TankUpdate:                  "TankUpdate",
	MesgNum_TankSummary:                 "TankSummary",
	MesgNum_SleepAssessment:             "SleepAssessment",
	MesgNum_HrvStatusSummary:            "HrvStatusSummary",
	MesgNum_HrvValue:                    "HrvValue",
	MesgNum_RawBbi:                      "RawBbi",
	MesgNum_DeviceAuxBatteryInfo:        "DeviceAuxBatteryInfo",
	MesgNum_HsaGyroscopeData:            "HsaGyroscopeData",
	MesgNum_ChronoShotSession:           "ChronoShotSession",
	MesgNum_ChronoShotData:              "ChronoShotData",
	MesgNum_HsaConfigurationData:        "HsaConfigurationData",
	MesgNum_DiveApneaAlarm:              "DiveApneaAlarm",
	MesgNum_SkinTempOvernight:           "SkinTempOvernight",
	MesgNum_HsaWristTemperatureData:     "HsaWristTemperatureData",
	MesgNum_MfgRangeMin:                 "MfgRangeMin",
	MesgNum_MfgRangeMax:                 "MfgRangeMax",
}

func (k MesgNum) String() string {
	if uint(k) < uint(len(MesgNumNames)) {
		if v, ok := MesgNumNames[k]; ok {
			return v
		}
	}
	return "MesgNum(" + strconv.Itoa(int(k)) + ")"
}

var MesgNumValues = map[string]MesgNum{
	"FileId":                      MesgNum_FileId,
	"Capabilities":                MesgNum_Capabilities,
	"DeviceSettings":              MesgNum_DeviceSettings,
	"UserProfile":                 MesgNum_UserProfile,
	"HrmProfile":                  MesgNum_HrmProfile,
	"SdmProfile":                  MesgNum_SdmProfile,
	"BikeProfile":                 MesgNum_BikeProfile,
	"ZonesTarget":                 MesgNum_ZonesTarget,
	"HrZone":                      MesgNum_HrZone,
	"PowerZone":                   MesgNum_PowerZone,
	"MetZone":                     MesgNum_MetZone,
	"Sport":                       MesgNum_Sport,
	"Goal":                        MesgNum_Goal,
	"Session":                     MesgNum_Session,
	"Lap":                         MesgNum_Lap,
	"Record":                      MesgNum_Record,
	"Event":                       MesgNum_Event,
	"DeviceInfo":                  MesgNum_DeviceInfo,
	"Workout":                     MesgNum_Workout,
	"WorkoutStep":                 MesgNum_WorkoutStep,
	"Schedule":                    MesgNum_Schedule,
	"WeightScale":                 MesgNum_WeightScale,
	"Course":                      MesgNum_Course,
	"CoursePoint":                 MesgNum_CoursePoint,
	"Totals":                      MesgNum_Totals,
	"Activity":                    MesgNum_Activity,
	"Software":                    MesgNum_Software,
	"FileCapabilities":            MesgNum_FileCapabilities,
	"MesgCapabilities":            MesgNum_MesgCapabilities,
	"FieldCapabilities":           MesgNum_FieldCapabilities,
	"FileCreator":                 MesgNum_FileCreator,
	"BloodPressure":               MesgNum_BloodPressure,
	"SpeedZone":                   MesgNum_SpeedZone,
	"Monitoring":                  MesgNum_Monitoring,
	"TrainingFile":                MesgNum_TrainingFile,
	"Hrv":                         MesgNum_Hrv,
	"AntRx":                       MesgNum_AntRx,
	"AntTx":                       MesgNum_AntTx,
	"AntChannelId":                MesgNum_AntChannelId,
	"Length":                      MesgNum_Length,
	"MonitoringInfo":              MesgNum_MonitoringInfo,
	"Pad":                         MesgNum_Pad,
	"SlaveDevice":                 MesgNum_SlaveDevice,
	"Connectivity":                MesgNum_Connectivity,
	"WeatherConditions":           MesgNum_WeatherConditions,
	"WeatherAlert":                MesgNum_WeatherAlert,
	"CadenceZone":                 MesgNum_CadenceZone,
	"Hr":                          MesgNum_Hr,
	"SegmentLap":                  MesgNum_SegmentLap,
	"MemoGlob":                    MesgNum_MemoGlob,
	"SegmentId":                   MesgNum_SegmentId,
	"SegmentLeaderboardEntry":     MesgNum_SegmentLeaderboardEntry,
	"SegmentPoint":                MesgNum_SegmentPoint,
	"SegmentFile":                 MesgNum_SegmentFile,
	"WorkoutSession":              MesgNum_WorkoutSession,
	"WatchfaceSettings":           MesgNum_WatchfaceSettings,
	"GpsMetadata":                 MesgNum_GpsMetadata,
	"CameraEvent":                 MesgNum_CameraEvent,
	"TimestampCorrelation":        MesgNum_TimestampCorrelation,
	"GyroscopeData":               MesgNum_GyroscopeData,
	"AccelerometerData":           MesgNum_AccelerometerData,
	"ThreeDSensorCalibration":     MesgNum_ThreeDSensorCalibration,
	"VideoFrame":                  MesgNum_VideoFrame,
	"ObdiiData":                   MesgNum_ObdiiData,
	"NmeaSentence":                MesgNum_NmeaSentence,
	"AviationAttitude":            MesgNum_AviationAttitude,
	"Video":                       MesgNum_Video,
	"VideoTitle":                  MesgNum_VideoTitle,
	"VideoDescription":            MesgNum_VideoDescription,
	"VideoClip":                   MesgNum_VideoClip,
	"OhrSettings":                 MesgNum_OhrSettings,
	"ExdScreenConfiguration":      MesgNum_ExdScreenConfiguration,
	"ExdDataFieldConfiguration":   MesgNum_ExdDataFieldConfiguration,
	"ExdDataConceptConfiguration": MesgNum_ExdDataConceptConfiguration,
	"FieldDescription":            MesgNum_FieldDescription,
	"DeveloperDataId":             MesgNum_DeveloperDataId,
	"MagnetometerData":            MesgNum_MagnetometerData,
	"BarometerData":               MesgNum_BarometerData,
	"OneDSensorCalibration":       MesgNum_OneDSensorCalibration,
	"MonitoringHrData":            MesgNum_MonitoringHrData,
	"TimeInZone":                  MesgNum_TimeInZone,
	"Set":                         MesgNum_Set,
	"StressLevel":                 MesgNum_StressLevel,
	"MaxMetData":                  MesgNum_MaxMetData,
	"DiveSettings":                MesgNum_DiveSettings,
	"DiveGas":                     MesgNum_DiveGas,
	"DiveAlarm":                   MesgNum_DiveAlarm,
	"ExerciseTitle":               MesgNum_ExerciseTitle,
	"DiveSummary":                 MesgNum_DiveSummary,
	"Spo2Data":                    MesgNum_Spo2Data,
	"SleepLevel":                  MesgNum_SleepLevel,
	"Jump":                        MesgNum_Jump,
	"AadAccelFeatures":            MesgNum_AadAccelFeatures,
	"BeatIntervals":               MesgNum_BeatIntervals,
	"RespirationRate":             MesgNum_RespirationRate,
	"HsaAccelerometerData":        MesgNum_HsaAccelerometerData,
	"HsaStepData":                 MesgNum_HsaStepData,
	"HsaSpo2Data":                 MesgNum_HsaSpo2Data,
	"HsaStressData":               MesgNum_HsaStressData,
	"HsaRespirationData":          MesgNum_HsaRespirationData,
	"HsaHeartRateData":            MesgNum_HsaHeartRateData,
	"Split":                       MesgNum_Split,
	"SplitSummary":                MesgNum_SplitSummary,
	"HsaBodyBatteryData":          MesgNum_HsaBodyBatteryData,
	"HsaEvent":                    MesgNum_HsaEvent,
	"ClimbPro":                    MesgNum_ClimbPro,
	"TankUpdate":                  MesgNum_TankUpdate,
	"TankSummary":                 MesgNum_TankSummary,
	"SleepAssessment":             MesgNum_SleepAssessment,
	"HrvStatusSummary":            MesgNum_HrvStatusSummary,
	"HrvValue":                    MesgNum_HrvValue,
	"RawBbi":                      MesgNum_RawBbi,
	"DeviceAuxBatteryInfo":        MesgNum_DeviceAuxBatteryInfo,
	"HsaGyroscopeData":            MesgNum_HsaGyroscopeData,
	"ChronoShotSession":           MesgNum_ChronoShotSession,
	"ChronoShotData":              MesgNum_ChronoShotData,
	"HsaConfigurationData":        MesgNum_HsaConfigurationData,
	"DiveApneaAlarm":              MesgNum_DiveApneaAlarm,
	"SkinTempOvernight":           MesgNum_SkinTempOvernight,
	"HsaWristTemperatureData":     MesgNum_HsaWristTemperatureData,
	"MfgRangeMin":                 MesgNum_MfgRangeMin,
	"MfgRangeMax":                 MesgNum_MfgRangeMax,
}

func MesgNumFromString(s string) MesgNum {
	if v, ok := MesgNumValues[s]; ok {
		return v
	}
	return MesgNum(MesgNum_Invalid)
}
