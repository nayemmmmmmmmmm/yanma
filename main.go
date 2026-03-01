package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/nayemmmmmmmmmm/yanma/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! THis is Yanma programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdin)
}
