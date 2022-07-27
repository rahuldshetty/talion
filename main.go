package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/rahuldshetty/talion/repl"
)

func main(){
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is talion language!\n", user.Username)
	fmt.Printf("You can now type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}