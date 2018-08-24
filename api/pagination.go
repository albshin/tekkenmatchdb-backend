package api

import (
	"net/http"
	"strconv"

	"github.com/albshin/tekkenmatchdb-backend/model"
)

const limitPerPage = 20

func withPagination(r *http.Request) (*model.Pagination, error) {
	q := r.URL.Query()
	qPage := q.Get("page")
	qLimit := q.Get("limit")
	var page uint64 = 1
	var limit uint64 = limitPerPage
	var err error

	if qPage != "" {
		page, err = strconv.ParseUint(qPage, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	if qLimit != "" {
		limit, err = strconv.ParseUint(qLimit, 10, 64)
		if err != nil {
			return nil, err
		}
		// TODO: Probably send error instead
		if limit > limitPerPage {
			limit = limitPerPage
		}
	}

	return &model.Pagination{
		Page:  page,
		Limit: limit,
	}, nil
}
