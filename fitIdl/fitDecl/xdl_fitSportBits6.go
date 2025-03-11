package fitDecl

import strconv "strconv"

type SportBits6 uint8

const (
	SportBits6_FloorClimbing SportBits6 = 1
	SportBits6_Invalid       SportBits6 = 0
)

var SportBits6Names = map[SportBits6]string{
	SportBits6_FloorClimbing: "FloorClimbing",
}

func (k SportBits6) String() string {
	if uint(k) < uint(len(SportBits6Names)) {
		if v, ok := SportBits6Names[k]; ok {
			return v
		}
	}
	return "SportBits6(" + strconv.Itoa(int(k)) + ")"
}

var SportBits6Values = map[string]SportBits6{
	"FloorClimbing": SportBits6_FloorClimbing,
}

func SportBits6FromString(s string) SportBits6 {
	if v, ok := SportBits6Values[s]; ok {
		return v
	}
	return SportBits6(SportBits6_Invalid)
}
