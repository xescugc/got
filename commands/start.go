package commands

import (
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Tracking",
		Long:  "Starts tracking the time for the current project",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	RootCmd.AddCommand(startCmd)
}
