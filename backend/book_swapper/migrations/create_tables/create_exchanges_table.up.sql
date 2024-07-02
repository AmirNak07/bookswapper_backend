CREATE TABLE IF NOT EXISTS exchanges
(
    id          SERIAL PRIMARY KEY,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id),
    book_name   VARCHAR(75),
    description VARCHAR(250) DEFAULT '',
    date        DATE NOT NULL,
    time        TIME NOT NULL
);
