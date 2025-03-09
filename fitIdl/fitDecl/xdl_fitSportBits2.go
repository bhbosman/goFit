package fitDecl

import strconv "strconv"

type SportBits2 uint8

const (
	SportBits2_Mountaineering SportBits2 = 1
	SportBits2_Hiking         SportBits2 = 2
	SportBits2_Multisport     SportBits2 = 4
	SportBits2_Paddling       SportBits2 = 8
	SportBits2_Flying         SportBits2 = 16
	SportBits2_EBiking        SportBits2 = 32
	SportBits2_Motorcycling   SportBits2 = 64
	SportBits2_Boating        SportBits2 = 128
	SportBits2_Invalid        SportBits2 = 0
)

var SportBits2Names = map[SportBits2]string{
	SportBits2_Mountaineering: "Mountaineering",
	SportBits2_Hiking:         "Hiking",
	SportBits2_Multisport:     "Multisport",
	SportBits2_Paddling:       "Paddling",
	SportBits2_Flying:         "Flying",
	SportBits2_EBiking:        "EBiking",
	SportBits2_Motorcycling:   "Motorcycling",
	SportBits2_Boating:        "Boating",
}

func (k SportBits2) String() string {
	if uint(k) < uint(len(SportBits2Names)) {
		if v, ok := SportBits2Names[k]; ok {
			return v
		}
	}
	return "SportBits2(" + strconv.Itoa(int(k)) + ")"
}

var SportBits2Values = map[string]SportBits2{
	"Mountaineering": SportBits2_Mountaineering,
	"Hiking":         SportBits2_Hiking,
	"Multisport":     SportBits2_Multisport,
	"Paddling":       SportBits2_Paddling,
	"Flying":         SportBits2_Flying,
	"EBiking":        SportBits2_EBiking,
	"Motorcycling":   SportBits2_Motorcycling,
	"Boating":        SportBits2_Boating,
}

func SportBits2FromString(s string) SportBits2 {
	if v, ok := SportBits2Values[s]; ok {
		return v
	}
	return SportBits2(SportBits2_Invalid)
}
