package config

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type Postgres struct {
	User    string `yaml:"user"`
	Port    string `yaml:"port"`
	Pass    string `yaml:"pass"`
	Host    string `yaml:"host"`
	Dbname  string `yaml:"dbname"`
	Sslmode string `yaml:"sslmode"`
}

type Migration struct {
	Path string `yaml:"path"`
}

type Config struct {
	Postgres  Postgres  `yaml:"postgres"`
	Migration Migration `yaml:"migration"`
}

func Init(path string) (*Config, error) {
	content, err := readYamlFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(content, cfg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal")
	}

	return cfg, nil
}

func readYamlFile(path string) ([]byte, error) {
	data, err := os.ReadFile(string(path))
	if err != nil {
		return nil, fmt.Errorf("no read file")
	}
	return data, nil
}
