package registry

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

/**
*@Author icepan
*@Date 7/19/21 15:44
*@Describe
**/

type ServiceCheck struct {
	Checker                        Checker
	Interval                       time.Duration
	Timeout                        time.Duration
	DeregisterCriticalServiceAfter time.Duration
	Address                        string //"127.0.0.1:9090"
}

type Checker interface {
	Ping() bool
}

type httpChecker struct {
	ServiceCheck
	Path string
}

func (c httpChecker) Ping() bool {
	u, _ := url.Parse(fmt.Sprintf("http://%s%s", c.Address, c.Path))
	resp, err := http.Get(u.String())
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}

type tcpChecker struct {
	ServiceCheck
}

func (c tcpChecker) Ping() bool {
	if _, err := net.DialTimeout("tcp", c.Address, c.Timeout); err != nil {
		return false
	}
	return true
}
