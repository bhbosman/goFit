package fitDecl

import strconv "strconv"

type BatteryStatus uint8

const (
	BatteryStatus_New      BatteryStatus = 1
	BatteryStatus_Good     BatteryStatus = 2
	BatteryStatus_Ok       BatteryStatus = 3
	BatteryStatus_Low      BatteryStatus = 4
	BatteryStatus_Critical BatteryStatus = 5
	BatteryStatus_Charging BatteryStatus = 6
	BatteryStatus_Unknown  BatteryStatus = 7
	BatteryStatus_Invalid  BatteryStatus = 255
)

var BatteryStatusNames = map[BatteryStatus]string{
	BatteryStatus_New:      "New",
	BatteryStatus_Good:     "Good",
	BatteryStatus_Ok:       "Ok",
	BatteryStatus_Low:      "Low",
	BatteryStatus_Critical: "Critical",
	BatteryStatus_Charging: "Charging",
	BatteryStatus_Unknown:  "Unknown",
}

func (k BatteryStatus) String() string {
	if uint(k) < uint(len(BatteryStatusNames)) {
		if v, ok := BatteryStatusNames[k]; ok {
			return v
		}
	}
	return "BatteryStatus(" + strconv.Itoa(int(k)) + ")"
}

var BatteryStatusValues = map[string]BatteryStatus{
	"New":      BatteryStatus_New,
	"Good":     BatteryStatus_Good,
	"Ok":       BatteryStatus_Ok,
	"Low":      BatteryStatus_Low,
	"Critical": BatteryStatus_Critical,
	"Charging": BatteryStatus_Charging,
	"Unknown":  BatteryStatus_Unknown,
}

func BatteryStatusFromString(s string) BatteryStatus {
	if v, ok := BatteryStatusValues[s]; ok {
		return v
	}
	return BatteryStatus(BatteryStatus_Invalid)
}
