%{
package internal

//import __constant__ "go/constant"
import __token__ "go/token"
import __ast__ "go/ast"

type IYYParserHandler interface {
	IOptions
	Call(Definitions []IBaseDefinition) error
	Declare(definition IBaseDefinition) error
	Find(scopeName string) (ITypeSpec, error)
}

type UnionData[TData interface{}] struct {
	fileName   string
	lineNumber int
	rowNumber  int
	data       TData
}


%}
%token None
%token SingleComment
%token RWstring
%token MultiLineComment
%token WhiteSpace
%token HashLoadDefinition
%token HashPragma
%token HashDefine
%token HashUnDefine
%token RWbyte
%token Hashifdef
%token Hashifndef
%token Hashendif
%token Hashelse
%token HashInclude
%token ScopeOp // "::"
%token ShlOp // "<<"
%token ShrOp // ">>"
%token Annotation //"@annotation"
%token String_literal
%token Identifier
%token Integer_literal
%token Floating_pt_literal
%token Fixed_pt_literal
%token Character_literal
%token Wide_Character_literal
%token Wide_String_literal
%token Hex_literal
%token ErrorFileNotFound
%token RWabstract
%token RWany
%token RWalias
%token RWattribute
%token RWbitfield
%token RWbitmask
%token RWbitset
%token RWboolean
%token RWcase
//%token RWchar
%token RWcomponent
%token RWconnector
%token RWconst
%token RWconsumes
%token RWcontext
%token RWdefault
%token RWdouble
%token RWexception
%token RWemits
%token RWenum
%token RWeventtype
%token RWfactory
%token RWFALSE
%token RWfinder
%token RWfixed
%token RWfloat
%token RWfloat32
%token RWfloat64
%token RWgetraises
%token RWhome
%token RWimport
%token RWin
%token RWinout
%token RWinterface
%token RWlocal
%token RWlong
%token RWmanages
%token RWmap
%token RWmirrorport
%token RWmodule
%token RWmultiple
%token RWnative
%token RWObject
%token RWoctet
%token RWoneway
%token RWout
%token RWprimarykey
%token RWprivate
%token RWport
%token RWporttype
%token RWprovides
%token RWpublic
%token RWpublishes
%token RWraises
%token RWreadonly
%token RWsetraises
%token RWsequence
%token RWshort

%token RWstruct
%token RWsupports
%token RWTRUE
%token RWtruncatable
%token RWtypedef
%token RWtypeid
%token RWtypename
%token RWtypeprefix
%token RWunsigned

%token RWuses
%token RWValueBase
%token RWvaluetype
%token RWvoid
//%token RWwchar
%token RWint8
%token RWuint8
%token RWint16
%token RWint32
%token RWint64
%token RWuint16
%token RWuint32
%token RWuint64
%token RWbool



%type <Identifiers> any_declaratorsPlus IdentifierPlus IdentifierStar

%type <Identifier> '+' '-' '/' '*' '%' HashDefine '{' '}' '\\'
%type <Identifier> any_declarator scoped_name Identifier RWshort RWint16 RWuint16 RWlong RWint32 RWint64 RWint8 RWuint8 RWsequence RWboolean RWbool
%type <Identifier> RWbyte RWuint32 RWfloat RWfloat32 RWfloat64 RWdouble RWunsigned RWuint64 simple_declarator RWconst RWstruct RWenum RWtypedef RWstring
%type <Definition> struct_def  struct_dcl constr_type_dcl type_dcl definition const_dcl enum_dcl typedef_dcl  bitset_dcl bitmask_dcl DefineExpression
%type <Definition> except_dcl interface_dcl value_dcl type_id_dcl type_prefix_dcl import_dcl component_dcl home_dcl event_dcl porttype_dcl connector_dcl template_module_dcl template_module_inst annotation_dcl module_dcl
%type <Node> Integer_literal literal primary_expr unary_expr mult_expr add_expr shift_expr and_expr xor_expr or_expr const_expr positive_int_const
%type <Node> Floating_pt_literal Fixed_pt_literal Character_literal Wide_Character_literal boolean_literal String_literal Wide_String_literal
%type <r> unary_operator
%type <TypeSpec> const_type integer_type base_type_spec simple_type_spec type_spec memberType subMmemberType
%type <TypeSpec> signed_short_int  signed_long_int  signed_longlong_int  unsigned_short_int
%type <TypeSpec> unsigned_long_int  signed_tiny_int  unsigned_tiny_int  signed_int
%type <TypeSpec> floating_pt_type fixed_pt_const_type  boolean_type octet_type string_type
%type <TypeSpec> any_type object_type value_base_type
%type <TypeSpec> unsigned_int unsigned_longlong_int sequence_type template_type_spec fixed_pt_type map_type
%type <Declarator> declarator
%type <Declarators> declaratorPlus declarators
%type <Member> member
%type <Members> memberPlus memberStar
%type <Enumerator> enumerator
%type <Enumerators> enumerators

%type <Definitions> definitionPlus
%type <TypeDeclarator> type_declarator
%union{
	TypeDeclarator UnionData[TypeDeclarator]
	TypeSpec UnionData[ITypeSpec]
	Identifier UnionData[string]
	Identifiers UnionData[[]string]
	Definition UnionData[IBaseDefinition]
	Definitions UnionData[[]IBaseDefinition]
	Declarator UnionData[Declarator]
	Declarators UnionData[[]Declarator]
	//rv UnionData[__ast__.Expr]
	r UnionData[rune]
	Member UnionData[FitMember]
	Members UnionData[[]FitMember]
	Enumerator UnionData[Enumerator]
	Enumerators UnionData[[]Enumerator]
	Node UnionData[__ast__.Expr]
}

%start specification
%%

//(1)
specification:
	definitionPlus
	{
		YYrcvr.data.Call($1.data)
	}

