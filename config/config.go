package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Api      Api
	Database Database
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return &Config{Api: API(), Database: DataStore()}
}
