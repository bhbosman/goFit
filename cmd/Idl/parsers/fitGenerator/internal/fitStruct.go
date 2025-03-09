package internal

import (
	"fmt"
	"go/ast"
	"io"
)

// fitStruct
// Must implement the following interfaces:
//  1. IDefinition
//  2. IFileInformation
//  3. ITypeSpec
type fitStruct struct {
	fileInformation
	identityNumber ast.Expr
	comment        ast.Expr
	name           string
	Members        []FitMember
}

type Declarator struct {
	Name      IdentifierType
	IsArray   bool
	ArraySize []int
}

type FitMember struct {
	fileInformation
	TypeDeclaration ITypeSpec
	Declarator      Declarator
}

func (rwStruct *fitStruct) BuildType(w io.Writer, params BuildTypeParams) (n int, err error) {
	structName := rwStruct.ExportName()
	if params.enforceName {
		structName = params.enforcedExportName
	}

	_, _ = io.WriteString(w, fmt.Sprintf("type %v struct {\n", structName))
	for _, member := range rwStruct.Members {
		if member.TypeDeclaration != nil {
			_, _ = io.WriteString(w, fmt.Sprintf("\t%v ", member.Declarator.Name.ExportName()))
			_, _ = member.TypeDeclaration.BuildType(w, BuildTypeParams{false, "", BtMemberTypeDecl})
			_, _ = io.WriteString(w, fmt.Sprintf("\n"))
		}
	}
	_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
	//

	_, _ = io.WriteString(w, fmt.Sprintf("func (obj * %v) MsgNumber() uint16 {\n", structName))
	_, _ = io.WriteString(w, fmt.Sprintf("\t return uint16(MesgNum_%v)\n", walkExpressionTree(rwStruct.identityNumber)))
	_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
	_, _ = io.WriteString(w, fmt.Sprintf("\n"))

	_, _ = io.WriteString(w, fmt.Sprintf("func (obj *%v) Read(data []byte, md *messages.MessageDefinition) error {\n", structName))
	_, _ = io.WriteString(w, fmt.Sprintf("\t return nil\n"))
	_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
	_, _ = io.WriteString(w, fmt.Sprintf("\n"))

	_, _ = io.WriteString(w, fmt.Sprintf("var fieldMapFor%v = map[uint16]registration.DefinedFields{\n", structName))
	for _, member := range rwStruct.Members {
		switch v := member.TypeDeclaration.(type) {
		case *MemberType:
			_, _ = io.WriteString(w, fmt.Sprintf("\t%v: {", walkExpressionTree(v.defNumber)))
			_, _ = io.WriteString(w, fmt.Sprintf("BaseType: messages.%v, ", walkExpressionTree(v.baseType)))
			_, _ = io.WriteString(w, fmt.Sprintf("FieldId: %v, ", walkExpressionTree(v.defNumber)))
			_, _ = io.WriteString(w, fmt.Sprintf("Name: \"%v\", ", member.Declarator.Name.ExportName()))
			arraySize := walkExpressionTree(v.arrayType)
			if arraySize == "-1" {
				_, _ = io.WriteString(w, fmt.Sprintf("IsArray: false, "))
				_, _ = io.WriteString(w, fmt.Sprintf("ArrayLength: 0, "))

			} else {
				_, _ = io.WriteString(w, fmt.Sprintf("IsArray: true, "))
				_, _ = io.WriteString(w, fmt.Sprintf("ArrayLength: %v, ", arraySize))
			}

			_, _ = io.WriteString(w, fmt.Sprintf("Components: \"%v\", ", walkExpressionTree(v.components)))
			_, _ = io.WriteString(w, fmt.Sprintf("Scale: %v, ", walkExpressionTree(v.scale)))
			_, _ = io.WriteString(w, fmt.Sprintf("Offset: %v, ", walkExpressionTree(v.offset)))
			_, _ = io.WriteString(w, fmt.Sprintf("Units: \"%v\", ", walkExpressionTree(v.units)))
			_, _ = io.WriteString(w, fmt.Sprintf("Bits: \"%v\", ", walkExpressionTree(v.defNumber)))
			_, _ = io.WriteString(w, fmt.Sprintf("Accumulate: \"%v\", ", walkExpressionTree(v.accumulate)))
			_, _ = io.WriteString(w, fmt.Sprintf("RefFieldName: \"%v\", ", walkExpressionTree(v.refFieldName)))
			_, _ = io.WriteString(w, fmt.Sprintf("RefFieldValue: \"%v\", ", walkExpressionTree(v.refFieldValue)))
			_, _ = io.WriteString(w, fmt.Sprintf("Comment: \"%v\", ", walkExpressionTree(v.comment)))
			_, _ = io.WriteString(w, fmt.Sprintf("Products: \"%v\", ", walkExpressionTree(v.products)))
			_, _ = io.WriteString(w, fmt.Sprintf("Example: \"%v\", ", walkExpressionTree(v.example)))
			_, _ = io.WriteString(w, fmt.Sprintf("\t},\n"))
		}
	}
	_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
	_, _ = io.WriteString(w, fmt.Sprintf("\n"))

	_, _ = io.WriteString(w, fmt.Sprintf("func init() {\n"))
	_, _ = io.WriteString(w, fmt.Sprintf("\t_ = registration.RegisterMessage(\n"))
	_, _ = io.WriteString(w, fmt.Sprintf("\t\tuint16(MesgNum_%v),\n", structName))
	_, _ = io.WriteString(w, fmt.Sprintf("\t\tfieldMapFor%v,\n", structName))
	_, _ = io.WriteString(w, fmt.Sprintf("\t\tfunc() (registration.IFitMessage, error) {\n"))
	_, _ = io.WriteString(w, fmt.Sprintf("\t\t\treturn &%v{}, nil\n", structName))
	_, _ = io.WriteString(w, fmt.Sprintf("\t\t},\n"))
	_, _ = io.WriteString(w, fmt.Sprintf("\t)\n"))
	_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
	_, _ = io.WriteString(w, fmt.Sprintf("\n"))

	_, _ = io.WriteString(w, fmt.Sprintf("\n"))

	return 0, nil
}

func (rwStruct *fitStruct) IncludeImports(m map[string]string) {
	m["messages"] = "github.com/bhbosman/goFit/fitIdl/messages"
	m["registration"] = "github.com/bhbosman/goFit/fitIdl/registration"
}

func (rwStruct *fitStruct) FileName() string {
	return rwStruct.fileName
}

func (rwStruct *fitStruct) LineNumber() int {
	return rwStruct.lineNumber
}

func (rwStruct *fitStruct) RowNumber() int {
	return rwStruct.rowNumber
}

func (rwStruct *fitStruct) Name() string {
	return rwStruct.ExportName()
}

func (rwStruct *fitStruct) ExportName() string {
	return __exportName(rwStruct.name)
}

func (rwStruct *fitStruct) DefinitionType() DefinitionType {
	return DtStruct
}
