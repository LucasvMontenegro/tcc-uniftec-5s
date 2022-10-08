create table if not exists prizes (
	id bigserial NOT NULL,
	edition_id BIGINT UNIQUE NOT NULL,
	name varchar(65) NULL DEFAULT NULL,
	description varchar(65) NULL DEFAULT NULL,

  CONSTRAINT prize_pkey PRIMARY KEY (id)
  );
