package commands

import (
	"accessdenied/helpers"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func taskEvaluator(cmd, desiredOutput string) {
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
		taskEvaluator(answerCmd, gameState.Task01Answer)
	case 2:
		taskEvaluator(answerCmd, gameState.Task02Answer)
	case 3:
		taskEvaluator(answerCmd, gameState.Task03Answer)
	case 4:
		taskEvaluator(answerCmd, gameState.Task04Answer)
	case 5:
		taskEvaluator(answerCmd, gameState.Task05Answer)
	case 6:
		taskEvaluator(answerCmd, gameState.Task06Answer)
	default:
		fmt.Println("invalid task number")
		os.Exit(1)
	}
}
