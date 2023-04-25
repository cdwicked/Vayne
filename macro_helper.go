package Vayne

import (
	"gopkg.in/yaml.v3"
	"os"
)

var globalMap map[string]string

func GetMap() map[string]string {
	return globalMap
}

func readConfig(configPath string) error {
	cp, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(cp, &globalMap)
	if err != nil {
		return err
	}
	return nil
}
