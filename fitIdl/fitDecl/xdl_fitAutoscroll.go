package fitDecl

import strconv "strconv"

type Autoscroll byte

const (
	Autoscroll_None    Autoscroll = 0
	Autoscroll_Slow    Autoscroll = 1
	Autoscroll_Medium  Autoscroll = 2
	Autoscroll_Fast    Autoscroll = 3
	Autoscroll_Invalid Autoscroll = 255
)

var AutoscrollNames = map[Autoscroll]string{
	Autoscroll_None:   "None",
	Autoscroll_Slow:   "Slow",
	Autoscroll_Medium: "Medium",
	Autoscroll_Fast:   "Fast",
}

func (k Autoscroll) String() string {
	if uint(k) < uint(len(AutoscrollNames)) {
		if v, ok := AutoscrollNames[k]; ok {
			return v
		}
	}
	return "Autoscroll(" + strconv.Itoa(int(k)) + ")"
}

var AutoscrollValues = map[string]Autoscroll{
	"None":   Autoscroll_None,
	"Slow":   Autoscroll_Slow,
	"Medium": Autoscroll_Medium,
	"Fast":   Autoscroll_Fast,
}

func AutoscrollFromString(s string) Autoscroll {
	if v, ok := AutoscrollValues[s]; ok {
		return v
	}
	return Autoscroll(Autoscroll_Invalid)
}
