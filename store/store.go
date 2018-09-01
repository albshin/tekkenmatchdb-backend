package store

import (
	"github.com/albshin/tekkenmatchdb-backend/model"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	GetMatches(mf *model.MatchFilter, pageParams *model.Pagination) ([]*model.GetMatch, error)
	GetMatchesByPlayerID(playerID int, mf *model.MatchFilter, pageParams *model.Pagination) ([]*model.GetMatch, error)
	CreateMatches(matches []*model.Match) ([]*model.Match, error)

	GetPlayer(playerID int) (*model.Player, error)
	GetPlayers(pageParams *model.Pagination) ([]*model.Player, error)
	GetPlayerNames() ([]*model.PlayerName, error)
	CreatePlayer(player *model.Player) (*model.Player, error)

	CreateMatchReport(mr *model.MatchReport) (*model.MatchReport, error)

	GetRanks() ([]*model.Rank, error)
	GetCharacters() ([]*model.Character, error)
}

type PGStore struct {
	*sqlx.DB
}

func Open(dsn string) (*PGStore, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &PGStore{db}, nil
}

func (db *PGStore) Close() error {
	return db.Close()
}
