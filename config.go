package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type DB struct {
	Driver     string `yaml:"driver"`
	Connection string `yaml:"connection"`
}

type Config struct {
	DB   `yaml:"db"`
	Host string `yaml:"host"`
}

func rebindEnv(c *Config) {
	strings := make(map[string]*string)
	strings["DB_CONNECTION"] = &c.DB.Connection
	strings["DB_DRIVER"] = &c.DB.Driver

	for name, pointer := range strings {
		value := os.Getenv(name)
		if value != "" {
			*pointer = value
		}
	}
}

var getConfig = func() func() *Config {
	c := new(Config)
	filepath := "config.yml"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	rebindEnv(c)
	return func() *Config {
		return c
	}
}()
