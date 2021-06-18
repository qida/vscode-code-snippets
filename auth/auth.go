package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type UserAuth struct {
	User interface{}
	tm   time.Time
	salt int
}

//加密
func AuthEncrypt(uid int, userAuth *UserAuth) string {
	userAuth.salt = rand.Intn(1000)
	userAuth.tm = time.Now()
	src, _ := json.Marshal(userAuth)
	encodeString := base64.URLEncoding.EncodeToString(Encrypt(false, src))
	return fmt.Sprintf("%d-%s", uid, encodeString)
}

//解密
func AuthDecrypt(uid int, auth string) (userAuth UserAuth, err error) {
	var decodeBytes []byte
	if decodeBytes, err = base64.URLEncoding.DecodeString(auth); err == nil {
		err = json.Unmarshal(Decrypt(false, decodeBytes), &userAuth)
	}
	return
}
