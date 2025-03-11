package fitDecl

import strconv "strconv"

type LanguageBits3 uint8

const (
	LanguageBits3_Bulgarian LanguageBits3 = 1
	LanguageBits3_Romanian  LanguageBits3 = 2
	LanguageBits3_Chinese   LanguageBits3 = 4
	LanguageBits3_Japanese  LanguageBits3 = 8
	LanguageBits3_Korean    LanguageBits3 = 16
	LanguageBits3_Taiwanese LanguageBits3 = 32
	LanguageBits3_Thai      LanguageBits3 = 64
	LanguageBits3_Hebrew    LanguageBits3 = 128
	LanguageBits3_Invalid   LanguageBits3 = 0
)

var LanguageBits3Names = map[LanguageBits3]string{
	LanguageBits3_Bulgarian: "Bulgarian",
	LanguageBits3_Romanian:  "Romanian",
	LanguageBits3_Chinese:   "Chinese",
	LanguageBits3_Japanese:  "Japanese",
	LanguageBits3_Korean:    "Korean",
	LanguageBits3_Taiwanese: "Taiwanese",
	LanguageBits3_Thai:      "Thai",
	LanguageBits3_Hebrew:    "Hebrew",
}

func (k LanguageBits3) String() string {
	if uint(k) < uint(len(LanguageBits3Names)) {
		if v, ok := LanguageBits3Names[k]; ok {
			return v
		}
	}
	return "LanguageBits3(" + strconv.Itoa(int(k)) + ")"
}

var LanguageBits3Values = map[string]LanguageBits3{
	"Bulgarian": LanguageBits3_Bulgarian,
	"Romanian":  LanguageBits3_Romanian,
	"Chinese":   LanguageBits3_Chinese,
	"Japanese":  LanguageBits3_Japanese,
	"Korean":    LanguageBits3_Korean,
	"Taiwanese": LanguageBits3_Taiwanese,
	"Thai":      LanguageBits3_Thai,
	"Hebrew":    LanguageBits3_Hebrew,
}

func LanguageBits3FromString(s string) LanguageBits3 {
	if v, ok := LanguageBits3Values[s]; ok {
		return v
	}
	return LanguageBits3(LanguageBits3_Invalid)
}
