package Vayne

import (
	"gopkg.in/yaml.v3"
	"os"
)

var globalMap map[string]string

// GetMap get macro config
// if map is nil , it will return map with length 0
func GetMap() map[string]string {
	if globalMap == nil {
		return make(map[string]string, 0)
	}
	return globalMap
}

// ReadMacroConfig
// read file from configPath,configPath file is yaml,if the file is not yaml,it will return error
// and file yaml like this
//
// value: 1
// baseurl: "http://127.0.0.1:8080"
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
