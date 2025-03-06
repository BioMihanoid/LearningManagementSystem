package config

type Config struct {
	DbConfig
}

type DbConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Dbname  string
	Sslmode string
}
