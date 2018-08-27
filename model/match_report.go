package model

import "github.com/jmoiron/sqlx/types"

type MatchReport struct {
	ID      int            `json:"id" db:"id"`
	MatchID int            `json:"match_id" db:"match_id"`
	Data    types.JSONText `json:"data" db:"data"`
}
