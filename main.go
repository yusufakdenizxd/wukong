package main

import (
	"fmt"
	"os"
	"os/user"
	"wukong.com/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if len(os.Args) == 0 {
		fmt.Printf("Hello %s! This is the Wukong programming language!\n",
			user.Username)
		fmt.Printf("Type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	} else {
		repl.StartScript(os.Stdin, os.Stdout, os.Args[1])

	}
}
