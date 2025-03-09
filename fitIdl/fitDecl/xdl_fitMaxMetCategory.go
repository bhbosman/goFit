package fitDecl

import strconv "strconv"

type MaxMetCategory byte

const (
	MaxMetCategory_Generic MaxMetCategory = 0
	MaxMetCategory_Cycling MaxMetCategory = 1
	MaxMetCategory_Invalid MaxMetCategory = 255
)

var MaxMetCategoryNames = map[MaxMetCategory]string{
	MaxMetCategory_Generic: "Generic",
	MaxMetCategory_Cycling: "Cycling",
}

func (k MaxMetCategory) String() string {
	if uint(k) < uint(len(MaxMetCategoryNames)) {
		if v, ok := MaxMetCategoryNames[k]; ok {
			return v
		}
	}
	return "MaxMetCategory(" + strconv.Itoa(int(k)) + ")"
}

var MaxMetCategoryValues = map[string]MaxMetCategory{
	"Generic": MaxMetCategory_Generic,
	"Cycling": MaxMetCategory_Cycling,
}

func MaxMetCategoryFromString(s string) MaxMetCategory {
	if v, ok := MaxMetCategoryValues[s]; ok {
		return v
	}
	return MaxMetCategory(MaxMetCategory_Invalid)
}
