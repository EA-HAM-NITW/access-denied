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
		fmt.Printf("invalid command: %s", os.Args[1])
		os.Exit(1)
	}
}
