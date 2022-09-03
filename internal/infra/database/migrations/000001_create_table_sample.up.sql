create table if not exists samples (
	id bigserial NOT NULL,
	reference_uuid uuid  UNIQUE NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL DEFAULT NULL,

  CONSTRAINT sample_pkey PRIMARY KEY (id)
);
