package nets

import (
	"fmt"
)

var q = NewQQwry("./go/nets/qqwry.dat")

func GetAddressByIp(ip string) string {
	q.Find(ip)
	return fmt.Sprintf("%s", q.Address)
}
