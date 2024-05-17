package gameoflife

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Rows       int `yaml:"rows"`
	Cols       int `yaml:"cols"`
	Iterations int `yaml:"iterations"`
}

func InitializeConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
