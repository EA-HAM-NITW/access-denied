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

type GameState struct {
	Task01Answer string
	Task04Answer string
}

func NewTaskPopulater(user, admin string, teamNumber int) TaskPopulater {
	return TaskPopulater{
		user:       user,
		admin:      admin,
		teamNumber: teamNumber,
	}
}

func (p TaskPopulater) Populate() {
	p.task01()
	p.task02()
	p.task04()
}

func (p TaskPopulater) task01() {
	task01Dir := fmt.Sprintf("/home/%s/404/01", p.user)

	ExecuteCmd(fmt.Sprintf("sudo mkdir -p %s", task01Dir))
	ExecuteCmd(fmt.Sprintf("sudo cp datasets/pincode.csv %s", task01Dir))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s %s/pincode.csv", p.admin, p.user, task01Dir))
	ExecuteCmd(fmt.Sprintf("sudo chmod 640 %s/pincode.csv", task01Dir))

	ExecuteCmd(fmt.Sprintf("sudo touch %s/script.sh", task01Dir))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s %s/script.sh", p.admin, p.user, task01Dir))
	ExecuteCmd(fmt.Sprintf("sudo chmod 770 %s/script.sh", task01Dir))

	pincodeDataset, err := os.Open("datasets/pincode.csv")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	gameStateFilePath := fmt.Sprintf("/home/%s/.game_state.json", p.user)
	gameStateBytes, err := ReadOrCreateFile(gameStateFilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var gameState GameState

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
			gameState = GameState{
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

func (p TaskPopulater) task02() {
	task02Dir := fmt.Sprintf("/home/%s/404/02", p.user)

	ExecuteCmd(fmt.Sprintf("sudo mkdir -p %s", task02Dir))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s %s", p.admin, p.admin, task02Dir))
	ExecuteCmd(fmt.Sprintf("sudo chmod -R 700 %s", task02Dir))
}

func (p TaskPopulater) task04() {
	task04Dir := fmt.Sprintf("/home/%s/404/04", p.user)
	ExecuteCmd(fmt.Sprintf("sudo mkdir -p %s", task04Dir))
	ExecuteCmd(fmt.Sprintf("sudo cp -r public/files_output %s", task04Dir))

	ExecuteCmd(fmt.Sprintf("sudo chown -R %s:%s %s", p.admin, p.user, task04Dir))
	ExecuteCmd(fmt.Sprintf("sudo chmod -R 640 %s", task04Dir))

	ExecuteCmd(fmt.Sprintf("sudo touch %s/script.sh", task04Dir))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s %s/script.sh", p.admin, p.user, task04Dir))
	ExecuteCmd(fmt.Sprintf("sudo chmod 770 %s/script.sh", task04Dir))

	gameStateFilePath := fmt.Sprintf("/home/%s/.game_state.json", p.user)
	gameStateBytes, err := ReadOrCreateFile(gameStateFilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var gameState GameState

	if err := json.Unmarshal(gameStateBytes, &gameState); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	gameState.Task04Answer = "phanek96 phaneq47 phangv03 phanif91 phanig96 phanlq97 phanmn94 phanou77 phanqw76 phanrj03 phanry94 phanvm88"
}
