package sms

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/qiniu/api.v7/sms"
)

type QiniuSMS struct {
	SignatureID string
	TemplateID  string
}

func SendQiniuSMS(qiniuSMS QiniuSMS, mobile string) (code int, err error) {
	if !CheckRegexMobile(mobile) {
		return "", errors.New("手机号码不正确！")
	}
	if RequestRegLimit <= 0 {
		RequestRegLimit = 0
		return "", errors.New("您的请求太过频繁")
	}
	RequestRegLimit--
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code = fmt.Sprintf("%04d", rnd.Int31n(10000))
	// SendMessage
	args := sms.MessagesRequest{
		SignatureID: qiniuSMS.SignatureID,
		TemplateID:  qiniuSMS.TemplateID,
		Mobiles:     []string{mobile},
		Parameters: map[string]interface{}{
			"code": code,
		},
	}
	ret, err := manager.SendMessage(args)
	if err != nil {
		err = errors.New("SendMessage() error: %v\n", err)
	}
	if len(ret.JobID) == 0 {
		err = errors.New("SendMessage() error: The job id cannot be empty")
	}
}
