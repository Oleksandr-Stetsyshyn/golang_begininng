CREATE DATABASE IF NOT EXISTS librarydb;

CREATE TABLE books (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255),
    is_borrowed BOOLEAN,
    borrower_id INT UNSIGNED, -- change this from int to UNSIGNED INT
    PRIMARY KEY (id)
);

CREATE TABLE managers (
    id int NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    PRIMARY KEY (id)
);

