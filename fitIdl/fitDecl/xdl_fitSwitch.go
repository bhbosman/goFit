package fitDecl

import strconv "strconv"

type Switch byte

const (
	Switch_Off     Switch = 0
	Switch_On      Switch = 1
	Switch_Auto    Switch = 2
	Switch_Invalid Switch = 255
)

var SwitchNames = map[Switch]string{
	Switch_Off:  "Off",
	Switch_On:   "On",
	Switch_Auto: "Auto",
}

func (k Switch) String() string {
	if uint(k) < uint(len(SwitchNames)) {
		if v, ok := SwitchNames[k]; ok {
			return v
		}
	}
	return "Switch(" + strconv.Itoa(int(k)) + ")"
}

var SwitchValues = map[string]Switch{
	"Off":  Switch_Off,
	"On":   Switch_On,
	"Auto": Switch_Auto,
}

func SwitchFromString(s string) Switch {
	if v, ok := SwitchValues[s]; ok {
		return v
	}
	return Switch(Switch_Invalid)
}
