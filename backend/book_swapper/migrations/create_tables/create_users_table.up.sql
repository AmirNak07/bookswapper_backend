CREATE TABLE IF NOT EXISTS users
(
    id           SERIAL PRIMARY KEY,

    username     VARCHAR(25) UNIQUE NOT NULL,
    fullname     VARCHAR(50)  default 'Guest',

    password     VARCHAR(50)        NOT NULL,

    city_id      INTEGER,
    FOREIGN KEY (city_id) REFERENCES cities (id),
    biography    VARCHAR(200) DEFAULT '',
    avatar       TEXT         DEFAULT '',
    phone_number VARCHAR(15)  DEFAULT '',

    auth_uuid VARCHAR(36) NOT NULL
);