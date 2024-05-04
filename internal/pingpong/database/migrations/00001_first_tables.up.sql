START TRANSACTION;

CREATE SCHEMA pingpong;

SET LOCAL search_path TO pingpong, "$user", public;

CREATE TABLE tournament (
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    is_draft   boolean NOT NULL,
    start_date date,
    end_date   date
);

CREATE TABLE player (
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name varchar(50)
);

CREATE TABLE tournaments_players (
    tournament_id  uuid REFERENCES tournament,
    participant_id uuid REFERENCES player
);

CREATE TABLE match (
    id uuid         PRIMARY KEY DEFAULT gen_random_uuid(),
    tournament_id   uuid REFERENCES tournament,
    due_date        date NOT NULL,
    opponent1_id    uuid REFERENCES player,
    opponent1_score integer,
    opponent2_id    uuid REFERENCES player,
    opponent2_score integer
);

COMMIT TRANSACTION;
