package internal

import (
	"fmt"
	"go/ast"
	"io"
	"strings"
)

// RwConstant
// Must implement the following interfaces:
// Must implement the following interfaces:
//  1. IDefinition
//  2. IFileInformation
//  3. ITypeSpec
type RwConstant struct {
	fileInformation
	name     string
	typeName ITypeSpec
	value    ast.Expr
}

func (rwConstant *RwConstant) BuildType(w io.Writer, _ BuildTypeParams) (n int, err error) {
	_, _ = io.WriteString(w, fmt.Sprintf("const %v ", rwConstant.ExportName()))
	_, _ = rwConstant.typeName.BuildType(w, BuildTypeParams{false, "", BtMemberTypeDecl})
	_, _ = io.WriteString(w, fmt.Sprintf(" = %v\n", walkExpressionTree(rwConstant.value)))
	return 0, nil
}

func (rwConstant *RwConstant) IncludeImports(map[string]string) {
}

func (rwConstant *RwConstant) FileName() string {
	return rwConstant.fileName
}

func (rwConstant *RwConstant) LineNumber() int {
	return rwConstant.lineNumber
}

func (rwConstant *RwConstant) RowNumber() int {
	return rwConstant.rowNumber
}

func (rwConstant *RwConstant) Name() string {
	return rwConstant.ExportName()
}

func (rwConstant *RwConstant) ExportName() string {
	return __exportName(rwConstant.name)
}

func (rwConstant *RwConstant) DefinitionType() DefinitionType {
	return DtConstant
}

func __exportName(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

type HashDefineExpression struct {
	fileInformation
	name  string
	value ast.Expr
}

func (hashDefine *HashDefineExpression) FileName() string {
	return hashDefine.fileName
}

func (hashDefine *HashDefineExpression) LineNumber() int {
	return hashDefine.lineNumber
}

func (hashDefine *HashDefineExpression) RowNumber() int {
	return hashDefine.rowNumber
}

func (hashDefine *HashDefineExpression) IncludeImports(m map[string]string) {
}

func (hashDefine *HashDefineExpression) BuildType(w io.Writer, params BuildTypeParams) (n int, err error) {
	return 0, err
}

func (hashDefine *HashDefineExpression) DefinitionType() DefinitionType {
	return DtDefineExpression

}
