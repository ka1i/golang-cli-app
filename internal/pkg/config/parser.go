package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	filename string
	Error    error
}

func (c *config) Init() {
	options, err := configParser(c.filename)
	if err != nil {
		c.Error = err
		return
	}
	Cfg.Set(&options)
}

var Configure = getConfig()

func getConfig() *config {
	return &config{
		filename: "configs.yaml",
		Error:    nil,
	}
}

func configParser(filename string) (Options, error) {
	options := Options{}

	configFile, err := os.ReadFile(filename)
	if err != nil {
		return options, fmt.Errorf("%s files Not Found: %v", filename, err)
	}

	err = yaml.Unmarshal(configFile, &options)
	if err != nil {
		return options, fmt.Errorf("%s files is corrupted: %v", filename, err)
	}
	return options, err
}
