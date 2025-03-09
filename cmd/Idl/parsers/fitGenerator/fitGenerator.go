package fitGenerator

import (
	"awesomeProject/cmd/Idl/parsers/common"
	"awesomeProject/cmd/Idl/parsers/fitGenerator/internal"
	"fmt"
	"io"
)

func GenerateFit(inputFile string, packageName string, readCloser io.ReadCloser, writer io.Writer, onReadIncludeFile common.OnReadIncludeFile) error {
	idlParse := func(data internal.IYYParserHandler, YYlex internal.YYLexer) (v int, err error) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case error:
					err = v
				case string:
					err = fmt.Errorf(v)
				}

			}
		}()
		v = internal.YYNewParser[internal.IYYParserHandler](data).Parse(YYlex)
		return v, err
	}

	handler := &internal.FitParseDataHandler{
		PackageName: packageName,
		FileName:    inputFile,
		Writer:      writer,
		Language:    internal.LanguageGo,
	}
	defer func(handler *internal.FitParseDataHandler) {
		err := handler.Close()
		if err != nil {

		}
	}(handler)

	lexer := internal.NewLexer(
		inputFile,
		readCloser,
		onReadIncludeFile,
	)
	defer func(lexer io.Closer) {
		err := lexer.Close()
		if err != nil {
		}
	}(lexer)
	_, err := idlParse(handler, lexer)
	if err != nil {
		return err
	}
	return nil
}

func SetDebugLevel(level int) {
	internal.YYDebug = level
}
