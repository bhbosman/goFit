package fitDecl

import strconv "strconv"

type SegmentLeaderboardType byte

const (
	SegmentLeaderboardType_Overall      SegmentLeaderboardType = 0
	SegmentLeaderboardType_PersonalBest SegmentLeaderboardType = 1
	SegmentLeaderboardType_Connections  SegmentLeaderboardType = 2
	SegmentLeaderboardType_Group        SegmentLeaderboardType = 3
	SegmentLeaderboardType_Challenger   SegmentLeaderboardType = 4
	SegmentLeaderboardType_Kom          SegmentLeaderboardType = 5
	SegmentLeaderboardType_Qom          SegmentLeaderboardType = 6
	SegmentLeaderboardType_Pr           SegmentLeaderboardType = 7
	SegmentLeaderboardType_Goal         SegmentLeaderboardType = 8
	SegmentLeaderboardType_Carrot       SegmentLeaderboardType = 9
	SegmentLeaderboardType_ClubLeader   SegmentLeaderboardType = 10
	SegmentLeaderboardType_Rival        SegmentLeaderboardType = 11
	SegmentLeaderboardType_Last         SegmentLeaderboardType = 12
	SegmentLeaderboardType_RecentBest   SegmentLeaderboardType = 13
	SegmentLeaderboardType_CourseRecord SegmentLeaderboardType = 14
	SegmentLeaderboardType_Invalid      SegmentLeaderboardType = 255
)

var SegmentLeaderboardTypeNames = map[SegmentLeaderboardType]string{
	SegmentLeaderboardType_Overall:      "Overall",
	SegmentLeaderboardType_PersonalBest: "PersonalBest",
	SegmentLeaderboardType_Connections:  "Connections",
	SegmentLeaderboardType_Group:        "Group",
	SegmentLeaderboardType_Challenger:   "Challenger",
	SegmentLeaderboardType_Kom:          "Kom",
	SegmentLeaderboardType_Qom:          "Qom",
	SegmentLeaderboardType_Pr:           "Pr",
	SegmentLeaderboardType_Goal:         "Goal",
	SegmentLeaderboardType_Carrot:       "Carrot",
	SegmentLeaderboardType_ClubLeader:   "ClubLeader",
	SegmentLeaderboardType_Rival:        "Rival",
	SegmentLeaderboardType_Last:         "Last",
	SegmentLeaderboardType_RecentBest:   "RecentBest",
	SegmentLeaderboardType_CourseRecord: "CourseRecord",
}

func (k SegmentLeaderboardType) String() string {
	if uint(k) < uint(len(SegmentLeaderboardTypeNames)) {
		if v, ok := SegmentLeaderboardTypeNames[k]; ok {
			return v
		}
	}
	return "SegmentLeaderboardType(" + strconv.Itoa(int(k)) + ")"
}

var SegmentLeaderboardTypeValues = map[string]SegmentLeaderboardType{
	"Overall":      SegmentLeaderboardType_Overall,
	"PersonalBest": SegmentLeaderboardType_PersonalBest,
	"Connections":  SegmentLeaderboardType_Connections,
	"Group":        SegmentLeaderboardType_Group,
	"Challenger":   SegmentLeaderboardType_Challenger,
	"Kom":          SegmentLeaderboardType_Kom,
	"Qom":          SegmentLeaderboardType_Qom,
	"Pr":           SegmentLeaderboardType_Pr,
	"Goal":         SegmentLeaderboardType_Goal,
	"Carrot":       SegmentLeaderboardType_Carrot,
	"ClubLeader":   SegmentLeaderboardType_ClubLeader,
	"Rival":        SegmentLeaderboardType_Rival,
	"Last":         SegmentLeaderboardType_Last,
	"RecentBest":   SegmentLeaderboardType_RecentBest,
	"CourseRecord": SegmentLeaderboardType_CourseRecord,
}

func SegmentLeaderboardTypeFromString(s string) SegmentLeaderboardType {
	if v, ok := SegmentLeaderboardTypeValues[s]; ok {
		return v
	}
	return SegmentLeaderboardType(SegmentLeaderboardType_Invalid)
}
