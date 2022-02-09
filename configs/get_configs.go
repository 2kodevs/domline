package configs

import (
	"io/ioutil"
	"strings"

	//"github.com/2kodevs/domline/internal/utils"

	"github.com/2kodevs/domline/internal/utils"
	"github.com/spf13/viper"
)

func GetConfigs(repo string, folderpath string) (Players, error) {
	err := utils.WGet(repo, folderpath)
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(folderpath)
	if err != nil {
		return nil, err
	}

	var players Players
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			viper.AddConfigPath(folderpath)
			viper.SetConfigType("yaml")
			viper.SetConfigName(file.Name())
			err := viper.ReadInConfig()
			if err != nil {
				return nil, err
			}

			var config repoConfig
			err = viper.Unmarshal(&config)
			if err != nil {
				return nil, err
			}
			players = append(players, config)
		}
	}

	return players, nil
}
