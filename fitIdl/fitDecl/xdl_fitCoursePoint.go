package fitDecl

import strconv "strconv"

type CoursePoint byte

const (
	CoursePoint_Generic         CoursePoint = 0
	CoursePoint_Summit          CoursePoint = 1
	CoursePoint_Valley          CoursePoint = 2
	CoursePoint_Water           CoursePoint = 3
	CoursePoint_Food            CoursePoint = 4
	CoursePoint_Danger          CoursePoint = 5
	CoursePoint_Left            CoursePoint = 6
	CoursePoint_Right           CoursePoint = 7
	CoursePoint_Straight        CoursePoint = 8
	CoursePoint_FirstAid        CoursePoint = 9
	CoursePoint_FourthCategory  CoursePoint = 10
	CoursePoint_ThirdCategory   CoursePoint = 11
	CoursePoint_SecondCategory  CoursePoint = 12
	CoursePoint_FirstCategory   CoursePoint = 13
	CoursePoint_HorsCategory    CoursePoint = 14
	CoursePoint_Sprint          CoursePoint = 15
	CoursePoint_LeftFork        CoursePoint = 16
	CoursePoint_RightFork       CoursePoint = 17
	CoursePoint_MiddleFork      CoursePoint = 18
	CoursePoint_SlightLeft      CoursePoint = 19
	CoursePoint_SharpLeft       CoursePoint = 20
	CoursePoint_SlightRight     CoursePoint = 21
	CoursePoint_SharpRight      CoursePoint = 22
	CoursePoint_UTurn           CoursePoint = 23
	CoursePoint_SegmentStart    CoursePoint = 24
	CoursePoint_SegmentEnd      CoursePoint = 25
	CoursePoint_Campsite        CoursePoint = 27
	CoursePoint_AidStation      CoursePoint = 28
	CoursePoint_RestArea        CoursePoint = 29
	CoursePoint_GeneralDistance CoursePoint = 30
	CoursePoint_Service         CoursePoint = 31
	CoursePoint_EnergyGel       CoursePoint = 32
	CoursePoint_SportsDrink     CoursePoint = 33
	CoursePoint_MileMarker      CoursePoint = 34
	CoursePoint_Checkpoint      CoursePoint = 35
	CoursePoint_Shelter         CoursePoint = 36
	CoursePoint_MeetingSpot     CoursePoint = 37
	CoursePoint_Overlook        CoursePoint = 38
	CoursePoint_Toilet          CoursePoint = 39
	CoursePoint_Shower          CoursePoint = 40
	CoursePoint_Gear            CoursePoint = 41
	CoursePoint_SharpCurve      CoursePoint = 42
	CoursePoint_SteepIncline    CoursePoint = 43
	CoursePoint_Tunnel          CoursePoint = 44
	CoursePoint_Bridge          CoursePoint = 45
	CoursePoint_Obstacle        CoursePoint = 46
	CoursePoint_Crossing        CoursePoint = 47
	CoursePoint_Store           CoursePoint = 48
	CoursePoint_Transition      CoursePoint = 49
	CoursePoint_Navaid          CoursePoint = 50
	CoursePoint_Transport       CoursePoint = 51
	CoursePoint_Alert           CoursePoint = 52
	CoursePoint_Info            CoursePoint = 53
	CoursePoint_Invalid         CoursePoint = 255
)

var CoursePointNames = map[CoursePoint]string{
	CoursePoint_Generic:         "Generic",
	CoursePoint_Summit:          "Summit",
	CoursePoint_Valley:          "Valley",
	CoursePoint_Water:           "Water",
	CoursePoint_Food:            "Food",
	CoursePoint_Danger:          "Danger",
	CoursePoint_Left:            "Left",
	CoursePoint_Right:           "Right",
	CoursePoint_Straight:        "Straight",
	CoursePoint_FirstAid:        "FirstAid",
	CoursePoint_FourthCategory:  "FourthCategory",
	CoursePoint_ThirdCategory:   "ThirdCategory",
	CoursePoint_SecondCategory:  "SecondCategory",
	CoursePoint_FirstCategory:   "FirstCategory",
	CoursePoint_HorsCategory:    "HorsCategory",
	CoursePoint_Sprint:          "Sprint",
	CoursePoint_LeftFork:        "LeftFork",
	CoursePoint_RightFork:       "RightFork",
	CoursePoint_MiddleFork:      "MiddleFork",
	CoursePoint_SlightLeft:      "SlightLeft",
	CoursePoint_SharpLeft:       "SharpLeft",
	CoursePoint_SlightRight:     "SlightRight",
	CoursePoint_SharpRight:      "SharpRight",
	CoursePoint_UTurn:           "UTurn",
	CoursePoint_SegmentStart:    "SegmentStart",
	CoursePoint_SegmentEnd:      "SegmentEnd",
	CoursePoint_Campsite:        "Campsite",
	CoursePoint_AidStation:      "AidStation",
	CoursePoint_RestArea:        "RestArea",
	CoursePoint_GeneralDistance: "GeneralDistance",
	CoursePoint_Service:         "Service",
	CoursePoint_EnergyGel:       "EnergyGel",
	CoursePoint_SportsDrink:     "SportsDrink",
	CoursePoint_MileMarker:      "MileMarker",
	CoursePoint_Checkpoint:      "Checkpoint",
	CoursePoint_Shelter:         "Shelter",
	CoursePoint_MeetingSpot:     "MeetingSpot",
	CoursePoint_Overlook:        "Overlook",
	CoursePoint_Toilet:          "Toilet",
	CoursePoint_Shower:          "Shower",
	CoursePoint_Gear:            "Gear",
	CoursePoint_SharpCurve:      "SharpCurve",
	CoursePoint_SteepIncline:    "SteepIncline",
	CoursePoint_Tunnel:          "Tunnel",
	CoursePoint_Bridge:          "Bridge",
	CoursePoint_Obstacle:        "Obstacle",
	CoursePoint_Crossing:        "Crossing",
	CoursePoint_Store:           "Store",
	CoursePoint_Transition:      "Transition",
	CoursePoint_Navaid:          "Navaid",
	CoursePoint_Transport:       "Transport",
	CoursePoint_Alert:           "Alert",
	CoursePoint_Info:            "Info",
}

