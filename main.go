package main

import (
	"accessdenied/commands"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . <command> [flags]")
		os.Exit(1)
	}

	genCmd := flag.NewFlagSet("gen", flag.ExitOnError)
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)

	teamsCsvFile := genCmd.String("teams", "", "teams' csv file")
	adminUsername := genCmd.String("admin", "", "admin username")

	switch os.Args[1] {
	case "gen":
		genCmd.Parse(os.Args[2:])
		commands.GenCmdHandler(*teamsCsvFile, *adminUsername)
	case "check":
		checkCmd.Parse(os.Args[2:])
		commands.EvalCmdHandler()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Available commands: gen, check")
		os.Exit(1)
	}
}
