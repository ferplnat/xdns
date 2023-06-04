package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port int `yaml:"port"`
}

func Load(path string) (Config, error) {
    config := Config{}

	configYaml, err := ioutil.ReadFile(path)
    if err != nil {
        return config, err 
    }
    err = yaml.Unmarshal(configYaml, &config)
    if err != nil {
        return config, err        
    }

    return config, nil
}
