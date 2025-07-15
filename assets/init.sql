

CREATE TABLE IF NOT EXISTS weather (
    city TEXT NOT NULL,
    temperature INT NOT NULL
);


INSERT INTO
    weather (city, temperature)
VALUES
    ('London', 15),
    ('Paris', 18),
    ('Tokyo', 25),
    ('Sydney', 20),
    ('Berlin', 16),
    ('Moscow', 12),
    ('Rome', 21),
    ('Madrid', 23),
    ('Dubai', 35),
    ('Singapore', 29);