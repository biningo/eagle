package agent

import (
	"context"
	"fmt"
	"github.com/biningo/eagle/registry"
	"log"
	"time"

	"github.com/biningo/eagle/docker"
	"github.com/biningo/eagle/etcd"
	"github.com/biningo/eagle/internal/config"
	"github.com/biningo/eagle/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Agent struct {
	registrars map[string]registry.Registrar
	services   map[string]*registry.ServiceInstance
	dockerCli  *client.Client
	etcdCli    *clientv3.Client
}

func (a *Agent) RegistryAndHealthCheck(container types.Container) {
	svc := utils.ContainerToServiceInstance(container, config.Conf.Namespace)
	etcdRegistry := etcd.NewRegistry(a.etcdCli, svc, etcd.WithPrefix(config.Conf.Prefix))
	a.registrars[svc.ID] = etcdRegistry
	a.services[svc.ID] = svc
	if err := etcdRegistry.Register(context.Background(), svc); err != nil {
		fmt.Println(err)
		return
	}
	go etcdRegistry.HealthCheck(svc)
}

func (a *Agent) Deregister(id string) {
	etcdRegistry := a.registrars[id]
	svc := a.services[id]
	if err := etcdRegistry.Deregister(context.Background(), svc); err != nil {
		fmt.Println(err)
		return
	}
}

func (a *Agent) Action(msg events.Message) {
	switch msg.Action {
	case "start":
		container, err := docker.GetContainer(context.Background(), a.dockerCli, msg.ID)
		if err != nil {
			log.Println(err)
			return
		}
		a.RegistryAndHealthCheck(container)
	case "die":
		a.Deregister(msg.ID)
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
		Endpoints:   config.Conf.Endpoints,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	agent := &Agent{
		registrars: make(map[string]registry.Registrar),
		services:   make(map[string]*registry.ServiceInstance),
		dockerCli:  dockerCli,
		etcdCli:    etcdCli,
	}

	containers, err := docker.ListContainerByLabels(context.Background(), dockerCli, config.Conf.Labels)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range containers {
		agent.RegistryAndHealthCheck(c)
	}
	msg, errCh := docker.ContainerEvents(context.Background(), dockerCli, config.Conf.Labels)
	fmt.Println("watch docker container.....")
	for {
		select {
		case v := <-msg:
			agent.Action(v)
		case err = <-errCh:
			fmt.Println(err)
			return
		}
	}
}
