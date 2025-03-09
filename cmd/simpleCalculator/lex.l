%{
// Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz


package main

import (
    "bufio"
    "log"
    "strconv"
)

type yylexer struct{
    src     *bufio.Reader
    buf     []byte
    empty   bool
    current byte
}

func newLexer(src *bufio.Reader) (y *yylexer) {
    y = &yylexer{src: src}
    if b, err := src.ReadByte(); err == nil {
        y.current = b
    }
    return
}

func (y *yylexer) getc() byte {
    if y.current != 0 {
        y.buf = append(y.buf, y.current)
    }
    y.current = 0
    if b, err := y.src.ReadByte(); err == nil {
        y.current = b
    }
    return y.current
}

func (y yylexer) Error(e string) {
    log.Fatal(e)
}

func (y *yylexer) Lex(lval *yySymType) int {
    var err error
    c := y.current
    if y.empty {
        c, y.empty = y.getc(), false
    }
%}

%yyc c
%yyn c = y.getc()

D  [0-9]+
E  [eE][-+]?{D}
F  {D}"."{D}?{E}?|{D}{E}?|"."{D}{E}?

%%
    y.buf = y.buf[:0]

[ \t\r]+


break   {return BREAK}
default {return DEFAULT}
func {return FUNC}
interface {return INTERFACE}
select {return SELECT}
case {return CASE}
defer {return DEFER}
go {return GO}
map {return MAP}
struct {return STRUCT}
chan {return CHAN}
else {return ELSE}
goto {return GOTO}
package {return PACKAGE}
switch {return SWITCH}
const {return CONST}
fallthrough {return FALLTHROUGH}
if {return IF}
range {return RANGE}
type {return TYPE}
continue {return CONTINUE}
for {return FOR}
import {return IMPORT}
return {return RETURN}
var {return VAR}



{F}
    if lval.value, err = strconv.ParseFloat(string(y.buf), 64); err != nil {
        log.Fatal(err)
    }

    return NUM

%%
    y.empty = true
    return int(c)
}
