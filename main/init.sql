CREATE TABLE books
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    Title VARCHAR(50) NOT NULL,
    Author VARCHAR(50) NOT NULL,
    Year DATE NOT NULL 
);