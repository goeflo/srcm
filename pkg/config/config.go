package config

import (
	"github.com/spf13/viper"
)

var GlobalConfig Config = Config{}

type Config struct {
	HomepageName     string `mapstructure:"HOMEPAGE_NAME"`
	DbSqliteFilename string `mapstructure:"DB_SQLITE_FILENAME"`
	HttpPort         string `mapstructure:"HTTP_PORT"`
	RestPort         string `mapstructure:"REST_PORT"`
	LogLevel         string `mapstructure:"LOG_LEVEL"`
}

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		return err
	}
	return nil
}
