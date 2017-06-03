package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/xescugc/got/entities"
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

			rf := entities.NewReportFilter(fo...)

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
		var from, to time.Time
		now := time.Now()
		currentYear, currentMonth, _ := now.Date()
		currentLocation := now.Location()

		if this == "month" {
			from = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
			to = from.AddDate(0, 1, -1)
		} else if this == "year" {
			from = time.Date(currentYear, 1, 1, 0, 0, 0, 0, currentLocation)
			to = from.AddDate(0, 12, -1)
		} else if this == "day" {
			from = now
			to = now
		} else {
			return fo, fmt.Errorf("Not a valid 'this': %s.\nThe valid ones are: month, year and day", this)
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
