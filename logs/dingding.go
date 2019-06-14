package logs

import (
	"github.com/ggicci/dingtalk/robot"
)

var rb *robot.Robot

func InitRobot(name string, webhook string) {
	rb = robot.New(name, webhook)
	return
}

func SendDingDing(content string) (err error) {
	m := rb.NewTextMessage()
	m.SetText(content)
	m.AtAll(true)
	err = m.Send()
	return
}
