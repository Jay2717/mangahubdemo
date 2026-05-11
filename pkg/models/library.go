package models

type LibraryItem struct {
	MangaID        int    `json:"manga_id"`
	Title          string `json:"title"`
	Status         string `json:"status"`
	CurrentChapter int    `json:"current_chapter"`
}
