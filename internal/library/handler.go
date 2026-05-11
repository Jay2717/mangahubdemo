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

// GET LIBRARY
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

// REMOVE FROM LIBRARY (FIX router error)
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

// UPDATE STATUS (FIX router error)
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
