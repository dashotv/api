package models

import "github.com/go-bongo/bongo"

type Search struct {
	query   interface{}
	results *bongo.ResultSet
	page    int
	limit   int
	sort    []string
}

//func (s *Search) Page(num int) *Search {
//	s.page = num
//	return s
//}

func (s *Search) Limit(num int) *Search {
	s.limit = num
	s.results.Query.Limit(num)
	return s
}

func (s *Search) Sort(fields ...string) *Search {
	s.sort = fields
	s.results.Query.Sort(fields...)
	return s
}

func (s *Search) Free() {
	s.results.Free()
}

func (s *Search) Page(num int, r Response) error {
	defer s.Free()

	if s.results.Error != nil {
		return s.results.Error
	}

	for i := 0; i < s.limit; i++ {
		m := r.New()

		if !s.results.Next(m) {
			break
		}

		r.Add(m)
	}

	pi, err := s.results.Paginate(s.limit, num)
	if err != nil {
		return err
	}

	r.Pagination(pi)

	return nil
}

func (s *Search) First(num int, r Response) error {
	defer s.Free()

	s.Limit(num)

	if s.results.Error != nil {
		return s.results.Error
	}

	pi, err := s.results.Paginate(num, 1)
	if err != nil {
		return err
	}
	r.Pagination(pi)

	for i := 0; i < num; i++ {
		m := r.New()

		if !s.results.Next(m) {
			break
		}

		r.Add(m)
	}

	return nil
}
