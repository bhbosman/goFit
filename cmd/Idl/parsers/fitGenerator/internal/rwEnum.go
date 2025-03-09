package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

type IdentifierType string

func (identifierType IdentifierType) ExportName() string {
	return __exportName(string(identifierType))
}

type Enumerator struct {
	fileInformation
	identifier IdentifierType
	rv         ast.Expr
}

// RwEnum
// Must implement the following interfaces:
//  1. IDefinition
//  2. IFileInformation
//  3. ITypeSpec
type RwEnum struct {
	fileInformation
	name     string
	items    []Enumerator
	typeSpec ITypeSpec
	node     ast.Expr
}

func internalWalkExpressionTree(expression ast.Expr) string {
	if expression == nil {
		return ""
	}
	switch expr := expression.(type) {
	case *ast.BinaryExpr:
		l := internalWalkExpressionTree(expr.X)
		r := internalWalkExpressionTree(expr.Y)
		return fmt.Sprintf("%v %v %v", l, expr.Op, r)
	case *ast.BasicLit:
		switch expr.Kind {
		case token.STRING:
			return fmt.Sprintf("%v", expr.Value)
		default:
			return fmt.Sprintf("%v", expr.Value)
		}

	case *ast.UnaryExpr:
		l := internalWalkExpressionTree(expr.X)
		//r := internalWalkExpressionTree(expr.Y)
		switch expr.Op {
		case token.SUB:
			return fmt.Sprintf("-%v", l)
		default:
			panic("implement case *ast.UnaryExpr:")
		}

	default:
		panic(fmt.Sprintf("unexpected expression type %T", expr))
	}
}

func walkExpressionTree(expression ast.Expr) string {
	return internalWalkExpressionTree(expression)
}

func (rwEnum *RwEnum) BuildType(w io.Writer, params BuildTypeParams) (n int, err error) {
	switch params.forWhere {
	case BtMemberTypeDecl:
		_, _ = io.WriteString(w, fmt.Sprintf("%v", rwEnum.ExportName()))
		return 0, err
	default:
		_, _ = io.WriteString(w, fmt.Sprintf("type %v ", rwEnum.ExportName()))
		_, _ = rwEnum.typeSpec.BuildType(w, params)
		_, _ = io.WriteString(w, fmt.Sprintf("\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("\n"))

		_, _ = io.WriteString(w, fmt.Sprintf("const (\n"))
		for idx := 0; idx < len(rwEnum.items); idx++ {
			_, _ = io.WriteString(w, fmt.Sprintf("\t%v_%v %v = %v\n",
				rwEnum.ExportName(),
				rwEnum.items[idx].identifier.ExportName(),
				rwEnum.ExportName(),
				walkExpressionTree(rwEnum.items[idx].rv)))
		}
		_, _ = io.WriteString(w, fmt.Sprintf("\t%v_Invalid %v = %v\n",
			rwEnum.ExportName(),
			rwEnum.ExportName(),
			walkExpressionTree(rwEnum.node)))

		_, _ = io.WriteString(w, fmt.Sprintf(")\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("\n"))

		_, _ = io.WriteString(w, fmt.Sprintf("var %vNames = map[%v]string{\n", rwEnum.name, rwEnum.ExportName()))
		for _, item := range rwEnum.items {
			_, _ = io.WriteString(w, fmt.Sprintf("\t\t%v_%v: \"%v\",\n",
				rwEnum.ExportName(),
				item.identifier.ExportName(),
				item.identifier))
		}
		_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("func (k %v) String() string {\n", rwEnum.name))
		_, _ = io.WriteString(w, fmt.Sprintf("\tif uint(k) < uint(len(%vNames)) {\n", rwEnum.name))
		_, _ = io.WriteString(w, fmt.Sprintf("\t\tif v, ok := %vNames[k]; ok {\n", rwEnum.ExportName()))
		_, _ = io.WriteString(w, fmt.Sprintf("\t\t\treturn v\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("\t\t}\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("\t}\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("\treturn \"%v(\" + strconv.Itoa(int(k)) + \")\"\n", rwEnum.ExportName()))
		_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("\n"))

		_, _ = io.WriteString(w, fmt.Sprintf("var %vValues = map[string]%v{\n", rwEnum.name, rwEnum.ExportName()))
		for _, item := range rwEnum.items {
			_, _ = io.WriteString(w, fmt.Sprintf("\t\t\"%v\": %v_%v,\n",
				item.identifier,
				rwEnum.ExportName(),
				item.identifier.ExportName(),
			))
		}
		_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("\n"))
		//
		_, _ = io.WriteString(w, fmt.Sprintf("func %vFromString(s string) %v {\n", rwEnum.ExportName(), rwEnum.ExportName()))
		_, _ = io.WriteString(w, fmt.Sprintf("\t\tif v, ok := %vValues[s]; ok {\n", rwEnum.ExportName()))
		_, _ = io.WriteString(w, fmt.Sprintf("\t\t\treturn v\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("\t\t}\n"))

		_, _ = io.WriteString(w, fmt.Sprintf("\treturn %v(", rwEnum.ExportName()))
		_, _ = io.WriteString(w, fmt.Sprintf("%v_Invalid", rwEnum.ExportName()))
		_, _ = io.WriteString(w, fmt.Sprintf(")\n"))

		_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
		_, _ = io.WriteString(w, fmt.Sprintf("\n"))

		return n, nil
	}
}

func (rwEnum *RwEnum) IncludeImports(m map[string]string) { m["strconv"] = "strconv" }

func (rwEnum *RwEnum) FileName() string {
	return rwEnum.fileName
}

func (rwEnum *RwEnum) LineNumber() int {
	return rwEnum.lineNumber
}

func (rwEnum *RwEnum) RowNumber() int {

	return rwEnum.rowNumber
}

func (rwEnum *RwEnum) ExportName() string {
	return __exportName(rwEnum.name)
}

func (rwEnum *RwEnum) Name() string {

	return rwEnum.ExportName()
}

func (rwEnum *RwEnum) DefinitionType() DefinitionType {
	return DtEnum
}
