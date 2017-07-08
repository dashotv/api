package main

import (
	"net/http"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"crypto/rand"
	"fmt"
	"github.com/dashotv/models"
	"io"
)

const expireSeconds = 1800

type UserLogin struct {
	Email    string
	Password string
}

type UserRefresh struct {
	Email   string
	Refresh string
}

func sessionIndex(c *gin.Context) {
	fmt.Printf("%#v\n", c)
}

func sessionCreate(c *gin.Context) {
	var login UserLogin

	if c.Bind(&login) != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not bind"})
		return
	}

	//fmt.Printf("login: %#v\n", login)

	user, err := models.UserFind(login.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if !user.CheckPassword(login.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email or password incorrect", "email": login.Email, "password": login.Password, "hash": user.PasswordHash})
		return
	}

	token, err := sessionToken(user.Email, c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not generate token"})
		return
	}

	refresh := refreshToken()
	user.Refresh = refresh

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not save refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh": refresh})
}

func sessionRefresh(c *gin.Context) {
	var params UserRefresh

	if c.Bind(&params) != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not bind"})
		return
	}

	user, err := models.UserRefresh(params.Email, params.Refresh)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	token, err := sessionToken(user.Email, c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh": refreshToken()})
}

func sessionDestroy(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "user not found",
	})
}

// refreshToken generates a random UUID according to RFC 4122
func refreshToken() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func sessionToken(id string, c *gin.Context) (string, error) {
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"Id":  id,
		"exp": time.Now().Add(time.Second * expireSeconds).Unix(),
	}
	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(tokenSecret))
}
