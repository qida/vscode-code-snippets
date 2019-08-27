package logs

import (
	"errors"
	"strings"

	"github.com/ggicci/dingtalk/robot"
)

const (
	Rb助手 = 0
	Rb调试 = 1
	Rb错误 = 2
	Rb重要 = 3
	Rb监控 = 4
	Rb日常 = 5
	Rb工作 = 6
	Rb打卡 = 7
)

var MapRobot map[int8]*robot.Robot

func init() {
	MapRobot = make(map[int8]*robot.Robot, 0)
	MapRobot[Rb助手] = robot.New("助手", "https://oapi.dingtalk.com/robot/send?access_token=6279fab6b59f75a86bf7b3475b909a0311bc41d5cafe3577fba7925ffdccc6ed")
	MapRobot[Rb调试] = robot.New("调试", "https://oapi.dingtalk.com/robot/send?access_token=53706b6a41b6817cfd5fe2c905a9bc845dd6e7226619febaeff54df1077934a7")
	MapRobot[Rb错误] = robot.New("错误", "https://oapi.dingtalk.com/robot/send?access_token=b20da9534552e2d056e1b585f9269ba1949bd49fbac02eca6b75beb7ab4d8895")
	MapRobot[Rb重要] = robot.New("重要", "https://oapi.dingtalk.com/robot/send?access_token=56e3fa5947b1ff099417a29b3d2ba27b5a89365872e0d52813adb2b01db1e344")
	MapRobot[Rb监控] = robot.New("监控", "https://oapi.dingtalk.com/robot/send?access_token=647c2211f593cf5b4713dfd8981f0a5ae581218efb1304a5b3470b4a2b435f1f")
	MapRobot[Rb日常] = robot.New("日常", "https://oapi.dingtalk.com/robot/send?access_token=8aa4f1eb02c677ca56619b7f1ef1901783a0086af27475dec1cc32bfc8984f9b")
	MapRobot[Rb工作] = robot.New("工作", "https://oapi.dingtalk.com/robot/send?access_token=192e63da2025b1ab579febc20200a8372318fb7b10e0e70fb0972bd57df1c777")
	MapRobot[Rb打卡] = robot.New("打卡", "https://oapi.dingtalk.com/robot/send?access_token=fb663ad34e7a3e71bf6169863734d7ddeb9058a8a66123bc8b66d45f5bac6083")

}

func SendDingDing(index int8, content ...string) (err error) {
	if robot, ok := MapRobot[index]; ok {
		m := robot.NewTextMessage()
		m.SetText(strings.Join(content, " "))
		// m.AtAll(true)
		err = m.Send()
	} else {
		err = errors.New("没有初始化机器人！")
		panic(err)
	}
	return
}
