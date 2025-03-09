package fitDecl

import strconv "strconv"

type SportBits4 uint8

const (
	SportBits4_Sailing               SportBits4 = 1
	SportBits4_IceSkating            SportBits4 = 2
	SportBits4_SkyDiving             SportBits4 = 4
	SportBits4_Snowshoeing           SportBits4 = 8
	SportBits4_Snowmobiling          SportBits4 = 16
	SportBits4_StandUpPaddleboarding SportBits4 = 32
	SportBits4_Surfing               SportBits4 = 64
	SportBits4_Wakeboarding          SportBits4 = 128
	SportBits4_Invalid               SportBits4 = 0
)

var SportBits4Names = map[SportBits4]string{
	SportBits4_Sailing:               "Sailing",
	SportBits4_IceSkating:            "IceSkating",
	SportBits4_SkyDiving:             "SkyDiving",
	SportBits4_Snowshoeing:           "Snowshoeing",
	SportBits4_Snowmobiling:          "Snowmobiling",
	SportBits4_StandUpPaddleboarding: "StandUpPaddleboarding",
	SportBits4_Surfing:               "Surfing",
	SportBits4_Wakeboarding:          "Wakeboarding",
}

func (k SportBits4) String() string {
	if uint(k) < uint(len(SportBits4Names)) {
		if v, ok := SportBits4Names[k]; ok {
			return v
		}
	}
	return "SportBits4(" + strconv.Itoa(int(k)) + ")"
}

var SportBits4Values = map[string]SportBits4{
	"Sailing":               SportBits4_Sailing,
	"IceSkating":            SportBits4_IceSkating,
	"SkyDiving":             SportBits4_SkyDiving,
	"Snowshoeing":           SportBits4_Snowshoeing,
	"Snowmobiling":          SportBits4_Snowmobiling,
	"StandUpPaddleboarding": SportBits4_StandUpPaddleboarding,
	"Surfing":               SportBits4_Surfing,
	"Wakeboarding":          SportBits4_Wakeboarding,
}

func SportBits4FromString(s string) SportBits4 {
	if v, ok := SportBits4Values[s]; ok {
		return v
	}
	return SportBits4(SportBits4_Invalid)
}
