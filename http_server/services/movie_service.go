package services

import (
	"errors"
	"http_server/models"
	repos "http_server/repository"
	"strings"
)

type MovieService struct {
	Repository *repos.Repository
}

func (u *MovieService) SaveANewFilm(newFilm models.Movie) (models.Movie, error) {
	if strings.TrimSpace(newFilm.Title) == "" {
		return models.Movie{}, errors.New("título é obrigatório")
	}
	if newFilm.Year < 1900 {
		return models.Movie{}, errors.New("ano inválido")
	}
	return u.Repository.Save(newFilm), nil
}

func (u MovieService) GetAllFilms() []models.Movie {
	return *u.Repository.FindAll()
}
