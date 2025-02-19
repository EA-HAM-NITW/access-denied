package commands

import (
	"accessdenied/helpers"
	"encoding/csv"
	"fmt"
	"os"
)

func GenCmdHandler() {
	var file string
	fmt.Println("enter path for the csv file")
	fmt.Scanln(&file)

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	var admin string
	fmt.Println("enter admin username")
	fmt.Scanln(&admin)

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for i := 1; i < len(records); i++ {
		name := records[i][0]
		password := records[i][1] + "RAM"

		fmt.Printf("generating an user with name %s and password %s\n", name, password)

		helpers.ExecuteCmd(fmt.Sprintf("sudo useradd -m -s /bin/bash %s", name))
		helpers.ExecuteCmd(fmt.Sprintf("echo %s:%s | sudo chpasswd", name, password))
		helpers.ExecuteCmd(fmt.Sprintf("sudo mkdir -p /home/%s/404", name))
		helpers.ExecuteCmd(fmt.Sprintf("sudo chown -R %s:%s /home/%s", admin, name, name))
		helpers.ExecuteCmd(fmt.Sprintf("sudo chmod 750 /home/%s", name))
		helpers.ExecuteCmd(fmt.Sprintf("sudo cp dist/cli /home/%s/cli", name))
		helpers.ExecuteCmd(fmt.Sprintf("sudo chown %s:%s /home/%s/cli", admin, name, name))
		helpers.ExecuteCmd(fmt.Sprintf("sudo chmod 750 /home/%s/cli", name))

		bashrcPath := fmt.Sprintf("/home/%s/.bashrc", name)
		f, err := os.OpenFile(bashrcPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		defer f.Close()

		fileInfo, err := f.Stat()
		if err == nil && fileInfo.Size() > 0 {
			if _, err := f.WriteString("\n"); err != nil {
				fmt.Println(err.Error())
				continue
			}
		}

		if _, err := f.WriteString("alias check='~/cli check'"); err != nil {
			fmt.Println(err.Error())
			continue
		}

		helpers.ExecuteCmd(fmt.Sprintf("sudo chown %s:%s /home/%s/.bashrc", name, admin, name))
		helpers.ExecuteCmd(fmt.Sprintf("sudo chmod 644 /home/%s/.bashrc", name))

		taskPopulater := helpers.NewTaskPopulater(name, admin, i)
		taskPopulater.Populate()
	}
}
