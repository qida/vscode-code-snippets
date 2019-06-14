package logs

import (
	"github.com/ggicci/dingtalk/robot"
)

var rb *robot.Robot

func InitRobot(name string, webhook string) *robot.Robot {
	rb = robot.New(name, webhook)
}

func SendDingDing(content string) (err error) {
	m := rb.NewTextMessage()
	m.SetText(content)
	m.AtAll(true)
	err = m.Send()
	return
}
