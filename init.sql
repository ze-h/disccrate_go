CREATE TABLE users (
    username CHAR(16),
    password CHAR(32),
    PRIMARY KEY (username)
);

CREATE TABLE records (
    title VARCHAR(64),
    artist VARCHAR(64),
    medium CHAR(8),
    format CHAR(8),
    label VARCHAR(64),
    genre VARCHAR(32),
    year YEAR,
    upc INT(64),
    username CHAR(16),
    PRIMARY KEY (title, artist, medium, upc),
    FOREIGN KEY (username) REFERENCES users(username)
);

INSERT INTO users (username, password) VALUES ('jimmy', '5f4dcc3b5aa765d61d8327deb882cf99'); -- password
INSERT INTO users (username, password) VALUES ('james', '47b7bfb65fa83ac9a71dcb0f6296bb6e'); -- Passw0rd!
INSERT INTO users (username, password) VALUES ('jordan', 'f830f69d23b8224b512a0dc2f5aec974'); -- thisisatest

INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES ('Our Love To Admire', 'Interpol', 'CD', 'Album', 'Matador', 'Rock', 2007, 0, 'jimmy');
INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES ('INNERMOST THOUGHT', 'Technical Itch', 'CD', 'DJ Mix', 'Tech Itch Recordings', 'Electronic', 2023, 0, 'james');
INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES ('Add Violence', 'Nine Inch Nails', 'CD', 'EP', 'Null Recordings', 'Rock', 2017, 0, 'jimmy');
INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES ('Parabola', 'Tool', 'DVD', 'Single', 'Volcano II', 'Metal', 2002, 0, 'jordan');
INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES ('Holy Roller', 'Spiritbox', 'Vinyl', 'Single', 'Pale Chord', 'Metal', 2021, 0, 'jordan');