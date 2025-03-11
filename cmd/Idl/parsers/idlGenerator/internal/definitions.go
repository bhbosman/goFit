package internal

import (
	"io"
)

type DefinitionType int

const (
	DtNotDefined DefinitionType = iota
	DtConstant
	DtStruct
	DtEnum
	DtTypedef
	DtDefineExpression
)

type IFileInformation interface {
	FileName() string
	LineNumber() int
	RowNumber() int
}

type BuildTypeParamsWhere int

const (
	BtNoDefined BuildTypeParamsWhere = iota
	BtMemberTypeDecl
)

type BuildTypeParams struct {
	enforceName        bool
	enforcedExportName string
	forWhere           BuildTypeParamsWhere
}

type ITypeSpec interface {
	IFileInformation
	IncludeImports(m map[string]string)
	BuildType(w io.Writer, params BuildTypeParams) (n int, err error)
}

type IBaseDefinition interface {
	ITypeSpec
	DefinitionType() DefinitionType
}
type IDefinition interface {
	IBaseDefinition
	Name() string
	ExportName() string
}

type fileInformation struct {
	fileName   string
	lineNumber int
	rowNumber  int
}

type IMultiItems interface {
	Count() int
	Item(i int) (ITypeSpec, string)
}
