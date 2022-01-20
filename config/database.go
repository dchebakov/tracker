package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Database struct {
	Host              string
	Port              string
	Name              string
	User              string
	Pass              string
	MaxConnectionPool int `default:"4"`
}

func DataStore() Database {
	var db Database
	envconfig.MustProcess("DB", &db)

	return db
}
