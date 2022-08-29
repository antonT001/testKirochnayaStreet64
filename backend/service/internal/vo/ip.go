package vo

import (
	"fmt"
	"net"
)

type Ip struct {
	value string
}

func ExamineIntIp(value string) (Ip, error) {
	var (
		ip  Ip
		err error
	)

	ip.value = value

	if net.ParseIP(ip.value) == nil {
		err = fmt.Errorf("id error")
	}

	return ip, err
}

func (ip *Ip) Ip() string {
	return ip.value
}
