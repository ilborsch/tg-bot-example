package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"strconv"
)

type Config struct {
	Token         string `yaml:"token"`
	Timeout       int    `yaml:"timeout"`
	BackendConfig `yaml:"backend"`
}

type BackendConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

func MustLoad() *Config {
	path := getConfigPath()
	return MustLoadByPath(path)
}

func getConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("")
	}
	return path
}

func MustLoadByPath(path string) *Config {
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(fmt.Sprintf("the file with path %s does not exist", path))
	}
	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("error parsing config file")
	}

	backendPort, err := strconv.Atoi(os.Getenv("BACKEND_PORT"))
	if err != nil {
		config.BackendConfig.Port = backendPort
	}

	backendHost := os.Getenv("BACKEND_HOST")
	if backendHost != "" {
		config.BackendConfig.Host = backendHost
	}

	return &config
}