//(2)(71)//(98)//(111)//(133)//(144)//(153)//(171)//(184)//(218)
definitionPlus:
	{}
	|definition
	{
	 	$$ = UnionData[[]IBaseDefinition]{$1.fileName, $1.lineNumber, $1.rowNumber, append($$.data, $1.data)}
	 	YYrcvr.data.Declare($1.data)
	 	$1 = UnionData[IBaseDefinition]{"", -1, -1, nil}
	}
	| definitionPlus definition
	{
		$$ = UnionData[[]IBaseDefinition]{$1.fileName, $1.lineNumber, $1.rowNumber, append($1.data, $2.data)}
		YYrcvr.data.Declare($2.data)
		$2 = UnionData[IBaseDefinition]{"", -1, -1, nil}
	}
definition:
	module_dcl ';'			{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| const_dcl ';'			{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| type_dcl ';'			{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| except_dcl ';'		{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| interface_dcl ';'		{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| value_dcl ';'			{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| type_id_dcl ';'		{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| type_prefix_dcl ';'		{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| import_dcl ';'		{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| component_dcl ';'		{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| home_dcl ';'			{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| event_dcl ';'			{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| porttype_dcl ';'		{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| connector_dcl ';'		{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| template_module_dcl ';'	{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| template_module_inst ';'	{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| annotation_dcl ';'		{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }
	| DefineExpression		{ $$ = $1; $1 = UnionData[IBaseDefinition]{"", -1, -1, nil} }

// custom rule
DefineExpression:
	// boolean
	HashDefine Identifier{
		$$ = UnionData[IBaseDefinition]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
		  	&HashDefineExpression{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$2.data,
				nil,
			},
		}
		$1 = UnionData[string]{"", -1, -1, ""}
		$2 = UnionData[string]{"", -1, -1, ""}
	}
	// some exprssion
	|HashDefine Identifier const_expr
	{
		$$ = UnionData[IBaseDefinition]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
		  	&HashDefineExpression{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$2.data,
				$3.data,
			},
		}
		$1 = UnionData[string]{"", -1, -1, ""}
		$2 = UnionData[string]{"", -1, -1, ""}
		$3 = UnionData[__ast__.Expr]{"", 0, 0, nil}
	}
	// function
	| HashDefine Identifier '(' IdentifierStar ')' hp
	{}
hp :
	const_expr{}
	| Identifier '(' macro_params ')'{}
macro_params:
	{}
	| macro_param{}
	| macro_params ',' macro_param {}
macro_param:
	Identifier{}
	| const_expr{}
	| hp {}

IdentifierStar:
	{
	}
	| Identifier
	{
		$$ = UnionData[[]string]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($$.data, $1.data),
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	| IdentifierStar ',' Identifier
	{
		$$ = UnionData[[]string]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($1.data, $3.data),
		}
		$3 = UnionData[string]{"", 0, 0, ""}
	}
//(3)
module_dcl:
	RWmodule Identifier '{' definitionPlus '}'
	{
	}
	| RWmodule Identifier '{'  '}'
	{
	}
//(4)
scoped_namePlus:
	scoped_name
	{
	}
	| scoped_namePlus ',' scoped_name
	{
	}
scoped_name:
	Identifier
	{
	}
	| ScopeOp Identifier
	{
	}
	| scoped_name ScopeOp Identifier
	{
	}
//(5)
const_dcl:
	RWconst const_type Identifier '=' const_expr
	{
		$$ = UnionData[IBaseDefinition]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
		  	&RwConstant{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$3.data,
				$2.data,
				$5.data,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
		$2 = UnionData[ITypeSpec]{"", -1, -1, nil}
		$3 = UnionData[string]{"", 0, 0, ""}
		$5 = UnionData[__ast__.Expr]{"", 0, 0, nil}
	}
//(6)
const_type:
	integer_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| floating_pt_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| fixed_pt_const_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| boolean_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| octet_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| string_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}

	| scoped_name
	{
		if v, err := YYrcvr.data.Find($1.data);  err == nil{
			$$ = UnionData[ITypeSpec]{"", -1, -1, v}
		} else {
			YYlex.Error(err.Error())
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
//(7)
const_expr:
	or_expr
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
//(8)
or_expr:
	xor_expr
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| or_expr '|' xor_expr
	{
	}
//(9)
xor_expr:
	and_expr
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| xor_expr '^' and_expr
	{
	}

//(10)
and_expr:
	shift_expr
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| and_expr '&' shift_expr
	{
	}

//(11)
shift_expr:
	add_expr
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| shift_expr ShrOp add_expr
	{
	}
	| shift_expr ShlOp add_expr
	{
	}
//(12)
add_expr:
	mult_expr
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| add_expr '+' mult_expr
	{
		$$ = UnionData[__ast__.Expr]{
			$2.fileName,
			$2.lineNumber,
			$2.rowNumber,
			&__ast__.BinaryExpr{
				X:     $1.data,
				OpPos: 0,
				Op:    __token__.ADD,
				Y:     $3.data,
			},
		}
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
		$3 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| add_expr '-' mult_expr
	{
		$$ = UnionData[__ast__.Expr]{
			$2.fileName,
			$2.lineNumber,
			$2.rowNumber,
			&__ast__.BinaryExpr{
				X:     $1.data,
				OpPos: 0,
				Op:    __token__.SUB,
				Y:     $3.data,
			},
		}
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
		$3 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
//(13)
mult_expr:
	unary_expr
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| mult_expr '*' unary_expr
	{
		$$ = UnionData[__ast__.Expr]{
			$2.fileName,
			$2.lineNumber,
			$2.rowNumber,
			&__ast__.BinaryExpr{
				X:     $1.data,
				OpPos: 0,
				Op:    __token__.MUL,
				Y:     $3.data,
			},
		}
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
		$3 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| mult_expr '/' unary_expr
	{
		$$ = UnionData[__ast__.Expr]{
			$2.fileName,
			$2.lineNumber,
			$2.rowNumber,
			&__ast__.BinaryExpr{
				X:     $1.data,
				OpPos: 0,
				Op:    __token__.QUO,
				Y:     $3.data,
			},
		}
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
		$3 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| mult_expr '%' unary_expr
	{
		$$ = UnionData[__ast__.Expr]{
			$2.fileName,
			$2.lineNumber,
			$2.rowNumber,
			&__ast__.BinaryExpr{
				X:     $1.data,
				OpPos: 0,
				Op:    __token__.REM,
				Y:     $3.data,
			},
		}
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
		$3 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
//(14)
unary_expr:
	unary_operator primary_expr
	{
		switch $1.data {
			case rune('-'):
				$$ = UnionData[__ast__.Expr]{
					$2.fileName,
					$2.lineNumber,
					$2.rowNumber,
					&__ast__.UnaryExpr{
						OpPos: 0,
						Op:    __token__.SUB,
						X:     	$2.data,
					},
				}
				$2 = UnionData[__ast__.Expr]{"", -1, -1, nil}
				break
			case rune('+') :
				$$ = $2
				$2 = UnionData[__ast__.Expr]{"", -1, -1, nil}
				break
			case rune('~') :
				$$ = UnionData[__ast__.Expr]{
					$2.fileName,
					$2.lineNumber,
					$2.rowNumber,
					&__ast__.UnaryExpr{
						OpPos: 0,
						Op:    __token__.Token('~'),
						X:     	$2.data,
					},
				}
				$2 = UnionData[__ast__.Expr]{"", -1, -1, nil}
				break
		}
	}
	| primary_expr
	{
		$$ = $1
	}
//(15)
unary_operator:
	'-'
	{
		$$ = UnionData[rune]{"", -1, -1, '-'}
	}
	| '+'
	{
		$$ = UnionData[rune]{"", -1, -1, '+'}
	}
	| '~'
	{
		$$ = UnionData[rune]{"", -1, -1, '~'}
	}
//(16)
primary_expr:
	scoped_name
	{
	}
	| literal
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| '(' const_expr ')'
	{
		$$ = $2
		$2 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
//(17)
literal:
	Integer_literal
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| Floating_pt_literal
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| Fixed_pt_literal
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| Character_literal
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| Wide_Character_literal
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| boolean_literal
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| String_literal
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
	| Wide_String_literal
	{
		$$ = $1
		$1 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}

//(18)
boolean_literal:
	RWTRUE
	{
	}
	| RWFALSE
	{
	}
//(19)
positive_int_const:
	const_expr
	{
	}
//(20)
type_dcl:
	constr_type_dcl
	{
		$$ = $1
		$1 = UnionData[IBaseDefinition]{"", -1, -1, nil}
	}
	| native_dcl
	{
	}
	| typedef_dcl
	{
		$$ = $1
		$1 = UnionData[IBaseDefinition]{"", -1, -1, nil}
	}
//(21)
type_spec:
	simple_type_spec
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	|template_type_spec
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
//(22)
simple_type_spec:
	base_type_spec
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| scoped_name
	{
		if v, err := YYrcvr.data.Find($1.data);  err == nil{
			$$ = UnionData[ITypeSpec]{"", -1, -1, v}
		} else {
			YYlex.Error(err.Error())
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
//(23)(69)//(117)//(131)
base_type_spec:
	integer_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
        }
	| floating_pt_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| boolean_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| octet_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| any_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| object_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| value_base_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}

//(24)
floating_pt_type:
	RWfloat
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$1.data,
				true,
				0,
				 false,
				 false,
		 	},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	|RWfloat32
         	{
         		$$ = UnionData[ITypeSpec]{
         			$1.fileName,
         			$1.lineNumber,
         			$1.rowNumber,
         			&ConstType{
         				fileInformation{
         					$1.fileName,
         					$1.lineNumber,
         					$1.rowNumber,
         				},
         				YYrcvr.data,
         				$1.data,
         				true,
         				0,
         				 false,
         				 false,
         		 	},
         		}
         		$1 = UnionData[string]{"", 0, 0, ""}
         	}
	|RWfloat64
        	{
        		$$ = UnionData[ITypeSpec]{
        			$1.fileName,
        			$1.lineNumber,
        			$1.rowNumber,
        			&ConstType{
        				fileInformation{
        					$1.fileName,
        					$1.lineNumber,
        					$1.rowNumber,
        				},
        				YYrcvr.data,
        				$1.data,
        				true,
        				0,
        				 false,
        				 false,
        		 	},
        		}
        		$1 = UnionData[string]{"", 0, 0, ""}
        	}

	| RWdouble
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,

		&ConstType{
			fileInformation{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			},
			YYrcvr.data,
			$1.data, true, 0, false, false}}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	| RWlong RWdouble {
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,

		&ConstType{
			fileInformation{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			},
			YYrcvr.data,
			$2.data,
			true, 0, false, true}}
		$1 = UnionData[string]{"", 0, 0, ""}
		$2 = UnionData[string]{"", 0, 0, ""}
	}
//(25)
integer_type:
	signed_int
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| unsigned_int
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}

//(26)//(206)
signed_int:
	signed_short_int
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| signed_long_int
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| signed_longlong_int
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| signed_tiny_int
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
//(27)//(210)
signed_short_int:
	RWshort
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			 &ConstType{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	| RWint16
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&ConstType{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
//(28)//(211)
signed_long_int:
	RWlong
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	| RWint32
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}

//(29)//(212)
signed_longlong_int:
	RWlong RWlong
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$2.data, true, 0, false, true,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
		$2 = UnionData[string]{"", 0, 0, ""}
	}
	|RWint64
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			&ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false}}
		$1 = UnionData[string]{"", 0, 0, ""}
	}

//(30)//(207)
unsigned_int:
	unsigned_short_int
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| unsigned_long_int
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| unsigned_longlong_int
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	|unsigned_tiny_int
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}

//(31)//(213)
unsigned_short_int:
	RWunsigned RWshort
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			&ConstType{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				YYrcvr.data,
				$2.data, true, 0, true, false}}
		$1 = UnionData[string]{"", 0, 0, ""}
		$2 = UnionData[string]{"", 0, 0, ""}
	}
	|RWuint16
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			 &ConstType{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false}}
		$1 = UnionData[string]{"", 0, 0, ""}
	}

//(32)//(214)
unsigned_long_int:
	RWunsigned RWlong
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			&ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$2.data, true, 0, true, false},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
		$2 = UnionData[string]{"", 0, 0, ""}
	}
	| RWuint32
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			 &ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}

//(33)//(215)
unsigned_longlong_int:
	RWunsigned RWlong RWlong
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			 &ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
			YYrcvr.data,
			$3.data, true, 0, true, true}}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	| RWuint64
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,

			&ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false}}
		$1 = UnionData[string]{"", 0, 0, ""}
	}

boolean_type:
	RWboolean
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,

			&ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false}}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	|RWbool
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,

			&ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false}}
		$1 = UnionData[string]{"", 0, 0, ""}
	}

