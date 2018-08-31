DROP TABLE IF EXISTS match_reports;
DROP TABLE IF EXISTS youtube_videos;
DROP TABLE IF EXISTS matches;
DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS t7_ranks;
DROP TABLE IF EXISTS t7_characters;

CREATE TABLE IF NOT EXISTS t7_characters (
    character_name text NOT NULL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS t7_ranks (
    rank_name text NOT NULL PRIMARY KEY
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
    p1_rank text REFERENCES t7_ranks(rank_name) ON UPDATE CASCADE, 
    p2_rank text REFERENCES t7_ranks(rank_name) ON UPDATE CASCADE,
    p1_character text NOT NULL REFERENCES t7_characters(character_name) ON UPDATE CASCADE,
    p2_character text NOT NULL REFERENCES t7_characters(character_name) ON UPDATE CASCADE,
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

CREATE TABLE IF NOT EXISTS match_reports (
    id serial NOT NULL PRIMARY KEY,
    match_id int NOT NULL,
    data jsonb,

    FOREIGN KEY (match_id) REFERENCES matches(id)
);

INSERT INTO t7_characters (character_name) VALUES
    ('Akuma'),
    ('Alisa'),
    ('Asuka'),
    ('Bob'),
    ('Bryan'),
    ('Claudio'),
    ('Devil Jin'),
    ('Dragunov'),
    ('Eddy'),
    ('Geese'),
    ('Gigas'),
    ('Heihachi'),
    ('Hwoarang'),
    ('Jack-7'),
    ('Jin'),
    ('Josie'),
    ('Katarina'),
    ('Kazumi'),
    ('Kazuya'),
    ('King'),
    ('Kuma'),
    ('Lars'),
    ('Law'),
    ('Lee'),
    ('Leo'),
    ('Lili'),
    ('Lucky Chloe'),
    ('Miguel'),
    ('Nina'),
    ('Noctis'),
    ('Panda'),
    ('Paul'),
    ('Master Raven'),
    ('Shaheen'),
    ('Steve'),
    ('Xiaoyu'),
    ('Yoshimitsu')
    ON CONFLICT DO NOTHING;

INSERT INTO t7_ranks (rank_name) VALUES
    ('Beginner'),
    ('1st kyu'),
    ('2nd kyu'),
    ('3rd kyu'),
    ('4th kyu'),
    ('5th kyu'),
    ('6th kyu'),
    ('7th kyu'),
    ('8th kyu'),
    ('9th kyu'),
    ('1st dan'),
    ('2nd dan'),
    ('3rd dan'),
    ('Initiate'),
    ('Mentor'),
    ('Expert'),
    ('Grandmaster'),
    ('Brawler'),
    ('Marauder'),
    ('Fighter'),
    ('Vanguard'),
    ('Warrior'),
    ('Vindicator'),
    ('Juggernaut'),
    ('Usurper'),
    ('Vanquisher'),
    ('Destroyer'),
    ('Savior'),
    ('Overlord'),
    ('Genbu'),
    ('Byakko'),
    ('Seiryu'),
    ('Suzaku'),
    ('Mighty Ruler'),
    ('Revered Ruler'),
    ('Divine Ruler'),
    ('Eternal Ruler'),
    ('Fujin'),
    ('Raijin'),
    ('Yaksa'),
    ('Ryujin'),
    ('Emperor'),
    ('Tekken King'),
    ('Tekken God'),
    ('True Tekken God'),
    ('Tekken God Prime')
    ON CONFLICT DO NOTHING;

INSERT INTO players (player_name, country) VALUES ('LowHigh', 'KR');
INSERT INTO players (player_name, country) VALUES ('Wecka', 'KR');
