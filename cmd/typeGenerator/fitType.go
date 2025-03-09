package main

type FitType string

func (ct FitType) InvalidValue() int {
	switch ct {
	case ctenum:
		return int(byte(0xFF))
	case ctuint8:

		return int(uint8(0xFF))
	case ctuint8z:
		return 0x00
	case ctuint16:
		return int(uint16(0xFFFF))
	case ctuint16z:
		return 0x00
	case ctuint32:
		return int(uint32(0xFFFFFFFF))
	case ctuint32z:
		return 0x00
	default:
		panic("invalid constant type")
	}
}

func StringToConstantType(s string) FitType {
	switch s {
	case "enum":
		return ctenum
	case "uint8":
		return ctuint8
	case "uint8z":
		return ctuint8z
	case "uint16":
		return ctuint16
	case "uint16z":
		return ctuint16z
	case "uint32":
		return ctuint32
	case "uint32z":
		return ctuint32z
	default:
		panic("invalid constant type")
	}
}

func (ct FitType) String() string {
	switch ct {
	case ctenum:
		return "byte"
	case ctuint8:
		return "uint8"
	case ctuint8z:
		return "uint8"
	case ctuint16:
		return "uint16"
	case ctuint16z:
		return "uint16"
	case ctuint32:
		return "uint32"
	case ctuint32z:
		return "uint32"
	}
	return "(Unknown)"
}

const (
	ctenum    FitType = "enum"
	ctuint8           = "uint8"
	ctuint8z          = "uint8z"
	ctuint16          = "uint16"
	ctuint16z         = "uint16z"
	ctuint32          = "uint32"
	ctuint32z         = "uint32z"
)
