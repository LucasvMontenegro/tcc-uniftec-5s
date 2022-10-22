create table if not exists scores (
	id bigserial NOT NULL,
	team_id BIGINT NOT NULL,
	five_s_id BIGINT NOT NULL,
	score float NOT NULL,

  	CONSTRAINT score_pkey PRIMARY KEY (id),
    CONSTRAINT fk_team FOREIGN KEY (team_id) REFERENCES teams(id),
    CONSTRAINT fk_five_s FOREIGN KEY (five_s_id) REFERENCES five_s(id),
	CONSTRAINT unique_team_five_s UNIQUE (team_id, five_s_id)
);
