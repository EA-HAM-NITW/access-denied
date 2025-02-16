package helpers

import (
	"fmt"
)

type TaskPopulater struct {
	user, admin string
}

func NewTaskPopulater(user, admin string) TaskPopulater {
	return TaskPopulater{
		user:  user,
		admin: admin,
	}
}

func (p TaskPopulater) Populate() {
	p.task01()
}

func (p TaskPopulater) task01() {
	// creating a directory for 1st task and copying csv file to that dir and assigning proper perms so that the users can view the file
	ExecuteCmd(fmt.Sprintf("sudo mkdir -p /home/%s/404/001", p.user))
	ExecuteCmd(fmt.Sprintf("sudo cp datasets/pincode.csv /home/%s/404/001", p.user))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s /home/%s/404/001/pincode.csv", p.admin, p.user, p.user))
	ExecuteCmd(fmt.Sprintf("sudo chmod 640 /home/%s/404/001/pincode.csv", p.user))

	// creating the `script.sh` beforehand so that the users can just edit it directly
	ExecuteCmd(fmt.Sprintf("sudo touch /home/%s/404/001/script.sh", p.user))
	ExecuteCmd(fmt.Sprintf("sudo chown %s:%s /home/%s/404/001/script.sh", p.admin, p.user, p.user))
	ExecuteCmd(fmt.Sprintf("sudo chmod 777 /home/%s/404/001/script.sh", p.user))
}
