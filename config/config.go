package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

// Global configuration variables
var Data Config

// Configuration structure
type Config struct {
	Server cServer `yaml:"server"`
	Path cPath `yaml:"path"`
}

// Path for Config
type cPath struct {
	Theme string `yaml:"theme"`
	Work string `yaml:"work"`
	Upload string `yaml:"upload"`
}

// Server for Config
type cServer struct {
	Title string `yaml:"title"`
	Addr string `yaml:"addr"`
	Password string `yaml:"password"`
}

// Get configuration information from a configuration file
func (c *Config) GetByFile(path string) error {
	if f, err := os.Open(path); err != nil {
		return err
	} else if err = yaml.NewDecoder(f).Decode(c); err != nil {
		return err
	}
	return nil
}

func (c *Config) GetByMap(args map[string]string) error {
	for k, v := range args {
		switch k {
		case "title":
			c.Server.Title = v
		case "port":
			c.Server.Addr = ":" + v
		case "password":
			c.Server.Password = v
		case "work":
			c.Path.Work = v
		case "upload":
			c.Path.Upload = v
		case "theme":
			c.Path.Theme = v
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

