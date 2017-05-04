package commands

import (
	"github.com/spf13/cobra"
	"github.com/xescugc/got/entities"
)

var (
	cleanCmd = &cobra.Command{
		Use:   "clean",
		Short: "Clean all the data",
		Long:  "Clean all the data related to 'got' (history, project files ...)",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := entities.NewEnv()
			if err != nil {
				return err
			}

			projects, err := entities.GetProjects(env)
			if err != nil {
				return err
			}

			projects.Clean(env)

			env.Clean()

			return nil
		},
	}
)

func init() {
	RootCmd.AddCommand(cleanCmd)
}
