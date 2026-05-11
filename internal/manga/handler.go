package manga

import (
	"net/http"

	"mangahub/pkg/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// CREATE MANGA
func (h *Handler) CreateManga(c *gin.Context) {
	var m models.Manga

	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateManga(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "manga created"})
}

// GET ALL MANGA
func (h *Handler) GetAll(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}

// GET MANGA BY ID
func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func (h *Handler) Search(c *gin.Context) {

	query := c.Query("q")
	author := c.Query("author")
	status := c.Query("status")
	genre := c.Query("genre")

	mangas, err := h.service.SearchManga(
		query,
		author,
		status,
		genre,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mangas)
}
