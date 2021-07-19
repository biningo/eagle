package docker

/**
*@Author icepan
*@Date 7/19/21 15:48
*@Describe
**/
import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func init() {
	SetUp()
}

var Cli *client.Client

func SetUp() {
	var err error
	Cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
}

func ListContainerByLabels(ctx context.Context, client *client.Client, labels []string) ([]types.Container, error) {
	args := filters.NewArgs()
	for _, l := range labels {
		args.Add("label", l)
	}
	return client.ContainerList(ctx, types.ContainerListOptions{
		Filters: args,
	})
}
