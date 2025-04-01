package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router with our REST API endpoints.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Songs endpoints
	router.HandleFunc("/jukebox", getJukeboxSongs).Methods("GET")
	router.HandleFunc("/songs/{name}", getSongsByName).Methods("GET")
	router.HandleFunc("/jukebox/{id}", addSongJukebox).Methods("POST")
	router.HandleFunc("/jukebox/skip", skipSongJukebox).Methods("POST")
	router.HandleFunc("/jukebox/clear", clearJukebox).Methods("POST")
	router.HandleFunc("/jukebox/shuffle", shuffleJukebox).Methods("POST")

	// Serve static files or documentation if needed
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("docs"))))

	return router
}
