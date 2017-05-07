package entities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/xescugc/got/utils"
)

type Task struct {
	Description string     `json:"description,omitempty"`
	Project     string     `json:"project"`
	Start       *time.Time `json:"start"`
	Stop        *time.Time `json:"stop,omitempty"`
	Seconds     int        `json:"seconds,omitempty"`
}

func NewTask(project string) *Task {
	t := time.Now()
	return &Task{
		Project: project,
		Start:   &t,
	}
}

func NewTaskFromCurrent(e *Env) (*Task, error) {
	task_path, err := ioutil.ReadFile(path.Join(e.DataHome, "current"))
	if err != nil {
		return nil, err
	}
	return NewTaskFromPath(string(task_path))
}

func NewTaskFromPath(p string) (*Task, error) {
	data, err := ioutil.ReadFile(string(p))
	if err != nil {
		return nil, err
	}

	var t Task
	json.Unmarshal(data, &t)

	return &t, nil
}

func IsWorking(e *Env) (bool, error) {
	return utils.ExistsPath(path.Join(e.DataHome, "current"))
}

func (t *Task) StopWorking(e *Env) error {
	stop := time.Now()
	t.Stop = &stop
	d := stop.Sub(*t.Start)

	s, err := strconv.Atoi(fmt.Sprintf("%.0f", d.Seconds()))
	if err != nil {
		return err
	}

	t.Seconds = s
	err = t.Save(e)
	if err != nil {
		return err
	}

	err = os.Remove(path.Join(e.DataHome, "current"))
	if err != nil {
		return err
	}

	fmt.Printf("You have worked: %s\n", t.Duration())
	return nil
}

func (t *Task) Duration() time.Duration {
	return time.Duration(t.Seconds) * time.Second
}

func (t *Task) StartWorking(e *Env) error {
	if err := t.Save(e); err != nil {
		return err
	}
	return t.setWorking(e)
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

func (t *Task) setWorking(e *Env) error {
	return utils.WriteTo(path.Join(e.DataHome, "current"), []byte(t.PathToTask(e)))
}

func (t *Task) filename() string {
	return fmt.Sprintf("%v-%v_%v.json", t.Start.Format("20060102"), t.Start.Format("150405"), t.Project)
}

func (t *Task) directory(e *Env) string {
	return path.Join(e.DataHome, t.Start.Format("2006"), t.Start.Format("01"))
}
