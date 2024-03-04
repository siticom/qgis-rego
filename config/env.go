package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	Port               uint   `env:"PORT" envDefault:"8080"`
	PluginDir          string `env:"PLUGIN_DIR" envDefault:"./plugins"`
	MinimumQgisVersion string `env:"MINIMUM_QGIS_VERSION" envDefault:"3.28"`
	ShowIcons          bool   `env:"SHOW_ICONS" envDefault:"true"`
}

var (
	Env Configuration
	// Store the external server url including protocol
	ServerURL string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Warn("Error loading .env file")
	}
	err = env.Parse(&Env)
	if err != nil {
		log.WithError(err).Fatal("Error parsing environment variables")
	}

	log.Infof("Configuration loaded: %+v", Env)
}
