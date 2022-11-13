package utils

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

func GetLocalConfig() (LocalConfig, error) {
	file, err := os.ReadFile("runtimer.yml")
	if err != nil {
		return LocalConfig{}, errors.New("cannot find config file (runrimer.yml)")
	}
	var config LocalConfig
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return LocalConfig{}, err
	}
	return config, err
}

func GetPrimaryRuntime(config LocalConfig) int {
	for i, r := range config.Runtimes {
		if r.Name == config.Primary {
			return i
		}
	}
	return 0
}

func GetRuntimeByName(name string) int {
	config, _ := GetLocalConfig()
	for i, r := range config.Runtimes {
		if r.Name == name {
			return i
		}
	}
	fallback := GetPrimaryRuntime(config)
	return fallback
}