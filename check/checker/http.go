package checker

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

/**
*@Author icepan
*@Date 7/22/21 11:04
*@Describe
**/

type HttpChecker struct {
	Address
	Path string
}

func (c HttpChecker) Ping(timeout uint8) bool {
	u, _ := url.Parse(fmt.Sprintf("http://%s:%s%s", c.Host, c.Port, c.Path))
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Get(u.String())
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}
