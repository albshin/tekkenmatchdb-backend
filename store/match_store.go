package store

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/albshin/tekkenmatchdb-backend/model"
	"github.com/jmoiron/sqlx"
)

func (db *PGStore) GetMatches(matchFilters *model.MatchFilter, pageParams *model.Pagination) ([]*model.GetMatch, error) {
	matches := make([]*model.GetMatch, 0)
	q := sq.Select(`
		matches.*,
		p1.player_name AS p1_name,
		p1.country AS p1_country,
		p2.player_name AS p2_name,
		p2.country AS p2_country
	`).From(`
		matches
	`).Join(`
		players p1 ON matches.p1_id = p1.id
	`).Join(`
		players p2 ON matches.p2_id = p2.id
	`)

	if matchFilters.P1Name != "" {
		q = q.Where(sq.Eq{"p1.player_name": matchFilters.P1Name})
	}
	if matchFilters.P2Name != "" {
		q = q.Where(sq.Eq{"p2.player_name": matchFilters.P2Name})
	}
	if matchFilters.P1Rank != "" {
		q = q.Where(sq.Eq{"matches.p1_rank": matchFilters.P1Rank})
	}
	if matchFilters.P2Rank != "" {
		q = q.Where(sq.Eq{"matches.p2_rank": matchFilters.P2Rank})
	}
	if matchFilters.P1Character != "" {
		q = q.Where(sq.Eq{"matches.p1_character": matchFilters.P1Character})
	}
	if matchFilters.P2Character != "" {
		q = q.Where(sq.Eq{"matches.p2_rank": matchFilters.P2Character})
	}
	if matchFilters.Winner != "" {
		q = q.Where(sq.Eq{"matches.winner": matchFilters.Winner})
	}

	if pageParams != nil {
		q = q.Limit(pageParams.Limit)
		q = q.Offset(pageParams.Offset())
	}

	sql, args, _ := q.PlaceholderFormat(sq.Dollar).ToSql()
	// TODO: Optimize to not be n+1
	if err := db.Select(&matches, sql, args...); err != nil {
		return nil, err
	}
	ytvq := `SELECT * FROM youtube_videos WHERE match_id=$1`
	for _, match := range matches {
		if err := db.Select(&match.YoutubeVideos, ytvq, match.ID); err != nil {
			return nil, err
		}
	}

	return matches, nil
}

func (db *PGStore) CreateMatches(matches []*model.Match) ([]*model.Match, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	for _, match := range matches {
		if err := createMatch(tx, match); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return matches, nil
}

func createMatch(tx *sqlx.Tx, req *model.Match) error {
	q := `
	INSERT INTO matches
	(match_date, event_name, p1_id, p2_id, p1_rank, p2_rank, p1_character, p2_character, winner)
	VALUES (:match_date, :event_name, :p1_id, :p2_id, :p1_rank, :p2_rank, :p1_character, :p2_character, :winner)
	RETURNING id
	`
	rows, err := tx.NamedQuery(q, req)
	if err != nil {
		return err
	}
	var id int
	if rows.Next() {
		rows.Scan(&id)
		rows.Close()
	}

	q = `
	INSERT INTO youtube_videos
	(match_id, player_side, video_id, video_timestamp)
	VALUES (:match_id, :player_side, :video_id, :video_timestamp)
	`
	for _, vid := range req.YoutubeVideos {
		vid.MatchID = id

		_, err = tx.NamedExec(q, vid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *PGStore) GetMatchesByPlayerID(playerID int, matchFilters *model.MatchFilter, pageParams *model.Pagination) ([]*model.GetMatch, error) {
	matches := make([]*model.GetMatch, 0)
	q := sq.Select(`
		matches.*,
		p1.player_name AS p1_name,
		p1.country AS p1_country,
		p2.player_name AS p2_name,
		p2.country AS p2_country
	`).From(`
		matches
	`).Join(`
		players p1 ON matches.p1_id = p1.id
	`).Join(`
		players p2 ON matches.p2_id = p2.id
	`).Where(sq.Or{
		sq.Eq{"matches.p1_id": playerID},
		sq.Eq{"matches.p2_id": playerID},
	})

	if matchFilters.P1Name != "" {
		q = q.Where(sq.Eq{"p1.player_name": matchFilters.P1Name})
	}
	if matchFilters.P2Name != "" {
		q = q.Where(sq.Eq{"p2.player_name": matchFilters.P2Name})
	}
	if matchFilters.P1Rank != "" {
		q = q.Where(sq.Eq{"matches.p1_rank": matchFilters.P1Rank})
	}
	if matchFilters.P2Rank != "" {
		q = q.Where(sq.Eq{"matches.p2_rank": matchFilters.P2Rank})
	}
	if matchFilters.P1Character != "" {
		q = q.Where(sq.Eq{"matches.p1_character": matchFilters.P1Character})
	}
	if matchFilters.P2Character != "" {
		q = q.Where(sq.Eq{"matches.p2_rank": matchFilters.P2Character})
	}
	if matchFilters.Winner != "" {
		q = q.Where(sq.Eq{"matches.winner": matchFilters.Winner})
	}

	if pageParams != nil {
		q = q.Limit(pageParams.Limit)
		q = q.Offset(pageParams.Offset())
	}

	sql, args, _ := q.PlaceholderFormat(sq.Dollar).ToSql()
	if err := db.Select(&matches, sql, args...); err != nil {
		return nil, err
	}
	return matches, nil
}
