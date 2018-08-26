package model

import (
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/gobuffalo/pop/nulls"
)

type GetMatch struct {
	Match
	P1Name    string `json:"p1_name" db:"p1_name"`
	P1Country string `json:"p1_country" db:"p1_country"`
	P2Name    string `json:"p2_name" db:"p2_name"`
	P2Country string `json:"p2_country" db:"p2_country"`
}

// Match is an object representing the database table.
type Match struct {
	ID            int             `json:"id" db:"id"`
	MatchDate     time.Time       `json:"match_date" db:"match_date"`
	EventName     string          `json:"event_name" db:"event_name"`
	P1ID          int             `json:"p1_id" db:"p1_id"`
	P2ID          int             `json:"p2_id" db:"p2_id"`
	P1Rank        nulls.String    `json:"p1_rank" db:"p1_rank"`
	P2Rank        nulls.String    `json:"p2_rank" db:"p2_rank"`
	P1Character   string          `json:"p1_character" db:"p1_character"`
	P2Character   string          `json:"p2_character" db:"p2_character"`
	Winner        string          `json:"winner" db:"winner"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" db:"updated_at"`
	YoutubeVideos []*YoutubeVideo `json:"youtube_videos"`
}

func (m *Match) Validate() error {
	return validation.ValidateStruct(m,
		//validation.Field(&m.MatchDate, validation.Required, validation.Date("YYYY-MM-DD")),
		validation.Field(&m.EventName, validation.Required, is.Alphanumeric),
		validation.Field(&m.P1ID, validation.Required, is.Digit),
		validation.Field(&m.P2ID, validation.Required, is.Digit),
		validation.Field(&m.P1Rank, is.Alpha),
		validation.Field(&m.P2Rank, is.Alpha),
		validation.Field(&m.P1Character, validation.Required, is.Alpha),
		validation.Field(&m.P2Character, validation.Required, is.Alpha),
		validation.Field(&m.Winner, validation.Required, validation.In("p1", "p2", "draw")),
		validation.Field(&m.YoutubeVideos),
	)
}
