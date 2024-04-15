package config

type Config struct {
	LocalURL string `mapstructure:"THIS_APP_URL"`
	Path     string `mapstructure:"PATH"`
	Logger   Logger `mapstructure:"LOGGER"`
}

type Logger struct {
	Production  string `mapstructure:"PRODUCTION"`
	Development string `mapstructure:"DEVELOPMENT"`
}
