package config

import (
	"log"

	"github.com/spf13/viper"
)

var Envs = LoadConfig()

type DBConfig struct {
	DB_URL string
}

type ServerConfig struct {
	PORT string
}

type JWTConfig struct {
	ACCESS_TOKEN_SECRET  string
	REFRESH_TOKEN_SECRET string
}

type CorsConfig struct {
	WHITELIST_DOMAINS string
}

type Config struct {
	DB     DBConfig
	Server ServerConfig
	JWT    JWTConfig
	CORS   CorsConfig
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	dbURL := GetEnv("DATABASE_URL", "")
	port := GetEnv("PORT", "3000")
	accessTokenSecret := GetEnv("ACCESS_TOKEN_SECRET", "")
	refreshTokenSecret := GetEnv("REFRESH_TOKEN_SECRET", "")
	whitelistDomains := GetEnv("WHITELIST_DOMAINS", "*")

	return &Config{
		DB: DBConfig{
			DB_URL: dbURL,
		},
		Server: ServerConfig{
			PORT: port,
		},
		JWT: JWTConfig{
			ACCESS_TOKEN_SECRET:  accessTokenSecret,
			REFRESH_TOKEN_SECRET: refreshTokenSecret,
		},
		CORS: CorsConfig{
			WHITELIST_DOMAINS: whitelistDomains,
		},
	}
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" && defaultValue != "" {
		return defaultValue
	}
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}
