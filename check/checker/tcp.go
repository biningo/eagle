package checker

import (
	"fmt"
	"net"
	"time"
)

/**
*@Author icepan
*@Date 7/22/21 11:04
*@Describe
**/

type TcpChecker struct {
	Address
}

func NewTcPChecker(host string, port uint16) *TcpChecker {
	return &TcpChecker{
		Address: Address{Host: host, Port: port},
	}
}

func (c TcpChecker) Ping(timeout uint8) bool {
	if _, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", c.Host, c.Port), time.Duration(timeout)*time.Second); err != nil {
		return false
	}
	return true
}
