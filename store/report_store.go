package store

import "github.com/albshin/tekkenmatchdb-backend/model"

func (db *PGStore) CreateMatchReport(mr *model.MatchReport) (*model.MatchReport, error) {
	q := `
	INSERT INTO match_reports
	(match_id, data)
	VALUES (:match_id, :data)
	`
	rows, err := db.NamedQuery(q, mr)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&mr.ID)
		rows.Close()
	}
	return mr, nil
}
