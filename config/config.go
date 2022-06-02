package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

type Env struct {
	Port string
}

// LoadConfig
func LoadConfig() (e Env) {

	_, result := os.LookupEnv("PORT")
	if !result {
		fmt.Println("Not found: PORT")
		// Todo: 5xx を返却
	}
	var goenv Env
	err := envconfig.Process("", &goenv)
	if err != nil {
		log.Fatal(err.Error())
	}

	return goenv
}
