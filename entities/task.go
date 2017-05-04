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

func (t *Task) Save(e *Env) error {
	err := utils.WriteStructTo(t.PathToTask(e), t)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(t.directory(e), 0755); err != nil {
			return err
		}
		return t.Save(e)
	}
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) PathToTask(e *Env) string {
	return path.Join(t.directory(e), t.filename())
}

func (t *Task) SetWorking(e *Env) error {
	return utils.WriteTo(path.Join(e.DataHome, "current"), []byte(t.PathToTask(e)))
}

func IsWorking(e *Env) (bool, error) {
	return utils.ExistsPath(path.Join(e.DataHome, "current"))
}

func (t *Task) filename() string {
	return fmt.Sprintf("%v-%v_%v.json", t.Start.Format("20060102"), t.Start.Format("150405"), t.Project)
}

func (t *Task) directory(e *Env) string {
	return path.Join(e.DataHome, t.Start.Format("2006"), t.Start.Format("01"))
}
