package internal

import (
	"fmt"
	"io"
)

// RwStruct
// Must implement the following interfaces:
//  1. IDefinition
//  2. IFileInformation
//  3. ITypeSpec
type RwStruct struct {
	fileInformation
	forward bool
	name    string
	Members []Member
}

type Declarator struct {
	Name      IdentifierType
	IsArray   bool
	ArraySize []int
}

type Member struct {
	fileInformation
	TypeDeclaration ITypeSpec
	Declarator      []Declarator
}

func (rwStruct *RwStruct) BuildType(w io.Writer, params BuildTypeParams) (n int, err error) {
	structName := rwStruct.ExportName()
	if params.enforceName {
		structName = params.enforcedExportName
	}

	_, _ = io.WriteString(w, fmt.Sprintf("type %v struct {\n", structName))
	for _, member := range rwStruct.Members {
		if member.TypeDeclaration != nil {
			for _, declarator := range member.Declarator {
				_, _ = io.WriteString(w, fmt.Sprintf("\t%v ", declarator.Name.ExportName()))
				_, _ = member.TypeDeclaration.BuildType(w, BuildTypeParams{false, "", BtMemberTypeDecl})
				_, _ = io.WriteString(w, fmt.Sprintf("\n"))
			}
		}
	}
	_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
	return 0, nil
}

func (rwStruct *RwStruct) IncludeImports(map[string]string) {
}

func (rwStruct *RwStruct) FileName() string {
	return rwStruct.fileName
}

func (rwStruct *RwStruct) LineNumber() int {
	return rwStruct.lineNumber
}

func (rwStruct *RwStruct) RowNumber() int {
	return rwStruct.rowNumber
}

func (rwStruct *RwStruct) Name() string {
	return rwStruct.ExportName()
}

func (rwStruct *RwStruct) ExportName() string {
	return __exportName(rwStruct.name)
}

func (rwStruct *RwStruct) DefinitionType() DefinitionType {
	return DtStruct
}
