package utils

import (
	"errors"
	"fmt"
	"os"
	"runtimer/constants"

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

func GetPrimaryRuntime(config LocalConfigType) (*constants.RuntimesType, error) {
	for i, r := range config.Runtimes {
		if r.Name == config.Primary {
			return &config.Runtimes[i], nil
		}
	}
	return nil, fmt.Errorf("cannot find runtime %s", config)
}

func GetCustomRuntimeByName(name string) (*constants.RuntimesType, error) {
	config, _ := GetLocalConfig()
	for i, r := range config.Runtimes {
		if r.Name == name {
			return &config.Runtimes[i], nil
		}
	}
	return nil, fmt.Errorf("cannot find runtime %s", name)
}