CREATE TABLE games (
    id serial NOT NULL PRIMARY KEY,
    title varchar(60) NOT NULL,
    genre varchar(9) NOT NULL,
    console varchar(16) NOT NULL,
    file_url varchar(29) NOT NULL,
    image_url varchar(29) NOT NULL,
    sorted integer NOT NULL,
    active boolean NOT NULL,
    lastSortedAt timestamp with time zone
)