package config

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env:"DB_HOST" default:"localhost"`
	Port     int    `yaml:"port" env:"DB_PORT" default:"5432"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Name     string `yaml:"name" env:"DB_NAME"`
	SSLMode  string `yaml:"sslMode" env:"DB_SSL_MODE" default:"disable"`
}
