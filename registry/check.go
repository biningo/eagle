package registry

import (
	"context"
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
	Address                        string //127.0.0.1:9090
}

type Checker interface {
	Ping() bool
}

type HttpChecker struct {
	ServiceCheck
	Path string
}

func (c HttpChecker) Ping() bool {
	u, _ := url.Parse(fmt.Sprintf("http://%s%s", c.Address, c.Path))
	resp, err := http.Get(u.String())
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}

type TcpChecker struct {
	ServiceCheck
}

func (c TcpChecker) Ping() bool {
	if _, err := net.DialTimeout("tcp", c.Address, c.Timeout); err != nil {
		return false
	}
	return true
}

func Check(svc *ServiceInstance, registry Registrar) {
	if svc.Check == nil {
		if err := registry.Deregister(context.Background(), svc); err != nil {
			fmt.Println(err)
			return
		}
	}
	tick := time.Tick(svc.Service.Check.Interval)
	for {
		select {
		case <-tick:
			if ok := registry.HealthCheck(context.Background(), svc); !ok {
				if err := registry.Deregister(context.Background(), svc); err != nil {
					fmt.Println(err)
					return
				}
			}
		default:
		}
	}
}
