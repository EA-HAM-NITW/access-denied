package main

import (
	"accessdenied/commands"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if args[len(args)-1] == "gen" {
		commands.GenCmdHandler()
	} else if args[len(args)-1] == "check" {
		commands.EvalCmdHandler()
	} else {
		fmt.Println("invalid command")
	}
}
