package fitDecl

import strconv "strconv"

type ExdDisplayType byte

const (
	ExdDisplayType_Numerical         ExdDisplayType = 0
	ExdDisplayType_Simple            ExdDisplayType = 1
	ExdDisplayType_Graph             ExdDisplayType = 2
	ExdDisplayType_Bar               ExdDisplayType = 3
	ExdDisplayType_CircleGraph       ExdDisplayType = 4
	ExdDisplayType_VirtualPartner    ExdDisplayType = 5
	ExdDisplayType_Balance           ExdDisplayType = 6
	ExdDisplayType_StringList        ExdDisplayType = 7
	ExdDisplayType_String            ExdDisplayType = 8
	ExdDisplayType_SimpleDynamicIcon ExdDisplayType = 9
	ExdDisplayType_Gauge             ExdDisplayType = 10
	ExdDisplayType_Invalid           ExdDisplayType = 255
)

var ExdDisplayTypeNames = map[ExdDisplayType]string{
	ExdDisplayType_Numerical:         "Numerical",
	ExdDisplayType_Simple:            "Simple",
	ExdDisplayType_Graph:             "Graph",
	ExdDisplayType_Bar:               "Bar",
	ExdDisplayType_CircleGraph:       "CircleGraph",
	ExdDisplayType_VirtualPartner:    "VirtualPartner",
	ExdDisplayType_Balance:           "Balance",
	ExdDisplayType_StringList:        "StringList",
	ExdDisplayType_String:            "String",
	ExdDisplayType_SimpleDynamicIcon: "SimpleDynamicIcon",
	ExdDisplayType_Gauge:             "Gauge",
}

func (k ExdDisplayType) String() string {
	if uint(k) < uint(len(ExdDisplayTypeNames)) {
		if v, ok := ExdDisplayTypeNames[k]; ok {
			return v
		}
	}
	return "ExdDisplayType(" + strconv.Itoa(int(k)) + ")"
}

var ExdDisplayTypeValues = map[string]ExdDisplayType{
	"Numerical":         ExdDisplayType_Numerical,
	"Simple":            ExdDisplayType_Simple,
	"Graph":             ExdDisplayType_Graph,
	"Bar":               ExdDisplayType_Bar,
	"CircleGraph":       ExdDisplayType_CircleGraph,
	"VirtualPartner":    ExdDisplayType_VirtualPartner,
	"Balance":           ExdDisplayType_Balance,
	"StringList":        ExdDisplayType_StringList,
	"String":            ExdDisplayType_String,
	"SimpleDynamicIcon": ExdDisplayType_SimpleDynamicIcon,
	"Gauge":             ExdDisplayType_Gauge,
}

func ExdDisplayTypeFromString(s string) ExdDisplayType {
	if v, ok := ExdDisplayTypeValues[s]; ok {
		return v
	}
	return ExdDisplayType(ExdDisplayType_Invalid)
}
