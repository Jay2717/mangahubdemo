package manga

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"mangahub/pkg/models"
)

func GetMangaHandler(c *gin.Context) {
	port := os.Getenv("PORT")

	mangas, err := GetMangaList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"server_port": port,
		"data": mangas,
	})
}

func CreateMangaHandler(c *gin.Context) {
	var m models.Manga

	// bind JSON từ client
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
		return
	}

	// gọi service
	err := CreateNewManga(m)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "manga created",
	})
}