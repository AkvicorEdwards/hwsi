package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

// Global configuration variables
var Data Config

// Configuration file directory
var path = "./config/config.yaml"

// Configuration structure
type Config struct {
	Server cServer `yaml:"server"`
	Path cPath `yaml:"path"`
}

// Path for Config
type cPath struct {
	Theme string `yaml:"theme"`
	Work string `yaml:"work"`
}

// Server for Config
type cServer struct {
	Title string `yaml:"title"`
	Addr string `yaml:"addr"`
	Password string `yaml:"password"`
}

// Get configuration information from a configuration file
func (c *Config) Get() error {
	if f, err := os.Open(path); err != nil {
		return err
	} else if err = yaml.NewDecoder(f).Decode(c); err != nil {
		return err
	}
	return nil
}

func (c *Config) String() string {
	byt,err := json.Marshal(c)
	if err != nil {
		log.Error(err)
	}
	return string(byt)
}

// Returns a new Config with configuration information
func New() (*Config, error) {
	conf := &Config{}
	if err := conf.Get(); err != nil {
		return nil, err
	}
	return conf, nil
}
