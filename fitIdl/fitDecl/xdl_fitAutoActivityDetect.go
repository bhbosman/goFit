package fitDecl

import strconv "strconv"

type AutoActivityDetect uint32

const (
	AutoActivityDetect_None       AutoActivityDetect = 0
	AutoActivityDetect_Running    AutoActivityDetect = 1
	AutoActivityDetect_Cycling    AutoActivityDetect = 2
	AutoActivityDetect_Swimming   AutoActivityDetect = 4
	AutoActivityDetect_Walking    AutoActivityDetect = 8
	AutoActivityDetect_Elliptical AutoActivityDetect = 32
	AutoActivityDetect_Sedentary  AutoActivityDetect = 1024
	AutoActivityDetect_Invalid    AutoActivityDetect = 4294967295
)

var AutoActivityDetectNames = map[AutoActivityDetect]string{
	AutoActivityDetect_None:       "None",
	AutoActivityDetect_Running:    "Running",
	AutoActivityDetect_Cycling:    "Cycling",
	AutoActivityDetect_Swimming:   "Swimming",
	AutoActivityDetect_Walking:    "Walking",
	AutoActivityDetect_Elliptical: "Elliptical",
	AutoActivityDetect_Sedentary:  "Sedentary",
}

func (k AutoActivityDetect) String() string {
	if uint(k) < uint(len(AutoActivityDetectNames)) {
		if v, ok := AutoActivityDetectNames[k]; ok {
			return v
		}
	}
	return "AutoActivityDetect(" + strconv.Itoa(int(k)) + ")"
}

var AutoActivityDetectValues = map[string]AutoActivityDetect{
	"None":       AutoActivityDetect_None,
	"Running":    AutoActivityDetect_Running,
	"Cycling":    AutoActivityDetect_Cycling,
	"Swimming":   AutoActivityDetect_Swimming,
	"Walking":    AutoActivityDetect_Walking,
	"Elliptical": AutoActivityDetect_Elliptical,
	"Sedentary":  AutoActivityDetect_Sedentary,
}

func AutoActivityDetectFromString(s string) AutoActivityDetect {
	if v, ok := AutoActivityDetectValues[s]; ok {
		return v
	}
	return AutoActivityDetect(AutoActivityDetect_Invalid)
}
