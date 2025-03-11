package internal

import (
	"fmt"
	"io"
)

type ForLanguage int

const (
	LanguageUnknown ForLanguage = iota
	LanguageGo
)

type ConstType struct {
	fileInformation
	options      IOptions
	name         string
	buildInType  bool
	size         int
	unsignedFlag bool
	isLong       bool
}

func (ct *ConstType) ExportName() string {
	return ct.calculateValue()
}

func (ct *ConstType) Name() string {
	return ct.ExportName()
}

func (ct *ConstType) IncludeImports(m map[string]string) {
}

func (ct *ConstType) FileName() string {
	return ct.fileName
}

func (ct *ConstType) LineNumber() int {
	return ct.lineNumber
}

func (ct *ConstType) RowNumber() int {
	return ct.rowNumber
}

func (ct *ConstType) calculateValue() string {
	switch ct.options.GetLanguage() {
	case LanguageGo:
		switch ct.buildInType {
		case true:
			switch ct.name {
			case "uint64":
				return "uint64"
			case "int64":
				return "int64"
			case "int32":
				return "int32"
			case "int16":
				return "int16"
			case "int8":
				return "int8"
			case "uint32":
				return "uint32"
			case "uint16":
				return "uint16"
			case "uint8":
				return "uint8"
			case "short":
				if ct.unsignedFlag {
					return "uint16"
				} else {
					return "int16"
				}
			case "long":
				if ct.unsignedFlag && ct.isLong {
					return "uint64"
				} else if ct.isLong {
					return "int64"
				} else if ct.unsignedFlag {
					return "uint32"
				} else {
					return "int32"
				}
			case "double":
				return "float64"
			case "float":
				return "float32"
			case "string":
				return "string"
			case "wstring":
				return "string"
			default:
				panic(ct.name + " is undefined")
			}
		default:
			return ct.name
		}
	default:
		return ct.name
	}

}

func (ct *ConstType) BuildType(w io.Writer, params BuildTypeParams) (n int, err error) {
	value := ct.calculateValue()
	if params.enforceName {
		_, _ = io.WriteString(w, fmt.Sprintf("type %v %v\n", params.enforcedExportName, value))
	} else {
		_, _ = io.WriteString(w, value)
	}
	return 0, nil
}

func (ct *ConstType) Size() int {
	return ct.size
}

func (ct *ConstType) UnsignedFlag() bool {
	return ct.unsignedFlag
}

func (ct *ConstType) IsLong() bool {
	return ct.isLong
}