//(37)
octet_type:
	RWoctet
	{
	}
//(38)//(197)
template_type_spec:
	sequence_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| string_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| fixed_pt_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| map_type
	{
		$$ = $1
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
//(39)
sequence_type:
	RWsequence '<' type_spec ',' positive_int_const '>'
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&RwSequence{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$3.data,
				$5.data,
			},
		}
		$3 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
	| RWsequence '<' type_spec '>'
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&RwSequence{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$3.data,
				nil,
			},
		}
		$3 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}
//(40)
string_type:
	RWstring
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			 &ConstType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false}}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
//(42)
fixed_pt_type:
	RWfixed '<' positive_int_const ',' positive_int_const '>'
	{
	}
//(43)
fixed_pt_const_type:
	RWfixed
	{
	}
//(44)//(198)
constr_type_dcl:
	struct_dcl
	{
		$$ = $1
		$1 = UnionData[IBaseDefinition]{"", -1, -1, nil}
	}
	| enum_dcl
	{
		$$ = $1
		$1 = UnionData[IBaseDefinition]{"", -1, -1, nil}
	}
	| bitset_dcl
	{
	}
	| bitmask_dcl
	{
	}

//(45)
struct_dcl:
	struct_def
	{
		$$ = $1
		$1 = UnionData[IBaseDefinition]{"", -1, -1, nil}
	}
