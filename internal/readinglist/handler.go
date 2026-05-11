package readinglist

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mangahub/pkg/database"
	"mangahub/pkg/models"
)

func AddToReadingList(c *gin.Context) {
	var item models.ReadingList

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	_, err := database.DB.Exec(
		"INSERT INTO reading_list(username, manga_id) VALUES (?, ?)",
		item.Username,
		item.MangaID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "added to reading list",
	})
}

func GetReadingList(c *gin.Context) {
	username := c.Query("username")

	rows, err := database.DB.Query(
		"SELECT id, username, manga_id FROM reading_list WHERE username = ?",
		username,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	var list []models.ReadingList

	for rows.Next() {
		var item models.ReadingList

		rows.Scan(
			&item.ID,
			&item.Username,
			&item.MangaID,
		)

		list = append(list, item)
	}

	c.JSON(http.StatusOK, list)
}