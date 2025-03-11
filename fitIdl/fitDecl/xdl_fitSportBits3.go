package fitDecl

import strconv "strconv"

type SportBits3 uint8

const (
	SportBits3_Driving         SportBits3 = 1
	SportBits3_Golf            SportBits3 = 2
	SportBits3_HangGliding     SportBits3 = 4
	SportBits3_HorsebackRiding SportBits3 = 8
	SportBits3_Hunting         SportBits3 = 16
	SportBits3_Fishing         SportBits3 = 32
	SportBits3_InlineSkating   SportBits3 = 64
	SportBits3_RockClimbing    SportBits3 = 128
	SportBits3_Invalid         SportBits3 = 0
)

var SportBits3Names = map[SportBits3]string{
	SportBits3_Driving:         "Driving",
	SportBits3_Golf:            "Golf",
	SportBits3_HangGliding:     "HangGliding",
	SportBits3_HorsebackRiding: "HorsebackRiding",
	SportBits3_Hunting:         "Hunting",
	SportBits3_Fishing:         "Fishing",
	SportBits3_InlineSkating:   "InlineSkating",
	SportBits3_RockClimbing:    "RockClimbing",
}

func (k SportBits3) String() string {
	if uint(k) < uint(len(SportBits3Names)) {
		if v, ok := SportBits3Names[k]; ok {
			return v
		}
	}
	return "SportBits3(" + strconv.Itoa(int(k)) + ")"
}

var SportBits3Values = map[string]SportBits3{
	"Driving":         SportBits3_Driving,
	"Golf":            SportBits3_Golf,
	"HangGliding":     SportBits3_HangGliding,
	"HorsebackRiding": SportBits3_HorsebackRiding,
	"Hunting":         SportBits3_Hunting,
	"Fishing":         SportBits3_Fishing,
	"InlineSkating":   SportBits3_InlineSkating,
	"RockClimbing":    SportBits3_RockClimbing,
}

func SportBits3FromString(s string) SportBits3 {
	if v, ok := SportBits3Values[s]; ok {
		return v
	}
	return SportBits3(SportBits3_Invalid)
}
