package config

import "time"

type Config struct {
	ServerConfig
	DbConfig
}

type ServerConfig struct {
	PortServ    string
	Timeout     time.Duration
	IdleTimeout time.Duration
}

type DbConfig struct {
	Host    string
	PortDb  string
	User    string
	Pass    string
	Dbname  string
	Sslmode string
}
