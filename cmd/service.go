package cmd

import (
	"context"
	"github.com/biningo/eagle/docker"
	"github.com/spf13/cobra"
	"log"
)

/**
*@Author icepan
*@Date 7/19/21 16:38
*@Describe
**/

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "show docker service info",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("[service] unrecognized command")
		labels := []string{"name", "age"}
		containers, _ := docker.ListContainerByLabels(context.Background(), docker.Cli, labels)
		for _, c := range containers {
			log.Println(c.Names[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
