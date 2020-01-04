package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Parent struct {
		Address string `yaml:"address"`
	} `yaml:"parent"`
}

func LoadConfig() Config {
	workingDirectory, _ := os.Getwd()
	f, err := os.Open(workingDirectory + "/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
