package commands

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "argo-events",
	Short: "Argo Events CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		spew.Dump(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(NewControllerCommand())
	rootCmd.AddCommand(NewEventSourceCommand())
	rootCmd.AddCommand(NewSensorCommand())
	rootCmd.AddCommand(NewWebhookCommand())
}
