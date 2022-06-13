package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Env struct {
	Port string
}

func LoadConfig() Env {
	if _, err := os.LookupEnv("PORT"); !err {
		zap.S().Error("Not found an env value: PORT")
	}

	var goenv Env

	err := envconfig.Process("", &goenv)
	if err != nil {
		log.Fatal(err.Error())
	}

	return goenv
}
