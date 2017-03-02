package commands

import (
	"errors"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Tracking",
		Long:  "Starts tracking the time for the current project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := InitializeConfig(); err != nil {
				return err
			}

			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			current_config_file := path.Join(wd, ".got.json")
			exists, err := ExistsPath(current_config_file)
			if err != nil {
				return err
			}
			if !exists {
				return errors.New("Could not find project")
			}

			return nil

		},
	}
)

func init() {
	RootCmd.AddCommand(startCmd)
}
