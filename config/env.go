package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	if accessTokenSecret == "" {
		log.Fatal("ACCESS_TOKEN_SECRET is not set")
	}
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	if refreshTokenSecret == "" {
		log.Fatal("REFRESH_TOKEN_SECRET is not set")
	}
	whitelistDomains := os.Getenv("WHITELIST_DOMAINS")
	if whitelistDomains == "" {
		whitelistDomains = "*"
	}
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
