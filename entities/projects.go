package entities

import (
	"encoding/json"
	"io/ioutil"

	"github.com/xescugc/got/utils"
)

type Projects map[string]string

func GetProjects(e *Env) (Projects, error) {
	exists, err := utils.ExistsPath(e.ProjectsPath)
	if err != nil {
		return nil, err
	}

	var projects Projects
	if exists {
		data, err := ioutil.ReadFile(e.ProjectsPath)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(data, &projects)
	} else {
		projects = make(map[string]string)
	}

	return projects, nil
}

func (p Projects) Save(e *Env) error {
	return utils.WriteStructTo(e.ProjectsPath, p)
}
