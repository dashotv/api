package models

import "gopkg.in/mgo.v2/bson"

type BaseQuery struct {
	Query bson.M
}

func (q *BaseQuery) M() bson.M {
	// TODO: handle non-exact queries
	return q.Query
}

func (q *BaseQuery) SetString(name, value string, ok bool) {
	if ok {
		q.Query[name] = value
	}
}

func (q *BaseQuery) SetInt(name string, value int, ok bool) {
	if ok {
		q.Query[name] = value
	}
}

func (q *BaseQuery) SetBool(name string, value bool, ok bool) {
	if ok {
		q.Query[name] = value
	}
}
