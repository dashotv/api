package models

import (
	"time"
)

type MediumResponse struct {
	BaseResponse `bson:",inline" json:",inline"`
	List         []*Medium
}

func (r *MediumResponse) New() interface{} {
	return &Medium{}
}

func (r *MediumResponse) Add(m interface{}) {
	r.List = append(r.List, m.(*Medium))
}

type Medium struct {
	Document `bson:",inline"`

	Type     *string `bson:"_type"`
	Source   *string
	SourceID *string `bson:"source_id"`

	Slug         *string
	Text         *string
	Display      *string
	Directory    *string
	Search       *string
	SearchParams struct {
		Group      *string
		Author     *string
		Resolution *string
		Type       *string
		Source     *string
		Verified   bool
		Uncensored bool
		Bluray     bool
	} `bson:"search_params"`

	Active      bool
	Downloaded  bool
	Completed   bool
	Skipped     bool
	Watched     bool
	Broken      bool
	Title       *string
	Description *string
	ReleaseDate time.Time `bson:"release_date"`

	Created time.Time `bson:"created_at"`
	Updated time.Time `bson:"updated_at"`
}

type MediumQuery struct {
	BaseQuery
}

func MediumFind(id string) (*Medium, error) {
	media := &Medium{}

	err := DB.Media.Find(id, media)
	if err != nil {
		return nil, err
	}

	return media, nil
}

func MediumMovies() ([]*Medium, error) {
	r := &MediumResponse{}

	//err := DB.Media.Where(bson.M{"_type": "Movie"}).Sort("-created_at").First(25, r)
	//if err != nil {
	//	return nil, err
	//}

	return r.List, nil
}

func (m *Medium) Save() error {
	return DB.Media.Save(m)
}

func (m *Medium) Find(id string) (*Medium, error) {
	err := DB.Media.Find(id, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
