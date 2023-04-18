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
	mongo struct {
		MongoHost      string `yaml:"mongoHost"`
		MongoHort      string `yaml:"mongoPort"`
		MongoUser      string `yaml:"mongoUser"`
		MongoPassword  string `yaml:"mongoPassword"`
		MongoDatabase  string `yaml:"mongoDatabase"`
		MongoUriScheme string `yaml:"mongoUriScheme" env-default:"mongodb"`
	}
}

type ConfigEnv struct {
	IsDebug        bool   `yaml:"isDebug" env:"IS_DEBUG" env-default:"true"`
	Host           string `yaml:"host" env:"HOST" env-default:"8181"`
	Port           string `yaml:"port" env:"PORT" env-default:"127.0.0.1"`
	MongoHost      string `env:"MONGO_HOST"`
	MongoHort      string `env:"MONGO_PORT"`
	MongoUser      string `env:"MONGO_USER"`
	MongoPassword  string `env:"MONGO_PASSWORD"`
	MongoDatabase  string `env:"MONGO_DATABASE"`
	MongoUriScheme string `env:"MONGO_URI_SCHEME" env-default:"mongodb"`
}

var (
	instance *ConfigEnv
	once     sync.Once
)

func GetConfig() *ConfigEnv {
	once.Do(func() { // do it once. Singleton pattern
		logger := logging.GetLogger()

		logger.Infoln("Read application's config.")
		instance = &ConfigEnv{}

		//if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
		if err := cleanenv.ReadEnv(instance); err != nil {
			//help, _ := cleanenv.GetDescription(instance, nil)

			//logger.Fatalln(err)
			//logger.Fatalln(help)
		}
	})

	return instance
}
