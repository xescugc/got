package commands

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/xescugc/got/utils"
)

var (
	RootCmd = &cobra.Command{
		Use:   "got",
		Short: "Tack your time",
		Long:  "Track the time you dedicate working to make exact reports",
	}
)

func InitializeConfig() error {
	utils.DataHome = os.Getenv("XDG_DATA_HOME")
	if len(utils.DataHome) == 0 {
		utils.DataHome = path.Join(os.Getenv("HOME"), ".local/share")
	}
	utils.DataHome = path.Join(utils.DataHome, "got")

	utils.ConfigHome = os.Getenv("XDG_CONFIG_HOME")
	if len(utils.ConfigHome) == 0 {
		utils.ConfigHome = path.Join(os.Getenv("HOME"), ".config")
	}
	utils.ConfigHome = path.Join(utils.ConfigHome, "got")

	utils.ProjectsPath = path.Join(utils.DataHome, "projects.json")
	utils.RootConfigPath = path.Join(utils.ConfigHome, "got.json")

	for _, path := range []string{utils.DataHome, utils.ConfigHome} {
		exists, err := utils.ExistsPath(path)
		if err != nil {
			return err
		}
		if !exists {
			os.Mkdir(path, 0755)
		}
	}

	exists, err := utils.ExistsPath(utils.RootConfigPath)
	if err != nil {
		return err
	}

	if !exists {
		err = utils.WriteTo(utils.RootConfigPath, []byte(`{}`))
		if err != nil {
			return err
		}
	}

	return nil
}
