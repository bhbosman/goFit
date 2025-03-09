package fitDecl

import strconv "strconv"

type TurnType byte

const (
	TurnType_ArrivingIdx             TurnType = 0
	TurnType_ArrivingLeftIdx         TurnType = 1
	TurnType_ArrivingRightIdx        TurnType = 2
	TurnType_ArrivingViaIdx          TurnType = 3
	TurnType_ArrivingViaLeftIdx      TurnType = 4
	TurnType_ArrivingViaRightIdx     TurnType = 5
	TurnType_BearKeepLeftIdx         TurnType = 6
	TurnType_BearKeepRightIdx        TurnType = 7
	TurnType_ContinueIdx             TurnType = 8
	TurnType_ExitLeftIdx             TurnType = 9
	TurnType_ExitRightIdx            TurnType = 10
	TurnType_FerryIdx                TurnType = 11
	TurnType_Roundabout45Idx         TurnType = 12
	TurnType_Roundabout90Idx         TurnType = 13
	TurnType_Roundabout135Idx        TurnType = 14
	TurnType_Roundabout180Idx        TurnType = 15
	TurnType_Roundabout225Idx        TurnType = 16
	TurnType_Roundabout270Idx        TurnType = 17
	TurnType_Roundabout315Idx        TurnType = 18
	TurnType_Roundabout360Idx        TurnType = 19
	TurnType_RoundaboutNeg45Idx      TurnType = 20
	TurnType_RoundaboutNeg90Idx      TurnType = 21
	TurnType_RoundaboutNeg135Idx     TurnType = 22
	TurnType_RoundaboutNeg180Idx     TurnType = 23
	TurnType_RoundaboutNeg225Idx     TurnType = 24
	TurnType_RoundaboutNeg270Idx     TurnType = 25
	TurnType_RoundaboutNeg315Idx     TurnType = 26
	TurnType_RoundaboutNeg360Idx     TurnType = 27
	TurnType_RoundaboutGenericIdx    TurnType = 28
	TurnType_RoundaboutNegGenericIdx TurnType = 29
	TurnType_SharpTurnLeftIdx        TurnType = 30
	TurnType_SharpTurnRightIdx       TurnType = 31
	TurnType_TurnLeftIdx             TurnType = 32
	TurnType_TurnRightIdx            TurnType = 33
	TurnType_UturnLeftIdx            TurnType = 34
	TurnType_UturnRightIdx           TurnType = 35
	TurnType_IconInvIdx              TurnType = 36
	TurnType_IconIdxCnt              TurnType = 37
	TurnType_Invalid                 TurnType = 255
)

var TurnTypeNames = map[TurnType]string{
	TurnType_ArrivingIdx:             "ArrivingIdx",
	TurnType_ArrivingLeftIdx:         "ArrivingLeftIdx",
	TurnType_ArrivingRightIdx:        "ArrivingRightIdx",
	TurnType_ArrivingViaIdx:          "ArrivingViaIdx",
	TurnType_ArrivingViaLeftIdx:      "ArrivingViaLeftIdx",
	TurnType_ArrivingViaRightIdx:     "ArrivingViaRightIdx",
	TurnType_BearKeepLeftIdx:         "BearKeepLeftIdx",
	TurnType_BearKeepRightIdx:        "BearKeepRightIdx",
	TurnType_ContinueIdx:             "ContinueIdx",
	TurnType_ExitLeftIdx:             "ExitLeftIdx",
	TurnType_ExitRightIdx:            "ExitRightIdx",
	TurnType_FerryIdx:                "FerryIdx",
	TurnType_Roundabout45Idx:         "Roundabout45Idx",
	TurnType_Roundabout90Idx:         "Roundabout90Idx",
	TurnType_Roundabout135Idx:        "Roundabout135Idx",
	TurnType_Roundabout180Idx:        "Roundabout180Idx",
	TurnType_Roundabout225Idx:        "Roundabout225Idx",
	TurnType_Roundabout270Idx:        "Roundabout270Idx",
	TurnType_Roundabout315Idx:        "Roundabout315Idx",
	TurnType_Roundabout360Idx:        "Roundabout360Idx",
	TurnType_RoundaboutNeg45Idx:      "RoundaboutNeg45Idx",
	TurnType_RoundaboutNeg90Idx:      "RoundaboutNeg90Idx",
	TurnType_RoundaboutNeg135Idx:     "RoundaboutNeg135Idx",
	TurnType_RoundaboutNeg180Idx:     "RoundaboutNeg180Idx",
	TurnType_RoundaboutNeg225Idx:     "RoundaboutNeg225Idx",
	TurnType_RoundaboutNeg270Idx:     "RoundaboutNeg270Idx",
	TurnType_RoundaboutNeg315Idx:     "RoundaboutNeg315Idx",
	TurnType_RoundaboutNeg360Idx:     "RoundaboutNeg360Idx",
	TurnType_RoundaboutGenericIdx:    "RoundaboutGenericIdx",
	TurnType_RoundaboutNegGenericIdx: "RoundaboutNegGenericIdx",
	TurnType_SharpTurnLeftIdx:        "SharpTurnLeftIdx",
	TurnType_SharpTurnRightIdx:       "SharpTurnRightIdx",
	TurnType_TurnLeftIdx:             "TurnLeftIdx",
	TurnType_TurnRightIdx:            "TurnRightIdx",
	TurnType_UturnLeftIdx:            "UturnLeftIdx",
	TurnType_UturnRightIdx:           "UturnRightIdx",
	TurnType_IconInvIdx:              "IconInvIdx",
	TurnType_IconIdxCnt:              "IconIdxCnt",
}

