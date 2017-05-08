package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/xescugc/got/utils"
)

// Projects hold the information of all the known projects
// initialized.
// The structure is project_name => path_to_config_file
type Projects map[string]string

// GetProjects returns all the Projects from the projects.json file
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

// Save stores the Projects to the the projects.json file
func (p Projects) Save(e *Env) error {
	return utils.WriteStructTo(e.ProjectsPath, p)
}

// Clean removes all the Projects config files that it knows of
func (p Projects) Clean(e *Env) {
	for _, path := range p {
		os.Remove(path)
	}
}
