create table if not exists session_history (
	id bigserial NOT NULL,
	account_id BIGINT UNIQUE NOT NULL,
	created_at timestamptz NOT NULL,
	expires_at timestamptz NULL DEFAULT NULL,

  CONSTRAINT session_history_pkey PRIMARY KEY (id),
  CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES accounts(id)
);
