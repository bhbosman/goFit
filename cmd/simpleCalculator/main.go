package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		if _, err := os.Stdout.WriteString("Salome, gee asb vir my n some om te doen. > "); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}

		buff := bytes.NewReader(line)
		yyParse(newLexer(bufio.NewReaderSize(buff, buff.Len()-1)))
	}
}