func (k CoursePoint) String() string {
	if uint(k) < uint(len(CoursePointNames)) {
		if v, ok := CoursePointNames[k]; ok {
			return v
		}
	}
	return "CoursePoint(" + strconv.Itoa(int(k)) + ")"
}

var CoursePointValues = map[string]CoursePoint{
	"Generic":         CoursePoint_Generic,
	"Summit":          CoursePoint_Summit,
	"Valley":          CoursePoint_Valley,
	"Water":           CoursePoint_Water,
	"Food":            CoursePoint_Food,
	"Danger":          CoursePoint_Danger,
	"Left":            CoursePoint_Left,
	"Right":           CoursePoint_Right,
	"Straight":        CoursePoint_Straight,
	"FirstAid":        CoursePoint_FirstAid,
	"FourthCategory":  CoursePoint_FourthCategory,
	"ThirdCategory":   CoursePoint_ThirdCategory,
	"SecondCategory":  CoursePoint_SecondCategory,
	"FirstCategory":   CoursePoint_FirstCategory,
	"HorsCategory":    CoursePoint_HorsCategory,
	"Sprint":          CoursePoint_Sprint,
	"LeftFork":        CoursePoint_LeftFork,
	"RightFork":       CoursePoint_RightFork,
	"MiddleFork":      CoursePoint_MiddleFork,
	"SlightLeft":      CoursePoint_SlightLeft,
	"SharpLeft":       CoursePoint_SharpLeft,
	"SlightRight":     CoursePoint_SlightRight,
	"SharpRight":      CoursePoint_SharpRight,
	"UTurn":           CoursePoint_UTurn,
	"SegmentStart":    CoursePoint_SegmentStart,
	"SegmentEnd":      CoursePoint_SegmentEnd,
	"Campsite":        CoursePoint_Campsite,
	"AidStation":      CoursePoint_AidStation,
	"RestArea":        CoursePoint_RestArea,
	"GeneralDistance": CoursePoint_GeneralDistance,
	"Service":         CoursePoint_Service,
	"EnergyGel":       CoursePoint_EnergyGel,
	"SportsDrink":     CoursePoint_SportsDrink,
	"MileMarker":      CoursePoint_MileMarker,
	"Checkpoint":      CoursePoint_Checkpoint,
	"Shelter":         CoursePoint_Shelter,
	"MeetingSpot":     CoursePoint_MeetingSpot,
	"Overlook":        CoursePoint_Overlook,
	"Toilet":          CoursePoint_Toilet,
	"Shower":          CoursePoint_Shower,
	"Gear":            CoursePoint_Gear,
	"SharpCurve":      CoursePoint_SharpCurve,
	"SteepIncline":    CoursePoint_SteepIncline,
	"Tunnel":          CoursePoint_Tunnel,
	"Bridge":          CoursePoint_Bridge,
	"Obstacle":        CoursePoint_Obstacle,
	"Crossing":        CoursePoint_Crossing,
	"Store":           CoursePoint_Store,
	"Transition":      CoursePoint_Transition,
	"Navaid":          CoursePoint_Navaid,
	"Transport":       CoursePoint_Transport,
	"Alert":           CoursePoint_Alert,
	"Info":            CoursePoint_Info,
}

func CoursePointFromString(s string) CoursePoint {
	if v, ok := CoursePointValues[s]; ok {
		return v
	}
	return CoursePoint(CoursePoint_Invalid)
}
