package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xescugc/got/entities"
)

var (
	currentCmd = &cobra.Command{
		Use:   "current",
		Short: "Show current working task",
		Long:  "Show inforammtion of the current task",
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
				return errors.New("You don't have any task running")
			}

			t, err := entities.NewTaskFromCurrent(env)
			if err != nil {
				return err
			}

			fmt.Println(t)

			return nil
		},
	}
)

func init() {
	RootCmd.AddCommand(currentCmd)
}
