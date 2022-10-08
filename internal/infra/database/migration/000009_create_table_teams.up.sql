create table if not exists teams (
	id bigserial NOT NULL,
	edition_id BIGINT NOT NULL,
	name varchar(65) NULL DEFAULT NULL,

  CONSTRAINT team_pkey PRIMARY KEY (id)
  );
