package model

type MovieWithGenres struct {
	Title    string   `json:"title"`
	Overview string   `json:"overview"`
	Genres   []string `json:"genres"`
}
