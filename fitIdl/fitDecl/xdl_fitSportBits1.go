package fitDecl

import strconv "strconv"

type SportBits1 uint8

const (
	SportBits1_Tennis             SportBits1 = 1
	SportBits1_AmericanFootball   SportBits1 = 2
	SportBits1_Training           SportBits1 = 4
	SportBits1_Walking            SportBits1 = 8
	SportBits1_CrossCountrySkiing SportBits1 = 16
	SportBits1_AlpineSkiing       SportBits1 = 32
	SportBits1_Snowboarding       SportBits1 = 64
	SportBits1_Rowing             SportBits1 = 128
	SportBits1_Invalid            SportBits1 = 0
)

var SportBits1Names = map[SportBits1]string{
	SportBits1_Tennis:             "Tennis",
	SportBits1_AmericanFootball:   "AmericanFootball",
	SportBits1_Training:           "Training",
	SportBits1_Walking:            "Walking",
	SportBits1_CrossCountrySkiing: "CrossCountrySkiing",
	SportBits1_AlpineSkiing:       "AlpineSkiing",
	SportBits1_Snowboarding:       "Snowboarding",
	SportBits1_Rowing:             "Rowing",
}

func (k SportBits1) String() string {
	if uint(k) < uint(len(SportBits1Names)) {
		if v, ok := SportBits1Names[k]; ok {
			return v
		}
	}
	return "SportBits1(" + strconv.Itoa(int(k)) + ")"
}

var SportBits1Values = map[string]SportBits1{
	"Tennis":             SportBits1_Tennis,
	"AmericanFootball":   SportBits1_AmericanFootball,
	"Training":           SportBits1_Training,
	"Walking":            SportBits1_Walking,
	"CrossCountrySkiing": SportBits1_CrossCountrySkiing,
	"AlpineSkiing":       SportBits1_AlpineSkiing,
	"Snowboarding":       SportBits1_Snowboarding,
	"Rowing":             SportBits1_Rowing,
}

func SportBits1FromString(s string) SportBits1 {
	if v, ok := SportBits1Values[s]; ok {
		return v
	}
	return SportBits1(SportBits1_Invalid)
}
