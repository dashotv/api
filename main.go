package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/dashotv/models"
)

var tokenSecret string

func main() {
	models.InitDB(os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_HOST"))
	tokenSecret = os.Getenv("TOKEN_SECRET")

	router := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins: []string{"https://dasho.tv", "http://localhost:8000"},
		AllowMethods: []string{"GET, POST, OPTIONS, PUT, DELETE, HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"if-modified-since",
		},
		ExposeHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"if-modified-since",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	corsMiddleware := cors.New(corsConfig)

	router.Use(corsMiddleware)
	router.GET("/", homeIndex)

	api := router.Group("/api")
	api.Use(corsMiddleware)
	api.Use(Auth(tokenSecret))

	torrents := api.Group("/torrents")
	{
		torrents.GET("/", torrentsList)
	}

	session := router.Group("/auth")
	{
		session.POST("/sign_in", sessionCreate)
		session.DELETE("/sign_out", sessionDestroy)
	}

	router.Run()
}
