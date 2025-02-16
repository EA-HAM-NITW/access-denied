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
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// setting all the files and folders under admin user home dir to be private
	// it takes some time to run the following command
	// it can be ran once manually before running this script
	helpers.ExecuteCmd(fmt.Sprintf("sudo chmod -R 700 /home/%s", admin))

	for i := 1; i < len(records); i++ {
		name := records[i][0]
		password := helpers.RandomString(10)

		fmt.Printf("generating an user with name %s and password %s\n", name, password)

		helpers.ExecuteCmd(fmt.Sprintf("sudo useradd -m -s /bin/bash %s", name))
		helpers.ExecuteCmd(fmt.Sprintf("echo %s:%s | sudo chpasswd", name, password))
		helpers.ExecuteCmd(fmt.Sprintf("sudo chown -R %s:%s /home/%s", admin, name, name))
		helpers.ExecuteCmd(fmt.Sprintf("sudo mkdir -p /home/%s/404", name))
		helpers.ExecuteCmd(fmt.Sprintf("sudo find /home/%s -type d -exec chmod 750 {} +", name))
		helpers.ExecuteCmd(fmt.Sprintf("sudo find /home/%s -type f -exec chmod 640 {} +", name))

		taskPopulater := helpers.NewTaskPopulater(name, admin, i)
		taskPopulater.Populate()
	}
}