//(46)//(195)
struct_def:
	RWstruct '<' literal ',' String_literal '>' Identifier    '{' memberStar '}'
	{
		$$ = UnionData[IBaseDefinition] {
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&fitStruct {
				fileInformation {
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$3.data,
				$5.data,
				$7.data,
				$9.data,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
		$3 = UnionData[__ast__.Expr]{"", 0, 0, nil}
		$5 = UnionData[__ast__.Expr]{"", 0, 0, nil}
		$7 = UnionData[string]{"", 0, 0, ""}
		$9 = UnionData[[]FitMember]{"", -1, -1, $9.data[:0]}
	}
//(47)
memberPlus:
	member
	{
		$$ = UnionData[[]FitMember]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
		 	append($$.data, $1.data),
		}
		$1 = UnionData[FitMember]{"", -1, -1, FitMember{}}
	}
	|memberPlus member
	{
		$$ = UnionData[[]FitMember]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($1.data, $2.data),
		}
		$2 = UnionData[FitMember]{"", -1, -1, FitMember{}}
	}
memberStar:
	{
	}
	|memberPlus
	{
		$$ = $1
		$1 = UnionData[[]FitMember]{"", -1, -1, $1.data[:0]}
	}


memberType:
	type_spec '<'
		String_literal ','
		Integer_literal ','
		const_expr ','
		String_literal ','
		const_expr ','
		const_expr ','
		String_literal ','
		String_literal ','
		String_literal ','
		String_literal ','
		String_literal ','
		String_literal ','
		String_literal ','
		String_literal
	'>'
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&MemberType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$1.data,
				$3.data,
				$5.data,
				$7.data,
				$9.data,
				$11.data,
				$13.data,
				$15.data,
				$17.data,
				$19.data,
				$21.data,
				$23.data,
				$25.data,
				$27.data,
				$29.data,
			},
		}
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}

subMmemberType:
	type_spec '<'
		String_literal ','
		String_literal ','
		unary_expr ','
		unary_expr ','
		String_literal ','
		String_literal ','
		String_literal ','
		String_literal ','
		String_literal ','
		String_literal ','
		String_literal ','
		String_literal
	'>'
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&SubMemberType{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$1.data,
			},
		}
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
	}

subMmemberPlus:
	subMmember
	{

	}
	| subMmemberPlus  subMmember

subMmember:
	subMmemberType declarator ';'
	{
	}


member:
	memberType declarator '{' subMmemberPlus '}' ';'
	{

	}
	| memberType declarator ';'
	{
		$$ = UnionData[FitMember]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			FitMember{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				$1.data, $2.data},
		}
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
		$2 = UnionData[Declarator]{"", -1, -1, Declarator{}}

	}
//(50)
//(51)//(196)
switch_type_spec:
	integer_type
	{
	}
	| boolean_type
	{
	}
	| scoped_name
	{
	}
	| octet_type
	{
	}
//(52)
switch_body:
	casePlus
	{
	}
casePlus:
	case
	{
	}
	|casePlus case
	{
	}
//(53)
case:
	case_labelPlus element_spec ';'
	{
	}

case_labelPlus:
	case_label
	{
	}
	|case_labelPlus case_label
	{
	}
//(54)
case_label:
	RWcase const_expr ':'
	{
	}
	| RWdefault ':'
	{
	}
//(55)
element_spec:
	type_spec declarator
	{
	}
