package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "got",
		Short: "Tack your time",
		Long:  "Track the time you dedicate working to make exact reports",
	}
)
