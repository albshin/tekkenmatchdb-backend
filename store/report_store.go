package store

import "github.com/albshin/tekkenmatchdb-backend/model"

func (db *PGStore) ReportMatch(req *model.MatchReport) (*model.MatchReport, error) {
	q := `
	INSERT INTO match_reports
	(match_id, data)
	VALUES (:match_id, :data)
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
