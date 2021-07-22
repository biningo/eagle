package docker

/**
*@Author icepan
*@Date 7/19/21 15:48
*@Describe
**/
import (
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func makeLabelFilter(labels []string) filters.Args {
	args := filters.NewArgs()
	for _, l := range labels {
		args.Add("label", l)
	}
	return args
}

func ListContainerByLabels(ctx context.Context, client *client.Client, labels []string) ([]types.Container, error) {
	return client.ContainerList(ctx, types.ContainerListOptions{
		Filters: makeLabelFilter(labels),
	})
}

func ContainerEvents(ctx context.Context, client *client.Client, labels []string) (<-chan events.Message, <-chan error) {
	return client.Events(ctx, types.EventsOptions{Filters: makeLabelFilter(labels)})
}

func GetContainer(ctx context.Context, client *client.Client, id string) (types.Container, error) {
	args := filters.NewArgs()
	args.Add("id", id)
	containers, err := client.ContainerList(ctx, types.ContainerListOptions{
		Filters: args,
	})
	if err != nil {
		return types.Container{}, err
	}
	if len(containers) == 0 {
		return types.Container{}, errors.New("container does not exist")
	}
	return containers[0], err
}
