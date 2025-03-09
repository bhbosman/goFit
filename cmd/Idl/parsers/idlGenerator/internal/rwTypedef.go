package internal

import (
	"fmt"
	"io"
)

// RwTypedef
// Must implement the following interfaces:
//  1. IBaseDefinition
//  2. IFileInformation
//  3. ITypeSpec
//  4. IMultiItems
type RwTypedef struct {
	fileInformation
	typeSpec ITypeSpec
	Items    []IdentifierType
}

func (rwTypeDef *RwTypedef) Count() int {
	return len(rwTypeDef.Items)
}

func (rwTypeDef *RwTypedef) Item(i int) (ITypeSpec, string) {
	v := &RwTypedefWithOne{rwTypeDef.typeSpec, rwTypeDef.Items[i]}
	return v, string(rwTypeDef.Items[i])
}

func (rwTypeDef *RwTypedef) BuildType(w io.Writer, _ BuildTypeParams) (n int, err error) {
	for index, item := range rwTypeDef.Items {
		_, _ = rwTypeDef.typeSpec.BuildType(
			w,
			BuildTypeParams{true, item.ExportName(), BtNoDefined})
		if index < len(rwTypeDef.Items)-1 {
			_, _ = io.WriteString(w, fmt.Sprintf("\n"))
		}
	}
	return 0, nil
}

func (rwTypeDef *RwTypedef) FileName() string {
	return rwTypeDef.fileName
}

func (rwTypeDef *RwTypedef) LineNumber() int {
	return rwTypeDef.lineNumber
}

func (rwTypeDef *RwTypedef) RowNumber() int {
	return rwTypeDef.rowNumber
}

func (rwTypeDef *RwTypedef) DefinitionType() DefinitionType {
	return DtTypedef
}

func (rwTypeDef *RwTypedef) IncludeImports(m map[string]string) {
	rwTypeDef.typeSpec.IncludeImports(m)
}

type RwTypedefWithOne struct {
	typeSpec ITypeSpec
	item     IdentifierType
}

func (typeDefWithOne *RwTypedefWithOne) FileName() string {
	return typeDefWithOne.typeSpec.FileName()
}

func (typeDefWithOne *RwTypedefWithOne) LineNumber() int {
	return typeDefWithOne.typeSpec.LineNumber()
}

func (typeDefWithOne *RwTypedefWithOne) RowNumber() int {
	return typeDefWithOne.typeSpec.RowNumber()
}

func (typeDefWithOne *RwTypedefWithOne) IncludeImports(m map[string]string) {
	typeDefWithOne.typeSpec.IncludeImports(m)
}

func (typeDefWithOne *RwTypedefWithOne) BuildType(w io.Writer, params BuildTypeParams) (n int, err error) {
	switch params.forWhere {
	case BtMemberTypeDecl:
		_, _ = io.WriteString(w, fmt.Sprintf("%v", typeDefWithOne.item.ExportName()))
		break
	default:
		_, _ = typeDefWithOne.typeSpec.BuildType(w, params)
	}
	return 0, nil
}

func (typeDefWithOne *RwTypedefWithOne) Name() string {
	return typeDefWithOne.ExportName()
}

func (typeDefWithOne *RwTypedefWithOne) ExportName() string {
	return typeDefWithOne.item.ExportName()
}
