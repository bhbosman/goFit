package main

import "unicode"

type FitIdentifier string

func (s FitIdentifier) String() string {
	result := [128]rune{}
	resultIndex := 0
	upperCase := true
	for idx, r := range string(s) {
		if idx == 0 {
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				result[resultIndex] = '_'
				resultIndex++
			}
		}
		switch r {
		case '_':
			upperCase = true
			continue
		default:
			if upperCase {
				result[resultIndex] = unicode.ToUpper(r)
			} else {
				result[resultIndex] = r
			}
			resultIndex++
			upperCase = false
		}
	}
	return string(result[:resultIndex])
}

//func (s FitIdentifier) Original() string {
//	switch s[0] {
//	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
//		return "_" + string(s)
//	default:
//		return string(s)
//	}
//
//}

type FitTypeIdentifier string

func (s FitTypeIdentifier) String() string {
	switch s {

	case "uint8z":
		return "uint8"
	case "sint8":
		return "int8"
	case "uint8":
		return "uint8"
	case "uint16z":
		return "uint16"
	case "sint16":
		return "int16"
	case "uint16":
		return "uint16"

	case "date_time":
		return "uint32"

	case "uint32z":
		return "uint32"
	case "sint32":
		return "int32"
	case "uint32":
		return "uint32"
	case "uint64z":
		return "uint64"
	case "sint64":
		return "int64"
	case "uint64":
		return "uint64"

	case "float32":
		return "float32"
	case "float64":
		return "float64"
	case "byte":
		return "byte"
	case "string":
		return "string"
	case "bool":
		return "bool"
	case "boolean":
		return "bool"

	default:
		dd := FitIdentifier(s)
		return dd.String()

	}

}

func (s FitTypeIdentifier) BaseType() string {
	switch s {

	case "uint8z":
		return "Uint8z"
	case "sint8":
		return "Sint8"
	case "uint8":
		return "Uint8"
	case "uint16z":
		return "Uint16z"
	case "sint16":
		return "Sint16"
	case "uint16":
		return "Uint16"
	case "date_time":
		return "Uint32"

	case "uint32z":
		return "Uint32z"
	case "sint32":
		return "Sint32"
	case "uint32":
		return "Uint32"
	case "uint64z":
		return "Uint64z"
	case "sint64":
		return "Sint64"
	case "uint64":
		return "Uint64"
	case "float32":
		return "Float32"
	case "float64":
		return "Float64"
	case "byte":
		return "Byte"
	case "string":
		return "String"
	case "bool":
		return "Bool"
	case "boolean":
		return "Bool"
	default:
		return "Enum"
	}
}
