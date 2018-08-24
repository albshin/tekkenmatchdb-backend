DROP TABLE IF EXISTS youtube_videos;
DROP TABLE IF EXISTS matches;
DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS t7_rank;
DROP TABLE IF EXISTS t7_character;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS t7_character (
    character_name text NOT NULL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS t7_rank (
    rank_name text NOT NULL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL PRIMARY KEY,
    email text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS players (
    id serial NOT NULL PRIMARY KEY,
    player_name text NOT NULL,
    country text REFERENCES countries(id),
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS matches (
    id serial NOT NULL PRIMARY KEY,
    match_date date NOT NULL,
    event_name text NOT NULL,
    p1_id int NOT NULL REFERENCES players(id),
    p2_id int NOT NULL REFERENCES players(id),
    p1_rank text REFERENCES t7_rank(rank_name) ON UPDATE CASCADE, 
    p2_rank text REFERENCES t7_rank(rank_name) ON UPDATE CASCADE,
    p1_character text NOT NULL REFERENCES t7_character(character_name) ON UPDATE CASCADE,
    p2_character text NOT NULL REFERENCES t7_character(character_name) ON UPDATE CASCADE,
    winner text NOT NULL CHECK (winner IN ('p1', 'p2', 'draw')), 
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS youtube_videos (
    id serial NOT NULL PRIMARY KEY,
    match_id  int NOT NULL,
    player_side text CHECK (player_side IN ('p1', 'p2')),
    video_id text NOT NULL,
    video_timestamp text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),

    FOREIGN KEY (match_id) REFERENCES matches(id)
);

INSERT INTO t7_character (character_name) VALUES
    ('akuma'),
    ('alisa'),
    ('anna'),
    ('asuka'),
    ('bob'),
    ('bryan'),
    ('chloe'),
    ('claudio'),
    ('devil_jin'),
    ('dragunov'),
    ('eddy'),
    ('geese'),
    ('gigas'),
    ('heihachi'),
    ('hwoarang'),
    ('jack-7'),
    ('jin'),
    ('josie'),
    ('katarina'),
    ('kazumi'),
    ('kazuya'),
    ('king'),
    ('kuma'),
    ('lars'),
    ('law'),
    ('lee'),
    ('lei'),
    ('leo'),
    ('lili'),
    ('miguel'),
    ('negan'),
    ('nina'),
    ('panda'),
    ('paul'),
    ('master_raven'),
    ('shaheen'),
    ('steve'),
    ('xiaoyu'),
    ('yoshimitsu')
    ON CONFLICT DO NOTHING;

INSERT INTO t7_rank (rank_name) VALUES
    ('beginner'),
    ('1st_kyu'),
    ('2nd_kyu'),
    ('3rd_kyu'),
    ('4th_kyu'),
    ('5th_kyu'),
    ('6th_kyu'),
    ('7th_kyu'),
    ('8th_kyu'),
    ('9th_kyu'),
    ('1st_dan'),
    ('2nd_dan'),
    ('3rd_dan'),
    ('initiate'),
    ('mentor'),
    ('expert'),
    ('grandmaster'),
    ('brawler'),
    ('marauder'),
    ('fighter'),
    ('vanguard'),
    ('warrior'),
    ('vindicator'),
    ('juggernaut'),
    ('usurper'),
    ('vanquisher'),
    ('destroyer'),
    ('savior'),
    ('overlord'),
    ('genbu'),
    ('byakko'),
    ('seiryu'),
    ('suzaku'),
    ('mighty_ruler'),
    ('revered_ruler'),
    ('divine_ruler'),
    ('eternal_ruler'),
    ('fujin'),
    ('raijin'),
    ('yaksa'),
    ('ryujin'),
    ('emperor'),
    ('tekken_king'),
    ('tekken_god'),
    ('true_tekken_god'),
    ('tekken_god_prime')
    ON CONFLICT DO NOTHING;

INSERT INTO players (player_name, country) VALUES ('LowHigh', 'KR');
INSERT INTO players (player_name, country) VALUES ('Qudans', 'KR');
INSERT INTO players (player_name, country) VALUES ('Knee', 'KR');

INSERT INTO matches (match_date, event_name, p1_id, p2_id, p1_rank, p2_rank, p1_character, p2_character, winner)
VALUES (NOW(), 'Evo', 1, 2, 'vanquisher', 'vindicator', 'shaheen', 'alisa', 'p1');

INSERT INTO matches (match_date, event_name, p1_id, p2_id, p1_rank, p2_rank, p1_character, p2_character, winner)
VALUES (NOW(), 'Evo', 1, 3, 'vanquisher', 'vindicator', 'shaheen', 'devil_jin', 'p1');

INSERT INTO matches (match_date, event_name, p1_id, p2_id, p1_rank, p2_rank, p1_character, p2_character, winner)
VALUES (NOW(), 'Evo', 2, 3, 'vanquisher', 'vindicator', 'bob', 'steve', 'p2');