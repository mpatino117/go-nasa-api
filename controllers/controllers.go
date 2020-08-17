package controllers

import (
	"log"
	"nasa-api/config"
	"net/http"
)

func TechPort(ev config.EnviromentVariables) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Print(ev.NasaAuthKey)
		log.Print(ev.NasaUrl)
	}
}
