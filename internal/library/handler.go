package library

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Add(c *gin.Context) {
	userID := c.GetInt("user_id")

	mangaID, err := strconv.Atoi(c.Param("manga_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid manga_id",
		})
		return
	}

	var body struct {
		Status         string `json:"status"`
		CurrentChapter int    `json:"current_chapter"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.AddToLibrary(
		userID,
		mangaID,
		body.Status,
		body.CurrentChapter,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "added",
	})
}

// Get library
func (h *Handler) Get(c *gin.Context) {
	userID := c.GetInt("user_id")

	data, err := h.service.GetLibrary(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Remove from lib
func (h *Handler) Remove(c *gin.Context) {
	userID := c.GetInt("user_id")

	mangaID, err := strconv.Atoi(c.Param("manga_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid manga_id"})
		return
	}

	err = h.service.RemoveFromLibrary(userID, mangaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "removed"})
}

// Update status
func (h *Handler) UpdateStatus(c *gin.Context) {
	userID := c.GetInt("user_id")

	mangaID, err := strconv.Atoi(c.Param("manga_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid manga_id"})
		return
	}

	var body struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdateStatus(userID, mangaID, body.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// Create progress
func (h *Handler) Create(c *gin.Context) {
	userID := c.GetInt("user_id")

	var body struct {
		MangaID int `json:"manga_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Create(userID, body.MangaID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "created"})
}

// Update progress
func (h *Handler) Update(c *gin.Context) {
	userID := c.GetInt("user_id")

	var body struct {
		MangaID int `json:"manga_id"`
		Chapter int `json:"chapter"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Update(userID, body.MangaID, body.Chapter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "updated"})
}

// Get progress
func (h *Handler) GetProgress(c *gin.Context) {
	userID := c.GetInt("user_id")

	mangaID, _ := strconv.Atoi(c.Param("manga_id"))

	data, err := h.service.GetProgress(userID, mangaID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}
