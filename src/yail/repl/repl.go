package repl

import (
	"bufio"
	"fmt"
	"io"
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
		fmt.Printf("You have typed: \"%s\"\n", input)
	}
}
