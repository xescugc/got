package entities

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

func (rf ReportFilter) FromYear() int {
	y, _ := strconv.Atoi(rf.From.Format("2006"))
	return y
}

func (rf ReportFilter) FromMonth() int {
	m, _ := strconv.Atoi(rf.From.Format("01"))
	return m
}

func (rf ReportFilter) ToYear() int {
	y, _ := strconv.Atoi(rf.To.Format("2006"))
	return y
}

func (rf ReportFilter) ToMonth() int {
	m, _ := strconv.Atoi(rf.To.Format("01"))
	return m
}

type FilterOption func(*ReportFilter)

func WithFilterFrom(f time.Time) FilterOption { return func(rf *ReportFilter) { rf.From = f } }
func WithFilterTo(t time.Time) FilterOption   { return func(rf *ReportFilter) { rf.To = t } }
func WithFilterProject(p string) FilterOption { return func(rf *ReportFilter) { rf.Project = p } }

func NewReportFilter(filters ...FilterOption) *ReportFilter {
	rf := &ReportFilter{
		From: time.Now(),
		To:   time.Now(),
	}

	for _, f := range filters {
		f(rf)
	}
	return rf
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

	dir := strings.Split(filepath.Dir(p), string(filepath.Separator))

	year, err := strconv.Atoi(dir[len(dir)-2])
	if err != nil {
		return false
	}

	month, err := strconv.Atoi(dir[len(dir)-1])
	if err != nil {
		return false
	}

	if year >= rf.FromYear() && month >= rf.FromMonth() && year <= rf.ToYear() && month <= rf.ToMonth() {
		return true
	}

	return false
}
