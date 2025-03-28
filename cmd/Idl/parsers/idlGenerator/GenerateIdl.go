package idlGenerator

import (
	"fmt"
	"github.com/bhbosman/goFit/cmd/Idl/parsers/common"
	"github.com/bhbosman/goFit/cmd/Idl/parsers/idlGenerator/internal"
	"io"
)

func GenerateIdl(inputFile string, packageName string, readCloser io.ReadCloser, writer io.Writer, onReadIncludeFile common.OnReadIncludeFile) error {
	idlParse := func(data internal.IYYParserHandler, lexer internal.YYLexer) (v int, err error) {
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
		v = internal.YYNewParser[internal.IYYParserHandler](data).Parse(lexer)
		return v, err
	}

	handler := &internal.IdlParseDataHandler{Writer: writer, Language: internal.LanguageGo}
	defer func(handler *internal.IdlParseDataHandler) {
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
