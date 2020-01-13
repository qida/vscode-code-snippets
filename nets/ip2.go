package nets

import (
	"fmt"
)

var q = NewQQwry("qqwry.dat")

func GetAddressByIp(ip string) string {
	q.Find(ip)
	return fmt.Sprintf("%s", q.Address)
}
