package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/dashotv/models"
)

var (
	router *gin.Engine      = gin.Default()
	api    *gin.RouterGroup = router.Group("/api")
)

func main() {
	models.InitDB(os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_HOST"))
	router.Run()
}
