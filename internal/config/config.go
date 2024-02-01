package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	MQ struct {
		URL  string `mapstructure:"url"`
		Name string `mapstructure:"name"`
	} `mapstructure:"mq"`
	Mongo struct {
		DbName   string `mapstructure:"dbname"`
		URI      string `mapstructure:"uri"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"db"`
}

func NewConfig(folder string, filename string) (*Config, error) {
	cfg := new(Config)

	viper.AddConfigPath(folder)
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.WithField("config", "wrong config").Fatal(err)

		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		log.WithField("config", "wrong unmarshalling").Fatal(err)

		return nil, err
	}

	return cfg, nil
}
