package commands

import (
	"errors"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/xescugc/got/entities"
	"github.com/xescugc/got/utils"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Tracking",
		Long:  "Starts tracking the time for the current project",
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			current_config_file := path.Join(wd, ".got.json")
			exists, err := utils.ExistsPath(current_config_file)
			if err != nil {
				return err
			}
			if !exists {
				return errors.New("Could not find project\nRun 'got init' to start")
			}

			if err := InitializeConfig(); err != nil {
				return err
			}

			cf, err := entities.GetConfig()
			if err != nil {
				return err
			}

			task := entities.NewTask(cf.Project)
			if err = task.Save(); err != nil {
				return err
			}

			current_task := path.Join(utils.DataHome, "current")
			exists, err = utils.ExistsPath(current_task)
			if err != nil {
				return err
			}
			if exists {
				return errors.New("Already in a task")
			}
			if err = utils.WriteTo(current_task, []byte(task.PathToTask())); err != nil {
				return err
			}

			return nil

		},
	}
)

func init() {
	RootCmd.AddCommand(startCmd)
}
