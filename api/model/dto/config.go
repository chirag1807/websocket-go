package dto

type Config struct {
	Database  Database `mapstructure:",squash"`
}

type Database struct {
	Username string `mapstructure:"DATABASE_USERNAME"`
	Password string `mapstructure:"DATABASE_PASSWORD"`
	Port     string `mapstructure:"DATABASE_PORT"`
	Name     string `mapstructure:"DATABASE_NAME"`
	SSLMode  string `mapstructure:"DATABASE_SSLMODE"`
}