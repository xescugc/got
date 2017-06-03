package cmd

import (
	"github.com/spf13/cobra"
)

var (
	project string
	RootCmd = &cobra.Command{
		Use:   "got",
		Short: "Tack your time",
		Long:  "Track the time you dedicate working to make exact reports",
	}
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&project, "project", "p", "", "Project name")
}
