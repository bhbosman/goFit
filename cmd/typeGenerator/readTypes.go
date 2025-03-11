package main

import (
	"fmt"
	"github.com/thedatashed/xlsxreader"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func readTypes() ([]ConstantDeclaration, error) {
	xlsxReader, err := xlsxreader.OpenFile("./cmd/typeGenerator/Profile.xlsx")
	if err != nil {
		os.Exit(1)
	}
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
		}
	}(xlsxReader)
	ch := xlsxReader.XlsxFile.ReadRows("Types")
	// skip line
	<-ch
	var constantDeclaration ConstantDeclaration
	var constantDeclarations []ConstantDeclaration
	for row := range ch {
		collValue := func(col string) string {
			for i := range row.Cells {
				if row.Cells[i].Column == col {
					return row.Cells[i].Value
				}
			}
			return ""
		}
		cv := collValue("A")
		if cv != "" {
			if constantDeclaration.Name != "" {
				constantDeclarations = append(constantDeclarations, constantDeclaration)
			}
			constantDeclaration.ConstantValue = nil
			constantDeclaration.Name = FitIdentifier(cv)
			constantDeclaration.Type = StringToConstantType(collValue("B"))
			constantDeclaration.Comment = collValue("E")
		} else {
			idx, err := strconv.ParseInt(strings.TrimSpace(collValue("D")), 0, 64)
			if err != nil {
				return nil, err
			}

			dd := collValue("E")
			if !strings.Contains(dd, "Deprecated") {
				constantDeclaration.ConstantValue = append(
					constantDeclaration.ConstantValue,
					ConstantValue{
						FitIdentifier(collValue("C")),
						int(idx),
						ConstantValueComment(collValue("E")),
					},
				)
			}
		}
	}
	if constantDeclaration.Name != "" {
		constantDeclarations = append(constantDeclarations, constantDeclaration)
	}
	return constantDeclarations, nil
}

func generateTypes(decls []ConstantDeclaration) error {
	packageName := "types"
	folder := path.Join(".", "fitIdl", packageName)
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		return err
	}

	file := path.Join(folder, fmt.Sprintf("fitAllTypes.idl"))
	combined, err := os.Create(file)
	if err != nil {
		return err
	}

	{
		for _, declaration := range decls {
			strings.Replace("ddd", "@", "", -1)
			file := path.Join(folder, fmt.Sprintf("fit%v.idl", declaration.Name))
			var constantFile *os.File
			constantFile, err = os.Create(file)
			if err != nil {
				return err
			}
			_, _ = io.WriteString(combined, fmt.Sprintf("#include \"fit%v.idl\"\n", declaration.Name))

			_, _ = io.WriteString(constantFile, fmt.Sprintf("#include \"fitTypeDecls.idl\"\n"))
			_, _ = io.WriteString(constantFile, fmt.Sprintf("\n"))
			_, _ = io.WriteString(constantFile, fmt.Sprintf("enum<"))
			_, _ = io.WriteString(constantFile, fmt.Sprintf("%v,", declaration.Type))

			switch string(declaration.Type) {
			case "enum", "byte", "uint8", "uint8z":
				_, _ = io.WriteString(constantFile, fmt.Sprintf("0x%02X", declaration.Type.InvalidValue()))
			case "uint16", "uint16z":
				_, _ = io.WriteString(constantFile, fmt.Sprintf("0x%04X", declaration.Type.InvalidValue()))
			default:
				_, _ = io.WriteString(constantFile, fmt.Sprintf("0x%08X", declaration.Type.InvalidValue()))
			}

			_, _ = io.WriteString(constantFile, fmt.Sprintf("> "))
			_, _ = io.WriteString(constantFile, fmt.Sprintf("%v\n", declaration.Name))
			_, _ = io.WriteString(constantFile, fmt.Sprintf("{\n"))
			for i, value := range declaration.ConstantValue {
				_, _ = io.WriteString(constantFile, fmt.Sprintf("\t@%v = ", value.Name))
				switch string(declaration.Type) {
				case "enum", "byte", "uint8", "uint8z":
					_, _ = io.WriteString(constantFile, fmt.Sprintf("0x%02X", value.Value))
				case "uint16", "uint16z":
					_, _ = io.WriteString(constantFile, fmt.Sprintf("0x%04X", value.Value))
				default:
					_, _ = io.WriteString(constantFile, fmt.Sprintf("0x%08X", value.Value))
				}
				if i < len(declaration.ConstantValue)-1 {
					_, _ = io.WriteString(constantFile, fmt.Sprintf(","))
				}
				_, _ = io.WriteString(constantFile, fmt.Sprintf(" // Value: %v, Comment: %v\n", value.Value, declaration.Comment))
			}
			_, _ = io.WriteString(constantFile, fmt.Sprintf("};\n"))
		}
	}
	return nil
}
