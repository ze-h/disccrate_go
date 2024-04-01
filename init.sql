CREATE TABLE records (
    title VARCHAR(50),
    artist VARCHAR(50),
    medium CHAR(10),
    format CHAR(10),
    label VARCHAR(50),
    genre VARCHAR(25),
    year YEAR,
    upc INT(64),
    PRIMARY KEY (title, artist, medium, upc)
);

CREATE TABLE users (
    username CHAR(15),
    password CHAR(16),
    PRIMARY KEY (username)
);

INSERT INTO records (title, artist, medium, format, label, genre, year, upc) VALUES ('Our Love To Admire', 'Interpol', 'CD', 'Album', 'Matador', 'Rock', 2007, 0);
INSERT INTO records (title, artist, medium, format, label, genre, year, upc) VALUES ('INNERMOST THOUGHT', 'Technical Itch', 'CD', 'DJ Mix', 'Tech Itch Recordings', 'Electronic', 2023, 0);
INSERT INTO records (title, artist, medium, format, label, genre, year, upc) VALUES ('Add Violence', 'Nine Inch Nails', 'CD', 'EP', 'Null Recordings', 'Rock', 2017, 0);
INSERT INTO records (title, artist, medium, format, label, genre, year, upc) VALUES ('Parabola', 'Tool', 'DVD', 'Single', 'Volcano II', 'Metal', 2002, 0);
INSERT INTO records (title, artist, medium, format, label, genre, year, upc) VALUES ('Holy Roller', 'Spiritbox', 'Vinyl', 'Single', 'Pale Chord', 'Metal', 2021, 0);

INSERT INTO users (username, password) VALUES ('jimmy', '5f4dcc3b5aa765d61d8327deb882cf99'); -- password
INSERT INTO users (username, password) VALUES ('james', '47b7bfb65fa83ac9a71dcb0f6296bb6e'); -- Passw0rd!
INSERT INTO users (username, password) VALUES ('jordan', 'f830f69d23b8224b512a0dc2f5aec974'); -- thisisatest