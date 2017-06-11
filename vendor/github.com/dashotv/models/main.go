package models

import (
	"fmt"

	"github.com/go-bongo/bongo"
)

const COLLECTION_USERS = "Users"
const COLLECTION_MEDIA = "Media"
const COLLECTION_TORRENTS = "Torrents"
const PER_PAGE = 25

var (
	// DB holds the db connection
	DB *Connector
)

type Connector struct {
	connection *bongo.Connection
	Users      *Collector
	Media      *Collector
	Torrents   *Collector
}

type Document struct {
	bongo.DocumentBase `bson:",inline"`
}

func InitDB(name, host string) {
	config := &bongo.Config{
		ConnectionString: host,
		Database:         name,
	}

	connection, err := bongo.Connect(config)
	if err != nil {
		panic(fmt.Sprintf("bongo error: (%s/%s) %s", host, name, err))
	}

	DB = &Connector{
		connection: connection,
		Users:      NewCollector(COLLECTION_USERS, connection),
		Media:      NewCollector(COLLECTION_MEDIA, connection),
		Torrents:   NewCollector(COLLECTION_TORRENTS, connection),
	}
}

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}
