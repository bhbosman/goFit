package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	for {
		x := rand.Int()%10 + 1
		y := rand.Int()%10 + 1

		if _, err := os.Stdout.WriteString(fmt.Sprintf("wat is %v + %v >", x, y)); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		line = line[:len(line)-1]
		n, err := strconv.ParseUint(string(line), 0, 64)
		if err != nil {
			log.Fatalf("ParseUint: %s", err)
		}
		if n == uint64(x+y) {
			if _, err := os.Stdout.WriteString(fmt.Sprintf("Korrek\n")); err != nil {
				log.Fatalf("WriteString: %s", err)
			}

		} else if _, err := os.Stdout.WriteString(fmt.Sprintf("Verkeerd" +
			"\n")); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
	}

}
