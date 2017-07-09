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
		return
	}

	c.JSON(http.StatusOK, r)
}

func torrentsSearch(c *gin.Context) {
	s := models.NewTorrentSearch()
	p := c.DefaultQuery("page", "1")
	name := c.Query("name")
	season := s.Int(c.Query("season"))
	episode := s.Int(c.Query("episode"))
	resolution := c.Query("resolution")
	source := c.Query("source")
	mtype := c.Query("type")
	bluray := c.Query("bluray") == "true"
	uncensored := c.Query("uncensored") == "true"
	verified := c.Query("verified") == "true"

	s.Name(name)
	s.Season(season)
	s.Episode(episode)
	s.Resolution(resolution)
	s.Source(source)
	s.Type(mtype)
	if bluray {
		s.Bluray(bluray)
	}
	if uncensored {
		s.Uncensored(uncensored)
	}
	if verified {
		s.Verified(verified)
	}

	//fmt.Printf("SEARCH: %#v\n", s)

	r, err := s.Results(s.Int(p))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run query"})
		return
	}

	c.JSON(http.StatusOK, r)
}

func torrentsShow(c *gin.Context) {
	id := c.Param("id")

	r, err := models.TorrentsFind(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run query: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}
