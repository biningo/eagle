package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {

}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eagle",
	Short: "Registry and configuration center agent",
	Long:  `Registry and configuration center agent.Use etcd as the base storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[eagle] unrecognized command")
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