func (k TurnType) String() string {
	if uint(k) < uint(len(TurnTypeNames)) {
		if v, ok := TurnTypeNames[k]; ok {
			return v
		}
	}
	return "TurnType(" + strconv.Itoa(int(k)) + ")"
}

var TurnTypeValues = map[string]TurnType{
	"ArrivingIdx":             TurnType_ArrivingIdx,
	"ArrivingLeftIdx":         TurnType_ArrivingLeftIdx,
	"ArrivingRightIdx":        TurnType_ArrivingRightIdx,
	"ArrivingViaIdx":          TurnType_ArrivingViaIdx,
	"ArrivingViaLeftIdx":      TurnType_ArrivingViaLeftIdx,
	"ArrivingViaRightIdx":     TurnType_ArrivingViaRightIdx,
	"BearKeepLeftIdx":         TurnType_BearKeepLeftIdx,
	"BearKeepRightIdx":        TurnType_BearKeepRightIdx,
	"ContinueIdx":             TurnType_ContinueIdx,
	"ExitLeftIdx":             TurnType_ExitLeftIdx,
	"ExitRightIdx":            TurnType_ExitRightIdx,
	"FerryIdx":                TurnType_FerryIdx,
	"Roundabout45Idx":         TurnType_Roundabout45Idx,
	"Roundabout90Idx":         TurnType_Roundabout90Idx,
	"Roundabout135Idx":        TurnType_Roundabout135Idx,
	"Roundabout180Idx":        TurnType_Roundabout180Idx,
	"Roundabout225Idx":        TurnType_Roundabout225Idx,
	"Roundabout270Idx":        TurnType_Roundabout270Idx,
	"Roundabout315Idx":        TurnType_Roundabout315Idx,
	"Roundabout360Idx":        TurnType_Roundabout360Idx,
	"RoundaboutNeg45Idx":      TurnType_RoundaboutNeg45Idx,
	"RoundaboutNeg90Idx":      TurnType_RoundaboutNeg90Idx,
	"RoundaboutNeg135Idx":     TurnType_RoundaboutNeg135Idx,
	"RoundaboutNeg180Idx":     TurnType_RoundaboutNeg180Idx,
	"RoundaboutNeg225Idx":     TurnType_RoundaboutNeg225Idx,
	"RoundaboutNeg270Idx":     TurnType_RoundaboutNeg270Idx,
	"RoundaboutNeg315Idx":     TurnType_RoundaboutNeg315Idx,
	"RoundaboutNeg360Idx":     TurnType_RoundaboutNeg360Idx,
	"RoundaboutGenericIdx":    TurnType_RoundaboutGenericIdx,
	"RoundaboutNegGenericIdx": TurnType_RoundaboutNegGenericIdx,
	"SharpTurnLeftIdx":        TurnType_SharpTurnLeftIdx,
	"SharpTurnRightIdx":       TurnType_SharpTurnRightIdx,
	"TurnLeftIdx":             TurnType_TurnLeftIdx,
	"TurnRightIdx":            TurnType_TurnRightIdx,
	"UturnLeftIdx":            TurnType_UturnLeftIdx,
	"UturnRightIdx":           TurnType_UturnRightIdx,
	"IconInvIdx":              TurnType_IconInvIdx,
	"IconIdxCnt":              TurnType_IconIdxCnt,
}

func TurnTypeFromString(s string) TurnType {
	if v, ok := TurnTypeValues[s]; ok {
		return v
	}
	return TurnType(TurnType_Invalid)
}
