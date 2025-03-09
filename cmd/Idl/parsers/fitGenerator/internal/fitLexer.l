%{
package internal

import (
    "strings"
    "fmt"
    "unicode"
)

import __strings__ "strings"
//import __constant__ "go/constant"
//import __strconv__ "strconv"
//import __token__ "go/token"
//import __ast__ "go/ast"


func (y *yylexer) internalLex(lval *YYSymType) (int, error) {
	var err error
	c := y.current
	var rowNumber int
	var lineNumber int
	var fileName string

	if y.empty {
		vv, err := y.getc()
		if err != nil {
			return 0, err
		}
		c, y.empty = vv, false
	}

%}

%yyc c
%yyn if c, err = y.getc(); err != nil{return 0, err}
%yyt y.startCond[y.startCondIndex]

D  [0-9]+
E  [eE][-+]?{D}
F  {D}"."{D}?{E}?|{D}{E}?|"."{D}{E}?
%x IncludeExpression
%x DefineExpression


%%
    y.buf = y.buf[:0]

<INITIAL>"+"
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "+"}
	//
	return int('+'), nil
<INITIAL>"-"
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "-"}
	//
	return int('-'), nil

<INITIAL>"//"
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "/"}
	//
	return int('/'), nil
<INITIAL>"*"
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "*"}
	//
	return int('*'), nil
<INITIAL>"%"
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "*"}
	//
	return int('/'), nil
<INITIAL,DefineExpression>[ \t\r]+
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()

<INITIAL>\n
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	y.streamData[y.streamDataIndex].lineNumber = 1
	y.streamData[y.streamDataIndex].rowNumber = 1
<DefineExpression>\n
	y.startCondIndex--
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	y.streamData[y.streamDataIndex].lineNumber = 1
	y.streamData[y.streamDataIndex].rowNumber = 1
<INITIAL>"//".*\n
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	y.streamData[y.streamDataIndex].lineNumber = 1
	y.streamData[y.streamDataIndex].rowNumber = 1

<INITIAL>"/"
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	//
	return int('/'), nil
")"
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	//
	return int(')'), nil
"("
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	//
	return int('('), nil
","
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	//
	return int(','), nil
<DefineExpression>\\
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	y.startCondIndex++
	y.startCond[y.startCondIndex]=DefineExpression
<INITIAL>abstract
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "abstract"}
	//
	return RWabstract, nil
<INITIAL>any
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "any"}
	//
	return RWany, nil
<INITIAL>alias
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "alias"}
	//
	return RWalias, nil
<INITIAL>attribute
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "attribute"}
	//
	return RWattribute, nil
<INITIAL>bitfield
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "bitfield"}
	//
	return RWbitfield, nil
<INITIAL>bitmask
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "bitmask"}
	//
	return RWbitmask, nil
<INITIAL>bitset
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier =  UnionData[string]{fileName, lineNumber, rowNumber, "bitset"}
	//
	return RWbitset, nil
<INITIAL>boolean
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "boolean"}
	//
	return RWboolean, nil
<INITIAL>bool
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "bool"}
	//
	return RWbool, nil
<INITIAL>case
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "case"}
	//
	return RWcase, nil
<INITIAL>component
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "component"}
	//
	return RWcomponent, nil
<INITIAL>connector
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "connector"}
	//
	return RWconnector, nil
<INITIAL>const
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "const"}
	//
	return RWconst, nil
<INITIAL>consumes
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "consumes"}
	//
	return RWconsumes, nil
<INITIAL>context
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "context"}
	//
	return RWcontext, nil
<INITIAL>default
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "default"}
	//
	return RWdefault, nil
<INITIAL>double
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "double"}
	//
	return RWdouble, nil
<INITIAL>exception
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "exception"}
	//
	return RWexception, nil
<INITIAL>emits
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "emits"}
	//
	return RWemits, nil
<INITIAL>enum
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "enum"}
	//
	return RWenum, nil
<INITIAL>eventtype
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "eventtype"}
	//
	return RWeventtype, nil
<INITIAL>factory
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "factory"}
	//
	return RWfactory, nil
<INITIAL>FALSE
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "FALSE"}
	//
	return RWFALSE, nil
<INITIAL>finder
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "finder"}
	//
	return RWfinder, nil
<INITIAL>fixed
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "fixed"}
	//
	return RWfixed, nil
<INITIAL>float
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "float"}
	//
	return RWfloat, nil
<INITIAL>float32
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "float32"}
	//
	return RWfloat32, nil
<INITIAL>float64
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "float64"}
	//
	return RWfloat64, nil
<INITIAL>getraises
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "getraises"}
	//
	return RWgetraises, nil
<INITIAL>home
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "home"}
	//
	return RWhome, nil
<INITIAL>import
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "import"}
	//
	return RWimport, nil
<INITIAL>in
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "in"}
	//
	return RWin, nil
<INITIAL>inout
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "inout"}
	//
	return RWinout, nil
<INITIAL>interface
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "interface"}
	//
	return RWinterface, nil
<INITIAL>local
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "local"}
	//
	return RWlocal, nil
<INITIAL>long
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "long"}
	//
	return RWlong, nil
<INITIAL>manages
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "manages"}
	//
	return RWmanages, nil
<INITIAL>map
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "map"}
	//
	return RWmap, nil
<INITIAL>mirrorport
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "mirrorport"}
	//
	return RWmirrorport, nil
