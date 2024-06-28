CREATE TABLE IF NOT EXISTS application_users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    gender     VARCHAR(1) NOT NULL,
    age        INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS swipes
(
    from_id    INT NOT NULL,
    to_id      INT NOT NULL,
    accept     BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (from_id, to_id)
);

CREATE TABLE IF NOT EXISTS matches
(
    user1_id   INT NOT NULL,
    user2_id   INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user1_id, user2_id)
);