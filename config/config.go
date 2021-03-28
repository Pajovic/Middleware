package config

import (
	"log"

	"github.com/mcuadros/go-defaults"

	"github.com/lytics/confl"
)

var Conf Config

type Config struct {
	ExternalEndpoints `confl:"external_endpoints"`
	Service           `confl:"service"`
	Limit             int `confl:"limit" default:"5"`
}

type ExternalEndpoints struct {
	ConfigTournamentsPath  string `confl:"config_tournaments"`
	FixturesTournamentPath string `confl:"fixtures_tournament"`
}

type Service struct {
	Port int `confl:"port" default:"8080"`
}

func Load(path string) {
	if _, err := confl.DecodeFile(path, &Conf); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}
	defaults.SetDefaults(&Conf)

	log.Printf("config initialized successfully:")
	log.Printf("Tournaments source: %s", Conf.ConfigTournamentsPath)
	log.Printf("Matches source: %s", Conf.FixturesTournamentPath)
	log.Printf("Port: %d", Conf.Port)
	log.Printf("Limit: %d", Conf.Limit)
}
