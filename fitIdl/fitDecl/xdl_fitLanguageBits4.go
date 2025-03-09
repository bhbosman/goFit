package fitDecl

import strconv "strconv"

type LanguageBits4 uint8

const (
	LanguageBits4_BrazilianPortuguese LanguageBits4 = 1
	LanguageBits4_Indonesian          LanguageBits4 = 2
	LanguageBits4_Malaysian           LanguageBits4 = 4
	LanguageBits4_Vietnamese          LanguageBits4 = 8
	LanguageBits4_Burmese             LanguageBits4 = 16
	LanguageBits4_Mongolian           LanguageBits4 = 32
	LanguageBits4_Invalid             LanguageBits4 = 0
)

var LanguageBits4Names = map[LanguageBits4]string{
	LanguageBits4_BrazilianPortuguese: "BrazilianPortuguese",
	LanguageBits4_Indonesian:          "Indonesian",
	LanguageBits4_Malaysian:           "Malaysian",
	LanguageBits4_Vietnamese:          "Vietnamese",
	LanguageBits4_Burmese:             "Burmese",
	LanguageBits4_Mongolian:           "Mongolian",
}

func (k LanguageBits4) String() string {
	if uint(k) < uint(len(LanguageBits4Names)) {
		if v, ok := LanguageBits4Names[k]; ok {
			return v
		}
	}
	return "LanguageBits4(" + strconv.Itoa(int(k)) + ")"
}

var LanguageBits4Values = map[string]LanguageBits4{
	"BrazilianPortuguese": LanguageBits4_BrazilianPortuguese,
	"Indonesian":          LanguageBits4_Indonesian,
	"Malaysian":           LanguageBits4_Malaysian,
	"Vietnamese":          LanguageBits4_Vietnamese,
	"Burmese":             LanguageBits4_Burmese,
	"Mongolian":           LanguageBits4_Mongolian,
}

func LanguageBits4FromString(s string) LanguageBits4 {
	if v, ok := LanguageBits4Values[s]; ok {
		return v
	}
	return LanguageBits4(LanguageBits4_Invalid)
}
