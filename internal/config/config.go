package config

/**
*@Author icepan
*@Date 7/22/21 09:32
*@Describe
**/
import (
	"fmt"
	"github.com/spf13/viper"
)

var Conf EagleConfig

type EagleConfig struct {
	Labels        []string `yaml:"labels"`
	*EtcdConfig   `mapstructure:"etcd"`
	*ServerConfig `mapstructure:"server"`
}

type EtcdConfig struct {
	Endpoints []string `yaml:"endpoints"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func InitConfigDefault() {
	sc := &ServerConfig{}
	sc.Port = "9999"
	sc.Host = "127.0.0.1"
	Conf.ServerConfig = sc

	ec := &EtcdConfig{}
	ec.Endpoints = []string{fmt.Sprintf("%s:%s", Conf.Host, "2380")}
	Conf.EtcdConfig = ec

	Conf.Labels = []string{"eagle"}
}

func InitConfigFromFile(filepath string) error {
	viper.SetConfigFile(filepath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		return err
	}
	return nil
}
