package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/xescugc/got/utils"
)

type ConfigFile struct {
	Project string `json:"project"`

	Exists bool   `json:"-"`
	Path   string `json:"-"`
}

func NewConfigFile(project string) *ConfigFile {
	return &ConfigFile{Project: project}
}

func GetConfig() (*ConfigFile, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	current_config_file := path.Join(wd, ".got.json")
	exists, err := utils.ExistsPath(current_config_file)
	if err != nil {
		return nil, err
	}

	if !exists {
		paths := strings.Split(wd, "/")
		project_name := paths[len(paths)-1]
		c := NewConfigFile(project_name)
		c.Path = current_config_file
		c.Exists = exists
		return c, nil
	} else {
		data, err := ioutil.ReadFile(current_config_file)
		if err != nil {
			return nil, err
		}

		var c ConfigFile
		json.Unmarshal(data, &c)

		c.Exists = true

		return &c, nil
	}

}

func (c *ConfigFile) Save() error {
	return utils.WriteStructTo(c.Path, c)
}
