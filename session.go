package main

import (
	"net/http"
	"os"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/dashotv/models"
)

var mysupersecretpassword string

func init() {
	mysupersecretpassword = os.Getenv("TOKEN_SECRET")
	session := router.Group("/session")
	{
		session.POST("/", sessionCreate)
		session.DELETE("/", sessionDestroy)
	}
}

type UserLogin struct {
	Email    string
	Password string
}

func sessionCreate(c *gin.Context) {
	var login UserLogin

	if c.Bind(&login) != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not bind"})
	}

	user, err := models.UserFind(login.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	}

	if !user.CheckPassword(login.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email or password incorrect"})
	}

	token, err := sessionToken(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not generate token"})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func sessionDestroy(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "user not found",
	})
}

func sessionToken(c *gin.Context) (string, error) {
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"Id":  "Christopher",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(mysupersecretpassword))
}
