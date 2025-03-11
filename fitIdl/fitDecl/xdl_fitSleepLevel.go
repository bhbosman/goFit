package fitDecl

import strconv "strconv"

type SleepLevel byte

const (
	SleepLevel_Unmeasurable SleepLevel = 0
	SleepLevel_Awake        SleepLevel = 1
	SleepLevel_Light        SleepLevel = 2
	SleepLevel_Deep         SleepLevel = 3
	SleepLevel_Rem          SleepLevel = 4
	SleepLevel_Invalid      SleepLevel = 255
)

var SleepLevelNames = map[SleepLevel]string{
	SleepLevel_Unmeasurable: "Unmeasurable",
	SleepLevel_Awake:        "Awake",
	SleepLevel_Light:        "Light",
	SleepLevel_Deep:         "Deep",
	SleepLevel_Rem:          "Rem",
}

func (k SleepLevel) String() string {
	if uint(k) < uint(len(SleepLevelNames)) {
		if v, ok := SleepLevelNames[k]; ok {
			return v
		}
	}
	return "SleepLevel(" + strconv.Itoa(int(k)) + ")"
}

var SleepLevelValues = map[string]SleepLevel{
	"Unmeasurable": SleepLevel_Unmeasurable,
	"Awake":        SleepLevel_Awake,
	"Light":        SleepLevel_Light,
	"Deep":         SleepLevel_Deep,
	"Rem":          SleepLevel_Rem,
}

func SleepLevelFromString(s string) SleepLevel {
	if v, ok := SleepLevelValues[s]; ok {
		return v
	}
	return SleepLevel(SleepLevel_Invalid)
}
