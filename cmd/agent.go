package cmd

import (
	"fmt"
	"github.com/biningo/eagle/agent"
	"github.com/spf13/cobra"
)

/**
*@Author icepan
*@Date 7/20/21 10:58
*@Describe
**/

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "run agent",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		InitConfigFromFilePath(rootCmd)
		switch args[0] {
		case "run":
			agent.Run()
		default:
			fmt.Println("[agent] unrecognized command")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)
}
