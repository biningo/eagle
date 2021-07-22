package agent

import (
	"context"
	"fmt"
	"log"

	"github.com/biningo/eagle/docker"
	"github.com/biningo/eagle/etcd"
	"github.com/biningo/eagle/internal/config"
	"github.com/biningo/eagle/registry"
	"github.com/biningo/eagle/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func RegistryAndHealthCheck(container types.Container, cli *clientv3.Client) {
	svc := utils.ContainerToServiceInstance(container)
	etcdRegistry := etcd.NewRegistry(cli)
	if err := etcdRegistry.Register(context.Background(), svc); err != nil {
		fmt.Println(err)
		return
	}
	go registry.Check(svc, etcdRegistry)
}

func Deregister(container types.Container, cli *clientv3.Client) {
	svc := utils.ContainerToServiceInstance(container)
	etcdRegistry := etcd.NewRegistry(cli)
	if err := etcdRegistry.Deregister(context.Background(), svc); err != nil {
		fmt.Println(err)
		return
	}
}

func Action(msg events.Message, dockerCli *client.Client, etcdCli *clientv3.Client) {
	container, err := docker.GetContainer(context.Background(), dockerCli, msg.ID)
	if err != nil {
		log.Println(err)
		return
	}

	switch msg.Action {
	case "start":
		RegistryAndHealthCheck(container, etcdCli)
	case "die":
		Deregister(container, etcdCli)
	default:
	}
}

func Run() {
	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println(err)
		return
	}

	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: config.Conf.Endpoints,
	})

	containers, err := docker.ListContainerByLabels(context.Background(), dockerCli, config.Conf.Labels)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range containers {
		RegistryAndHealthCheck(c, etcdCli)
	}
	msg, errCh := docker.ContainerEvents(context.Background(), dockerCli, config.Conf.Labels)
	for {
		select {
		case v := <-msg:
			Action(v, dockerCli, etcdCli)
		case err = <-errCh:
			fmt.Println(err)
			return
		}
	}
}
