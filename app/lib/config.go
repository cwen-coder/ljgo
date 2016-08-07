package lib

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
}

type AuthorConfig struct {
	Name string `yaml:"name"`
}

type BuildConfig struct {
	Port    string   `yaml:"port"`
	Copy    []string `yaml:"copy"`
	Publish string   `yaml:"publish"`
}

type Config struct {
	Site   SiteConfig   `yaml:"site"`
	Author AuthorConfig `yaml:"author"`
	Build  BuildConfig  `yaml:"build"`
}

//func newConfig() *Config {
//return &Config{}
//}

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
