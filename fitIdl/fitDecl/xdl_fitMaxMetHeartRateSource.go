package fitDecl

import strconv "strconv"

type MaxMetHeartRateSource byte

const (
	MaxMetHeartRateSource_Whr     MaxMetHeartRateSource = 0
	MaxMetHeartRateSource_Hrm     MaxMetHeartRateSource = 1
	MaxMetHeartRateSource_Invalid MaxMetHeartRateSource = 255
)

var MaxMetHeartRateSourceNames = map[MaxMetHeartRateSource]string{
	MaxMetHeartRateSource_Whr: "Whr",
	MaxMetHeartRateSource_Hrm: "Hrm",
}

func (k MaxMetHeartRateSource) String() string {
	if uint(k) < uint(len(MaxMetHeartRateSourceNames)) {
		if v, ok := MaxMetHeartRateSourceNames[k]; ok {
			return v
		}
	}
	return "MaxMetHeartRateSource(" + strconv.Itoa(int(k)) + ")"
}

var MaxMetHeartRateSourceValues = map[string]MaxMetHeartRateSource{
	"Whr": MaxMetHeartRateSource_Whr,
	"Hrm": MaxMetHeartRateSource_Hrm,
}

func MaxMetHeartRateSourceFromString(s string) MaxMetHeartRateSource {
	if v, ok := MaxMetHeartRateSourceValues[s]; ok {
		return v
	}
	return MaxMetHeartRateSource(MaxMetHeartRateSource_Invalid)
}
