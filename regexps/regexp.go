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
)

const (
	regular = `(1\d{10}`
)

func IsTelephone(src string) (err error) {
	reg := regexp.MustCompile(regular)
	if !reg.MatchString(src) {
		err = errors.New("手机号不正确")
	}
	return
}
