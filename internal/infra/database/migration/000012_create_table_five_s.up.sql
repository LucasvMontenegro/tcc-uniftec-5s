CREATE TYPE five_s_names AS ENUM ('SEIRI', 'SEITON', 'SEISO', 'SEIKETSU', 'SHITSUKE');

create table if not exists five_s (
	id bigserial NOT NULL,
	name five_s_names UNIQUE NOT NULL,
	description varchar(255) NULL DEFAULT NULL,

	CONSTRAINT five_s_pkey PRIMARY KEY (id)
);
