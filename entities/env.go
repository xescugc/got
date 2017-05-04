package entities

import (
	"os"
	"path"

	"github.com/xescugc/got/utils"
)

type Env struct {
	DataHome       string
	ConfigHome     string
	ProjectsPath   string
	RootConfigPath string
}

func NewEnv() (*Env, error) {
	env := &Env{}

	env.DataHome = os.Getenv("XDG_DATA_HOME")
	if len(env.DataHome) == 0 {
		env.DataHome = path.Join(os.Getenv("HOME"), ".local/share")
	}
	env.DataHome = path.Join(env.DataHome, "got")

	env.ConfigHome = os.Getenv("XDG_CONFIG_HOME")
	if len(env.ConfigHome) == 0 {
		env.ConfigHome = path.Join(os.Getenv("HOME"), ".config")
	}
	env.ConfigHome = path.Join(env.ConfigHome, "got")

	env.ProjectsPath = path.Join(env.DataHome, "projects.json")
	env.RootConfigPath = path.Join(env.ConfigHome, "got.json")

	err := env.SetUp()

	return env, err
}

func (e *Env) SetUp() error {
	for _, path := range []string{e.DataHome, e.ConfigHome} {
		exists, err := utils.ExistsPath(path)
		if err != nil {
			return err
		}
		if !exists {
			os.Mkdir(path, 0755)
		}
	}

	exists, err := utils.ExistsPath(e.RootConfigPath)
	if err != nil {
		return err
	}

	if !exists {
		err = utils.WriteTo(e.RootConfigPath, []byte(`{}`))
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Env) Clean() {
	os.RemoveAll(e.DataHome)
	os.RemoveAll(e.ConfigHome)
}
