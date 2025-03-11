build_fit_yacc:
	cp ./parsers/fitGenerator/internal/fit.y ./fit.y
	goyacc -p YY -o ./parsers/fitGenerator/internal/fit.yy.go ./fit.y
	rm ./fit.y

build_idl_yacc:
	cp ./parsers/idlGenerator/internal/idl.y ./idl.y
	goyacc -p YY -o ./parsers/idlGenerator/internal/idl.yy.go ./idl.y
	rm ./idl.y

build_idl_lex:
	golex   -o ./parsers/idlGenerator/internal/idl.ll.go ./parsers/idlGenerator/internal/idlLexer.l

build_fit_lex:
	golex   -o ./parsers/fitGenerator/internal/fitLexer.ll.go ./parsers/fitGenerator/internal/fitLexer.l

build: build_idl_yacc build_idl_lex build_fit_yacc build_fit_lex
	go build
	go install

runFromFiles: build
#	idl -o../../ddd/xdl_fitMappings.go ../../idlResources/fitMappings.idl
#	idl -o../../ddd/xdl_struct.go ../../idlResources/struct.idl
#	idl -o../../ddd/xdl_defines.go ../../idlResources/defines.idl
	idl  -gfit -d4 -o./ddd/xdl_fit.go ../../fitIdl/types/fitRecord.idl

runEcho: build
	echo "struct ses {};"  |  ./idl
	echo "struct one {};struct two {};"  |  ./idl
