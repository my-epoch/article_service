package config

import (
	configReader "github.com/ShuffleBoy/goapp/config_reader"
	"github.com/ShuffleBoy/goapp/logger"
)

func NewConfigProvider(log *logger.Logger) (*Config, error) {
	config := new(Config)
	err := configReader.LoadConfig("config.yaml", &config, new(configReader.YamlProvider), new(configReader.IoReader))
	if err != nil {
		log.Fatalf("cannot load config: %e", err)
	}
	log.Info("config loaded")
	return config, err
}
