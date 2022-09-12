create table if not exists users (
	id bigserial NOT NULL,
	account_id BIGINT UNIQUE NOT NULL,
	name varchar(255) NOT NULL,
	is_admin boolean NOT NULL DEFAULT false,
	status varchar(65) NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,

  CONSTRAINT users_pkey PRIMARY KEY (id)
);
