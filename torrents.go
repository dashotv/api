package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/dashotv/models"
)

func torrentsList(c *gin.Context) {
	p := c.DefaultQuery("page", "1")

	page, err := strconv.Atoi(p)
	if err != nil {
		page = 1
	}

	//q := models.NewTorrentQuery()
	r, err := models.TorrentIndex(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run query"})
	}

	c.JSON(http.StatusOK, r)
}
