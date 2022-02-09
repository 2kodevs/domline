package configs

// LogConfig : Struct Represent configuration data for logging
type LogConfig struct {
	Level string `mapstructure:"logLevel"`
}

type repoConfig struct {
	URL        string `mapstructure:"url"`
	PlayerName string `mapstructure:"playerName"`
}

type Players []repoConfig
