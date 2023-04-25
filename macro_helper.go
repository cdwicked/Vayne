package Vayne

import (
	"gopkg.in/yaml.v3"
	"os"
)

var globalMap map[string]string

// GetMap get macro config
func GetMap() map[string]string {
	if globalMap == nil {
		return make(map[string]string, 0)
	}
	return globalMap
}

// ReadMacroConfig read file
func ReadMacroConfig(configPath string) error {
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
