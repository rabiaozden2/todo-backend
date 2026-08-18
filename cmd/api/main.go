package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/rabiaozden/todo-backend/internal/auth"
	"github.com/rabiaozden/todo-backend/internal/db"
	"github.com/rabiaozden/todo-backend/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Bilgi: .env dosyasi bulunamadi, sistem ortam degiskenleri kullanilacak.")
	}

	database := db.Connect()
	r := gin.Default()

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check and root paths
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Welcome to Todo API! 🚀"})
	})
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Todo API v1.0 is running! 🚀", "endpoints": []string{"/api/auth/register", "/api/auth/login", "/api/tasks"}})
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	authHandler := &handlers.AuthHandler{DB: database}
	api := r.Group("/api")
	api.POST("/auth/register", authHandler.Register)
	api.POST("/auth/login", authHandler.Login)

	taskHandler := &handlers.TaskHandler{DB: database}
	tasks := api.Group("/tasks")
	tasks.Use(auth.RequireAuth())
	tasks.GET("", taskHandler.List)
	tasks.POST("", taskHandler.Create)
	tasks.PATCH("/:id", taskHandler.Update)
	tasks.DELETE("/:id", taskHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("Backend API running on http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}