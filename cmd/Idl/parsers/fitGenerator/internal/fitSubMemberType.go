package internal

import (
	"io"
)

type SubMemberType struct {
	fileInformation
	typeSpec ITypeSpec
}

func (subMemberType *SubMemberType) FileName() string {
	return subMemberType.fileName
}

func (subMemberType *SubMemberType) LineNumber() int {
	return subMemberType.lineNumber
}

func (subMemberType *SubMemberType) RowNumber() int {
	return subMemberType.rowNumber
}

func (subMemberType *SubMemberType) IncludeImports(m map[string]string) {
}

func (subMemberType *SubMemberType) BuildType(w io.Writer, params BuildTypeParams) (n int, err error) {
	return subMemberType.typeSpec.BuildType(w, params)
}
