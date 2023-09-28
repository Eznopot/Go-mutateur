package config

import (
	"log"
	"os"
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
	ScrollSpeed int  `yaml:"scroll_speed"`
}

type Developpement struct {
	Replication bool `yaml:"replication"`
}

type conf struct {
	Server        serverConf    `yaml:"server"`
	Client        clientConf    `yaml:"client"`
	Config        config        `yaml:"config"`
	Developpement Developpement `yaml:"developpement"`
}

// The GetConfig function reads a YAML file, unmarshals it into a struct, and returns an instance of
// that struct.
func GetConfig() *conf {
	once.Do(func() {
		var confTmp conf
		yamlFile, err := os.ReadFile("config.yml")
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
