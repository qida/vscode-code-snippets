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

var MapRobot map[int8]*Robot

func init() {
	MapRobot = make(map[int8]*Robot, 0)
	MapRobot[Rb助手] = New("6279fab6b59f75a86bf7b3475b909a0311bc41d5cafe3577fba7925ffdccc6ed", "SEC1a8494b16e83a5af27742163371cc0dfbcac30feac4406f22e42f0d7758c6c6f")
	MapRobot[Rb调试] = New("53706b6a41b6817cfd5fe2c905a9bc845dd6e7226619febaeff54df1077934a7", "SEC6231e2382520b2ffe7bfe08d86e737c30213d4165acdaaa66836924d3f882dc2")
	MapRobot[Rb错误] = New("b20da9534552e2d056e1b585f9269ba1949bd49fbac02eca6b75beb7ab4d8895", "SECe5ce88434067c58ebb50053fcab938e0f70647c9c0ec90c02e65463539ecf9bd")
	MapRobot[Rb重要] = New("56e3fa5947b1ff099417a29b3d2ba27b5a89365872e0d52813adb2b01db1e344", "SEC58f19f278f1276e574af9a7535f5ae33422b25fef549737bc12cce213977fbc3")
	MapRobot[Rb监控] = New("647c2211f593cf5b4713dfd8981f0a5ae581218efb1304a5b3470b4a2b435f1f", "SECde04ab7fce8c2d4ef8ad314b1661382d5910c5c125ee8e26dfb05f0dfd042c8e")
	MapRobot[Rb日常] = New("8aa4f1eb02c677ca56619b7f1ef1901783a0086af27475dec1cc32bfc8984f9b", "SECa7e0c151d4fcc068876ea6beacf088f390625e9c78156255554b77d91dfc9604")
	MapRobot[Rb工作] = New("192e63da2025b1ab579febc20200a8372318fb7b10e0e70fb0972bd57df1c777", "SEC2f9e8dfffa1a8bb6f49fcf63dde06ba4ba9d2f9eee53ea08e9b6f47fead0621b")
	MapRobot[Rb打卡] = New("fc9b672e899bbe5ec5a26a22a817b58592d0591e4633abec25fa1488487ace94", "SECf75affd09c87af253e7dec33f8216b6950a15914be01430dd0bc7f3d3726a800")
	MapRobot[Rb服务] = New("2ac60f0670075770953396dd0bbdbc9dc26ef1a733c803d25a1fb443b846861f", "SEC6d49fb17116c51ae524107eed839d86a542299a0176e79b168a529d66f32563e")
	MapRobot[Rb正贤] = New("8fcb17718f3c525f5930ddea5fa175d10bc36368841f3bdb8603524c92b8396a", "SEC1b0d251c2d12bf3f980623aab21aa3080aaf28fab5ffbd70eb4e8a4a4b42bddf")
	MapRobot[Rb积分] = New("439bec8b44c6ddd99b0522d29b6618fa129aa202127a1d004accccbf0f24a843", "SEC7d000ed14858f32ab9d3f5343d0f2527c796c0dfe304ccb068c7da488d936e14")
}

func Send2Ding(index int8, content string) (err error) {
	if ding, ok := MapRobot[index]; ok {
		err = ding.SendTextMessage(content, []string{}, false)
	} else {
		err = errors.New("没有初始化机器人！")
		panic(err)
	}
	return
}
func Send2Dingf(index int8, format string, content ...interface{}) (err error) {
	if ding, ok := MapRobot[index]; ok {
		err = ding.SendTextMessage(fmt.Sprintf(format, content...), []string{}, false)
	} else {
		err = errors.New("没有初始化机器人！")
		panic(err)
	}
	return
}
