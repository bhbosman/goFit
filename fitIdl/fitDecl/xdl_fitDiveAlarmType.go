package fitDecl

import strconv "strconv"

type DiveAlarmType byte

const (
	DiveAlarmType_Depth   DiveAlarmType = 0
	DiveAlarmType_Time    DiveAlarmType = 1
	DiveAlarmType_Speed   DiveAlarmType = 2
	DiveAlarmType_Invalid DiveAlarmType = 255
)

var DiveAlarmTypeNames = map[DiveAlarmType]string{
	DiveAlarmType_Depth: "Depth",
	DiveAlarmType_Time:  "Time",
	DiveAlarmType_Speed: "Speed",
}

func (k DiveAlarmType) String() string {
	if uint(k) < uint(len(DiveAlarmTypeNames)) {
		if v, ok := DiveAlarmTypeNames[k]; ok {
			return v
		}
	}
	return "DiveAlarmType(" + strconv.Itoa(int(k)) + ")"
}

var DiveAlarmTypeValues = map[string]DiveAlarmType{
	"Depth": DiveAlarmType_Depth,
	"Time":  DiveAlarmType_Time,
	"Speed": DiveAlarmType_Speed,
}

func DiveAlarmTypeFromString(s string) DiveAlarmType {
	if v, ok := DiveAlarmTypeValues[s]; ok {
		return v
	}
	return DiveAlarmType(DiveAlarmType_Invalid)
}
