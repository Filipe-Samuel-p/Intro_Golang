package handlers

import (
	"encoding/json"
	"http_server/models"
	"http_server/services"
	"net/http"
)

type MovieHandler struct {
	MovieService *services.MovieService
}

func (m MovieHandler) GetFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	films := m.MovieService.GetAllFilms()
	json.NewEncoder(w).Encode(films)
}

func (m MovieHandler) AddFilm(w http.ResponseWriter, r *http.Request) {
	var f models.Movie
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	filmCreated, err := m.MovieService.SaveANewFilm(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(filmCreated)
}
