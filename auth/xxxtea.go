package auth

import (
  "encoding/base64"
  "strings"

	"github.com/xxtea/xxtea-go/xxtea"
)

const (
	KEY = "78sk89w8dh09sll892oa9n8hs9s892jb"
)

//加密
func EncryptStr(src []byte) (dst string) {
	dst = base64.URLEncoding.EncodeToString(xxtea.Encrypt(src, []byte(KEY)))
	return
}

///解密
func DecryptStr(src string) (dst string) {
	src = strings.TrimSpace(src)
	d, _ := base64.URLEncoding.DecodeString(src)
	dst = string(xxtea.Decrypt(d, []byte(KEY)))
	return
}

func Encrypt(debug bool, src []byte) (dst []byte) {
	if debug {
		dst = src
	} else {
		dst = xxtea.Encrypt(src, []byte(KEY))
	}
	return
}

///解密
func Decrypt(debug bool, src []byte) (dst []byte) {
	if debug {
		dst = src
	} else {
		dst = xxtea.Decrypt(src, []byte(KEY))
	}
	return
}
