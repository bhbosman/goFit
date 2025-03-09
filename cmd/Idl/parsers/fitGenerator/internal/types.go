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
	return ct.name
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
