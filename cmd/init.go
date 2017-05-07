package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xescugc/got/entities"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes the current project",
		Long:  "Initializes the current project adding the .got.json that stores the information for this project",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := entities.NewEnv()
			if err != nil {
				return err
			}

			c, err := entities.GetConfig()

			if c.Exists {
				fmt.Println("This directory is already set up.")
				return nil
			}

			if err := c.Save(); err != nil {
				return err
			}

			projects, err := entities.GetProjects(env)
			if err != nil {
				return err
			}

			if _, ok := projects[c.Project]; !ok {
				projects[c.Project] = c.Path
				err = projects.Save(env)
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
