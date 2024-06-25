package configs

import (
	"fmt"
	"os"

	"github.com/richardktran/grpc-golang/configs"
	"gopkg.in/yaml.v2"
)

type ConfigFilePath string

type Config struct {
	HTTP HTTP `yaml:"http"`
	GRPC GRPC `yaml:"grpc"`
}

func NewConfig(filePath ConfigFilePath) (Config, error) {
	var (
		defaultConfig = configs.DefaultConfig
		config        = Config{}
		err           error
	)

	if filePath != "" {
		defaultConfig, err = os.ReadFile(string(filePath))

		if err != nil {
			return Config{}, fmt.Errorf("failed to read yaml file: %w", err)
		}
	}

	err = yaml.Unmarshal(defaultConfig, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal Yaml: %w", err)
	}

	return config, nil
}
