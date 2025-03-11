package fitDecl

import strconv "strconv"

type TissueModelType byte

const (
	TissueModelType_Zhl16c  TissueModelType = 0
	TissueModelType_Invalid TissueModelType = 255
)

var TissueModelTypeNames = map[TissueModelType]string{
	TissueModelType_Zhl16c: "Zhl16c",
}

func (k TissueModelType) String() string {
	if uint(k) < uint(len(TissueModelTypeNames)) {
		if v, ok := TissueModelTypeNames[k]; ok {
			return v
		}
	}
	return "TissueModelType(" + strconv.Itoa(int(k)) + ")"
}

var TissueModelTypeValues = map[string]TissueModelType{
	"Zhl16c": TissueModelType_Zhl16c,
}

func TissueModelTypeFromString(s string) TissueModelType {
	if v, ok := TissueModelTypeValues[s]; ok {
		return v
	}
	return TissueModelType(TissueModelType_Invalid)
}
