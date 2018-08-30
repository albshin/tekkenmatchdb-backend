package store

import "github.com/albshin/tekkenmatchdb-backend/model"

func (db *PGStore) GetRanks() ([]*model.Rank, error) {
	ranks := make([]*model.Rank, 0)
	if err := db.Select(&ranks, "SELECT rank_name FROM t7_ranks"); err != nil {
		return nil, err
	}
	return ranks, nil
}

func (db *PGStore) GetCharacters() ([]*model.Character, error) {
	characters := make([]*model.Character, 0)
	if err := db.Select(&characters, "SELECT character_name FROM t7_characters"); err != nil {
		return nil, err
	}
	return characters, nil
}
