package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/dashotv/models"
	"gopkg.in/mgo.v2/bson"
)

var tokenSecret string

func main() {
	tokenSecret = os.Getenv("TOKEN_SECRET")
	mode := os.Getenv("GIN_MODE")
	fmt.Printf("starting dashotv/api server (%s)...\n", mode)

	corsMiddleware := configureCors()

	router := gin.Default()

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

	initDB()
	fmt.Println("running router")
	router.Run()
}

func initDB() {
	//host := os.Getenv("DATABASE_HOST")
	//name := os.Getenv("DATABASE_NAME")

	mode := os.Getenv("GIN_MODE")
	env := "development"
	if mode == "release" {
		env = "production"
	}
	config := &models.Config{
		Host: "127.0.0.1",
		Torrents: &models.ConfigEntry{
			Database:   fmt.Sprintf("torch_%s", env),
			Collection: "torrents",
		},
		Media: &models.ConfigEntry{
			Database:   fmt.Sprintf("seer_%s", env),
			Collection: "media",
		},
		Users: &models.ConfigEntry{
			Database:   "dashotv",
			Collection: "users",
		},
	}

	t1 := make(chan bool, 1)
	go func() {
		fmt.Printf("intializing db connection: %s\n", config.Host)
		models.InitDB(config)
		t1 <- true
	}()

	select {
	case _ = <-t1:
		fmt.Println("database initialized")
	case <-time.After(time.Second * 30):
		fmt.Println("timed out waiting for database")
		os.Exit(1)
	}

	user := &models.User{}
	err := models.DB.Users.FindOne(bson.M{"email": "shawn@dasho.net"}, user)
	if err != nil {
		fmt.Printf("error getting user: %s\n", err)
		user.Email = "shawn@dasho.net"
		user.Password = "blarg"
		err = user.Save()
		if err != nil {
			fmt.Printf("error while saving user: %s\n", err)
		}
	}
}

func configureCors() gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowOrigins: []string{"http://dasho.tv", "http://www.dasho.tv", "http://localhost:8000", "http://localhost:4200"},
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

	return cors.New(corsConfig)
}
