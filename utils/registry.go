package utils

import (
	"github.com/biningo/eagle/internal/config"
	"github.com/biningo/eagle/registry"
	"github.com/docker/docker/api/types"
	"strings"
)

/**
*@Author icepan
*@Date 7/20/21 15:22
*@Describe
**/

func ContainerToServiceInstance(container types.Container) *registry.ServiceInstance {
	s := strings.Split(container.Image, "/")
	namespace := "default"
	serviceName := container.Names[0][1:]
	serviceID := container.ID
	if len(s) > 1 {
		serviceName = s[1]
	} else if len(s) > 0 {
		serviceName = s[0]
	}
	if n, ok := container.Labels["namespace"]; ok {
		namespace = n
	}
	if sn, ok := container.Labels["serviceName"]; ok {
		serviceName = sn
	}
	if id, ok := container.Labels["serviceID"]; ok {
		serviceID = id
	}
	svc := registry.NewServiceInstance(
		namespace,
		serviceName,
		serviceID,
		container.NetworkSettings.Networks[config.Conf.DockerConfig.Network].IPAddress,
		container.Ports[0].PrivatePort,
		config.Conf.Host,
		container.Ports[0].PublicPort,
	)
	svc.Labels = container.Labels
	return svc
}
