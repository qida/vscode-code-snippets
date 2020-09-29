/*
 * @Author: sunqida
 * @Date: 2019-06-14 13:12:45
 * @LastEditors: sunqida
 * @LastEditTime: 2019-06-14 13:12:45
 * @Description:
 */
package logs

import (
	"errors"
	"fmt"

	"github.com/qida/go/logs/ding"
)

const (
	Rb助手 = iota
	Rb调试
	Rb错误
	Rb重要
	Rb监控
	Rb日常
	Rb工作
	Rb打卡
	Rb服务
	Rb正贤
	Rb积分
)
const (
	DingUrl = "https://oapi.dingtalk.com/ding/send?access_token="
)

var MapRobot map[int8]*ding.Robot

func init() {
	MapRobot = make(map[int8]*ding.Robot, 0)
	MapRobot[Rb助手] = ding.New("助手", DingUrl+"6279fab6b59f75a86bf7b3475b909a0311bc41d5cafe3577fba7925ffdccc6ed", "")
	MapRobot[Rb调试] = ding.New("调试", DingUrl+"53706b6a41b6817cfd5fe2c905a9bc845dd6e7226619febaeff54df1077934a7", "")
	MapRobot[Rb错误] = ding.New("错误", DingUrl+"b20da9534552e2d056e1b585f9269ba1949bd49fbac02eca6b75beb7ab4d8895", "")
	MapRobot[Rb重要] = ding.New("重要", DingUrl+"56e3fa5947b1ff099417a29b3d2ba27b5a89365872e0d52813adb2b01db1e344", "")
	MapRobot[Rb监控] = ding.New("监控", DingUrl+"647c2211f593cf5b4713dfd8981f0a5ae581218efb1304a5b3470b4a2b435f1f", "")
	MapRobot[Rb日常] = ding.New("日常", DingUrl+"8aa4f1eb02c677ca56619b7f1ef1901783a0086af27475dec1cc32bfc8984f9b", "")
	MapRobot[Rb工作] = ding.New("工作", DingUrl+"192e63da2025b1ab579febc20200a8372318fb7b10e0e70fb0972bd57df1c777", "")
	MapRobot[Rb打卡] = ding.New("打卡", DingUrl+"fc9b672e899bbe5ec5a26a22a817b58592d0591e4633abec25fa1488487ace94", "")
	MapRobot[Rb服务] = ding.New("服务", DingUrl+"2ac60f0670075770953396dd0bbdbc9dc26ef1a733c803d25a1fb443b846861f", "")
	MapRobot[Rb正贤] = ding.New("正贤", DingUrl+"8fcb17718f3c525f5930ddea5fa175d10bc36368841f3bdb8603524c92b8396a", "sunqida")
	MapRobot[Rb积分] = ding.New("积分", DingUrl+"edb68dd0585f7feaf05defff9d78eefb767916a91715a648fed2367a8680d7ba", "sunqida")
}

func Send2Ding(index int8, content string) (err error) {
	if ding, ok := MapRobot[index]; ok {
		m := ding.NewTextMessage()
		m.SetText(content)
		// m.AtAll(true)
		err = m.Send()
	} else {
		err = errors.New("没有初始化机器人！")
		panic(err)
	}
	return
}
func Send2Dingf(index int8, format string, content ...interface{}) (err error) {
	if ding, ok := MapRobot[index]; ok {
		m := ding.NewTextMessage()
		m.SetText(fmt.Sprintf(format, content...))
		// m.AtAll(true)
		err = m.Send()
	} else {
		err = errors.New("没有初始化机器人！")
		panic(err)
	}
	return
}
