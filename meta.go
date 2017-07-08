package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var META = map[string][]string{
	"sources": []string{
		"anidex",
		"extratorrent",
		"eztv",
		"horrible",
		"kickass",
		"lime",
		"monova",
		"nyaa",
		"piratebay",
		"rarbg",
		"shana",
		"showrss",
		"yify",
	},
	"resolutions": []string{
		"2160",
		"1080",
		"720",
	},
	"types": []string{
		"tv",
		"anime",
		"movies",
	},
}

func metaIndex(c *gin.Context) {
	c.JSON(http.StatusOK, META)
}
