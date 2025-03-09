package internal

import (
	"errors"
	"fmt"
	"github.com/bhbosman/goFit/cmd/Idl/parsers/common"
	"go.uber.org/multierr"
	"go/ast"
	"go/token"
	"io"
	"strconv"
	"strings"
)

const (
	Initial = iota
	IncludeExpression
	DefineExpression
)

type yylexer struct {
	filesUsed       map[string]int
	streamData      [64]StreamData
	streamDataIndex int
	startCond       [64]int
	startCondIndex  int

	onReadIncludeFile func(fileName string) (io.ReadCloser, string, error)
	fileName          string
	buf               []byte // the current token been read
	empty             bool
	current           byte
}

func NewLexer(fileName string, src io.ReadCloser, onReadIncludeFile common.OnReadIncludeFile) *yylexer {
	y := &yylexer{
		filesUsed:         map[string]int{fileName: 0},
		streamData:        [64]StreamData{{src, fileName, bytesData{}, 1, 1}},
		streamDataIndex:   0,
		onReadIncludeFile: onReadIncludeFile,
		fileName:          fileName,
	}
	tmp := [1]byte{0}
	if _, err := y.readFull(tmp[:1]); err == nil {
		y.current = tmp[0]
	}
	return y
}

func (y *yylexer) Close() error {
	var err error
	for idx := 0; idx < len(y.streamData); idx++ {
		err = multierr.Append(err, y.streamData[idx].Close())
	}
	return err
}

func (y *yylexer) readFull(p []byte) (int, error) {
	// Unread the overshot bytes, if any.

	for {
		var streamData *StreamData = &y.streamData[y.streamDataIndex]
		n := copy(p, streamData.bytes.buf[streamData.bytes.i:streamData.bytes.j])
		p = p[n:]
		streamData.bytes.i += n
		streamData.bytes.pos += n
		if len(p) == 0 {
			break
		}
		for {
			if err := y.fill(); err != nil {
				return n, err
			}
			break
		}
	}
	return len(p), nil
}

func (y *yylexer) fill() error {
	var streamData *StreamData = &y.streamData[y.streamDataIndex]
	if streamData.bytes.i != streamData.bytes.j {
		panic("fill called when unread bytes exist")
	}
	if streamData.bytes.j > 1 {
		streamData.bytes.buf[0] = streamData.bytes.buf[streamData.bytes.j-1]
		streamData.bytes.i = 1
		streamData.bytes.j = 1
		streamData.bytes.unread = 0
	}

	n, err := streamData.Read(streamData.bytes.buf[streamData.bytes.j:])
	streamData.bytes.j += n
	if n > 0 {
		return nil
	}
	if err == io.EOF {
		err = io.ErrUnexpectedEOF
	}
	return err
}

type FileAlreadyReadError struct {
	files []string
}

func (fileAlreadyReadError *FileAlreadyReadError) Error() string {
	files := strings.Join(fileAlreadyReadError.files, ",")
	return fmt.Sprintf("the following files created a circular reference.  %s", files)
}

func (y *yylexer) LoadFile(fileName string) error {
	if _, ok := y.filesUsed[fileName]; ok {
		arr := make([]string, len(y.filesUsed))
		for key, value := range y.filesUsed {
			arr[value] = key
		}
		err := &FileAlreadyReadError{arr}
		return err
	}

	// the parser already read the next char on the current file
	// unread one byte that will be used when the stream is read again
	y.streamData[y.streamDataIndex].bytes.i--
	y.streamData[y.streamDataIndex].bytes.unread++

	if y.onReadIncludeFile != nil {
		readCloser, fullFileName, err := y.onReadIncludeFile(fileName)
		if err != nil {
			return err
		}
		y.streamDataIndex++
		y.streamData[y.streamDataIndex] = StreamData{readCloser, fullFileName, bytesData{}, 1, 1}
		y.filesUsed[fullFileName] = len(y.filesUsed)
	}
	return nil
}

