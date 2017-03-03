package entities

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/xescugc/got/utils"
)

type Task struct {
	Description string     `json:"description,omitempty"`
	Project     string     `json:"project"`
	Start       *time.Time `json:"start"`
	Stop        *time.Time `json:"stop,omitempty"`
	Duration    string     `json:"duration,omitempty"`
}

func NewTask(project string) *Task {
	t := time.Now()
	return &Task{
		Project: project,
		Start:   &t,
	}
}

func (t *Task) filename() string {
	return fmt.Sprintf("%v-%v_%v.json", t.Start.Format("20060102"), t.Start.Format("150405"), t.Project)
}

func (t *Task) directory() string {
	return path.Join(utils.DataHome, t.Start.Format("2006"), t.Start.Format("01"))
}

func (t *Task) PathToTask() string {
	return path.Join(t.directory(), t.filename())
}

func (t *Task) Save() error {
	err := utils.WriteStructTo(t.PathToTask(), t)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(t.directory(), 0755); err != nil {
			return err
		}
		return t.Save()
	}
	if err != nil {
		return err
	}
	return nil
}
