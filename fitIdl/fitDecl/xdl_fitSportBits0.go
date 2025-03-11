package fitDecl

import strconv "strconv"

type SportBits0 uint8

const (
	SportBits0_Generic          SportBits0 = 1
	SportBits0_Running          SportBits0 = 2
	SportBits0_Cycling          SportBits0 = 4
	SportBits0_Transition       SportBits0 = 8
	SportBits0_FitnessEquipment SportBits0 = 16
	SportBits0_Swimming         SportBits0 = 32
	SportBits0_Basketball       SportBits0 = 64
	SportBits0_Soccer           SportBits0 = 128
	SportBits0_Invalid          SportBits0 = 0
)

var SportBits0Names = map[SportBits0]string{
	SportBits0_Generic:          "Generic",
	SportBits0_Running:          "Running",
	SportBits0_Cycling:          "Cycling",
	SportBits0_Transition:       "Transition",
	SportBits0_FitnessEquipment: "FitnessEquipment",
	SportBits0_Swimming:         "Swimming",
	SportBits0_Basketball:       "Basketball",
	SportBits0_Soccer:           "Soccer",
}

func (k SportBits0) String() string {
	if uint(k) < uint(len(SportBits0Names)) {
		if v, ok := SportBits0Names[k]; ok {
			return v
		}
	}
	return "SportBits0(" + strconv.Itoa(int(k)) + ")"
}

var SportBits0Values = map[string]SportBits0{
	"Generic":          SportBits0_Generic,
	"Running":          SportBits0_Running,
	"Cycling":          SportBits0_Cycling,
	"Transition":       SportBits0_Transition,
	"FitnessEquipment": SportBits0_FitnessEquipment,
	"Swimming":         SportBits0_Swimming,
	"Basketball":       SportBits0_Basketball,
	"Soccer":           SportBits0_Soccer,
}

func SportBits0FromString(s string) SportBits0 {
	if v, ok := SportBits0Values[s]; ok {
		return v
	}
	return SportBits0(SportBits0_Invalid)
}
