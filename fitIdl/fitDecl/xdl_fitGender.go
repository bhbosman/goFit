package fitDecl

import strconv "strconv"

type Gender byte

const (
	Gender_Female  Gender = 0
	Gender_Male    Gender = 1
	Gender_Invalid Gender = 255
)

var GenderNames = map[Gender]string{
	Gender_Female: "Female",
	Gender_Male:   "Male",
}

func (k Gender) String() string {
	if uint(k) < uint(len(GenderNames)) {
		if v, ok := GenderNames[k]; ok {
			return v
		}
	}
	return "Gender(" + strconv.Itoa(int(k)) + ")"
}

var GenderValues = map[string]Gender{
	"Female": Gender_Female,
	"Male":   Gender_Male,
}

func GenderFromString(s string) Gender {
	if v, ok := GenderValues[s]; ok {
		return v
	}
	return Gender(Gender_Invalid)
}
