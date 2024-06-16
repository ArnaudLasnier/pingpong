START TRANSACTION;

CREATE TYPE tournament_status AS ENUM (
    'draft',
    'started',
    'ended'
);

CREATE TABLE tournament (
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title      text UNIQUE NOT NULL CHECK (length(title) > 0),
    status     tournament_status NOT NULL,
    started_at timestamptz,
    ended_at   timestamptz
);

CREATE TABLE player (
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name text NOT NULL CHECK (length(first_name) > 0),
    last_name  text NOT NULL CHECK (length(last_name) > 0),
    email      text UNIQUE NOT NULL
);

CREATE TABLE tournament_participation (
    tournament_id  uuid REFERENCES tournament,
    participant_id uuid REFERENCES player,
    PRIMARY KEY (tournament_id, participant_id)
);

CREATE TABLE match (
    id                uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    tournament_id     uuid NOT NULL REFERENCES tournament,
    parent_match_1_id uuid REFERENCES match,
    parent_match_2_id uuid REFERENCES match,
    due_at            timestamptz NOT NULL,
    opponent_1_id     uuid REFERENCES player,
    opponent_1_score  integer,
    opponent_2_id     uuid REFERENCES player,
    opponent_2_score  integer
);

COMMIT TRANSACTION;
