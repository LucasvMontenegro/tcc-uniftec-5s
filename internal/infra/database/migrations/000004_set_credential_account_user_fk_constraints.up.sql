alter table credentials
  ADD CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES accounts(id);

alter table accounts
 ADD CONSTRAINT fk_credential FOREIGN KEY (credential_id) REFERENCES credentials(id),
 ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id);

alter table users
 ADD CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES accounts(id);