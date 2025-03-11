package fitDecl

import strconv "strconv"

type LanguageBits0 uint8

const (
	LanguageBits0_English  LanguageBits0 = 1
	LanguageBits0_French   LanguageBits0 = 2
	LanguageBits0_Italian  LanguageBits0 = 4
	LanguageBits0_German   LanguageBits0 = 8
	LanguageBits0_Spanish  LanguageBits0 = 16
	LanguageBits0_Croatian LanguageBits0 = 32
	LanguageBits0_Czech    LanguageBits0 = 64
	LanguageBits0_Danish   LanguageBits0 = 128
	LanguageBits0_Invalid  LanguageBits0 = 0
)

var LanguageBits0Names = map[LanguageBits0]string{
	LanguageBits0_English:  "English",
	LanguageBits0_French:   "French",
	LanguageBits0_Italian:  "Italian",
	LanguageBits0_German:   "German",
	LanguageBits0_Spanish:  "Spanish",
	LanguageBits0_Croatian: "Croatian",
	LanguageBits0_Czech:    "Czech",
	LanguageBits0_Danish:   "Danish",
}

func (k LanguageBits0) String() string {
	if uint(k) < uint(len(LanguageBits0Names)) {
		if v, ok := LanguageBits0Names[k]; ok {
			return v
		}
	}
	return "LanguageBits0(" + strconv.Itoa(int(k)) + ")"
}

var LanguageBits0Values = map[string]LanguageBits0{
	"English":  LanguageBits0_English,
	"French":   LanguageBits0_French,
	"Italian":  LanguageBits0_Italian,
	"German":   LanguageBits0_German,
	"Spanish":  LanguageBits0_Spanish,
	"Croatian": LanguageBits0_Croatian,
	"Czech":    LanguageBits0_Czech,
	"Danish":   LanguageBits0_Danish,
}

func LanguageBits0FromString(s string) LanguageBits0 {
	if v, ok := LanguageBits0Values[s]; ok {
		return v
	}
	return LanguageBits0(LanguageBits0_Invalid)
}
