/*
 * @Author: qida
 * @LastEditors: qida
 */
package sms

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/qida/go/logs"
)

const (
	LimitMax = 2
)

var (
	RequestRegLimit int = LimitMax
)

func init() {
	go TimerAddPoolLimit()
}

func TimerAddPoolLimit() {
	timer1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-time.After(2 * time.Second): //超时
		case <-timer1.C:
			go func() {
				if RequestRegLimit < LimitMax {
					RequestRegLimit++
					fmt.Println(RequestRegLimit)
				}
			}()
		}
	}
}

// 15633331111
// 13855554444
// 16677778888
// 17011112222
// 19911114444
const (
	regular  = `^1((3[0-9])|(4[1,4-9])|(5[0-3,5-9])|(6[6,7])|(7[0,1,3-8])|(8[0-9])|(9[1,5,8,9]))\d{8}$`
	duration = time.Minute * 10 //手机验证码超时时间
)

func CheckRegexMobile(mobile string) (err error) {
	if mobile == "" {
		errors.New("手机号不能为空！")
		return
	}
	reg := regexp.MustCompile(regular)
	if !reg.MatchString(mobile) {
		err = errors.New("手机号不满足格式要求！")
		logs.Send2Dingf(logs.Rb重要, "手机号：%s 不符合格式", mobile)
	}
	return
}
