package config

import "github.com/kelseyhightower/envconfig"

type Postgres struct {
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
	Name     string `envconfig:"NAME"`
}

func NewPostgresConf() *Postgres {
	config := new(Postgres)
	envconfig.MustProcess("postgres", config)
	return config
}
