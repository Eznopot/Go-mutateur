package config

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

var once sync.Once
var instance *conf

type serverConf struct {
	Port string `yaml:"port"`
}

type clientConf struct {
	Port    string `yaml:"port"`
	Address string `yaml:"address"`
}

type config struct {
	SmoothMode  bool `yaml:"smooth_mode"`
	SmoothDelay int  `yaml:"smooth_delay"`
}

type conf struct {
	Server serverConf `yaml:"server"`
	Client clientConf `yaml:"client"`
	Config config     `yaml:"config"`
}

func GetConfig() *conf {
	once.Do(func() {
		var confTmp conf
		yamlFile, err := ioutil.ReadFile("config.yml")
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, &confTmp)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
		instance = &confTmp
	})
	return instance
}
