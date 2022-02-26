set client_encoding = 'UTF8';

CREATE TABLE users (
  id serial primary key,
  name varchar not null,
  age integer not null
);

INSERT INTO users(name, age) VALUES
  ('Kudo', 29),
  ('Yamada', 28),
  ('Tanaka', 27)
;
