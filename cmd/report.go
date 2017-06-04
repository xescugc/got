package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/xescugc/got/entities"
	"github.com/xescugc/got/utils"
)

var (
	from string
	to   string
	this string

	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "Report all the work",
		Long:  "Report all the work filtered with flags",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := entities.NewEnv()
			if err != nil {
				return err
			}

			fo, err := getFilterOption()
			if err != nil {
				return err
			}

			rf, err := entities.NewReportFilter(fo...)
			if err != nil {
				return err
			}

			ts, err := entities.NewReport(env, rf)
			if err != nil {
				return err
			}

			fmt.Println(ts)

			return nil
		},
	}
)

func init() {
	reportCmd.Flags().StringVarP(&from, "from", "f", "", "The FROM date where the report will start (dd/mm/yyyy)")
	reportCmd.Flags().StringVarP(&to, "to", "t", "", "The TO date where the report will end (dd/mm/yyyy)")
	reportCmd.Flags().StringVarP(&this, "this", "", "", "Fast way to report units of time (year, month, day)")

	RootCmd.AddCommand(reportCmd)
}

func getFilterOption() ([]entities.FilterOption, error) {
	fo := make([]entities.FilterOption, 0)

	if len(this) > 0 {
		from, to, err := utils.DateRange(time.Now(), this)
		if err != nil {
			return fo, err
		}
		fo = append(fo, entities.WithFilterTo(to), entities.WithFilterFrom(from))
	} else {
		if len(from) > 0 {
			f, err := time.Parse("02/01/2006", from)
			if err != nil {
				return fo, err
			}
			fo = append(fo, entities.WithFilterFrom(f))
		}

		if len(to) > 0 {
			t, err := time.Parse("02/01/2006", to)
			if err != nil {
				return fo, err
			}
			fo = append(fo, entities.WithFilterTo(t))
		}
	}

	return fo, nil
}
