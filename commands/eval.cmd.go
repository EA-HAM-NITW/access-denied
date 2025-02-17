package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func EvalCmdHandler() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	parentDir := filepath.Dir(currentDir)
	grandparentDir := filepath.Base(parentDir)

	if grandparentDir != "404" {
		fmt.Println("invalid location. run the `check` command in the tasks' subfolders")
		os.Exit(1)
	}

	answerCmdBytes, err := os.ReadFile(filepath.Join(currentDir, "script.sh"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	answerCmd := string(answerCmdBytes)
	cmdParts := strings.Split(answerCmd, " ")

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(output))
}
