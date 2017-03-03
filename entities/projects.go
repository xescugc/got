package entities

import (
	"encoding/json"
	"io/ioutil"

	"github.com/xescugc/got/utils"
)

type Projects map[string]string

func (p Projects) Save() error {
	return utils.WriteStructTo(utils.ProjectsPath, p)
}

func GetProjects() (Projects, error) {
	exists, err := utils.ExistsPath(utils.ProjectsPath)
	var projects Projects
	if err != nil {
		return nil, err
	}
	if exists {
		data, err := ioutil.ReadFile(utils.ProjectsPath)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(data, &projects)
	} else {
		projects = make(map[string]string)
	}

	return projects, nil
}
