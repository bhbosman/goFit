package fitDecl

import strconv "strconv"

type File byte

const (
	File_Device           File = 1
	File_Settings         File = 2
	File_Sport            File = 3
	File_Activity         File = 4
	File_Workout          File = 5
	File_Course           File = 6
	File_Schedules        File = 7
	File_Weight           File = 9
	File_Totals           File = 10
	File_Goals            File = 11
	File_BloodPressure    File = 14
	File_MonitoringA      File = 15
	File_ActivitySummary  File = 20
	File_MonitoringDaily  File = 28
	File_MonitoringB      File = 32
	File_Segment          File = 34
	File_SegmentList      File = 35
	File_ExdConfiguration File = 40
	File_MfgRangeMin      File = 247
	File_MfgRangeMax      File = 254
	File_Invalid          File = 255
)

var FileNames = map[File]string{
	File_Device:           "Device",
	File_Settings:         "Settings",
	File_Sport:            "Sport",
	File_Activity:         "Activity",
	File_Workout:          "Workout",
	File_Course:           "Course",
	File_Schedules:        "Schedules",
	File_Weight:           "Weight",
	File_Totals:           "Totals",
	File_Goals:            "Goals",
	File_BloodPressure:    "BloodPressure",
	File_MonitoringA:      "MonitoringA",
	File_ActivitySummary:  "ActivitySummary",
	File_MonitoringDaily:  "MonitoringDaily",
	File_MonitoringB:      "MonitoringB",
	File_Segment:          "Segment",
	File_SegmentList:      "SegmentList",
	File_ExdConfiguration: "ExdConfiguration",
	File_MfgRangeMin:      "MfgRangeMin",
	File_MfgRangeMax:      "MfgRangeMax",
}

func (k File) String() string {
	if uint(k) < uint(len(FileNames)) {
		if v, ok := FileNames[k]; ok {
			return v
		}
	}
	return "File(" + strconv.Itoa(int(k)) + ")"
}

var FileValues = map[string]File{
	"Device":           File_Device,
	"Settings":         File_Settings,
	"Sport":            File_Sport,
	"Activity":         File_Activity,
	"Workout":          File_Workout,
	"Course":           File_Course,
	"Schedules":        File_Schedules,
	"Weight":           File_Weight,
	"Totals":           File_Totals,
	"Goals":            File_Goals,
	"BloodPressure":    File_BloodPressure,
	"MonitoringA":      File_MonitoringA,
	"ActivitySummary":  File_ActivitySummary,
	"MonitoringDaily":  File_MonitoringDaily,
	"MonitoringB":      File_MonitoringB,
	"Segment":          File_Segment,
	"SegmentList":      File_SegmentList,
	"ExdConfiguration": File_ExdConfiguration,
	"MfgRangeMin":      File_MfgRangeMin,
	"MfgRangeMax":      File_MfgRangeMax,
}

func FileFromString(s string) File {
	if v, ok := FileValues[s]; ok {
		return v
	}
	return File(File_Invalid)
}
