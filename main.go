package main

import (
	"accessdenied/helpers"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
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
	records, _ := reader.ReadAll()

	for i := 1; i < len(records); i++ {
		go func() {
			// have to change the structure a bit based on the final csv file format
			name := records[i][0]
			password := helpers.RandomString(10)

			fmt.Printf("generating an user with name %s and password %s\n", name, password)

			// create a new user account with bash as the shell
			helpers.ExecuteCmd(fmt.Sprintf("sudo useradd -m -s /bin/bash %s", name))
			// set the password
			helpers.ExecuteCmd(fmt.Sprintf("echo %s:%s | sudo chpasswd", name, password))
			// set the admin account and team account as the owners for team account's home dir
			helpers.ExecuteCmd(fmt.Sprintf("sudo chown -R %s:%s /home/%s", admin, name, name))
			// make `404` directory
			helpers.ExecuteCmd(fmt.Sprintf("sudo mkdir -p /home/%s/404", name))
			// give only read perms for team user and all perms for admin account
			helpers.ExecuteCmd(fmt.Sprintf("sudo find /home/%s -type d -exec chmod 750 {} +", name))
			helpers.ExecuteCmd(fmt.Sprintf("sudo find /home/%s -type f -exec chmod 640 {} +", name))
			// allows only admin to view files and dirs under the admin home dir
			helpers.ExecuteCmd(fmt.Sprintf("sudo chmod -R 700 /home/%s", admin))

			taskPopulater := helpers.NewTaskPopulater(name, admin)
			taskPopulater.Populate()
		}()
	}
}
