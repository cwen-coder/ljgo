package library

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type SiteConfig struct {
	Title     string `yaml:"title"`
	Introduce string `yaml:"introduce"`
	Limit     int    `yaml:"limit"`
	Theme     string `yaml:"theme"`
	URL       string `yaml:"url"`
	Comment   string `yaml:"comment"`
	Github    string `yaml:"github"`
	Facebook  string `yaml:"facebook"`
	Twitter   string `yaml:"twitter"`
}

type BuildConfig struct {
	Port    string   `yaml:"port"`
	Copy    []string `yaml:"copy"`
	Publish string   `yaml:"publish"`
}

type Config struct {
	Site  SiteConfig  `yaml:"site"`
	Build BuildConfig `yaml:"build"`
}

func ParseConfig(path string) (*Config, error) {
	var configT *Config

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Read config: %v", err)
	}

	if err = yaml.Unmarshal(data, &configT); err != nil {
		return nil, fmt.Errorf("Unmarshal config: %v", err)
	}

	return configT, nil
}
