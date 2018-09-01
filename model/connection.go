package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Pagination holds the parameters necessary for pagination.
type Pagination struct {
	Page  uint64
	Limit uint64
}

// MatchFilter holds the parameters for match filtering.
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

// Offset returns the value by which the offset value should be defined
// as based on the parameters inside a Pagination object.
func (p *Pagination) Offset() uint64 {
	return (p.Page - 1) * p.Limit
}

func (mf *MatchFilter) Validate() error {
	// Do any pre validation transforms
	mf.transform()
	return validation.ValidateStruct(mf,
		validation.Field(&mf.P1Rank, validation.Match(regexp.MustCompile(`[a-zA-Z,]+`))),
		validation.Field(&mf.P2Rank, validation.Match(regexp.MustCompile(`[a-zA-Z,]+`))),
		validation.Field(&mf.P1Character, validation.Match(regexp.MustCompile(`[a-zA-Z,]+`))),
		validation.Field(&mf.P2Character, validation.Match(regexp.MustCompile(`[a-zA-Z,]+`))),
		validation.Field(&mf.Winner, validation.In("p1", "p2", "draw")),
	)
}

func (mf *MatchFilter) transform() {
	if mf.P1Character == "Kuma" || mf.P1Character == "Panda" {
		mf.P1Character = "Kuma,Panda"
	}
	if mf.P2Character == "Kuma" || mf.P2Character == "Panda" {
		mf.P2Character = "Kuma,Panda"
	}
}
