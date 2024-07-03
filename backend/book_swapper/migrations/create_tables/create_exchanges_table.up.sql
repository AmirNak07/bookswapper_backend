CREATE TABLE IF NOT EXISTS exchanges
(
    id          SERIAL PRIMARY KEY,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users (id),
    book_name   VARCHAR(100),
    description VARCHAR(250) DEFAULT '',
    author VARCHAR(75) DEFAULT '',
    date        DATE NOT NULL,
    time        TIME NOT NULL
);
