package model

import (
	"time"

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
