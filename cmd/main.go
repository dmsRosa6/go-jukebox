package main

import (
	"log"
	"net/http"

	"github.com/dmsrosa/jukebox/config"
	"github.com/dmsrosa/jukebox/internal/api"
	"github.com/dmsrosa/jukebox/internal/service"
)

func main() {
	options := config.LoadDefaultOptions()
	conf := config.LoadDefaultConfig()

	jukebox := service.NewJukebox(options)
	router := api.NewRouter(jukebox, conf)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(conf.Port, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
