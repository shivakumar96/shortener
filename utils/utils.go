package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Counter struct {
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		MaxCount int    `yaml:"maxcount"`
		Ranges   int    `yaml:"ranges"`
	} `yaml:"counter"`
	Worker struct {
		Port  string `yaml:"port"`
		Host  string `yaml:"host"`
		Count int    `yaml:"count"`
	} `yaml:"worker"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

func ReadConfig() (*Config, error) {
	f, err := os.Open("config.yml")
	if err != nil {
		return nil, err
	}

	defer f.Close()
	var config Config
	decoder := yaml.NewDecoder(f)
	decoder.Decode(&config)
	return &config, nil
}
