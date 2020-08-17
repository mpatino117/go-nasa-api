package config

import (
	"github.com/kelseyhightower/envconfig"
)

type EnviromentVariables struct {
	NasaAuthKey string `required:"true" envconfig:"NASA_AUTH_KEY"`
	NasaUrl     string `required:"true" envconfig:"NASA_URL"`
}

func LoadEnv() (EnviromentVariables, error) {
	ev := EnviromentVariables{}
	err := envconfig.Process("", &ev)
	if err != nil {
		return EnviromentVariables{}, err
	}

	return ev, nil
}
