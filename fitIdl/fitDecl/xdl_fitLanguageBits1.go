package fitDecl

import strconv "strconv"

type LanguageBits1 uint8

const (
	LanguageBits1_Dutch      LanguageBits1 = 1
	LanguageBits1_Finnish    LanguageBits1 = 2
	LanguageBits1_Greek      LanguageBits1 = 4
	LanguageBits1_Hungarian  LanguageBits1 = 8
	LanguageBits1_Norwegian  LanguageBits1 = 16
	LanguageBits1_Polish     LanguageBits1 = 32
	LanguageBits1_Portuguese LanguageBits1 = 64
	LanguageBits1_Slovakian  LanguageBits1 = 128
	LanguageBits1_Invalid    LanguageBits1 = 0
)

var LanguageBits1Names = map[LanguageBits1]string{
	LanguageBits1_Dutch:      "Dutch",
	LanguageBits1_Finnish:    "Finnish",
	LanguageBits1_Greek:      "Greek",
	LanguageBits1_Hungarian:  "Hungarian",
	LanguageBits1_Norwegian:  "Norwegian",
	LanguageBits1_Polish:     "Polish",
	LanguageBits1_Portuguese: "Portuguese",
	LanguageBits1_Slovakian:  "Slovakian",
}

func (k LanguageBits1) String() string {
	if uint(k) < uint(len(LanguageBits1Names)) {
		if v, ok := LanguageBits1Names[k]; ok {
			return v
		}
	}
	return "LanguageBits1(" + strconv.Itoa(int(k)) + ")"
}

var LanguageBits1Values = map[string]LanguageBits1{
	"Dutch":      LanguageBits1_Dutch,
	"Finnish":    LanguageBits1_Finnish,
	"Greek":      LanguageBits1_Greek,
	"Hungarian":  LanguageBits1_Hungarian,
	"Norwegian":  LanguageBits1_Norwegian,
	"Polish":     LanguageBits1_Polish,
	"Portuguese": LanguageBits1_Portuguese,
	"Slovakian":  LanguageBits1_Slovakian,
}

func LanguageBits1FromString(s string) LanguageBits1 {
	if v, ok := LanguageBits1Values[s]; ok {
		return v
	}
	return LanguageBits1(LanguageBits1_Invalid)
}
