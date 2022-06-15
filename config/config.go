package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	ServerPort   string `yaml:"ServerPort"`
	PingInterval int    `yaml:"PingInterval"`
	Servers      []struct {
		IP string `yaml:"IP"`
	} `yaml:"Servers"`
}

func (c *AppConfig) LoadConf() *AppConfig {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
