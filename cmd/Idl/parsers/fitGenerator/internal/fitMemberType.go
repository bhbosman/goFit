package internal

import (
	"go/ast"
	"io"
)

type MemberType struct {
	fileInformation
	typeSpec      ITypeSpec
	baseType      ast.Expr
	defNumber     ast.Expr
	arrayType     ast.Expr
	components    ast.Expr
	scale         ast.Expr
	offset        ast.Expr
	units         ast.Expr
	bits          ast.Expr
	accumulate    ast.Expr
	refFieldName  ast.Expr
	refFieldValue ast.Expr
	comment       ast.Expr
	products      ast.Expr
	example       ast.Expr
}

func (memberType *MemberType) FileName() string {
	return memberType.fileName
}

func (memberType *MemberType) LineNumber() int {
	return memberType.lineNumber
}

func (memberType *MemberType) RowNumber() int {
	return memberType.rowNumber
}

func (memberType *MemberType) IncludeImports(m map[string]string) {
}

func (memberType *MemberType) BuildType(w io.Writer, params BuildTypeParams) (n int, err error) {
	return memberType.typeSpec.BuildType(w, params)
}
