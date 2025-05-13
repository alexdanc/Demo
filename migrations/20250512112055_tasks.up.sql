CREATE TABLE RequestBody (
                       id SERIAL PRIMARY KEY,
                       task VARCHAR(255) NOT NULL,
                       is_done BOOLEAN DEFAULT FALSE,

);