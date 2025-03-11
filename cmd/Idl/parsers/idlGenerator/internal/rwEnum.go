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
	name  string
	items []Enumerator
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
			return fmt.Sprintf("\"%v\"", expr.Value)
		default:
			return fmt.Sprintf("%v", expr.Value)
		}

	case *ast.UnaryExpr:
		return fmt.Sprintf("(%v%v)", expr.Op, expr.X)
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
		_, _ = io.WriteString(w, fmt.Sprintf("type %v  int\n", rwEnum.ExportName()))
		b := true
		for _, item := range rwEnum.items {
			b = b && item.rv == nil
		}
		if len(rwEnum.items) > 0 {
			if b {
				_, _ = io.WriteString(w, fmt.Sprintf("const (\n"))
				for idx := 0; idx < len(rwEnum.items); idx++ {
					switch idx {
					case 0:
						_, _ = io.WriteString(w, fmt.Sprintf("\t%v_%v %v = iota\n",
							rwEnum.ExportName(),
							rwEnum.items[idx].identifier.ExportName(),
							rwEnum.ExportName()))
						break
					default:
						_, _ = io.WriteString(w, fmt.Sprintf("\t%v_%v\n",
							rwEnum.ExportName(),
							rwEnum.items[idx].identifier.ExportName()))
						break
					}
				}
				_, _ = io.WriteString(w, fmt.Sprintf(")\n"))

				_, _ = io.WriteString(w, fmt.Sprintf("var %vNames = []string{\n", rwEnum.ExportName()))
				for _, item := range rwEnum.items {
					_, _ = io.WriteString(w, fmt.Sprintf("\t\t%v_%v: \"%v\",\n", rwEnum.ExportName(), item.identifier.ExportName(), item.identifier))
				}
				_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
				_, _ = io.WriteString(w, fmt.Sprintf("func (k %v) String() string {\n", rwEnum.ExportName()))
				_, _ = io.WriteString(w, fmt.Sprintf("\tif uint(k) < uint(len(%vNames)) {\n", rwEnum.ExportName()))
				_, _ = io.WriteString(w, fmt.Sprintf("\t\treturn %vNames[uint(k)]\n", rwEnum.ExportName()))
				_, _ = io.WriteString(w, fmt.Sprintf("\t}\n"))
				_, _ = io.WriteString(w, fmt.Sprintf("\treturn \"%v(\" + strconv.Itoa(int(k)) + \")\"\n", rwEnum.ExportName()))
				_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
			} else {
				_, _ = io.WriteString(w, fmt.Sprintf("const (\n"))
				for idx := 0; idx < len(rwEnum.items); idx++ {
					_, _ = io.WriteString(w, fmt.Sprintf("\t%v_%v %v = %v\n",
						rwEnum.ExportName(),
						rwEnum.items[idx].identifier.ExportName(),
						rwEnum.ExportName(),
						walkExpressionTree(rwEnum.items[idx].rv)))
				}
				_, _ = io.WriteString(w, fmt.Sprintf(")\n"))
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
				_, _ = io.WriteString(w, fmt.Sprintf("\treturn v\n"))
				_, _ = io.WriteString(w, fmt.Sprintf("\t}\n"))
				_, _ = io.WriteString(w, fmt.Sprintf("\t}\n"))
				_, _ = io.WriteString(w, fmt.Sprintf("\treturn \"%v(\" + strconv.Itoa(int(k)) + \")\"\n", rwEnum.ExportName()))
				_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
				_, _ = io.WriteString(w, fmt.Sprintf("\n"))
			}
		}

		_, _ = io.WriteString(w, fmt.Sprintf("var %vValues = map[string]%v{\n", rwEnum.name, rwEnum.ExportName()))
		for _, item := range rwEnum.items {
			_, _ = io.WriteString(w, fmt.Sprintf("\t\t\"%v\": %v_%v,\n",
				item.identifier,
				rwEnum.ExportName(),
				item.identifier.ExportName(),
			))
		}
		_, _ = io.WriteString(w, fmt.Sprintf("}\n"))
		return n, nil
	}
}

func (rwEnum *RwEnum) IncludeImports(m map[string]string) {
	m["strconv"] = "strconv"
}

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
