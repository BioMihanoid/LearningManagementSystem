package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DbConfig     `yaml:"db"`
}

type ServerConfig struct {
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DbConfig struct {
	Host    string `yaml:"host" env-required:"true"`
	Port    string `yaml:"port" env-required:"true"`
	User    string `yaml:"user" env-required:"true"`
	Pass    string `yaml:"password" env-required:"true"`
	Dbname  string `yaml:"dbname" env-required:"true"`
	Sslmode string `yaml:"sslmode" env-default:"disable"`
}

//TODO: create fn do dsn

func ParseConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		logrus.Panic("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logrus.Panic("config file does not exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		logrus.Panic("cannot read configuration")
	}

	return &cfg
}
