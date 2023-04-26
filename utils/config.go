package utils

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PORT                       int
	UPDATE_INTERVAL_MILLIS     int
	SERVICE_COMMISSION_PERCENT float64
}

func LoadConfig(path string) (config Config, err error) {
	if err = envconfig.Process("", &config); err != nil {
		log.Fatal(err.Error())
	}
	return
}
