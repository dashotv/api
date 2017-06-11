package models

import "github.com/go-bongo/bongo"

type Response interface {
	New() interface{}
	Add(interface{})
	Pagination(*bongo.PaginationInfo)
}

type BaseResponse struct {
	Page *bongo.PaginationInfo `json:"header"`
}

func (b *BaseResponse) Pagination(p *bongo.PaginationInfo) {
	b.Page = p
}
