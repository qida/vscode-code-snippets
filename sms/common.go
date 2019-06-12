package sms

import (
	"fmt"
	"regexp"
	"time"
)

const (
	LimitMax = 2
)

var (
	RequestRegLimit int = LimitMax
)

func init() {
	TimerAddPoolLimit()
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

const (
	regular  = "^((13[0-9])|(14[1,4-9])|(15([0-3]|[5-9]))|(166)|(17[0|1,3-8])|(18[0-9])|(19[8|9]))\\d{8}$"
	duration = time.Minute * 10 //手机验证码超时时间
)

type VerificationParam struct {
	CodeType  string `json:"codetype"`
	SendMsgID string `json:"sendmsgid"`
}

func CheckRegexMobile(mobile string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}
