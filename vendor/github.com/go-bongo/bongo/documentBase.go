package bongo

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type DocumentBase struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Created  time.Time     `bson:"_created" json:"_created"`
	Modified time.Time     `bson:"_modified" json:"_modified"`

	// We want this to default to false without any work. So this will be the opposite of isNew. We want it to be new unless set to existing
	exists bool
}

// Satisfy the new tracker interface
func (d *DocumentBase) SetIsNew(isNew bool) {
	d.exists = !isNew
}

func (d *DocumentBase) IsNew() bool {
	return !d.exists
}

// Satisfy the document interface
func (d *DocumentBase) GetId() bson.ObjectId {
	return d.Id
}

func (d *DocumentBase) SetId(id bson.ObjectId) {
	d.Id = id
}

func (d *DocumentBase) SetCreated(t time.Time) {
	d.Created = t
}

func (d *DocumentBase) SetModified(t time.Time) {
	d.Modified = t
}
