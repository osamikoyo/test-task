package config

import "os"

type Config struct{
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
	SERVER_HOST string
	SERVER_PORT string
}

func Load() *Config {
	dbhost := os.Getenv("DB_HOST")
	dbport:= os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_POST")
	return &Config{
		DB_HOST: dbhost,
		DB_PORT: dbport,
		DB_PASSWORD: dbpassword,
		DB_USER: dbuser,
		DB_NAME: dbname,
		SERVER_HOST: host,
		SERVER_PORT: port,
	}
}