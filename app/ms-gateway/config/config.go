package config

import (
	"ms-gateway/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug bool `yaml:"isDebug" env-default:"true"`
	Listen  struct {
		Host string `yaml:"host" env:"HOST" env-default:"8181"`
		Port string `yaml:"port" env:"PORT" env-default:"127.0.0.1"`
	} `yaml:"listen"`
}

// var instance *Config
var once = sync.Once{} // for parsing this file just once

func GetConfig() *Config {
	instance := &Config{}

	once.Do(func() { // do it once. Singleton pattern
		logger := logging.GetLogger()

		logger.Infoln("Read application's config.")

		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)

			logger.Fatalln(err)
			logger.Fatalln(help)
		}
	})

	return instance
}
