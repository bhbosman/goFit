package fitDecl

import strconv "strconv"

type DigitalWatchfaceLayout byte

const (
	DigitalWatchfaceLayout_Traditional DigitalWatchfaceLayout = 0
	DigitalWatchfaceLayout_Modern      DigitalWatchfaceLayout = 1
	DigitalWatchfaceLayout_Bold        DigitalWatchfaceLayout = 2
	DigitalWatchfaceLayout_Invalid     DigitalWatchfaceLayout = 255
)

var DigitalWatchfaceLayoutNames = map[DigitalWatchfaceLayout]string{
	DigitalWatchfaceLayout_Traditional: "Traditional",
	DigitalWatchfaceLayout_Modern:      "Modern",
	DigitalWatchfaceLayout_Bold:        "Bold",
}

func (k DigitalWatchfaceLayout) String() string {
	if uint(k) < uint(len(DigitalWatchfaceLayoutNames)) {
		if v, ok := DigitalWatchfaceLayoutNames[k]; ok {
			return v
		}
	}
	return "DigitalWatchfaceLayout(" + strconv.Itoa(int(k)) + ")"
}

var DigitalWatchfaceLayoutValues = map[string]DigitalWatchfaceLayout{
	"Traditional": DigitalWatchfaceLayout_Traditional,
	"Modern":      DigitalWatchfaceLayout_Modern,
	"Bold":        DigitalWatchfaceLayout_Bold,
}

func DigitalWatchfaceLayoutFromString(s string) DigitalWatchfaceLayout {
	if v, ok := DigitalWatchfaceLayoutValues[s]; ok {
		return v
	}
	return DigitalWatchfaceLayout(DigitalWatchfaceLayout_Invalid)
}
