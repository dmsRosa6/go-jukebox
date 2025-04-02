package api

import (
	"github.com/dmsrosa/jukebox/config"
	"github.com/dmsrosa/jukebox/internal/service"
	"github.com/gorilla/mux"
)

// NewRouter creates a new Gorilla Mux router and registers routes.
func NewRouter(jb *service.Jukebox, conf config.Config) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	h := NewHandlers(jb)

	// Register endpoints.
	router.HandleFunc("/api/songs", h.GetSongs).Methods("GET")
	router.HandleFunc("/api/current", h.GetCurrent).Methods("GET")
	router.HandleFunc("/api/queue/enqueue", h.EnqueueSong).Methods("PUT")
	router.HandleFunc("/api/queue/dequeue", h.DequeueSong).Methods("DELETE")

	return router
}
