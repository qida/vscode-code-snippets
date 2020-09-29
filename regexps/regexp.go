/*
 * @Author: sunqida
 * @Date: 2019-05-25 07:56:43
 * @LastEditors: sunqida
 * @LastEditTime: 2019-05-25 07:56:43
 * @Description:
 */
package regexps

import (
	"errors"
	"regexp"

	"github.com/qida/go/logs"
)

const (
	regular = `1\d{10}`
)

func IsTelephone(mobile string) (err error) {
	reg := regexp.MustCompile(regular)
	if !reg.MatchString(mobile) {
		logs.Send2Dingf(logs.Rb监控, "手机号不正确：%s", mobile)
		err = errors.New("手机号不正确")
	}
	return
}
