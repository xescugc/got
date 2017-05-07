package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/xescugc/got/entities"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Tracking",
		Long:  "Starts tracking the time for the current project",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := entities.NewEnv()
			if err != nil {
				return err
			}

			c, err := entities.GetConfig()

			if !c.Exists {
				return errors.New("Could not find project\nRun 'got init' to start")
			}

			working, err := entities.IsWorking(env)
			if err != nil {
				return err
			}
			if working {
				return errors.New("Already in a task")
			}

			task := entities.NewTask(c.Project)

			if err := task.StartWorking(env); err != nil {
				return err
			}

			return nil

		},
	}
)

func init() {
	RootCmd.AddCommand(startCmd)
}
