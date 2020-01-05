package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

var Data Config

var path = "./config/config/config.yaml"

type Config struct {
	Server cServer `yaml:"server"`
	Path cPath `yaml:"path"`
}

type cPath struct {
	Theme string `yaml:"theme"`
	Work string `yaml:"work"`
}

type cServer struct {
	Title string `yaml:"title"`
	Addr string `yaml:"addr"`
	Password string `yaml:"password"`
}

func (c *Config) Get() error {
	if f, err := os.Open(path); err != nil {
		return err
	} else {
		if err = yaml.NewDecoder(f).Decode(c); err != nil {
			return err
		}
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

func New() (*Config, error) {
	conf := &Config{}
	if err := conf.Get(); err != nil {
		return nil, err
	}
	return conf, nil
}


