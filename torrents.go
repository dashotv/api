package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/dashotv/models"
)

func torrentsList(c *gin.Context) {
	p, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		p = 1
	}

	q := models.NewTorrentQuery()
	r, err := models.TorrentSearch(p, q.M())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run query"})
	}

	c.JSON(http.StatusOK, r)
}
