package cmd

import (
	"github.com/biningo/eagle/app/router"
	"github.com/biningo/eagle/app/server"
	"github.com/spf13/cobra"
)

/**
*@Author icepan
*@Date 7/19/21 16:38
*@Describe
**/

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run agent server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(":9090", router.Init())
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
