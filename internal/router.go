package internal

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"mangahub/internal/auth"
	"mangahub/internal/chat"
	"mangahub/internal/library"
	"mangahub/internal/manga"
	"mangahub/internal/websocket"
	"mangahub/pkg/middleware"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// AUTH
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// MANGA
	mangaRepo := manga.NewRepository(db)
	mangaService := manga.NewService(mangaRepo)
	mangaHandler := manga.NewHandler(mangaService)

	mangaGroup := r.Group("/manga")
	{
		mangaGroup.GET("/", mangaHandler.GetAll)
		mangaGroup.GET("/search", mangaHandler.Search)
		mangaGroup.GET("/:id", mangaHandler.GetByID)
		mangaGroup.POST("/", mangaHandler.CreateManga)
	}

	// PROTECTED ROUTES
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// LIBRARY
	libRepo := library.NewRepository(db)
	libService := library.NewService(libRepo)
	libHandler := library.NewHandler(libService)

	libraryGroup := protected.Group("/library")
	{
		libraryGroup.POST("/:manga_id", libHandler.Add)
		libraryGroup.GET("/", libHandler.Get)
		libraryGroup.DELETE("/:manga_id", libHandler.Remove)
		libraryGroup.PUT("/:manga_id/status", libHandler.UpdateStatus)
	}

	// CHAT
	chatHub := chat.NewHub()
	go chatHub.Run()

	chatHandler := chat.NewHandler(chatHub)

	chatGroup := protected.Group("/chat")
	{
		chatGroup.GET("/ws/:room_id", chatHandler.HandleWebSocket)
	}

	// WEBSOCKET
	wsHandler := websocket.NewHandler()

	wsGroup := protected.Group("/ws")
	{
		wsGroup.GET("/manga/:id", wsHandler.HandleMangaRoom)
	}

	return r
}
