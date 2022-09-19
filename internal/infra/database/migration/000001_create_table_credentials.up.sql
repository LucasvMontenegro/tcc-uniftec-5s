create table if not exists credentials (
	id bigserial NOT NULL,
	account_id BIGINT UNIQUE NULL DEFAULT NULL,
	email varchar(65) UNIQUE NOT NULL,
	password varchar(65) NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,

  CONSTRAINT credential_pkey PRIMARY KEY (id)
);
