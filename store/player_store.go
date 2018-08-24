package store

import "github.com/albshin/tekkenmatchdb-backend/model"

func (db *PGStore) CreatePlayer(req *model.Player) (*model.Player, error) {
	q := `
	INSERT INTO players
	(player_name, country)
	VALUES (:player_name, :country)
	`
	rows, err := db.NamedQuery(q, req)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&req.ID)
		rows.Close()
	}
	return req, nil
}
