package sms

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var ali *dysmsapi.Client

var request *dysmsapi.SendSmsRequest

func init() {
	ali, _ = dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4FfBkTvJuEFVbr7URrV3", "zqLXNCSWVrOSRKG5lg5AGySZHzUK9B")
	request = dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = "UPM"
	request.TemplateCode = "SMS_175475271"
}
func AliSend(mobile string) (code string, err error) {
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
	request.PhoneNumbers = mobile
	request.TemplateParam = fmt.Sprintf(`{"code":"%s"}`, code)
	response, err := ali.SendSms(request)
	if err != nil {
		return
	}
	fmt.Printf("response is %#v\n", response)
	return
}
