package fitDecl

import strconv "strconv"

type AnalogWatchfaceLayout byte

const (
	AnalogWatchfaceLayout_Minimal     AnalogWatchfaceLayout = 0
	AnalogWatchfaceLayout_Traditional AnalogWatchfaceLayout = 1
	AnalogWatchfaceLayout_Modern      AnalogWatchfaceLayout = 2
	AnalogWatchfaceLayout_Invalid     AnalogWatchfaceLayout = 255
)

var AnalogWatchfaceLayoutNames = map[AnalogWatchfaceLayout]string{
	AnalogWatchfaceLayout_Minimal:     "Minimal",
	AnalogWatchfaceLayout_Traditional: "Traditional",
	AnalogWatchfaceLayout_Modern:      "Modern",
}

func (k AnalogWatchfaceLayout) String() string {
	if uint(k) < uint(len(AnalogWatchfaceLayoutNames)) {
		if v, ok := AnalogWatchfaceLayoutNames[k]; ok {
			return v
		}
	}
	return "AnalogWatchfaceLayout(" + strconv.Itoa(int(k)) + ")"
}

var AnalogWatchfaceLayoutValues = map[string]AnalogWatchfaceLayout{
	"Minimal":     AnalogWatchfaceLayout_Minimal,
	"Traditional": AnalogWatchfaceLayout_Traditional,
	"Modern":      AnalogWatchfaceLayout_Modern,
}

func AnalogWatchfaceLayoutFromString(s string) AnalogWatchfaceLayout {
	if v, ok := AnalogWatchfaceLayoutValues[s]; ok {
		return v
	}
	return AnalogWatchfaceLayout(AnalogWatchfaceLayout_Invalid)
}
