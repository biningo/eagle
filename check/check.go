package check

import (
	"github.com/biningo/eagle/check/checker"
)

/**
*@Author icepan
*@Date 7/19/21 15:44
*@Describe
**/

type ServiceCheck struct {
	Checker  checker.Checker
	Interval uint8
	Timeout  uint8
	Host     string //127.0.0.1:9090
	Port     uint16
}
