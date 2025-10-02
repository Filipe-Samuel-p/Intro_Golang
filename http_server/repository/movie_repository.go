package repository

import (
	"http_server/models"
)

type Repository struct {
	db []models.Movie
}

func (r Repository) FindAll() *[]models.Movie {

	return &r.db
}

func (r *Repository) Save(movie models.Movie) models.Movie {
	r.db = append(r.db, movie)
	return movie
}