//(57)
enum_dcl:
	RWenum '<' const_type ',' Integer_literal  '>'  Identifier  '{'  '}' {
		$$ = UnionData[IBaseDefinition]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&RwEnum{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$7.data,
				nil,
				$3.data,
				$5.data,
			}}
		$1 = UnionData[string]{"", 0, 0, ""}
		$3 = UnionData[ITypeSpec]{"", 0, 0, nil}
		$5 = UnionData[__ast__.Expr]{"", 0, 0, nil}
		$7 = UnionData[string]{"", 0, 0, ""}

	}
	| RWenum '<' const_type ',' Integer_literal  '>'  Identifier  '{' enumerators '}'
	{
		$$ = UnionData[IBaseDefinition]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&RwEnum{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$7.data,
				$9.data,
				$3.data,
				$5.data,
			}}
		$1 = UnionData[string]{"", 0, 0, ""}
		$3 = UnionData[ITypeSpec]{"", 0, 0, nil}
		$5 = UnionData[__ast__.Expr]{"", 0, 0, nil}
		$7 = UnionData[string]{"", 0, 0, ""}
		$9 = UnionData[[]Enumerator]{"", -1, -1, $9.data[:0]}
	}
enumerators:
	enumerator
	{
		$$ = UnionData[[]Enumerator]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($$.data, $1.data),
		}
	}
	| enumerators ',' enumerator
	{
		$$ = UnionData[[]Enumerator]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($1.data, $3.data),
		}
	}
//(58)
enumerator:
	Identifier
	{
		$$ = UnionData[Enumerator]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			Enumerator{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				IdentifierType($1.data),
				nil,
			}}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	| Identifier '=' const_expr
	{
		$$ = UnionData[Enumerator]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			Enumerator{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				IdentifierType($1.data),
				$3.data,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
		$3 = UnionData[__ast__.Expr]{"", -1, -1, nil}
	}
//(59)
array_declarator:
	simple_declarator fixed_array_sizePlus
	{
	}
//(60)
fixed_array_sizePlus:
	fixed_array_size
	{
	}
	|fixed_array_sizePlus fixed_array_size
	{
	}
fixed_array_size:
	'[' positive_int_const ']'
	{
	}
//(61)
native_dcl:
	RWnative simple_declarator
	{
	}
//(62)
simple_declaratorPlus:
	simple_declarator
	{
	}
	|simple_declaratorPlus ',' simple_declarator
	{
	}

simple_declarator:
	Identifier
	{
		$$ = $1
		$1 = UnionData[string]{"", 0, 0, ""}
	}
//(63)
typedef_dcl:
	RWtypedef type_declarator
	{
		v := make([]IdentifierType, len($2.data.Items), len($2.data.Items))
		for i, item := range $2.data.Items{
			v[i] = IdentifierType(item)
		}
		$$ = UnionData[IBaseDefinition]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			&RwTypedef{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				$2.data.TypeSpec,
				v,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
		$2 = UnionData[TypeDeclarator]{"", -1, -1, TypeDeclarator{}}
	}
//(64)
type_declarator:
	simple_type_spec any_declaratorsPlus
	{
		$$ = UnionData[TypeDeclarator]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			TypeDeclarator{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				$1.data,
				$2.data,
			},
		}
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
		$2 = UnionData[[]string]{"", -1, -1, $2.data[:0]}

	}
	| template_type_spec any_declaratorsPlus
	{
		$$ = UnionData[TypeDeclarator]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			TypeDeclarator{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				$1.data,
				$2.data,
			},
		}
		$1 = UnionData[ITypeSpec]{"", -1, -1, nil}
		$2 = UnionData[[]string]{"", -1, -1, $2.data[:0]}

	}
	| constr_type_dcl any_declaratorsPlus
	{
		$$ = UnionData[TypeDeclarator]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			TypeDeclarator{
				fileInformation{
					$1.fileName,
					$1.lineNumber,
					$1.rowNumber,
				},
				$1.data,
				$2.data,
			},
		}
		$1 = UnionData[IBaseDefinition]{"", -1, -1, nil}
		$2 = UnionData[[]string]{"", -1, -1, $2.data[:0]}
	}
//(65)
any_declaratorsPlus:
	any_declarator
	{
		$$ = UnionData[[]string]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($$.data, $1.data),
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	| any_declaratorsPlus ',' any_declarator
	{
		$$ = UnionData[[]string]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($1.data, $3.data),
		}
		$3 = UnionData[string]{"", 0, 0, ""}
	}

//(66)
any_declarator:
	simple_declarator
	{
		$$ = $1
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	| array_declarator
	{

	}
//(67)
declarators:
	declaratorPlus
	{
		$$ = $1
		$1 = UnionData[[]Declarator]{"", -1, -1, $1.data[:0]}
	}
//(68)//(217)
declaratorPlus:
	declarator
	{
		$$ = UnionData[[]Declarator]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($$.data, $1.data),
		}
		$1 = UnionData[Declarator]{"", -1, -1, Declarator{}}
	}
	| declaratorPlus ',' declarator
	{
		$$ = UnionData[[]Declarator]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($1.data, $3.data),
		}
		$3  = UnionData[Declarator]{"", -1, -1, Declarator{}}
	}
declarator:
	simple_declarator
	{
		$$ = UnionData[Declarator]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			Declarator{
				IdentifierType($1.data),
				false,
				nil,
			}}
	}
	|array_declarator
	{
        }
//(70)
any_type:
	RWany
	{
	}
//(72)
except_dcl:
	RWexception Identifier '{' memberStar '}'
	{
	}
//(73)
interface_dcl:
	interface_def
	{
	}
	|
	interface_forward_dcl
	{
	}
//(74)
interface_def:
	interface_header '{' interface_body '}'
	{
	}
//(75)
interface_forward_dcl:
	interface_kind Identifier
	{

	}
//(76)
interface_header:
	interface_kind Identifier
	{
	}

	|
	interface_kind Identifier interface_inheritance_spec
	{
	}
//(77)//(119)//(129)
interface_kind:
	RWinterface
	{
	}
	| RWlocal RWinterface
	{
	}
	|RWabstract RWinterface
	{
	}

