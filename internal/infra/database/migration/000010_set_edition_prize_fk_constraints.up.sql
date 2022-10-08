alter table prizes
  ADD CONSTRAINT fk_edition FOREIGN KEY (edition_id) REFERENCES editions(id);

alter table editions
  ADD CONSTRAINT fk_team FOREIGN KEY (winner_id) REFERENCES teams(id);

alter table teams
  ADD CONSTRAINT fk_edition FOREIGN KEY (edition_id) REFERENCES editions(id);
