package configs

import (
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

func GetConfigs(folderpath string) (Players, error) {
	files, err := ioutil.ReadDir(folderpath)
	if err != nil {
		return nil, err
	}

	var players Players
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yml") {
			viper.AddConfigPath(folderpath)
			viper.SetConfigType("yml")
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
