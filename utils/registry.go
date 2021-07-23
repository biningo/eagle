package utils

import (
	"github.com/biningo/eagle/registry"
	"github.com/docker/docker/api/types"
)

/**
*@Author icepan
*@Date 7/20/21 15:22
*@Describe
**/

func ContainerToServiceInstance(container types.Container) *registry.ServiceInstance {
	svc := registry.NewServiceInstance(
		container.Image,
		container.ID,
		container.NetworkSettings.Networks["bridge"].IPAddress,
		container.Ports[0].PrivatePort,
	)
	return svc
}
