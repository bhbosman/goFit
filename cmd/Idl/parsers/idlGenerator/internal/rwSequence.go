package internal

import (
	"fmt"
	"go/ast"
	"io"
)

type RwSequence struct {
	fileInformation
	typeSpec ITypeSpec
	size     ast.Expr
}

func (rwSequence *RwSequence) IncludeImports(m map[string]string) {
	rwSequence.typeSpec.IncludeImports(m)
}

func (rwSequence *RwSequence) BuildType(w io.Writer, params BuildTypeParams) (n int, err error) {

	if params.enforceName {
		_, _ = io.WriteString(w, fmt.Sprintf("type %v ", params.enforcedExportName))
		_, _ = rwSequence.BuildType(w, BuildTypeParams{false, "", BtNoDefined})
		_, _ = io.WriteString(w, fmt.Sprintf("\n"))
	} else {
		size := walkExpressionTree(rwSequence.size)
		if size == "" {
			_, _ = io.WriteString(w, "[]")
			_, _ = rwSequence.typeSpec.BuildType(w, BuildTypeParams{false, "", BtMemberTypeDecl})
		} else {
			_, _ = io.WriteString(w, fmt.Sprintf("[%v]", size))
			_, _ = rwSequence.typeSpec.BuildType(w, BuildTypeParams{false, "", BtMemberTypeDecl})
		}
	}
	return 0, nil
}

func (rwSequence *RwSequence) FileName() string {
	return rwSequence.fileName
}

func (rwSequence *RwSequence) LineNumber() int {
	return rwSequence.lineNumber
}

func (rwSequence *RwSequence) RowNumber() int {
	return rwSequence.rowNumber
}
