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
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig() // Ignore error in Railway
	viper.AutomaticEnv() 


	dbURL := viper.GetString("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	port := viper.GetString("PORT")
	if port == "" {
		port = "3000"
	}
	accessTokenSecret := viper.GetString("ACCESS_TOKEN_SECRET")
	if accessTokenSecret == "" {
		log.Fatal("ACCESS_TOKEN_SECRET is not set")
	}
	refreshTokenSecret := viper.GetString("REFRESH_TOKEN_SECRET")
	if refreshTokenSecret == "" {
		log.Fatal("REFRESH_TOKEN_SECRET is not set")
	}
	whitelistDomains := viper.GetString("WHITELIST_DOMAINS")
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
