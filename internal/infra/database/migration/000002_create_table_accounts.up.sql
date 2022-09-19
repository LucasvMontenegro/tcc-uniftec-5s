create table if not exists accounts (
	id bigserial NOT NULL,
	credential_id BIGINT UNIQUE NOT NULL,
	user_id BIGINT UNIQUE,
	email varchar(65) UNIQUE NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,

  CONSTRAINT accounts_pkey PRIMARY KEY (id)
);
