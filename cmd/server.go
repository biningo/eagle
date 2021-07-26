package cmd

import (
	"fmt"
	"github.com/biningo/eagle/app/router"
	"github.com/biningo/eagle/app/server"
	"github.com/biningo/eagle/internal/config"
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
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		InitConfigFromFilePath(rootCmd)
		switch args[0] {
		case "run":
			server.Run(fmt.Sprintf("%s:%s", config.Conf.Host, config.Conf.Port), router.Init())
		default:
			fmt.Println("[server] unrecognized command")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
