package models

type UserMangaLibrary struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	MangaID int    `json:"manga_id"`
	Status  string `json:"status"`
}
