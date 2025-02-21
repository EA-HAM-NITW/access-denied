package commands

import (
	"accessdenied/helpers"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func task01Evaluator(cmd, desiredOutput string) {
	if !helpers.CheckForCmdKeywords(cmd, "grep", "awk") {
		fmt.Println("wrong cmd keywords. use either `grep` or `awk`")
		os.Exit(1)
	}

	output := helpers.GetCmdOutput(cmd)

	if output == desiredOutput {
		fmt.Println("correct")
	} else {
		fmt.Println("wrong")
	}
}

func task04Evaluator(cmd, desiredOutput string) {
	if !helpers.CheckForCmdKeywords(cmd, "grep", "awk") {
		fmt.Println("wrong cmd keywords. use either `grep` or `awk`")
		os.Exit(1)
	}

	output := helpers.GetCmdOutput(cmd)

	if output == desiredOutput {
		fmt.Println("correct")
	} else {
		fmt.Println("wrong")
	}
}

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

	homeDir, _ := os.UserHomeDir()
	gameStateBytes, err := os.ReadFile(filepath.Join(homeDir, ".game_state.json"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var gameState helpers.GameState

	if err := json.Unmarshal(gameStateBytes, &gameState); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	currentTask, err := strconv.Atoi(filepath.Base(currentDir))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	switch currentTask {
	case 1:
		task01Evaluator(answerCmd, gameState.Task01Answer)
	case 4:
		task04Evaluator(answerCmd, gameState.Task04Answer)
	default:
		fmt.Println("invalid task number")
		os.Exit(1)
	}
}
