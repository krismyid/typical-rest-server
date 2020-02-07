package dbkit

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type sort struct {
	column  string
	orderBy OrderBy
}

type OrderBy int

const (
	Asc OrderBy = iota
	Desc
)

func (o OrderBy) String() string {
	switch o {
	case Asc:
		return "ASC"
	case Desc:
		return "DESC"
	}
	return "ASC"
}

// Sort is find option to sort by column and order
func Sort(column string, orderBy OrderBy) FindOption {
	return &sort{
		column:  column,
		orderBy: orderBy,
	}
}

// CompileQuery to compile select query for sorting
func (s *sort) CompileQuery(base sq.SelectBuilder) (sq.SelectBuilder, error) {
	if s.column == "" {
		return base, errors.New("Sort column can't be empty")
	}
	base = base.OrderBy(fmt.Sprintf("%s %s", s.column, s.orderBy))
	return base, nil
}