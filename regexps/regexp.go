/*
 * @Author: sunqida
 * @Date: 2019-05-25 07:56:43
 * @LastEditors: sunqida
 * @LastEditTime: 2019-05-25 07:56:43
 * @Description:
 */
package regexps

import "regexp"

const (
	regular = `(13[0-9]|14[57]|15[0-35-9]|18[07-9]|17[7])\d{8}`
)

func GetTelphone(text string) (tel []string) {

	reg := regexp.MustCompile(regular)
	tel = reg.FindAllString(text, -1)
	return
}
