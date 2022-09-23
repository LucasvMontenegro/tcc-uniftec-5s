DROP TYPE IF EXISTS user_status;
CREATE TYPE user_status AS ENUM ('ACTIVE', 'INACTIVE');

create table if not exists users (
	id bigserial NOT NULL,
	account_id BIGINT UNIQUE NOT NULL,
	name varchar(255) NOT NULL,
	is_admin boolean NOT NULL DEFAULT false,
	status user_status NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,

  CONSTRAINT users_pkey PRIMARY KEY (id)
);
