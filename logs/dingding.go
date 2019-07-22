package logs

import (
	"errors"

	"github.com/ggicci/dingtalk/robot"
)

const (
	Rb助手 = 0
	Rb调试 = 1
	Rb错误 = 2
	Rb重要 = 3
)

var MapRobot map[int8]*robot.Robot

func init() {
	MapRobot = make(map[int8]*robot.Robot, 0)
	MapRobot[Rb助手] = robot.New("助手", "https://oapi.dingtalk.com/robot/send?access_token=376653979aee75c1cc07a08544817207e6f7a63db018130911c5d50dd5df5fed")
	MapRobot[Rb调试] = robot.New("调试", "https://oapi.dingtalk.com/robot/send?access_token=53706b6a41b6817cfd5fe2c905a9bc845dd6e7226619febaeff54df1077934a7")
	MapRobot[Rb错误] = robot.New("错误", "https://oapi.dingtalk.com/robot/send?access_token=b20da9534552e2d056e1b585f9269ba1949bd49fbac02eca6b75beb7ab4d8895")
	MapRobot[Rb重要] = robot.New("重要", "https://oapi.dingtalk.com/robot/send?access_token=56e3fa5947b1ff099417a29b3d2ba27b5a89365872e0d52813adb2b01db1e344")
}

func SendDingDing(index int8, content string) (err error) {
	if robot, ok := MapRobot[index]; ok {
		m := robot.NewTextMessage()
		m.SetText(content)
		// m.AtAll(true)
		err = m.Send()
	} else {
		err = errors.New("没有初始化机器人！")
	}
	return
}
