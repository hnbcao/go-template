package config

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Host           string `envconfig:"HTTP_HOST"`
		Port           int    `envconfig:"HTTP_PORT"`
		Mode           string `envconfig:"LOG_MODE"`
		DatabaseSource string `envconfig:"DATABASE_SOURCE"`
		DatabaseSecret string `envconfig:"DATABASE_SECRET"`
	}
)

// String returns the configuration in string format.
func (c Config) String() string {
	out, _ := yaml.Marshal(c)
	return string(out)
}

// Environ returns the settings from the environment.
func environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if cfg.Host == "" {
		cfg.Host = "0.0.0.0"
	}
	if cfg.Port == 0 {
		cfg.Port = 8080
	}
	return cfg, err
}

//
func InitializeConfig() (Config, error) {
	var envFile string
	flag.StringVar(&envFile, "config", "config", "Read in a file of environment variables")
	flag.Parse()

	err := godotenv.Load(envFile)
	if err != nil {
		logrus.WithError(err).Warn("no config file")
	}

	return environ()
}
