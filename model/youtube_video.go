package model

import (
	"regexp"
	"time"

	"github.com/go-ozzo/ozzo-validation"

	"github.com/gobuffalo/pop/nulls"
)

// YoutubeVideo is an object representing the database table.
type YoutubeVideo struct {
	ID             int          `json:"id" db:"id"`
	MatchID        int          `json:"match_id" db:"match_id"`
	VideoID        string       `json:"video_id" db:"video_id"`
	VideoTimestamp string       `json:"video_timestamp" db:"video_timestamp"`
	PlayerSide     nulls.String `json:"player_side" db:"player_side"`
	CreatedAt      time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at" db:"updated_at"`
}

func (y *YoutubeVideo) Validate() error {
	// YoutubeID regex is subject to change
	return validation.ValidateStruct(y,
		validation.Field(&y.VideoID, validation.Required, validation.Match(regexp.MustCompile(`[a-zA-Z0-9_-]{11}`))),
		validation.Field(&y.VideoTimestamp, validation.Required, validation.Match(regexp.MustCompile(`^[0-9]{2}h[0-9]{2}m[0-9]{2}s$`))),
		validation.Field(&y.PlayerSide, validation.In("p1", "p2")),
	)
}