//(78)
interface_inheritance_spec:
	':' interface_namePlus
	{
	}
//(79)
interface_namePlus:
	interface_name
	{
	}
	| interface_namePlus ',' interface_name
	{
	}
interface_name:
	scoped_name
	{
	}
//(80)
interface_body:
	exportStar
	{
	}
//(81)//(97)//(112)
exportStar:
	{
	}
	| exportPlus
	{
	}

exportPlus:
	exportPlus export {
	}
	| export {
	}
export:
	op_dcl ';'
	{
	}
	| attr_dcl ';'
	{
	}
	| type_dcl ';'
	{
	}
	| const_dcl ';'
	{
	}
	| except_dcl ';'
	{
	}
	| type_id_dcl ';'
	{
	}
	| type_prefix_dcl ';'
	{
	}
	| import_dcl ';'
	{
	}
	| op_oneway_dcl ';'
	{
	}
	| op_with_context ';'
	{
	}

//(82)
op_dcl:
	op_type_spec Identifier '('  ')'
	{
	}
	|op_type_spec Identifier '('  ')'  raises_expr
	{
	}
	|op_type_spec Identifier '('  parameter_dcls ')'
	{
	}
	|op_type_spec Identifier '('  parameter_dcls ')'  raises_expr
	{
	}
//(83)
op_type_spec:
	type_spec
	{
	}
	| RWvoid
	{
	}
//(84)
parameter_dcls:
	param_dclPlus
	{

	}
//(85)
param_dclPlus:
	param_dcl {
	}
	| param_dclPlus ',' param_dcl
	{
	}
param_dcl:
	param_attribute type_spec simple_declarator
	{
	}
//(86)
param_attribute:
	RWin
	{
	}
	| RWout
	{
	}
	| RWinout
	{
	}
//(87)
raises_expr:
	RWraises '(' scoped_namePlus ')'
	{
	}
//(88)
attr_dcl:
	readonly_attr_spec
	{
	}
	| attr_spec
	{
	}
//(89)
readonly_attr_spec:
	RWreadonly RWattribute type_spec readonly_attr_declarator
	{
	}
//(90)
readonly_attr_declarator:
	simple_declarator raises_expr
	{
	}
	| simple_declaratorPlus
	{
	}
//(91)
attr_spec:
	RWattribute type_spec attr_declarator
	{
	}
//(92)
attr_declarator:
	simple_declarator attr_raises_expr
	{
	}
	| simple_declaratorPlus
	{
	}
//(93)
attr_raises_expr:
	get_excep_expr set_excep_expr
	{
	}
	| get_excep_expr
	{
	}
	| set_excep_expr
	{
	}
//(94)
get_excep_expr:
	RWgetraises exception_list
	{
	}
//(95)
set_excep_expr:
	RWsetraises exception_list
	{
	}
//(96)
exception_list:
 	'(' scoped_namePlus ')'
	{
	}
//(99)//(125)
value_dcl:
	value_def
	{
	}
	| value_forward_dcl
	{
	}
	| value_box_def
	{
	}
// TODO: FIX
//	| value_abs_def
//	{
//		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
		//yaccLogger.Log("value_dcl/value_abs_def")
//}
//		$$ = $1
//	}
//(100)
value_def:
	value_header '{' value_elementStar '}'
	{
	}
//(101)
value_header:
	value_kind Identifier value_inheritance_spec
	{
	}
	| value_kind Identifier
	{
	}
//(102) //(128) //(127)
value_kind:
	RWvaluetype
	{
	}
	|RWabstract RWvaluetype
	{
	}

//(103)//(130)
value_inheritance_spec:
	{
	}
	|   RWsupports interface_name
	{
	}
	|  ':' value_name
	{
	}
	|  ':' value_name RWsupports interface_name
	{
	}
	|  ':' value_namePlus
	{
	}
	|  ':' value_namePlus RWsupports interface_namePlus
	{
	}
	|  ':' RWtruncatable value_namePlus
	{
	}
	|  ':' RWtruncatable value_namePlus RWsupports interface_namePlus
	{
	}

//(104)
value_namePlus:
	value_name
	{
	}
	| value_namePlus ',' value_name
	{
	}
value_name:
	scoped_name
	{
	}
//(105)
value_elementPlus:
	value_element
	{
	}
	|value_elementPlus value_element
	{
	}

value_elementStar:
	{
	}
	|value_elementPlus
	{
	}

value_element:
	export
	{
	}
	| state_member
	{
	}
	| init_dcl
	{
	}
//(106)
state_member:
	RWpublic type_spec declarators ';'
	{
	}
	| RWprivate type_spec declarators ';'
	{
	}
//(107)
init_dcl:
	RWfactory Identifier '('  ')'  ';'
	{
	}
	| RWfactory Identifier '('  ')'  raises_expr  ';'
	{
	}
	| RWfactory Identifier '('  init_param_dcls  ')'  ';'
	{
	}
	| RWfactory Identifier '('  init_param_dcls  ')'  raises_expr  ';'
	{
	}
//(108)
init_param_dcls:
	init_param_dclPlus
	{
	}
//(109)
init_param_dclPlus:
	init_param_dcl
	{
	}
	|init_param_dclPlus ',' init_param_dcl
	{
	}
init_param_dcl:
	RWin type_spec simple_declarator
	{
	}
//(110)
value_forward_dcl:
	value_kind Identifier
	{
	}
//(113)
type_id_dcl: RWtypeid scoped_name String_literal {}
//(114)
type_prefix_dcl:
	RWtypeprefix scoped_name String_literal
	{
	}
//(115)
import_dcl:
	RWimport imported_scope
	{
	}
//(116)
imported_scope:
	scoped_name
	{
	}
	| String_literal
	{
	}
//(118)
object_type:
	RWObject
	{
	}
//(120)
op_oneway_dcl: RWoneway RWvoid Identifier '('  in_parameter_dclsStar  ')' {}
//(121)
in_parameter_dclsStar:
	 {
	 }
	| in_parameter_dcls
	{
	}
	| in_parameter_dclsStar in_parameter_dcls
	{
	}
