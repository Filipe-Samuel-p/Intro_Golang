package models

type Movie struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Genre     string `json:"genre"`
	Year      int    `json:"year"`
	Available bool   `json:"available"`
}
