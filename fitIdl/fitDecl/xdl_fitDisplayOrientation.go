package fitDecl

import strconv "strconv"

type DisplayOrientation byte

const (
	DisplayOrientation_Auto             DisplayOrientation = 0
	DisplayOrientation_Portrait         DisplayOrientation = 1
	DisplayOrientation_Landscape        DisplayOrientation = 2
	DisplayOrientation_PortraitFlipped  DisplayOrientation = 3
	DisplayOrientation_LandscapeFlipped DisplayOrientation = 4
	DisplayOrientation_Invalid          DisplayOrientation = 255
)

var DisplayOrientationNames = map[DisplayOrientation]string{
	DisplayOrientation_Auto:             "Auto",
	DisplayOrientation_Portrait:         "Portrait",
	DisplayOrientation_Landscape:        "Landscape",
	DisplayOrientation_PortraitFlipped:  "PortraitFlipped",
	DisplayOrientation_LandscapeFlipped: "LandscapeFlipped",
}

func (k DisplayOrientation) String() string {
	if uint(k) < uint(len(DisplayOrientationNames)) {
		if v, ok := DisplayOrientationNames[k]; ok {
			return v
		}
	}
	return "DisplayOrientation(" + strconv.Itoa(int(k)) + ")"
}

var DisplayOrientationValues = map[string]DisplayOrientation{
	"Auto":             DisplayOrientation_Auto,
	"Portrait":         DisplayOrientation_Portrait,
	"Landscape":        DisplayOrientation_Landscape,
	"PortraitFlipped":  DisplayOrientation_PortraitFlipped,
	"LandscapeFlipped": DisplayOrientation_LandscapeFlipped,
}

func DisplayOrientationFromString(s string) DisplayOrientation {
	if v, ok := DisplayOrientationValues[s]; ok {
		return v
	}
	return DisplayOrientation(DisplayOrientation_Invalid)
}