in_parameter_dcls:
	in_param_dclPlus
	{
	}
//(122)
in_param_dclPlus:
	in_param_dcl
	{
	}
	| in_param_dclPlus ',' in_param_dcl
	{
	}
in_param_dcl:
	RWin type_spec simple_declarator
	{
	}
//(123)
op_with_context:
 	 op_oneway_dcl context_expr
 	 {
 	 }
 	| op_dcl context_expr
 	{
 	}
//(124)
String_literalPlus:
	String_literal
	{
	}
	| String_literalPlus ',' String_literal
	{
	}
context_expr:
	RWcontext '(' String_literalPlus  ')'
	{
	}
//(126)
value_box_def:
	value_kind Identifier type_spec
	{
	}
//(132)
value_base_type:
	RWValueBase
	{
	}
//(134)
component_dcl:
	component_def
	{
	}
	| component_forward_dcl
	{
	}
//(135)
component_forward_dcl:
	RWcomponent Identifier
	{
	}
//(136)
component_def:
	component_header '{' component_body '}'
	{
	}

//(137) //(154)
component_header:
	RWcomponent Identifier
	{
	}
	| RWcomponent Identifier component_inheritance_spec
	{
	}
	| RWcomponent Identifier supported_interface_spec
	{
	}
	| RWcomponent Identifier component_inheritance_spec supported_interface_spec
	{
	}

//(138)
component_inheritance_spec:
	':' scoped_name
	{
	}
//(139)
component_body:
	component_exportStar
	{
	}
//(140)//(156)//(179)
component_exportStar:
	{
	}
	| component_export
	{
	}
	| component_exportStar component_export
	{
	}
component_export:
	provides_dcl ';'
	{
	}
	| uses_dcl ';'
	{
	}
	| attr_dcl ';'
	{
	}
	| emits_dcl ';'
	{
	}
	| publishes_dcl ';'
	{
	}
	| consumes_dcl ';'
	{
	}
	| port_dcl ';'
	{
	}

//(141)
provides_dcl:
	RWprovides interface_type Identifier
	{
	}
//(142)//(157)
interface_type:
	scoped_name
	{
	}
	| RWObject
	{
	}

//(143)//(158)
uses_dcl:
	RWuses interface_type Identifier
	{
	}
	|RWuses RWmultiple interface_type Identifier
	{
	}

//(145)
home_dcl:
	home_header '{' home_body '}'
	{
	}
//(146)//(162)
home_header:
	RWhome Identifier RWmanages scoped_name
	{
	}
	| RWhome Identifier home_inheritance_spec RWmanages scoped_name
	{
	}
	| RWhome Identifier RWmanages scoped_name
	{
	}
	| RWhome Identifier RWmanages scoped_name primary_key_spec
	{
	}
	| RWhome Identifier   supported_interface_spec RWmanages scoped_name
	{
	}
	| RWhome Identifier   supported_interface_spec RWmanages scoped_name primary_key_spec
	{
	}
	| RWhome Identifier home_inheritance_spec   RWmanages scoped_name
	{
	}
	| RWhome Identifier home_inheritance_spec   RWmanages scoped_name primary_key_spec
	{
	}
	| RWhome Identifier home_inheritance_spec   supported_interface_spec RWmanages scoped_name
	{
	}
	| RWhome Identifier home_inheritance_spec   supported_interface_spec RWmanages scoped_name primary_key_spec
	{
	}


//(147)
home_inheritance_spec:
	':' scoped_name
	{
	}
//(148)
home_body:
	home_exportStar
	{
	}
//(149)//(164)
home_exportStar:
	{
	}
	| home_export
	{
	}
	| home_exportStar home_export
	{
	}
home_export:
	export
	{
	}
	| factory_dcl ';'
	{
	}
	| finder_dcl ';'
	{
	}

//(150)
factory_dcl:
	RWfactory Identifier '('  ')'
	{
	}
	|RWfactory Identifier '('  ')'  raises_expr
	{
	}
	|RWfactory Identifier '('  factory_param_dcls  ')'
	{
	}
	|RWfactory Identifier '('  factory_param_dcls  ')'  raises_expr
	{
	}

//(151)
factory_param_dcls:
	factory_param_dclPlus
	{
	}
//(152)
factory_param_dclPlus:
	factory_param_dcl
	{
	}
	| factory_param_dclPlus ',' factory_param_dcl
	{
	}
factory_param_dcl:
	RWin type_spec simple_declarator
	{
	}
//(155)
supported_interface_spec:
	RWsupports scoped_namePlus
	{
	}
//(159)
emits_dcl:
	RWemits scoped_name Identifier
	{
	}
//(160)
publishes_dcl:
	RWpublishes scoped_name Identifier
	{
	}
//(161)
consumes_dcl:
	RWconsumes scoped_name Identifier
	{
	}
//(163)
primary_key_spec:
	RWprimarykey scoped_name
	{
	}
//(165)
finder_dcl:
	RWfinder Identifier '('  ')'
	{
	}
	|RWfinder Identifier '('  ')'  raises_expr
	{
	}
	|RWfinder Identifier '('  init_param_dcls  ')'
	{
	}
	|RWfinder Identifier '('  init_param_dcls  ')'  raises_expr
	{
	}
//(166)
event_dcl:
	{
	}
	| event_def
	{
	}
	| event_abs_def
	{
	}
	| event_forward_dcl
	{
	}
//(167)
event_forward_dcl:
	RWeventtype Identifier
	{
	}
	|  RWabstract RWeventtype Identifier
	{
	}
//(168)
event_abs_def:
	RWabstract RWeventtype Identifier  '{' exportStar '}'
	{
	}
	| RWabstract RWeventtype Identifier value_inheritance_spec  '{' exportStar '}'
	{
	}
//(169)
event_def:
	event_header '{' value_elementStar '}'
	{
	}
