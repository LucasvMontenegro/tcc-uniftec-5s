INSERT INTO credentials (id, account_id, email, password, created_at, updated_at) values (1, null, 'admin@admin.com', 'admin', now(), now());

INSERT INTO accounts (id, credential_id, user_id, email, created_at, updated_at) VALUES (1,1,null,'admin@admin.com', now(), now());

INSERT INTO users (id, account_id, name, is_admin, status, created_at, updated_at) values (1, 1, 'admin', true, 'ACTIVE', now(), now());

UPDATE credentials SET account_id = 1 WHERE id = 1;

UPDATE accounts SET user_id = 1 WHERE id = 1;