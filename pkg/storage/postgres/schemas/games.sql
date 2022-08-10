CREATE TABLE games(
  id SERIAL  NOT NULL PRIMARY KEY,
  title     VARCHAR(60) NOT NULL,
  genre     VARCHAR(9) NOT NULL,
  console   VARCHAR(16) NOT NULL,
  file_url  VARCHAR(29) NOT NULL,
  image_url VARCHAR(29) NOT NULL,
  sorted    INTEGER  NOT NULL,
  active    BOOLEAN NOT NULL
);