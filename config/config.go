package config

import (
	"log"

	"github.com/lytics/confl"
)

var Conf Config

type Config struct {
	ExternalEndpoints `confl:"external_endpoints"`
	Service           `confl:"service"`
}

type ExternalEndpoints struct {
	ConfigTournamentsPath  string `confl:"config_tournaments"`
	FixturesTournamentPath string `confl:"fixtures_tournament"`
}

type Service struct {
	Port int `confl:"port"`
}

func Load(path string) {
	if _, err := confl.DecodeFile(path, &Conf); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}
	log.Printf("config initializes successfully")
}
