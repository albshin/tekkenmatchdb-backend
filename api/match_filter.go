package api

import (
	"net/http"

	"github.com/albshin/tekkenmatchdb-backend/model"
)

func getMatchFilters(r *http.Request) (*model.MatchFilter, error) {
	q := r.URL.Query()
	mf := &model.MatchFilter{
		P1Name:      q.Get("p1_name"),
		P2Name:      q.Get("p2_name"),
		P1Rank:      q.Get("p1_rank"),
		P2Rank:      q.Get("p2_rank"),
		P1Character: q.Get("p1_char"),
		P2Character: q.Get("p2_char"),
		Winner:      q.Get("winner"),
	}
	return mf, mf.Validate()
}
