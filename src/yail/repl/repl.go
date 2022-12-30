package repl

import (
	"bufio"
	"fmt"
	"io"
	"yail/evaluator"
	"yail/lexer"
	"yail/parser"	
)

const (
	PROMPT = ">> "
	EXIT   = "exit"
	QUIT   = "q"
)

func Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		input := scanner.Text()
		if input == EXIT || input == QUIT {
			return
		}
		l := lexer.New(input)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Failed to execute the given source code for following reasons.\n")
	for _, msg := range errors {
		io.WriteString(out, "\t[ERROR] " + msg + "\n")
	}
}