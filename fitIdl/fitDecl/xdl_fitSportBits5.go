package fitDecl

import strconv "strconv"

type SportBits5 uint8

const (
	SportBits5_WaterSkiing SportBits5 = 1
	SportBits5_Kayaking    SportBits5 = 2
	SportBits5_Rafting     SportBits5 = 4
	SportBits5_Windsurfing SportBits5 = 8
	SportBits5_Kitesurfing SportBits5 = 16
	SportBits5_Tactical    SportBits5 = 32
	SportBits5_Jumpmaster  SportBits5 = 64
	SportBits5_Boxing      SportBits5 = 128
	SportBits5_Invalid     SportBits5 = 0
)

var SportBits5Names = map[SportBits5]string{
	SportBits5_WaterSkiing: "WaterSkiing",
	SportBits5_Kayaking:    "Kayaking",
	SportBits5_Rafting:     "Rafting",
	SportBits5_Windsurfing: "Windsurfing",
	SportBits5_Kitesurfing: "Kitesurfing",
	SportBits5_Tactical:    "Tactical",
	SportBits5_Jumpmaster:  "Jumpmaster",
	SportBits5_Boxing:      "Boxing",
}

func (k SportBits5) String() string {
	if uint(k) < uint(len(SportBits5Names)) {
		if v, ok := SportBits5Names[k]; ok {
			return v
		}
	}
	return "SportBits5(" + strconv.Itoa(int(k)) + ")"
}

var SportBits5Values = map[string]SportBits5{
	"WaterSkiing": SportBits5_WaterSkiing,
	"Kayaking":    SportBits5_Kayaking,
	"Rafting":     SportBits5_Rafting,
	"Windsurfing": SportBits5_Windsurfing,
	"Kitesurfing": SportBits5_Kitesurfing,
	"Tactical":    SportBits5_Tactical,
	"Jumpmaster":  SportBits5_Jumpmaster,
	"Boxing":      SportBits5_Boxing,
}

func SportBits5FromString(s string) SportBits5 {
	if v, ok := SportBits5Values[s]; ok {
		return v
	}
	return SportBits5(SportBits5_Invalid)
}
