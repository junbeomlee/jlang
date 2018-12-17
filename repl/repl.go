package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/junbeomlee/jlang"
	"github.com/junbeomlee/jlang/parser"
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
		p := parser.New(l)
		program := p.Parse()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
