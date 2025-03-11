package fitDecl

import strconv "strconv"

type SwimStroke byte

const (
	SwimStroke_Freestyle    SwimStroke = 0
	SwimStroke_Backstroke   SwimStroke = 1
	SwimStroke_Breaststroke SwimStroke = 2
	SwimStroke_Butterfly    SwimStroke = 3
	SwimStroke_Drill        SwimStroke = 4
	SwimStroke_Mixed        SwimStroke = 5
	SwimStroke_Im           SwimStroke = 6
	SwimStroke_Invalid      SwimStroke = 255
)

var SwimStrokeNames = map[SwimStroke]string{
	SwimStroke_Freestyle:    "Freestyle",
	SwimStroke_Backstroke:   "Backstroke",
	SwimStroke_Breaststroke: "Breaststroke",
	SwimStroke_Butterfly:    "Butterfly",
	SwimStroke_Drill:        "Drill",
	SwimStroke_Mixed:        "Mixed",
	SwimStroke_Im:           "Im",
}

func (k SwimStroke) String() string {
	if uint(k) < uint(len(SwimStrokeNames)) {
		if v, ok := SwimStrokeNames[k]; ok {
			return v
		}
	}
	return "SwimStroke(" + strconv.Itoa(int(k)) + ")"
}

var SwimStrokeValues = map[string]SwimStroke{
	"Freestyle":    SwimStroke_Freestyle,
	"Backstroke":   SwimStroke_Backstroke,
	"Breaststroke": SwimStroke_Breaststroke,
	"Butterfly":    SwimStroke_Butterfly,
	"Drill":        SwimStroke_Drill,
	"Mixed":        SwimStroke_Mixed,
	"Im":           SwimStroke_Im,
}

func SwimStrokeFromString(s string) SwimStroke {
	if v, ok := SwimStrokeValues[s]; ok {
		return v
	}
	return SwimStroke(SwimStroke_Invalid)
}