<INITIAL>module
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "module"}
	//
	return RWmodule, nil
<INITIAL>multiple
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "multiple"}
	//
	return RWmultiple, nil
<INITIAL>native
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "native"}
	//
	return RWnative, nil
<INITIAL>Object
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "Object"}
	//
	return RWObject, nil
<INITIAL>octet
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "octet"}
	//
	return RWoctet, nil
<INITIAL>oneway
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "oneway"}
	//
	return RWoneway, nil
<INITIAL>out
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "out"}
	//
	return RWout, nil
<INITIAL>primarykey
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "primarykey"}
	//
	return RWprimarykey, nil
<INITIAL>private
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "private"}
	//
	return RWprivate, nil
<INITIAL>port
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "port"}
	//
	return RWport, nil
<INITIAL>porttype
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "porttype"}
	//
	return RWporttype, nil
<INITIAL>provides
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "provides"}
	//
	return RWprovides, nil
<INITIAL>public
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "public"}
	//
	return RWpublic, nil
<INITIAL>publishes
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "publishes"}
	//
	return RWpublishes, nil
<INITIAL>raises
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "raises"}
	//
	return RWraises, nil
<INITIAL>readonly
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "readonly"}
	//
	return RWreadonly, nil
<INITIAL>setraises
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "setraises"}
	//
	return RWsetraises, nil
<INITIAL>sequence
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "sequence"}
	//
	return RWsequence, nil
<INITIAL>short
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "short"}
	//
	return RWshort, nil
<INITIAL>string
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "string"}
	//
	return RWstring, nil
<INITIAL>struct
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "struct"}
	//
	return RWstruct, nil
<INITIAL>supports
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "supports"}
	//
	return RWsupports, nil
<INITIAL>TRUE
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "TRUE"}
	//
	return RWTRUE, nil
<INITIAL>truncatable
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "truncatable"}
	//
	return RWtruncatable, nil
<INITIAL>typedef
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "typedef"}
	//
	return RWtypedef, nil
<INITIAL>typeid
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "typeid"}
	//
	return RWtypeid, nil
<INITIAL>typename
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "typename"}
	//
	return RWtypename, nil
<INITIAL>typeprefix
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "typeprefix"}
	//
	return RWtypeprefix, nil
<INITIAL>unsigned
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "unsigned"}
	//
	return RWunsigned, nil
<INITIAL>uses
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "uses"}
	//
	return RWuses, nil
<INITIAL>ValueBase
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "ValueBase"}
	//
	return RWValueBase, nil
<INITIAL>valuetype
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "valuetype"}
	//
	return RWvaluetype, nil
<INITIAL>void
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "void"}
	//
	return RWvoid, nil
<INITIAL>byte
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "byte"}
	//
	return RWbyte, nil
<INITIAL>int8
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "int8"}
	//
	return RWint8, nil
<INITIAL>uint8
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "uint8"}
	//
	return RWuint8, nil
<INITIAL>int16
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "int16"}
	//
	return RWint16, nil
<INITIAL>int32
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "int32"}
	//
	return RWint32, nil
<INITIAL>int64
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "int64"}
	//
	return RWint64, nil
<INITIAL>uint16
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "uint16"}
	//
	return RWuint16, nil
<INITIAL>uint32
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "uint32"}
	//
	return RWuint32, nil
<INITIAL>uint64
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, "uint64"}
	//
	return RWuint64, nil
<INITIAL,DefineExpression>0|[1-9][0-9]*
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	//
	return y.onCreateInteger(lval)
<INITIAL,DefineExpression>0x[0-9A-Fa-f]+
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	//
	return y.OnHexValue(lval)
<INITIAL,DefineExpression>{F}
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	//
	return y.onCreateFloat(lval)
<INITIAL,DefineExpression>L?\"(\\.|[^\\"])*\"
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	//
	return y.OnStringLiteral(lval)
<INITIAL,DefineExpression>[@]?[a-zA-Z_][0-9a-zA-Z_]*
	s := __strings__.TrimSpace(string(y.buf[:len(y.buf)]))
	s = strings.Replace(s, "@", "", -1)
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, s}
	//
	return Identifier, nil
<INITIAL>"#"include
	y.startCondIndex++
	y.startCond[y.startCondIndex] = IncludeExpression
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
<INITIAL>"#"define
	y.startCondIndex++
	y.startCond[y.startCondIndex]=DefineExpression
	fileName = y.FileName()
	lineNumber = y.LineNumber()
	rowNumber = y.RowNumber()
	lval.Identifier = UnionData[string]{fileName, lineNumber, rowNumber, __strings__.TrimSpace(string(y.buf[:len(y.buf)]))}
	//
	return HashDefine, nil
<IncludeExpression>(.*)\n
	// to previous conditionstate
	y.startCondIndex--

	s := strings.TrimFunc(
		string(y.buf[:len(y.buf)-1]),
		func(r rune) bool {
			if b := unicode.IsSpace(r); b {
				return true
			}
			switch r {
			case '<', '>', '"', '\'':
				return true
			}
			return false
		})
	err = y.LoadFile(s)
	if err != nil {
		return 0, err
	}
	if c, err = y.getc(); err != nil {
		return 0, err
	}
	y.streamData[y.streamDataIndex].lineNumber++
	
%%
    y.empty = true
    return int(c), nil
}


