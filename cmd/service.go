package cmd

import (
	"context"
	"fmt"
	"github.com/biningo/eagle/docker"
	"github.com/biningo/eagle/internal/config"
	"github.com/biningo/eagle/utils"
	"github.com/docker/docker/client"
	"github.com/liushuochen/gotable"
	"github.com/spf13/cobra"
)

/**
*@Author icepan
*@Date 7/19/21 16:38
*@Describe
**/

func listService() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println(err)
		return
	}
	tb, err := gotable.Create("Namespace", "Name", "ID", "PublicIP", "PublicPort", "PrivateIP", "PrivatePort", "Labels")
	if err != nil {
		fmt.Println(err)
		return
	}
	containers, _ := docker.ListContainerByLabels(context.Background(), cli, config.Conf.Labels)
	for _, c := range containers {
		utils.ShowServiceInstance(c, tb)
	}
}

func getService(serviceName string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println(err)
		return
	}
	tb, err := gotable.Create("Namespace", "Name", "ID", "PublicIP", "PublicPort", "PrivateIP", "PrivatePort", "Labels")
	if err != nil {
		fmt.Println(err)
		return
	}
	containers, _ := docker.ListContainerByLabels(context.Background(), cli, config.Conf.Labels)
	for _, c := range containers {
		if c.Image == serviceName {
			utils.ShowServiceInstance(c, tb)
		}
	}
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "show docker service info",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("[service] unrecognized command")
			return
		}
		command := args[0]
		switch command {
		case "list":
			listService()
		case "get":
			if len(args) < 2 {
				fmt.Println("[service] unrecognized command")
				return
			}
			getService(args[1])
		default:
			fmt.Println("[service] unrecognized command")
		}
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
