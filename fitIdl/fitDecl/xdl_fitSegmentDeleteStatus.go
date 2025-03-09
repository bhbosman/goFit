package fitDecl

import strconv "strconv"

type SegmentDeleteStatus byte

const (
	SegmentDeleteStatus_DoNotDelete SegmentDeleteStatus = 0
	SegmentDeleteStatus_DeleteOne   SegmentDeleteStatus = 1
	SegmentDeleteStatus_DeleteAll   SegmentDeleteStatus = 2
	SegmentDeleteStatus_Invalid     SegmentDeleteStatus = 255
)

var SegmentDeleteStatusNames = map[SegmentDeleteStatus]string{
	SegmentDeleteStatus_DoNotDelete: "DoNotDelete",
	SegmentDeleteStatus_DeleteOne:   "DeleteOne",
	SegmentDeleteStatus_DeleteAll:   "DeleteAll",
}

func (k SegmentDeleteStatus) String() string {
	if uint(k) < uint(len(SegmentDeleteStatusNames)) {
		if v, ok := SegmentDeleteStatusNames[k]; ok {
			return v
		}
	}
	return "SegmentDeleteStatus(" + strconv.Itoa(int(k)) + ")"
}

var SegmentDeleteStatusValues = map[string]SegmentDeleteStatus{
	"DoNotDelete": SegmentDeleteStatus_DoNotDelete,
	"DeleteOne":   SegmentDeleteStatus_DeleteOne,
	"DeleteAll":   SegmentDeleteStatus_DeleteAll,
}

func SegmentDeleteStatusFromString(s string) SegmentDeleteStatus {
	if v, ok := SegmentDeleteStatusValues[s]; ok {
		return v
	}
	return SegmentDeleteStatus(SegmentDeleteStatus_Invalid)
}
