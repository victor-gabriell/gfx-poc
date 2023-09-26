package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ApplicationConfig struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

func NewApplicationConfig() *ApplicationConfig {
	read, err := os.ReadFile("./dev.yaml")
	if err != nil {
		os.Exit(1)
	}

	var config ApplicationConfig

	if err = yaml.Unmarshal(read, &config); err != nil {
		log.Fatalf("error while unmarshal config %+v", err)
	}

	return &config
}
