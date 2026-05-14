package internal

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"mangahub/internal/auth"
	"mangahub/internal/library"
	"mangahub/internal/manga"
	"mangahub/internal/middleware"
	"mangahub/internal/websocket"

	"mangahub/internal/progress"

	"time"

	"github.com/gin-contrib/cors"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")

	// Auth
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// Manga
	mangaRepo := manga.NewRepository(db)
	mangaService := manga.NewService(mangaRepo)
	mangaHandler := manga.NewHandler(mangaService)

	mangaGroup := api.Group("/manga")
	{
		mangaGroup.GET("/", mangaHandler.GetAll)
		mangaGroup.GET("/search", mangaHandler.Search)
		mangaGroup.GET("/:id", mangaHandler.GetByID)
		mangaGroup.POST("/", mangaHandler.CreateManga)
	}

	// Proteced route
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(db))

	// Library
	libRepo := library.NewRepository(db)
	progressRepo := progress.NewRepository(db)

	mangaRepo = manga.NewRepository(db)

	libService := library.NewService(
		libRepo,
		progressRepo,
		mangaRepo,
	)

	libHandler := library.NewHandler(libService)

	libraryGroup := protected.Group("/library")
	{
		libraryGroup.POST("/:manga_id", libHandler.Add)
		libraryGroup.GET("/", libHandler.Get)
		libraryGroup.DELETE("/:manga_id", libHandler.Remove)
		libraryGroup.PUT("/:manga_id/status", libHandler.UpdateStatus)
		libraryGroup.POST("/progress", libHandler.Create)
		libraryGroup.PUT("/progress", libHandler.Update)
		libraryGroup.GET("/progress/:manga_id", libHandler.GetProgress)
	}

	// Chat
	//chatHub := chat.NewHub()
	//go chatHub.Run()

	//chatHandler := chat.ChatHandler

	//chatGroup := protected.Group("/chat")
	//{
	//	chatGroup.GET("/ws/:room_id", chatHandler.HandleWebSocket)
	//}

	// Websocket
	wsHandler := websocket.NewHandler()

	wsGroup := protected.Group("/ws")
	{
		wsGroup.GET("/manga/:id", wsHandler.HandleMangaRoom)
	}

	return r
}
