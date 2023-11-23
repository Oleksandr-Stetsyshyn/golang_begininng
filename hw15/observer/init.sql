CREATE DATABASE IF NOT EXISTS playersdb;

CREATE TABLE game_rooms (
    id int NOT NULL AUTO_INCREMENT,
    level_name varchar(30),
    PRIMARY KEY (id)
);

CREATE TABLE players (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(30)

    room_id int,
    PRIMARY KEY (id),
    FOREIGN KEY (room_id) REFERENCES game_rooms(id)
);
