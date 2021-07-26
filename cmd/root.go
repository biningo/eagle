package cmd

import (
	"fmt"
	"github.com/biningo/eagle/internal/config"
	"os"

	"github.com/biningo/eagle/agent"
	"github.com/biningo/eagle/app/router"
	"github.com/biningo/eagle/app/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().String("config", "config.yaml", "Configuration file for eagle.")
}

func InitConfigFromFilePath(cmd *cobra.Command) {
	filepath, err := cmd.PersistentFlags().GetString("config")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := config.InitConfigFromFile(filepath); err != nil {
		fmt.Println(err)
		return
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eagle",
	Short: "Registry and configuration center agent",
	Long:  `Registry and configuration center agent.Use etcd as the base storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("[eagle] unrecognized command")
		InitConfigFromFilePath(cmd)
		go agent.Run()
		server.Run(fmt.Sprintf("%s:%s", config.Conf.Host, config.Conf.Port), router.Init())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
