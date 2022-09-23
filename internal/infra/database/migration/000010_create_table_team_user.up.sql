create table if not exists team_user (
	user_id BIGINT NOT NULL,
	team_id BIGINT NOT NULL,

  CONSTRAINT unique_user_team UNIQUE (user_id, team_id),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_team FOREIGN KEY (team_id) REFERENCES teams(id)
);