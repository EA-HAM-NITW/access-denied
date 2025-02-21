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
	Task02Answer string
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

func (p TaskPopulater) copyFile(path, dir, filename string) {
	ExecuteCmd(fmt.Sprintf("sudo cp %s %s/%s", path, dir, filename))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s %s/%s", p.admin, p.user, dir, filename))
	ExecuteCmd(fmt.Sprintf("sudo chmod 640 %s/%s", dir, filename))
}

func (p TaskPopulater) updateGameState(updateFunc func(s *GameState)) {
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

	updateFunc(&gameState)

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

func (p TaskPopulater) copyTaskInfoFile(dir string) {
	taskNumberStr := filepath.Base(dir)
	taskInfoFilePath := fmt.Sprintf("public/tasks/%s/t%s.md", taskNumberStr, taskNumberStr)

	p.copyFile(taskInfoFilePath, dir, "readme.md")
}

func (p TaskPopulater) task01() {
	task01Dir := fmt.Sprintf("/home/%s/404/01", p.user)

	ExecuteCmd(fmt.Sprintf("sudo mkdir -p %s", task01Dir))

	p.copyFile("datasets/pincode.csv", task01Dir, "pincode.csv")
	p.copyTaskInfoFile(task01Dir)
	p.createScriptFile(task01Dir)

	pincodeDataset, err := os.Open("datasets/pincode.csv")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	p.updateGameState(func(s *GameState) {
		reader := csv.NewReader(pincodeDataset)
		records, _ := reader.ReadAll()

		for _, record := range records {
			pincode := record[4]
			lat := record[9]
			long := record[10]

			if pincode == strconv.Itoa(110000+p.teamNumber) {
				s.Task01Answer = lat + " " + long
				break
			}
		}
	})
}

func (p TaskPopulater) task02() {
	task02Dir := fmt.Sprintf("/home/%s/404/02", p.user)

	ExecuteCmd(fmt.Sprintf("sudo mkdir -p %s", task02Dir))

	p.copyFile("public/tasks/02/index.html", task02Dir, "index.html")
	p.createScriptFile(task02Dir)
	p.copyTaskInfoFile(task02Dir)

	p.updateGameState(func(s *GameState) {
		s.Task02Answer = "phan"
	})
}

func (p TaskPopulater) task03() {
	task03Dir := fmt.Sprintf("/home/%s/404/03", p.user)

	ExecuteCmd(fmt.Sprintf("sudo mkdir -p %s", task03Dir))

	ExecuteCmd(fmt.Sprintf("sudo cp -r public/tasks/03/files %s/files", task03Dir))

	p.createScriptFile(task03Dir)
	p.copyTaskInfoFile(task03Dir)

	p.updateGameState(func(s *GameState) {
		s.Task03Answer = "a4c9vx.phan bfn5wy.phan i3s1fd.phan j5d6ch.phan lm6t0g.phan oy2p8n.phan rk7w4q.phan u8e2hj.phan xqr7p2.phan zt9k3l.phan"
	})
}

func (p TaskPopulater) task04() {
	task04Dir := fmt.Sprintf("/home/%s/404/04", p.user)

	ExecuteCmd(fmt.Sprintf("sudo mkdir -p %s", task04Dir))
	ExecuteCmd(fmt.Sprintf("sudo cp -r public/tasks/04/files %s/files", task04Dir))

	p.createScriptFile(task04Dir)
	p.copyTaskInfoFile(task04Dir)
	p.updateGameState(func(s *GameState) {
		s.Task04Answer = "phanek96 phaneq47 phangv03 phanif91 phanig96 phanlq97 phanmn94 phanou77 phanqw76 phanrj03 phanry94 phanvm88"
	})
}
