package gameoflife

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Rows       int     `yaml:"rows"`
	Cols       int     `yaml:"cols"`
	Iterations int     `yaml:"iterations"`
	Matrix     [][]int `yaml:"matrix"`
}

func InitializeConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
