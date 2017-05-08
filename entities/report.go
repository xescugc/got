package entities

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
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
type Report map[string]*Project

// Add adds a t.Seconds to the corresponding project or creates a new one
func (r Report) Add(t *Task) {
	if pj, ok := r[t.Project]; ok {
		pj.Add(t.Seconds)
	} else {
		r[t.Project] = &Project{
			Name:    t.Project,
			Seconds: t.Seconds,
		}
	}
}

// NewReport initializes a new reporter and populates it
func NewReport(e *Env) (*Report, error) {
	var r = make(Report)
	filepath.Walk(path.Join(e.DataHome, strconv.Itoa(time.Now().Year())), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(p) == ".json" {
			t, err := NewTaskFromPath(p)
			if err != nil {
				return err
			}
			r.Add(t)
		}
		return nil
	})
	return &r, nil
}

// String formats the output of the Reporter
func (r Report) String() string {
	result := "Projects worked and time:\n\n"
	for _, p := range r {
		result += fmt.Sprintf("%s: \t%s\n", p.Name, p.Duration())
	}
	return result
}
