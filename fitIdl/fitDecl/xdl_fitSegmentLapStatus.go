package fitDecl

import strconv "strconv"

type SegmentLapStatus byte

const (
	SegmentLapStatus_End     SegmentLapStatus = 0
	SegmentLapStatus_Fail    SegmentLapStatus = 1
	SegmentLapStatus_Invalid SegmentLapStatus = 255
)

var SegmentLapStatusNames = map[SegmentLapStatus]string{
	SegmentLapStatus_End:  "End",
	SegmentLapStatus_Fail: "Fail",
}

func (k SegmentLapStatus) String() string {
	if uint(k) < uint(len(SegmentLapStatusNames)) {
		if v, ok := SegmentLapStatusNames[k]; ok {
			return v
		}
	}
	return "SegmentLapStatus(" + strconv.Itoa(int(k)) + ")"
}

var SegmentLapStatusValues = map[string]SegmentLapStatus{
	"End":  SegmentLapStatus_End,
	"Fail": SegmentLapStatus_Fail,
}

func SegmentLapStatusFromString(s string) SegmentLapStatus {
	if v, ok := SegmentLapStatusValues[s]; ok {
		return v
	}
	return SegmentLapStatus(SegmentLapStatus_Invalid)
}
