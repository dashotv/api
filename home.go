package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func init() {
	router.GET("/", homeIndex)
}

func homeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
}
