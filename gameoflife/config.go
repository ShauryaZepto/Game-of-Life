package gameoflife

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Rows       int     `yaml:"rows"`
	Cols       int     `yaml:"cols"`
	Iterations int     `yaml:"iterations"`
	Matrix     [][]int `yaml:"matrix"`
}

func InitializeConfig(path string) (*Config, error) {
	data, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}

	var config Config
	e = yaml.Unmarshal(data, &config)
	if e != nil {
		return nil, e
	}

	return &config, nil
}
