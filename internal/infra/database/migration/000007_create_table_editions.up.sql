create table if not exists editions (
	id bigserial NOT NULL,
	winner_id BIGINT NULL DEFAULT NULL,
	name varchar(65) NULL DEFAULT NULL,
	description varchar(65) NULL DEFAULT NULL,
	status varchar(65) NULL DEFAULT NULL,
	start_date timestamptz NOT NULL,
	end_date timestamptz NOT NULL,

  CONSTRAINT edition_pkey PRIMARY KEY (id)
  );
