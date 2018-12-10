package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/junbeomlee/jlang"
)

const PROMPT = ">>"
const EXIT = "exit"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == EXIT {
			fmt.Print("bye")
			return
		}

		l := jlang.New(line)
		for token := l.NextToken(); token.Type != jlang.EOF; token = l.NextToken() {
			fmt.Printf("%+v\n", token)
		}
	}
}
