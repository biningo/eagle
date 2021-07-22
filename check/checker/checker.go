package checker

/**
*@Author icepan
*@Date 7/22/21 11:03
*@Describe
**/

type Checker interface {
	Ping(timeout uint8) bool
}

type Address struct {
	Host string
	Port uint16
}
