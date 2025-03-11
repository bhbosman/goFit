package fitDecl

import strconv "strconv"

type HrvStatus byte

const (
	HrvStatus_None       HrvStatus = 0
	HrvStatus_Poor       HrvStatus = 1
	HrvStatus_Low        HrvStatus = 2
	HrvStatus_Unbalanced HrvStatus = 3
	HrvStatus_Balanced   HrvStatus = 4
	HrvStatus_Invalid    HrvStatus = 255
)

var HrvStatusNames = map[HrvStatus]string{
	HrvStatus_None:       "None",
	HrvStatus_Poor:       "Poor",
	HrvStatus_Low:        "Low",
	HrvStatus_Unbalanced: "Unbalanced",
	HrvStatus_Balanced:   "Balanced",
}

func (k HrvStatus) String() string {
	if uint(k) < uint(len(HrvStatusNames)) {
		if v, ok := HrvStatusNames[k]; ok {
			return v
		}
	}
	return "HrvStatus(" + strconv.Itoa(int(k)) + ")"
}

var HrvStatusValues = map[string]HrvStatus{
	"None":       HrvStatus_None,
	"Poor":       HrvStatus_Poor,
	"Low":        HrvStatus_Low,
	"Unbalanced": HrvStatus_Unbalanced,
	"Balanced":   HrvStatus_Balanced,
}

func HrvStatusFromString(s string) HrvStatus {
	if v, ok := HrvStatusValues[s]; ok {
		return v
	}
	return HrvStatus(HrvStatus_Invalid)
}
