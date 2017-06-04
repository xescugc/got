package entities

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/xescugc/got/utils"
)

// Project hold the data of one Project
type Project struct {
	Name    string
	Seconds int
}

// Add adds seconds to the internal seconds of the Project
func (p *Project) Add(s int) {
	p.Seconds += s
}

// Duration returns the Seconds as a Duration
func (p *Project) Duration() time.Duration {
	return time.Duration(p.Seconds) * time.Second
}

// Report holds all the projects aggregated
type Report struct {
	rf    *ReportFilter
	tasks map[string]*Project
}

// Add adds a t.Seconds to the corresponding project or creates a new one
func (r Report) Add(t *Task) {
	if pj, ok := r.tasks[t.Project]; ok {
		pj.Add(t.Seconds)
	} else {
		r.tasks[t.Project] = &Project{
			Name:    t.Project,
			Seconds: t.Seconds,
		}
	}
}

type ReportFilter struct {
	From    time.Time
	To      time.Time
	Project string
}

type FilterOption func(*ReportFilter)

func WithFilterFrom(f time.Time) FilterOption { return func(rf *ReportFilter) { rf.From = f } }
func WithFilterTo(t time.Time) FilterOption   { return func(rf *ReportFilter) { rf.To = t } }
func WithFilterProject(p string) FilterOption { return func(rf *ReportFilter) { rf.Project = p } }

func NewReportFilter(filters ...FilterOption) (*ReportFilter, error) {
	from, to, err := utils.DateRange(time.Now(), "month")
	if err != nil {
		return nil, err
	}
	rf := &ReportFilter{From: from, To: to}

	for _, f := range filters {
		f(rf)
	}
	return rf, nil
}

// NewReport initializes a new reporter and populates it
func NewReport(e *Env, rf *ReportFilter) (*Report, error) {
	var r = &Report{
		rf:    rf,
		tasks: make(map[string]*Project),
	}
	filepath.Walk(path.Join(e.DataHome), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if isValidPath(rf, p) {
			t, err := NewTaskFromPath(p)
			if err != nil {
				return err
			}
			r.Add(t)
		}
		return nil
	})
	return r, nil
}

// String formats the output of the Reporter
func (r Report) String() string {
	result := fmt.Sprintf("Projects worked and time (%s - %s ):\n\n", r.rf.From.Format("02/01/2006"), r.rf.To.Format("02/01/2006"))
	for _, p := range r.tasks {
		result += fmt.Sprintf("%s: \t%s\n", p.Name, p.Duration())
	}
	return result
}

// isValidPath validates if the path follows the filters specified/default by the user
func isValidPath(rf *ReportFilter, p string) bool {
	if filepath.Ext(p) != ".json" {
		return false
	}

	_, file := filepath.Split(p)
	sfile := strings.Split(file, "_")
	if len(sfile) < 2 {
		return false
	}
	d := sfile[0]
	project := strings.TrimSuffix(sfile[1], ".json")
	date, err := time.Parse("20060102-150405", d)
	if err != nil {
		return false
	}
	if date.After(rf.From) && date.Before(rf.To) {
		if len(rf.Project) > 0 {
			if rf.Project == project {
				return true
			}
		} else {
			return true
		}
	}

	return false
}
