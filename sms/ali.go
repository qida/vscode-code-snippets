package sms

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/qida/aliyun_sms"
)

func SendAliSMS(aliyunSMS *aliyun_sms.AliyunSms, mobile string) (code string, err error) {
	if ok, err1 := CheckRegexMobile(mobile); !ok {
		err = err1
		return
	}
	if RequestRegLimit <= 0 {
		RequestRegLimit = 0
		return "", errors.New("您的请求太过频繁")
	}
	RequestRegLimit--
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code = fmt.Sprintf("%04v", rnd.Int31n(10000))
	err = aliyunSMS.Send(mobile, fmt.Sprintf(`{"code":"%s"}`, code))
	return
}
