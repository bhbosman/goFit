package fitDecl

import strconv "strconv"

type Language byte

const (
	Language_English             Language = 0
	Language_French              Language = 1
	Language_Italian             Language = 2
	Language_German              Language = 3
	Language_Spanish             Language = 4
	Language_Croatian            Language = 5
	Language_Czech               Language = 6
	Language_Danish              Language = 7
	Language_Dutch               Language = 8
	Language_Finnish             Language = 9
	Language_Greek               Language = 10
	Language_Hungarian           Language = 11
	Language_Norwegian           Language = 12
	Language_Polish              Language = 13
	Language_Portuguese          Language = 14
	Language_Slovakian           Language = 15
	Language_Slovenian           Language = 16
	Language_Swedish             Language = 17
	Language_Russian             Language = 18
	Language_Turkish             Language = 19
	Language_Latvian             Language = 20
	Language_Ukrainian           Language = 21
	Language_Arabic              Language = 22
	Language_Farsi               Language = 23
	Language_Bulgarian           Language = 24
	Language_Romanian            Language = 25
	Language_Chinese             Language = 26
	Language_Japanese            Language = 27
	Language_Korean              Language = 28
	Language_Taiwanese           Language = 29
	Language_Thai                Language = 30
	Language_Hebrew              Language = 31
	Language_BrazilianPortuguese Language = 32
	Language_Indonesian          Language = 33
	Language_Malaysian           Language = 34
	Language_Vietnamese          Language = 35
	Language_Burmese             Language = 36
	Language_Mongolian           Language = 37
	Language_Custom              Language = 254
	Language_Invalid             Language = 255
)

var LanguageNames = map[Language]string{
	Language_English:             "English",
	Language_French:              "French",
	Language_Italian:             "Italian",
	Language_German:              "German",
	Language_Spanish:             "Spanish",
	Language_Croatian:            "Croatian",
	Language_Czech:               "Czech",
	Language_Danish:              "Danish",
	Language_Dutch:               "Dutch",
	Language_Finnish:             "Finnish",
	Language_Greek:               "Greek",
	Language_Hungarian:           "Hungarian",
	Language_Norwegian:           "Norwegian",
	Language_Polish:              "Polish",
	Language_Portuguese:          "Portuguese",
	Language_Slovakian:           "Slovakian",
	Language_Slovenian:           "Slovenian",
	Language_Swedish:             "Swedish",
	Language_Russian:             "Russian",
	Language_Turkish:             "Turkish",
	Language_Latvian:             "Latvian",
	Language_Ukrainian:           "Ukrainian",
	Language_Arabic:              "Arabic",
	Language_Farsi:               "Farsi",
	Language_Bulgarian:           "Bulgarian",
	Language_Romanian:            "Romanian",
	Language_Chinese:             "Chinese",
	Language_Japanese:            "Japanese",
	Language_Korean:              "Korean",
	Language_Taiwanese:           "Taiwanese",
	Language_Thai:                "Thai",
	Language_Hebrew:              "Hebrew",
	Language_BrazilianPortuguese: "BrazilianPortuguese",
	Language_Indonesian:          "Indonesian",
	Language_Malaysian:           "Malaysian",
	Language_Vietnamese:          "Vietnamese",
	Language_Burmese:             "Burmese",
	Language_Mongolian:           "Mongolian",
	Language_Custom:              "Custom",
}

func (k Language) String() string {
	if uint(k) < uint(len(LanguageNames)) {
		if v, ok := LanguageNames[k]; ok {
			return v
		}
	}
	return "Language(" + strconv.Itoa(int(k)) + ")"
}

var LanguageValues = map[string]Language{
	"English":             Language_English,
	"French":              Language_French,
	"Italian":             Language_Italian,
	"German":              Language_German,
	"Spanish":             Language_Spanish,
	"Croatian":            Language_Croatian,
	"Czech":               Language_Czech,
	"Danish":              Language_Danish,
	"Dutch":               Language_Dutch,
	"Finnish":             Language_Finnish,
	"Greek":               Language_Greek,
	"Hungarian":           Language_Hungarian,
	"Norwegian":           Language_Norwegian,
	"Polish":              Language_Polish,
	"Portuguese":          Language_Portuguese,
	"Slovakian":           Language_Slovakian,
	"Slovenian":           Language_Slovenian,
	"Swedish":             Language_Swedish,
	"Russian":             Language_Russian,
	"Turkish":             Language_Turkish,
	"Latvian":             Language_Latvian,
	"Ukrainian":           Language_Ukrainian,
	"Arabic":              Language_Arabic,
	"Farsi":               Language_Farsi,
	"Bulgarian":           Language_Bulgarian,
	"Romanian":            Language_Romanian,
	"Chinese":             Language_Chinese,
	"Japanese":            Language_Japanese,
	"Korean":              Language_Korean,
	"Taiwanese":           Language_Taiwanese,
	"Thai":                Language_Thai,
	"Hebrew":              Language_Hebrew,
	"BrazilianPortuguese": Language_BrazilianPortuguese,
	"Indonesian":          Language_Indonesian,
	"Malaysian":           Language_Malaysian,
	"Vietnamese":          Language_Vietnamese,
	"Burmese":             Language_Burmese,
	"Mongolian":           Language_Mongolian,
	"Custom":              Language_Custom,
}

func LanguageFromString(s string) Language {
	if v, ok := LanguageValues[s]; ok {
		return v
	}
	return Language(Language_Invalid)
}
