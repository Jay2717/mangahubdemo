package models

type Manga struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Genres      string `json:"genres"`
}
