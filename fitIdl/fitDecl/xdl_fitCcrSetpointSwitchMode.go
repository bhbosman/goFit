package fitDecl

import strconv "strconv"

type CcrSetpointSwitchMode byte

const (
	CcrSetpointSwitchMode_Manual    CcrSetpointSwitchMode = 0
	CcrSetpointSwitchMode_Automatic CcrSetpointSwitchMode = 1
	CcrSetpointSwitchMode_Invalid   CcrSetpointSwitchMode = 255
)

var CcrSetpointSwitchModeNames = map[CcrSetpointSwitchMode]string{
	CcrSetpointSwitchMode_Manual:    "Manual",
	CcrSetpointSwitchMode_Automatic: "Automatic",
}

func (k CcrSetpointSwitchMode) String() string {
	if uint(k) < uint(len(CcrSetpointSwitchModeNames)) {
		if v, ok := CcrSetpointSwitchModeNames[k]; ok {
			return v
		}
	}
	return "CcrSetpointSwitchMode(" + strconv.Itoa(int(k)) + ")"
}

var CcrSetpointSwitchModeValues = map[string]CcrSetpointSwitchMode{
	"Manual":    CcrSetpointSwitchMode_Manual,
	"Automatic": CcrSetpointSwitchMode_Automatic,
}

func CcrSetpointSwitchModeFromString(s string) CcrSetpointSwitchMode {
	if v, ok := CcrSetpointSwitchModeValues[s]; ok {
		return v
	}
	return CcrSetpointSwitchMode(CcrSetpointSwitchMode_Invalid)
}
