package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xescugc/got/entities"
)

var (
	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "Report all the work",
		Long:  "Report all the work filtered with flags",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := entities.NewEnv()
			if err != nil {
				return err
			}

			ts, err := entities.NewReport(env)
			if err != nil {
				return err
			}

			fmt.Println(ts)

			return nil
		},
	}
)

func init() {
	RootCmd.AddCommand(reportCmd)
}
