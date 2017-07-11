package main

import (
	"fmt"
	"os"
	"time"

	//"github.com/MAD-GooZe/jaderender"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/dashotv/models"
)

var tokenSecret string

func main() {
	tokenSecret = os.Getenv("TOKEN_SECRET")
	host := os.Getenv("DATABASE_HOST")
	name := os.Getenv("DATABASE_NAME")
	mode := os.Getenv("GIN_MODE")
	fmt.Printf("starting dashotv/api server (%s)...\n", mode)

	router := gin.Default()
	//router.HTMLRender = jaderender.Default()
	corsConfig := cors.Config{
		AllowOrigins: []string{"https://dasho.tv", "http://localhost:8000", "http://localhost:4200"},
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

	meta := router.Group("/meta")
	meta.Use(corsMiddleware)
	meta.Use(Auth(tokenSecret))
	{
		meta.GET("/", metaIndex)
	}

	torrents := router.Group("/torrents")
	{
		torrents.GET("/", torrentsSearch)
		torrents.GET("/:id", torrentsShow)
	}

	session := router.Group("/auth")
	{
		session.POST("/", sessionIndex)
		session.POST("/refresh", sessionRefresh)
		session.POST("/sign_in", sessionCreate)
		session.DELETE("/sign_out", sessionDestroy)
	}

	t1 := make(chan bool, 1)
	go func() {
		fmt.Printf("intializing db connection: %s/%s\n", host, name)
		models.InitDB(name, host)
		t1 <- true
	}()

	select {
	case _ = <-t1:
		fmt.Println("database initialized")
	case <-time.After(time.Second * 30):
		fmt.Println("timed out waiting for database")
		os.Exit(1)
	}

	fmt.Println("running router")
	router.Run()
}
