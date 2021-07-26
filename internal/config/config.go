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
	Labels             []string `yaml:"labels"`
	*HealthCheckConfig `mapstructure:"health"`
	*EtcdConfig        `mapstructure:"etcd"`
	*ServerConfig      `mapstructure:"server"`
	*DockerConfig      `mapstructure:"docker"`
}

type CheckerConfig struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}

type HealthCheckConfig struct {
	Timeout        uint8 `yaml:"timeout"`
	Interval       uint8 `yaml:"interval"`
	*CheckerConfig `mapstructure:"checker"`
}

type EtcdConfig struct {
	Endpoints []string `yaml:"endpoints"`
	Prefix    string   `yaml:"prefix"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DockerConfig struct {
	Network string `yaml:"network"`
}

func InitConfigDefault() {
	sc := &ServerConfig{
		Host: "127.0.0.1",
		Port: "9999",
	}
	Conf.ServerConfig = sc

	ec := &EtcdConfig{
		Endpoints: []string{fmt.Sprintf("%s:%s", "127.0.0.1", "2380")},
		Prefix:    "eagle",
	}
	Conf.EtcdConfig = ec

	health := &HealthCheckConfig{
		Timeout:  5,
		Interval: 2,
		CheckerConfig: &CheckerConfig{
			Type: "tcp",
		},
	}
	Conf.HealthCheckConfig = health

	docker := &DockerConfig{Network: "bridge"}
	Conf.DockerConfig = docker

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
