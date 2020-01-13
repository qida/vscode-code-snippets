package nets

import (
	"fmt"

	"github.com/yinheli/qqwry"
)

var q = qqwry.NewQQwry("github.com/qida/go/nets/qqwry.dat")

func GetAddressByIp(ip string) string {
	q.Find(ip)
	return fmt.Sprintf("%s,%s", q.City, q.Country)
}
