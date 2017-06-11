package models

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type Collector struct {
	For        string
	Connection *bongo.Connection
	Collection *bongo.Collection
}

func NewCollector(col string, con *bongo.Connection) *Collector {
	return &Collector{
		For:        col,
		Connection: con,
		Collection: con.Collection(col),
	}
}

func (c *Collector) Find(id string, doc interface{}) error {
	err := c.Collection.FindById(bson.ObjectIdHex(id), doc)
	if err != nil {
		return err
	}

	return nil
}

func (c *Collector) FindOne(query interface{}, doc interface{}) error {
	return c.Collection.FindOne(query, doc)
}

func (c *Collector) Save(doc bongo.Document) error {
	return c.Collection.Save(doc)
}

func (c *Collector) Where(query interface{}) *Search {
	s := &Search{
		query:   query,
		results: c.Collection.Find(query),
	}

	s.page = 1
	s.Limit(PER_PAGE)

	return s
}
