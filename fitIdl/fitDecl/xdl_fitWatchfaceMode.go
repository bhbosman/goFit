package fitDecl

import strconv "strconv"

type WatchfaceMode byte

const (
	WatchfaceMode_Digital   WatchfaceMode = 0
	WatchfaceMode_Analog    WatchfaceMode = 1
	WatchfaceMode_ConnectIq WatchfaceMode = 2
	WatchfaceMode_Disabled  WatchfaceMode = 3
	WatchfaceMode_Invalid   WatchfaceMode = 255
)

var WatchfaceModeNames = map[WatchfaceMode]string{
	WatchfaceMode_Digital:   "Digital",
	WatchfaceMode_Analog:    "Analog",
	WatchfaceMode_ConnectIq: "ConnectIq",
	WatchfaceMode_Disabled:  "Disabled",
}

func (k WatchfaceMode) String() string {
	if uint(k) < uint(len(WatchfaceModeNames)) {
		if v, ok := WatchfaceModeNames[k]; ok {
			return v
		}
	}
	return "WatchfaceMode(" + strconv.Itoa(int(k)) + ")"
}

var WatchfaceModeValues = map[string]WatchfaceMode{
	"Digital":   WatchfaceMode_Digital,
	"Analog":    WatchfaceMode_Analog,
	"ConnectIq": WatchfaceMode_ConnectIq,
	"Disabled":  WatchfaceMode_Disabled,
}

func WatchfaceModeFromString(s string) WatchfaceMode {
	if v, ok := WatchfaceModeValues[s]; ok {
		return v
	}
	return WatchfaceMode(WatchfaceMode_Invalid)
}
