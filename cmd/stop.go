package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/xescugc/got/entities"
)

var (
	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop Tracking",
		Long:  "Stops tracking time for the current project",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := entities.NewEnv()
			if err != nil {
				return err
			}

			working, err := entities.IsWorking(env)
			if err != nil {
				return err
			}
			if !working {
				return errors.New("You don't have any task")
			}

			t, err := entities.NewTaskFromCurrent(env)
			if err != nil {
				return err
			}

			t.StopWorking(env)

			return nil
		},
	}
)

func init() {
	RootCmd.AddCommand(stopCmd)
}
