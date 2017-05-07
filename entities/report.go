package entities

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

type Project struct {
	Name    string
	Seconds int
}

func (p *Project) Add(s int) {
	p.Seconds += s
}

func (p *Project) Duration() time.Duration {
	return time.Duration(p.Seconds) * time.Second
}

type Report map[string]*Project

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

func (r Report) String() string {
	result := "Projects worked and time:\n\n"
	for _, p := range r {
		result += fmt.Sprintf("%s: \t%s\n", p.Name, p.Duration())
	}
	return result
}
