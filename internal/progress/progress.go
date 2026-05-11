package progress

type UserProgress struct {
	ID             int `json:"id"`
	UserID         int `json:"user_id"`
	MangaID        int `json:"manga_id"`
	CurrentChapter int `json:"current_chapter"`
	LastPage       int `json:"last_page"`
}
