package fitDecl

import strconv "strconv"

type SegmentSelectionType byte

const (
	SegmentSelectionType_Starred   SegmentSelectionType = 0
	SegmentSelectionType_Suggested SegmentSelectionType = 1
	SegmentSelectionType_Invalid   SegmentSelectionType = 255
)

var SegmentSelectionTypeNames = map[SegmentSelectionType]string{
	SegmentSelectionType_Starred:   "Starred",
	SegmentSelectionType_Suggested: "Suggested",
}

func (k SegmentSelectionType) String() string {
	if uint(k) < uint(len(SegmentSelectionTypeNames)) {
		if v, ok := SegmentSelectionTypeNames[k]; ok {
			return v
		}
	}
	return "SegmentSelectionType(" + strconv.Itoa(int(k)) + ")"
}

var SegmentSelectionTypeValues = map[string]SegmentSelectionType{
	"Starred":   SegmentSelectionType_Starred,
	"Suggested": SegmentSelectionType_Suggested,
}

func SegmentSelectionTypeFromString(s string) SegmentSelectionType {
	if v, ok := SegmentSelectionTypeValues[s]; ok {
		return v
	}
	return SegmentSelectionType(SegmentSelectionType_Invalid)
}
