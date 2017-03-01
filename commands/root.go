package commands

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "got",
		Short: "Tack your time",
		Long:  "Track the time you dedicate working to make exact reports",
	}
	DataHome       string
	ConfigHome     string
	ProjectsPath   string
	RootConfigPath string
)

type ConfigFile struct {
	Project string `json:"project"`
}

func NewConfigFile(project string) *ConfigFile {
	return &ConfigFile{Project: project}
}

type Projects map[string]string

func (p Projects) Save() error {
	return WriteStructTo(ProjectsPath, p)
}

func GetProjects() (Projects, error) {
	exists, err := ExistsPath(ProjectsPath)
	var projects Projects
	if err != nil {
		return nil, err
	}
	if exists {
		data, err := ioutil.ReadFile(ProjectsPath)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(data, &projects)
	} else {
		projects = make(map[string]string)
	}

	return projects, nil
}

func InitializeConfig() error {
	DataHome = os.Getenv("XDG_DATA_HOME")
	if len(DataHome) == 0 {
		DataHome = path.Join(os.Getenv("HOME"), ".local/share")
	}
	DataHome = path.Join(DataHome, "got")

	ConfigHome = os.Getenv("XDG_CONFIG_HOME")
	if len(ConfigHome) == 0 {
		ConfigHome = path.Join(os.Getenv("HOME"), ".config")
	}
	ConfigHome = path.Join(ConfigHome, "got")

	ProjectsPath = path.Join(DataHome, "projects.json")
	RootConfigPath = path.Join(ConfigHome, "got.json")

	for _, path := range []string{DataHome, ConfigHome} {
		exists, err := ExistsPath(path)
		if err != nil {
			return err
		}
		if !exists {
			os.Mkdir(path, 0755)
		}
	}

	exists, err := ExistsPath(RootConfigPath)
	if err != nil {
		return err
	}

	if !exists {
		err = WriteTo(RootConfigPath, []byte(`{}`))
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteTo(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	defer f.Close()
	f.Write(data)

	return nil
}

func WriteStructTo(path string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return WriteTo(path, b)
}

func ExistsPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