func (y *yylexer) getc() (byte, error) {
	if y.current != 0 {
		y.buf = append(y.buf, y.current)
	}
	y.current = 0
	tmp := [1]byte{0}
	index := 0

	for {
		if lengthRead, err := y.readFull(tmp[index : index+1]); err != nil {
			switch {
			case errors.Is(err, io.ErrUnexpectedEOF):
				index += lengthRead
				if y.streamDataIndex > 0 {
					err = y.streamData[y.streamDataIndex].Close()
					delete(y.filesUsed, y.streamData[y.streamDataIndex].fileName)
					y.streamData[y.streamDataIndex] = StreamData{nil, "", bytesData{}, 0, 0}
					y.streamDataIndex--
					continue
				}
				return 0, err
			default:
				return 0, err
			}
		}
		y.current = tmp[0]
		return y.current, nil
	}
}

func (y *yylexer) Error(e string) {
	currentStream := y.streamData[y.streamDataIndex]
	s := fmt.Sprintf("error: %v, fileName: %v (%v,%v)", e, currentStream.fileName, currentStream.lineNumber, currentStream.rowNumber-len(y.buf))
	panic(errors.New(s))
}

func (y *yylexer) onCreateInteger(val *YYSymType) (int, error) {
	intValue, _ := strconv.ParseInt(string(y.buf[:len(y.buf)]), 10, 64)
	astValue := &ast.BasicLit{Value: strconv.Itoa(int(intValue)), Kind: token.INT}
	val.Node = UnionData[ast.Expr]{
		y.FileName(),
		y.streamData[y.streamDataIndex].lineNumber,
		y.streamData[y.streamDataIndex].rowNumber,
		astValue,
	}
	return Integer_literal, nil

}

func (y *yylexer) onCreateFloat(val *YYSymType) (int, error) {
	v, _ := strconv.ParseFloat(string(y.buf), 64)
	astValue := &ast.BasicLit{Value: strconv.FormatFloat(v, 'g', -1, 64), Kind: token.FLOAT}
	val.Node = UnionData[ast.Expr]{
		y.FileName(),
		y.streamData[y.streamDataIndex].lineNumber,
		y.streamData[y.streamDataIndex].rowNumber,
		astValue,
	}
	return Floating_pt_literal, nil
}

func (y *yylexer) Lex(val *YYSymType) int {
	c, err := y.internalLex(val)
	if err != nil {
		switch {
		case errors.Is(err, io.ErrUnexpectedEOF):
			return 0
		default:
			panic(err)
		}
	}
	return c
}

func (y *yylexer) OnHexValue(val *YYSymType) (int, error) {
	intValue, _ := strconv.ParseInt(string(y.buf[2:len(y.buf)]), 16, 64)
	astValue := &ast.BasicLit{
		ValuePos: 0,
		Kind:     token.INT,
		Value:    strconv.Itoa(int(intValue)),
	}
	val.Node = UnionData[ast.Expr]{
		y.FileName(),
		y.streamData[y.streamDataIndex].lineNumber,
		y.streamData[y.streamDataIndex].rowNumber,
		astValue,
	}
	return Integer_literal, nil
}

func (y *yylexer) OnStringLiteral(val *YYSymType) (int, error) {
	s := strings.TrimSpace(string(y.buf[1 : len(y.buf)-1]))
	astValue := &ast.BasicLit{
		ValuePos: 0,
		Kind:     token.STRING,
		Value:    s,
	}

	val.Node = UnionData[ast.Expr]{
		y.FileName(),
		y.streamData[y.streamDataIndex].lineNumber,
		y.streamData[y.streamDataIndex].rowNumber,
		astValue}
	return String_literal, nil
}

func (y *yylexer) FileName() string {
	return y.streamData[y.streamDataIndex].fileName
}

func (y *yylexer) RowNumber() int {
	n := y.streamData[y.streamDataIndex].rowNumber
	y.streamData[y.streamDataIndex].rowNumber += len(y.buf)
	return n
}

func (y *yylexer) LineNumber() int {
	return y.streamData[y.streamDataIndex].lineNumber
}

type bytesData struct {
	buf    [4096]byte
	i, j   int
	pos    int
	unread byte
}

type StreamData struct {
	rr         io.ReadCloser
	fileName   string
	bytes      bytesData
	lineNumber int
	rowNumber  int
}

func (sd *StreamData) Read(p []byte) (n int, err error) {
	if sd.rr != nil {
		return sd.rr.Read(p)
	}
	return 0, io.ErrUnexpectedEOF
}

func (sd *StreamData) Close() error {
	if sd.rr != nil {
		if err := sd.rr.Close(); err != nil {
			return err
		}
	}
	return nil
}
