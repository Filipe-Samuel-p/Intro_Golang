package main

import (
	"http_server/handlers"
	"http_server/repository"
	"http_server/services"
	"log"
	"net/http"
)

func main() {

	movieHandler := &handlers.MovieHandler{
		MovieService: &services.MovieService{
			Repository: &repository.Repository{},
		},
	}

	http.HandleFunc("/movie", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			movieHandler.GetFilms(w, r)
		} else if r.Method == http.MethodPost {
			movieHandler.AddFilm(w, r)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
