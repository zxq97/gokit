package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func LoadYaml(path string, v interface{}) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bs, v)
}
