package fitDecl

import strconv "strconv"

type DiveGasStatus byte

const (
	DiveGasStatus_Disabled   DiveGasStatus = 0
	DiveGasStatus_Enabled    DiveGasStatus = 1
	DiveGasStatus_BackupOnly DiveGasStatus = 2
	DiveGasStatus_Invalid    DiveGasStatus = 255
)

var DiveGasStatusNames = map[DiveGasStatus]string{
	DiveGasStatus_Disabled:   "Disabled",
	DiveGasStatus_Enabled:    "Enabled",
	DiveGasStatus_BackupOnly: "BackupOnly",
}

func (k DiveGasStatus) String() string {
	if uint(k) < uint(len(DiveGasStatusNames)) {
		if v, ok := DiveGasStatusNames[k]; ok {
			return v
		}
	}
	return "DiveGasStatus(" + strconv.Itoa(int(k)) + ")"
}

var DiveGasStatusValues = map[string]DiveGasStatus{
	"Disabled":   DiveGasStatus_Disabled,
	"Enabled":    DiveGasStatus_Enabled,
	"BackupOnly": DiveGasStatus_BackupOnly,
}

func DiveGasStatusFromString(s string) DiveGasStatus {
	if v, ok := DiveGasStatusValues[s]; ok {
		return v
	}
	return DiveGasStatus(DiveGasStatus_Invalid)
}
