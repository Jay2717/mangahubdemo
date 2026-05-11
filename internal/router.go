package internal

import (
    "github.com/gin-gonic/gin"
    "mangahub/internal/manga"
    "mangahub/internal/auth"
    "mangahub/internal/middleware"
    "mangahub/internal/health"
    "mangahub/internal/readinglist"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    r.POST("/register", auth.RegisterHandler)
    r.POST("/login", auth.LoginHandler)

    r.GET("/manga", middleware.AuthMiddleware(), manga.GetMangaHandler)
    r.POST("/manga", middleware.AuthMiddleware(), manga.CreateMangaHandler)

    r.GET("/health", health.HealthHandler)

    r.POST("/reading-list", readinglist.AddToReadingList)
    r.GET("/reading-list", readinglist.GetReadingList)

    return r
}