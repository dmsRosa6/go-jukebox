package api

import (
	"encoding/json"
	"net/http"

	"github.com/dmsrosa/jukebox/internal/service"
)

// Handlers wraps your jukebox service for use in API endpoints.
type Handlers struct {
	Service *service.Jukebox
}

// NewHandlers returns a new Handlers instance with the given service.
func NewHandlers(s *service.Jukebox) *Handlers {
	return &Handlers{Service: s}
}

// GetSongs returns all songs in the jukebox.
func (h *Handlers) GetSongs(w http.ResponseWriter, r *http.Request) {
	songs := h.Service.GetSongs() 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}

// GetCurrent gets the current song playing
func (h *Handlers) GetCurrent(w http.ResponseWriter, r *http.Request) {

	song, err := h.Service.Current()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)	
		return 
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}

// CreateSong adds a new song to the jukebox.
func (h *Handlers) EnqueueSong(w http.ResponseWriter, r *http.Request) {
	var song service.Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Service.Enqueue(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(song)
}

// DeleteSong removes a song from the jukebox.
func (h *Handlers) DequeueSong(w http.ResponseWriter, r *http.Request) {	
	if _ , err := h.Service.Dequeue(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
