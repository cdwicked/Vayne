package util

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

var values, release map[string]interface{}

// GetMap get macro config
// if map is nil , it will return map with length 0
func GetMap(name string) map[string]interface{} {
	ret := func(m map[string]interface{}) map[string]interface{} {
		if m == nil {
			return make(map[string]interface{}, 0)
		}
		return m
	}
	switch strings.ToLower(name) {
	case "values":
		return ret(values)
	case "release":
		return ret(release)
	default:
		return make(map[string]interface{}, 0)
	}
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
	err = yaml.Unmarshal(cp, &values)
	if err != nil {
		return err
	}
	return nil
}

// ReadValue .Values.ccc.bbb
func ReadValue(xPath string) (interface{}, error) {
	xPaths := strings.Split(xPath, ".")
	temp := GetMap(xPaths[1])
	var a interface{}
	for _, path := range xPaths[2:] {
		a = temp[path]
		switch a.(type) {
		case map[string]interface{}:
			temp = a.(map[string]interface{})
		default:
			return a, nil
		}
	}
	return a, nil
}

// InitTestConfig
// replace the macro in test config
func InitTestConfig(configPath string) error {
	return nil
}
