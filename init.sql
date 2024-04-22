CREATE TABLE users (
    username CHAR(16),
    password VARCHAR(64),
    PRIMARY KEY (username)
);

CREATE TABLE records (
    title VARCHAR(64),
    artist VARCHAR(64),
    medium CHAR(8),
    format CHAR(16),
    label VARCHAR(64),
    genre VARCHAR(32),
    year YEAR,
    upc VARCHAR(32),
    username CHAR(16),
    PRIMARY KEY (upc, username),
    FOREIGN KEY (username) REFERENCES users(username)
);

INSERT INTO users (username, password) VALUES ('jimmy', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8'); -- password
INSERT INTO users (username, password) VALUES ('jordan', 'a7c96262c21db9a06fd49e307d694fd95f624569f9b35bb3ffacd880440f9787'); -- thisisatest

INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES ('Our Love To Admire', 'Interpol', 'CD', 'Album', 'Matador', 'Rock', 2007, '0094639624829', 'jimmy');
INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES ('Add Violence', 'Nine Inch Nails', 'CD', 'EP', 'Null Recordings', 'Rock', 2017, '602557897975', 'jimmy');
INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES ('Parabola', 'Tool', 'DVD', 'Single', 'Volcano II', 'Metal', 2002, '828765759199', 'jordan');
INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES ('Holy Roller', 'Spiritbox', 'Vinyl', 'Single', 'Pale Chord', 'Metal', 2021, '4050538642933', 'jordan');