//(170)
event_header:
	RWeventtype Identifier
	{
	}
	|  RWeventtype Identifier value_inheritance_spec
	{
	}
//(172)
porttype_dcl:
	porttype_def
	{
	}
	| porttype_forward_dcl
	{
	}
//(173)
porttype_forward_dcl:
	RWporttype Identifier
	{
	}
//(174)
porttype_def:
	RWporttype Identifier '{' port_body '}'
	{
	}
//(175)
port_body: port_ref port_exportStar{}
//(176)
port_ref:
	provides_dcl ';'
	{
	}
	| uses_dcl ';'
	{
	}
	| port_dcl ';'
	{
	}
//(177)
port_exportStar:
	{
	}
port_export:
	port_ref
	{
	}
	| attr_dcl ';'
	{
	}
//(178)
port_dcl:
 	RWmirrorport scoped_name Identifier
 	{
 	}
 	| RWport scoped_name Identifier
 	{
 	}
//(180)
connector_dcl:
	connector_header '{' connector_exportPlus '}'
	{
	}
//(181)
connector_header:
	RWconnector Identifier
	{
	}
	|RWconnector Identifier connector_inherit_spec
	{
	}
//(182)
connector_inherit_spec: ':'
	scoped_name
	{
	}
//(183)
connector_exportPlus:{}
connector_export:
	port_ref
	| attr_dcl ';'
//(185)
template_module_dcl:
	RWmodule Identifier '<' formal_parameters '>' '{' tpl_definitionPlus'}'{}
//(186)
formal_parameters:
	formal_parameterPlus
//(187)
formal_parameterPlus:
	formal_parameter
	| formal_parameterPlus ',' formal_parameter
formal_parameter:
	formal_parameter_type Identifier
//(188)
formal_parameter_type:
	RWtypename
	| RWinterface
	| RWvaluetype
	| RWeventtype
	| RWstruct
	| RWexception
	| RWenum
	| RWsequence
	| RWconst const_type
	| sequence_type
//(189)
tpl_definitionPlus:
tpl_definition:
	definition
	| template_module_ref ';'
//(190)
template_module_inst:
	RWmodule scoped_name '<' actual_parameters '>' Identifier {}

//(191)
actual_parameters:
	actual_parameterPlus

//(192)
actual_parameterPlus:
	actual_parameter
	| actual_parameterPlus ',' actual_parameter
actual_parameter:
 	type_spec
 	| const_expr
//(193)
template_module_ref: RWalias scoped_name '<' formal_parameter_names '>' Identifier{}
//(194)
IdentifierPlus:
	Identifier
	{
		$$ = UnionData[[]string]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($$.data, $1.data),
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	| IdentifierPlus ',' Identifier
	{
		$$ = UnionData[[]string]{
			$1.fileName,
			$1.lineNumber,
			$1.rowNumber,
			append($1.data, $3.data),
		}
		$3 = UnionData[string]{"", 0, 0, ""}
	}
formal_parameter_names: IdentifierPlus{}
//(199)
map_type:
 	RWmap '<' type_spec ',' type_spec ',' positive_int_const '>'{}
 	| RWmap '<' type_spec ',' type_spec '>'{}
//(200)
bitset_dcl:
	RWbitset Identifier  '{' bitfieldStar '}'{}
	|RWbitset Identifier ':' scoped_name '{' bitfieldStar '}'{}
//(201)

bitfieldStar:
bitfield:
	bitfield_spec IdentifierStar ';'
//(202)
bitfield_spec:
	RWbitfield '<' positive_int_const '>'
	| RWbitfield '<' positive_int_const ',' destination_type '>'
//(203)
destination_type:
	boolean_type
	| octet_type
	| integer_type
//(204)
bitmask_dcl:
	{}
	//RWbitmask Identifier '{' bit_valueStar '}'
//(205)
bit_valueStar:

bit_value: Identifier

//(208)
signed_tiny_int:
	RWint8
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			 &ConstType{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				YYrcvr.data,
				$1.data, true, 0, false, false}}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
//(209)
unsigned_tiny_int:
	RWuint8
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			 &ConstType{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				YYrcvr.data,
				$1.data,
				true,
				0,
				false,
				false,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
	| RWbyte
	{
		$$ = UnionData[ITypeSpec]{
			$1.fileName,
			$1.lineNumber,
			 $1.rowNumber,
			 &ConstType{
				fileInformation{
				$1.fileName,
				$1.lineNumber,
				$1.rowNumber,
				},
				YYrcvr.data,
				$1.data,
				true,
				0,
				false,
				false,
			},
		}
		$1 = UnionData[string]{"", 0, 0, ""}
	}
//(216)

//(219)
annotation_dcl: annotation_header '{' annotation_body '}'{}
//(220)
annotation_header: Annotation Identifier {}
//(221)
annotation_bodyEntry:annotation_member{}| enum_dcl ';'{}| const_dcl ';'{}| typedef_dcl ';'{}

annotation_bodyEntryStar:{}
annotation_body:
//(222)
annotation_member:
	annotation_member_type simple_declarator  ';'
	| annotation_member_type simple_declarator RWdefault const_expr  ';'
//(223)
annotation_member_type:
	const_type
	| any_const_type
	| scoped_name
//(224)
any_const_type:
	RWany
	{
	}
//(225)
annotation_appl:
	'@' scoped_name
	{
	}
	| '@' scoped_name  '(' annotation_appl_params ')'
	{
	}
//(226)
annotation_appl_params:
	const_expr
	{
	}
	| annotation_appl_paramPlus
	{
	}
//(227)
annotation_appl_paramPlus:
	annotation_appl_param
	{
	}
	| annotation_appl_paramPlus ',' annotation_appl_param
	{
	}
annotation_appl_param:
	Identifier '=' const_expr
	{
	}
%%

func SetDebugLevel(level int){
	YYDebug = level
}

func init() {
	YYErrorVerbose = true
}

