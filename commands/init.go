package commands

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xescugc/got/entities"
	"github.com/xescugc/got/utils"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes the current project",
		Long:  "Initializes the current project adding the .got.json that stores the informatio for this project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := InitializeConfig(); err != nil {
				return err
			}

			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			current_config_file := path.Join(wd, ".got.json")
			exists, err := utils.ExistsPath(current_config_file)
			if err != nil {
				return err
			}
			if exists {
				fmt.Println("This directory is already set up.\nTry 'got show_config' to see the current aggregated config.")
				return nil
			}

			paths := strings.Split(wd, "/")
			project_name := paths[len(paths)-1]
			cf := entities.NewConfigFile(project_name)
			err = utils.WriteStructTo(current_config_file, cf)
			if err != nil {
				return err
			}

			projects, err := entities.GetProjects()
			if err != nil {
				return err
			}

			if _, ok := projects[project_name]; !ok {
				projects[project_name] = current_config_file
				err = projects.Save()
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
)

func init() {
	RootCmd.AddCommand(initCmd)
}
