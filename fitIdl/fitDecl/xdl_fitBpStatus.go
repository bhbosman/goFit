package fitDecl

import strconv "strconv"

type BpStatus byte

const (
	BpStatus_NoError                 BpStatus = 0
	BpStatus_ErrorIncompleteData     BpStatus = 1
	BpStatus_ErrorNoMeasurement      BpStatus = 2
	BpStatus_ErrorDataOutOfRange     BpStatus = 3
	BpStatus_ErrorIrregularHeartRate BpStatus = 4
	BpStatus_Invalid                 BpStatus = 255
)

var BpStatusNames = map[BpStatus]string{
	BpStatus_NoError:                 "NoError",
	BpStatus_ErrorIncompleteData:     "ErrorIncompleteData",
	BpStatus_ErrorNoMeasurement:      "ErrorNoMeasurement",
	BpStatus_ErrorDataOutOfRange:     "ErrorDataOutOfRange",
	BpStatus_ErrorIrregularHeartRate: "ErrorIrregularHeartRate",
}

func (k BpStatus) String() string {
	if uint(k) < uint(len(BpStatusNames)) {
		if v, ok := BpStatusNames[k]; ok {
			return v
		}
	}
	return "BpStatus(" + strconv.Itoa(int(k)) + ")"
}

var BpStatusValues = map[string]BpStatus{
	"NoError":                 BpStatus_NoError,
	"ErrorIncompleteData":     BpStatus_ErrorIncompleteData,
	"ErrorNoMeasurement":      BpStatus_ErrorNoMeasurement,
	"ErrorDataOutOfRange":     BpStatus_ErrorDataOutOfRange,
	"ErrorIrregularHeartRate": BpStatus_ErrorIrregularHeartRate,
}

func BpStatusFromString(s string) BpStatus {
	if v, ok := BpStatusValues[s]; ok {
		return v
	}
	return BpStatus(BpStatus_Invalid)
}
