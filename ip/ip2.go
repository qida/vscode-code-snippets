package ip

import (
	"fmt"

	"github.com/yinheli/qqwry"
)

var q = qqwry.NewQQwry("go/ip/qqwry.dat")

func GetAddressByIp(ip string) string {
	q.Find(ip)
	return fmt.Sprintf("%s,%s", q.City, q.Country)
}
