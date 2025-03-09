package fitDecl

import strconv "strconv"

type LanguageBits2 uint8

const (
	LanguageBits2_Slovenian LanguageBits2 = 1
	LanguageBits2_Swedish   LanguageBits2 = 2
	LanguageBits2_Russian   LanguageBits2 = 4
	LanguageBits2_Turkish   LanguageBits2 = 8
	LanguageBits2_Latvian   LanguageBits2 = 16
	LanguageBits2_Ukrainian LanguageBits2 = 32
	LanguageBits2_Arabic    LanguageBits2 = 64
	LanguageBits2_Farsi     LanguageBits2 = 128
	LanguageBits2_Invalid   LanguageBits2 = 0
)

var LanguageBits2Names = map[LanguageBits2]string{
	LanguageBits2_Slovenian: "Slovenian",
	LanguageBits2_Swedish:   "Swedish",
	LanguageBits2_Russian:   "Russian",
	LanguageBits2_Turkish:   "Turkish",
	LanguageBits2_Latvian:   "Latvian",
	LanguageBits2_Ukrainian: "Ukrainian",
	LanguageBits2_Arabic:    "Arabic",
	LanguageBits2_Farsi:     "Farsi",
}

func (k LanguageBits2) String() string {
	if uint(k) < uint(len(LanguageBits2Names)) {
		if v, ok := LanguageBits2Names[k]; ok {
			return v
		}
	}
	return "LanguageBits2(" + strconv.Itoa(int(k)) + ")"
}

var LanguageBits2Values = map[string]LanguageBits2{
	"Slovenian": LanguageBits2_Slovenian,
	"Swedish":   LanguageBits2_Swedish,
	"Russian":   LanguageBits2_Russian,
	"Turkish":   LanguageBits2_Turkish,
	"Latvian":   LanguageBits2_Latvian,
	"Ukrainian": LanguageBits2_Ukrainian,
	"Arabic":    LanguageBits2_Arabic,
	"Farsi":     LanguageBits2_Farsi,
}

func LanguageBits2FromString(s string) LanguageBits2 {
	if v, ok := LanguageBits2Values[s]; ok {
		return v
	}
	return LanguageBits2(LanguageBits2_Invalid)
}
