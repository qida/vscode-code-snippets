package nets

import (
	"fmt"
)

var q = NewQQwry("github.com/qida/go/nets/qqwry.dat")

func GetAddressByIp(ip string) string {
	q.Find(ip)
	return fmt.Sprintf("%s", q.Address)
}
