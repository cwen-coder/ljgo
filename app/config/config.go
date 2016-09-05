package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/urfave/cli"

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

type ServeConfig struct {
	Addr string `yaml:"addr"`
}

type PublishConfig struct {
	Cmd string `yaml:"cmd"`
}

type Config struct {
	Site       SiteConfig    `yaml:"site"`
	Serve      ServeConfig   `yaml:"serve"`
	Publish    PublishConfig `yaml:"publish"`
	RootPath   string
	ThemePath  string
	SourcePath string
	PublicPath string
}

func (c *Config) parseConfig() error {
	data, err := ioutil.ReadFile(filepath.Join(c.RootPath, "config.yml"))
	if err != nil {
		return fmt.Errorf("Read config: %v", err)
	}

	if err = yaml.Unmarshal(data, &c); err != nil {
		return fmt.Errorf("Unmarshal config: %v", err)
	}

	return nil
}

func New(c *cli.Context) (*Config, error) {
	var config = &Config{
		RootPath: ".",
	}
	if len(c.Args()) > 0 {
		config.RootPath = c.Args()[0]
	}
	err := config.parseConfig()
	if err != nil {
		return nil, fmt.Errorf("parse config: %v", err)
	}
	config.ThemePath = filepath.Join(config.RootPath, config.Site.Theme)
	config.SourcePath = filepath.Join(config.RootPath, "source")
	config.PublicPath = filepath.Join(config.RootPath, "public")
	if c.String("addr") != "" {
		config.Serve.Addr = c.String("addr")
	}
	return config, nil
}
