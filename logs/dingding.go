package logs

import (
	"errors"

	"github.com/ggicci/dingtalk/robot"
)

var rb *robot.Robot

func InitRobot(name string, webhook string) {
	rb = robot.New(name, webhook)
	return
}

func SendDingDing(content string) (err error) {
	if rb == nil || rb.Name == "" {
		err = errors.New("没有初始化机器人！")
		return
	}
	m := rb.NewTextMessage()
	m.SetText(content)
	m.AtAll(true)
	err = m.Send()
	return
}
