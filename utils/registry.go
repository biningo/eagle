package utils

import (
	"github.com/biningo/eagle/internal/config"
	"github.com/biningo/eagle/registry"
	"github.com/docker/docker/api/types"
)

/**
*@Author icepan
*@Date 7/20/21 15:22
*@Describe
**/

func ContainerToServiceInstance(container types.Container, namespace string) *registry.ServiceInstance {
	svc := registry.NewServiceInstance(
		namespace,
		container.Image,
		container.ID,
		container.NetworkSettings.Networks[config.Conf.DockerConfig.Network].IPAddress,
		container.Ports[0].PrivatePort,
		container.NetworkSettings.Networks[config.Conf.DockerConfig.Network].IPAddress,
		container.Ports[1].PrivatePort,
	)
	return svc
}
