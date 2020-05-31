CREATE TABLE IF NOT EXISTS reports (
  id bigserial NOT NULL,
  user_id bigint NOT NULL,
  message varchar(1024) NOT NULL,
  date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)