package main

import (
	"fmt"
	"github.com/fandan-nyc/all-interpretors/monkey/repl"
	"os"
	user "os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello %s ! this is the monkey programming language\n", user.Username)
	fmt.Printf("Feel free to type in commands below.\n")
	repl.Start(os.Stdin, os.Stdout)
}
