package fitDecl

import strconv "strconv"

type FitBaseType uint8

const (
	FitBaseType_Enum    FitBaseType = 0
	FitBaseType_Sint8   FitBaseType = 1
	FitBaseType_Uint8   FitBaseType = 2
	FitBaseType_Sint16  FitBaseType = 131
	FitBaseType_Uint16  FitBaseType = 132
	FitBaseType_Sint32  FitBaseType = 133
	FitBaseType_Uint32  FitBaseType = 134
	FitBaseType_String  FitBaseType = 7
	FitBaseType_Float32 FitBaseType = 136
	FitBaseType_Float64 FitBaseType = 137
	FitBaseType_Uint8z  FitBaseType = 10
	FitBaseType_Uint16z FitBaseType = 139
	FitBaseType_Uint32z FitBaseType = 140
	FitBaseType_Byte    FitBaseType = 13
	FitBaseType_Sint64  FitBaseType = 142
	FitBaseType_Uint64  FitBaseType = 143
	FitBaseType_Uint64z FitBaseType = 144
	FitBaseType_Invalid FitBaseType = 255
)

var FitBaseTypeNames = map[FitBaseType]string{
	FitBaseType_Enum:    "Enum",
	FitBaseType_Sint8:   "Sint8",
	FitBaseType_Uint8:   "Uint8",
	FitBaseType_Sint16:  "Sint16",
	FitBaseType_Uint16:  "Uint16",
	FitBaseType_Sint32:  "Sint32",
	FitBaseType_Uint32:  "Uint32",
	FitBaseType_String:  "String",
	FitBaseType_Float32: "Float32",
	FitBaseType_Float64: "Float64",
	FitBaseType_Uint8z:  "Uint8z",
	FitBaseType_Uint16z: "Uint16z",
	FitBaseType_Uint32z: "Uint32z",
	FitBaseType_Byte:    "Byte",
	FitBaseType_Sint64:  "Sint64",
	FitBaseType_Uint64:  "Uint64",
	FitBaseType_Uint64z: "Uint64z",
}

func (k FitBaseType) String() string {
	if uint(k) < uint(len(FitBaseTypeNames)) {
		if v, ok := FitBaseTypeNames[k]; ok {
			return v
		}
	}
	return "FitBaseType(" + strconv.Itoa(int(k)) + ")"
}

var FitBaseTypeValues = map[string]FitBaseType{
	"Enum":    FitBaseType_Enum,
	"Sint8":   FitBaseType_Sint8,
	"Uint8":   FitBaseType_Uint8,
	"Sint16":  FitBaseType_Sint16,
	"Uint16":  FitBaseType_Uint16,
	"Sint32":  FitBaseType_Sint32,
	"Uint32":  FitBaseType_Uint32,
	"String":  FitBaseType_String,
	"Float32": FitBaseType_Float32,
	"Float64": FitBaseType_Float64,
	"Uint8z":  FitBaseType_Uint8z,
	"Uint16z": FitBaseType_Uint16z,
	"Uint32z": FitBaseType_Uint32z,
	"Byte":    FitBaseType_Byte,
	"Sint64":  FitBaseType_Sint64,
	"Uint64":  FitBaseType_Uint64,
	"Uint64z": FitBaseType_Uint64z,
}

func FitBaseTypeFromString(s string) FitBaseType {
	if v, ok := FitBaseTypeValues[s]; ok {
		return v
	}
	return FitBaseType(FitBaseType_Invalid)
}
