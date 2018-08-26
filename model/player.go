package model

import (
	"time"
)

// Player is an object representing the database table.
type Player struct {
	ID         int        `json:"id" db:"id"`
	PlayerName string     `json:"player_name" db:"player_name"`
	Country    string     `json:"country,omitempty" db:"country"`
	CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
