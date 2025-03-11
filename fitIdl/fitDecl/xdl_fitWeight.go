package fitDecl

import strconv "strconv"

type Weight uint16

const (
	Weight_Calculating Weight = 65534
	Weight_Invalid     Weight = 65535
)

var WeightNames = map[Weight]string{
	Weight_Calculating: "Calculating",
}

func (k Weight) String() string {
	if uint(k) < uint(len(WeightNames)) {
		if v, ok := WeightNames[k]; ok {
			return v
		}
	}
	return "Weight(" + strconv.Itoa(int(k)) + ")"
}

var WeightValues = map[string]Weight{
	"Calculating": Weight_Calculating,
}

func WeightFromString(s string) Weight {
	if v, ok := WeightValues[s]; ok {
		return v
	}
	return Weight(Weight_Invalid)
}
