package main

import (
	"fmt"
	"os"
	"os/user"
	"yail/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the interactive mode for YAIL!\n", user.Username)
	repl.Run(os.Stdin, os.Stdout)
}
