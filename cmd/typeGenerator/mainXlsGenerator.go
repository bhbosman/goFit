package main

import (
	"fmt"
	"github.com/thedatashed/xlsxreader"
	"io"
	"strconv"
	"strings"

	//"strconv"
	//"strings"

	//"github.com/thedatashed/xlsxreader"
	//"io"
	"os"
	"path"
)

func main() {
	messages, err := readMessages()
	if err != nil {
		return
	}
	err = generateMessages(messages)
	if err != nil {
		return
	}

	cd, err := readTypes()
	if err != nil {
		return
	}
	err = generateTypes(cd)
	if err != nil {
		return
	}
}

func generateMessages(messages []*FitMessage) error {
	packageName := "types"
	folder := path.Join(".", "fitIdl", packageName)
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		return err
	}

	allFile := path.Join(folder, "fitAllMessages.idl")
	all, err := os.Create(allFile)
	if err != nil {
		return err
	}

	for _, declaration := range messages {
		fileName := fmt.Sprintf("fit%v.idl", declaration.Name)
		file := path.Join(folder, fileName)

		_, _ = io.WriteString(all, fmt.Sprintf("#include \"%v\"\n", fileName))
		var constantFile *os.File
		constantFile, err = os.Create(file)
		if err != nil {
			return err
		}

		_, _ = io.WriteString(constantFile, fmt.Sprintf("#include \"fitTypeDecls.idl\"\n"))
		_, _ = io.WriteString(constantFile, fmt.Sprintf("#include \"fitAllTypes.idl\"\n"))
		_, _ = io.WriteString(constantFile, fmt.Sprintf("\n"))

		_, _ = io.WriteString(constantFile, fmt.Sprintf("struct<\"%v\", \"%v\"> %v {\n", declaration.Name, declaration.comment, declaration.Name))
		for _, field := range declaration.Fields {

			_, _ = io.WriteString(constantFile,
				fmt.Sprintf("\t"+
					"%v<\"%v\", %v, %v, \"%v\", %v, %v, \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\"> ",
					field.fieldType,
					field.fieldType.BaseType(),
					field.defNumber,
					func(s string) int {
						ss := strings.TrimSpace(s)
						if ss == "" {
							return -1
						}
						if ss == "[N]" {
							return 0
						}
						n, _ := strconv.ParseInt(ss, 10, 64)
						return int(n)
					}(field.arrayType),
					field.components,
					func(s string) float64 {
						ss := strings.TrimSpace(s)
						if ss == "" {
							return 1
						}
						n, _ := strconv.ParseFloat(ss, 64)
						return n
					}(field.scale),
					func(s string) float64 {
						ss := strings.TrimSpace(s)
						if ss == "" {
							return 1
						}
						n, _ := strconv.ParseFloat(ss, 64)
						return n

					}(field.offset),
					field.units,
					field.bits,
					field.accumulate,
					field.refFieldName,
					field.refFieldValue,
					strings.Replace(field.comment, "\n", "", -1),
					field.products,
					field.example,
				),
			)
			_, _ = io.WriteString(constantFile, fmt.Sprintf("@%v", field.fieldName))
			if len(field.Fields) > 0 {
				_, _ = io.WriteString(constantFile, fmt.Sprintf("\n"))
				_, _ = io.WriteString(constantFile, fmt.Sprintf("\t{\n"))
				for _, subField := range field.Fields {

					_, _ = io.WriteString(constantFile,
						fmt.Sprintf("\t\t"+
							"%v<\"%v\", \"%v\", %v, %v, \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\"> ",
							field.fieldType,
							field.fieldType.BaseType(),
							subField.components,
							func(s string) float64 {
								ss := strings.TrimSpace(s)
								if ss == "" {
									return 1
								}
								n, _ := strconv.ParseFloat(ss, 64)
								return n
							}(subField.scale),
							func(s string) float64 {
								ss := strings.TrimSpace(s)
								if ss == "" {
									return 1
								}
								n, _ := strconv.ParseFloat(ss, 64)
								return n

							}(subField.offset),
							subField.units,
							subField.bits,
							subField.accumulate,
							subField.refFieldName,
							subField.refFieldValue,
							strings.Replace(subField.comment, "\n", "", -1),
							subField.products,
							subField.example,
						),
					)
					_, _ = io.WriteString(constantFile, fmt.Sprintf("@%v;\n", subField.fieldName))

				}
				_, _ = io.WriteString(constantFile, fmt.Sprintf("\t};\n"))
			} else {
				_, _ = io.WriteString(constantFile, fmt.Sprintf(";\n"))
			}

		}

		_, _ = io.WriteString(constantFile, fmt.Sprintf("};\n"))

		err = constantFile.Close()
		if err != nil {
			return err
		}
	}
	_ = all.Close()
	return nil
}

func readMessages() ([]*FitMessage, error) {
	xlsxReader, err := xlsxreader.OpenFile("./cmd/typeGenerator/Profile.xlsx")
	if err != nil {
		os.Exit(1)
	}
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
		}
	}(xlsxReader)
	ch := xlsxReader.XlsxFile.ReadRows("Messages")
	// skip line
	<-ch

	fitMessage := &FitMessage{}
	var fitMessageField *FitMessageField
	var fitMessages []*FitMessage
	for row := range ch {
		collValue := func(col string) string {
			for i := range row.Cells {
				if row.Cells[i].Column == col {
					return row.Cells[i].Value
				}
			}
			return ""
		}

		if collValue("A") == "" && collValue("B") == "" && collValue("C") == "" {
			continue
		}

		cv := collValue("A")
		if cv != "" {
			if fitMessage.Name != "" {
				fitMessages = append(fitMessages, fitMessage)
				//return nil, err
			}
			fitMessage = &FitMessage{}
			fitMessage.Name = FitIdentifier(cv)

		} else {
			defNumberAsString := collValue("B")
			if defNumberAsString != "" {
				defNumber, err := strconv.ParseInt(strings.TrimSpace(defNumberAsString), 0, 64)
				if err != nil {
					return nil, err
				}
				fitMessageField = &FitMessageField{
					int(defNumber),
					FitIdentifier(collValue("C")),
					FitTypeIdentifier(collValue("D")),
					collValue("E"),
					collValue("F"),
					collValue("G"),
					collValue("H"),
					collValue("I"),
					collValue("J"),
					collValue("K"),
					collValue("L"),
					collValue("M"),
					collValue("N"),
					collValue("O"),
					collValue("P"),
					[]*FitSubMessageField{},
				}

				fitMessage.Fields = append(fitMessage.Fields, fitMessageField)

			} else {
				fitSubMessageField := &FitSubMessageField{
					FitIdentifier(collValue("C")),
					collValue("D"),
					collValue("E"),
					collValue("F"),
					collValue("G"),
					collValue("H"),
					collValue("I"),
					collValue("J"),
					collValue("K"),
					collValue("L"),
					collValue("M"),
					collValue("N"),
					collValue("O"),
					collValue("P"),
				}
				fitMessageField.Fields = append(fitMessageField.Fields, fitSubMessageField)

			}

			//idx, _ := strconv.Atoi(collValue("D"))
			//constantDeclaration.ConstantValue = append(
			//	constantDeclaration.ConstantValue,
			//	ConstantValue{
			//		FitIdentifier(collValue("C")),
			//		idx,
			//		ConstantValueComment(collValue("E")),
			//	},
			//)
		}
	}
	if fitMessage.Name != "" {
		fitMessages = append(fitMessages, fitMessage)
		//return nil, err
	}

	return fitMessages, nil
}
