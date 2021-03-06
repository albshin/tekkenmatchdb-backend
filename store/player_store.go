package store

import (
	"github.com/albshin/tekkenmatchdb-backend/model"
)

func (db *PGStore) GetPlayer(playerID int) (*model.Player, error) {
	var player model.Player
	if err := db.Get(&player, "SELECT * FROM players WHERE id=$1", playerID); err != nil {
		return nil, err
	}
	return &player, nil
}

func (db *PGStore) GetPlayers(pageParams *model.Pagination) ([]*model.Player, error) {
	players := make([]*model.Player, 0)
	if err := db.Select(&players, "SELECT * FROM players LIMIT $1 OFFSET $2", pageParams.Limit, pageParams.Offset()); err != nil {
		return nil, err
	}
	return players, nil
}

func (db *PGStore) GetPlayerNames() ([]*model.PlayerName, error) {
	names := make([]*model.PlayerName, 0)
	if err := db.Select(&names, "SELECT player_name FROM players"); err != nil {
		return nil, err
	}
	return names, nil
}

func (db *PGStore) CreatePlayer(player *model.Player) (*model.Player, error) {
	q := `
	INSERT INTO players
	(player_name, country)
	VALUES (:player_name, :country)
	`
	rows, err := db.NamedQuery(q, player)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&player.ID)
		rows.Close()
	}
	return player, nil
}
