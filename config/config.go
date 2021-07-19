package config

/**
*@Author lyer
*@Date 7/19/21 13:49
*@Describe
**/

// Config is a config interface.
type Config interface {
	Load(namespace string, name string, version string) error
}
