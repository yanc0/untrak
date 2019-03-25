package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Load Kuntrak config from path
func Load(path string) (*Config, error) {
	var cfg Config

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
