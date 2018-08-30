package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Pagination struct {
	Page  uint64
	Limit uint64
}

type MatchFilter struct {
	P1Name      string `json:"p1_name"`
	P2Name      string `json:"p2_name"`
	P1Rank      string `json:"p1_rank"`
	P2Rank      string `json:"p2_rank"`
	P1Character string `json:"p1_character"`
	P2Character string `json:"p2_character"`
	Winner      string `json:"winner"`
}

type Rank struct {
	Name string `json:"rank_name" db:"rank_name"`
}

type Character struct {
	Name string `json:"character_name" db:"character_name"`
}

func (p *Pagination) Offset() uint64 {
	return (p.Page - 1) * p.Limit
}

func (m *MatchFilter) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.P1Rank, validation.Match(regexp.MustCompile(`[a-zA-Z,]+`))),
		validation.Field(&m.P2Rank, validation.Match(regexp.MustCompile(`[a-zA-Z,]+`))),
		validation.Field(&m.P1Character, is.Alpha),
		validation.Field(&m.P2Character, is.Alpha),
		validation.Field(&m.Winner, validation.In("p1", "p2", "draw")),
	)
}
