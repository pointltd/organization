package config

import "os"

type AppConfig struct {
	env         string
	port        string
	databaseUrl string
	jwtSecret   string
}

func NewAppConfig() AppConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return AppConfig{
		env:         os.Getenv("ENV"),
		port:        port,
		databaseUrl: os.Getenv("DATABASE_URL"),
		jwtSecret:   os.Getenv("JWT_SECRET"),
	}
}

func (c AppConfig) Env() string {
	return c.env
}

func (c AppConfig) Port() string {
	return c.port
}

func (c AppConfig) DatabaseUrl() string {
	return c.databaseUrl
}

func (c AppConfig) JwtSecret() string {
	return c.jwtSecret
}
