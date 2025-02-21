package helpers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type TaskPopulater struct {
	user, admin string
	teamNumber  int
}

type GameState struct {
	Task01Answer string
	Task03Answer string
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
	p.task03()
	p.task04()
}

func (p TaskPopulater) createScriptFile(dir string) {
	ExecuteCmd(fmt.Sprintf("sudo touch %s/script.sh", dir))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s %s/script.sh", p.admin, p.user, dir))
	ExecuteCmd(fmt.Sprintf("sudo chmod 770 %s/script.sh", dir))
}

func (p TaskPopulater) copyTaskInfoFile(dir string) {
	taskNumberStr := filepath.Base(dir)
	taskInfoFilePath := fmt.Sprintf("public/tasks/%s/t%s.md", taskNumberStr, taskNumberStr)

	ExecuteCmd(fmt.Sprintf("sudo cp %s %s/readme.md", taskInfoFilePath, dir))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s %s/readme.md", p.admin, p.user, dir))
	ExecuteCmd(fmt.Sprintf("sudo chmod 640 %s/readme.md", dir))
}

func (p TaskPopulater) task01() {
	task01Dir := fmt.Sprintf("/home/%s/404/01", p.user)

	ExecuteCmd(fmt.Sprintf("sudo mkdir -p %s", task01Dir))

	ExecuteCmd(fmt.Sprintf("sudo cp datasets/pincode.csv %s", task01Dir))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s %s/pincode.csv", p.admin, p.user, task01Dir))
	ExecuteCmd(fmt.Sprintf("sudo chmod 640 %s/pincode.csv", task01Dir))

	p.copyTaskInfoFile(task01Dir)
	p.createScriptFile(task01Dir)

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
			gameState.Task01Answer = lat + " " + long
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

	p.createScriptFile(task02Dir)
	p.copyTaskInfoFile(task02Dir)
}

func (p TaskPopulater) task03() {
	task03Dir := fmt.Sprintf("/home/%s/404/03", p.user)

	ExecuteCmd(fmt.Sprintf("sudo mkdir -p %s", task03Dir))

	ExecuteCmd(fmt.Sprintf("sudo cp -r public/tasks/03/files %s/files", task03Dir))

	p.createScriptFile(task03Dir)
	p.copyTaskInfoFile(task03Dir)

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

	gameState.Task03Answer = "a4c9vx.phan bfn5wy.phan i3s1fd.phan j5d6ch.phan lm6t0g.phan oy2p8n.phan rk7w4q.phan u8e2hj.phan xqr7p2.phan zt9k3l.phan"

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

func (p TaskPopulater) task04() {
	task04Dir := fmt.Sprintf("/home/%s/404/04", p.user)

	ExecuteCmd(fmt.Sprintf("sudo mkdir -p %s", task04Dir))
	ExecuteCmd(fmt.Sprintf("sudo cp -r public/tasks/04/files %s/files", task04Dir))

	p.createScriptFile(task04Dir)
	p.copyTaskInfoFile(task04Dir)

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
