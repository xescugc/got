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

// Task holds the information of a task that you have workder (or currently working)
type Task struct {
	Description string     `json:"description,omitempty"`
	Project     string     `json:"project"`
	Start       *time.Time `json:"start"`
	Stop        *time.Time `json:"stop,omitempty"`
	Seconds     int        `json:"seconds,omitempty"`
}

// NewTask creates a Task for the project
func NewTask(project string) *Task {
	return &Task{
		Project: project,
	}
}

// NewTaskFromCurrent loads the current task working on (on the 'current' file)
func NewTaskFromCurrent(e *Env) (*Task, error) {
	task_path, err := ioutil.ReadFile(path.Join(e.DataHome, "current"))
	if err != nil {
		return nil, err
	}
	return NewTaskFromPath(string(task_path))
}

// NewTaskFromPath loads a task from a path
func NewTaskFromPath(p string) (*Task, error) {
	data, err := ioutil.ReadFile(string(p))
	if err != nil {
		return nil, err
	}

	var t Task
	json.Unmarshal(data, &t)

	return &t, nil
}

// IsWorking validates if some task is working
func IsWorking(e *Env) (bool, error) {
	return utils.ExistsPath(path.Join(e.DataHome, "current"))
}

// StartWorking saves the Task and sets it to the 'current'
func (t *Task) StartWorking(e *Env) error {
	ti := time.Now()
	t.Start = &ti
	if err := t.Save(e); err != nil {
		return err
	}
	return t.setWorking(e)
}

// StopWorking sets the task as finished and removes the 'current'
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

// Duration returns the Seconds as a Duration
func (t *Task) Duration() time.Duration {
	return time.Duration(t.Seconds) * time.Second
}

// Save stores the Task to disk to the corresponding location
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

// PathToTask returns the path in which the Task is
func (t *Task) PathToTask(e *Env) string {
	return path.Join(t.directory(e), t.filename())
}

// setWorking creates the 'current' file
func (t *Task) setWorking(e *Env) error {
	return utils.WriteTo(path.Join(e.DataHome, "current"), []byte(t.PathToTask(e)))
}

// filename returns the filename of the Task
func (t *Task) filename() string {
	return fmt.Sprintf("%v-%v_%v.json", t.Start.Format("20060102"), t.Start.Format("150405"), t.Project)
}

// direcotry returns the directory i which the Task is
func (t *Task) directory(e *Env) string {
	return path.Join(e.DataHome, t.Start.Format("2006"), t.Start.Format("01"))
}
