
%{
package main

import (
    "fmt"
    "math"
)

%}

%union{
    value float64
}

%token	NUM
%token	BREAK
%token	DEFAULT
%token	FUNC
%token	INTERFACE
%token	SELECT
%token	CASE
%token	DEFER
%token	GO
%token	MAP
%token	STRUCT
%token	CHAN
%token	ELSE
%token	GOTO
%token	PACKAGE
%token	SWITCH
%token	CONST
%token	FALLTHROUGH
%token	IF
%token	RANGE
%token	TYPE
%token	CONTINUE
%token	FOR
%token	IMPORT
%token	RETURN
%token	VAR

%left	'-' '+'
%left	'*' '/'
%left	NEG     /* negation--unary minus */
%right	'^'     /* exponentiation */

%type	<value>	NUM
%type	<value> exp
%start input

%% /* The grammar follows.  */

input:    /* empty */
        | input line
;

line:     '\n'
        | exp '\n'  { fmt.Printf("\t%.10g\n", $1) }
;

exp:      NUM                { $$ = $1          }
        | exp '+' exp        { $$ = $1 + $3     }
        | exp '-' exp        { $$ = $1 - $3     }
        | exp '*' exp        { $$ = $1 * $3     }
        | exp '/' exp        { $$ = $1 / $3     }
        | '-' exp  %prec NEG { $$ = -$2         }
        | exp '^' exp        { $$ = math.Pow($1, $3) }
        | '(' exp ')'        { $$ = $2;         }
;
%%
