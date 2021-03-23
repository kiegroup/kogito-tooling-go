package config

import (
	"io/ioutil"
	"log"

	"github.com/adrielparedes/kogito-local-server/pkg/utils"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Runner struct {
		Location string `yaml:"location"`
	} `yaml:"runner"`
	Proxy struct {
		IP   string `yaml:"ip"`
		Port string `yaml:"port"`
	} `yaml:"proxy"`
	Modeler struct {
		Link string `yaml:"link"`
	} `yaml:"modeler"`
}

func (c *Config) GetConfig() *Config {

	yamlFile, err := ioutil.ReadFile(utils.GetBaseDir() + "/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
