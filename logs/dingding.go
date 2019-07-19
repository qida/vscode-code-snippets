package logs

import (
	"errors"

	"github.com/ggicci/dingtalk/robot"
)

var rb []*robot.Robot = make([]*robot.Robot, 0)

func InitRobot(name string, webhook string) {
	rb = append(rb, robot.New(name, webhook))
	return
}

func SendDingDing(level int8, content string) (err error) {
	if rb[level] == nil || rb[level].Name == "" {
		err = errors.New("没有初始化机器人！")
		return
	}
	m := rb[level].NewTextMessage()
	m.SetText(content)
	// m.AtAll(true)
	err = m.Send()
	return
}
