package helpers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type TaskPopulater struct {
	user, admin string
	teamNumber  int
}

type GameStateValue struct {
	CurrentTask  int
	Task01Answer string
}

type GameState map[int]GameStateValue

func NewTaskPopulater(user, admin string, teamNumber int) TaskPopulater {
	return TaskPopulater{
		user:       user,
		admin:      admin,
		teamNumber: teamNumber,
	}
}

func (p TaskPopulater) Populate() {
	p.task01()
}

func (p TaskPopulater) task01() {
	ExecuteCmd(fmt.Sprintf("sudo mkdir -p /home/%s/404/001", p.user))
	ExecuteCmd(fmt.Sprintf("sudo cp datasets/pincode.csv /home/%s/404/001", p.user))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s /home/%s/404/001/pincode.csv", p.admin, p.user, p.user))
	ExecuteCmd(fmt.Sprintf("sudo chmod 640 /home/%s/404/001/pincode.csv", p.user))

	ExecuteCmd(fmt.Sprintf("sudo touch /home/%s/404/001/script.sh", p.user))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s /home/%s/404/001/script.sh", p.admin, p.user, p.user))
	ExecuteCmd(fmt.Sprintf("sudo chmod 777 /home/%s/404/001/script.sh", p.user))

	pincodeDataset, err := os.Open("datasets/pincode.csv")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	gameStateFilePath := fmt.Sprintf("/home/%s/.ad_game_state.json", p.admin)
	gameStateBytes, err := ReadOrCreateFile(gameStateFilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	gameState := make(GameState)

	if err := json.Unmarshal(gameStateBytes, &gameState); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	reader := csv.NewReader(pincodeDataset)
	records, _ := reader.ReadAll()

	for _, record := range records {
		pincode := record[4]
		lat := record[9]
		long := record[10]

		if pincode == strconv.Itoa(110000+p.teamNumber) {
			gameState[p.teamNumber] = GameStateValue{
				CurrentTask:  1,
				Task01Answer: lat + " " + long,
			}
			break
		}
	}

	gameStateBytes, err = json.Marshal(gameState)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := os.WriteFile(gameStateFilePath, gameStateBytes, 0644); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
