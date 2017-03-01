package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop Tracking",
		Long:  "Stops tracking time for the current project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hi")
		},
	}
)

func init() {
	RootCmd.AddCommand(stopCmd)
}
