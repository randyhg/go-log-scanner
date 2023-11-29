package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host              string `yaml:"Host"`
	ShowSql           bool   `yaml:"ShowSql"`
	MySqlUrl          string `yaml:"MySqlUrl"`
	MySqlMaxIdle      int    `yaml:"MySqlMaxIdle"`
	MySqlMaxOpen      int    `yaml:"MySqlMaxOpen"`
	SlaveMySqlUrl     string `yaml:"SlaveMySqlUrl"`
	SlaveMySqlMaxIdle int    `yaml:"SlaveMySqlMaxIdle"`
	SlaveMySqlMaxOpen int    `yaml:"SlaveMySqlMaxOpen"`
	BaseURL           string `yaml:"BaseURL"`
}

var Instance Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&Instance)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
