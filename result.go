package pg

import (
	"bytes"
	"strconv"

	"github.com/rustjason/pg/internal"
	"github.com/rustjason/pg/orm"
)

// A result summarizes an executed SQL command.
type result struct {
	model orm.Model

	affected int
	returned int
}

var _ orm.Result = (*result)(nil)

func (res *result) parse(b []byte) error {
	res.affected = -1

	ind := bytes.LastIndexByte(b, ' ')
	if ind == -1 {
		return nil
	}

	s := internal.BytesToString(b[ind+1 : len(b)-1])

	affected, err := strconv.Atoi(s)
	if err == nil {
		res.affected = affected
	}

	return nil
}

func (res *result) Model() orm.Model {
	return res.model
}

func (res *result) RowsAffected() int {
	return res.affected
}

func (res *result) RowsReturned() int {
	return res.returned
}
