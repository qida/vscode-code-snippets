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

	"github.com/qida/go/logs/dingtalk"
	// "github.com/blinkbean/dingtalk"
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

var MapRobot map[int8]*dingtalk.Robot

func init() {
	MapRobot = make(map[int8]*dingtalk.Robot, 0)
	MapRobot[Rb助手] = dingtalk.New("6279fab6b59f75a86bf7b3475b909a0311bc41d5cafe3577fba7925ffdccc6ed", "")
	MapRobot[Rb调试] = dingtalk.New("53706b6a41b6817cfd5fe2c905a9bc845dd6e7226619febaeff54df1077934a7", "")
	MapRobot[Rb错误] = dingtalk.New("b20da9534552e2d056e1b585f9269ba1949bd49fbac02eca6b75beb7ab4d8895", "")
	MapRobot[Rb重要] = dingtalk.New("56e3fa5947b1ff099417a29b3d2ba27b5a89365872e0d52813adb2b01db1e344", "")
	MapRobot[Rb监控] = dingtalk.New("647c2211f593cf5b4713dfd8981f0a5ae581218efb1304a5b3470b4a2b435f1f", "")
	MapRobot[Rb日常] = dingtalk.New("8aa4f1eb02c677ca56619b7f1ef1901783a0086af27475dec1cc32bfc8984f9b", "")
	MapRobot[Rb工作] = dingtalk.New("192e63da2025b1ab579febc20200a8372318fb7b10e0e70fb0972bd57df1c777", "")
	MapRobot[Rb打卡] = dingtalk.New("fc9b672e899bbe5ec5a26a22a817b58592d0591e4633abec25fa1488487ace94", "")
	MapRobot[Rb服务] = dingtalk.New("2ac60f0670075770953396dd0bbdbc9dc26ef1a733c803d25a1fb443b846861f", "")
	MapRobot[Rb正贤] = dingtalk.New("8fcb17718f3c525f5930ddea5fa175d10bc36368841f3bdb8603524c92b8396a", "")
	MapRobot[Rb积分] = dingtalk.New("439bec8b44c6ddd99b0522d29b6618fa129aa202127a1d004accccbf0f24a843", "SEC7d000ed14858f32ab9d3f5343d0f2527c796c0dfe304ccb068c7da488d936e14")
}

func Send2Ding(index int8, content string) (err error) {
	if ding, ok := MapRobot[index]; ok {
		m := ding.NewTextMessage()
		m.SetText(content)
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
		err = m.Send()
	} else {
		err = errors.New("没有初始化机器人！")
		panic(err)
	}
	return
}
