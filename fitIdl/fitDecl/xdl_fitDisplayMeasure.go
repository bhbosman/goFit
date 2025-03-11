package fitDecl

import strconv "strconv"

type DisplayMeasure byte

const (
	DisplayMeasure_Metric   DisplayMeasure = 0
	DisplayMeasure_Statute  DisplayMeasure = 1
	DisplayMeasure_Nautical DisplayMeasure = 2
	DisplayMeasure_Invalid  DisplayMeasure = 255
)

var DisplayMeasureNames = map[DisplayMeasure]string{
	DisplayMeasure_Metric:   "Metric",
	DisplayMeasure_Statute:  "Statute",
	DisplayMeasure_Nautical: "Nautical",
}

func (k DisplayMeasure) String() string {
	if uint(k) < uint(len(DisplayMeasureNames)) {
		if v, ok := DisplayMeasureNames[k]; ok {
			return v
		}
	}
	return "DisplayMeasure(" + strconv.Itoa(int(k)) + ")"
}

var DisplayMeasureValues = map[string]DisplayMeasure{
	"Metric":   DisplayMeasure_Metric,
	"Statute":  DisplayMeasure_Statute,
	"Nautical": DisplayMeasure_Nautical,
}

func DisplayMeasureFromString(s string) DisplayMeasure {
	if v, ok := DisplayMeasureValues[s]; ok {
		return v
	}
	return DisplayMeasure(DisplayMeasure_Invalid)
}
