package store

import (
	"testing"
	"time"

	"github.com/albshin/tekkenmatchdb-backend/model"
	"github.com/gobuffalo/pop/nulls"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TestSuite struct {
	tests map[string]func(*testing.T)
	store Storage
	db    *sqlx.DB
}

func TestMatchStore(t *testing.T) {
	db, err := sqlx.Open("postgres", "user=postgres password=password dbname=matches sslmode=disable")
	if err != nil {
		t.Fatalf("Could not establish a connection with the database: %s\n", err.Error())
	}
	ts := TestSuite{store: &PGStore{db}, db: db}
	tests := map[string]func(*testing.T){
		"get matches 1 player filter": ts.testGetMatches,
		"get matches by player ID":    ts.testGetMatchesByPlayerID,
		"create matches":              ts.testCreateMatches,
	}
	ts.tests = tests

	for name, test := range ts.tests {
		t.Run(name, test)
	}
}

func (ts *TestSuite) testGetMatches(t *testing.T) {
}

func (ts *TestSuite) testGetMatchesByPlayerID(t *testing.T) {
}

func (ts *TestSuite) testCreateMatches(t *testing.T) {
	models := []*model.Match{
		&model.Match{
			MatchDate:   time.Now(),
			EventName:   "Evo",
			P1ID:        1,
			P2ID:        2,
			P1Rank:      nulls.NewString("Vindicator"),
			P2Rank:      nulls.NewString("Vanquisher"),
			P1Character: "Shaheen",
			P2Character: "Alisa",
			Winner:      "p1",
		},
	}
	ts.store.CreateMatches(models)
}
