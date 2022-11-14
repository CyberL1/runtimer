package utils

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

func GetLocalConfig() (LocalConfigType, error) {
	file, err := os.ReadFile("runtimer.yml")
	if err != nil {
		return LocalConfigType{}, errors.New("cannot find config file (runtimer.yml)")
	}
	var config LocalConfigType
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return LocalConfigType{}, err
	}
	return config, err
}

func GetPrimaryRuntime(config LocalConfigType) int {
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