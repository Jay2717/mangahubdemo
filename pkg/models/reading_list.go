package models

type ReadingList struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	MangaID  string `json:"manga_id"`
}