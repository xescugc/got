package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type ConfigFile struct {
	Project string `json:"project"`
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
	data, err := ioutil.ReadFile(current_config_file)
	if err != nil {
		return nil, err
	}

	var cf ConfigFile
	json.Unmarshal(data, &cf)

	return &cf, nil
}
